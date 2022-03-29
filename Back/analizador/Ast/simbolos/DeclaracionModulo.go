package simbolos

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"
)

type DeclaracionModulo struct {
	Tipo    Ast.TipoDato
	Modulo  interface{}
	Fila    int
	Columna int
	Publico bool
	Entorno *Ast.Scope
}

func NewDeclaracionModulo(modulo interface{}, publico bool, fila, columna int) DeclaracionModulo {
	nD := DeclaracionModulo{
		Tipo:    Ast.DECLARACION,
		Modulo:  modulo,
		Fila:    fila,
		Columna: columna,
	}
	return nD
}

func (d DeclaracionModulo) Run(scope *Ast.Scope) interface{} {
	var existe bool
	var identificador interface{}
	var idString string
	var modulo Ast.TipoRetornado
	var newSimbolo Ast.Simbolo
	var tipoParticular Ast.TipoDato
	//Declarar el módulo en el scope actual

	//Verificar el id
	identificador = d.Modulo.(Modulo).Identificador
	_, tipoParticular = identificador.(Ast.Abstracto).GetTipo()

	if tipoParticular != Ast.IDENTIFICADOR {
		//Error, se espera un identificador
		fila := identificador.(Ast.Abstracto).GetFila()
		columna := identificador.(Ast.Abstracto).GetColumna()
		msg := "Semantic error, an IDENTIFICADOR was expected, found " + Ast.ValorTipoDato[tipoParticular] +
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
	} else {
		idString = identificador.(expresiones.Identificador).Valor
	}

	//Verificar que no exista el id
	existe = scope.Exist(idString)
	if existe {
		msg := "Semantic error, the element \"" + idString + "\" already exist." +
			" -- Line:" + strconv.Itoa(d.Fila) + " Column: " + strconv.Itoa(d.Columna)
		nError := errores.NewError(d.Fila, d.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}

	//Get el valor del modulo
	modulo = d.Modulo.(Ast.Expresion).GetValue(scope)

	//Crear el nuevo símbolo

	newSimbolo = Ast.Simbolo{
		Identificador: idString,
		Valor:         modulo,
		Fila:          d.Fila,
		Columna:       d.Columna,
		Tipo:          modulo.Tipo,
		Mutable:       false,
		Publico:       modulo.Valor.(Modulo).Publico,
		Entorno:       scope,
	}
	//Agregar el nuevo símbolo
	scope.Add(newSimbolo)
	//Agregar a la lista fms
	scope.Addfms(newSimbolo)

	return Ast.TipoRetornado{
		Tipo:  Ast.EJECUTADO,
		Valor: true,
	}
}

func (a DeclaracionModulo) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, a.Tipo
}

func (a DeclaracionModulo) GetFila() int {
	return a.Fila
}
func (a DeclaracionModulo) GetColumna() int {
	return a.Columna
}
