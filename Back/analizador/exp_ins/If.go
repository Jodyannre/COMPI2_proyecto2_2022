package exp_ins

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"strconv"

	"github.com/colegno/arraylist"
)

type IF struct {
	Condicion     Ast.Expresion
	Instrucciones *arraylist.List
	Lista_if_else *arraylist.List
	Tipo          Ast.TipoDato
	Fila          int
	Columna       int
	Expresion     bool
}

func NewIF(condicion Ast.Expresion, instrucciones *arraylist.List, lista *arraylist.List, tipo Ast.TipoDato,
	fila, columna int, expresion bool) IF {
	nif := IF{
		Condicion:     condicion,
		Instrucciones: instrucciones,
		Lista_if_else: lista,
		Tipo:          tipo,
		Fila:          fila,
		Columna:       columna,
		Expresion:     expresion,
	}
	return nif
}

func (i IF) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, i.Tipo
}

func (i IF) Run(scope *Ast.Scope) interface{} {
	/*************************OBJETOS 3D******************************/
	var obj3dTransferencia Ast.O3D
	var obj3d Ast.O3D
	/****************************************************************/
	//Crear el nuevo scope
	newScope := Ast.NewScope("IF_I", scope)
	newScope.Posicion = scope.Size + scope.Posicion
	//Aumento el size para apartar un espacio para el posible valor retornado
	newScope.Size++
	//Inicializar la lista de respuestas
	//Ejecutar la instrucción if
	resultado := GetResultado3D(i, &newScope, -1, i.Expresion)
	obj3d = resultado.Valor.(Ast.O3D)
	if Ast.EsTransferencia(resultado.Tipo) || obj3d.SaltoContinue != "" ||
		obj3d.SaltoBreak != "" || obj3d.SaltoReturn != "" {
		obj3dTransferencia = resultado.Valor.(Ast.O3D)
		//saltos := obj3dTransferencia.SaltoTranferencia
		//saltos = strings.Replace(saltos, ",", ":\n", -1)
		//saltos += ":\n"
		//obj3dTransferencia.SaltoTranferencia = saltos
		obj3dTransferencia.TranferenciaAgregada = true
		obj3dTransferencia.Valor.Tipo = resultado.Tipo

		resultado.Valor = obj3dTransferencia
	} else {
		obj3dTransferencia = resultado.Valor.(Ast.O3D)
		if obj3dTransferencia.RetornoIf != "" {
			//RECUPERAR EL POSIBLE VALOR GUARDADO DEL BREAK
			entornoAnterior := Ast.GetTemp()
			posicionBreak := Ast.GetTemp()
			valorIf := Ast.GetTemp()
			codigo3d := ""
			//Recuperar el valor de la variable que se va a retornar
			codigo3d += "/*********RECUPERAR EL VALOR DEL RETORNO DEL IF*/\n"
			codigo3d += entornoAnterior + " = P; //guardar entorno anterior \n"
			codigo3d += "P = P + " + strconv.Itoa(scope.Size) + "; //Cambio de entorno \n"
			codigo3d += posicionBreak + " = P + 0; //Posicion valor \n"
			codigo3d += valorIf + " = stack[(int)" + posicionBreak + "]; //Valor del retorno \n"
			codigo3d += "P = P - " + entornoAnterior + "; //Regresar al entorno anterior \n"
			codigo3d += "/***********************************************/\n"
			obj3dTransferencia.Codigo += codigo3d
			obj3dTransferencia.Referencia = valorIf
			resultado.Valor = obj3dTransferencia
		}
	}

	//actualizar el scope global con los resultados
	newScope.UpdateScopeGlobal()
	return resultado
}

func GetResultado(i IF, scope *Ast.Scope, pos int, expresion bool) Ast.TipoRetornado {
	/*******************************VARIABLES 3D**********************************/
	var obj3dCondicion Ast.O3D
	//var codigo3d string

	/*****************************************************************************/

	var condicion1 Ast.TipoRetornado
	var resultado Ast.TipoRetornado
	var ultimoTipo Ast.TipoDato
	var instruccion interface{}
	if pos == -1 {
		condicion1 = i.Condicion.GetValue(scope)
		obj3dCondicion = condicion1.Valor.(Ast.O3D)
		condicion1 = obj3dCondicion.Valor
	} else {
		//Conseguir el if/entonces
		//elemento := i.Lista_if_else.GetValue(pos).(IF)
		//Evaluar si es if o entonces
		if i.Tipo == Ast.ELSEIF || i.Tipo == Ast.ELSEIF_EXPRESION {
			//Es un if
			//i = elemento
			scope.Nombre = "ELSEIF"
			condicion1 = i.Condicion.GetValue(scope)
		} else if i.Tipo == Ast.ELSE || i.Tipo == Ast.ELSE_EXPRESION {
			//Es un else
			//i = elemento
			scope.Nombre = "ELSE"
			condicion1 = Ast.TipoRetornado{
				Tipo:  Ast.BOOLEAN,
				Valor: true,
			}
		}
	}

	if i.Tipo == Ast.IF_EXPRESION ||
		i.Tipo == Ast.ELSE_EXPRESION || i.Tipo == Ast.ELSEIF_EXPRESION {
		expresion = true
	}

	if condicion1.Tipo == Ast.BOOLEAN {
		if condicion1.Valor.(bool) {
			//Es verdadera; ejecutar las instrucciones
			n := 0

			for n < i.Instrucciones.Len() {

				//Verificar que la instrucción no sea null
				if i.Instrucciones.GetValue(n) == nil {
					n++
					continue
				}

				instruccion = i.Instrucciones.GetValue(n).(Ast.Abstracto)
				tipo_abstracto, _ := instruccion.(Ast.Abstracto).GetTipo()
				ultimoTipo = tipo_abstracto
				if tipo_abstracto == Ast.EXPRESION {
					instruccion = i.Instrucciones.GetValue(n)
					resultado = instruccion.(Ast.Expresion).GetValue(scope)

				} else if tipo_abstracto == Ast.INSTRUCCION {
					instruccion = i.Instrucciones.GetValue(n)
					resultado = instruccion.(Ast.Instruccion).Run(scope).(Ast.TipoRetornado)
				}

				if Ast.EsTransferencia(resultado.Tipo) &&
					expresion {
					temp := instruccion.(Ast.Abstracto)
					msg := "Semantic error, transfer statements are not allowed within a if expression statement." +
						" -- Line:" + strconv.Itoa(temp.GetFila()) + " Column: " + strconv.Itoa(temp.GetColumna())
					nError := errores.NewError(temp.GetFila(), temp.GetColumna(), msg)
					nError.Tipo = Ast.ERROR_SEMANTICO
					nError.Ambito = scope.GetTipoScope()
					scope.Errores.Add(nError)
					scope.Consola += msg + "\n"
					scope.UpdateScopeGlobal()
					return Ast.TipoRetornado{
						Valor: nError,
						Tipo:  Ast.ERROR,
					}
				}
				if Ast.EsTransferencia(resultado.Tipo) {
					//Si es transferencia, terminar con el if y retornarlo
					return resultado
				}
				n++
			}
			//Termino el for, retornar la ultima expresion
			//Verificar si hay algun retorno o retornar un error
			if ultimoTipo != Ast.EXPRESION && expresion {
				msg := "Semantic error, the if clause is not returning any value." +
					" -- Line:" + strconv.Itoa(i.Fila) + " Column: " + strconv.Itoa(i.Columna)
				nError := errores.NewError(i.Fila, i.Columna, msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				nError.Ambito = scope.GetTipoScope()
				scope.Errores.Add(nError)
				scope.Consola += msg + "\n"
				return Ast.TipoRetornado{
					Valor: nError,
					Tipo:  Ast.ERROR,
				}

			} else if expresion && ultimoTipo == Ast.EXPRESION {
				//Si esta retornado algún valor
				return resultado
			}

		} else {
			//Es falsa, buscar en la lista si hay otras
			//Llamada recursiva o fin
			//Recorrer la lista de ifs y else
			for j := 0; j < i.Lista_if_else.Len(); j++ {
				newScope := Ast.NewScope("if", scope)
				resultado := GetResultado(i.Lista_if_else.GetValue(j).(IF), &newScope, 0, expresion)
				if resultado.Tipo == Ast.EJECUTADO && resultado.Valor == true && !expresion {
					newScope.UpdateScopeGlobal()
					return Ast.TipoRetornado{
						Valor: true,
						Tipo:  Ast.EJECUTADO,
					}
				}
				if resultado.Tipo == Ast.ERROR {
					//newScope.Errores.Add(resultado.Valor)
					//newScope.Consola += resultado.Valor.(errores.CustomSyntaxError).Msg + "\n"
					newScope.UpdateScopeGlobal()
					return Ast.TipoRetornado{
						Valor: resultado.Valor,
						Tipo:  Ast.ERROR,
					}
				}

				if Ast.EsTransferencia(resultado.Tipo) {
					newScope.UpdateScopeGlobal()
					return resultado
				}

				if resultado.Tipo != Ast.EJECUTADO && (i.Tipo == Ast.IF_EXPRESION ||
					i.Tipo == Ast.ELSE_EXPRESION || i.Tipo == Ast.ELSEIF_EXPRESION) {
					scope.UpdateScopeGlobal()
					return resultado
				}

			}
			return Ast.TipoRetornado{
				Valor: false,
				Tipo:  Ast.EJECUTADO,
			}
		}
	} else {
		//No es booleano, entonces generar un error semántico
		//fmt.Println("Error semántico, la expresión no es un booleano")
		msg := "Semantic error, the condition of the expression is not a boolean expression." +
			" -- Line:" + strconv.Itoa(i.Fila) + " Column: " + strconv.Itoa(i.Columna)
		nError := errores.NewError(i.Fila, i.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}
	//Se acabo todo y retornar un true de finalizado
	return Ast.TipoRetornado{
		Valor: true,
		Tipo:  Ast.EJECUTADO,
	}
}

func GetResultado3D(i IF, scope *Ast.Scope, pos int, expresion bool) Ast.TipoRetornado {
	/*******************************VARIABLES 3D**********************************/
	var obj3dCondicion, objResultadoInstruccion, objResultadoIfs, obj3d Ast.O3D
	var codigo3d string
	var salto string
	var saltosFin string
	var ultimoFalso string
	var resultadoTranferencia Ast.TipoRetornado
	var hayTranferencia bool = false
	var hayRetornoIf bool = false
	//var saltosTransferencia string
	var saltosContinue string
	var saltosBreak string
	var saltosBreakExp string
	var saltosReturn string
	var saltoReturnExp string
	var guardarScope string
	var posicionGuardar string
	var scopeOrigen *Ast.Scope
	var posCorrectaNoElse bool = false
	/*****************************************************************************/
	codigo3d += "/********************************CONDICIONAL IF*/\n"
	var condicion1 Ast.TipoRetornado
	var resultado Ast.TipoRetornado
	var ultimoTipo Ast.TipoDato
	var instruccion interface{}
	if pos == -1 {
		var condicionTemp interface{} = i.Condicion
		_, tipoParticular := condicionTemp.(Ast.Abstracto).GetTipo()
		condicion1 = i.Condicion.GetValue(scope)
		obj3dCondicion = condicion1.Valor.(Ast.O3D)
		condicion1 = obj3dCondicion.Valor
		codigo3d += obj3dCondicion.Codigo
		if tipoParticular == Ast.IDENTIFICADOR {
			obj3dCondicion = GenerarCod3DCondicionEspecial(obj3dCondicion.Referencia)
			codigo3d += obj3dCondicion.Codigo
		}

	} else {
		//Conseguir el if/entonces
		//elemento := i.Lista_if_else.GetValue(pos).(IF)
		//Evaluar si es if o entonces
		if i.Tipo == Ast.ELSEIF || i.Tipo == Ast.ELSEIF_EXPRESION {
			//Es un if
			//i = elemento
			scope.Nombre = "ELSEIF"
			var condicionTemp interface{} = i.Condicion
			_, tipoParticular := condicionTemp.(Ast.Abstracto).GetTipo()
			condicion1 = i.Condicion.GetValue(scope)
			obj3dCondicion = condicion1.Valor.(Ast.O3D)
			condicion1 = obj3dCondicion.Valor
			codigo3d += obj3dCondicion.Codigo
			if tipoParticular == Ast.IDENTIFICADOR {
				obj3dCondicion = GenerarCod3DCondicionEspecial(obj3dCondicion.Referencia)
				codigo3d += obj3dCondicion.Codigo
			}
		} else if i.Tipo == Ast.ELSE || i.Tipo == Ast.ELSE_EXPRESION {
			//Es un else
			//i = elemento
			scope.Nombre = "ELSE"
			condicion1 = Ast.TipoRetornado{
				Tipo:  Ast.BOOLEAN,
				Valor: true,
			}
		}
	}

	if i.Tipo == Ast.IF_EXPRESION ||
		i.Tipo == Ast.ELSE_EXPRESION || i.Tipo == Ast.ELSEIF_EXPRESION {
		expresion = true
	}

	if condicion1.Tipo == Ast.BOOLEAN {
		////////////////////////////VERDADERAS//////////////////////////////////////
		//Es verdadera; ejecutar las instrucciones
		n := 0
		if obj3dCondicion.Lt != "" {
			codigo3d += obj3dCondicion.Lt + ":\n"
		}
		for n < i.Instrucciones.Len() {

			//Verificar que la instrucción no sea null
			if i.Instrucciones.GetValue(n) == nil {
				n++
				continue
			}

			instruccion = i.Instrucciones.GetValue(n).(Ast.Abstracto)
			tipo_abstracto, _ := instruccion.(Ast.Abstracto).GetTipo()
			ultimoTipo = tipo_abstracto
			if tipo_abstracto == Ast.EXPRESION {
				//instruccion = i.Instrucciones.GetValue(n)
				resultado = instruccion.(Ast.Expresion).GetValue(scope)
				objResultadoInstruccion = resultado.Valor.(Ast.O3D)
				resultado = objResultadoInstruccion.Valor

			} else if tipo_abstracto == Ast.INSTRUCCION {
				//instruccion = i.Instrucciones.GetValue(n)
				resultado = instruccion.(Ast.Instruccion).Run(scope).(Ast.TipoRetornado)
				objResultadoInstruccion = resultado.Valor.(Ast.O3D)
				resultado = objResultadoInstruccion.Valor
			}
			codigo3d += objResultadoInstruccion.Codigo

			if Ast.EsTransferencia(resultado.Tipo) &&
				expresion {
				temp := instruccion.(Ast.Abstracto)
				msg := "Semantic error, transfer statements are not allowed within a if expression statement." +
					" -- Line:" + strconv.Itoa(temp.GetFila()) + " Column: " + strconv.Itoa(temp.GetColumna())
				nError := errores.NewError(temp.GetFila(), temp.GetColumna(), msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				nError.Ambito = scope.GetTipoScope()
				scope.Errores.Add(nError)
				scope.Consola += msg + "\n"
				scope.UpdateScopeGlobal()
				return Ast.TipoRetornado{
					Valor: nError,
					Tipo:  Ast.ERROR,
				}
			}
			if Ast.EsTransferencia(resultado.Tipo) {
				//Si es transferencia, agregar el salto y guardar el resultado para retornarlo
				if condicion1.Valor == true {
					if !posCorrectaNoElse {
						resultadoTranferencia = resultado
						posCorrectaNoElse = true
					}
				}
				if !objResultadoInstruccion.TranferenciaAgregada {
					switch resultado.Tipo {
					case Ast.BREAK:
						codigo3d += "goto " + objResultadoInstruccion.SaltoBreak + ";\n"
						saltosBreak += objResultadoInstruccion.SaltoBreak + ","
						saltosBreakExp += objResultadoInstruccion.SaltoBreakExp + ","
						resultadoTranferencia = resultado
					case Ast.CONTINUE:
						codigo3d += "goto " + objResultadoInstruccion.SaltoContinue + ";\n"
						saltosContinue += objResultadoInstruccion.SaltoContinue + ","
						resultadoTranferencia = resultado
					case Ast.RETURN, Ast.RETURN_EXPRESION:
						codigo3d += "goto " + objResultadoInstruccion.SaltoReturn + ";\n"
						saltosReturn += objResultadoInstruccion.SaltoReturn + ","
						saltoReturnExp += objResultadoInstruccion.SaltoReturnExp + ","
						resultadoTranferencia = resultado
					}
					//saltosTransferencia += objResultadoInstruccion.SaltoTranferencia + ","
				} else {
					switch resultado.Tipo {
					case Ast.BREAK:
						saltosBreak += objResultadoInstruccion.SaltoBreak
						saltosBreakExp += objResultadoInstruccion.SaltoBreakExp
						resultadoTranferencia = resultado
					case Ast.CONTINUE:
						saltosContinue += objResultadoInstruccion.SaltoContinue
						resultadoTranferencia = resultado
					case Ast.RETURN, Ast.RETURN_EXPRESION:
						saltosReturn += objResultadoInstruccion.SaltoReturn
						saltoReturnExp += objResultadoInstruccion.SaltoReturnExp
						resultadoTranferencia = resultado
					}
					//saltosTransferencia += objResultadoInstruccion.SaltoTranferencia
				}
				hayTranferencia = true
				//return resultado
			} else if resultado.Tipo != Ast.EJECUTADO && (i.Tipo == Ast.IF_EXPRESION ||
				i.Tipo == Ast.ELSE_EXPRESION || i.Tipo == Ast.ELSEIF_EXPRESION) {
				scope.UpdateScopeGlobal()
				guardarScope = Ast.GetTemp()
				posicionGuardar = Ast.GetTemp()
				scopeOrigen = scope.GetEntornoPadreIF()
				resultadoTranferencia = resultado
				codigo3d += "/***********************GUARDAR EL VALOR DEL IF*/\n"
				codigo3d += guardarScope + " = P; //Guardar scope anterior \n"
				codigo3d += "P = " + strconv.Itoa(scopeOrigen.Posicion) + "; //Cambio a entorno simulado \n"
				codigo3d += posicionGuardar + " = P + 0; //Pos para el valor del if \n"
				codigo3d += "stack[(int)" + posicionGuardar + "] = " + objResultadoInstruccion.Referencia + "; //Guardar valor \n"
				codigo3d += "P = " + guardarScope + "; //Retornar al entorno anterior \n"
				codigo3d += "/***********************************************/\n"
				//return resultado
				hayRetornoIf = true
			}
			n++
		}
		//Termino el for, retornar la ultima expresion
		//Verificar si hay algun retorno o retornar un error
		if ultimoTipo != Ast.EXPRESION && expresion {
			msg := "Semantic error, the if clause is not returning any value." +
				" -- Line:" + strconv.Itoa(i.Fila) + " Column: " + strconv.Itoa(i.Columna)
			nError := errores.NewError(i.Fila, i.Columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			nError.Ambito = scope.GetTipoScope()
			scope.Errores.Add(nError)
			scope.Consola += msg + "\n"
			return Ast.TipoRetornado{
				Valor: nError,
				Tipo:  Ast.ERROR,
			}

		}

		/*else if expresion && ultimoTipo == Ast.EXPRESION {
			//Si esta retornado algún valor
			return resultado
		}
		*/

		////////////////////////////VERDADERAS//////////////////////////////////////
		if obj3dCondicion.Lf != "" {
			salto = Ast.GetLabel()
			codigo3d += "goto " + salto + ";\n"
			saltosFin += salto + ":\n" //Voy construyendo los saltos para salir de los ifs verdaderos
		}

		///////////////////////////////FALSAS//////////////////////////////////////
		//Es falsa, buscar en la lista si hay otras
		//Llamada recursiva o fin
		//Recorrer la lista de ifs y else
		ultimoFalso = obj3dCondicion.Lf
		faltaActual := obj3dCondicion.Lf
		for j := 0; j < i.Lista_if_else.Len(); j++ {
			codigo3d += faltaActual + ":\n"
			newScope := Ast.NewScope("if", scope)
			newScope.Posicion = scope.Size
			resultado := GetResultado3D(i.Lista_if_else.GetValue(j).(IF), &newScope, 0, expresion)
			objResultadoIfs = resultado.Valor.(Ast.O3D)
			codigo3d += objResultadoIfs.Codigo
			resultado = objResultadoIfs.Valor

			if resultado.Tipo == Ast.ERROR {
				//newScope.Errores.Add(resultado.Valor)
				//newScope.Consola += resultado.Valor.(errores.CustomSyntaxError).Msg + "\n"
				newScope.UpdateScopeGlobal()
				return Ast.TipoRetornado{
					Valor: resultado.Valor,
					Tipo:  Ast.ERROR,
				}
			}

			if Ast.EsTransferencia(resultado.Tipo) {
				newScope.UpdateScopeGlobal()
				//resultadoTranferencia = resultado
				if !posCorrectaNoElse && objResultadoIfs.IfCorrecto {
					resultadoTranferencia = resultado
					posCorrectaNoElse = true
				}

				switch resultado.Tipo {
				case Ast.BREAK:
					saltosBreak += objResultadoIfs.SaltoBreak
					resultadoTranferencia = resultado
				case Ast.CONTINUE:
					saltosContinue += objResultadoIfs.SaltoContinue
					resultadoTranferencia = resultado
				case Ast.RETURN:
					saltosReturn += objResultadoIfs.SaltoReturn
					saltoReturnExp += objResultadoIfs.SaltoReturnExp
					resultadoTranferencia = resultado
				}
				//saltosTransferencia += objResultadoIfs.SaltoBreak
				hayTranferencia = true
			}

			/*
				else if resultado.Tipo != Ast.EJECUTADO && (i.Tipo == Ast.IF_EXPRESION ||
					i.Tipo == Ast.ELSE_EXPRESION || i.Tipo == Ast.ELSEIF_EXPRESION) {
					scope.UpdateScopeGlobal()
					guardarScope = Ast.GetTemp()
					posicionGuardar = Ast.GetTemp()
					scopeOrigen = scope.GetEntornoPadreIF()
					resultadoTranferencia = resultado*/
			//codigo3d += "/***********************GUARDAR EL VALOR DEL IF*/\n"
			//codigo3d += guardarScope + " = P; //Guardar scope anterior \n"
			//codigo3d += "P = " + strconv.Itoa(scopeOrigen.Posicion) + "; //Cambio a entorno simulado \n"
			//codigo3d += posicionGuardar + " = P + 0; //Pos para el valor del if \n"
			//codigo3d += "stack[(int)" + posicionGuardar + "] = " + objResultadoIfs.Referencia + "; //Guardar valor \n"
			//codigo3d += "P = " + guardarScope + "; //Retornar al entorno anterior \n"
			//codigo3d += "/***********************************************/\n"
			//return resultado
			//hayRetornoIf = true
			//}

			faltaActual = objResultadoIfs.Lf
			ultimoFalso = faltaActual
			saltosFin += objResultadoIfs.Salto
			newScope.UpdateScopeGlobal()
		}
		///////////////////////////////FALSAS//////////////////////////////////////
		if i.Tipo == Ast.IF || i.Tipo == Ast.IF_EXPRESION {
			if ultimoFalso != "" {
				codigo3d += ultimoFalso + ":\n"
			}
			codigo3d += saltosFin
		}
		codigo3d += "/***********************************************/\n"
		obj3d.Salto = saltosFin
		obj3d.Codigo = codigo3d
		obj3d.Valor = Ast.TipoRetornado{
			Valor: true,
			Tipo:  Ast.EJECUTADO,
		}
		obj3d.Lt = obj3dCondicion.Lt
		obj3d.Lf = obj3dCondicion.Lf

		if hayTranferencia {
			obj3d.SaltoBreak = saltosBreak
			obj3d.SaltoBreakExp = saltosBreakExp
			obj3d.SaltoContinue = saltosContinue
			obj3d.SaltoReturn = saltosReturn
			obj3d.SaltoReturnExp = saltoReturnExp
			obj3d.Valor.Tipo = resultadoTranferencia.Tipo
			obj3d.Valor = resultadoTranferencia
			if posCorrectaNoElse {
				obj3d.IfCorrecto = true
			}
			return Ast.TipoRetornado{
				Valor: obj3d,
				Tipo:  resultadoTranferencia.Tipo,
			}
		} else if hayRetornoIf {
			obj3d.Valor = resultadoTranferencia
			obj3d.Valor.Tipo = resultadoTranferencia.Tipo
			obj3d.RetornoIf = "Si hay"
			if posCorrectaNoElse {
				obj3d.IfCorrecto = true
			}
			return Ast.TipoRetornado{
				Tipo:  resultadoTranferencia.Tipo,
				Valor: obj3d,
			}
		}

		return Ast.TipoRetornado{
			Valor: obj3d,
			Tipo:  Ast.EJECUTADO,
		}

	} else {
		//No es booleano, entonces generar un error semántico
		//fmt.Println("Error semántico, la expresión no es un booleano")
		msg := "Semantic error, the condition of the expression is not a boolean expression." +
			" -- Line:" + strconv.Itoa(i.Fila) + " Column: " + strconv.Itoa(i.Columna)
		nError := errores.NewError(i.Fila, i.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}
	//Se acabo todo y retornar un true de finalizado
	/*
		return Ast.TipoRetornado{
			Valor: true,
			Tipo:  Ast.EJECUTADO,
		}
	*/
}

func (op IF) GetFila() int {
	return op.Fila
}
func (op IF) GetColumna() int {
	return op.Columna
}

func GenerarCod3DCondicionEspecial(referencia string) Ast.O3D {
	var codigo3d string
	var lt, lf string
	var obj3d Ast.O3D
	lt = Ast.GetLabel()
	lf = Ast.GetLabel()
	codigo3d += "if (" + referencia + " == 1) goto " + lt + ";\n"
	codigo3d += "goto " + lf + ";\n"
	obj3d.Referencia = referencia
	obj3d.Codigo = codigo3d
	obj3d.Lt = lt
	obj3d.Lf = lf
	return obj3d
}
