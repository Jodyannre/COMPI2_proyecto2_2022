package expresiones

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"fmt"
	"math"
	"strconv"
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
	} else {
		tipo_izq = op.operando_izq.GetValue(entorno)
		tipo_der = op.operando_der.GetValue(entorno)
	}
	//Verificar primero si no hay error
	if op.unario {
		if tipo_izq.Tipo == Ast.ERROR {
			return tipo_izq
		}
	} else {
		if tipo_izq.Tipo == Ast.ERROR {
			return tipo_izq
		}
		if tipo_der.Tipo == Ast.ERROR {
			return tipo_der
		}
	}

	if tipo_izq.Valor.(Ast.O3D).Valor.Tipo > 7 || tipo_der.Valor.(Ast.O3D).Valor.Tipo > 7 {
		//Error, no se pueden operar porque no es ningún valor operable
		return errores.GenerarError(1, op.operando_izq, op.operando_der, entorno)
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

			//Parseo por si viene algún int
			if tipo_izq.Tipo == Ast.I64 {
				valIzq = float64(valorIzq.(int))
			} else {
				valIzq = valorIzq.(float64)
			}
			if tipo_der.Tipo == Ast.I64 {
				valDer = float64(valorDer.(int))
			} else {
				valDer = valorDer.(float64)
			}

			//Get el valor
			valor := Ast.TipoRetornado{
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
			cadena_izq := fmt.Sprintf("%v", tipo_izq.Valor)
			cadena_der := fmt.Sprintf("%v", tipo_der.Valor)
			return Ast.TipoRetornado{
				Tipo:  result_dominante,
				Valor: cadena_izq + cadena_der,
			}
		} else if result_dominante == Ast.NULL {
			return errores.GenerarError(1, op.operando_izq, op.operando_der, entorno)

		}

	case "-", "!":
		if op.unario {
			//Es una operación unaria
			if tipo_izq.Tipo == Ast.F64 && op.operador == "-" {
				//Es un F64
				return Ast.TipoRetornado{
					Tipo:  Ast.F64,
					Valor: tipo_izq.Valor.(float64) * -1,
				}
			} else if tipo_izq.Tipo == Ast.I64 {
				//Es un int
				if op.operador == "-" {
					return Ast.TipoRetornado{
						Tipo:  Ast.I64,
						Valor: tipo_izq.Valor.(int) * -1,
					}
				} else {
					var valorFinal = tipo_izq.Valor.(int)
					//Verificar la regla del bitwise
					if valorFinal >= 0 {
						valorFinal++
					} else {
						valorFinal++
					}
					return Ast.TipoRetornado{
						Tipo:  Ast.I64,
						Valor: valorFinal * -1,
					}
				}

			} else if tipo_izq.Tipo == Ast.BOOLEAN && op.operador == "!" {
				return Ast.TipoRetornado{
					Tipo:  Ast.BOOLEAN,
					Valor: !tipo_izq.Valor.(bool),
				}

			} else if tipo_izq.Tipo == Ast.USIZE && op.operador == "-" {
				//Error, no se puede aplicar el menos a un usize
				msg := "Semantic error, can't apply unary operator `-` to type `usize`." +
					" -- Line: " + strconv.Itoa(op.Fila) +
					" Column: " + strconv.Itoa(op.Columna)
				nError := errores.NewError(op.Fila, op.Columna, msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				nError.Ambito = entorno.GetTipoScope()
				entorno.Errores.Add(nError)
				entorno.Consola += msg + "\n"
				return Ast.TipoRetornado{
					Tipo:  Ast.ERROR,
					Valor: nError,
				}
			} else {
				//Tipo no operable
				msg := "Semantic error, can't operate (!) with a " + Ast.ValorTipoDato[tipo_izq.Tipo] +
					" type. -- Line: " + strconv.Itoa(op.Fila) +
					" Column: " + strconv.Itoa(op.Columna)
				nError := errores.NewError(op.Fila, op.Columna, msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				nError.Ambito = entorno.GetTipoScope()
				entorno.Errores.Add(nError)
				entorno.Consola += msg + "\n"
				return Ast.TipoRetornado{
					Tipo:  Ast.ERROR,
					Valor: nError,
				}
			}

		}

		result_dominante = resta_dominante[tipo_izq.Tipo][tipo_der.Tipo]

		if result_dominante == Ast.I64 {
			return Ast.TipoRetornado{
				Tipo:  result_dominante,
				Valor: tipo_izq.Valor.(int) - tipo_der.Valor.(int),
			}
		} else if result_dominante == Ast.USIZE {
			preValor := tipo_izq.Valor.(int) - tipo_der.Valor.(int)
			if preValor < 0 {
				//Error, el usize no puede ser negativo
				msg := "Semantic error, attempt to subtract with overflow." +
					" -- Line: " + strconv.Itoa(op.Fila) +
					" Column: " + strconv.Itoa(op.Columna)
				nError := errores.NewError(op.Fila, op.Columna, msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				nError.Ambito = entorno.GetTipoScope()
				entorno.Errores.Add(nError)
				entorno.Consola += msg + "\n"
				return Ast.TipoRetornado{
					Tipo:  Ast.ERROR,
					Valor: nError,
				}
			}
			return Ast.TipoRetornado{
				Tipo:  result_dominante,
				Valor: tipo_izq.Valor.(int) - tipo_der.Valor.(int),
			}

		} else if result_dominante == Ast.F64 {

			if tipo_izq.Tipo == Ast.I64 {
				tipo_izq = Ast.TipoRetornado{
					Valor: float64(tipo_izq.Valor.(int)),
					Tipo:  Ast.F64,
				}
			}
			if tipo_der.Tipo == Ast.I64 {
				tipo_der = Ast.TipoRetornado{
					Valor: float64(tipo_der.Valor.(int)),
					Tipo:  Ast.F64,
				}
			}

			return Ast.TipoRetornado{
				Tipo:  result_dominante,
				Valor: tipo_izq.Valor.(float64) - tipo_der.Valor.(float64),
			}
		} else if result_dominante == Ast.NULL {
			/*
				return Ast.TipoRetornado{
					Tipo:  result_dominante,
					Valor: nil,
				}
			*/
			msg := "Semantic error, can't subtract " + Ast.ValorTipoDato[tipo_izq.Tipo] +
				" type to " + Ast.ValorTipoDato[tipo_der.Tipo] +
				" type. -- Line: " + strconv.Itoa(op.Fila) +
				" Column: " + strconv.Itoa(op.Columna)
			nError := errores.NewError(op.Fila, op.Columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			nError.Ambito = entorno.GetTipoScope()
			entorno.Errores.Add(nError)
			entorno.Consola += msg + "\n"
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}
		}

	case "*":
		result_dominante = mul_div_dominante[tipo_izq.Tipo][tipo_der.Tipo]
		if result_dominante == Ast.I64 || result_dominante == Ast.USIZE {
			return Ast.TipoRetornado{
				Tipo:  result_dominante,
				Valor: tipo_izq.Valor.(int) * tipo_der.Valor.(int),
			}
		} else if result_dominante == Ast.F64 {
			if tipo_izq.Tipo == Ast.I64 {
				tipo_izq = Ast.TipoRetornado{
					Valor: float64(tipo_izq.Valor.(int)),
					Tipo:  Ast.F64,
				}
			}
			if tipo_der.Tipo == Ast.I64 {
				tipo_der = Ast.TipoRetornado{
					Valor: float64(tipo_der.Valor.(int)),
					Tipo:  Ast.F64,
				}
			}
			return Ast.TipoRetornado{
				Tipo:  result_dominante,
				Valor: tipo_izq.Valor.(float64) * tipo_der.Valor.(float64),
			}

		} else if result_dominante == Ast.NULL {
			/*
				return Ast.TipoRetornado{
					Tipo:  result_dominante,
					Valor: nil,
				}
			*/
			msg := "Semantic error, can't multiply " + Ast.ValorTipoDato[tipo_izq.Tipo] +
				" type to " + Ast.ValorTipoDato[tipo_der.Tipo] +
				" type. -- Line: " + strconv.Itoa(op.Fila) +
				" Column: " + strconv.Itoa(op.Columna)
			nError := errores.NewError(op.Fila, op.Columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			nError.Ambito = entorno.GetTipoScope()
			entorno.Errores.Add(nError)
			entorno.Consola += msg + "\n"
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}

		}
	case "/":
		result_dominante = mul_div_dominante[tipo_izq.Tipo][tipo_der.Tipo]
		if result_dominante == Ast.I64 {
			if tipo_der.Valor.(int) == 0 {
				//Error, no se puede dividir dentro de 0
				msg := "Semantic error, can't be divided by zero." +
					" -- Line: " + strconv.Itoa(op.Fila) +
					" Column: " + strconv.Itoa(op.Columna)
				nError := errores.NewError(op.Fila, op.Columna, msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				nError.Ambito = entorno.GetTipoScope()
				entorno.Errores.Add(nError)
				entorno.Consola += msg + "\n"
				return Ast.TipoRetornado{
					Tipo:  Ast.ERROR,
					Valor: nError,
				}
			}
			return Ast.TipoRetornado{
				Tipo:  result_dominante,
				Valor: tipo_izq.Valor.(int) / tipo_der.Valor.(int),
			}

		} else if result_dominante == Ast.USIZE {
			if tipo_der.Valor.(int) == 0 {
				//Error, no se puede dividir dentro de 0
				msg := "Semantic error, can't be divided by zero." +
					" -- Line: " + strconv.Itoa(op.Fila) +
					" Column: " + strconv.Itoa(op.Columna)
				nError := errores.NewError(op.Fila, op.Columna, msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				nError.Ambito = entorno.GetTipoScope()
				entorno.Errores.Add(nError)
				entorno.Consola += msg + "\n"
				return Ast.TipoRetornado{
					Tipo:  Ast.ERROR,
					Valor: nError,
				}
			}
			return Ast.TipoRetornado{
				Tipo:  result_dominante,
				Valor: int(math.Trunc(float64(tipo_izq.Valor.(int)) / float64(tipo_der.Valor.(int)))),
			}

		} else if result_dominante == Ast.F64 {

			if tipo_izq.Tipo == Ast.I64 {
				tipo_izq = Ast.TipoRetornado{
					Valor: float64(tipo_izq.Valor.(int)),
					Tipo:  Ast.F64,
				}
			}
			if tipo_der.Tipo == Ast.I64 {
				tipo_der = Ast.TipoRetornado{
					Valor: float64(tipo_der.Valor.(int)),
					Tipo:  Ast.F64,
				}
			}

			if tipo_der.Valor.(float64) == 0 {
				//Error, no se puede dividir dentro de 0
				msg := "Semantic error, can't be divided by zero." +
					" -- Line: " + strconv.Itoa(op.Fila) +
					" Column: " + strconv.Itoa(op.Columna)
				nError := errores.NewError(op.Fila, op.Columna, msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				nError.Ambito = entorno.GetTipoScope()
				entorno.Errores.Add(nError)
				entorno.Consola += msg + "\n"
				return Ast.TipoRetornado{
					Tipo:  Ast.ERROR,
					Valor: nError,
				}
			}
			return Ast.TipoRetornado{
				Tipo:  result_dominante,
				Valor: tipo_izq.Valor.(float64) / tipo_der.Valor.(float64),
			}

		} else if result_dominante == Ast.NULL {
			/*
				return Ast.TipoRetornado{
					Tipo:  result_dominante,
					Valor: nil,
				}
			*/
			msg := "Semantic error, can't divide " + Ast.ValorTipoDato[tipo_izq.Tipo] +
				" type by " + Ast.ValorTipoDato[tipo_der.Tipo] +
				" type. -- Line: " + strconv.Itoa(op.Fila) +
				" Column: " + strconv.Itoa(op.Columna)
			nError := errores.NewError(op.Fila, op.Columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			nError.Ambito = entorno.GetTipoScope()
			entorno.Errores.Add(nError)
			entorno.Consola += msg + "\n"
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}

		}
	case "%":
		result_dominante = mul_div_dominante[tipo_izq.Tipo][tipo_der.Tipo]
		if result_dominante == Ast.I64 || result_dominante == Ast.USIZE {
			return Ast.TipoRetornado{
				Tipo:  result_dominante,
				Valor: tipo_izq.Valor.(int) % tipo_der.Valor.(int),
			}

		} else if result_dominante == Ast.F64 {
			if tipo_izq.Tipo == Ast.I64 {
				tipo_izq = Ast.TipoRetornado{
					Valor: float64(tipo_izq.Valor.(int)),
					Tipo:  Ast.F64,
				}
			}
			if tipo_der.Tipo == Ast.I64 {
				tipo_der = Ast.TipoRetornado{
					Valor: float64(tipo_der.Valor.(int)),
					Tipo:  Ast.F64,
				}
			}
			return Ast.TipoRetornado{
				Tipo:  result_dominante,
				Valor: math.Mod(tipo_izq.Valor.(float64), tipo_der.Valor.(float64)),
			}

		} else if result_dominante == Ast.NULL {
			/*
				return Ast.TipoRetornado{
					Tipo:  result_dominante,
					Valor: nil,
				}
			*/
			msg := "Semantic error, can't divide " + Ast.ValorTipoDato[tipo_izq.Tipo] +
				" type by " + Ast.ValorTipoDato[tipo_der.Tipo] +
				" type. -- Line: " + strconv.Itoa(op.Fila) +
				" Column: " + strconv.Itoa(op.Columna)
			nError := errores.NewError(op.Fila, op.Columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			nError.Ambito = entorno.GetTipoScope()
			entorno.Errores.Add(nError)
			entorno.Consola += msg + "\n"
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}

		}
	case "&&", "||":
		if tipo_der.Tipo != Ast.BOOLEAN || tipo_izq.Tipo != Ast.BOOLEAN {
			msg := "Semantic error, can't logically operate " + Ast.ValorTipoDato[tipo_izq.Tipo] +
				" type with " + Ast.ValorTipoDato[tipo_der.Tipo] +
				" type. -- Line: " + strconv.Itoa(op.Fila) +
				" Column: " + strconv.Itoa(op.Columna)
			nError := errores.NewError(op.Fila, op.Columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			nError.Ambito = entorno.GetTipoScope()
			entorno.Errores.Add(nError)
			entorno.Consola += msg + "\n"
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}

		}
		if op.operador == "&&" {
			return Ast.TipoRetornado{
				Tipo:  Ast.BOOLEAN,
				Valor: tipo_izq.Valor.(bool) && tipo_der.Valor.(bool),
			}
		} else {
			return Ast.TipoRetornado{
				Tipo:  Ast.BOOLEAN,
				Valor: tipo_izq.Valor.(bool) || tipo_der.Valor.(bool),
			}
		}

	case ">", "<", ">=", "<=", "==", "!=":
		var val_der interface{}
		var val_izq interface{}
		result_dominante = suma_dominante_comparacion[tipo_izq.Tipo][tipo_der.Tipo]
		if result_dominante == Ast.I64 || result_dominante == Ast.F64 ||
			result_dominante == Ast.USIZE {

			if tipo_izq.Tipo == Ast.F64 || tipo_der.Tipo == Ast.F64 {

				if tipo_izq.Tipo == Ast.F64 {
					val_izq = tipo_izq.Valor.(float64)
				} else {
					val_izq = (float64)(tipo_izq.Valor.(int))
				}

				if tipo_izq.Tipo == Ast.F64 {
					val_der = tipo_der.Valor.(float64)
				} else {
					val_der = (float64)(tipo_der.Valor.(int))
				}
				switch op.operador {
				case ">":
					return Ast.TipoRetornado{
						Tipo:  Ast.BOOLEAN,
						Valor: val_izq.(float64) > val_der.(float64),
					}
				case "<":
					return Ast.TipoRetornado{
						Tipo:  Ast.BOOLEAN,
						Valor: val_izq.(float64) < val_der.(float64),
					}
				case ">=":
					return Ast.TipoRetornado{
						Tipo:  Ast.BOOLEAN,
						Valor: val_izq.(float64) >= val_der.(float64),
					}
				case "<=":
					return Ast.TipoRetornado{
						Tipo:  Ast.BOOLEAN,
						Valor: val_izq.(float64) <= val_der.(float64),
					}
				case "==":
					return Ast.TipoRetornado{
						Tipo:  Ast.BOOLEAN,
						Valor: val_izq.(float64) == val_der.(float64),
					}
				case "!=":
					return Ast.TipoRetornado{
						Tipo:  Ast.BOOLEAN,
						Valor: val_izq.(float64) != val_der.(float64),
					}
				}
			} else {
				val_izq = tipo_izq.Valor.(int)
				val_der = tipo_der.Valor.(int)
				switch op.operador {
				case ">":
					return Ast.TipoRetornado{
						Tipo:  Ast.BOOLEAN,
						Valor: val_izq.(int) > val_der.(int),
					}
				case "<":
					return Ast.TipoRetornado{
						Tipo:  Ast.BOOLEAN,
						Valor: val_izq.(int) < val_der.(int),
					}
				case ">=":
					return Ast.TipoRetornado{
						Tipo:  Ast.BOOLEAN,
						Valor: val_izq.(int) >= val_der.(int),
					}
				case "<=":
					return Ast.TipoRetornado{
						Tipo:  Ast.BOOLEAN,
						Valor: val_izq.(int) <= val_der.(int),
					}
				case "==":
					return Ast.TipoRetornado{
						Tipo:  Ast.BOOLEAN,
						Valor: val_izq.(int) == val_der.(int),
					}
				case "!=":
					return Ast.TipoRetornado{
						Tipo:  Ast.BOOLEAN,
						Valor: val_izq.(int) != val_der.(int),
					}
				}
			}
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
				return Ast.TipoRetornado{
					Tipo:  Ast.BOOLEAN,
					Valor: len(val_izq.(string)) == len(val_der.(string)),
				}
			case "!=":
				return Ast.TipoRetornado{
					Tipo:  Ast.BOOLEAN,
					Valor: len(val_izq.(string)) != len(val_der.(string)),
				}
			default:
				msg := "Semantic error, can't compare a " + Ast.ValorTipoDato[tipo_izq.Tipo] +
					" using " + op.operador +
					" type. -- Line: " + strconv.Itoa(op.Fila) +
					" Column: " + strconv.Itoa(op.Columna)
				nError := errores.NewError(op.Fila, op.Columna, msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				nError.Ambito = entorno.GetTipoScope()
				entorno.Errores.Add(nError)
				entorno.Consola += msg + "\n"
				return Ast.TipoRetornado{
					Tipo:  Ast.ERROR,
					Valor: nError,
				}
			}
			/*
				return Ast.TipoRetornado{
					Tipo:  Ast.BOOLEAN,
					Valor: len(val_izq.(string)) > len(val_der.(string)),
				}
			*/
		}
		msg := "Semantic error, can't compare " + Ast.ValorTipoDato[tipo_izq.Tipo] +
			" with " + Ast.ValorTipoDato[tipo_der.Tipo] +
			" type. -- Line: " + strconv.Itoa(op.Fila) +
			" Column: " + strconv.Itoa(op.Columna)
		nError := errores.NewError(op.Fila, op.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = entorno.GetTipoScope()
		entorno.Errores.Add(nError)
		entorno.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}
	return Ast.TipoRetornado{Tipo: Ast.NULL, Valor: nil}
}
