package Ast

import (
	"fmt"
	"strconv"
	"strings"
)

var P, H = 0, 0
var Label, Temporal int = 1, 1
var Temporales string = ""
var FuncionesC3D string = ""

type TipoDato int

var ValorTipoDato = [100]string{
	"I64",
	"F64",
	"STRING_OWNED",
	"STRING",
	"STR",
	"BOOLEAN",
	"USIZE",
	"CHAR",
	"ARRAY",
	"VECTOR",
	"STRUCT",
	"IDENTIFICADOR",
	"LOOP_EXPRESION",
	"IF_EXPRESION",
	"ELSE_EXPRESION",
	"ELSEIF_EXPRESION",
	"MATCH_EXPRESION",
	"ACCESO_VECTOR",
	"ACCESO_MODULO",
	"LLAMADA",
	"MODULO",
	"FUNCION",
	"RETURN",
	"CONTINUE",
	"BREAK",
	"LOOP",
	"WHILE",
	"FOR",
	"MATCH",
	"DECLARACION_VARIABLE",
	"DECLARACION_FUNCION",
	"DECLARACION_MODULO",
	"DECLARACION_VECTOR",
	"DECLARACION",
	"ASIGNACION",
	"IF",
	"ELSE",
	"ELSEIF",
	"PRINT",
	"EXPRESION",
	"INSTRUCCION",
	"PRIMITIVO",
	"NULL",
	"ERROR",
	"ERROR_SEMANTICO",
	"ERROR_LEXICO",
	"ERROR_SINTACTICO",
	"VARIABLE",
	"INDEFINIDO",
	"VOID",
	"RETORNO",
	"EJECUTADO",
	"CASE",
	"CASE_EXPRESION",
	"DEFAULT",
	"BREAK_EXPRESION",
	"RETURN_EXPRESION",
	"PRINTF",
	"PRINT_PRIMITIVOS",
	"PRINT_ARRAY",
	"ERROR_SEMANTICO_NO",
	"LLAMADA_FUNCION",
	"PARAMETRO",
	"CAST",
	"SIMBOLO",
	"VALOR",
	"ERROR_NO_EXISTE",
	"ERROR_ACCESO_PRIVADO",
	"MUTABLE",
	"NOT_MUTABLE",
	"OCUPADO",
	"LIBRE",
	"VEC_NEW",
	"VEC_PUSH",
	"VEC_LEN",
	"VEC_CONTAINS",
	"VEC_CAPACITY",
	"VEC_INSERT",
	"VEC_REMOVE",
	"VEC_ACCESO",
	"VEC_FAC",
	"VEC_ELEMENTOS",
	"VEC_WITH_CAPACITY",
	"POW",
}

const (
	I64 TipoDato = iota
	F64
	STRING_OWNED
	STRING
	STR
	BOOLEAN
	USIZE
	CHAR
	ARRAY
	VECTOR
	STRUCT
	IDENTIFICADOR
	LOOP_EXPRESION
	IF_EXPRESION
	ELSE_EXPRESION
	ELSEIF_EXPRESION
	MATCH_EXPRESION
	ACCESO_VECTOR
	ACCESO_MODULO
	LLAMADA
	MODULO
	FUNCION
	RETURN
	CONTINUE
	BREAK
	LOOP
	WHILE
	FOR
	MATCH
	DECLARACION_VARIABLE
	DECLARACION_FUNCION
	DECLARACION_MODULO
	DECLARACION_VECTOR
	DECLARACION
	ASIGNACION
	IF
	ELSE
	ELSEIF
	PRINT
	EXPRESION
	INSTRUCCION
	PRIMITIVO
	NULL
	ERROR
	ERROR_SEMANTICO
	ERROR_LEXICO
	ERROR_SINTACTICO
	VARIABLE
	INDEFINIDO
	VOID
	RETORNO
	EJECUTADO
	CASE
	CASE_EXPRESION
	DEFAULT
	BREAK_EXPRESION
	RETURN_EXPRESION
	PRINTF
	PRINT_PRIMITIVOS
	PRINT_ARRAY
	ERROR_SEMANTICO_NO
	LLAMADA_FUNCION
	PARAMETRO
	CAST
	SIMBOLO
	VALOR
	ERROR_NO_EXISTE
	ERROR_ACCESO_PRIVADO
	MUTABLE
	NOT_MUTABLE
	OCUPADO
	LIBRE
	VEC_NEW
	VEC_PUSH
	VEC_LEN
	VEC_CONTAINS
	VEC_CAPACITY
	VEC_INSERT
	VEC_REMOVE
	VEC_ACCESO
	VEC_FAC
	VEC_ELEMENTOS
	VEC_WITH_CAPACITY
	POW
	ARRAY_ELEMENTOS
	ARRAY_FAC
	DIMENSION_ARRAY
	DECLARACION_ARRAY
	ACCESO_ARRAY
	ATRIBUTO
	STRUCT_TEMPLATE
	TIPO
	ACCESO_STRUCT
	ASIGNACION_STRUCT
	FUNCION_MAIN
	TO_CHARS
	RANGE_EXPRESION
	RANGE_RANGO
	CLONE
	ARITMETICA
	RELACIONAL
	LOGICA
	HEAP
	STACK
)

type TipoRetornado struct {
	Tipo    TipoDato
	Valor   interface{}
	Fila    int
	Columna int
}

type O3D struct {
	Lt         string
	Lf         string
	Valor      TipoRetornado
	Codigo     string
	Referencia string
}

func (t TipoRetornado) GetTipo() (TipoDato, TipoDato) {
	return t.Tipo, t.Tipo
}
func (t TipoRetornado) GetFila() int {
	return t.Fila
}
func (t TipoRetornado) GetColumna() int {
	return t.Columna
}

/* Método solo para clonar un string*/
func (t TipoRetornado) Clonar(scope *Scope) interface{} {
	return CrearString3D(t)
}

func CrearString3D(valor TipoRetornado) TipoRetornado {

	obj := O3D{
		Lt:         "",
		Lf:         "",
		Valor:      valor,
		Codigo:     "",
		Referencia: Primitivo_To_String(valor.Valor, valor.Tipo),
	}

	//Verificar que sea un string o un str

	temp := GetTemp()
	cadenaAscii := obj.Referencia
	arrayAscii := strings.Split(cadenaAscii, ",")
	codigo := ""
	//Inicializar la cadena con el valor inicial del H guardado en el temporal
	codigo += "/***************AGREGANDO UN STRING/STR AL HEAP*/\n"
	codigo += temp + " = " + "H;\n"

	for _, valor := range arrayAscii {
		codigo += "heap[(int)H] = " + valor + "; //Letra\n"
		codigo += "H = H + 1;\n"
		GetH()
	}
	//Agregar caracter para saber que la cadena ha terminado
	codigo += "heap[(int)H] = 0;\n"
	codigo += "H = H + 1;\n"
	GetH()
	codigo += "/***********************************************/\n"

	obj.Codigo = codigo
	obj.Referencia = temp

	return TipoRetornado{
		Tipo:  PRIMITIVO,
		Valor: obj,
	}

}

func Primitivo_To_String(valor interface{}, tipo TipoDato) string {
	var salida string = ""
	switch tipo {
	case I64, USIZE:
		salida = strconv.Itoa(valor.(int))
	case F64:
		salida = fmt.Sprintf("%f", valor.(float64))
	case STR, STRING:
		primera := true
		tmp := valor.(string)
		runes := []rune(tmp)
		salida = ""
		for i := 0; i < len(runes); i++ {
			if primera {
				salida += strconv.Itoa(int(runes[i]))
				primera = false
			} else {
				salida += "," + strconv.Itoa(int(runes[i]))
			}

		}
	case CHAR:
		tmp := valor.(string)
		char := int(tmp[0])
		salida = strconv.Itoa(char)
	case BOOLEAN:
		if valor.(bool) {
			salida = "1"
		} else {
			salida = "0"
		}
	}
	return salida
}
