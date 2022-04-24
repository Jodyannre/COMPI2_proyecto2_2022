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

	return salida
}

func GetFinFuncionMain() string {
	salida := "\treturn 0;\n"
	salida += "}\n"
	return salida
}

func GetInicioMain() string {
	salida := "int main(){\n"
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
	var relacionalDirecto string
	var referencia string
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

	codigo += "if (" + codIzq.Referencia + " " + operador + " " + codDer.Referencia + ") goto " + lt + ";\n"
	codigo += "goto " + lf + ";\n"

	/*******************************CODIGO RELACIONAL DIRECTO***********************************/
	salto := GetTemp()
	referencia = GetTemp()
	relacionalDirecto += "/**************************RESULTADO RELACIONAL*/\n"
	relacionalDirecto += lt + ": \n"
	relacionalDirecto += referencia + " = 1;\n"
	relacionalDirecto += "goto " + salto + ";\n"
	relacionalDirecto += lf + ": \n"
	relacionalDirecto += referencia + " = 0;\n"
	relacionalDirecto += salto + ":\n"
	relacionalDirecto += "/***********************************************/\n"
	/*******************************************************************************************/

	//Crear el nuevo objeto 3D
	obj := O3D{
		Lt:            lt,
		Lf:            lf,
		Valor:         TipoRetornado{},
		Codigo:        codigo,
		Referencia:    referencia,
		RelacionalExp: relacionalDirecto,
	}

	return obj
}

func SumaStrings(op1 O3D, op2 O3D) O3D {
	var obj3d O3D
	var referenciaIzq = op1.Referencia
	var referenciaDer = op2.Referencia
	var codigo3d string
	var letraActual string
	var contadorPrimerString string
	var inicioNuevoString string
	var contadorNuevoString string
	var aumentarPrimerStringContador string
	var contadorSegundoString string
	var aumentarSegundoStringContador string
	var lt1, lf1, salto1 string
	var lt2, lf2, salto2 string

	/********************INICIALIZAR VARIABLES 3D*****************/
	inicioNuevoString = GetTemp()
	contadorNuevoString = GetTemp()
	contadorPrimerString = GetTemp()
	letraActual = GetTemp()
	aumentarPrimerStringContador = GetTemp()
	contadorSegundoString = GetTemp()
	aumentarSegundoStringContador = GetTemp()
	lt1 = GetLabel()
	lf1 = GetLabel()
	salto1 = GetLabel()
	lt2 = GetLabel()
	lf2 = GetLabel()
	salto2 = GetLabel()

	/*************************************************************/
	codigo3d += "/**************************OPERACION ARITMETICA*/\n"
	codigo3d += "/*******************************PRIMER ELEMENTO*/\n"
	codigo3d += inicioNuevoString + " = H; //Guardar inicio del nuevo string \n"
	codigo3d += contadorNuevoString + " = H; //Iniciar contador nuevo string \n"
	codigo3d += contadorPrimerString + " = " + referenciaIzq + "; //Copiar referencia valor\n"
	codigo3d += salto1 + ":\n"
	codigo3d += letraActual + " = heap[(int)" + contadorPrimerString + "]; //Get letra\n"
	codigo3d += "if (" + letraActual + "!=0) goto " + lt1 + ";\n"
	codigo3d += "goto " + lf1 + ";\n"
	codigo3d += lt1 + ":\n"
	codigo3d += "heap[(int)" + contadorNuevoString + "] = " + letraActual + ";\n"
	codigo3d += "H = H + 1; \n"
	codigo3d += contadorNuevoString + " = H; //Aumentar contador nuevo string \n"
	codigo3d += aumentarPrimerStringContador + " = " + contadorPrimerString + " + 1; //Aumentar contador primer string \n"
	codigo3d += contadorPrimerString + " = " + aumentarPrimerStringContador + "; //Actualizar contador \n"
	GetH()
	codigo3d += "goto " + salto1 + ";\n"
	codigo3d += "/***********************************************/\n"
	codigo3d += lf1 + ":\n"
	codigo3d += "/******************************SEGUNDO ELEMENTO*/\n"
	codigo3d += contadorSegundoString + " = " + referenciaDer + "; //Copiar referencia valor\n"
	codigo3d += salto2 + ":\n"
	codigo3d += letraActual + " = heap[(int)" + contadorSegundoString + "]; //Get letra\n"
	codigo3d += "if (" + letraActual + "!=0) goto " + lt2 + ";\n"
	codigo3d += "goto " + lf2 + ";\n"
	codigo3d += lt2 + ":\n"
	codigo3d += "heap[(int)" + contadorNuevoString + "] = " + letraActual + ";\n"
	codigo3d += "H = H + 1; \n"
	codigo3d += contadorNuevoString + " = H; //Aumentar contador nuevo string \n"
	codigo3d += aumentarSegundoStringContador + " = " + contadorSegundoString + " + 1; //Aumentar contador segundo string \n"
	codigo3d += contadorSegundoString + " = " + aumentarSegundoStringContador + "; //Actualizar contador \n"
	GetH()
	codigo3d += "goto " + salto2 + ";\n"
	codigo3d += lf2 + ":\n"
	codigo3d += "heap[(int)" + contadorNuevoString + "] = 0 ; //Agregar el fin de la cadena \n"
	codigo3d += "H = H + 1;\n"
	codigo3d += "/***********************************************/\n"
	codigo3d += "/***********************************************/\n"
	obj3d.Codigo = codigo3d
	obj3d.Referencia = inicioNuevoString
	return obj3d
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
			//Verificar que no sean solo identificadores
			if codIzq.Lt == "" {
				//Es un valor por identificador
				codIzq = GetCod3DLogicaId(codIzq)
			}

			if codDer.Lt == "" {
				//Es un valor por identificador
				codDer = GetCod3DLogicaId(codDer)
			}

			codigo += codIzq.Codigo
			codigo += codIzq.Lt + ":\n"
			codigo += codDer.Codigo
			/*Actualizar los labels*/
			lt += codDer.Lt
			lf += codIzq.Lf + ":\n"
			lf += codDer.Lf
		} else {
			/*OR*/
			if codIzq.Lt == "" {
				//Es un valor por identificador
				codIzq = GetCod3DLogicaId(codIzq)
			}

			if codDer.Lt == "" {
				//Es un valor por identificador
				codDer = GetCod3DLogicaId(codDer)
			}
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
		if codIzq.Lt != "" {
			codigo = codIzq.Codigo
			lt = codIzq.Lf
			lf = codIzq.Lt
		} else if op1.Tipo != PRIMITIVO {
			ltTemp := GetLabel()
			lfTemp := GetLabel()
			codigo = codIzq.Codigo
			codigo += "if (" + codIzq.Referencia + " == 1) goto " + ltTemp + ";\n"
			codigo += "goto " + lfTemp + ";\n"
			lt = ltTemp
			lf = lfTemp
		}

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

	if codIzq.EsContains != "" {
		obj.EsContains = codIzq.EsContains
	}
	if codDer.EsContains != "" {
		obj.EsContains = codDer.EsContains
	}

	return obj
}

func MathError(refIzq, refDer, refActual, operador string) string {
	lT := GetLabel()
	lF := GetLabel()
	salto := GetLabel()
	//ident := "    "
	p := "%%"
	c := "c"
	if operador == "%" {
		operador = p
	}
	salida := "if (" + refDer + " == 0) goto " + lT + ";\n"
	salida += "goto " + lF + ";\n"
	salida += lT + ":\n"
	salida += "printf(\"" + p + c + "\",  77);\n"
	salida += "printf(\"" + p + c + "\",  97);\n"
	salida += "printf(\"" + p + c + "\", 116);\n"
	salida += "printf(\"" + p + c + "\", 104);\n"
	salida += "printf(\"" + p + c + "\",  69);\n"
	salida += "printf(\"" + p + c + "\", 114);\n"
	salida += "printf(\"" + p + c + "\", 114);\n"
	salida += "printf(\"" + p + c + "\", 111);\n"
	salida += "printf(\"" + p + c + "\", 114);\n"
	salida += "printf(\"" + p + c + "\", 10);\n"
	salida += refActual + " = 0; //Valor de error 0 \n"
	salida += "goto " + salto + ";\n"
	salida += lF + ":\n"
	salida += refActual + " = (int)" + refIzq + " " + operador + " (int)" + refDer + ";\n"
	salida += salto + ":"
	return salida
}

func GetCod3DLogicaId(obj3d O3D) O3D {
	var lt, lf, codigo3d string
	lt = GetLabel()
	lf = GetLabel()
	codigo3d = obj3d.Codigo
	codigo3d += "if (" + obj3d.Referencia + " == 1) goto " + lt + ";\n"
	codigo3d += "goto " + lf + ";\n"
	obj3d.Lt = lt
	obj3d.Lf = lf
	obj3d.Codigo = codigo3d
	return obj3d
}
