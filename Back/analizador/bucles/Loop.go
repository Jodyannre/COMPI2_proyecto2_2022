package bucles

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"strconv"

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
	newScope := Ast.NewScope("Loop", scope)
	var contadorSeguridad int = 0
	var retornarBreak bool = false
	var instruccion, resultado interface{}
	var tipoGeneral, _ Ast.TipoDato
	//var respuesta interface{}
	var i int = 0

	for {

		if contadorSeguridad == 1000 {
			msg := "Semantic error, infinite loop." +
				" -- Line:" + strconv.Itoa(l.Fila) + " Column: " +
				strconv.Itoa(l.Columna)
			nError := errores.NewError(l.Fila, l.Columna, msg)
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
				if resultado.(Ast.TipoRetornado).Tipo == Ast.ERROR {
					return resultado
				}
			} else if tipoGeneral == Ast.EXPRESION {
				resultado = instruccion.(Ast.Expresion).GetValue(&newScope)
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
					newScope.UpdateScopeGlobal()
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
				newScope.UpdateScopeGlobal()
				return resultado.(Ast.TipoRetornado).Valor
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
					newScope.UpdateScopeGlobal()
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
				newScope.UpdateScopeGlobal()
				newScope.Errores.Clear()
				newScope.Consola = ""
				continue
			}

			if resultado.(Ast.TipoRetornado).Tipo == Ast.CONTINUE {
				//Siguiente iteración
				newScope.UpdateScopeGlobal()
				newScope.Errores.Clear()
				newScope.Consola = ""
				break
			}

			if resultado.(Ast.TipoRetornado).Tipo == Ast.BREAK {
				//Terminar el loop
				newScope.UpdateScopeGlobal()
				return Ast.TipoRetornado{
					Tipo:  Ast.EJECUTADO,
					Valor: true,
				}
			}

			if resultado.(Ast.TipoRetornado).Tipo == Ast.RETURN ||
				resultado.(Ast.TipoRetornado).Tipo == Ast.RETURN_EXPRESION ||
				resultado.(Ast.TipoRetornado).Tipo == Ast.BREAK_EXPRESION {
				newScope.UpdateScopeGlobal()
				//Terminar loop y retornar el return
				return resultado
			}
		}
		contadorSeguridad++
	}
}

func (op Loop) GetFila() int {
	return op.Fila
}
func (op Loop) GetColumna() int {
	return op.Columna
}
