package instrucciones

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/colegno/arraylist"
)

type Print struct {
	Expresiones Ast.Expresion
	Tipo        Ast.TipoDato
	Fila        int
	Columna     int
}

type PrintF struct {
	Expresiones *arraylist.List
	Cadena      string
	Tipo        Ast.TipoDato
	Fila        int
	Columna     int
}

func NewPrint(val Ast.Expresion, tipo Ast.TipoDato, fila, columna int) Print {
	nP := Print{
		Expresiones: val,
		Fila:        fila,
		Columna:     columna,
		Tipo:        tipo,
	}
	return nP
}

func NewPrintF(expresiones *arraylist.List, cadena string, tipo Ast.TipoDato, fila, columna int) PrintF {
	nP := PrintF{
		Expresiones: expresiones,
		Cadena:      cadena,
		Tipo:        tipo,
		Fila:        fila,
		Columna:     columna,
	}
	return nP
}

func (i Print) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, Ast.PRINT
}

func (i PrintF) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, Ast.PRINTF
}

func (p Print) Run(scope *Ast.Scope) interface{} {
	resultado_expresion := p.Expresiones.GetValue(scope)
	valor := ""
	//Verificar que no sea un identificador
	_, tipoParticular := p.Expresiones.(Ast.Abstracto).GetTipo()

	if tipoParticular == Ast.IDENTIFICADOR {
		//Error, con este tipo de print solo se puedfen imprimir literales
		msg := "Semantic error, a literal was expected, " + Ast.ValorTipoDato[resultado_expresion.Tipo] +
			" type was found." +
			" -- Line:" + strconv.Itoa(p.Fila) + " Column: " + strconv.Itoa(p.Columna)
		nError := errores.NewError(p.Fila, p.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}

	//Si resultado es error, que lo retorne
	if resultado_expresion.Tipo == Ast.ERROR {
		return resultado_expresion
	}
	if resultado_expresion.Tipo == Ast.STR {
		valor = resultado_expresion.Valor.(string)
	} else {
		//Error, no es un tipo que se pueda imprimir
		//O es una operación que dio como resultado null
		//No existe, generar un error semántico
		msg := "Semantic error, a literal was expected, " + Ast.ValorTipoDato[resultado_expresion.Tipo] +
			" type was found." +
			" -- Line:" + strconv.Itoa(p.Fila) + " Column: " + strconv.Itoa(p.Columna)
		nError := errores.NewError(p.Fila, p.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}
	//Actualizar consola del scope global directamente
	//scope.Consola += valor + "\n"
	scope.AgregarPrint(valor + "\n")
	return Ast.TipoRetornado{
		Tipo:  Ast.EJECUTADO,
		Valor: true,
	}
}

func (p PrintF) Run(scope *Ast.Scope) interface{} {
	//Formatos de los regex
	var salida string
	var valor interface{}
	regex, _ := regexp.Compile("{ *}|{:[\x3F]}")
	posiciones_regex := regex.FindAllStringIndex(p.Cadena, -1)
	encontrados := regex.MatchString(p.Cadena)    //Formato para encontrar los {} y {:?}
	elementos_string := regex.Split(p.Cadena, -1) //Cadena cortada por elementos
	//posciciones := regex.FindAllStringIndex(p.Cadena, -1) //Array de posiciones de los elementos {} encontrados
	retorno := Ast.TipoRetornado{}
	retorno.Tipo = Ast.STRING
	if !encontrados {
		//No se encontraron, error
		msg := "Semantic error, the number of expressions expected (" + strconv.Itoa(len(elementos_string)-1) + ")" +
			" is different within the print statement (" + strconv.Itoa(p.Expresiones.Len()) + ")." +
			" -- Line:" + strconv.Itoa(p.Fila) + " Column: " + strconv.Itoa(p.Columna)
		nError := errores.NewError(p.Fila, p.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}
	if len(elementos_string)-1 != p.Expresiones.Len() {
		//Error, la cantidad de expresiones es diferente de la que se esperaba
		msg := "Semantic error, the number of expressions expected (" + strconv.Itoa(len(elementos_string)-1) + ")" +
			"is different within the print statement (" + strconv.Itoa(p.Expresiones.Len()) + ")." +
			" -- Line:" + strconv.Itoa(p.Fila) + " Column: " + strconv.Itoa(p.Columna)
		nError := errores.NewError(p.Fila, p.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}

	}
	var cadena = ""
	var preCadena Ast.TipoRetornado

	for i := range elementos_string {
		//verificar los tipos

		if elementos_string[i] == "" && i < 1 {
			//En el primero agrego el primer elemento
			resultado := p.GetCompareValues(scope, i, posiciones_regex[i])
			if resultado.Tipo != Ast.ERROR {
				salida += resultado.Valor.(string)
			} else {
				return resultado
			}
		} else if elementos_string[i] == "" && i == len(elementos_string)-1 {
			//En el último no hago nada
		} else {

			if i >= p.Expresiones.Len() {
				salida += elementos_string[i]
			} else {
				resultado := p.GetCompareValues(scope, i, posiciones_regex[i])
				if resultado.Tipo != Ast.ERROR {
					valor = p.Expresiones.GetValue(i).(Ast.Expresion).GetValue(scope)
					preCadena = To_String(valor.(Ast.TipoRetornado), scope).(Ast.TipoRetornado)
					//valor = p.Expresiones.GetValue(i).(Ast.Expresion).GetValue(scope)
					if preCadena.Tipo == Ast.ERROR {
						//Crear el error y retornarlo
						msg := "Semantic error, a literal was expected," +
							Ast.ValorTipoDato[valor.(Ast.TipoRetornado).Tipo] + " was found" +
							" -- Line:" + strconv.Itoa(p.Fila) + " Column: " + strconv.Itoa(p.Columna)
						nError := errores.NewError(p.Fila, p.Columna, msg)
						nError.Tipo = Ast.ERROR_SEMANTICO
						nError.Ambito = scope.GetTipoScope()
						scope.Errores.Add(nError)
						scope.Consola += msg + "\n"
						return Ast.TipoRetornado{
							Tipo:  Ast.ERROR,
							Valor: nError,
						}
					} else {
						cadena = preCadena.Valor.(string)
					}
				} else {
					return resultado
				}
				salida += elementos_string[i] + cadena
			}
		}
	}
	//scope.Consola += salida + "\n"
	scope.AgregarPrint(salida + "\n")
	return Ast.TipoRetornado{
		Valor: true,
		Tipo:  Ast.EJECUTADO,
	}
}

func To_String(valor Ast.TipoRetornado, scope *Ast.Scope) interface{} {
	salida := ""
	preSalida := Ast.TipoRetornado{
		Tipo:  Ast.STRING,
		Valor: "",
	}
	switch valor.Tipo {
	case Ast.I64, Ast.USIZE:
		salida = strconv.Itoa(valor.Valor.(int))
	case Ast.F64:
		salida = fmt.Sprintf("%f", valor.Valor.(float64))
	case Ast.STR:
		salida = valor.Valor.(string)
	case Ast.CHAR:
		salida = valor.Valor.(string)
	case Ast.BOOLEAN:
		salida = strconv.FormatBool(valor.Valor.(bool))
	case Ast.STRING | Ast.STRING_OWNED:
		salida = valor.Valor.(string)
	case Ast.VECTOR:
		//De momento no tengo idea, pendiente
		//Recorrer todos sus elementos e irlos convirtiendo en string
		lista := valor.Valor.(expresiones.Vector).Valor
		var tipoAnterior Ast.TipoDato
		var elemento Ast.TipoRetornado
		salida += "[ "
		for i := 0; i < lista.Len(); i++ {
			if i != 0 && tipoAnterior != Ast.LIBRE {
				salida += ", "
			}
			elemento = lista.GetValue(i).(Ast.TipoRetornado)
			resultado := To_String(elemento, scope)
			tipoAnterior = elemento.Tipo
			if elemento.Tipo == Ast.STRING ||
				elemento.Tipo == Ast.STR ||
				elemento.Tipo == Ast.STRING_OWNED {
				salida += "\"" + resultado.(Ast.TipoRetornado).Valor.(string) + "\""
			} else if elemento.Tipo == Ast.CHAR {
				salida += "'" + resultado.(Ast.TipoRetornado).Valor.(string) + "'"
			} else {
				salida += resultado.(Ast.TipoRetornado).Valor.(string)
			}

		}
		if salida[len(salida)-1] == ',' {
			//Si hay coma al final, eliminarla
			salida = salida[:len(salida)-1]
		}
		salida += " ]"
	case Ast.LIBRE:
		//Espacios libres en un vector
		salida += ""
	case Ast.STRUCT:
		salida += valor.Valor.(Ast.Structs).GetPlantilla(scope)
	case Ast.ARRAY:
		//Recorrer todos sus elementos e irlos convirtiendo en string
		lista := valor.Valor.(expresiones.Array).Elementos
		var tipoAnterior Ast.TipoDato
		var elemento Ast.TipoRetornado
		salida += "[ "
		for i := 0; i < lista.Len(); i++ {
			if i != 0 && tipoAnterior != Ast.LIBRE {
				salida += ", "
			}
			elemento = lista.GetValue(i).(Ast.TipoRetornado)
			resultado := To_String(elemento, scope)
			tipoAnterior = elemento.Tipo
			if elemento.Tipo == Ast.STRING ||
				elemento.Tipo == Ast.STR ||
				elemento.Tipo == Ast.STRING_OWNED {
				salida += "\"" + resultado.(Ast.TipoRetornado).Valor.(string) + "\""
			} else if elemento.Tipo == Ast.CHAR {
				salida += "'" + resultado.(Ast.TipoRetornado).Valor.(string) + "'"
			} else {
				salida += resultado.(Ast.TipoRetornado).Valor.(string)
			}
		}
		if salida[len(salida)-1] == ',' {
			//Si hay coma al final, eliminarla
			salida = salida[:len(salida)-1]
		}
		salida += " ]"
	default:
		preSalida = Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: "",
		}
	}
	if preSalida.Tipo != Ast.ERROR {
		preSalida.Valor = salida
	}

	return preSalida
}

func TypeString(tipo Ast.TipoDato, cadena string) Ast.TipoDato {
	var tipoPrint Ast.TipoDato
	if tipo > 10 {
		return Ast.ERROR
	}
	if cadena == "{}" {
		tipoPrint = validacion_String[0][tipo]
	} else {
		tipoPrint = validacion_String[1][tipo]
	}

	return tipoPrint
}

var validacion_String = [2][11]Ast.TipoDato{
	{Ast.STRING, Ast.STRING, Ast.STRING, Ast.STRING, Ast.STRING, Ast.STRING, Ast.STRING, Ast.STRING, Ast.ERROR, Ast.ERROR, Ast.ERROR},
	{Ast.STRING, Ast.STRING, Ast.STRING, Ast.STRING, Ast.STRING, Ast.STRING, Ast.STRING, Ast.STRING, Ast.STRING, Ast.STRING, Ast.STRING},
	//I64-F64-String_owned-string-str-boolean-char-vector-array-struct
}

func (p PrintF) GetCompareValues(scope *Ast.Scope, i int, posiciones []int) Ast.TipoRetornado {
	salida := ""
	//En el primero agrego el primer elemento
	valor := p.Expresiones.GetValue(i).(Ast.Expresion).GetValue(scope)
	//Verificar que el tipo espera es el que se va a imprimir
	subString := p.Cadena[posiciones[0]:posiciones[1]]
	subString = strings.Replace(subString, " ", "", -1)
	//Verificar si el tipo es correcto dentro del string
	resultado := TypeString(valor.Tipo, subString)

	if resultado != Ast.ERROR {
		salida = ""
		//Convertir los valores a string

		switch valor.Tipo {
		case Ast.I64, Ast.USIZE:
			salida += strconv.Itoa(valor.Valor.(int))
		case Ast.F64:
			salida += strconv.FormatFloat(valor.Valor.(float64), 'f', -1, 64)
		case Ast.STRING:
			salida += valor.Valor.(string)
		case Ast.STR:
			salida += valor.Valor.(string)
		case Ast.STRING_OWNED:
			salida += valor.Valor.(string)
		case Ast.CHAR:
			salida += valor.Valor.(string)
		case Ast.BOOLEAN:
			salida += strconv.FormatBool(valor.Valor.(bool))
		case Ast.VECTOR, Ast.ARRAY:
			salida = To_String(valor, scope).(Ast.TipoRetornado).Valor.(string)
		case Ast.STRUCT:
			salida += valor.Valor.(Ast.Structs).GetPlantilla(scope)
		default:
		}

	} else {
		//Error, no se puede imprimir eso
		if valor.Tipo == Ast.NULL {
			msg := "Semantic error, can't print a NULL value" +
				" -- Line:" + strconv.Itoa(p.Fila) + " Column: " + strconv.Itoa(p.Columna)
			nError := errores.NewError(p.Fila, p.Columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			nError.Ambito = scope.GetTipoScope()
			scope.Errores.Add(nError)
			scope.Consola += msg + "\n"
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}
		}
		//Verificar que no es un error
		if valor.Tipo == Ast.ERROR {
			return valor
		}
		msg := "Semantic error, can't format " + Ast.ValorTipoDato[valor.Tipo] +
			" type with " + subString + "." +
			" -- Line:" + strconv.Itoa(p.Fila) + " Column: " + strconv.Itoa(p.Columna)
		nError := errores.NewError(p.Fila, p.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}
	return Ast.TipoRetornado{
		Valor: salida,
		Tipo:  Ast.STRING,
	}
}

func (op Print) GetFila() int {
	return op.Fila
}
func (op Print) GetColumna() int {
	return op.Columna
}

func (op PrintF) GetFila() int {
	return op.Fila
}
func (op PrintF) GetColumna() int {
	return op.Columna
}
