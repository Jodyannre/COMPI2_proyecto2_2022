package simbolos

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"

	"github.com/colegno/arraylist"
)

type AccesoModulo struct {
	Tipo      Ast.TipoDato
	Fila      int
	Columna   int
	Elementos *arraylist.List
}

func NewAccesoModulo(elementos *arraylist.List, fila, columna int) AccesoModulo {
	nA := AccesoModulo{
		Tipo:      Ast.ACCESO_MODULO,
		Fila:      fila,
		Columna:   columna,
		Elementos: elementos,
	}
	return nA
}

func (a AccesoModulo) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	/******************************VARIABLES 3D**********************************/
	var obj3d, obj3dValor Ast.O3D
	var resultadoFuncion Ast.TipoRetornado
	/****************************************************************************/

	// FUNC::ID::ID
	var idElementoGlobal interface{}
	var idModuloGlobal string
	var tipoParticular Ast.TipoDato
	var simboloGlobal Ast.Simbolo
	var simboloLocal Ast.Simbolo
	var moduloGlobal Modulo
	var scopeValido *Ast.Scope
	var idElementoActual interface{}
	var idElementString string
	var simboloElementoActual Ast.Simbolo
	//var estructura interface{}
	//Verificar que el primer módulo sea global y exista en el ámbito global
	//Get el id del módulo global
	//En la última posición porque vienen al reves ya que es recursiva por derecha
	//Ya no esta más al reves
	idElementoGlobal = a.Elementos.GetValue(0)
	idModuloGlobal = idElementoGlobal.(expresiones.Identificador).Valor

	//Verificar que el módulo existe, obtener el símbolo donde esta guardado
	simboloGlobal = scope.Exist_fms_declaracion(idModuloGlobal)
	simboloLocal = scope.Exist_fms_local(idModuloGlobal)

	if simboloGlobal.Tipo != Ast.MODULO {
		simboloGlobal = simboloLocal
	}

	//Verificar que el símbolo exista
	if simboloGlobal.Tipo == Ast.ERROR_NO_EXISTE {
		//Crear error de que el módulo no existe en el ámbito global al menos
		fila := idElementoGlobal.(Ast.Abstracto).GetFila()
		columna := idElementoGlobal.(Ast.Abstracto).GetColumna()
		msg := "Semantic error, the MODULE: \"" + idModuloGlobal + "\" doesn't exist." +
			" -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
		nError := errores.NewError(fila, columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		//scope.UpdateScopeGlobal()
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}

	//Verificar que no sea privado

	if simboloGlobal.Tipo == Ast.ERROR_ACCESO_PRIVADO {
		//Crear error de que el módulo no existe en el ámbito global al menos
		fila := idElementoGlobal.(Ast.Abstracto).GetFila()
		columna := idElementoGlobal.(Ast.Abstracto).GetColumna()
		msg := "Semantic error, the element: \"" + idModuloGlobal + "\" is private." +
			" -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
		nError := errores.NewError(fila, columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		//scope.UpdateScopeGlobal()
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}

	//Verificar que el elemento que esta en el simbolo sea un módulo
	if simboloGlobal.Tipo != Ast.MODULO {
		fila := idElementoGlobal.(Ast.Abstracto).GetFila()
		columna := idElementoGlobal.(Ast.Abstracto).GetColumna()
		msg := "Semantic error, a MODULE expected, found" + Ast.ValorTipoDato[simboloGlobal.Tipo] +
			". -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
		nError := errores.NewError(fila, columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		//scope.UpdateScopeGlobal()
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}

	}
	//Get el módulo
	moduloGlobal = simboloGlobal.Valor.(Ast.TipoRetornado).Valor.(Modulo)
	scopeValido = moduloGlobal.Entorno

	//Verificar si la lista de elemento esta vacia
	if a.Elementos.Len() <= 1 {
		//Error, no se esta accediendo a nada
		fila := idElementoGlobal.(Ast.Abstracto).GetFila()
		columna := idElementoGlobal.(Ast.Abstracto).GetColumna()
		msg := "Semantic error, expected value, found module \"" + idModuloGlobal + "\"." +
			" -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
		nError := errores.NewError(fila, columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		//scope.UpdateScopeGlobal()
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}

	//Iterar los elementos para buscar el acceso
	for i := 1; i < a.Elementos.Len(); i++ {
		elementoActual := a.Elementos.GetValue(i)
		_, tipoParticular = elementoActual.(Ast.Abstracto).GetTipo()
		if tipoParticular == Ast.IDENTIFICADOR {
			//Es un id y puede ser un struct u otro módulo
			idElementString = elementoActual.(expresiones.Identificador).Valor
			//Verificar que exista el elemento
			simboloElementoActual = scopeValido.Exist_fms_local(idElementString)
			if simboloElementoActual.Tipo == Ast.ERROR_ACCESO_PRIVADO {
				fila := elementoActual.(Ast.Abstracto).GetFila()
				columna := elementoActual.(Ast.Abstracto).GetColumna()
				_, tipoParticular := elementoActual.(Ast.Abstracto).GetTipo()
				if tipoParticular == Ast.FUNCION {
					idElementString = elementoActual.(Funcion).Nombre
				}
				msg := "Semantic error, the elemento \"" + idElementString + "\", is private." +
					" -- Line: " + strconv.Itoa(fila) +
					" Column: " + strconv.Itoa(columna)
				nError := errores.NewError(fila, columna, msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				nError.Ambito = scope.GetTipoScope()
				scope.Errores.Add(nError)
				scope.Consola += msg + "\n"
				//scope.UpdateScopeGlobal()
				return Ast.TipoRetornado{
					Tipo:  Ast.ERROR,
					Valor: nError,
				}
			}
			if simboloElementoActual.Tipo == Ast.ERROR_NO_EXISTE {

				fila := idElementoGlobal.(Ast.Abstracto).GetFila()
				columna := idElementoGlobal.(Ast.Abstracto).GetColumna()
				msg := "Semantic error, the element: \"" + idElementString + "\" doesn't exist." +
					" -- Line: " + strconv.Itoa(fila) +
					" Column: " + strconv.Itoa(columna)
				nError := errores.NewError(fila, columna, msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				nError.Ambito = scope.GetTipoScope()
				scope.Errores.Add(nError)
				scope.Consola += msg + "\n"
				//scope.UpdateScopeGlobal()
				return Ast.TipoRetornado{
					Tipo:  Ast.ERROR,
					Valor: nError,
				}
			}

			if simboloElementoActual.Tipo == Ast.STRUCT_TEMPLATE && i != a.Elementos.Len()-1 {
				fila := idElementoGlobal.(Ast.Abstracto).GetFila()
				columna := idElementoGlobal.(Ast.Abstracto).GetColumna()
				msg := "Semantic error, an (STRUCT|FUNCION) expected, found" + Ast.ValorTipoDato[simboloElementoActual.Tipo] +
					". -- Line: " + strconv.Itoa(fila) +
					" Column: " + strconv.Itoa(columna)
				nError := errores.NewError(fila, columna, msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				nError.Ambito = scope.GetTipoScope()
				scope.Errores.Add(nError)
				scope.Consola += msg + "\n"
				//scope.UpdateScopeGlobal()
				return Ast.TipoRetornado{
					Tipo:  Ast.ERROR,
					Valor: nError,
				}
			}

			if simboloElementoActual.Tipo != Ast.STRUCT_TEMPLATE && i == a.Elementos.Len()-1 {
				fila := idElementoGlobal.(Ast.Abstracto).GetFila()
				columna := idElementoGlobal.(Ast.Abstracto).GetColumna()
				msg := "Semantic error, an STRUCT expected, found" + Ast.ValorTipoDato[simboloElementoActual.Tipo] +
					". -- Line: " + strconv.Itoa(fila) +
					" Column: " + strconv.Itoa(columna)
				nError := errores.NewError(fila, columna, msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				nError.Ambito = scope.GetTipoScope()
				scope.Errores.Add(nError)
				scope.Consola += msg + "\n"
				//scope.UpdateScopeGlobal()
				return Ast.TipoRetornado{
					Tipo:  Ast.ERROR,
					Valor: nError,
				}
			}
			if simboloElementoActual.Tipo == Ast.MODULO {
				scopeValido = simboloElementoActual.Valor.(Ast.TipoRetornado).Valor.(Modulo).Entorno
			}

			if simboloElementoActual.Tipo == Ast.STRUCT_TEMPLATE && i == a.Elementos.Len()-1 {
				//Llegamos al último, retornar el struct
				//estructura = simboloElementoActual.Valor.(Ast.TipoRetornado).Valor.(Ast.Structs).GetPlantilla()
				return Ast.TipoRetornado{
					Tipo:  Ast.STRUCT_TEMPLATE,
					Valor: simboloElementoActual,
				}
			}

		} else {
			//Definitivamente es un acceso a una funcion
			idElementoActual = elementoActual.(LlamadaFuncion).Identificador

			//Verificar que sea un id
			_, tipoParticular = idElementoActual.(Ast.Abstracto).GetTipo()
			if tipoParticular != Ast.IDENTIFICADOR {
				//Error se espera un identificador
				fila := idElementoActual.(Ast.Abstracto).GetFila()
				columna := idElementoActual.(Ast.Abstracto).GetColumna()
				msg := "Semantic error, an IDENTIFICADOR expected, found" + Ast.ValorTipoDato[tipoParticular] +
					". -- Line: " + strconv.Itoa(fila) +
					" Column: " + strconv.Itoa(columna)
				nError := errores.NewError(fila, columna, msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				nError.Ambito = scope.GetTipoScope()
				scope.Errores.Add(nError)
				scope.Consola += msg + "\n"
				//scope.UpdateScopeGlobal()
				return Ast.TipoRetornado{
					Tipo:  Ast.ERROR,
					Valor: nError,
				}
			}
			//Continuar, ahora conseguir el string del id
			idElementString = idElementoActual.(expresiones.Identificador).Valor
			//Verificar que exista el elemento
			simboloElementoActual = scopeValido.Exist_fms_local(idElementString)

			if simboloElementoActual.Tipo == Ast.ERROR_ACCESO_PRIVADO {
				fila := idElementoActual.(Ast.Abstracto).GetFila()
				columna := idElementoActual.(Ast.Abstracto).GetColumna()
				msg := "Semantic error, the elemento \"" + idElementString + "\", is private." +
					" -- Line: " + strconv.Itoa(fila) +
					" Column: " + strconv.Itoa(columna)
				nError := errores.NewError(fila, columna, msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				nError.Ambito = scope.GetTipoScope()
				scope.Errores.Add(nError)
				scope.Consola += msg + "\n"
				//scope.UpdateScopeGlobal()
				return Ast.TipoRetornado{
					Tipo:  Ast.ERROR,
					Valor: nError,
				}
			}
			if simboloElementoActual.Tipo == Ast.ERROR_NO_EXISTE {

				fila := idElementoGlobal.(Ast.Abstracto).GetFila()
				columna := idElementoGlobal.(Ast.Abstracto).GetColumna()
				msg := "Semantic error, the element: \"" + idModuloGlobal + "\" doesn't exist." +
					" -- Line: " + strconv.Itoa(fila) +
					" Column: " + strconv.Itoa(columna)
				nError := errores.NewError(fila, columna, msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				nError.Ambito = scope.GetTipoScope()
				scope.Errores.Add(nError)
				scope.Consola += msg + "\n"
				//scope.UpdateScopeGlobal()
				return Ast.TipoRetornado{
					Tipo:  Ast.ERROR,
					Valor: nError,
				}
			}

			if simboloElementoActual.Tipo != Ast.FUNCION {
				fila := idElementoGlobal.(Ast.Abstracto).GetFila()
				columna := idElementoGlobal.(Ast.Abstracto).GetColumna()
				msg := "Semantic error, the element: \"" + idModuloGlobal + "\" is not a FUNCION." +
					" -- Line: " + strconv.Itoa(fila) +
					" Column: " + strconv.Itoa(columna)
				nError := errores.NewError(fila, columna, msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				nError.Ambito = scope.GetTipoScope()
				scope.Errores.Add(nError)
				scope.Consola += msg + "\n"
				//scope.UpdateScopeGlobal()
				return Ast.TipoRetornado{
					Tipo:  Ast.ERROR,
					Valor: nError,
				}
			}

			if simboloElementoActual.Tipo == Ast.FUNCION && i != a.Elementos.Len()-1 {
				fila := idElementoGlobal.(Ast.Abstracto).GetFila()
				columna := idElementoGlobal.(Ast.Abstracto).GetColumna()
				msg := "Semantic error, an (STRUCT|FUNCION) expected, found" + Ast.ValorTipoDato[simboloElementoActual.Tipo] +
					". -- Line: " + strconv.Itoa(fila) +
					" Column: " + strconv.Itoa(columna)
				nError := errores.NewError(fila, columna, msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				nError.Ambito = scope.GetTipoScope()
				scope.Errores.Add(nError)
				scope.Consola += msg + "\n"
				//scope.UpdateScopeGlobal()
				return Ast.TipoRetornado{
					Tipo:  Ast.ERROR,
					Valor: nError,
				}
			}
			funcion := elementoActual.(LlamadaFuncion)
			funcion.ScopeOriginal = scope
			elementoActual = funcion
			resultadoFuncion = elementoActual.(Ast.Expresion).GetValue(scopeValido)
			obj3dValor = resultadoFuncion.Valor.(Ast.O3D)
			return Ast.TipoRetornado{
				Tipo:  resultadoFuncion.Tipo,
				Valor: obj3dValor,
			}

		}
	}

	obj3d.Valor = Ast.TipoRetornado{
		Tipo:  Ast.EJECUTADO,
		Valor: true,
	}
	return Ast.TipoRetornado{
		Tipo:  Ast.EJECUTADO,
		Valor: obj3d,
	}
}

func (a AccesoModulo) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, a.Tipo
}

func (a AccesoModulo) GetFila() int {
	return a.Fila
}
func (a AccesoModulo) GetColumna() int {
	return a.Columna
}
