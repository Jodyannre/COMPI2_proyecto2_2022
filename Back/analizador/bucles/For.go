package bucles

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"
	"strings"

	"github.com/colegno/arraylist"
)

type For struct {
	Tipo          Ast.TipoDato
	Variable      interface{}
	Range         interface{}
	Instrucciones *arraylist.List
	Fila          int
	Columna       int
}

func NewFor(variable interface{}, condicion interface{}, instrucciones *arraylist.List,
	fila, columna int) For {
	nF := For{
		Tipo:          Ast.FOR,
		Instrucciones: instrucciones,
		Range:         condicion,
		Fila:          fila,
		Columna:       columna,
		Variable:      variable,
	}
	return nF
}

func (f For) Run(scope *Ast.Scope) interface{} {
	/**************************************VARIABLES 3D*************************************/
	var obj3dRange, obj3dVariableFor, obj3dResultado, obj3d Ast.O3D
	var codigo3d string
	var codigo3dFor string
	var direccion int
	var lt, lf, salto string
	var limiteSuperior, limiteInferior string
	var iteracionPara3D bool = true
	var saltoBreak, saltoContinue, saltoReturn, saltoReturnExp string
	/**************************************************************************************/

	var variable expresiones.Identificador
	var nombreVariable string
	var tipoGeneral Ast.TipoDato
	var primerValor Ast.TipoRetornado
	var rango Ast.TipoRetornado
	var nSimbolo Ast.Simbolo
	var vector interface{}
	//var simboloTemp Ast.Simbolo
	//var variableTemp Ast.TipoRetornado
	var valorActual Ast.TipoRetornado
	var instruccion interface{}
	var tipoParticular Ast.TipoDato
	var resultadoInstruccion Ast.TipoRetornado
	//var primeraIteracion bool
	newScope := Ast.NewScope("For", scope)

	//Inicializar algunas variables de 3d
	lt = Ast.GetLabel()
	lf = Ast.GetLabel()
	salto = Ast.GetLabel()

	//Verificar que la expresión sea un identificador o error
	_, tipoParticular = f.Variable.(Ast.Abstracto).GetTipo()

	if tipoParticular != Ast.IDENTIFICADOR {
		//Error se espera un identificador
		fila := f.Variable.(Ast.Abstracto).GetFila()
		columna := f.Variable.(Ast.Abstracto).GetColumna()
		msg := "Semantic error, IDENTIFICADOR expected, found " + Ast.ValorTipoDato[tipoParticular] +
			". -- Line:" + strconv.Itoa(fila) + " Column: " +
			strconv.Itoa(columna)
		nError := errores.NewError(fila, columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		scope.UpdateScopeGlobal()
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}

	//Recuperar la variable y sus datos para crearla
	variable = f.Variable.(expresiones.Identificador)
	nombreVariable = variable.Valor
	//error, identifier expected

	//Ejecutar el range para obtener el vector que se va a iterar
	rango = f.Range.(Ast.Expresion).GetValue(scope)
	obj3dRange = rango.Valor.(Ast.O3D)
	codigo3d += obj3dRange.Codigo
	limiteInferior = obj3dRange.Lt
	limiteSuperior = obj3dRange.Lf
	rango = obj3dRange.Valor
	//Verificar error
	if rango.Tipo == Ast.ERROR {
		return rango
	}
	if rango.Tipo == Ast.VECTOR {
		vector = rango.Valor.(expresiones.Vector)
	} else {
		vector = rango.Valor.(expresiones.Array)
	}

	//Get el primer elemento
	if rango.Tipo == Ast.VECTOR {
		if rango.Valor.(expresiones.Vector).Valor.Len() > 0 {
			primerValor = rango.Valor.(expresiones.Vector).Valor.GetValue(0).(Ast.TipoRetornado)
		} else {
			primerValor = Ast.TipoRetornado{
				Tipo:  Ast.I64,
				Valor: 0,
			}
		}

	} else {
		primerValor = rango.Valor.(expresiones.Array).Elementos.GetValue(0).(Ast.TipoRetornado)

	}

	//Crear el símbolo de la variable que se va a utilizar en el for
	nSimbolo = Ast.Simbolo{
		Identificador: nombreVariable,
		Valor:         primerValor,
		Fila:          variable.Fila,
		Columna:       variable.Columna,
		Tipo:          primerValor.Tipo,
		Mutable:       true,
		Publico:       false,
		Entorno:       &newScope,
	}

	/*******************************CODIGO 3D PARA CREAR EL SIMBOLO**************************/
	codigo3d += "/*************CAMBIO A ENTORNO SIMULADO DEL FOR*/ \n"
	codigo3d += "P = P + " + strconv.Itoa(scope.Size) + "; //Cambio de entorno \n"
	newScope.Posicion = scope.Size + scope.Posicion
	codigo3d += "/***************DECLARACIÓN DE VARIABLE DEL FOR*/ \n"
	temp2 := Ast.GetTemp()
	direccion = newScope.Size
	codigo3d += temp2 + " = " + " P + " + strconv.Itoa(direccion) + ";\n"
	_, tipoParticular = f.Range.(Range).ValorInf.(Ast.Abstracto).GetTipo()

	if expresiones.EsArray(tipoParticular) == Ast.ARRAY {
		temp := Ast.GetTemp()
		codigo3d += temp + " = " + "heap[(int)" + limiteInferior + "]; //Get elemento desde el heap \n"
		codigo3d += "stack[(int)" + temp2 + "] = " + temp + ";\n"
	} else {
		codigo3d += "stack[(int)" + temp2 + "] = " + limiteInferior + ";\n"
	}

	newScope.Size++
	codigo3d += "/***********************************************/\n"

	nSimbolo.Direccion = direccion
	nSimbolo.TipoDireccion = Ast.STACK
	//Agregar el nuevo simbolo al scope del for
	newScope.Add(nSimbolo)
	/****************************************************************************************/

	//primeraIteracion = true
	if rango.Tipo == Ast.VECTOR {

		//for i := 0; i < vector.(expresiones.Vector).Valor.Len(); i++ {
		//Hasta que sea verdadero y termine de iterar toda la lista
		//Actualizar el valor de la variable al siguiente elemento
		//Verificar la variable por si fue modifica en la iteración anterior

		//Recupero la variable tal como esta luego de la iteracion
		//valorActual = vector.(expresiones.Vector).Valor.GetValue(0).(Ast.TipoRetornado)

		if rango.Valor.(expresiones.Vector).Valor.Len() > 0 {
			valorActual = rango.Valor.(expresiones.Vector).Valor.GetValue(0).(Ast.TipoRetornado)
		} else {
			valorActual = Ast.TipoRetornado{
				Tipo:  Ast.I64,
				Valor: 0,
			}
		}
		//simboloTemp = newScope.GetSimbolo(nombreVariable)
		//variableTemp = simboloTemp.Valor.(Ast.TipoRetornado)
		nSimbolo.Valor = vector.(expresiones.Vector).Valor.GetValue(0)
		//Verifico el valor antes de actualizar
		nSimbolo.Valor = valorActual
		newScope.UpdateSimbolo(nombreVariable, nSimbolo)

		/********************************CODIGO 3D PARA ACTUALIZAR LA VARIABLE***********************/

		/********************************************************************************************/

		//Ejectuar todas las instrucciones dentro del for en la n iteración
		for j := 0; j < f.Instrucciones.Len(); j++ {
			instruccion = f.Instrucciones.GetValue(j)

			//Verificar los tipos para saber que comportamiento tiene que tener
			tipoGeneral, _ = instruccion.(Ast.Abstracto).GetTipo()

			if tipoGeneral == Ast.INSTRUCCION {
				//Ejecutar run
				resultadoInstruccion = instruccion.(Ast.Instruccion).Run(&newScope).(Ast.TipoRetornado)
				obj3dResultado = resultadoInstruccion.Valor.(Ast.O3D)
				resultadoInstruccion = obj3dResultado.Valor
				if iteracionPara3D {
					codigo3dFor += obj3dResultado.Codigo
				}

			} else if tipoGeneral == Ast.EXPRESION {
				//Ejecutar getvalue
				resultadoInstruccion = instruccion.(Ast.Expresion).GetValue(&newScope)
				obj3dResultado = resultadoInstruccion.Valor.(Ast.O3D)
				resultadoInstruccion = obj3dResultado.Valor
				for iteracionPara3D {
					codigo3dFor += obj3dResultado.Codigo
				}
			}

			//Verificar las instrucciones de transferencia
			if resultadoInstruccion.Tipo == Ast.CONTINUE ||
				resultadoInstruccion.Tipo == Ast.BREAK ||
				resultadoInstruccion.Tipo == Ast.RETURN {
				//Siguiente iteración
				newScope.UpdateScopeGlobal()
				newScope.Errores.Clear()
				newScope.Consola = ""

				if !obj3dResultado.TranferenciaAgregada {
					codigo3d += "goto " + obj3dResultado.Salto + ";\n"
				}
				saltoContinue += obj3dResultado.SaltoContinue
				saltoContinue = strings.Replace(saltoContinue, ",", ":\n", -1)

				saltoBreak += obj3dResultado.SaltoBreak
				//println(saltoBreak)
				saltoBreak = strings.Replace(saltoBreak, ",", ":\n", -1)

				saltoReturn += obj3dResultado.SaltoReturn
				saltoReturnExp += obj3dResultado.SaltoReturnExp
				//saltoReturn = strings.Replace(saltoReturn, ",", ":\n", -1)
			}

		}
		iteracionPara3D = false
		//primeraIteracion = false
		//}

	} else {
		//for i := 0; i < vector.(expresiones.Array).Elementos.Len(); i++ {
		//Hasta que sea verdadero y termine de iterar toda la lista
		//Actualizar el valor de la variable al siguiente elemento
		//Verificar la variable por si fue modifica en la iteración anterior

		//Recupero la variable tal como esta luego de la iteracion
		valorActual = vector.(expresiones.Array).Elementos.GetValue(0).(Ast.TipoRetornado)
		//simboloTemp = newScope.GetSimbolo(nombreVariable)
		//variableTemp = simboloTemp.Valor.(Ast.TipoRetornado)
		nSimbolo.Valor = vector.(expresiones.Array).Elementos.GetValue(0)
		//Verifico el valor antes de actualizar
		nSimbolo.Valor = valorActual

		newScope.UpdateSimbolo(nombreVariable, nSimbolo)

		//Ejectuar todas las instrucciones dentro del for en la n iteración
		for j := 0; j < f.Instrucciones.Len(); j++ {
			instruccion = f.Instrucciones.GetValue(j)

			//Verificar los tipos para saber que comportamiento tiene que tener
			tipoGeneral, _ = instruccion.(Ast.Abstracto).GetTipo()

			if tipoGeneral == Ast.INSTRUCCION {
				//Ejecutar run
				resultadoInstruccion = instruccion.(Ast.Instruccion).Run(&newScope).(Ast.TipoRetornado)
				obj3dResultado = resultadoInstruccion.Valor.(Ast.O3D)
				resultadoInstruccion = obj3dResultado.Valor
				codigo3dFor += obj3dResultado.Codigo

			} else if tipoGeneral == Ast.EXPRESION {
				//Ejecutar getvalue
				resultadoInstruccion = instruccion.(Ast.Expresion).GetValue(&newScope)
				obj3dResultado = resultadoInstruccion.Valor.(Ast.O3D)
				resultadoInstruccion = obj3dResultado.Valor
				codigo3dFor += obj3dResultado.Codigo
			}
			//Verificar las instrucciones de transferencia
			if Ast.EsTransferencia(resultadoInstruccion.Tipo) {

				//Primero verificar que no sea un return normal, el cual si es permitido
				if resultadoInstruccion.Tipo == Ast.CONTINUE {
					//Rompemos y vamos a la siguiente iteración del for
					break
				}
				switch resultadoInstruccion.Tipo {
				case Ast.BREAK_EXPRESION, Ast.RETURN_EXPRESION, Ast.RETURN:
					return resultadoInstruccion
				case Ast.BREAK:
					return Ast.TipoRetornado{
						Tipo:  Ast.EJECUTADO,
						Valor: true,
					}
				}
			}

		}
		//primeraIteracion = false
		//}
	}

	posActualizacion := Ast.GetTemp()
	nuevoValor := Ast.GetTemp()
	codigo3d += "/*********************************EJECUCION FOR*/ \n"
	codigo3d += "//#aquiVaElSaltoContinue\n"
	codigo3d += salto + ":\n"
	codigo3d += "/***********CONSEGUIR VALOR DE VARIABLE DEL FOR*/\n"
	idExp := expresiones.NewIdentificador(nombreVariable, Ast.IDENTIFICADOR, 0, 0)
	obj3dVariableFor = idExp.GetValue(&newScope).Valor.(Ast.O3D)
	codigo3d += obj3dVariableFor.Codigo
	codigo3d += "if (" + obj3dVariableFor.Referencia + " < " + limiteSuperior + ") goto " + lt + ";\n"
	codigo3d += "goto " + lf + ";\n"
	codigo3d += lt + ":\n"
	codigo3d += "/***********************************************/\n"
	codigo3d += codigo3dFor
	codigo3d += "/**********ACTUALIZAR VALOR DE VARIABLE DEL FOR*/\n"
	codigo3d += posActualizacion + " = P + 0;\n"
	codigo3d += nuevoValor + " = " + obj3dVariableFor.Referencia + " + 1;//Nuevo valor \n"
	codigo3d += "stack[(int)" + posActualizacion + "] = " + nuevoValor + "; //Actualizar valor \n"
	codigo3d += "goto " + salto + "; \n"
	codigo3d += lf + ":\n"

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

	newScope.UpdateScopeGlobal()

	obj3d.Valor = Ast.TipoRetornado{
		Tipo:  Ast.EJECUTADO,
		Valor: true,
	}
	obj3d.Codigo = codigo3d

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

func (op For) GetFila() int {
	return op.Fila
}
func (op For) GetColumna() int {
	return op.Columna
}
func (f For) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, f.Tipo
}
