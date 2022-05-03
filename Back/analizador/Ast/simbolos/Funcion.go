package simbolos

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"Back/analizador/instrucciones"
	"strconv"
	"strings"

	"github.com/colegno/arraylist"
)

type Funcion struct {
	Nombre        string
	Tipo          Ast.TipoDato
	Fila          int
	Columna       int
	Instrucciones *arraylist.List
	Publica       bool
	Parametros    *arraylist.List
	Retorno       Ast.TipoRetornado
	Entorno       *Ast.Scope
}

func NewFuncion(nombre string, tipo Ast.TipoDato, instrucciones *arraylist.List,
	parametros *arraylist.List, retorno Ast.TipoRetornado, publica bool, fila, columna int) Funcion {
	nF := Funcion{
		Nombre:        nombre,
		Tipo:          tipo,
		Fila:          fila,
		Columna:       columna,
		Instrucciones: instrucciones,
		Publica:       publica,
		Parametros:    parametros,
		Retorno:       retorno,
	}
	return nF
}

func (f Funcion) Run(scope *Ast.Scope) interface{} {
	/********************VARIABLES 3D***********************/
	var obj3d, obj3dValor Ast.O3D
	var codigo3d string
	var codigoFuncion string
	var saltoReturn string
	var saltoReturnExp string
	var valorReturn string
	var posicionReturn string
	var algunValorParaRetornar interface{}
	//Set la funcion actual en stack
	Ast.SetFuncionStack(f.Nombre)
	/*******************************************************/

	var actual interface{}
	var tipoGeneral interface{}
	var respuesta interface{}

	//Ejecutar instrucciones

	for i := 0; i < f.Instrucciones.Len(); i++ {
		actual = f.Instrucciones.GetValue(i)
		if actual != nil {
			tipoGeneral, _ = actual.(Ast.Abstracto).GetTipo()
		} else {
			continue
		}

		if tipoGeneral == Ast.INSTRUCCION {
			//Declarar variables globales
			respuesta = actual.(Ast.Instruccion).Run(scope)
			obj3dValor = respuesta.(Ast.TipoRetornado).Valor.(Ast.O3D)
			respuesta = obj3dValor.Valor
			if respuesta.(Ast.TipoRetornado).Tipo == Ast.ERROR {
				return respuesta
			}
			codigo3d += obj3dValor.Codigo
		} else if tipoGeneral == Ast.EXPRESION {
			respuesta = actual.(Ast.Expresion).GetValue(scope)
			obj3dValor = respuesta.(Ast.TipoRetornado).Valor.(Ast.O3D)
			respuesta = obj3dValor.Valor
			if respuesta.(Ast.TipoRetornado).Tipo == Ast.ERROR {
				return respuesta
			}
			codigo3d += obj3dValor.Codigo
		}

		if obj3dValor.SaltoReturn != "" {
			if !obj3dValor.TranferenciaAgregada {
				saltoTemp := strings.Replace(obj3dValor.SaltoReturn, ",", "", -1)
				codigo3d += "goto " + saltoTemp + ";\n"
			}
			if !obj3dValor.TranferenciaAgregada {
				if obj3dValor.SaltoReturn[len(obj3dValor.SaltoReturn)-1] != ',' {
					saltoReturn += obj3dValor.SaltoReturn + ","
				} else {
					saltoReturn += obj3dValor.SaltoReturn
				}
			} else {
				saltoReturn += obj3dValor.SaltoReturn
			}
			//saltoReturn = strings.Replace(saltoReturn, ",", ":\n", -1)
			saltoReturnExp += obj3dValor.SaltoReturnExp
			algunValorParaRetornar = respuesta
			//saltoReturn = strings.Replace(saltoReturn, ",", ":\n", -1)

		}

		//scope.UpdateReferencias()
		/*
			if Ast.EsTransferencia(respuesta.(Ast.TipoRetornado).Tipo) {
				if respuesta.(Ast.TipoRetornado).Tipo == Ast.BREAK ||
					respuesta.(Ast.TipoRetornado).Tipo == Ast.BREAK_EXPRESION ||
					respuesta.(Ast.TipoRetornado).Tipo == Ast.CONTINUE {
					//Error de break, break_expresion y continue
					valor := actual.(Ast.Abstracto)
					fila := valor.GetFila()
					columna := valor.GetColumna()
					msg := "Semantic error," + Ast.ValorTipoDato[respuesta.(Ast.TipoRetornado).Tipo] +
						" statement not allowed outside a loop." +
						" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
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
				if f.Retorno.Tipo == Ast.VOID && respuesta.(Ast.TipoRetornado).Tipo == Ast.RETURN_EXPRESION {
					//Error de break, break_expresion
					valor := actual.(Ast.Abstracto)
					fila := valor.GetFila()
					columna := valor.GetColumna()
					msg := "Semantic error, this function can't return a value." +
						" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
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
				if f.Retorno.Tipo != Ast.VOID && respuesta.(Ast.TipoRetornado).Tipo == Ast.RETURN {
					//Error, la función espera retornar algo y no esta retornando nada
					valor := actual.(Ast.Abstracto)
					fila := valor.GetFila()
					columna := valor.GetColumna()
					msg := "Semantic error, this function is not returning any value." +
						" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
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

				if f.Retorno.Tipo != Ast.VOID && respuesta.(Ast.TipoRetornado).Tipo == Ast.RETURN_EXPRESION {
					//Verificar que los tipos sean correctos
					var tipoEntrante Ast.TipoRetornado
					var valorRespueta Ast.TipoRetornado = respuesta.(Ast.TipoRetornado).Valor.(Ast.TipoRetornado)
					if valorRespueta.Tipo == Ast.STRUCT {
						tipoEntrante.Tipo = Ast.STRUCT
						tipoEntrante.Valor = valorRespueta.Valor.(StructInstancia).GetPlantilla(scope)
					} else {
						if valorRespueta.Tipo == Ast.VECTOR {
							tipoEntrante = Ast.TipoRetornado{
								Tipo:  Ast.VECTOR,
								Valor: valorRespueta.Valor.(expresiones.Vector).TipoVector,
							}
						}
						if valorRespueta.Tipo == Ast.ARRAY {
							tipoEntrante = Ast.TipoRetornado{
								Tipo:  Ast.ARRAY,
								Valor: valorRespueta.Valor.(expresiones.Array).TipoDelArray,
							}
						}
					}
					if f.Retorno.Tipo == Ast.DIMENSION_ARRAY && tipoEntrante.Tipo == Ast.ARRAY {
						//Comparar las dimensiones solicitadas con el array de salida
						//Tengo que comparar el dimension de retorno y el valorRespuesta del array
						resultado := instrucciones.CompararDimensiones(f.Retorno.Valor.(expresiones.DimensionArray),
							valorRespueta.Valor.(expresiones.Array), scope)
						if resultado.Tipo == Ast.ERROR {
							return resultado
						}

					} else if !CompararTipos(f.Retorno, tipoEntrante) {
						//Error, retorna un tipo diferente
						valor := actual.(Ast.Abstracto)
						fila := valor.GetFila()
						columna := valor.GetColumna()
						msg := "Semantic error, expected" + expresiones.Tipo_String(f.Retorno) + " found " +
							expresiones.Tipo_String(respuesta.(Ast.TipoRetornado).Valor.(Ast.TipoRetornado)) +
							" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
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
					//Ejecutar el return y retornar el valor que trae
					return respuesta.(Ast.TipoRetornado).Valor.(Ast.TipoRetornado)
				}
			}
		*/
	}

	//Verificar que la funcion no sea de return y no este retornando nada
	/*
		if f.Retorno.Tipo != Ast.VOID {
			//Error la funcion debe retornar algo
			valor := actual.(Ast.Abstracto)
			fila := valor.GetFila()
			columna := valor.GetColumna()
			msg := "Semantic error, expected " + expresiones.Tipo_String(f.Retorno) + " , found ()." +
				" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
			nError := errores.NewError(fila, columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			nError.Ambito = scope.GetTipoScope()
			scope.Errores.Add(nError)
			scope.Consola += msg + "\n"
		}
	*/

	if saltoReturn != "" && saltoReturnExp == "" {
		if saltoReturn[len(saltoReturn)-1] != ',' {
			saltoReturn += ","
		}
		saltoReturn = strings.Replace(saltoReturn, ",", ":\n", -1)
		codigo3d += saltoReturn
	}

	if saltoReturnExp != "" && len(saltoReturnExp) >= 2 {
		posicionReturn = Ast.GetTemp()
		valorReturn = Ast.GetTemp()
		if saltoReturnExp[len(saltoReturnExp)-1] != ',' {
			saltoReturnExp += ","
		}
		saltoReturnExp = strings.Replace(saltoReturnExp, ",", ":\n", -1)
		codigo3d += saltoReturnExp
		//Recuperar el valor de la variable que se va a retornar
		codigo3d += "/*****************RECUPERAR EL VALOR DEL RETURN*/\n"
		codigo3d += posicionReturn + " = P + 0; //Posicion return \n"
		codigo3d += valorReturn + " = stack[(int)" + posicionReturn + "]; //Valor return \n"
		codigo3d += "/***********************************************/\n"
	}

	//Crear el código de la función
	codigoFuncion += "void " + f.Nombre + "(){\n"
	codigoFuncion += Ast.Indentar(scope.GetNivel(), codigo3d)
	codigoFuncion += Ast.Indentar(scope.GetNivel(), "return; \n")
	codigoFuncion += "}\n"
	codigo3d = f.Nombre + "();\n"

	if !strings.Contains(Ast.GetFuncionesC3D(), "void "+f.Nombre+"(){\n") {
		Ast.AgregarFuncion(codigoFuncion)
	}

	obj3d.Valor = Ast.TipoRetornado{
		Tipo:  Ast.EJECUTADO,
		Valor: true,
	}
	obj3d.Codigo = codigo3d
	if saltoReturnExp != "" && len(saltoReturnExp) > 2 {
		obj3d.Valor = Ast.TipoRetornado{
			Tipo:  f.Retorno.Tipo,
			Valor: algunValorParaRetornar.(Ast.TipoRetornado).Valor,
		}
		obj3d.Referencia = valorReturn
		return Ast.TipoRetornado{
			Tipo:  f.Retorno.Tipo,
			Valor: obj3d,
		}
	}

	return Ast.TipoRetornado{
		Tipo:  Ast.EJECUTADO,
		Valor: obj3d,
	}
}

func (f Funcion) RunParametros(scope *Ast.Scope, scopeOrigen *Ast.Scope, parametrosIN *arraylist.List) Ast.TipoRetornado {
	//var tipos Ast.TipoRetornado             //Variable para verificar que los tipos son correctos
	var parametrosCreados Ast.TipoRetornado //Variaable para verificar que los parametros fueron creados
	// Primero revisar que la cantidad de parámetros sea la misma
	if parametrosIN.Len() != f.Parametros.Len() {
		//Error, la cantidad de parámetros no es la esperada
		////////////////////////////ERROR//////////////////////////////////
		return errores.GenerarError(53, f, f, "",
			"",
			"",
			scope)
		//////////////////////////////////////////////////////////////////
	}

	// Revisar que los tipos sean correctos
	/*
		tipos = TiposCorrectos(scope, f.Parametros, parametrosIN)
		if tipos.Tipo != Ast.BOOLEAN {
			return tipos
		}
	*/

	// crear los parámetros
	parametrosCreados = CrearParametros(scope, scopeOrigen, f.Parametros, parametrosIN)

	return parametrosCreados

}

func TiposCorrectos(scope *Ast.Scope, parametros, parametrosIN *arraylist.List) Ast.TipoRetornado {
	var iterador int
	var resultadoParametroIN, resultadoParametro Ast.TipoRetornado
	var parametroIN, parametro interface{}
	for iterador = 0; iterador < parametros.Len(); iterador++ {
		parametro = parametros.GetValue(iterador)
		parametroIN = parametrosIN.GetValue(iterador)
		resultadoParametroIN = parametroIN.(Ast.Expresion).GetValue(scope)
		resultadoParametro = parametro.(Ast.Expresion).GetValue(scope)

		//Verificar errores en los resultados
		if resultadoParametroIN.Tipo == Ast.ERROR {
			return resultadoParametroIN
		}
		if resultadoParametro.Tipo == Ast.ERROR {
			return resultadoParametro
		}

		if resultadoParametroIN.Tipo != resultadoParametro.Tipo {
			//Error no son iguales los tipos
			fila := parametroIN.(Ast.Abstracto).GetFila()
			columna := parametroIN.(Ast.Abstracto).GetColumna()
			msg := "Semantic error, " + Ast.ValorTipoDato[resultadoParametro.Tipo] +
				" type expected, " + Ast.ValorTipoDato[resultadoParametroIN.Tipo] + " type found." +
				" type. -- Line: " + strconv.Itoa(fila) +
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

		//Verificar la mutabilidad en los valores que se mandan por valor
		_, tipoParticular := parametroIN.(Valor).Valor.(Ast.Abstracto).GetTipo()
		if tipoParticular == Ast.IDENTIFICADOR {

			//Verificar la mutabilidad en los valores que se mandan por referencia

			//Conseguir el símbolo del identificador
			simbolo := scope.GetSimbolo(parametroIN.(Valor).Valor.(expresiones.Identificador).Valor)
			//Primero verificar que el símbolo sea un struct, un vector o un array
			if simbolo.Tipo == Ast.STRUCT ||
				simbolo.Tipo == Ast.ARRAY ||
				simbolo.Tipo == Ast.VECTOR {
				if parametro.(Parametro).Mutable != parametroIN.(Valor).Mutable {
					fila := parametroIN.(Ast.Abstracto).GetFila()
					columna := parametroIN.(Ast.Abstracto).GetColumna()
					var mut1, mut2 string
					if parametro.(Parametro).Mutable {
						mut1 = "Mutable"
					} else {
						mut1 = "Not-Mutable"
					}
					if parametroIN.(Valor).Mutable {
						mut2 = "Mutable"
					} else {
						mut2 = "Not-Mutable"
					}
					msg := "Semantic error, " + mut1 + " value expected, " +
						mut2 + " value found." +
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
				if simbolo.Mutable != parametroIN.(Valor).Mutable {
					fila := parametroIN.(Ast.Abstracto).GetFila()
					columna := parametroIN.(Ast.Abstracto).GetColumna()
					var mut1, mut2 string
					if simbolo.Mutable {
						mut1 = "Mutable"
					} else {
						mut1 = "Not-Mutable"
					}
					if parametroIN.(Valor).Mutable {
						mut2 = "Mutable"
					} else {
						mut2 = "Not-Mutable"
					}
					msg := "Semantic error, " + mut2 + " value expected, " +
						mut1 + " value found." +
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

			} else if simbolo.Tipo == Ast.MODULO {
				fila := parametroIN.(Ast.Abstracto).GetFila()
				columna := parametroIN.(Ast.Abstracto).GetColumna()
				msg := "Semantic error, a module can't be a parameter." +
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
			} else

			//Es cualquier otra variable
			if parametroIN.(Valor).Mutable {
				// Error, las variables no pueden ser mut
				fila := parametroIN.(Ast.Abstracto).GetFila()
				columna := parametroIN.(Ast.Abstracto).GetColumna()
				msg := "Semantic error, " + Ast.ValorTipoDato[simbolo.Tipo] + " variable can´t be mut." +
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

		}

	}
	return Ast.TipoRetornado{
		Tipo:  Ast.BOOLEAN,
		Valor: true,
	}
}

func CrearParametros(scope *Ast.Scope, scopeOrigen *Ast.Scope, parametros, parametrosIN *arraylist.List) Ast.TipoRetornado {
	/**********************VARIABLES 3d*************************/
	var obj3dParametroIn, obj3dparametro, obj3dValor, obj3d Ast.O3D
	var codigo3d string
	/***********************************************************/

	var iterador int
	var resultadoParametro, tipoParametro, tipoParametroIN Ast.TipoRetornado
	var parametroIN, parametro, resultadoDeclaracion interface{}
	var paramTemp Parametro
	var verificacionDeReferencia Ast.TipoRetornado

	/***********************REINICIO EL SIZE DEL SCOPE*********/
	//scope.Size = 1
	/**********************************************************/
	for iterador = 0; iterador < parametros.Len(); iterador++ {
		parametro = parametros.GetValue(iterador)
		tipoParametro = parametro.(Parametro).FormatearTipo(scope)
		paramTemp = parametro.(Parametro)
		paramTemp.TipoDeclaracion = tipoParametro
		parametro = paramTemp
		parametroIN = parametrosIN.GetValue(iterador)
		/***************************************************************/
		tipoParametroIN = parametroIN.(Ast.Expresion).GetValue(scopeOrigen)
		obj3dParametroIn = tipoParametroIN.Valor.(Ast.O3D)
		tipoParametroIN = obj3dParametroIn.Valor
		/***************************************************************/
		resultadoParametro = parametro.(Ast.Expresion).GetValue(scope)
		obj3dparametro = resultadoParametro.Valor.(Ast.O3D)
		resultadoParametro = obj3dparametro.Valor
		/***************************************************************/

		if resultadoParametro.Tipo == Ast.ERROR {
			return resultadoParametro
		}
		if tipoParametroIN.Tipo == Ast.ERROR {
			return tipoParametroIN
		}
		//Crear un objeto declaración
		//Obtener el tipo del parámetro
		/*
			nuevaDeclaracion := instrucciones.NewDeclaracion(resultadoParametro.Valor.(string),
				resultadoParametro.Tipo, parametro.(Parametro).Mutable, false, Ast.VOID, parametroIN,
				parametro.(Ast.Abstracto).GetFila(), parametro.(Ast.Abstracto).GetColumna())
		*/
		//Verificar si es referencia y que todo este correcto
		verificacionDeReferencia = VerificarReferencia(parametro, parametroIN, scope, scopeOrigen)
		//Verificar algún posible error
		if verificacionDeReferencia.Tipo == Ast.ERROR {
			return verificacionDeReferencia
		}

		if parametroIN.(Valor).Referencia || !expresiones.EsVAS(tipoParametroIN.Tipo) {
			if tipoParametro.Tipo == Ast.DIMENSION_ARRAY {
				nuevaDeclaracion := instrucciones.NewDeclaracionArray(
					resultadoParametro.Valor.(string), tipoParametro.Valor, parametro.(Parametro).Mutable,
					false, parametroIN, parametro.(Ast.Abstracto).GetFila(),
					parametro.(Ast.Abstracto).GetColumna())
				//Ejecutar declaración
				nuevaDeclaracion.ScopeOriginal = scopeOrigen
				resultadoDeclaracion = nuevaDeclaracion.Run(scope)
			} else if tipoParametro.Tipo == Ast.VECTOR {
				nuevaDeclaracion := instrucciones.NewDeclaracionVector(resultadoParametro.Valor.(string), tipoParametro.Valor.(Ast.TipoRetornado), parametroIN,
					parametro.(Parametro).Mutable, false, parametro.(Ast.Abstracto).GetFila(),
					parametro.(Ast.Abstracto).GetColumna())
				//Ejecutar declaración
				nuevaDeclaracion.ScopeOriginal = scopeOrigen
				resultadoDeclaracion = nuevaDeclaracion.Run(scope)
			} else {
				nuevaDeclaracion := instrucciones.NewDeclaracionTotal(resultadoParametro.Valor.(string), parametroIN, tipoParametro,
					parametro.(Parametro).Mutable, false, parametro.(Ast.Abstracto).GetFila(),
					parametro.(Ast.Abstracto).GetColumna())
				//Ejecutar declaración
				nuevaDeclaracion.ScopeOriginal = scopeOrigen
				resultadoDeclaracion = nuevaDeclaracion.Run(scope)

			}
			obj3dValor = resultadoDeclaracion.(Ast.TipoRetornado).Valor.(Ast.O3D)
			if obj3dValor.Valor.Tipo == Ast.ERROR {
				return Ast.TipoRetornado{
					Tipo:  obj3dValor.Valor.Tipo,
					Valor: obj3dValor,
				}
			}
			resultadoDeclaracion = obj3dValor.Valor
			codigo3d += obj3dValor.Codigo

		} else if expresiones.EsVAS(tipoParametroIN.Tipo) {

			if tipoParametro.Tipo == Ast.DIMENSION_ARRAY {
				nuevaDeclaracion := instrucciones.NewDeclaracionArrayNoRef(
					resultadoParametro.Valor.(string), tipoParametro.Valor, parametro.(Parametro).Mutable,
					false, parametroIN, parametro.(Ast.Abstracto).GetFila(),
					parametro.(Ast.Abstracto).GetColumna())
				//Ejecutar declaración
				nuevaDeclaracion.ScopeOriginal = scopeOrigen
				resultadoDeclaracion = nuevaDeclaracion.Run(scope)

			} else if tipoParametro.Tipo == Ast.VECTOR {
				nuevaDeclaracion := instrucciones.NewDeclaracionVectorNoRef(resultadoParametro.Valor.(string), tipoParametro.Valor.(Ast.TipoRetornado), parametroIN,
					parametro.(Parametro).Mutable, false, parametro.(Ast.Abstracto).GetFila(),
					parametro.(Ast.Abstracto).GetColumna())
				//Ejecutar declaración
				nuevaDeclaracion.ScopeOriginal = scopeOrigen
				resultadoDeclaracion = nuevaDeclaracion.Run(scope)
			} else {
				nuevaDeclaracion := instrucciones.NewDeclaracionNoRef(resultadoParametro.Valor.(string), parametroIN, tipoParametro,
					parametro.(Parametro).Mutable, false, parametro.(Ast.Abstracto).GetFila(),
					parametro.(Ast.Abstracto).GetColumna())
				//Ejecutar declaración
				nuevaDeclaracion.ScopeOriginal = scopeOrigen
				resultadoDeclaracion = nuevaDeclaracion.Run(scope)
			}
			obj3dValor = resultadoDeclaracion.(Ast.TipoRetornado).Valor.(Ast.O3D)
			if obj3dValor.Valor.Tipo == Ast.ERROR {
				return Ast.TipoRetornado{
					Tipo:  obj3dValor.Valor.Tipo,
					Valor: obj3dValor,
				}
			}
			resultadoDeclaracion = obj3dValor.Valor
			codigo3d += obj3dValor.Codigo
		}

		if resultadoDeclaracion.(Ast.TipoRetornado).Tipo == Ast.ERROR {
			return resultadoDeclaracion.(Ast.TipoRetornado)
		}
		//De forma sucia asignar la referencia si el valor es por referencia
		if parametroIN.(Valor).Referencia {
			simbolo := scope.GetSimbolo(resultadoParametro.Valor.(string))
			referencia := parametroIN.(Valor).Valor.(expresiones.Identificador).Valor
			simboloRef := scope.GetSimboloReferencia(referencia)
			simbolo.Referencia = true
			simbolo.Referencia_puntero = &simboloRef
			scope.UpdateSimbolo(resultadoParametro.Valor.(string), simbolo)
		}

	}
	obj3d.Valor = Ast.TipoRetornado{
		Tipo:  Ast.BOOLEAN,
		Valor: true,
	}
	obj3d.Codigo = codigo3d
	return Ast.TipoRetornado{
		Tipo:  Ast.BOOLEAN,
		Valor: obj3d,
	}
}

func (f Funcion) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, f.Tipo
}

func (f Funcion) GetFila() int {
	return f.Fila
}
func (f Funcion) GetColumna() int {
	return f.Columna
}

func (f Funcion) GetTipoRetornado(scope *Ast.Scope) string {
	tipo := GetTipoEstructura(f.Retorno, scope, f)
	cadena := Tipo_String(tipo)
	return cadena
}

func EsAVelementos(tipo Ast.TipoDato) bool {
	switch tipo {
	case Ast.DIMENSION_ARRAY:
		return true
	default:
		return false
	}
}

func GetValorPredeterminado(tipo Ast.TipoDato) interface{} {
	switch tipo {
	case Ast.I64:
		return 0
	case Ast.F64:
		return 1.1
	case Ast.BOOLEAN:
		return true
	case Ast.CHAR:
		return "a"
	case Ast.STRING, Ast.STR:
		return "cadena"
	default:
		return true
	}

}

func GetValorPredeterminadoTipoRet(tipo Ast.TipoDato) Ast.TipoRetornado {
	var retorno Ast.TipoRetornado
	switch tipo {
	case Ast.I64:
		retorno.Valor = 0
		retorno.Tipo = Ast.I64
	case Ast.F64:
		retorno.Valor = 1.1
		retorno.Tipo = Ast.F64
	case Ast.BOOLEAN:
		retorno.Valor = true
		retorno.Tipo = Ast.BOOLEAN
	case Ast.CHAR:
		retorno.Valor = "a"
		retorno.Tipo = Ast.CHAR
	case Ast.STRING, Ast.STR:
		retorno.Valor = "cadena"
		retorno.Tipo = tipo
	default:
		retorno.Valor = true
		retorno.Tipo = Ast.EJECUTADO
	}
	return retorno
}
