package fn_array

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"

	"github.com/colegno/arraylist"
)

type AccesoArray struct {
	Identificador interface{}
	Posiciones    *arraylist.List
	Tipo          Ast.TipoDato
	Fila          int
	Columna       int
}

func NewAccesoArray(id interface{}, posiciones *arraylist.List, fila, columna int) AccesoArray {
	nA := AccesoArray{
		Identificador: id,
		Tipo:          Ast.ACCESO_ARRAY,
		Posiciones:    posiciones,
		Fila:          fila,
		Columna:       columna,
	}
	return nA
}

func (p AccesoArray) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	var simbolo Ast.Simbolo
	var array interface{}
	var resultado Ast.TipoRetornado
	var id string
	var posicion interface{}
	var copiaLista *arraylist.List
	var valorPosicion Ast.TipoRetornado
	posiciones := arraylist.New()
	//Primero verificar que sea un identificador el id
	_, tipoParticular := p.Identificador.(Ast.Abstracto).GetTipo()
	if tipoParticular != Ast.IDENTIFICADOR {
		return p.AccesoPorExpresion(scope).(Ast.TipoRetornado)
		//Error se espera un identificador
		/*
			msg := "Semantic error, expected IDENTIFICADOR, found. " + Ast.ValorTipoDato[tipoParticular] +
				". -- Line: " + strconv.Itoa(p.Fila) +
				" Column: " + strconv.Itoa(p.Columna)
			nError := errores.NewError(p.Fila, p.Columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			scope.Errores.Add(nError)
			scope.Consola += msg + "\n"
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}
		*/
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
	if simbolo.Tipo != Ast.ARRAY && simbolo.Tipo != Ast.VECTOR {
		msg := "Semantic error, expected " + Ast.ValorTipoDato[simbolo.Tipo] + ", found " + Ast.ValorTipoDato[simbolo.Tipo] + "." +
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
	if simbolo.Tipo == Ast.ARRAY {
		array = simbolo.Valor.(Ast.TipoRetornado).Valor.(expresiones.Array)
	} else {
		array = simbolo.Valor.(Ast.TipoRetornado).Valor.(expresiones.Vector)
	}

	//Get las posiciones
	for i := 0; i < p.Posiciones.Len(); i++ {
		posicion = p.Posiciones.GetValue(i)
		valorPosicion = posicion.(Ast.Expresion).GetValue(scope)
		_, tipoParticular := posicion.(Ast.Abstracto).GetTipo()
		if valorPosicion.Tipo == Ast.ERROR {
			return posicion.(Ast.TipoRetornado)
		}
		//Verificar que el número en el acceso sea usize
		resultado := expresiones.EsUsize(valorPosicion, tipoParticular, posicion, scope)
		if resultado.Tipo == Ast.ERROR {
			return resultado
		}
		posiciones.Add(valorPosicion.Valor)
	}

	//Buscar la posición
	//Copiar lista
	copiaLista = p.Posiciones.Clone()
	if simbolo.Tipo == Ast.ARRAY {
		resultado = GetElemento(array.(expresiones.Array), copiaLista, posiciones, scope)
	} else {
		resultado = GetElementoVector(array.(expresiones.Vector), copiaLista, posiciones, scope)
	}

	return resultado
}

func (p AccesoArray) AccesoPorExpresion(scope *Ast.Scope) interface{} {
	var resultado Ast.TipoRetornado
	var posicion interface{}
	var valorPosicion Ast.TipoRetornado
	var copiaLista *arraylist.List
	posiciones := arraylist.New()
	//Conseguir el array
	array := p.Identificador.(Ast.Expresion).GetValue(scope)
	//Verificar que sea un array
	if array.Tipo != Ast.ARRAY {
		fila := p.Identificador.(Ast.Abstracto).GetFila()
		columna := p.Identificador.(Ast.Abstracto).GetColumna()
		msg := "Semantic error, expected ARRAY, found " + Ast.ValorTipoDato[array.Tipo] + "." +
			" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
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

	//Get las posiciones
	for i := 0; i < p.Posiciones.Len(); i++ {
		posicion = p.Posiciones.GetValue(i)
		valorPosicion = posicion.(Ast.Expresion).GetValue(scope)
		_, tipoParticular := posicion.(Ast.Abstracto).GetTipo()
		if valorPosicion.Tipo == Ast.ERROR {
			return posicion.(Ast.TipoRetornado)
		}
		//Verificar que el número en el acceso sea usize
		resultado := expresiones.EsUsize(valorPosicion, tipoParticular, posicion, scope)
		if resultado.Tipo == Ast.ERROR {
			return resultado
		}
		posiciones.Add(valorPosicion.Valor)
	}

	//Buscar la posición
	//Copiar lista
	copiaLista = p.Posiciones.Clone()
	resultado = GetElemento(array.Valor.(expresiones.Array), copiaLista, posiciones, scope)
	return resultado
}

func (v AccesoArray) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, v.Tipo
}

func (v AccesoArray) GetFila() int {
	return v.Fila
}
func (v AccesoArray) GetColumna() int {
	return v.Columna
}

//p.Posiciones
func GetElemento(array expresiones.Array, elementos *arraylist.List, posiciones *arraylist.List, scope *Ast.Scope) Ast.TipoRetornado {
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
		//Si es 0, entonces retornar la posición actual
		return array.Elementos.GetValue(posicion).(Ast.TipoRetornado)
	}
	if posiciones.Len() > 0 && array.TipoArray != Ast.ARRAY {
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
	next := array.Elementos.GetValue(posicion).(Ast.TipoRetornado)

	//Validar que el siguiente sea un array y que todavía existan posiciones que buscar
	return GetElemento(next.Valor.(expresiones.Array), elementos, posiciones, scope)

}

//Get elemento vector

func GetElementoVector(array expresiones.Vector, elementos *arraylist.List, posiciones *arraylist.List, scope *Ast.Scope) Ast.TipoRetornado {
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
		//Si es 0, entonces retornar la posición actual
		return array.Valor.GetValue(posicion).(Ast.TipoRetornado)
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

	//Validar que el siguiente sea un array y que todavía existan posiciones que buscar
	return GetElementoVector(next.Valor.(expresiones.Vector), elementos, posiciones, scope)

}

func UpdateElemento(array expresiones.Array, elementos *arraylist.List, posiciones *arraylist.List, scope *Ast.Scope, objeto interface{}) Ast.TipoRetornado {
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
		//Actualizar el elemento simpre y cuando no este en un tipo array en la siguiente columna
		next := array.Elementos.GetValue(posicion).(Ast.TipoRetornado)
		if next.Tipo == Ast.ARRAY {
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
		for i := 0; i < array.Elementos.Len(); i++ {
			if i == posicion {
				//Reemplazar el elemento
				nuevaLista.Add(objeto)
				continue
			}
			nuevaLista.Add(array.Elementos.GetValue(i))
		}
		array.Elementos.Clear()
		for i := 0; i < nuevaLista.Len(); i++ {
			array.Elementos.Add(nuevaLista.GetValue(i))
		}
		return Ast.TipoRetornado{Valor: true, Tipo: Ast.EJECUTADO}
	}
	if posiciones.Len() > 0 && array.TipoArray != Ast.ARRAY {
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
	next := array.Elementos.GetValue(posicion).(Ast.TipoRetornado)
	valorNext := next.Valor.(expresiones.Array)
	//Validar que el siguiente sea un array y que todavía existan posiciones que buscar
	return UpdateElemento(valorNext, elementos, posiciones, scope, objeto)

}
