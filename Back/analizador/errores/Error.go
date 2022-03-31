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

func GenerarError(tipoError int, elemento1, elemento2 interface{}, scope *Ast.Scope) Ast.TipoRetornado {

	_, tipoI := elemento1.(Ast.Abstracto).GetTipo()
	_, tipoD := elemento2.(Ast.Abstracto).GetTipo()
	fila := elemento1.(Ast.Abstracto).GetFila()
	columna := elemento1.(Ast.Abstracto).GetColumna()

	switch tipoError {
	/*Error de tipos entre operaciones*/
	case 1:

		msg := "Semantic error, can't operate " + Ast.ValorTipoDato[tipoI] +
			" type with " + Ast.ValorTipoDato[tipoD] +
			" type. -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
		nError := NewError(fila, columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	/*Error en tipos de suma*/
	case 2:
		msg := "Semantic error, can't add " + Ast.ValorTipoDato[tipoI] +
			" type to " + Ast.ValorTipoDato[tipoD] +
			" type. -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
		nError := NewError(fila, columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}

	}

	return Ast.TipoRetornado{}
}
