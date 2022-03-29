package bucles

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"strconv"

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
	newScope := Ast.NewScope("Loop", scope)
	var contadorSeguridad int = 0
	var condicion = true
	var condicionResultado Ast.TipoRetornado
	var instruccion, resultado interface{}
	var tipoGeneral Ast.TipoDato
	var i int = 0

	//Validar la condición de inicio
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

	for condicion {

		if contadorSeguridad == 1000 {
			msg := "Semantic error, infinite loop." +
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

		//Ejecutar las instrucciones
		for i = 0; i < w.Instrucciones.Len(); i++ {
			instruccion = w.Instrucciones.GetValue(i)
			//Recuperar el tipo de la instrucción
			tipoGeneral, _ = instruccion.(Ast.Abstracto).GetTipo()
			//Verificar tipos

			/*
				if tipoGeneral == Ast.EXPRESION {
					//Error, no pueden existir expresiones aisladas
					msg := "Semantic error, an instruction was expected." +
						" -- Line:" + strconv.Itoa(instruccion.(Ast.Abstracto).GetFila()) + " Column: " +
						strconv.Itoa(instruccion.(Ast.Abstracto).GetColumna())
					nError := errores.NewError(instruccion.(Ast.Abstracto).GetFila(),
						instruccion.(Ast.Abstracto).GetColumna(), msg)
					nError.Tipo = Ast.ERROR_SEMANTICO
					newScope.Errores.Add(nError)
					newScope.Consola += msg + "\n"
					newScope.UpdateScopeGlobal()
					return Ast.TipoRetornado{
						Valor: nError,
						Tipo:  Ast.ERROR,
					}
				}
			*/
			/*
				if tipoGeneral == Ast.INSTRUCCION {

					//Ejecutar la instrucción
					resultado = instruccion.(Ast.Instruccion).Run(scope)

				}
			*/
			if tipoGeneral == Ast.INSTRUCCION {
				//Declarar variables globales
				resultado = instruccion.(Ast.Instruccion).Run(&newScope)
				if resultado.(Ast.TipoRetornado).Tipo == Ast.ERROR {
					return resultado
				}
			} else if tipoGeneral == Ast.EXPRESION {
				resultado = instruccion.(Ast.Expresion).GetValue(&newScope)
				if resultado.(Ast.TipoRetornado).Tipo == Ast.ERROR {
					return resultado
				}
			}
			/*
				if resultado.(Ast.TipoRetornado).Tipo == Ast.ERROR ||
					resultado.(Ast.TipoRetornado).Tipo == Ast.EJECUTADO {
					//Siguiente instrucción
					newScope.UpdateScopeGlobal()
					newScope.Errores.Clear()
					newScope.Consola = ""
					continue
				}
			*/
			if resultado.(Ast.TipoRetornado).Tipo == Ast.CONTINUE {
				//Siguiente iteración
				newScope.UpdateScopeGlobal()
				newScope.Errores.Clear()
				newScope.Consola = ""
				break
			}

			if Ast.EsTransferencia(resultado.(Ast.TipoRetornado).Tipo) {
				if resultado.(Ast.TipoRetornado).Tipo == Ast.CONTINUE {
					newScope.UpdateScopeGlobal()
					newScope.Errores.Clear()
					newScope.Consola = ""
					break
				}
				if resultado.(Ast.TipoRetornado).Tipo == Ast.BREAK {
					newScope.UpdateScopeGlobal()
					return Ast.TipoRetornado{
						Tipo:  Ast.EJECUTADO,
						Valor: true,
					}
				}
				//Terminar el loop
				return resultado
			}

		}
		contadorSeguridad++
		newScope.UpdateScopeGlobal()
		newScope.Errores.Clear()
		newScope.Consola = ""
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
	}
	return Ast.TipoRetornado{
		Tipo:  Ast.EJECUTADO,
		Valor: true,
	}
}

func (op While) GetFila() int {
	return op.Fila
}
func (op While) GetColumna() int {
	return op.Columna
}
