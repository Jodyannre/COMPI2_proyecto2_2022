package fn_vectores

import (
	"Back/analizador/Ast"
	"Back/analizador/expresiones"

	"github.com/colegno/arraylist"
)

type VecNew struct {
	Tipo    Ast.TipoDato
	Fila    int
	Columna int
}

func NewVecNew(fila, columna int) VecNew {
	//Crear el vector dependiendo de las banderas
	nV := VecNew{
		Tipo:    Ast.VEC_NEW,
		Fila:    fila,
		Columna: columna,
	}
	return nV
}

func (w VecNew) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	elementos := arraylist.New()
	vector := expresiones.NewVector(elementos, Ast.TipoRetornado{Tipo: Ast.INDEFINIDO, Valor: true}, 0, 0, true, w.Fila, w.Columna)
	return Ast.TipoRetornado{
		Tipo:  Ast.VECTOR,
		Valor: vector,
	}
}

func (v VecNew) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, v.Tipo
}

func (v VecNew) GetFila() int {
	return v.Fila
}
func (v VecNew) GetColumna() int {
	return v.Columna
}
