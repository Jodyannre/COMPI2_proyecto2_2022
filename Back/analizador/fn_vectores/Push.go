package fn_vectores

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"
)

type Push struct {
	Identificador interface{}
	Valor         interface{}
	Tipo          Ast.TipoDato
	Fila          int
	Columna       int
}

func NewPush(id interface{}, valor interface{}, tipo Ast.TipoDato, fila, columna int) Push {
	nP := Push{
		Valor:         valor,
		Identificador: id,
		Tipo:          tipo,
		Fila:          fila,
		Columna:       columna,
	}
	return nP
}

func (p Push) Run(scope *Ast.Scope) interface{} {
	var simbolo Ast.Simbolo
	var vector expresiones.Vector
	var valor Ast.TipoRetornado
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
			" in a not mutable VECTOR<" + expresiones.Tipo_String(vector.TipoVector) + ">." +
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

	//Paso todas las pruebas, entonces guardar el elemento
	vector.Valor.Add(valor)
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

func (v Push) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, v.Tipo
}

func (v Push) GetFila() int {
	return v.Fila
}
func (v Push) GetColumna() int {
	return v.Columna
}
