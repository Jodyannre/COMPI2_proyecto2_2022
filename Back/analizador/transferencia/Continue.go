package transferencia

import "Back/analizador/Ast"

type Continue struct {
	Tipo    Ast.TipoDato
	Fila    int
	Columna int
}

func (c Continue) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, c.Tipo
}

func NewContinue(fila, columna int) Continue {
	nC := Continue{
		Tipo:    Ast.CONTINUE,
		Fila:    fila,
		Columna: columna,
	}
	return nC
}

func (c Continue) Run(scope *Ast.Scope) interface{} {
	return Ast.TipoRetornado{
		Valor: Ast.TipoRetornado{
			Tipo:  Ast.CONTINUE,
			Valor: c,
		},
		Tipo: Ast.CONTINUE,
	}
}

func (op Continue) GetFila() int {
	return op.Fila
}
func (op Continue) GetColumna() int {
	return op.Columna
}
