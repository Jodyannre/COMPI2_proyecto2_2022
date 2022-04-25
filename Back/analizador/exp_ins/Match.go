package exp_ins

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"strconv"
	"strings"

	"github.com/colegno/arraylist"
)

type Match struct {
	Expresion Ast.Expresion
	Cases     *arraylist.List
	Tipo      Ast.TipoDato
	Fila      int
	Columna   int
}

type Case struct {
	Expresion     *arraylist.List
	Instrucciones *arraylist.List
	Tipo          Ast.TipoDato
	Default       bool
	Fila          int
	Columna       int
}

func NewMatch(expresion Ast.Expresion, cases *arraylist.List, tipo Ast.TipoDato, fila, columna int) Match {
	nMatch := Match{
		Expresion: expresion,
		Cases:     cases,
		Tipo:      tipo,
		Fila:      fila,
		Columna:   columna,
	}
	return nMatch
}

func NewCase(expresion *arraylist.List, instrucciones *arraylist.List, tipo Ast.TipoDato,
	fila, columna int, Default bool) Case {
	nCase := Case{
		Expresion:     expresion,
		Instrucciones: instrucciones,
		Tipo:          tipo,
		Default:       Default,
		Fila:          fila,
		Columna:       columna,
	}
	return nCase
}

func (m Match) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, m.Tipo
}

func (c Case) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, c.Tipo
}

func (m Match) Run(scope *Ast.Scope) interface{} {
	/*******************************VARIABLES 3D*********************************/
	var codigo3d string
	var obj3d, obj3dvalor, obj3dExpresion, obj3dResultadoExpresion Ast.O3D
	var referenciaExpresion, referenciaExpCase string
	var labelsTrue, saltos, salto string
	var falsoAnterior string = ""
	var resultadoTranferencia Ast.TipoRetornado
	var hayTranferencia bool = false
	var saltosContinue string
	var saltosBreak string
	var saltosReturn string
	var saltoReturnExp string
	var resultado Ast.TipoRetornado
	pos_correcta := false
	pos_correta_no_default := false
	/****************************************************************************/

	//Primero conseguir el exp_match de la expresion
	buscar_default := false
	default_econtrado := false
	///////////////////////////VALOR EXPRESION////////////////////////////////
	exp_match := m.Expresion.GetValue(scope)
	obj3dExpresion = exp_match.Valor.(Ast.O3D)
	exp_match = obj3dExpresion.Valor
	codigo3d += obj3dExpresion.Codigo
	referenciaExpresion = obj3dExpresion.Referencia
	var hayRetornoCase bool = false
	//////////////////////////////////////////////////////////////////////////
	//pos_expresion_correcta := -1
	var i, j int

	//Validar si es un boolean y si tiene todos los casos
	if exp_match.Tipo == Ast.BOOLEAN && m.Cases.Len() < 2 {
		//Error, no estan todos los casos incluidos
		msg := "Semantic error, a default case was expected." +
			" -- Line:" + strconv.Itoa(m.Fila) + " Column: " + strconv.Itoa(m.Columna)
		nError := errores.NewError(m.Fila, m.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Valor: nError,
			Tipo:  Ast.ERROR,
		}
	}

	//Verificar si no es boolean y se espera un default
	if exp_match.Tipo != Ast.BOOLEAN {
		buscar_default = true
	}

	for i = 0; i < m.Cases.Len(); i++ {
		//Recorrer la lista de cases y verificar que los exp_matches de las expresiones concuerdan
		caso := m.Cases.GetValue(i).(Case)
		listaExpresiones := caso.Expresion
		var expresion Ast.TipoRetornado
		for j = 0; j < listaExpresiones.Len(); j++ {
			expresion = listaExpresiones.GetValue(j).(Ast.Expresion).GetValue(scope)
			obj3dExpresion = expresion.Valor.(Ast.O3D)
			expresion = obj3dExpresion.Valor
			referenciaExpCase = obj3dExpresion.Referencia
			if expresion.Tipo != exp_match.Tipo && !caso.Default {
				//Error de tipos
				msg := "Semantic error, the expression is not a " + Ast.ValorTipoDato[exp_match.Tipo] +
					" value. -- Line:" + strconv.Itoa(caso.Fila) + " Column: " + strconv.Itoa(caso.Columna)
				nError := errores.NewError(caso.Fila, caso.Columna, msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				nError.Ambito = scope.GetTipoScope()
				scope.Errores.Add(nError)
				scope.Consola += msg + "\n"
				return Ast.TipoRetornado{
					Valor: nError,
					Tipo:  Ast.ERROR,
				}
			}
			if !caso.Default {
				obj3dResultadoExpresion = CrearOpRelacional3D(referenciaExpresion, referenciaExpCase)
			} else {
				obj3dResultadoExpresion = CrearOpRelacional3D(referenciaExpresion, referenciaExpresion)
			}
			labelsTrue += obj3dResultadoExpresion.Lt + ":\n"
			codigo3d += falsoAnterior
			codigo3d += obj3dResultadoExpresion.Codigo
			falsoAnterior = obj3dResultadoExpresion.Lf + ": \n"
			if exp_match.Valor == expresion.Valor {
				if !pos_correta_no_default {
					pos_correcta = true
					pos_correta_no_default = true
				}
			}
			if j == listaExpresiones.Len()-1 && !pos_correta_no_default {
				pos_correcta = true
			}
		}
		//Verificar si viene default
		if buscar_default {
			if caso.Default {
				default_econtrado = true
			}
		}

		if default_econtrado && i != m.Cases.Len()-1 {
			//Error, el default tiene que venir de último
			msg := "Semantic error, the default case was expected in the last position." +
				" -- Line:" + strconv.Itoa(caso.Fila) + " Column: " + strconv.Itoa(caso.Columna)
			nError := errores.NewError(caso.Fila, caso.Columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			nError.Ambito = scope.GetTipoScope()
			scope.Errores.Add(nError)
			scope.Consola += msg + "\n"
			return Ast.TipoRetornado{
				Valor: nError,
				Tipo:  Ast.ERROR,
			}
		}

		////////////////////EJECUTAR CASE PARA CONSEGUIR EL C3D/////////////////////////////////
		resultado = caso.Run(scope).(Ast.TipoRetornado)
		obj3dvalor = resultado.Valor.(Ast.O3D)
		resultado = obj3dvalor.Valor

		codigo3d += labelsTrue
		labelsTrue = ""
		codigo3d += obj3dvalor.Codigo
		salto = Ast.GetLabel()
		codigo3d += "goto " + salto + ";\n"
		saltos += salto + ","
		codigo3d += falsoAnterior
		falsoAnterior = ""

		switch resultado.Tipo {
		case Ast.BREAK:
			saltosBreak += obj3dvalor.SaltoBreak + ","
			hayTranferencia = true
			//resultadoTranferencia = resultado
		case Ast.CONTINUE:
			saltosContinue += obj3dvalor.SaltoContinue + ","
			hayTranferencia = true
			//resultadoTranferencia = resultado
		case Ast.RETURN:
			saltosReturn += obj3dvalor.SaltoReturn + ","
			saltoReturnExp += obj3dvalor.SaltoReturnExp
			hayTranferencia = true
			//resultadoTranferencia = resultado
		}

		if obj3dvalor.RetornoMatch != "" {
			hayRetornoCase = true
		}
		if pos_correcta {
			resultadoTranferencia = resultado
			pos_correcta = false
		}

		///////////////////////////////////////////////////////////////////////////////

	}

	/************************AGREGAR TODOS LOS SALTOS VERDADERSO*****************************/
	saltos = strings.Replace(saltos, ",", ":\n", -1)
	codigo3d += saltos

	/***************************************************************************************/

	//Verificar que se necesite default y que lo traiga
	if buscar_default {
		if !default_econtrado {
			msg := "Semantic error, default case not found." +
				" -- Line:" + strconv.Itoa(m.Fila) + " Column: " + strconv.Itoa(m.Columna)
			nError := errores.NewError(m.Fila, m.Columna, msg)
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
	}

	//Todos los casos son correctos
	//Recorrer los cases
	//var resultado_retornado Ast.TipoRetornado
	//Recuperar el caso en donde se encontró el valor igual

	//Aqui ya no va a ejecutar los cases
	/*
		if pos_expresion_correcta != -1 {
			caso := m.Cases.GetValue(pos_expresion_correcta).(Case)
			resultado_retornado = caso.Run(scope).(Ast.TipoRetornado)
		} else {
			//Ejecutar default
			caso := m.Cases.GetValue(m.Cases.Len() - 1).(Case)
			resultado_retornado = caso.Run(scope).(Ast.TipoRetornado)
		}

		if Ast.EsTransferencia(resultado_retornado.Tipo) &&
			m.Tipo != Ast.MATCH_EXPRESION {
			return resultado_retornado
		}

		if resultado_retornado.Tipo == Ast.ERROR {
			return resultado_retornado
		}

		if resultado_retornado.Tipo != Ast.EJECUTADO && m.Tipo == Ast.MATCH_EXPRESION {
			return resultado_retornado
		}

		if m.Tipo == Ast.MATCH_EXPRESION && resultado_retornado.Tipo == Ast.EJECUTADO {
			//Error, el match no esta retornando nada
			msg := "Semantic error, Match statement is not returning any value." +
				" -- Line:" + strconv.Itoa(m.Fila) + " Column: " + strconv.Itoa(m.Columna)
			nError := errores.NewError(m.Fila, m.Columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			nError.Ambito = scope.GetTipoScope()
			scope.Errores.Add(nError)
			scope.Consola += msg + "\n"
			return Ast.TipoRetornado{
				Valor: nError,
				Tipo:  Ast.ERROR,
			}
		}
	*/

	obj3d.Codigo = codigo3d

	if hayTranferencia {
		obj3d.SaltoBreak = saltosBreak
		obj3d.SaltoContinue = saltosContinue
		obj3d.SaltoReturn = saltosReturn
		obj3d.Valor.Tipo = resultadoTranferencia.Tipo
		obj3d.SaltoReturnExp = saltoReturnExp
		obj3d.Valor = resultadoTranferencia
		return Ast.TipoRetornado{
			Valor: obj3d,
			Tipo:  resultadoTranferencia.Tipo,
		}
	} else if hayRetornoCase {
		entornoAnterior := Ast.GetTemp()
		posicionBreak := Ast.GetTemp()
		valorCase := Ast.GetTemp()
		codigo3d := ""
		//Recuperar el valor de la variable que se va a retornar
		codigo3d += "/******RECUPERAR EL VALOR DEL RETORNO DEL MATCH*/\n"
		codigo3d += entornoAnterior + " = P; //guardar entorno anterior \n"
		codigo3d += "P = P + " + strconv.Itoa(scope.Size) + "; //Cambio de entorno \n"
		codigo3d += posicionBreak + " = P + 0; //Posicion valor \n"
		codigo3d += valorCase + " = stack[(int)" + posicionBreak + "]; //Valor del retorno \n"
		codigo3d += "P = P - " + strconv.Itoa(scope.Size) + "; //Regresar al entorno anterior \n"
		codigo3d += "/***********************************************/\n"
		obj3d.Codigo += codigo3d
		obj3d.Referencia = valorCase
		obj3d.Valor = resultadoTranferencia
		return Ast.TipoRetornado{
			Valor: obj3d,
			Tipo:  obj3dvalor.Valor.Tipo,
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

func (c Case) Run(scope *Ast.Scope) interface{} {

	/********************************CODIGO 3D***********************************/
	var codigo3d string
	var obj3d, obj3dValor Ast.O3D
	var resultadoTranferencia Ast.TipoRetornado
	var saltosBreak, saltosContinue, saltosReturn, saltoReturnExp string
	var hayTranferencia bool
	var hayRetornoCase bool
	/****************************************************************************/

	//Crear un nuevo scope y otras variables
	newScope := Ast.NewScope("Case", scope)

	//Direccion del nuevo entorno
	newScope.Posicion = scope.Size + scope.Posicion
	//Aumentar +1 por el posible retorno

	newScope.Size++
	codigo3d += "/******************************************CASE*/\n"
	codigo3d += "P = P + " + strconv.Itoa(scope.Size) + "; //Cambio de ambito \n"

	var expresion = false
	//ultimaExpresion := Ast.TipoRetornado{Valor: nil, Tipo: Ast.NULL}
	var ultimoTipo Ast.TipoDato
	var instruccion interface{}
	var resultado Ast.TipoRetornado
	i := 0

	//Verificar el tipo del caso, si es exp o ins
	if c.Tipo == Ast.CASE_EXPRESION {
		expresion = true
	}

	//Ejecutar las instrucciones

	for i = 0; i < c.Instrucciones.Len(); i++ {
		//Verificar el tipo de entrada
		instruccion = c.Instrucciones.GetValue(i).(Ast.Abstracto)
		tipo_abstracto, _ := instruccion.(Ast.Abstracto).GetTipo()
		ultimoTipo = tipo_abstracto
		if tipo_abstracto == Ast.EXPRESION {
			//Si es un if expresión, tiene que retornar algo
			expresion := c.Instrucciones.GetValue(i).(Ast.Expresion)
			resultado = expresion.GetValue(&newScope)
			obj3dValor = resultado.Valor.(Ast.O3D)
			resultado = obj3dValor.Valor
			codigo3d += obj3dValor.Codigo
		} else if tipo_abstracto == Ast.INSTRUCCION {
			instruccion := c.Instrucciones.GetValue(i).(Ast.Instruccion)
			resultado = instruccion.Run(&newScope).(Ast.TipoRetornado)
			obj3dValor = resultado.Valor.(Ast.O3D)
			resultado = obj3dValor.Valor
			codigo3d += obj3dValor.Codigo
		}

		//Error si viene un break o un return dentro de un case expresion
		if Ast.EsTransferencia(resultado.Tipo) &&
			c.Tipo == Ast.CASE_EXPRESION {
			temp := instruccion.(Ast.Abstracto)
			msg := "Semantic error, transfer statements are not allowed within a case expression statement." +
				" -- Line:" + strconv.Itoa(temp.GetFila()) + " Column: " + strconv.Itoa(temp.GetColumna())
			nError := errores.NewError(temp.GetFila(), temp.GetColumna(), msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			nError.Ambito = scope.GetTipoScope()
			newScope.Errores.Add(nError)
			newScope.Consola += msg + "\n"
			newScope.UpdateScopeGlobal()
			return Ast.TipoRetornado{
				Valor: nError,
				Tipo:  Ast.ERROR,
			}
		}

		if Ast.EsTransferencia(resultado.Tipo) {
			//Retornar la transferencia
			newScope.UpdateScopeGlobal()
			resultadoTranferencia = resultado
			switch resultado.Tipo {
			case Ast.BREAK:
				saltosBreak += obj3dValor.SaltoBreak
			case Ast.CONTINUE:
				saltosContinue += obj3dValor.SaltoContinue
			case Ast.RETURN:
				saltosReturn += obj3dValor.SaltoReturn
				saltoReturnExp += obj3dValor.SaltoReturnExp
			}
			//saltosTransferencia += objResultadoIfs.SaltoBreak
			hayTranferencia = true
			//return resultado
		} else if resultado.Tipo != Ast.EJECUTADO && c.Tipo == Ast.CASE_EXPRESION {
			scope.UpdateScopeGlobal()
			guardarScope := Ast.GetTemp()
			posicionGuardar := Ast.GetTemp()
			//scopeOrigen := scope.GetEntornoPadreBreak()
			resultadoTranferencia = resultado
			codigo3d += "/*********************GUARDAR EL VALOR DEL CASE*/\n"
			codigo3d += guardarScope + " = P; //Guardar scope anterior \n"
			codigo3d += "P = " + strconv.Itoa(scope.Posicion+scope.Size) + "; //Cambio a entorno simulado \n"
			codigo3d += posicionGuardar + " = P + 0; //Pos para el valor del case \n"
			codigo3d += "stack[(int)" + posicionGuardar + "] = " + obj3dValor.Referencia + "; //Guardar valor \n"
			codigo3d += "P = " + guardarScope + "; //Retornar al entorno anterior \n"
			codigo3d += "/***********************************************/\n"
			//return resultado
			hayRetornoCase = true
		}
	}

	//Termino el for, retornar la ultima expresion
	//Verificar si hay algun retorno o retornar un error
	if ultimoTipo != Ast.EXPRESION && expresion {
		msg := "Semantic error, the match clause is not returning any value." +
			" -- Line:" + strconv.Itoa(c.Fila) + " Column: " + strconv.Itoa(c.Columna)
		nError := errores.NewError(c.Fila, c.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		newScope.UpdateScopeGlobal()
		return Ast.TipoRetornado{
			Valor: nError,
			Tipo:  Ast.ERROR,
		}
	}
	/*
		else if expresion && ultimoTipo == Ast.EXPRESION {
			//Si esta retornado algún valor
			newScope.UpdateScopeGlobal()
			return resultado
		}
	*/

	codigo3d += "P = P - " + strconv.Itoa(scope.Size) + "; //Retornar al ambito anterior \n"
	codigo3d += "/***********************************************/\n"

	obj3d.Codigo = codigo3d
	obj3d.Valor = Ast.TipoRetornado{
		Tipo:  Ast.EJECUTADO,
		Valor: true,
	}

	if hayTranferencia {
		obj3d.SaltoBreak = saltosBreak
		obj3d.SaltoContinue = saltosContinue
		obj3d.SaltoReturn = saltosReturn
		obj3d.Valor.Tipo = resultadoTranferencia.Tipo
		obj3d.SaltoReturnExp = saltoReturnExp
		obj3d.Valor = resultadoTranferencia
		newScope.UpdateScopeGlobal()
		return Ast.TipoRetornado{
			Valor: obj3d,
			Tipo:  resultadoTranferencia.Tipo,
		}
	} else if hayRetornoCase {
		obj3d.Valor = resultadoTranferencia
		obj3d.Valor.Tipo = resultadoTranferencia.Tipo
		obj3d.RetornoMatch = "Si hay"
		return Ast.TipoRetornado{
			Tipo:  resultadoTranferencia.Tipo,
			Valor: obj3d,
		}
	}

	newScope.UpdateScopeGlobal()
	return Ast.TipoRetornado{
		Tipo:  Ast.EJECUTADO,
		Valor: obj3d,
	}
}

func (op Match) GetFila() int {
	return op.Fila
}
func (op Match) GetColumna() int {
	return op.Columna
}

func CrearOpRelacional3D(refExp, refExpCase string) Ast.O3D {
	var lt, lf string
	var codigo3d string
	var obj3d Ast.O3D
	lt = Ast.GetLabel()
	lf = Ast.GetLabel()

	codigo3d += "/********************************CONDICION CASE*/\n"
	codigo3d += "if (" + refExp + " == " + refExpCase + ") goto " + lt + "; \n"
	codigo3d += "goto " + lf + ";\n"
	codigo3d += "/***********************************************/\n"

	obj3d.Valor = Ast.TipoRetornado{
		Tipo:  Ast.EJECUTADO,
		Valor: true,
	}

	obj3d.Codigo = codigo3d
	obj3d.Lt = lt
	obj3d.Lf = lf

	return obj3d
}
