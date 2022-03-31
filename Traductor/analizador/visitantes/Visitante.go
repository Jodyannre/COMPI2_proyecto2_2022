package visitantes

import (
	"Traductor/analizador/ast"
	"Traductor/analizador/errores"
	"Traductor/parser"
	"fmt"
	"strconv"

	"github.com/colegno/arraylist"
)

type Visitador struct {
	*parser.BaseNparserListener
	Consola       string
	Errores       arraylist.List
	EntornoGlobal ast.Scope
}

func NewVisitor() *Visitador {
	return new(Visitador)
}

func (v *Visitador) ExitInicio(ctx *parser.InicioContext) {
	fmt.Println("Ingreso>>>>>>>>>>>>>>>")
	instrucciones := ctx.GetLista()
	//Verificar que no existan errores sintácticos o semánticos
	if v.Errores.Len() > 0 {
		EntornoGlobal := ast.NewScope("global", nil)
		EntornoGlobal.Global = true
		for i := 0; i < v.Errores.Len(); i++ {
			err := v.Errores.GetValue(i).(errores.CustomSyntaxError)
			nError := errores.NewError(err.Fila, err.Columna, err.Msg)
			nError.Tipo = ast.ERROR_SEMANTICO
			nError.Ambito = EntornoGlobal.GetTipoScope()
			EntornoGlobal.Errores.Add(nError)
		}
		//EntornoGlobal.GenerarTablaErrores()
		return
	}

	//Desde aquí ya tengo todo el resultado del parser, listo para ejecutar

	/*
		1) Declarar el scope
		2) Primera pasada para declarar todas los elementos y buscar el main o si hay más de 1 main
		3) Correr todas las instrucciones
	*/

	//Creando el scope global
	EntornoGlobal := ast.NewScope("global", nil)
	EntornoGlobal.Global = true
	//Variables para las declaraciones
	var actual interface{}
	var tipoParticular, tipoGeneral ast.TipoDato
	var metodoMain interface{}
	var mainEncontrado bool = false
	var contadorMain int = 0
	var posicionMain int = 0

	//Primera pasada para agregar todas las declaraciones de las variables y los elemenos
	//Buscar el main y verificar que solo exista uno o no ejecutar y error de que no hay main

	//Pasada de declaraciones
	for i := 0; i < instrucciones.Len(); i++ {
		actual = instrucciones.GetValue(i)
		if actual != nil {
			tipoGeneral, tipoParticular = actual.(ast.Abstracto).GetTipo()
		} else {
			continue
		}
		if tipoParticular == ast.DECLARACION {
			//Declarar variables globales
			actual.(ast.Instruccion).Run(&EntornoGlobal)
		} else if tipoGeneral == ast.EXPRESION {
			//Verificar que sea el main
			if tipoParticular == ast.FUNCION_MAIN {
				mainEncontrado = true
				contadorMain++
				posicionMain = i
			}
		}
	}

	//Verificar que se haya encontrado el main y que no sea más de 1 main
	if !mainEncontrado {
		//Error no hay metodo main
		fila := 0
		columna := 0
		msg := "Semantic error, MAIN method not found" +
			". -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
		nError := errores.NewError(fila, columna, msg)
		nError.Tipo = ast.ERROR_SEMANTICO
		nError.Ambito = EntornoGlobal.GetTipoScope()
		EntornoGlobal.Errores.Add(nError)
		EntornoGlobal.Consola += msg + "\n"
	}

	if contadorMain > 1 {
		//Error hay varios metodos main
		fila := 0
		columna := 0
		msg := "Semantic error, the program cannot have more than 1 MAIN method" +
			". -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
		nError := errores.NewError(fila, columna, msg)
		nError.Tipo = ast.ERROR_SEMANTICO
		nError.Ambito = EntornoGlobal.GetTipoScope()
		EntornoGlobal.Errores.Add(nError)
		EntornoGlobal.Consola += msg + "\n"
	}
	if mainEncontrado && contadorMain == 1 {
		//Get el metodo main
		metodoMain = instrucciones.GetValue(posicionMain)

		//Ejetuar el método main
		metodoMain.(ast.Expresion).GetValue(&EntornoGlobal)

	}

	EntornoGlobal.UpdateScopeGlobal()
	v.Consola += EntornoGlobal.Consola
	for i := 0; i < EntornoGlobal.Errores.Len(); i++ {
		v.Errores.Add(EntornoGlobal.Errores.GetValue(i))
	}
	//EntornoGlobal.GenerarTablaSimbolos()
	//EntornoGlobal.GenerarTablaErrores()
	//EntornoGlobal.GenerarTablaBD()
	//EntornoGlobal.GenerarTablaTablas()
}

func (v *Visitador) GetConsola() string {
	//return v.Consola
	return v.Consola
}

func (v *Visitador) UpdateConsola(entrada string) {
	v.Consola += entrada + "\n"
}
