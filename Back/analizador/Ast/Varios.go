package Ast

import (
	"strconv"
	"strings"
)

func AgregarFuncion(funcion string) {
	FuncionesC3D += funcion + "\n"
}

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

func GetFuncionesC3D() string {
	return FuncionesC3D
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

func ResetAll() {
	Label = 1
	Temporal = 1
	Temporales = ""
	P = 0
	H = 0
	FuncionesC3D = ""
}

func GetEncabezado() string {
	salida := "#include <stdio.h>\n"
	salida += "float stack[10000];\n"
	salida += "float heap[10000];\n"
	salida += "float P;\n"
	salida += "float H;\n"
	salida += Temporales + ";\n\n"
	salida += "int main(){\n"
	return salida
}

func GetFinEncabezado() string {
	salida := "\treturn 0;\n"
	salida += "}\n"
	return salida
}

func Indentar(nivel int, cadena string) string {
	ident := ""
	salida := ""
	for i := 0; i < nivel; i++ {
		ident += "\t"
	}

	for _, linea := range strings.Split(strings.TrimSuffix(cadena, "\n"), "\n") {
		salida += ident + linea + "\n"
	}
	return salida
}

/*Retorna la siguiente dirección del apuntador P (int)*/
func GetP() int {
	newP := P
	P++
	return newP
}

func GetPactual() int {
	return P
}

/* Reiniciar P cada vez que entra en un nuevo entorno*/
func ReiniciarP() {
	P = 0
}

/* Retorna la siguiente dirección del apuntador H (int)*/
func GetH() int {
	newH := H
	H++
	return newH
}

func GetValorH() int {
	return H
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
		/*
			if op1.Tipo != PRIMITIVO {
				codigo += codIzq.Codigo + "\n"
			}
			if op2.Tipo != PRIMITIVO {
				codigo += codDer.Codigo + "\n"
			}
		*/

		//Agregar el código anterior
		codigo += codIzq.Codigo + "\n"
		codigo += codDer.Codigo + "\n"
		codigo += "/**************************OPERACION ARITMETICA*/\n"
		//Crear el nuevo código con las referencias de los anteriores
		referencia = GetTemp()
		if operador == "/" || operador == "%" {
			codigo += MathError(codIzq.Referencia, codDer.Referencia, referencia, operador)
		} else {
			codigo += referencia + " = " + codIzq.Referencia + " " + operador + " " + codDer.Referencia + " ;\n"
		}

	} else {
		codIzq = op1.Valor.(O3D)
		if op1.Tipo != PRIMITIVO {
			codigo += codIzq.Codigo + "\n"
			codigo += "/**************************OPERACION ARITMETICA*/\n"
		}
		referencia = GetTemp()
		codigo += referencia + " = " + codIzq.Referencia + " * " + " -1 ;"
	}
	codigo += "/***********************************************/\n"
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

	/*
		if op1.Tipo != PRIMITIVO {
			codigo += codIzq.Codigo + "\n"
		}
		if op2.Tipo != PRIMITIVO {
			codigo += codDer.Codigo + "\n"
		}
	*/

	//Agregar el código anterior
	codigo += codIzq.Codigo + "\n"
	codigo += codDer.Codigo + "\n"
	codigo += "/**************************OPERACION RELACIONAL*/\n"
	//Get labels
	lt = GetLabel()
	lf = GetLabel()

	codigo += "if " + codIzq.Referencia + " " + operador + " " + codDer.Referencia + " goto " + lt + ";\n"
	codigo += "goto " + lf + ";\n"

	codigo += "/***********************************************/\n"
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
	codigo += "/******************************OPERACION LOGICA*/\n"
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
	codigo += "/***********************************************/\n"
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
