package simbolos

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"strconv"

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

	//Primero crear el nuevo scope main
	newScope := Ast.NewScope("Main", scope)
	var actual interface{}
	var tipoGeneral interface{}
	var respuesta interface{}

	//Recorrer y ejecutar todas las instrucciones
	for i := 0; i < f.Instrucciones.Len(); i++ {
		actual = f.Instrucciones.GetValue(i)
		if actual != nil {
			tipoGeneral, _ = actual.(Ast.Abstracto).GetTipo()
		} else {
			continue
		}
		if tipoGeneral == Ast.INSTRUCCION {
			//Declarar variables globales
			respuesta = actual.(Ast.Instruccion).Run(&newScope)
			//No es necesario verificar si trae error o es ejecutada, lo único que hay que verificar es
			//Que traiga retornos que puede generar errores
			if Ast.EsTransferencia(respuesta.(Ast.TipoRetornado).Tipo) {
				//Variables para el msg del error
				valor := actual.(Ast.Abstracto)
				fila := valor.GetFila()
				columna := valor.GetColumna()
				msg := ""
				//Primero verificar que no sea un return normal, el cual si es permitido
				if respuesta.(Ast.TipoRetornado).Tipo == Ast.RETURN {
					return Ast.TipoRetornado{
						Tipo:  Ast.EJECUTADO,
						Valor: true,
					}
				}
				switch respuesta.(Ast.TipoRetornado).Tipo {
				case Ast.BREAK, Ast.BREAK_EXPRESION, Ast.CONTINUE:
					msg = "Semantic error, cannot break outside of a loop." +
						" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
				case Ast.RETURN_EXPRESION:
					msg = "Semantic error, MAIN method cannot return a value." +
						" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
				}
				nError := errores.NewError(fila, columna, msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				nError.Ambito = scope.GetTipoScope()
				newScope.Errores.Add(nError)
				newScope.Consola += msg + "\n"
			}

		} else if tipoGeneral == Ast.EXPRESION {
			respuesta = actual.(Ast.Expresion).GetValue(&newScope)

			println(Ast.Temporales)
			println(respuesta.(Ast.TipoRetornado).Valor.(Ast.O3D).Codigo)

			if Ast.EsTransferencia(respuesta.(Ast.TipoRetornado).Tipo) {
				//Variables para el msg del error
				valor := actual.(Ast.Abstracto)
				fila := valor.GetFila()
				columna := valor.GetColumna()
				msg := ""
				//Primero verificar que no sea un return normal, el cual si es permitido
				if respuesta.(Ast.TipoRetornado).Tipo == Ast.RETURN {
					return Ast.TipoRetornado{
						Tipo:  Ast.EJECUTADO,
						Valor: true,
					}
				}
				switch respuesta.(Ast.TipoRetornado).Tipo {
				case Ast.BREAK, Ast.BREAK_EXPRESION, Ast.CONTINUE:
					msg = "Semantic error, cannot break outside of a loop." +
						" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
				case Ast.RETURN_EXPRESION:
					msg = "Semantic error, MAIN method cannot return a value." +
						" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
				}
				nError := errores.NewError(fila, columna, msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				nError.Ambito = scope.GetTipoScope()
				newScope.Errores.Add(nError)
				newScope.Consola += msg + "\n"
			}

		}
	}
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
