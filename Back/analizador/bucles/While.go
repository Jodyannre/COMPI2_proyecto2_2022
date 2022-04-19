package bucles

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"strconv"
	"strings"

	"github.com/colegno/arraylist"
)

type While struct {
	Tipo          Ast.TipoDato
	Condicion     Ast.Expresion
	Instrucciones *arraylist.List
	Fila          int
	Columna       int
}

func (w While) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, w.Tipo
}

func NewWhile(tipo Ast.TipoDato, condicion Ast.Expresion, instrucciones *arraylist.List, fila, columna int) While {
	nW := While{
		Tipo:          tipo,
		Instrucciones: instrucciones,
		Condicion:     condicion,
		Fila:          fila,
		Columna:       columna,
	}
	return nW
}

func (w While) Run(scope *Ast.Scope) interface{} {
	/******************************VARIABLES 3D*********************************/
	var obj3d, obj3dcondicion, obj3dresultadoInstruccion Ast.O3D
	var codigo3d, lt, lf, saltoWhile string
	var saltoContinue, saltoBreak, saltoReturn string
	/***************************************************************************/

	newScope := Ast.NewScope("Loop", scope)
	newScope.Posicion = scope.Size

	var condicionResultado Ast.TipoRetornado
	var instruccion, resultado interface{}
	var tipoGeneral Ast.TipoDato
	var i int = 0
	saltoWhile = Ast.GetLabel()
	//Validar la condición de inicio
	codigo3d += "/***********CAMBIO A ENTORNO SIMULADO DEL WHILE*/ \n"
	codigo3d += "P = P + " + strconv.Itoa(scope.Size) + "; //Cambio de entorno \n"
	codigo3d += "//#aquiVaElSaltoContinue\n"
	codigo3d += "/*******************************EJECUCION WHILE*/ \n"
	codigo3d += saltoWhile + " :\n"
	condicionResultado = w.Condicion.GetValue(&newScope)
	obj3dcondicion = condicionResultado.Valor.(Ast.O3D)
	condicionResultado = obj3dcondicion.Valor
	lt = obj3dcondicion.Lt
	lf = obj3dcondicion.Lf

	codigo3d += obj3dcondicion.Codigo

	if condicionResultado.Tipo == Ast.ERROR {
		newScope.UpdateScopeGlobal()
		return condicionResultado
	}
	if condicionResultado.Tipo != Ast.BOOLEAN {
		//Error, la condición no es booleano
		msg := "Semantic error, While condition is not a boolean type." +
			" -- Line:" + strconv.Itoa(w.Fila) + " Column: " +
			strconv.Itoa(w.Columna)
		nError := errores.NewError(w.Fila, w.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		newScope.Errores.Add(nError)
		newScope.Consola += msg + "\n"
		newScope.UpdateScopeGlobal()
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}

	codigo3d += lt + ":\n"
	//Ejecutar las instrucciones
	for i = 0; i < w.Instrucciones.Len(); i++ {
		instruccion = w.Instrucciones.GetValue(i)
		//Recuperar el tipo de la instrucción
		tipoGeneral, _ = instruccion.(Ast.Abstracto).GetTipo()
		//Verificar tipos

		if tipoGeneral == Ast.INSTRUCCION {
			//Declarar variables globales
			resultado = instruccion.(Ast.Instruccion).Run(&newScope)
			obj3dresultadoInstruccion = resultado.(Ast.TipoRetornado).Valor.(Ast.O3D)
			codigo3d += obj3dresultadoInstruccion.Codigo
			resultado = obj3dresultadoInstruccion.Valor
			if resultado.(Ast.TipoRetornado).Tipo == Ast.ERROR {
				return resultado
			}
		} else if tipoGeneral == Ast.EXPRESION {
			resultado = instruccion.(Ast.Expresion).GetValue(&newScope)
			obj3dresultadoInstruccion = resultado.(Ast.TipoRetornado).Valor.(Ast.O3D)
			codigo3d += obj3dresultadoInstruccion.Codigo
			resultado = obj3dresultadoInstruccion.Valor
			if resultado.(Ast.TipoRetornado).Tipo == Ast.ERROR {
				return resultado
			}
		}
		/*
			if resultado.(Ast.TipoRetornado).Tipo == Ast.CONTINUE {
				//Siguiente iteración
				newScope.UpdateScopeGlobal()
				newScope.Errores.Clear()
				newScope.Consola = ""
				break
			}
		*/

		if resultado.(Ast.TipoRetornado).Tipo == Ast.CONTINUE ||
			resultado.(Ast.TipoRetornado).Tipo == Ast.BREAK ||
			resultado.(Ast.TipoRetornado).Tipo == Ast.RETURN {
			//Siguiente iteración
			newScope.UpdateScopeGlobal()
			newScope.Errores.Clear()
			newScope.Consola = ""

			if !obj3dresultadoInstruccion.TranferenciaAgregada {
				codigo3d += "goto " + obj3dresultadoInstruccion.Salto + ";\n"
			}
			saltoContinue += obj3dresultadoInstruccion.SaltoContinue
			saltoContinue = strings.Replace(saltoContinue, ",", ":\n", -1)

			saltoBreak += obj3dresultadoInstruccion.SaltoBreak
			//println(saltoBreak)
			saltoBreak = strings.Replace(saltoBreak, ",", ":\n", -1)

			saltoReturn += obj3dresultadoInstruccion.SaltoReturn
			saltoReturn = strings.Replace(saltoReturn, ",", ":\n", -1)

			continue
		}

	}
	codigo3d += "goto " + saltoWhile + "; \n"
	codigo3d += lf + ":\n"
	newScope.UpdateScopeGlobal()
	newScope.Errores.Clear()
	newScope.Consola = ""
	/*
		condicionResultado = w.Condicion.GetValue(&newScope)

		if condicionResultado.Tipo == Ast.ERROR {
			newScope.UpdateScopeGlobal()
			return condicionResultado
		}
		if condicionResultado.Tipo != Ast.BOOLEAN {
			//Error, la condición no es booleano
			msg := "Semantic error, While condition is not a boolean type." +
				" -- Line:" + strconv.Itoa(w.Fila) + " Column: " +
				strconv.Itoa(w.Columna)
			nError := errores.NewError(w.Fila, w.Columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			nError.Ambito = scope.GetTipoScope()
			newScope.Errores.Add(nError)
			newScope.Consola += msg + "\n"
			newScope.UpdateScopeGlobal()
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}
		} else {
			condicion = condicionResultado.Valor.(bool)
		}
	*/

	if saltoBreak != "" {
		codigo3d += saltoBreak
	}

	if saltoContinue != "" {
		codigo3d = strings.Replace(codigo3d, "//#aquiVaElSaltoContinue", saltoContinue, -1)
	} else {
		codigo3d = strings.Replace(codigo3d, "//#aquiVaElSaltoContinue", "", -1)
	}

	codigo3d += "P = P - " + strconv.Itoa(scope.Size) + "; //Regresar al entorno anterior \n"
	codigo3d += "/***********************************************/\n"

	obj3d.Codigo = codigo3d
	obj3d.Valor = Ast.TipoRetornado{
		Tipo:  Ast.EJECUTADO,
		Valor: true,
	}

	if saltoReturn != "" {
		obj3d.SaltoReturn = saltoReturn
		obj3d.Valor.Tipo = Ast.RETURN
		return Ast.TipoRetornado{
			Tipo:  Ast.RETURN,
			Valor: obj3d,
		}
	}

	return Ast.TipoRetornado{
		Tipo:  Ast.EJECUTADO,
		Valor: obj3d,
	}
}

func (op While) GetFila() int {
	return op.Fila
}
func (op While) GetColumna() int {
	return op.Columna
}
