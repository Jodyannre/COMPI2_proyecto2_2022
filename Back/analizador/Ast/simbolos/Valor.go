package simbolos

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"strconv"
)

type Valor struct {
	Valor      Ast.Expresion
	Tipo       Ast.TipoDato
	Fila       int
	Columna    int
	Referencia bool
	Mutable    bool
}

func NewValor(valor Ast.Expresion, tipo Ast.TipoDato, referencia bool, mutable bool, fila, columna int) Valor {
	nV := Valor{
		Tipo:       tipo,
		Valor:      valor,
		Referencia: referencia,
		Mutable:    mutable,
		Fila:       fila,
		Columna:    columna,
	}
	return nV
}

func (v Valor) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	//Verificar que los valores de referencia sean correctos
	valor := v.Valor.GetValue(scope)

	if !EsPosibleReferencia(valor.Tipo) && v.Referencia {
		//Error, ese tipo de valor no se puede enviar como referencia
		fila := v.Fila
		columna := v.Columna
		msg := "Semantic error, " + Ast.ValorTipoDato[valor.Tipo] + " can't be send as a reference." +
			" -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
		nError := errores.NewError(fila, columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}
	return valor
}
func (p Valor) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, p.Tipo
}

func (p Valor) GetFila() int {
	return p.Fila
}
func (p Valor) GetColumna() int {
	return p.Columna
}
