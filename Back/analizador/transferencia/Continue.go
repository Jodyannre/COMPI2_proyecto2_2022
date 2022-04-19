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
	/************************VARIABLES 3D**************************************/
	var obj3d Ast.O3D
	var salto string = Ast.GetLabel()
	/**************************************************************************/
	salto = Ast.GetLabel()
	obj3d.Valor = Ast.TipoRetornado{
		Tipo:  Ast.CONTINUE,
		Valor: c,
	}
	obj3d.SaltoTranferencia = salto
	obj3d.SaltoContinue = salto
	return Ast.TipoRetornado{
		Valor: obj3d,
		Tipo:  Ast.CONTINUE,
	}
}

func (op Continue) GetFila() int {
	return op.Fila
}
func (op Continue) GetColumna() int {
	return op.Columna
}
