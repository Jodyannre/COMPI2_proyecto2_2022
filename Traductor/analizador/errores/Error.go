package errores

import (
	"Traductor/analizador/ast"
	"time"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type CustomSyntaxError struct {
	Fila    int
	Columna int
	Msg     string
	Tipo    ast.TipoDato
	Ambito  string
	Fecha   string
}

type CustomError struct {
	Fila    int
	Columna int
	Msg     string
	Tipo    ast.TipoDato
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

func (e CustomSyntaxError) GetTipo() (ast.TipoDato, ast.TipoDato) {
	return ast.INSTRUCCION, e.Tipo
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
