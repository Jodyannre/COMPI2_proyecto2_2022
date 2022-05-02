package errores

import (
	"Back/analizador/Ast"
	"strconv"
	"time"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type CustomSyntaxError struct {
	Fila    int
	Columna int
	Msg     string
	Tipo    Ast.TipoDato
	Ambito  string
	Fecha   string
}

type CustomError struct {
	Fila    int
	Columna int
	Msg     string
	Tipo    Ast.TipoDato
	Ambito  string
	Fecha   string
}

func NewError(fila int, columna int, msg string) CustomSyntaxError {
	//Generar la fecha y la hora
	dt := time.Now()
	fecha := dt.Format("01-02-2006 15:04:05")
	return CustomSyntaxError{
		Fila:    fila,
		Columna: columna,
		Msg:     msg,
		Fecha:   fecha,
		Ambito:  "Local",
	}
}

type CustomErrorListener struct {
	*antlr.DefaultErrorListener
	Errores []CustomSyntaxError
}

func (c *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	c.Errores = append(c.Errores, CustomSyntaxError{
		Fila:    line,
		Columna: column,
		Msg:     msg,
	})
}

func (op CustomSyntaxError) GetFila() int {
	return op.Fila
}
func (op CustomSyntaxError) GetColumna() int {
	return op.Columna
}

func (e CustomSyntaxError) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, e.Tipo
}

func (op CustomSyntaxError) GetAmbito() string {
	return op.Ambito
}
func (op CustomSyntaxError) GetDescripcion() string {
	return op.Msg
}
func (op CustomSyntaxError) GetFecha() string {
	return op.Fecha
}

func GenerarError(tipoError int, elemento1, elemento2 interface{}, operador string,
	tipoIString, tipoDString string, scope *Ast.Scope) Ast.TipoRetornado {
	var obj3d Ast.O3D
	_, tipoI := elemento1.(Ast.Abstracto).GetTipo()
	_, tipoD := elemento2.(Ast.Abstracto).GetTipo()
	fila := elemento1.(Ast.Abstracto).GetFila()
	columna := elemento1.(Ast.Abstracto).GetColumna()
	var msg = ""

	switch tipoError {
	/*Errores en operaciones*/
	case 1:
		/*Error de tipos entre operaciones*/
		msg = "Semantic error, can't operate " + Ast.ValorTipoDato[tipoI] +
			" type with " + Ast.ValorTipoDato[tipoD] +
			" type. -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
		/*Error en tipos de suma*/
	case 2:
		msg = "Semantic error, can't add " + Ast.ValorTipoDato[tipoI] +
			" type to " + Ast.ValorTipoDato[tipoD] +
			" type. -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
	case 3:
		/*Error unario con usize*/
		msg = "Semantic error, can't apply unary operator `-` to type `usize`." +
			" -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
	case 4:
		/*Error de unario con un no booleano*/
		msg = "Semantic error, can't operate (!) with a " + Ast.ValorTipoDato[tipoI] +
			" type. -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
	case 5:
		/*Error overflow negativo en resta de usize*/
		msg = "Semantic error, attempt to subtract with overflow." +
			" -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
	case 6:
		/*Error en tipos de resta */
		msg = "Semantic error, can't subtract " + Ast.ValorTipoDato[tipoI] +
			" type to " + Ast.ValorTipoDato[tipoD] +
			" type. -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
	case 7:
		/*Error en tipos de multiplicación*/
		msg = "Semantic error, can't multiply " + Ast.ValorTipoDato[tipoI] +
			" type to " + Ast.ValorTipoDato[tipoD] +
			" type. -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
	case 8:
		/*Error en tipos de división*/
		msg = "Semantic error, can't divide " + Ast.ValorTipoDato[tipoI] +
			" type by " + Ast.ValorTipoDato[tipoD] +
			" type. -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
	case 9:
		/*Error de tipos en operación lógica*/
		msg = "Semantic error, can't logically operate " + Ast.ValorTipoDato[tipoI] +
			" type with " + Ast.ValorTipoDato[tipoD] +
			" type. -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
	case 10:
		msg = "Semantic error, can't compare a " + Ast.ValorTipoDato[tipoI] +
			" using " + operador +
			" type. -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
	case 11:
		msg = "Semantic error, can't compare " + Ast.ValorTipoDato[tipoI] +
			" with " + Ast.ValorTipoDato[tipoD] +
			" type. -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)

	/* Errores en declaraciones*/
	case 12:
		msg = "Semantic error, the element \"" + operador + "\" already exist in this scope." +
			" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
	case 13:
		msg = "Semantic error, can't initialize a " + tipoIString +
			" with " + tipoDString + " value." +
			" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
	case 14:
		msg = "Semantic error, type error." +
			" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)

		/*Errores en asignaciones*/

	case 15:
		msg = "Semantic error, the element \"" + operador + "\" doesn't exist in any scope." +
			" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
		/*Error en identificadores*/
	case 16:
		msg = "Semantic error, can't modify a non-mutable " + tipoIString +
			" type. -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
	case 17:
		msg = "Semantic error, can't assign " + tipoIString +
			" type to ARRAY[" + tipoDString + "]" +
			" type. -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
	case 18:
		msg = "Semantic error, can't assign Vector<" + tipoIString + ">" +
			" to Vector<" + tipoDString + ">" +
			" type. -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
	case 19:
		msg = "Semantic error, can't store " + tipoIString +
			" to an ARRAY[" + tipoDString + "]" +
			". -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
	case 20:
		msg = "Semantic error, \"" + operador + "\" variable doesn't not exist." +
			" -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
		/*Error de mutabilidad en asignación*/
	case 21:
		msg = "Semantic error, can't modify a non-mutable " + operador +
			" type. -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
	case 22:
		msg = "Semantic error, can't assign Vec<" + tipoIString + ">" +
			" to Vec<" + tipoDString + ">" +
			" type. -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
	case 23:
		msg = "Semantic error, the element \"" + operador + "\" doesn't exist in any scope." +
			" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
	case 24:
		msg = "Semantic error, can't assign " + tipoIString +
			" type to " + tipoDString +
			" type. -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
	case 25:
		msg = "Semantic error, ARRAY dimensions don't match." +
			" -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
	case 26:
		msg = "Semantic error, can't assign ARRAY[" + tipoIString + "]" +
			" to ARRAY[" + tipoDString + "]" +
			" type. -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
	case 27:
		msg = "Semantic error, can't store Vec<" + tipoIString +
			"> to a Vec<" + tipoDString + ">" +
			". -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
		/*Errores en transferencia en lugares incorrectos */
	case 30:
		msg = "Semantic error, cannot break outside of a loop." +
			" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
	case 31:
		msg = "Semantic error, MAIN method cannot return a value." +
			" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
		/*Continuacion de errores en asignaciones*/
	case 32:
		msg = "Semantic error, can't assign " + tipoIString +
			" type to Vector<" + tipoDString + ">" +
			" type. -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
	case 33:
		msg = "Semantic error, expected USIZE, found. " + tipoIString +
			". -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
	case 34:
		msg = "Semantic error, expected VECTOR, found " + tipoIString +
			". -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
	case 35:
		msg = "Semantic error, can't store \"" + tipoIString +
			"\" to a VECTOR<" + tipoDString + ">" +
			". -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
	case 36:
		msg = "Semantic error, expected IDENTIFICADOR, found " + tipoIString +
			". -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
		/*Errores en declaraciones simples*/
	case 37:
		msg = "Semantic error, can't initialize a Vector with " + tipoIString + " type" +
			" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
	case 38:
		msg = "Semantic error, can't initialize a Vector<" + tipoIString + "> with Vector<" +
			tipoDString + "> type" +
			" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
	case 39:
		msg = "Semantic error, can't initialize an ARRAY with " +
			tipoIString + " value." +
			" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
	case 40:
		msg = "Semantic error, can't initialize an ARRAY[" + tipoIString + "] with ARRAY[" +
			tipoDString + "> type" +
			" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
	case 41:
		msg = "Semantic error, expected  ARRAY[" + tipoIString + "] " +
			" found ARRAY[" + tipoDString + "]." +
			" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
	}

	nError := NewError(fila, columna, msg)
	nError.Tipo = Ast.ERROR_SEMANTICO
	nError.Ambito = scope.GetTipoScope()
	scope.Errores.Add(nError)
	scope.Consola += msg + "\n"

	obj3d.Valor = Ast.TipoRetornado{
		Tipo:  Ast.ERROR,
		Valor: nError,
	}

	return Ast.TipoRetornado{
		Tipo:  Ast.ERROR,
		Valor: obj3d,
	}
}
