package fn_vectores

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"

	"github.com/colegno/arraylist"
)

type RemoveVec struct {
	Identificador interface{}
	Posicion      interface{}
	Tipo          Ast.TipoDato
	Fila          int
	Columna       int
}

func NewRemoveVec(id interface{}, posicion interface{}, tipo Ast.TipoDato, fila, columna int) RemoveVec {
	nP := RemoveVec{
		Identificador: id,
		Posicion:      posicion,
		Tipo:          tipo,
		Fila:          fila,
		Columna:       columna,
	}
	return nP
}

func (p RemoveVec) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	var simbolo Ast.Simbolo
	var vector expresiones.Vector
	var posicion Ast.TipoRetornado
	var removido Ast.TipoRetornado
	var id string
	//Primero verificar que sea un identificador el id
	_, tipoParticular := p.Identificador.(Ast.Abstracto).GetTipo()
	if tipoParticular != Ast.IDENTIFICADOR {
		//Error se espera un identificador
		msg := "Semantic error, expected IDENTIFICADOR, found. " + Ast.ValorTipoDato[tipoParticular] +
			". -- Line: " + strconv.Itoa(p.Fila) +
			" Column: " + strconv.Itoa(p.Columna)
		nError := errores.NewError(p.Fila, p.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}
	//Recuperar el id del identificador
	id = p.Identificador.(expresiones.Identificador).Valor

	//Verificar que el id exista
	if !scope.Exist(id) {
		//Error la variable no existe
		msg := "Semantic error, the element \"" + id + "\" doesn't exist in any scope." +
			" -- Line:" + strconv.Itoa(p.Fila) + " Column: " + strconv.Itoa(p.Columna)
		nError := errores.NewError(p.Fila, p.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}
	//Conseguir el simbolo y el vector
	simbolo = scope.GetSimbolo(id)
	//Verificar que sea un vector
	if simbolo.Tipo != Ast.VECTOR {
		msg := "Semantic error, expected Vector, found " + Ast.ValorTipoDato[simbolo.Tipo] + "." +
			" -- Line:" + strconv.Itoa(p.Fila) + " Column: " + strconv.Itoa(p.Columna)
		nError := errores.NewError(p.Fila, p.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}
	vector = simbolo.Valor.(Ast.TipoRetornado).Valor.(expresiones.Vector)

	//Verificar que el vector sea mutable
	if !simbolo.Mutable {
		msg := "Semantic error, can't remove an elment from an inmutable Vector<" +
			expresiones.Tipo_String(vector.TipoVector) + ">." +
			" -- Line: " + strconv.Itoa(p.Fila) +
			" Column: " + strconv.Itoa(p.Columna)
		nError := errores.NewError(p.Fila, p.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}

	//Get la posición de donde se va a extraer el elemento
	posicion = p.Posicion.(Ast.Expresion).GetValue(scope)
	_, tipoParticular = p.Posicion.(Ast.Abstracto).GetTipo()
	if posicion.Tipo == Ast.ERROR {
		return posicion
	}
	//Verificar que el número en el acceso sea usize
	if (posicion.Tipo != Ast.USIZE && posicion.Tipo != Ast.I64) ||
		tipoParticular == Ast.IDENTIFICADOR && posicion.Tipo == Ast.I64 {
		//Error, se espera un usize
		fila := p.Posicion.(Ast.Abstracto).GetFila()
		columna := p.Posicion.(Ast.Abstracto).GetColumna()
		msg := "Semantic error, expected USIZE, found. " + Ast.ValorTipoDato[posicion.Tipo] +
			". -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
		nError := errores.NewError(fila, columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}
	//Verificar que la posición exista en el vector
	if posicion.Valor.(int) > vector.Size {
		//Error, fuera de rango
		fila := p.Posicion.(Ast.Abstracto).GetFila()
		columna := p.Posicion.(Ast.Abstracto).GetColumna()
		msg := "Semantic error, index (" + strconv.Itoa(posicion.Valor.(int)) + ") out of bounds." +
			". -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
		nError := errores.NewError(fila, columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}

	//Paso todas las pruebas, entonces guardar el elemento
	//Crear la nueva lista que contendrá los valores después de la eliminación
	nLista := arraylist.New()

	for i := 0; i <= vector.Valor.Len(); i++ {
		if i == posicion.Valor.(int) {
			removido = vector.Valor.GetValue(i).(Ast.TipoRetornado)
			continue
		}
		if i < vector.Valor.Len() {
			nLista.Add(vector.Valor.GetValue(i))
		}
	}
	//Limpiar la lista
	vector.Valor.Clear()

	//Regresar los valores a la lista
	for i := 0; i <= nLista.Len(); i++ {
		vector.Valor.Add(nLista.GetValue(i))
	}

	//Agregar la nueva lista
	vector.Valor = nil
	vector.Valor = nLista
	//Disminuir el tamaño
	vector.Size--
	//vector.Capacity = vector.CalcularCapacity(vector.Size, vector.Capacity)
	//Cambiar el estado de vacio
	if vector.Size == 0 {
		vector.Vacio = true
	}
	simbolo.Valor = Ast.TipoRetornado{Tipo: Ast.VECTOR, Valor: vector}
	scope.UpdateSimbolo(id, simbolo)
	return removido
}

func (v RemoveVec) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, v.Tipo
}

func (v RemoveVec) GetFila() int {
	return v.Fila
}
func (v RemoveVec) GetColumna() int {
	return v.Columna
}
