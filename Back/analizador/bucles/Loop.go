package bucles

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"strconv"
	"strings"

	"github.com/colegno/arraylist"
)

type Loop struct {
	Tipo          Ast.TipoDato
	Instrucciones *arraylist.List
	Fila          int
	Columna       int
}

func (l Loop) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, l.Tipo
}

func NewLoop(tipo Ast.TipoDato, instrucciones *arraylist.List, fila, columna int) Loop {
	nL := Loop{
		Tipo:          tipo,
		Instrucciones: instrucciones,
		Fila:          fila,
		Columna:       columna,
	}
	return nL
}

func (l Loop) Run(scope *Ast.Scope) interface{} {
	/*********************************VARIABLES 3D**********************************/
	var obj3d, obj3dresultadoInstruccion Ast.O3D
	var codigo3d string
	var saltoBreak, saltoContinue, saltoReturn, saltoLoop, saltoReturnExp string
	var valorReturn Ast.TipoRetornado
	//var reemplazoContinue string
	/*******************************************************************************/

	newScope := Ast.NewScope("Loop", scope)
	newScope.Posicion = scope.Size
	var retornarBreak bool = false
	var instruccion, resultado interface{}
	var tipoGeneral, _ Ast.TipoDato
	//var respuesta interface{}
	var i int = 0
	codigo3d += "/************CAMBIO A ENTORNO SIMULADO DEL LOOP*/ \n"
	codigo3d += "P = P + " + strconv.Itoa(scope.Size) + "; //Cambio de entorno \n"
	saltoLoop = Ast.GetLabel()
	codigo3d += saltoLoop + ":\n"
	codigo3d += "/********************************EJECUCION LOOP*/ \n"
	codigo3d += "//#aquiVaElSaltoContinue\n"
	//Ejecutar las instrucciones
	for i = 0; i < l.Instrucciones.Len(); i++ {
		instruccion = l.Instrucciones.GetValue(i)
		//Recuperar el tipo de la instrucción
		//tipoGeneral, tipoParticular = instruccion.(Ast.Abstracto).GetTipo()
		//Verificar tipos
		if instruccion != nil {
			tipoGeneral, _ = instruccion.(Ast.Abstracto).GetTipo()
		} else {
			continue
		}

		if tipoGeneral == Ast.INSTRUCCION {
			//Declarar variables globales
			resultado = instruccion.(Ast.Instruccion).Run(&newScope)
			obj3dresultadoInstruccion = resultado.(Ast.TipoRetornado).Valor.(Ast.O3D)
			codigo3d += obj3dresultadoInstruccion.Codigo
			resultado = obj3dresultadoInstruccion.Valor

			if resultado.(Ast.TipoRetornado).Tipo == Ast.RETURN ||
				resultado.(Ast.TipoRetornado).Tipo == Ast.RETURN_EXPRESION {
				valorReturn = resultado.(Ast.TipoRetornado)
			}

			if resultado.(Ast.TipoRetornado).Tipo == Ast.ERROR {
				return resultado
			}
		} else if tipoGeneral == Ast.EXPRESION {
			resultado = instruccion.(Ast.Expresion).GetValue(&newScope)
			obj3dresultadoInstruccion = resultado.(Ast.TipoRetornado).Valor.(Ast.O3D)
			codigo3d += obj3dresultadoInstruccion.Codigo
			resultado = obj3dresultadoInstruccion.Valor

			if resultado.(Ast.TipoRetornado).Tipo == Ast.RETURN ||
				resultado.(Ast.TipoRetornado).Tipo == Ast.RETURN_EXPRESION {
				valorReturn = resultado.(Ast.TipoRetornado)
			}

			if resultado.(Ast.TipoRetornado).Tipo == Ast.ERROR {
				return resultado
			}
		}

		//Verificar los breaks y los return

		//Error, solo puede venir un break expresion
		if Ast.EsTransferencia(resultado.(Ast.TipoRetornado).Tipo) &&
			l.Tipo == Ast.LOOP_EXPRESION {
			if resultado.(Ast.TipoRetornado).Tipo != Ast.BREAK_EXPRESION &&
				resultado.(Ast.TipoRetornado).Tipo != Ast.CONTINUE {
				msg := "Semantic error," + Ast.ValorTipoDato[resultado.(Ast.TipoRetornado).Tipo] + " statement not allowed inside this kind of loop." +
					" -- Line:" + strconv.Itoa(instruccion.(Ast.Abstracto).GetFila()) + " Column: " +
					strconv.Itoa(instruccion.(Ast.Abstracto).GetColumna())
				nError := errores.NewError(instruccion.(Ast.Abstracto).GetFila(),
					instruccion.(Ast.Abstracto).GetColumna(), msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				nError.Ambito = scope.GetTipoScope()
				newScope.Errores.Add(nError)
				newScope.Consola += msg + "\n"
				//newScope.UpdateScopeGlobal()
				return Ast.TipoRetornado{
					Valor: nError,
					Tipo:  Ast.ERROR,
				}
			} else if resultado.(Ast.TipoRetornado).Tipo == Ast.BREAK_EXPRESION {
				//Si ambos son de tipo expresion
				retornarBreak = true
			}

		}

		if retornarBreak {
			//Terminar loop y retornar el valor del break
			//newScope.UpdateScopeGlobal()
			return resultado.(Ast.TipoRetornado)
		}

		if Ast.EsTransferencia(resultado.(Ast.TipoRetornado).Tipo) &&
			l.Tipo == Ast.LOOP_EXPRESION {
			if resultado.(Ast.TipoRetornado).Tipo != Ast.BREAK_EXPRESION &&
				resultado.(Ast.TipoRetornado).Tipo != Ast.CONTINUE {
				msg := "Semantic error," + Ast.ValorTipoDato[resultado.(Ast.TipoRetornado).Tipo] + " statement not allowed inside this kind of loop." +
					" -- Line:" + strconv.Itoa(instruccion.(Ast.Abstracto).GetFila()) + " Column: " +
					strconv.Itoa(instruccion.(Ast.Abstracto).GetColumna())
				nError := errores.NewError(instruccion.(Ast.Abstracto).GetFila(),
					instruccion.(Ast.Abstracto).GetColumna(), msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				nError.Ambito = scope.GetTipoScope()
				newScope.Errores.Add(nError)
				newScope.Consola += msg + "\n"
				//newScope.UpdateScopeGlobal()
				return Ast.TipoRetornado{
					Valor: nError,
					Tipo:  Ast.ERROR,
				}
			} else if resultado.(Ast.TipoRetornado).Tipo == Ast.BREAK_EXPRESION {
				//Si ambos son de tipo expresion
				retornarBreak = true
			}

		}

		if resultado.(Ast.TipoRetornado).Tipo == Ast.ERROR ||
			resultado.(Ast.TipoRetornado).Tipo == Ast.EJECUTADO {
			//Siguiente instrucción
			//newScope.UpdateScopeGlobal()
			newScope.Errores.Clear()
			newScope.Consola = ""
			continue
		}

		if resultado.(Ast.TipoRetornado).Tipo == Ast.CONTINUE ||
			resultado.(Ast.TipoRetornado).Tipo == Ast.BREAK ||
			resultado.(Ast.TipoRetornado).Tipo == Ast.RETURN {
			//Siguiente iteración
			//newScope.UpdateScopeGlobal()
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
			saltoReturnExp += obj3dresultadoInstruccion.SaltoReturnExp
			//saltoReturn = strings.Replace(saltoReturn, ",", ":\n", -1)

			continue
		}

		/*
			if resultado.(Ast.TipoRetornado).Tipo == Ast.BREAK {
				//Terminar el loop
				newScope.UpdateScopeGlobal()
				if !obj3dresultadoInstruccion.TranferenciaAgregada {
					codigo3d += "goto " + obj3dresultadoInstruccion.Salto + ";\n"
				}
				saltoBreak += obj3dresultadoInstruccion.SaltoBreak
				//println(saltoBreak)
				saltoBreak = strings.Replace(saltoBreak, ",", ":\n", -1)
				//saltoBreak += ":\n"
				continue
			}
		*/

		if resultado.(Ast.TipoRetornado).Tipo == Ast.RETURN ||
			resultado.(Ast.TipoRetornado).Tipo == Ast.RETURN_EXPRESION ||
			resultado.(Ast.TipoRetornado).Tipo == Ast.BREAK_EXPRESION {
			//newScope.UpdateScopeGlobal()
			//Terminar loop y retornar el return
			return resultado
		}
	}
	codigo3d += "goto " + saltoLoop + ";\n"

	if saltoBreak != "" {
		codigo3d += saltoBreak
	}

	if saltoContinue != "" {
		codigo3d = strings.Replace(codigo3d, "//#aquiVaElSaltoContinue", saltoContinue, -1)
	} else {
		codigo3d = strings.Replace(codigo3d, "//#aquiVaElSaltoContinue", "", -1)
	}

	if saltoReturnExp != "" {
		saltoReturnExp = strings.Replace(saltoReturnExp, ",", ":\n", -1)
		if saltoReturnExp[len(saltoReturnExp)-1] != '\n' {
			saltoReturnExp += ","
		}
		codigo3d += saltoReturnExp
	}
	codigo3d += "P = P - " + strconv.Itoa(scope.Size) + "; //Regresar al entorno anterior \n"
	codigo3d += "/***********************************************/\n"
	obj3d.Valor = Ast.TipoRetornado{
		Valor: true,
		Tipo:  Ast.EJECUTADO,
	}

	obj3d.Codigo = codigo3d

	if saltoReturnExp != "" {
		saltoNuevo := Ast.GetLabel()
		codigo3d += "goto " + saltoNuevo + ";\n"
		obj3d.Codigo = codigo3d
		obj3d.SaltoReturn = saltoReturn + ","
		obj3d.SaltoReturnExp = saltoNuevo + ","
		obj3d.Valor.Tipo = Ast.RETURN
		obj3d.Valor = valorReturn
		obj3d.TranferenciaAgregada = true
		return Ast.TipoRetornado{
			Tipo:  Ast.RETURN,
			Valor: obj3d,
		}
	}

	if saltoReturn != "" {
		obj3d.SaltoReturn = saltoReturn
		obj3d.SaltoReturnExp = saltoReturnExp
		obj3d.Valor.Tipo = Ast.RETURN
		obj3d.Valor = valorReturn
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

func (op Loop) GetFila() int {
	return op.Fila
}
func (op Loop) GetColumna() int {
	return op.Columna
}
