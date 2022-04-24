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
	var obj3d, obj3dValor, obj3dExp Ast.O3D
	var codigo3d string
	var preCodigo3d string
	var newExp expresiones.Primitivo
	var resExp Ast.TipoRetornado
	resultado_expresion := p.Expresiones.GetValue(scope)
	obj3dValor = resultado_expresion.Valor.(Ast.O3D)
	resultado_expresion = obj3dValor.Valor
	valor := ""
	var nuevaLinea string = "10"
	c := "c"
	porcentaje := "%%"
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

	/*Trabajar todo el código 3d aquí */
	/************************************************/
	//Conseguir el código 3d de estos elementos
	newExp = expresiones.NewPrimitivo(valor, Ast.STRING, 0, 0)
	resExp = newExp.GetValue(scope)
	obj3dExp = resExp.Valor.(Ast.O3D)
	codigo3d += obj3dExp.Codigo
	preCodigo3d, _ = GetC3DExpresion(obj3dExp)
	if obj3dExp.Lf != "" {
		codigo3d += obj3dExp.Lt + ":\n"
	}
	codigo3d += preCodigo3d
	//codigo3d += "printf (\"\\n\");\n"
	codigo3d += "printf(\"" + porcentaje + c + "\",(int)" + nuevaLinea + "); //Imprimir nueva linea\n"
	if obj3dExp.Lf != "" {
		codigo3d += obj3dExp.Lf + ":\n"
	}
	/************************************************/
	obj3d.Codigo = codigo3d
	obj3d.Valor = Ast.TipoRetornado{Tipo: Ast.PRINT, Valor: true}
	return Ast.TipoRetornado{
		Tipo:  Ast.PRINT,
		Valor: obj3d,
	}
}

func (p PrintF) Run(scope *Ast.Scope) interface{} {
	//Formatos de los regex
	var salida string
	var valor interface{}
	var obj3d Ast.O3D
	var obj3dValor Ast.O3D
	var codigo3d string
	var preCodigo3d string
	var salto string
	//var retornoImpresion string = "13"
	var nuevaLinea string = "10"
	c := "c"
	porcentaje := "%%"
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
				/************************************************/
				//Conseguir el código 3d de estos elementos
				obj3dExp := resultado.Valor.(Ast.O3D)
				cadena := To_String(obj3dExp.Valor, scope)
				codigo3d += obj3dExp.Codigo
				preCodigo3d, _ = GetC3DExpresion(obj3dExp)
				if obj3dExp.Lt != "" && obj3dExp.EsContains == "" {
					codigo3d += obj3dExp.Lt + ":\n"
					salto = Ast.GetLabel()
					codigo3d += "goto " + salto + ";\n"
				}
				codigo3d += preCodigo3d
				if obj3dExp.Lf != "" && obj3dExp.EsContains == "" {
					codigo3d += obj3dExp.Lf + ":\n"
					codigo3d += salto + ":\n"
				}
				/************************************************/
				salida += cadena.(Ast.TipoRetornado).Valor.(string)
			} else {
				return resultado
			}
		} else if elementos_string[i] == "" && i == len(elementos_string)-1 {
			//En el último no hago nada
		} else {

			if i >= p.Expresiones.Len() {
				salida += elementos_string[i]
				/************************************************/
				//Conseguir el código 3d de estos elementos
				newExp := expresiones.NewPrimitivo(elementos_string[i], Ast.STRING, 0, 0)
				resExp := newExp.GetValue(scope)
				obj3dExp := resExp.Valor.(Ast.O3D)
				codigo3d += obj3dExp.Codigo
				preCodigo3d, _ = GetC3DExpresion(obj3dExp)
				if obj3dExp.Lt != "" && obj3dExp.EsContains == "" {
					codigo3d += obj3dExp.Lt + ":\n"
					salto = Ast.GetLabel()
					codigo3d += "goto " + salto + ";\n"
				}
				codigo3d += preCodigo3d
				if obj3dExp.Lf != "" && obj3dExp.EsContains == "" {
					codigo3d += obj3dExp.Lf + ":\n"
					codigo3d += salto + ":\n"
				}
				/************************************************/
			} else {
				resultado := p.GetCompareValues(scope, i, posiciones_regex[i])
				if resultado.Tipo != Ast.ERROR {
					//valor = p.Expresiones.GetValue(i).(Ast.Expresion).GetValue(scope)
					obj3dValor = resultado.Valor.(Ast.O3D)
					valor = obj3dValor.Valor
					preCadena = To_String(obj3dValor.Valor, scope).(Ast.TipoRetornado)
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
				/************************************************/
				//Conseguir el código 3d de los elementos
				newExp := expresiones.NewPrimitivo(elementos_string[i], Ast.STRING, 0, 0)
				resExp := newExp.GetValue(scope)
				obj3dExp := resExp.Valor.(Ast.O3D)
				codigo3d += obj3dExp.Codigo
				if obj3dExp.Lt != "" && obj3dExp.EsContains == "" {
					codigo3d += obj3dExp.Lt + ":\n"
					salto = Ast.GetLabel()
					codigo3d += "goto " + salto + ";\n"
				}
				preCodigo3d, _ = GetC3DExpresion(obj3dExp)
				codigo3d += preCodigo3d
				if obj3dExp.Lf != "" && obj3dExp.EsContains == "" {
					codigo3d += obj3dExp.Lf + ":\n"
					codigo3d += salto + ":\n"
				}
				/************************************************/
				codigo3d += obj3dValor.Codigo
				preCodigo3d, _ = GetC3DExpresion(obj3dValor)
				if obj3dExp.Lt != "" && obj3dExp.EsContains == "" {
					codigo3d += obj3dValor.Lt + ":\n"
					salto = Ast.GetLabel()
					codigo3d += "goto " + salto + ";\n"
				}
				codigo3d += preCodigo3d
				if obj3dExp.Lt != "" && obj3dExp.EsContains == "" {
					codigo3d += obj3dValor.Lt + ":\n"
					codigo3d += salto + ":\n"
				}
			}
		}
	}
	//scope.Consola += salida + "\n"
	scope.AgregarPrint(salida + "\n")
	//codigo3d += "printf (\"\\n\");\n"
	codigo3d += "printf(\"" + porcentaje + c + "\",(int)" + nuevaLinea + "); //Imprimir nueva linea\n"
	//codigo3d += "printf(\"" + porcentaje + c + "\",(int)" + retornoImpresion + "); //Imprimir retorno carro\n"
	obj3d.Codigo = codigo3d
	obj3d.Valor = Ast.TipoRetornado{Tipo: Ast.STRING, Valor: "true"}
	return Ast.TipoRetornado{
		Valor: obj3d,
		Tipo:  Ast.PRINTF,
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
	var obj3d Ast.O3D
	salida := ""
	//En el primero agrego el primer elemento
	valor := p.Expresiones.GetValue(i).(Ast.Expresion).GetValue(scope)
	obj3d = valor.Valor.(Ast.O3D)
	valor = obj3d.Valor
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
			salida += To_String(valor, scope).(Ast.TipoRetornado).Valor.(string)
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
	/*
		obj3d.Valor = Ast.TipoRetornado{
			Valor: valor,
			Tipo:  valor.Tipo,
		}
	*/
	return Ast.TipoRetornado{
		Valor: obj3d,
		Tipo:  valor.Tipo,
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

func GetC3DExpresion(obj3d Ast.O3D) (string, string) {
	var codigo3d, resultado, siguientePos string
	var obj3dTemp Ast.O3D
	referencia := obj3d.Referencia
	valor := obj3d.Valor
	c := "c"
	d := "d"
	f := "f"
	p := "%%"
	corcheteIzq := "91"
	corcheteDer := "93"
	coma := "44"
	comillaDoble := "34"
	comillaSimple := "44"
	switch valor.Tipo {
	case Ast.STRING, Ast.STR:
		temp := Ast.GetTemp()
		temp2 := Ast.GetTemp()
		lt := Ast.GetLabel()
		lf := Ast.GetLabel()
		salto := Ast.GetLabel()
		codigo3d += "/******************************IMPRESION CADENA*/\n"
		codigo3d += salto + ":\n"
		codigo3d += temp + " = heap[(int)" + referencia + "]; //Get letra\n"
		codigo3d += "if (" + temp + "!=0) goto " + lt + ";\n"
		codigo3d += "goto " + lf + ";\n"
		codigo3d += lt + ":\n"
		codigo3d += "printf(\"" + p + c + "\",(int)" + temp + "); //Imprimir la letra\n"
		codigo3d += temp2 + " = " + referencia + "+ 1; //Actualizar posicion\n"
		codigo3d += referencia + " = " + temp2 + ";\n"
		codigo3d += "goto " + salto + ";\n"
		codigo3d += lf + ":\n"
		codigo3d += "/***********************************************/\n"
		siguientePos = referencia
	case Ast.I64, Ast.USIZE:
		codigo3d += "/*********************************IMPRESION I64*/\n"
		codigo3d += "printf(\"" + p + d + "\",(int)" + referencia + "); //Imprimir el numero\n"
		codigo3d += "/***********************************************/\n"
	case Ast.F64:
		codigo3d += "/*********************************IMPRESION F64*/\n"
		codigo3d += "printf(\"" + p + f + "\"," + referencia + "); //Imprimir el numero\n"
		codigo3d += "/***********************************************/\n"
	case Ast.BOOLEAN:
		codigo3d += "/********************************IMPRESION BOOL*/\n"
		if obj3d.Lt != "" {
			salto := Ast.GetLabel()
			if obj3d.RelacionalExp != "" {
				codigo3d += obj3d.RelacionalExp
			}
			codigo3d += obj3d.Lt + ":\n"
			codigo3d += "printf(\"" + p + c + "\",116); //Imprimir el booleano\n"
			codigo3d += "printf(\"" + p + c + "\",114); //Imprimir el booleano\n"
			codigo3d += "printf(\"" + p + c + "\",117); //Imprimir el booleano\n"
			codigo3d += "printf(\"" + p + c + "\",101); //Imprimir el booleano\n"
			codigo3d += "goto " + salto + ";\n"
			codigo3d += obj3d.Lf + ":\n"
			codigo3d += "printf(\"" + p + c + "\",102); //Imprimir el booleano\n"
			codigo3d += "printf(\"" + p + c + "\",97); //Imprimir el booleano\n"
			codigo3d += "printf(\"" + p + c + "\",108); //Imprimir el booleano\n"
			codigo3d += "printf(\"" + p + c + "\",115); //Imprimir el booleano\n"
			codigo3d += "printf(\"" + p + c + "\",101); //Imprimir el booleano\n"
			codigo3d += salto + ":\n"
			//codigo3d += "printf(\"" + p + d + "\",(int)" + referencia + "); //Imprimir el numero\n"
		} else {
			lt := Ast.GetLabel()
			lf := Ast.GetLabel()
			salto := Ast.GetLabel()
			if obj3d.RelacionalExp != "" {
				codigo3d += obj3d.RelacionalExp
			}
			codigo3d += "if (" + referencia + " == 1) goto " + lt + ";\n"
			codigo3d += "goto " + lf + ";\n"
			codigo3d += lt + ":\n"
			codigo3d += "printf(\"true\"); //Imprimir el booleano\n"
			codigo3d += "goto " + salto + ";\n"
			codigo3d += lf + ":\n"
			codigo3d += "printf(\"false\"); //Imprimir el booleano\n"
			codigo3d += salto + ":\n"
			//codigo3d += "printf(\"" + p + d + "\",(int)" + referencia + "); //Imprimir el numero\n"
		}

		codigo3d += "/***********************************************/\n"
	case Ast.CHAR:
		codigo3d += "/********************************IMPRESION CHAR*/\n"
		codigo3d += "printf(\"" + p + c + "\",(int)" + referencia + "); //Imprimir el numero\n"
		codigo3d += "/***********************************************/\n"
	case Ast.VECTOR:
		//De momento no tengo idea, pendiente
		//Recorrer todos sus elementos e irlos convirtiendo en string
		vector := obj3d.Valor.Valor.(expresiones.Vector)
		primeraPos := true
		lista := vector.Valor
		temp := Ast.GetTemp()
		posicionEnVector := Ast.GetTemp()
		contador := ""
		var tipoAnterior Ast.TipoDato
		var elemento Ast.TipoRetornado
		codigo3d += "/******************************IMPRESION VECTOR*/\n"
		codigo3d += temp + " = " + referencia + " + 1;//Subir una pos por el size del vec\n"
		contador = temp
		referencia = temp
		codigo3d += posicionEnVector + " = " + referencia + "; //Guardar posicion de vec actual \n"
		//codigo3d += "printf(\"[\");\n"
		codigo3d += "printf(\"" + p + c + "\",(int)" + corcheteIzq + "); //Imprimir corchete\n"
		for i := 0; i < lista.Len(); i++ {
			if i != 0 && tipoAnterior != Ast.LIBRE {
				//codigo3d += "printf(\",\");\n"
				codigo3d += "printf(\"" + p + c + "\",(int)" + coma + "); //Imprimir coma\n"
			}
			elemento = lista.GetValue(i).(Ast.TipoRetornado)
			if elemento.Tipo != Ast.STRING && elemento.Tipo != Ast.STR && elemento.Tipo != Ast.VECTOR {
				codigo3d += "/*********************GET ELEMENTO DESDE VECTOR*/\n"
				temp = Ast.GetTemp()
				codigo3d += temp + " = heap[(int)" + referencia + "];//Get valor\n"
				codigo3d += "/***********************************************/\n"
			} else if (elemento.Tipo == Ast.STRING || elemento.Tipo == Ast.STR) && primeraPos {
				codigo3d += "/*********************GET ELEMENTO DESDE VECTOR*/\n"
				temp = Ast.GetTemp()
				codigo3d += temp + " = heap[(int)" + referencia + "];//Get valor\n"
				codigo3d += "/***********************************************/\n"
				primeraPos = false
			} else if elemento.Tipo == Ast.VECTOR {
				codigo3d += "/*********************GET ELEMENTO DESDE VECTOR*/\n"
				temp2 := Ast.GetTemp()
				//temp3 := Ast.GetTemp()
				codigo3d += temp2 + " = heap[(int)" + contador + "];//Get direccion del otro vec\n"
				//codigo3d += temp3 + " = heap[(int)" + temp2 + "];//Get Valor del vec\n"
				temp = temp2
				codigo3d += "/***********************************************/\n"
			} else {
				temp = referencia
			}
			obj3dTemp.Valor = elemento
			obj3dTemp.Referencia = temp
			siguientePos = referencia
			resultado, referencia = GetC3DExpresion(obj3dTemp)
			tipoAnterior = elemento.Tipo
			if elemento.Tipo == Ast.STRING ||
				elemento.Tipo == Ast.STR ||
				elemento.Tipo == Ast.STRING_OWNED {
				//codigo3d += "printf(\"\\\"\");\n"
				codigo3d += "printf(\"" + p + c + "\",(int)" + comillaDoble + "); //Imprimir comillas\n"
				codigo3d += resultado
				//codigo3d += "printf(\"\\\"\");\n"
				codigo3d += "printf(\"" + p + c + "\",(int)" + comillaDoble + "); //Imprimir comillas\n"
				siguientePos = referencia
				codigo3d += posicionEnVector + " = " + posicionEnVector + " + 1; //sig pos vec\n"
				siguientePos = posicionEnVector
				primeraPos = true
			} else if elemento.Tipo == Ast.CHAR {
				//codigo3d += "printf(\"'\");\n"
				codigo3d += "printf(\"" + p + c + "\",(int)" + comillaSimple + "); //Imprimir comilla\n"
				codigo3d += resultado
				//codigo3d += "printf(\"'\");\n"
				codigo3d += "printf(\"" + p + c + "\",(int)" + comillaSimple + "); //Imprimir comilla\n"
			} else {
				codigo3d += resultado
			}

			/*Actualizar la referencia para la próxima posición*/
			codigo3d += "/****************************SIGUIENTE POSICION*/\n"
			if elemento.Tipo == Ast.VECTOR {

				temp := Ast.GetTemp()
				codigo3d += temp + " = " + contador + " +  1" + ";\n"
				contador = temp

			} else if elemento.Tipo == Ast.STRUCT {

			} else if elemento.Tipo == Ast.ARRAY {

			} else if elemento.Tipo == Ast.STR {
				temp := Ast.GetTemp()
				codigo3d += temp + " = " + siguientePos + ";\n"
				referencia = temp
			} else {
				temp := Ast.GetTemp()
				codigo3d += temp + " = " + siguientePos + " + 1;\n"
				referencia = temp
			}
			codigo3d += "/***********************************************/\n"

		}
		//codigo3d += "printf(\"]\");\n"
		codigo3d += "printf(\"" + p + c + "\",(int)" + corcheteDer + "); //Imprimir corchete\n"
		codigo3d += "/***********************************************/\n"
	case Ast.ARRAY:
		//De momento no tengo idea, pendiente
		//Recorrer todos sus elementos e irlos convirtiendo en string
		vector := obj3d.Valor.Valor.(expresiones.Array)
		lista := vector.Elementos
		temp := Ast.GetTemp()
		contador := ""
		var tipoAnterior Ast.TipoDato
		var elemento Ast.TipoRetornado
		codigo3d += "/*******************************IMPRESION ARRAY*/\n"
		codigo3d += temp + " = " + referencia + " + 1;//Subir una pos por el size del array\n"
		contador = temp
		referencia = temp
		//codigo3d += "printf(\"[\");\n"
		codigo3d += "printf(\"" + p + c + "\",(int)" + corcheteIzq + "); //Imprimir corchete\n"
		for i := 0; i < lista.Len(); i++ {
			if i != 0 && tipoAnterior != Ast.LIBRE {
				//codigo3d += "printf(\",\");\n"
				codigo3d += "printf(\"" + p + c + "\",(int)" + coma + "); //Imprimir coma\n"
			}
			elemento = lista.GetValue(i).(Ast.TipoRetornado)
			if elemento.Tipo != Ast.STRING && elemento.Tipo != Ast.STR && elemento.Tipo != Ast.ARRAY {
				codigo3d += "/**********************GET ELEMENTO DESDE ARRAY*/\n"
				temp = Ast.GetTemp()
				codigo3d += temp + " = heap[(int)" + referencia + "];//Get valor\n"
				codigo3d += "/***********************************************/\n"
			} else if (elemento.Tipo == Ast.STRING || elemento.Tipo == Ast.STR) && i == 0 {
				codigo3d += "/**********************GET ELEMENTO DESDE ARRAY*/\n"
				temp = Ast.GetTemp()
				codigo3d += temp + " = heap[(int)" + referencia + "];//Get valor\n"
				codigo3d += "/***********************************************/\n"
			} else if elemento.Tipo == Ast.ARRAY {
				codigo3d += "/**********************GET ELEMENTO DESDE ARRAY*/\n"
				temp2 := Ast.GetTemp()
				//temp3 := Ast.GetTemp()
				codigo3d += temp2 + " = heap[(int)" + contador + "];//Get direccion del otro array\n"
				//codigo3d += temp3 + " = heap[(int)" + temp2 + "];//Get Valor del vec\n"
				temp = temp2
				codigo3d += "/***********************************************/\n"
			} else {
				temp = referencia
			}
			obj3dTemp.Valor = elemento
			obj3dTemp.Referencia = temp
			siguientePos = referencia
			resultado, referencia = GetC3DExpresion(obj3dTemp)
			tipoAnterior = elemento.Tipo
			if elemento.Tipo == Ast.STRING ||
				elemento.Tipo == Ast.STR ||
				elemento.Tipo == Ast.STRING_OWNED {
				//codigo3d += "printf(\"\\\"\");\n"
				codigo3d += "printf(\"" + p + c + "\",(int)" + comillaDoble + "); //Imprimir comillas\n"
				codigo3d += resultado
				//codigo3d += "printf(\"\\\"\");\n"
				codigo3d += "printf(\"" + p + c + "\",(int)" + comillaDoble + "); //Imprimir comillas\n"
				siguientePos = referencia
			} else if elemento.Tipo == Ast.CHAR {
				codigo3d += "printf(\"" + p + c + "\",(int)" + comillaSimple + "); //Imprimir comilla\n"
				//codigo3d += "printf(\"'\");\n"
				codigo3d += resultado
				codigo3d += "printf(\"" + p + c + "\",(int)" + comillaSimple + "); //Imprimir comilla\n"
				//codigo3d += "printf(\"'\");\n"
			} else {
				codigo3d += resultado
			}

			/*Actualizar la referencia para la próxima posición*/
			codigo3d += "/****************************SIGUIENTE POSICION*/\n"
			if elemento.Tipo == Ast.VECTOR {

			} else if elemento.Tipo == Ast.STRUCT {

			} else if elemento.Tipo == Ast.ARRAY {
				temp := Ast.GetTemp()
				codigo3d += temp + " = " + contador + " +  1" + ";\n"
				contador = temp

			} else {
				temp := Ast.GetTemp()
				codigo3d += temp + " = " + siguientePos + " + 1;\n"
				referencia = temp
			}
			codigo3d += "/***********************************************/\n"

		}
		//codigo3d += "printf(\"]\");\n"
		codigo3d += "printf(\"" + p + c + "\",(int)" + corcheteDer + "); //Imprimir corchete\n"
		codigo3d += "/***********************************************/\n"
	}
	return codigo3d, siguientePos
}
