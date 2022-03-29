package fn_vectores

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"

	"github.com/colegno/arraylist"
)

type InsertVec struct {
	Identificador interface{}
	Posicion      interface{}
	Valor         interface{}
	Tipo          Ast.TipoDato
	Fila          int
	Columna       int
}

func NewInsertVec(id interface{}, valor interface{}, posicion interface{}, tipo Ast.TipoDato, fila, columna int) InsertVec {
	nP := InsertVec{
		Valor:         valor,
		Identificador: id,
		Posicion:      posicion,
		Tipo:          tipo,
		Fila:          fila,
		Columna:       columna,
	}
	return nP
}

func (p InsertVec) Run(scope *Ast.Scope) interface{} {
	var simbolo Ast.Simbolo
	var vector expresiones.Vector
	var valor Ast.TipoRetornado
	var posicion Ast.TipoRetornado
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

	//Verificar que el elemento que se va a agregar sea del mismo tipo que el que guarda el vector
	//Primero calcular el valor
	valor = p.Valor.(Ast.Expresion).GetValue(scope)
	if valor.Tipo == Ast.ERROR {
		return valor
	}

	//Verificar que el vector sea mutable
	if !simbolo.Mutable {
		msg := "Semantic error, can't store " + Ast.ValorTipoDato[valor.Tipo] + " value" +
			" in a not mutable Vec<" + Ast.ValorTipoDato[vector.Tipo] + ">." +
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

	if valor.Tipo != vector.TipoVector.Tipo {
		//Error de tipos dentro del vector
		msg := "Semantic error, can't store " + Ast.ValorTipoDato[valor.Tipo] + " value" +
			" in a Vec<" + Ast.ValorTipoDato[vector.Tipo] + ">." +
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

	//Verificar si es vector el que se va a agregar y el tipo del vector
	if valor.Tipo == Ast.VECTOR {
		if !expresiones.CompararTipos(valor.Valor.(expresiones.Vector).TipoVector, vector.TipoVector) {
			//Error, no se puede guardar ese tipo de vector en este vector
			msg := "Semantic error, can't store " + expresiones.Tipo_String(valor.Valor.(expresiones.Vector).TipoVector) + " value" +
				" in a VEC< " + expresiones.Tipo_String(vector.TipoVector) + ">." +
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
	}
	//Verificar si es un struct el que se va a agregar
	if valor.Tipo == Ast.STRUCT {
		plantilla := valor.Valor.(Ast.Structs).GetPlantilla(scope)
		tipoStruct := Ast.TipoRetornado{Valor: plantilla, Tipo: Ast.STRUCT}
		if !expresiones.CompararTipos(tipoStruct, vector.TipoVector) {
			//Error, no se puede guardar ese tipo de vector en este vector
			msg := "Semantic error, can't store " + plantilla + " value" +
				" in a VEC< " + expresiones.Tipo_String(vector.TipoVector) + ">." +
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
	}

	//Get la posición en donde se quiere agregar el nuevo valor
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
	//Crear la nueva lista que contendrá los valores
	nLista := arraylist.New()

	for i := 0; i <= vector.Valor.Len(); i++ {
		if i == posicion.Valor.(int) {
			nLista.Add(valor)
		}
		if i < vector.Valor.Len() {
			nLista.Add(vector.Valor.GetValue(i))
		}
	}
	//Agregar la nueva lista
	vector.Valor = nil
	vector.Valor = nLista
	//Aumentar el tamaño
	vector.Size++
	vector.Capacity = vector.CalcularCapacity(vector.Size, vector.Capacity)
	//Cambiar el estado de vacio
	if vector.Vacio {
		vector.Vacio = false
	}
	simbolo.Valor = Ast.TipoRetornado{Tipo: Ast.VECTOR, Valor: vector}
	scope.UpdateSimbolo(id, simbolo)
	return Ast.TipoRetornado{
		Tipo:  Ast.EJECUTADO,
		Valor: true,
	}
}

func (v InsertVec) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, v.Tipo
}

func (v InsertVec) GetFila() int {
	return v.Fila
}
func (v InsertVec) GetColumna() int {
	return v.Columna
}
