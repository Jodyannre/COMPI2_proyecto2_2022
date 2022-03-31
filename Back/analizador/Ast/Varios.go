package Ast

import "strconv"

func EsTransferencia(tipo TipoDato) bool {
	if tipo == BREAK ||
		tipo == BREAK_EXPRESION ||
		tipo == RETURN ||
		tipo == RETURN_EXPRESION ||
		tipo == CONTINUE {
		return true
	} else {
		return false
	}
}

func EsPrimitivo(tipo TipoDato) bool {
	if tipo <= 6 {
		return true
	} else {
		return false
	}
}

func EsFuncion(tipo interface{}) bool {
	validador := false

	switch tipo {
	case FUNCION, VEC_NEW, VEC_ACCESO,
		VEC_LEN, VEC_CONTAINS,
		VEC_CAPACITY, VEC_REMOVE, ARRAY_FAC, ARRAY_ELEMENTOS, ARRAY,
		VEC_ELEMENTOS, VEC_FAC, VEC_WITH_CAPACITY, DIMENSION_ARRAY, LLAMADA_FUNCION:
		validador = true
	default:
		validador = false
	}

	return validador
}

func GetLabel() string {
	newLabel := "L" + strconv.Itoa(Label)
	Label++
	return newLabel
}

/*Método que crea, guarda y retorna un nuevo temporal*/
func GetTemp() string {
	//Crear el nuevo temporal
	newTemporal := "T" + strconv.Itoa(Temporal)
	if Temporales == "" {
		Temporales += "float " + newTemporal
	} else {
		Temporales += ", " + newTemporal
	}
	Temporal++
	return newTemporal
}

/*Retorna la siguiente dirección del apuntador P (int)*/
func GetP() int {
	newP := P
	P++
	return newP
}

/* Retorna la siguiente dirección del apuntador H (int)*/
func GetH() int {
	newH := H
	H++
	return newH
}

/* Actualizar el código actual con los que se van retornando de las operaciones anteriores.
-->Retonar un objeto O3D<--*/
func ActualizarCodigoAritmetica(op1 TipoRetornado, op2 TipoRetornado, operador string,
	unario bool) O3D {
	var codigo, referencia = "", ""
	var codIzq, codDer O3D

	if !unario {
		codIzq = op1.Valor.(O3D)
		codDer = op2.Valor.(O3D)
		//Recuperar el codigo anterior de las 2 expresiones
		if op1.Tipo != PRIMITIVO {
			codigo += codIzq.Codigo + "\n"
		}
		if op2.Tipo != PRIMITIVO {
			codigo += codDer.Codigo + "\n"
		}
		//Crear el nuevo código con las referencias de los anteriores
		referencia = GetTemp()
		codigo += referencia + " = " + codIzq.Referencia + " " + operador + " " + codDer.Referencia + " ;"

	} else {
		codIzq = op1.Valor.(O3D)
		if op1.Tipo != PRIMITIVO {
			codigo += codIzq.Codigo + "\n"
		}
		referencia = GetTemp()
		codigo += referencia + " = " + codIzq.Referencia + " * " + " -1 ;"
	}

	//Crear el nuevo objeto 3D
	obj := O3D{
		Lt:         "",
		Lf:         "",
		Valor:      TipoRetornado{},
		Codigo:     codigo,
		Referencia: referencia,
	}

	return obj
}
