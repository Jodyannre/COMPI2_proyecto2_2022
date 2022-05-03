package simbolos

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"

	"github.com/colegno/arraylist"
)

type LlamadaFuncion struct {
	Identificador Ast.Expresion
	Parametros    *arraylist.List
	Tipo          Ast.TipoDato
	Fila          int
	ScopeOriginal *Ast.Scope
	Columna       int
}

func NewLlamadaFuncion(id Ast.Expresion, parametros *arraylist.List, tipo Ast.TipoDato, fila, columna int) LlamadaFuncion {
	nF := LlamadaFuncion{
		Identificador: id,
		Parametros:    parametros,
		Tipo:          tipo,
		Fila:          fila,
		Columna:       columna,
	}
	return nF
}

func (l LlamadaFuncion) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	/**************************VARIABLES 3D***************************/
	var codigo3d string
	var obj3dValor Ast.O3D
	var obj3dTemp Ast.O3D
	var codigoTemp string
	var contadorDeclaraciones int
	/*****************************************************************/

	var simbolo Ast.Simbolo
	var funcion Funcion
	var parametrosCreados Ast.TipoRetornado
	var resultadoFuncion Ast.TipoRetornado
	//Crear el scope para la nueva función
	newScope := Ast.NewScope("funcion", scope)

	//Verificar que la función existe en el ámbilo global
	simbolo = newScope.Exist_fms_local(l.Identificador.(expresiones.Identificador).Valor)
	//Sino verificar que exista en el local

	if simbolo.Tipo != Ast.FUNCION ||
		simbolo.Tipo == Ast.ERROR_ACCESO_PRIVADO ||
		simbolo.Tipo == Ast.ERROR_NO_EXISTE {
		//La función no existe en el scope global, puede que exista en un módulo
		//simbolo = newScope.Exist_fms(l.Identificador.(Identificador).Valor)

		//Verificar si el símbolo es privado
		if simbolo.Tipo == Ast.ERROR_ACCESO_PRIVADO {
			//Error el símbolo tiene acceso privado
			msg := "Semantic error, the function is private." +
				" -- Line: " + strconv.Itoa(l.Fila) +
				" Column: " + strconv.Itoa(l.Columna)
			nError := errores.NewError(l.Fila, l.Columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			nError.Ambito = scope.GetTipoScope()
			newScope.Errores.Add(nError)
			newScope.Consola += msg + "\n"
			//newScope.UpdateScopeGlobal()
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}

		}

		//Verificar si el símbolo existe
		if simbolo.Tipo == Ast.ERROR_NO_EXISTE {
			//Error el símbolo no existe
			msg := "Semantic error, the function doesn't exist." +
				" -- Line: " + strconv.Itoa(l.Fila) +
				" Column: " + strconv.Itoa(l.Columna)
			nError := errores.NewError(l.Fila, l.Columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			nError.Ambito = scope.GetTipoScope()
			newScope.Errores.Add(nError)
			newScope.Consola += msg + "\n"
			//newScope.UpdateScopeGlobal()
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}
		}

		//Verificar que el símbolo sea una función
		if simbolo.Tipo != Ast.FUNCION {
			//Error, el símbolo no es una función
			msg := "Semantic error, " + l.Identificador.(expresiones.Identificador).Valor + " is not a function." +
				" -- Line: " + strconv.Itoa(l.Fila) +
				" Column: " + strconv.Itoa(l.Columna)
			nError := errores.NewError(l.Fila, l.Columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			nError.Ambito = scope.GetTipoScope()
			newScope.Errores.Add(nError)
			newScope.Consola += msg + "\n"
			//newScope.UpdateScopeGlobal()
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}
		}
	} else {
		//simbolo = newScope.GetSimbolo(l.Identificador.(Identificador).Valor)
		funcion = simbolo.Valor.(Ast.TipoRetornado).Valor.(Funcion)

		for i := 0; i < funcion.Instrucciones.Len(); i++ {
			actual := funcion.Instrucciones.GetValue(i)
			if actual != nil {
				_, tipoParticular := actual.(Ast.Abstracto).GetTipo()
				if tipoParticular == Ast.DECLARACION {
					contadorDeclaraciones++
				}
			} else {
				continue
			}
		}

		//Contar los parametros
		contadorDeclaraciones += funcion.Parametros.Len()
		if l.ScopeOriginal != nil {
			newScope.Posicion = l.ScopeOriginal.Size + l.ScopeOriginal.Posicion
			//Primera posicion para el return
			newScope.Size = contadorDeclaraciones
			newScope.Size++
			newScope.ContadorDeclaracion++
			simbolo.Direccion = l.ScopeOriginal.Size + l.ScopeOriginal.Posicion
		} else {
			newScope.Posicion = scope.Size + scope.Posicion
			//Primera posicion para el return
			newScope.Size = contadorDeclaraciones
			newScope.Size++
			newScope.ContadorDeclaracion++
			simbolo.Direccion = scope.Size + scope.Posicion
		}

	}

	//Verificar que la función reciba o no parámetros y se estén enviando parámetros

	if l.Parametros.Len() > 0 && funcion.Parametros.Len() == 0 {
		//Error, se estan enviando parámetros y la función no pide parámetros
		msg := "Semantic error, " + l.Identificador.(expresiones.Identificador).Valor + " function doesn't expect parameters." +
			" -- Line: " + strconv.Itoa(l.Fila) +
			" Column: " + strconv.Itoa(l.Columna)
		nError := errores.NewError(l.Fila, l.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		newScope.Errores.Add(nError)
		newScope.Consola += msg + "\n"
		//newScope.UpdateScopeGlobal()
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}

	if l.Parametros.Len() == 0 && funcion.Parametros.Len() > 0 {
		//Error, la función espera parámetros y no se están enviando parámetros
		msg := "Semantic error, " + l.Identificador.(expresiones.Identificador).Valor + " function expects parameters." +
			" -- Line: " + strconv.Itoa(l.Fila) +
			" Column: " + strconv.Itoa(l.Columna)
		nError := errores.NewError(l.Fila, l.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		newScope.Errores.Add(nError)
		newScope.Consola += msg + "\n"
		//newScope.UpdateScopeGlobal()
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}

	//Verificar que el scopeOriginal no sea null
	if l.ScopeOriginal == nil {
		l.ScopeOriginal = scope
	}
	codigo3d += "/*******************************LLAMADA FUNCION*/\n"
	codigo3d += "/*********************DECLARACION DE PARAMETROS*/\n"
	/********************************CAMBIO DE AMBITO PARA DECLARAR PARAMENTROS********************/
	if l.ScopeOriginal != nil {
		codigo3d += "P = P + " + strconv.Itoa(l.ScopeOriginal.Size) + "; //Set direccion ambito simulado \n"
	} else {
		codigo3d += "P = P + " + strconv.Itoa(scope.Size) + "; //Set direccion ambito simulado \n"
	}

	/**********************************************************************************************/
	//Crear los parámetros de las funciones
	parametrosCreados = funcion.RunParametros(&newScope, l.ScopeOriginal, l.Parametros)
	obj3dValor = parametrosCreados.Valor.(Ast.O3D)
	if parametrosCreados.Tipo == Ast.ERROR {
		//newScope.UpdateScopeGlobal()
		return parametrosCreados
	}
	parametrosCreados = obj3dValor.Valor
	codigo3d += obj3dValor.Codigo

	//Ejecutar la función
	//Verificar la recursividad
	/*
		if !Ast.CompararFuncionStack(funcion.Nombre) {
			resultadoFuncion = funcion.Run(&newScope).(Ast.TipoRetornado)
			Ast.SetFuncionStack(funcion.Nombre)
			Ast.SetResultadoFuncionStack(resultadoFuncion)
		} else {
			Ast.SetFuncionStack("")
			resultadoFuncion = Ast.GetResultadoFuncionStack()
		}
	*/
	resultadoFuncion = funcion.Run(&newScope).(Ast.TipoRetornado)
	//Agregar funcion a stack por tema de recursividad

	obj3dTemp = resultadoFuncion.Valor.(Ast.O3D)
	codigoTemp += "/**************************EJECUCION DE FUNCION*/\n"
	codigoTemp += obj3dTemp.Codigo
	codigoTemp += "/***********************************************/\n"
	codigo3d += codigoTemp
	/********************************RETORNO AL AMBITO ANTERIOR************************************/

	if l.ScopeOriginal != nil {
		codigo3d += "P = P - " + strconv.Itoa(l.ScopeOriginal.Size) + "; //Retorno al ambito anterior \n"
	} else {
		codigo3d += "P = P - " + strconv.Itoa(scope.Size) + "; //Retorno al ambito anterior \n"
	}

	/**********************************************************************************************/
	codigo3d += "/***********************************************/\n"
	codigo3d += "/***********************************************/\n"
	obj3dTemp.Codigo = codigo3d
	resultadoFuncion.Valor = obj3dTemp

	/*********************************ACTUALIZAR SIMBOLO PARA QUE NO GENERE MÁS CODIGO3D***************/
	if !simbolo.CodigoGenerado {
		simbolo.CodigoGenerado = true
		simbolo.ReferenciaRetorno = obj3dTemp.Referencia
		scope.Update_fms_local(simbolo.Identificador, simbolo)
	} else {
		obj3dTemp.Referencia = simbolo.ReferenciaRetorno
		resultadoFuncion.Valor = obj3dTemp
	}
	/**************************************************************************************************/

	//newScope.Codigo += codigo3d
	//newScope.UpdateScopeGlobal()
	/*
		if newScope.Errores.Len() > 0 {
			msg := "Semantic error, " + l.Identificador.(expresiones.Identificador).Valor + " function expects parameters." +
				" -- Line: " + strconv.Itoa(l.Fila) +
				" Column: " + strconv.Itoa(l.Columna)
			nError := errores.NewError(l.Fila, l.Columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}
		}*/
	return resultadoFuncion
	/*
		return Ast.TipoRetornado{
			Tipo:  Ast.EJECUTADO,
			Valor: true,
		}*/
}

func (l LlamadaFuncion) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, l.Tipo
}

func (l LlamadaFuncion) GetFila() int {
	return l.Fila
}
func (l LlamadaFuncion) GetColumna() int {
	return l.Columna
}
