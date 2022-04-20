package simbolos

import (
	"Back/analizador/Ast"
	"strings"

	"github.com/colegno/arraylist"
)

type FuncionMain struct {
	Tipo          Ast.TipoDato
	Fila          int
	Columna       int
	Instrucciones *arraylist.List
}

func NewFuncionMain(instrucciones *arraylist.List, fila, columna int) FuncionMain {
	nM := FuncionMain{
		Tipo:          Ast.FUNCION_MAIN,
		Fila:          fila,
		Columna:       columna,
		Instrucciones: instrucciones,
	}
	return nM
}

func (f FuncionMain) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	/************************************VARIABLES 3D ***************************************/
	var obj3dRespuestaInstruccion Ast.O3D
	var saltoReturn string
	var codigo3d string
	/****************************************************************************************/

	//Primero crear el nuevo scope main
	newScope := Ast.NewScope("Main", scope)
	//Le agrego su posición en el stack, reinicio el puntero p para simular el nuevo ambito
	newScope.Posicion = 0

	var actual interface{}
	var tipoGeneral interface{}
	var respuesta interface{}
	codigo3d += Ast.Indentar(newScope.GetNivel(), "P = 0; //Volver p a 0 para ejecutar el main \n")
	//Recorrer y ejecutar todas las instrucciones
	for i := 0; i < f.Instrucciones.Len(); i++ {
		actual = f.Instrucciones.GetValue(i)
		if actual != nil {
			tipoGeneral, _ = actual.(Ast.Abstracto).GetTipo()
		} else {
			continue
		}

		if tipoGeneral == Ast.INSTRUCCION {
			respuesta = actual.(Ast.Instruccion).Run(&newScope)
			obj3dRespuestaInstruccion = respuesta.(Ast.TipoRetornado).Valor.(Ast.O3D)
			//saltoBreak = obj3dRespuestaInstruccion.SaltoBreak
			//saltoContinue = obj3dRespuestaInstruccion.SaltoContinue
			saltoReturn = obj3dRespuestaInstruccion.SaltoReturn
			codigo3d += Ast.Indentar(newScope.GetNivel(), obj3dRespuestaInstruccion.Codigo)
		} else if tipoGeneral == Ast.EXPRESION {
			respuesta = actual.(Ast.Expresion).GetValue(&newScope)
			obj3dRespuestaInstruccion = respuesta.(Ast.TipoRetornado).Valor.(Ast.O3D)
			//saltoBreak = obj3dRespuestaInstruccion.SaltoBreak
			//saltoContinue = obj3dRespuestaInstruccion.SaltoContinue
			saltoReturn = obj3dRespuestaInstruccion.SaltoReturn
			codigo3d += Ast.Indentar(newScope.GetNivel(), obj3dRespuestaInstruccion.Codigo)
		}

		/*No es necesario verificar si trae error o es ejecutada, lo único que hay que verificar es
		Que traiga retornos que puede generar errores*/
		/*
			if Ast.EsTransferencia(respuesta.(Ast.TipoRetornado).Valor.(Ast.O3D).Valor.Tipo) {
				//Primero verificar que no sea un return normal, el cual si es permitido
				if respuesta.(Ast.TipoRetornado).Valor.(Ast.O3D).Valor.Tipo == Ast.RETURN {
					return respuesta.(Ast.TipoRetornado)
				}
				switch respuesta.(Ast.TipoRetornado).Valor.(Ast.O3D).Valor.Tipo {
				case Ast.BREAK, Ast.BREAK_EXPRESION, Ast.CONTINUE:
					/////////////////////////////ERROR/////////////////////////////////////
					errores.GenerarError(30, actual, actual, "", "", "", &newScope)
				case Ast.RETURN_EXPRESION:
					/////////////////////////////ERROR/////////////////////////////////////
					errores.GenerarError(31, actual, actual, "", "", "", &newScope)
				}
			}
		*/

		//Agregar el código que trae la respuesta al entorno global para luego imprimirlo en la web
		//newScope.Codigo += Ast.Indentar(newScope.GetNivel(), respuesta.(Ast.TipoRetornado).Valor.(Ast.O3D).Codigo)
	}

	if saltoReturn != "" {
		saltoReturn = strings.Replace(saltoReturn, ",", ":\n", -1)
		codigo3d += Ast.Indentar(newScope.GetNivel(), saltoReturn)
	}

	newScope.Codigo = codigo3d
	newScope.UpdateScopeGlobal()
	return Ast.TipoRetornado{
		Tipo:  Ast.EJECUTADO,
		Valor: true,
	}
}

func (v FuncionMain) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, v.Tipo
}

func (v FuncionMain) GetFila() int {
	return v.Fila
}
func (v FuncionMain) GetColumna() int {
	return v.Columna
}
