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
		if operador == "/" || operador == "%" {
			codigo += MathError(codIzq.Referencia, codDer.Referencia, referencia, operador)
		} else {
			codigo += referencia + " = " + codIzq.Referencia + " " + operador + " " + codDer.Referencia + " ;"
		}

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

/* Actualizar el código actual con los que se van retornando de las operaciones anteriores.
-->Retonar un objeto O3D<--*/
func ActualizarCodigoRelacional(op1 TipoRetornado, op2 TipoRetornado, operador string,
	unario bool) O3D {
	var codigo, lt, lf string = "", "", ""
	var codIzq, codDer O3D

	codIzq = op1.Valor.(O3D)
	codDer = op2.Valor.(O3D)

	if op1.Tipo != PRIMITIVO {
		codigo += codIzq.Codigo + "\n"
	}
	if op2.Tipo != PRIMITIVO {
		codigo += codDer.Codigo + "\n"
	}

	//Get labels
	lt = GetLabel()
	lf = GetLabel()

	codigo += "if " + codIzq.Referencia + " " + operador + " " + codDer.Referencia + " goto " + lt + ";\n"
	codigo += "goto " + lf + ";\n"

	//Crear el nuevo objeto 3D
	obj := O3D{
		Lt:         lt,
		Lf:         lf,
		Valor:      TipoRetornado{},
		Codigo:     codigo,
		Referencia: "",
	}

	return obj
}

/* Actualizar el código actual con los que se van retornando de las operaciones anteriores.
-->Retonar un objeto O3D<--*/
func ActualizarCodigoLogica(op1 TipoRetornado, op2 TipoRetornado, operador string,
	unario bool) O3D {
	var codigo, lt, lf string = "", "", ""
	var codIzq, codDer O3D

	codIzq = op1.Valor.(O3D)
	codDer = op2.Valor.(O3D)

	if !unario {
		if operador == "&&" {
			/*AND*/
			codigo += codIzq.Codigo
			codigo += codIzq.Lt + ":\n"
			codigo += codDer.Codigo
			/*Actualizar los labels*/
			lt += codDer.Lt
			lf += codIzq.Lt + ":\n"
			lf += codDer.Lt
		} else {
			/*OR*/
			codigo += codIzq.Codigo
			codigo += codIzq.Lf + ":\n"
			codigo += codDer.Codigo
			/*Actualizar los labels*/
			lt += codIzq.Lt + ":\n"
			lt += codDer.Lt
			lf += codDer.Lf
		}

	} else {
		/*NOT*/
		codigo = codIzq.Codigo
		lt = codIzq.Lf
		lf = codIzq.Lt
	}

	//Crear el nuevo objeto 3D
	obj := O3D{
		Lt:         lt,
		Lf:         lf,
		Valor:      TipoRetornado{},
		Codigo:     codigo,
		Referencia: "",
	}

	return obj
}

func MathError(refIzq, refDer, refActual, operador string) string {
	lT := GetLabel()
	lF := GetLabel()
	//ident := "    "
	p := "%"
	c := "c"
	salida := "if (" + refDer + " != 0) goto " + lT + ";\n"
	salida += "Printf(\"" + p + c + "\",  77);\n"
	salida += "Printf(\"" + p + c + "\",  97);\n"
	salida += "Printf(\"" + p + c + "\", 116);\n"
	salida += "Printf(\"" + p + c + "\", 104);\n"
	salida += "Printf(\"" + p + c + "\",  69);\n"
	salida += "Printf(\"" + p + c + "\", 114);\n"
	salida += "Printf(\"" + p + c + "\", 114);\n"
	salida += "Printf(\"" + p + c + "\", 111);\n"
	salida += "Printf(\"" + p + c + "\", 114);\n"
	salida += refActual + " = 0;\n"
	salida += "goto " + lF + ";\n"
	salida += lT + ":\n"
	salida += refActual + " = " + refIzq + " " + operador + " " + refDer + ";\n"
	salida += lF + ":"
	return salida
}
