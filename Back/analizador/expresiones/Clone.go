package expresiones

import (
	"Back/analizador/Ast"
)

type Clone struct {
	Expresion interface{}
	Tipo      Ast.TipoDato
	Fila      int
	Columna   int
}

func NewClone(expresion interface{}, fila, columna int) Clone {
	nC := Clone{
		Expresion: expresion,
		Tipo:      Ast.CLONE,
		Fila:      fila,
		Columna:   columna,
	}
	return nC
}

func (c Clone) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	//El nuevo valor
	var nValor interface{}
	//Conseguir el valor
	valor := c.Expresion.(Ast.Expresion).GetValue(scope)

	if valor.Tipo == Ast.ERROR {
		return valor
	}
	switch valor.Tipo {
	case Ast.STRUCT, Ast.VECTOR, Ast.ARRAY:
		nValor = valor.Valor.(Ast.Clones).Clonar(scope)
	default:
		nValor = valor.Valor
	}

	return Ast.TipoRetornado{
		Tipo:  valor.Tipo,
		Valor: nValor,
	}
}

func (c Clone) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, c.Tipo
}

func (c Clone) GetFila() int {
	return c.Fila
}
func (c Clone) GetColumna() int {
	return c.Columna
}
