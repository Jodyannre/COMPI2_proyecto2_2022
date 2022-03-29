package fn_vectores

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"

	"github.com/colegno/arraylist"
)

type AccesoVec struct {
	Identificador interface{}
	Posicion      interface{}
	Tipo          Ast.TipoDato
	Fila          int
	Columna       int
}

func NewAccesoVec(id interface{}, posicion interface{}, tipo Ast.TipoDato, fila, columna int) AccesoVec {
	nA := AccesoVec{
		Identificador: id,
		Tipo:          tipo,
		Posicion:      posicion,
		Fila:          fila,
		Columna:       columna,
	}
	return nA
}

func (p AccesoVec) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	var simbolo Ast.Simbolo
	var vector interface{}
	var posicion Ast.TipoRetornado
	var resultado Ast.TipoRetornado
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
	if simbolo.Tipo != Ast.VECTOR && simbolo.Tipo != Ast.ARRAY {
		msg := "Semantic error, expected (VECTOR|ARRAY), found " + Ast.ValorTipoDato[simbolo.Tipo] + "." +
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
	if simbolo.Tipo == Ast.VECTOR {
		vector = simbolo.Valor.(Ast.TipoRetornado).Valor.(expresiones.Vector)
	} else {
		vector = simbolo.Valor.(Ast.TipoRetornado).Valor.(expresiones.Array)
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
	if simbolo.Tipo == Ast.VECTOR {
		if posicion.Valor.(int) >= vector.(expresiones.Vector).Size {
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

		//Acceder al elemento
		resultado = vector.(expresiones.Vector).Valor.GetValue(posicion.Valor.(int)).(Ast.TipoRetornado)
	}

	if simbolo.Tipo == Ast.ARRAY {
		if posicion.Valor.(int) >= vector.(expresiones.Array).Size {
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
		//Validar que el array no sea de más dimensiones
		/*
			if vector.(expresiones.Array).TipoArray == Ast.ARRAY {
				//Es de más dimensiones
				fila := p.Posicion.(Ast.Abstracto).GetFila()
				columna := p.Posicion.(Ast.Abstracto).GetColumna()
				msg := "Semantic error, you are trying to access a position in one dimension and this ARRAY has multiple dimensions." +
					". -- Line: " + strconv.Itoa(fila) +
					" Column: " + strconv.Itoa(columna)
				nError := errores.NewError(fila, columna, msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				scope.Errores.Add(nError)
				scope.Consola += msg + "\n"
				return Ast.TipoRetornado{
					Tipo:  Ast.ERROR,
					Valor: nError,
				}

			}
		*/
		//Acceder al elemento
		resultado = vector.(expresiones.Array).Elementos.GetValue(posicion.Valor.(int)).(Ast.TipoRetornado)
	}
	return resultado
}

func UpdateElemento(array expresiones.Vector, elementos *arraylist.List, posiciones *arraylist.List, scope *Ast.Scope, objeto interface{}) Ast.TipoRetornado {
	posicion := posiciones.GetValue(0).(int)
	elemento := elementos.GetValue(0)
	posiciones.RemoveAtIndex(0)
	elementos.RemoveAtIndex(0)
	if posicion >= array.Size || posicion < 0 {
		//Error, out of bounds
		fila := elemento.(Ast.Abstracto).GetFila()
		columna := elemento.(Ast.Abstracto).GetColumna()
		msg := "Semantic error, index (" + strconv.Itoa(posicion) + ") out of bounds." +
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
	if posiciones.Len() == 0 {
		//Actualizar el elemento siempre y cuando no sea una capa vector
		next := array.Valor.GetValue(posicion).(Ast.TipoRetornado)
		if next.Tipo == Ast.VECTOR {
			//No se puede guardar en esa posición porque es posición de array
			fila := elemento.(Ast.Abstracto).GetFila()
			columna := elemento.(Ast.Abstracto).GetColumna()
			msg := "Semantic error, index (" + strconv.Itoa(posicion) + "). Can't access to that position." +
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
		nuevaLista := *arraylist.New()
		for i := 0; i < array.Valor.Len(); i++ {
			if i == posicion {
				//Reemplazar el elemento
				nuevaLista.Add(objeto)
				continue
			}
			nuevaLista.Add(array.Valor.GetValue(i))
		}
		array.Valor.Clear()
		for i := 0; i < nuevaLista.Len(); i++ {
			array.Valor.Add(nuevaLista.GetValue(i))
		}
		return Ast.TipoRetornado{Valor: true, Tipo: Ast.EJECUTADO}
	}
	if posiciones.Len() > 0 && array.TipoVector.Tipo != Ast.VECTOR {
		//Error, no hay más dimensiones
		fila := elemento.(Ast.Abstracto).GetFila()
		columna := elemento.(Ast.Abstracto).GetColumna()
		msg := "Semantic error, index (" + strconv.Itoa(posicion) + ") out of bounds." +
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
	next := array.Valor.GetValue(posicion).(Ast.TipoRetornado)
	valorNext := next.Valor.(expresiones.Vector)
	//Validar que el siguiente sea un array y que todavía existan posiciones que buscar
	return UpdateElemento(valorNext, elementos, posiciones, scope, objeto)

}

func (v AccesoVec) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, v.Tipo
}

func (v AccesoVec) GetFila() int {
	return v.Fila
}
func (v AccesoVec) GetColumna() int {
	return v.Columna
}
