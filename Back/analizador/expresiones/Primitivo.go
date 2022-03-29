package expresiones

import (
	"Back/analizador/Ast"
)

type Primitivo struct {
	Tipo    Ast.TipoDato
	Valor   interface{}
	Fila    int
	Columna int
}

func (p Primitivo) GetValue(entorno *Ast.Scope) Ast.TipoRetornado {
	return Ast.TipoRetornado{
		Tipo:  p.Tipo,
		Valor: p.Valor,
	}
}

func (p Primitivo) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, p.Tipo
}

func NewPrimitivo(val interface{}, tipo Ast.TipoDato, fila, columna int) Primitivo {
	nuevo := Primitivo{Tipo: tipo, Valor: val, Fila: fila, Columna: columna}
	return nuevo
}

func (p Primitivo) GetFila() int {
	return p.Fila
}
func (p Primitivo) GetColumna() int {
	return p.Columna
}
