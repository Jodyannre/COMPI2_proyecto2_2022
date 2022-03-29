package simbolos

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"reflect"
	"strconv"
)

type AsignacionAccesoStruct struct {
	Acceso  interface{}
	Valor   interface{}
	Fila    int
	Columna int
}

func NewAsignacionAccesoStruct(acceso, valor interface{}, fila, columna int) AsignacionAccesoStruct {
	nA := AsignacionAccesoStruct{
		Acceso:  acceso,
		Valor:   valor,
		Fila:    fila,
		Columna: columna,
	}
	return nA
}

func (a AsignacionAccesoStruct) Run(scope *Ast.Scope) interface{} {
	//Ejecutar el acceso

	nombreAtributo := a.Acceso.(AccesoStruct).NombreAtributo
	nombreStruct := a.Acceso.(AccesoStruct).NombreStruct
	_, tipoParticular := nombreStruct.(Ast.Abstracto).GetTipo()
	var resultadoAtributo Ast.TipoRetornado
	var idAtributo, idStruct string

	if reflect.TypeOf(nombreAtributo) != reflect.TypeOf(expresiones.Identificador{}) {
		fila := nombreAtributo.(Ast.Abstracto).GetFila()
		columna := nombreAtributo.(Ast.Abstracto).GetColumna()
		_, tipo := nombreAtributo.(Ast.Abstracto).GetTipo()
		msg := "Semantic error, expected IDENTIFICADOR, found. " + Ast.ValorTipoDato[tipo] +
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

	}

	//Get id del atributo
	idAtributo = nombreAtributo.(expresiones.Identificador).Valor

	//Verificar de que sean identificadores

	if tipoParticular != Ast.IDENTIFICADOR {
		resultadoAtributo = nombreStruct.(Ast.Expresion).GetValue(scope)
		_, tipoParticular = resultadoAtributo.Valor.(Ast.Abstracto).GetTipo()

		if reflect.TypeOf(resultadoAtributo.Valor) != reflect.TypeOf(StructInstancia{}) {
			//Error, no es un struct
			fila := a.Valor.(Ast.Abstracto).GetFila()
			columna := a.Valor.(Ast.Abstracto).GetColumna()
			msg := "Semantic error, expected STRUCT, found. " + Ast.ValorTipoDato[tipoParticular] +
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
		}

		structInstancia := resultadoAtributo.Valor.(StructInstancia)

		//Verificar que exista el atributo
		if !structInstancia.Entorno.Exist_actual(idAtributo) {
			fila := a.Valor.(Ast.Abstracto).GetFila()
			columna := a.Valor.(Ast.Abstracto).GetColumna()
			msg := "Semantic error, the element " + idAtributo + " doesn't exist." +
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
		}

		//Get el simbolo del atributo
		simboloAtributo := structInstancia.Entorno.GetSimbolo(idAtributo)

		//Get el valor nuevo del atributo
		newValue := a.Valor.(Ast.Expresion).GetValue(scope)

		if newValue.Tipo == Ast.ERROR {
			return newValue
		}

		//Si los tipos son avanzados
		if simboloAtributo.TipoEspecial.Tipo != Ast.INDEFINIDO {
			_, tipoParticular = newValue.Valor.(Ast.Abstracto).GetTipo()
			errorTipos := false
			var tipoEntrante Ast.TipoRetornado
			switch tipoParticular {
			case Ast.VECTOR:
				valor := newValue.Valor.(expresiones.Vector)
				if !CompararTipos(simboloAtributo.TipoEspecial, valor.TipoVector) {
					errorTipos = true
					tipoEntrante = Ast.TipoRetornado{
						Tipo:  Ast.VECTOR,
						Valor: valor.TipoVector,
					}
				}
			case Ast.ARRAY:
				valor := newValue.Valor.(expresiones.Array)
				if !CompararTipos(simboloAtributo.TipoEspecial, valor.TipoDelArray) {
					errorTipos = true
					tipoEntrante = Ast.TipoRetornado{
						Tipo:  Ast.ARRAY,
						Valor: valor.TipoDelArray,
					}
				}

			case Ast.STRUCT:
				valor := newValue.Valor.(StructInstancia)
				nTipo := Ast.TipoRetornado{
					Tipo:  Ast.STRUCT,
					Valor: valor.GetPlantilla(scope),
				}
				if !CompararTipos(simboloAtributo.TipoEspecial, nTipo) {
					errorTipos = true
					tipoEntrante = nTipo
				}
			}

			if errorTipos {
				if newValue.Tipo != simboloAtributo.Tipo {
					fila := a.Valor.(Ast.Abstracto).GetFila()
					columna := a.Valor.(Ast.Abstracto).GetColumna()
					msg := "Semantic error, expected " + Tipo_String(simboloAtributo.TipoEspecial) +
						" found " + Tipo_String(tipoEntrante) +
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
				}
			}

		} else

		//Si los tipos son diferentes
		if newValue.Tipo != simboloAtributo.Tipo {
			fila := a.Valor.(Ast.Abstracto).GetFila()
			columna := a.Valor.(Ast.Abstracto).GetColumna()
			msg := "Semantic error, expected " + Ast.ValorTipoDato[simboloAtributo.Tipo] +
				" found " + Ast.ValorTipoDato[newValue.Tipo] +
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
		}
		//Actualizar el valor del atributo
		simboloAtributo.Valor = newValue

		//Actualizar el valor del símbolo en el scope

		structInstancia.Entorno.UpdateSimbolo(idAtributo, simboloAtributo)

	} else {
		idStruct = nombreStruct.(expresiones.Identificador).Valor
		idAtributo = nombreAtributo.(expresiones.Identificador).Valor

		if !scope.Exist_actual(idStruct) {
			fila := a.Valor.(Ast.Abstracto).GetFila()
			columna := a.Valor.(Ast.Abstracto).GetColumna()
			msg := "Semantic error, the element " + idStruct + " doesn't exist." +
				" -- Line: " + strconv.Itoa(fila) +
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
		}

		//Get el simbolo del struct
		simboloStruct := scope.GetSimbolo(idStruct)

		if simboloStruct.Tipo != Ast.STRUCT {
			fila := a.Valor.(Ast.Abstracto).GetFila()
			columna := a.Valor.(Ast.Abstracto).GetColumna()
			msg := "Semantic error, STRUCT expected " + Ast.ValorTipoDato[simboloStruct.Tipo] + " found." +
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
		}

		//Verificar mutabilidad del struct
		if !simboloStruct.Mutable {
			fila := a.Valor.(Ast.Abstracto).GetFila()
			columna := a.Valor.(Ast.Abstracto).GetColumna()
			msg := "Semantic error, can't modify a non-mutable STRUCT." +
				" -- Line: " + strconv.Itoa(fila) +
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
		}

		//Get el struc
		structInstancia := simboloStruct.Valor.(Ast.TipoRetornado).Valor.(StructInstancia)

		if !structInstancia.Entorno.Exist_actual(idAtributo) {
			fila := a.Valor.(Ast.Abstracto).GetFila()
			columna := a.Valor.(Ast.Abstracto).GetColumna()
			msg := "Semantic error, the element " + idAtributo + " doesn't exist." +
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
		}

		//Verificar que exista el atributo
		if !structInstancia.Entorno.Exist_actual(idAtributo) {
			fila := a.Valor.(Ast.Abstracto).GetFila()
			columna := a.Valor.(Ast.Abstracto).GetColumna()
			msg := "Semantic error, the element " + idAtributo + " doesn't exist." +
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
		}
		//Get el simbolo del atributo
		simboloAtributo := structInstancia.Entorno.GetSimbolo(idAtributo)

		//Get el valor nuevo del atributo
		newValue := a.Valor.(Ast.Expresion).GetValue(scope)

		if newValue.Tipo == Ast.ERROR {
			return newValue
		}

		//Si los tipos son avanzados
		if simboloAtributo.TipoEspecial.Tipo != Ast.INDEFINIDO {
			_, tipoParticular = newValue.Valor.(Ast.Abstracto).GetTipo()
			errorTipos := false
			var tipoEntrante Ast.TipoRetornado
			switch tipoParticular {
			case Ast.VECTOR:
				valor := newValue.Valor.(expresiones.Vector)
				if !CompararTipos(simboloAtributo.TipoEspecial, valor.TipoVector) {
					errorTipos = true
					tipoEntrante = Ast.TipoRetornado{
						Tipo:  Ast.VECTOR,
						Valor: valor.TipoVector,
					}
				}
			case Ast.ARRAY:
				valor := newValue.Valor.(expresiones.Array)
				if !CompararTipos(simboloAtributo.TipoEspecial, valor.TipoDelArray) {
					errorTipos = true
					tipoEntrante = Ast.TipoRetornado{
						Tipo:  Ast.ARRAY,
						Valor: valor.TipoDelArray,
					}
				}

			case Ast.STRUCT:
				valor := newValue.Valor.(StructInstancia)
				nTipo := Ast.TipoRetornado{
					Tipo:  Ast.STRUCT,
					Valor: valor.GetPlantilla(scope),
				}
				if !CompararTipos(simboloAtributo.TipoEspecial, nTipo) {
					errorTipos = true
					tipoEntrante = nTipo
				}
			}

			if errorTipos {
				if newValue.Tipo != simboloAtributo.Tipo {
					fila := a.Valor.(Ast.Abstracto).GetFila()
					columna := a.Valor.(Ast.Abstracto).GetColumna()
					msg := "Semantic error, expected " + Tipo_String(simboloAtributo.TipoEspecial) +
						" found " + Tipo_String(tipoEntrante) +
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
				}
			}

		} else
		//Si los tipos son diferentes
		if newValue.Tipo != simboloAtributo.Tipo {
			fila := a.Valor.(Ast.Abstracto).GetFila()
			columna := a.Valor.(Ast.Abstracto).GetColumna()
			msg := "Semantic error, expected " + Ast.ValorTipoDato[simboloAtributo.Tipo] +
				" found " + Ast.ValorTipoDato[newValue.Tipo] +
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
		}

		//Actualizar el valor del atributo
		simboloAtributo.Valor = newValue

		//Actualizar el valor del símbolo en el scope
		structInstancia.Entorno.UpdateSimbolo(idAtributo, simboloAtributo)

	}
	return Ast.TipoRetornado{
		Valor: true,
		Tipo:  Ast.EJECUTADO,
	}
}

func (op AsignacionAccesoStruct) GetFila() int {
	return op.Fila
}
func (op AsignacionAccesoStruct) GetColumna() int {
	return op.Columna
}

func (d AsignacionAccesoStruct) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, Ast.ASIGNACION
}
