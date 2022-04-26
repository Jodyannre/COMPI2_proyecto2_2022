package expresiones

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"fmt"
	"math"
)

var suma_dominante = [7][7]Ast.TipoDato{
	{Ast.I64, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.USIZE},
	{Ast.NULL, Ast.F64, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL},
	{Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.STRING_OWNED, Ast.NULL, Ast.NULL},
	{Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.STRING, Ast.NULL, Ast.NULL},
	{Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL},
	{Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL},
	{Ast.USIZE, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.USIZE},
}

var suma_dominante_comparacion = [8][8]Ast.TipoDato{
	{Ast.I64, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.I64},
	{Ast.NULL, Ast.F64, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL},
	{Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.STRING_OWNED, Ast.NULL, Ast.NULL},
	{Ast.NULL, Ast.NULL, Ast.NULL, Ast.STRING, Ast.NULL, Ast.NULL, Ast.NULL},
	{Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.STR, Ast.NULL, Ast.NULL},
	{Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.BOOLEAN, Ast.NULL},
	{Ast.I64, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.USIZE},
	{Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.CHAR},
}

var resta_dominante = [7][7]Ast.TipoDato{
	{Ast.I64, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.USIZE},
	{Ast.NULL, Ast.F64, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL},
	{Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL},
	{Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL},
	{Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL},
	{Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL},
	{Ast.USIZE, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.USIZE},
}

var mul_div_dominante = [7][7]Ast.TipoDato{
	{Ast.I64, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.USIZE},
	{Ast.NULL, Ast.F64, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL},
	{Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL},
	{Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL},
	{Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL},
	{Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL},
	{Ast.USIZE, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.USIZE},
}

type Operacion struct {
	operando_der Ast.Expresion
	operador     string
	operando_izq Ast.Expresion
	unario       bool
	Fila         int
	Columna      int
}

func NewOperation(op_izq Ast.Expresion, operador string,
	op_der Ast.Expresion, unario bool, fila, columna int) Operacion {
	nuevo := Operacion{
		operando_der: op_der,
		operando_izq: op_izq,
		operador:     operador,
		unario:       unario,
		Fila:         fila,
		Columna:      columna,
	}
	return nuevo
}

func (op Operacion) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, Ast.PRIMITIVO
}

func (op Operacion) GetFila() int {
	return op.Fila
}
func (op Operacion) GetColumna() int {
	return op.Columna
}

func (op Operacion) GetValue(entorno *Ast.Scope) Ast.TipoRetornado {
	var tipo_izq Ast.TipoRetornado
	var tipo_der Ast.TipoRetornado
	var result_dominante Ast.TipoDato

	if op.unario {
		tipo_izq = op.operando_izq.GetValue(entorno)
		tipo_der = tipo_izq
	} else {
		tipo_izq = op.operando_izq.GetValue(entorno)
		tipo_der = op.operando_der.GetValue(entorno)
	}
	//Verificar primero si no hay error
	if op.unario {
		if tipo_izq.Valor.(Ast.O3D).Valor.Tipo == Ast.ERROR {
			return tipo_izq
		}
	} else {
		if tipo_izq.Valor.(Ast.O3D).Valor.Tipo == Ast.ERROR {
			return tipo_izq
		}
		if tipo_der.Valor.(Ast.O3D).Valor.Tipo == Ast.ERROR {
			return tipo_der
		}
	}

	if tipo_izq.Valor.(Ast.O3D).Valor.Tipo > 7 || tipo_der.Valor.(Ast.O3D).Valor.Tipo > 7 {
		////////////////////////////ERROR////////////////////////////////
		//Error, no se pueden operar porque no es ningún valor operable
		vI := tipo_izq.Valor.(Ast.O3D).Valor
		vD := tipo_der.Valor.(Ast.O3D).Valor
		vI.Fila = op.Fila
		vI.Columna = op.Columna
		vD.Fila = op.Fila
		vD.Columna = op.Columna
		return errores.GenerarError(1, vI, vD, "", "", "", entorno)
	}

	switch op.operador {
	case "+":
		result_dominante = suma_dominante[tipo_izq.Valor.(Ast.O3D).Valor.Tipo][tipo_der.Valor.(Ast.O3D).Valor.Tipo]

		if result_dominante == Ast.I64 || result_dominante == Ast.USIZE {
			//Get los valores
			valorIzq := tipo_izq.Valor.(Ast.O3D).Valor.Valor
			valorDer := tipo_der.Valor.(Ast.O3D).Valor.Valor

			//Get el resultado de la operación
			valor := Ast.TipoRetornado{
				Tipo:  result_dominante,
				Valor: valorIzq.(int) + valorDer.(int),
			}

			//Actualizar el código y conseguir el obj O3D
			obj := Ast.ActualizarCodigoAritmetica(tipo_izq, tipo_der, op.operador, false)
			obj.Valor = valor

			return Ast.TipoRetornado{
				Tipo:  Ast.ARITMETICA,
				Valor: obj,
			}
		} else if result_dominante == Ast.F64 {
			//Get los valores
			valorIzq := tipo_izq.Valor.(Ast.O3D).Valor.Valor
			valorDer := tipo_der.Valor.(Ast.O3D).Valor.Valor
			var valIzq, valDer float64
			var valor Ast.TipoRetornado

			//Parseo por si viene algún int
			if tipo_izq.Valor.(Ast.O3D).Valor.Tipo == Ast.I64 {
				valIzq = float64(valorIzq.(int))
			} else {
				valIzq = valorIzq.(float64)
			}
			if tipo_der.Valor.(Ast.O3D).Valor.Tipo == Ast.I64 {
				valDer = float64(valorDer.(int))
			} else {
				valDer = valorDer.(float64)
			}

			//Get el valor
			valor = Ast.TipoRetornado{
				Tipo:  result_dominante,
				Valor: valIzq + valDer,
			}

			//Actualizar el código y conseguir el obj O3D
			obj := Ast.ActualizarCodigoAritmetica(tipo_izq, tipo_der, op.operador, false)
			obj.Valor = valor

			return Ast.TipoRetornado{
				Tipo:  Ast.ARITMETICA,
				Valor: obj,
			}
			/////////////////////////////////////////// SUMA DE STRINGS PENDIENTE ///////////////////////
		} else if result_dominante == Ast.STRING || result_dominante == Ast.STRING_OWNED {
			//Get los valores
			var obj3dValorIzq, obj3dValorDer, obj3d Ast.O3D
			var codigo3d string
			valorIzq := tipo_izq.Valor.(Ast.O3D).Valor.Valor
			valorDer := tipo_der.Valor.(Ast.O3D).Valor.Valor
			obj3dValorIzq = tipo_izq.Valor.(Ast.O3D)
			obj3dValorDer = tipo_der.Valor.(Ast.O3D)

			codigo3d += obj3dValorIzq.Codigo
			codigo3d += obj3dValorDer.Codigo

			obj3d = Ast.SumaStrings(obj3dValorIzq, obj3dValorDer)
			codigo3d += obj3d.Codigo
			obj3d.Codigo = codigo3d

			cadena_izq := fmt.Sprintf("%v", valorIzq)
			cadena_der := fmt.Sprintf("%v", valorDer)

			obj3d.Valor = Ast.TipoRetornado{
				Tipo:  result_dominante,
				Valor: cadena_izq + cadena_der,
			}

			return Ast.TipoRetornado{
				Tipo:  result_dominante,
				Valor: obj3d,
			}
		} else if result_dominante == Ast.NULL {
			//////////////////////////ERROR////////////////////////////////
			//Error los tipo no se pueden operar
			vI := tipo_izq.Valor.(Ast.O3D).Valor
			vD := tipo_der.Valor.(Ast.O3D).Valor
			vI.Fila = op.Fila
			vI.Columna = op.Columna
			vD.Fila = op.Fila
			vD.Columna = op.Columna
			return errores.GenerarError(2, vI, vD, "", "", "", entorno)

		}

	case "-", "!":
		valorIzq := tipo_izq.Valor.(Ast.O3D).Valor.Valor
		var valor Ast.TipoRetornado
		var valIzq, valDer float64

		if op.unario {
			//Es una operación unaria
			if op.operador == "-" {

				if tipo_izq.Valor.(Ast.O3D).Valor.Tipo == Ast.F64 {
					//Es un F64
					valor.Tipo = Ast.F64
					valor.Valor = valorIzq.(float64) * -1
				}
				if tipo_izq.Valor.(Ast.O3D).Valor.Tipo == Ast.I64 {
					//Es un I64
					valor.Tipo = Ast.I64
					valor.Valor = valorIzq.(int) * -1
				}

				//Actualizar el código y conseguir el obj O3D
				obj := Ast.ActualizarCodigoAritmetica(tipo_izq, tipo_izq, op.operador, true)
				obj.Valor = valor

				return Ast.TipoRetornado{
					Tipo:  Ast.ARITMETICA,
					Valor: obj,
				}

				////////////////////////PENDIENTE LOGICA///////////////////////////////
			} else if tipo_izq.Valor.(Ast.O3D).Valor.Tipo == Ast.BOOLEAN && op.operador == "!" {
				valorIzq = tipo_izq.Valor.(Ast.O3D).Valor.Valor
				valor.Valor = !valorIzq.(bool)
				valor.Tipo = Ast.BOOLEAN
				/************VERIFICAR QUE SEA UN IDENTIFICADOR************************/
				var opAbstracto interface{} = op.operando_izq
				_, tipoParticular := opAbstracto.(Ast.Abstracto).GetTipo()
				if tipoParticular == Ast.IDENTIFICADOR || tipoParticular == Ast.PRIMITIVO &&
					tipo_izq.Tipo != Ast.LOGICA {
					var objTemp Ast.O3D
					var codTemp string = tipo_izq.Valor.(Ast.O3D).Codigo
					var referencia string = tipo_izq.Valor.(Ast.O3D).Referencia
					var objEnviar Ast.TipoRetornado
					objTemp = GenerarCod3DLogicaEspecial(referencia)
					codTemp += objTemp.Codigo
					objTemp.Codigo = codTemp
					objEnviar.Valor = objTemp
					obj := Ast.ActualizarCodigoLogica(objEnviar, objEnviar, op.operador, true)
					obj.Valor = valor
					obj.Referencia = referencia
					return Ast.TipoRetornado{
						Tipo:  Ast.LOGICA,
						Valor: obj,
					}
				}

				/**********************************************************************/

				//Actualizar el código y conseguir el obj O3D
				obj := Ast.ActualizarCodigoLogica(tipo_izq, tipo_izq, op.operador, true)
				obj.Valor = valor
				obj.Referencia = tipo_izq.Valor.(Ast.O3D).Referencia
				return Ast.TipoRetornado{
					Tipo:  Ast.LOGICA,
					Valor: obj,
				}
				////////////////////////PENDIENTE LOGICA///////////////////////////////
			} else if tipo_izq.Valor.(Ast.O3D).Valor.Tipo == Ast.USIZE && op.operador == "-" {
				//Error, no se puede aplicar el menos a un usize
				vI := tipo_izq.Valor.(Ast.O3D).Valor
				vI.Fila = op.Fila
				vI.Columna = op.Columna
				return errores.GenerarError(3, vI, vI, "", "", "", entorno)
			} else {
				//Error, tipo no operable
				vI := tipo_izq.Valor.(Ast.O3D).Valor
				vI.Fila = op.Fila
				vI.Columna = op.Columna
				return errores.GenerarError(4, vI, vI, "", "", "", entorno)
			}

		}

		//Calcular resultado de la resta dominante
		valorDer := tipo_der.Valor.(Ast.O3D).Valor.Valor
		result_dominante = resta_dominante[tipo_izq.Valor.(Ast.O3D).Valor.Tipo][tipo_der.Valor.(Ast.O3D).Valor.Tipo]

		if result_dominante != Ast.NULL {

			if result_dominante == Ast.I64 {

				valor.Tipo = result_dominante
				valor.Valor = valorIzq.(int) - valorDer.(int)
			} else if result_dominante == Ast.USIZE {

				preValor := valorIzq.(int) - valorDer.(int)
				if preValor < 0 {
					///////////////////////////////ERROR/////////////////////////////////////
					//Error, el usize no puede ser negativo
					return errores.GenerarError(5, op, op, "", "", "", entorno)
				}
				valor.Tipo = result_dominante
				valor.Valor = valorIzq.(int) - valorDer.(int)
			} else if result_dominante == Ast.F64 {

				if tipo_izq.Valor.(Ast.O3D).Valor.Tipo == Ast.I64 {
					valIzq = float64(valorIzq.(int))
				} else {
					valIzq = valorIzq.(float64)
				}
				if tipo_der.Valor.(Ast.O3D).Valor.Tipo == Ast.I64 {
					valDer = float64(valorDer.(int))
				} else {
					valDer = valorDer.(float64)
				}
				valor.Tipo = result_dominante
				valor.Valor = valIzq - valDer
			}

			//Actualizar el código y conseguir el obj O3D
			obj := Ast.ActualizarCodigoAritmetica(tipo_izq, tipo_der, op.operador, false)
			obj.Valor = valor
			return Ast.TipoRetornado{
				Tipo:  Ast.ARITMETICA,
				Valor: obj,
			}

		} else if result_dominante == Ast.NULL {
			////////////////////////////////////ERROR/////////////////////////////////////
			vI := tipo_izq.Valor.(Ast.O3D).Valor
			vD := tipo_der.Valor.(Ast.O3D).Valor
			vI.Fila = op.Fila
			vI.Columna = op.Columna
			vD.Fila = op.Fila
			vD.Columna = op.Columna
			return errores.GenerarError(6, vI, vD, "", "", "", entorno)
		}

	case "*":
		result_dominante = mul_div_dominante[tipo_izq.Valor.(Ast.O3D).Valor.Tipo][tipo_der.Valor.(Ast.O3D).Valor.Tipo]

		if result_dominante != Ast.NULL {
			//Get los valores
			valorIzq := tipo_izq.Valor.(Ast.O3D).Valor.Valor
			valorDer := tipo_der.Valor.(Ast.O3D).Valor.Valor
			var valIzq, valDer float64
			var valor Ast.TipoRetornado

			if result_dominante == Ast.I64 || result_dominante == Ast.USIZE {
				valor.Tipo = result_dominante
				valor.Valor = valorIzq.(int) * valorDer.(int)
			} else if result_dominante == Ast.F64 {
				if tipo_izq.Valor.(Ast.O3D).Valor.Tipo == Ast.I64 {
					valIzq = float64(valorIzq.(int))
				} else {
					valIzq = valorIzq.(float64)
				}
				if tipo_der.Valor.(Ast.O3D).Valor.Tipo == Ast.I64 {
					valDer = float64(valorDer.(int))
				} else {
					valDer = valorDer.(float64)
				}
				valor.Tipo = result_dominante
				valor.Valor = valIzq * valDer
			}

			//Actualizar el código y conseguir el obj O3D
			obj := Ast.ActualizarCodigoAritmetica(tipo_izq, tipo_der, op.operador, false)
			obj.Valor = valor
			return Ast.TipoRetornado{
				Tipo:  Ast.ARITMETICA,
				Valor: obj,
			}

		} else if result_dominante == Ast.NULL {
			//////////////////////////////////ERROR///////////////////////////////////////
			vI := tipo_izq.Valor.(Ast.O3D).Valor
			vD := tipo_der.Valor.(Ast.O3D).Valor
			vI.Fila = op.Fila
			vI.Columna = op.Columna
			vD.Fila = op.Fila
			vD.Columna = op.Columna
			return errores.GenerarError(7, vI, vD, "", "", "", entorno)

		}
	case "/":
		result_dominante = mul_div_dominante[tipo_izq.Valor.(Ast.O3D).Valor.Tipo][tipo_der.Valor.(Ast.O3D).Valor.Tipo]

		if result_dominante != Ast.NULL {
			valorIzq := tipo_izq.Valor.(Ast.O3D).Valor.Valor
			valorDer := tipo_der.Valor.(Ast.O3D).Valor.Valor
			var valor Ast.TipoRetornado
			var valIzq, valDer float64
			var obj Ast.O3D

			if result_dominante == Ast.I64 {
				if tipo_der.Valor.(Ast.O3D).Valor.Valor.(int) == 0 {
					valor.Valor = 0
				} else {
					valor.Valor = valorIzq.(int) / valorDer.(int)
				}
				valor.Tipo = result_dominante

			} else if result_dominante == Ast.USIZE {
				if tipo_der.Valor.(Ast.O3D).Valor.Valor.(int) == 0 {
					valor.Valor = 0
				} else {
					valor.Valor = valorIzq.(int) / valorDer.(int)
				}
				valor.Tipo = result_dominante

			} else if result_dominante == Ast.F64 {

				if tipo_izq.Valor.(Ast.O3D).Valor.Tipo == Ast.I64 {
					valIzq = float64(valorIzq.(int))
				} else {
					valIzq = valorIzq.(float64)
				}
				if tipo_der.Valor.(Ast.O3D).Valor.Tipo == Ast.I64 {
					valDer = float64(valorDer.(int))
				} else {
					valDer = valorDer.(float64)
				}

				if valDer == 0 {
					valor.Valor = 0
				} else {
					valor.Valor = valIzq / valDer
				}
			}

			valor.Tipo = result_dominante

			//Actualizar el código y conseguir el obj O3D
			obj = Ast.ActualizarCodigoAritmetica(tipo_izq, tipo_der, op.operador, false)
			obj.Valor = valor
			return Ast.TipoRetornado{
				Tipo:  Ast.ARITMETICA,
				Valor: obj,
			}

		} else if result_dominante == Ast.NULL {
			//////////////////////////////////////ERROR//////////////////////////////////////
			vI := tipo_izq.Valor.(Ast.O3D).Valor
			vD := tipo_der.Valor.(Ast.O3D).Valor
			vI.Fila = op.Fila
			vI.Columna = op.Columna
			vD.Fila = op.Fila
			vD.Columna = op.Columna
			return errores.GenerarError(8, vI, vD, "", "", "", entorno)

		}
	case "%":
		result_dominante = mul_div_dominante[tipo_izq.Valor.(Ast.O3D).Valor.Tipo][tipo_der.Valor.(Ast.O3D).Valor.Tipo]
		if result_dominante != Ast.NULL {
			valorIzq := tipo_izq.Valor.(Ast.O3D).Valor.Valor
			valorDer := tipo_der.Valor.(Ast.O3D).Valor.Valor
			var valor Ast.TipoRetornado
			var valIzq, valDer float64

			if result_dominante == Ast.I64 || result_dominante == Ast.USIZE {
				if tipo_der.Valor.(Ast.O3D).Valor.Valor.(int) == 0 {
					valor.Valor = 0
				} else {
					valor.Valor = valorIzq.(int) % valorDer.(int)
				}

			} else if result_dominante == Ast.F64 {

				if tipo_izq.Valor.(Ast.O3D).Valor.Tipo == Ast.I64 {
					valIzq = float64(valorIzq.(int))
				} else {
					valIzq = valorIzq.(float64)
				}
				if tipo_der.Valor.(Ast.O3D).Valor.Tipo == Ast.I64 {
					valDer = float64(valorDer.(int))
				} else {
					valDer = valorDer.(float64)
				}
				if valDer == 0 {
					valor.Valor = 0
				} else {
					valor.Valor = math.Mod(valIzq, valDer)
				}

			}
			valor.Tipo = result_dominante
			//Actualizar el código y conseguir el obj O3D
			obj := Ast.ActualizarCodigoAritmetica(tipo_izq, tipo_der, op.operador, false)
			obj.Valor = valor
			return Ast.TipoRetornado{
				Tipo:  Ast.ARITMETICA,
				Valor: obj,
			}

		} else if result_dominante == Ast.NULL {
			/////////////////////////////////////ERROR////////////////////////////////////
			vI := tipo_izq.Valor.(Ast.O3D).Valor
			vD := tipo_der.Valor.(Ast.O3D).Valor
			vI.Fila = op.Fila
			vI.Columna = op.Columna
			vD.Fila = op.Fila
			vD.Columna = op.Columna
			return errores.GenerarError(8, vI, vD, "", "", "", entorno)

		}
	case "&&", "||":
		valorIzq := tipo_izq.Valor.(Ast.O3D).Valor.Valor
		valorDer := tipo_der.Valor.(Ast.O3D).Valor.Valor
		var valor Ast.TipoRetornado
		if tipo_der.Valor.(Ast.O3D).Valor.Tipo != Ast.BOOLEAN || tipo_izq.Valor.(Ast.O3D).Valor.Tipo != Ast.BOOLEAN {
			/////////////////////////////////////ERROR////////////////////////////////////
			vI := tipo_izq.Valor.(Ast.O3D).Valor
			vD := tipo_der.Valor.(Ast.O3D).Valor
			vI.Fila = op.Fila
			vI.Columna = op.Columna
			vD.Fila = op.Fila
			vD.Columna = op.Columna
			return errores.GenerarError(9, vI, vD, "", "", "", entorno)

		}

		if op.operador == "&&" {
			valor.Valor = valorIzq.(bool) && valorDer.(bool)
		} else {
			valor.Valor = valorIzq.(bool) || valorDer.(bool)
		}
		valor.Tipo = Ast.BOOLEAN
		//Actualizar el código y conseguir el obj O3D
		obj := Ast.ActualizarCodigoLogica(tipo_izq, tipo_der, op.operador, false)
		obj.Valor = valor

		return Ast.TipoRetornado{
			Tipo:  Ast.LOGICA,
			Valor: obj,
		}

	case ">", "<", ">=", "<=", "==", "!=":
		var val_der interface{}
		var val_izq interface{}
		valorIzq := tipo_izq.Valor.(Ast.O3D).Valor.Valor
		valorDer := tipo_der.Valor.(Ast.O3D).Valor.Valor
		var valor Ast.TipoRetornado
		result_dominante = suma_dominante_comparacion[tipo_izq.Valor.(Ast.O3D).Valor.Tipo][tipo_der.Valor.(Ast.O3D).Valor.Tipo]
		if result_dominante == Ast.I64 || result_dominante == Ast.F64 ||
			result_dominante == Ast.USIZE {

			if tipo_izq.Valor.(Ast.O3D).Valor.Tipo == Ast.F64 || tipo_der.Valor.(Ast.O3D).Valor.Tipo == Ast.F64 {

				if tipo_izq.Valor.(Ast.O3D).Valor.Tipo == Ast.F64 {
					val_izq = valorIzq.(float64)
				} else {
					val_izq = (float64)(valorIzq.(int))
				}

				if tipo_izq.Valor.(Ast.O3D).Valor.Tipo == Ast.F64 {
					val_der = valorDer.(float64)
				} else {
					val_der = (float64)(valorDer.(int))
				}
				switch op.operador {
				case ">":
					valor.Valor = val_izq.(float64) > val_der.(float64)
				case "<":
					valor.Valor = val_izq.(float64) < val_der.(float64)
				case ">=":
					valor.Valor = val_izq.(float64) >= val_der.(float64)
				case "<=":
					valor.Valor = val_izq.(float64) <= val_der.(float64)
				case "==":
					valor.Valor = val_izq.(float64) == val_der.(float64)
				case "!=":
					valor.Valor = val_izq.(float64) != val_der.(float64)
				}
			} else {
				val_izq = valorIzq.(int)
				val_der = valorDer.(int)
				switch op.operador {
				case ">":
					valor.Valor = val_izq.(int) > val_der.(int)
				case "<":
					valor.Valor = val_izq.(int) < val_der.(int)
				case ">=":
					valor.Valor = val_izq.(int) >= val_der.(int)
				case "<=":
					valor.Valor = val_izq.(int) <= val_der.(int)
				case "==":
					valor.Valor = val_izq.(int) == val_der.(int)
				case "!=":
					valor.Valor = val_izq.(int) != val_der.(int)
				}
			}
			valor.Tipo = Ast.BOOLEAN

			//Actualizar el código y conseguir el obj O3D
			obj := Ast.ActualizarCodigoRelacional(tipo_izq, tipo_der, op.operador, false)
			obj.Valor = valor
			obj.RelacionalExp = ""
			obj.EsRelacionalSimple = "si"
			return Ast.TipoRetornado{
				Tipo:  Ast.RELACIONAL,
				Valor: obj,
			}

			//////////////////////////////////PENDIENTE RELACIONAL CON STRING/////////////////////////////
		} else if result_dominante == Ast.STR || result_dominante == Ast.STRING || result_dominante == Ast.CHAR {
			//Es una comparación entre STR
			val_izq = tipo_izq.Valor.(string)
			val_der = tipo_der.Valor.(string)
			switch op.operador {
			case ">":
				return Ast.TipoRetornado{
					Tipo:  Ast.BOOLEAN,
					Valor: len(val_izq.(string)) > len(val_der.(string)),
				}
			case "<":
				return Ast.TipoRetornado{
					Tipo:  Ast.BOOLEAN,
					Valor: len(val_izq.(string)) < len(val_der.(string)),
				}
			case ">=":
				return Ast.TipoRetornado{
					Tipo:  Ast.BOOLEAN,
					Valor: len(val_izq.(string)) >= len(val_der.(string)),
				}
			case "<=":
				return Ast.TipoRetornado{
					Tipo:  Ast.BOOLEAN,
					Valor: len(val_izq.(string)) <= len(val_der.(string)),
				}
			case "==":
				return Ast.TipoRetornado{
					Tipo:  Ast.BOOLEAN,
					Valor: len(val_izq.(string)) == len(val_der.(string)),
				}
			case "!=":
				return Ast.TipoRetornado{
					Tipo:  Ast.BOOLEAN,
					Valor: len(val_izq.(string)) != len(val_der.(string)),
				}
			}
			return Ast.TipoRetornado{
				Tipo:  Ast.BOOLEAN,
				Valor: len(val_izq.(string)) > len(val_der.(string)),
			}
		} else if result_dominante == Ast.BOOLEAN {

			switch op.operador {

			case "==":
				valor.Valor = valorIzq.(bool) == valorDer.(bool)
			case "!=":
				valor.Valor = valorIzq.(bool) != valorDer.(bool)
			}

			valor.Tipo = Ast.BOOLEAN

			//Actualizar el código y conseguir el obj O3D
			obj := Ast.ActualizarCodigoRelacional(tipo_izq, tipo_der, op.operador, false)
			obj.Valor = valor
			obj.RelacionalExp = ""
			obj.EsRelacionalSimple = "si"
			return Ast.TipoRetornado{
				Tipo:  Ast.RELACIONAL,
				Valor: obj,
			}

		}
		/////////////////////////////////////ERROR////////////////////////////////////
		vI := tipo_izq.Valor.(Ast.O3D).Valor
		vD := tipo_der.Valor.(Ast.O3D).Valor
		vI.Fila = op.Fila
		vI.Columna = op.Columna
		vD.Fila = op.Fila
		vD.Columna = op.Columna
		return errores.GenerarError(11, vI, vD, "", "", "", entorno)
	}
	/*Nunca debería de llegar hasta aquí*/
	return Ast.TipoRetornado{Tipo: Ast.NULL, Valor: nil}
}
