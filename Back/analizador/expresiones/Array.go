package expresiones

import (
	"Back/analizador/Ast"

	"github.com/colegno/arraylist"
)

type Array struct {
	Tipo          Ast.TipoDato
	TipoArray     Ast.TipoDato
	Elementos     *arraylist.List
	Fila          int
	Columna       int
	Mutable       bool
	TipoDelVector Ast.TipoDato
	TipoDelArray  Ast.TipoRetornado
	TipoDelStruct string
	Size          int //Tamaño de la dimensión
}

func NewArray(elementos *arraylist.List, TipoArray Ast.TipoDato, size, fila, columna int) Array {
	nA := Array{
		Tipo:          Ast.ARRAY,
		Elementos:     elementos,
		Fila:          fila,
		Columna:       columna,
		TipoArray:     TipoArray,
		TipoDelVector: Ast.INDEFINIDO,
		TipoDelArray:  Ast.TipoRetornado{Tipo: Ast.INDEFINIDO, Valor: true},
		TipoDelStruct: "INDEFINIDO",
		Size:          size,
	}
	return nA
}

func (a Array) GetValue(scope *Ast.Scope) Ast.TipoRetornado {

	return Ast.TipoRetornado{
		Tipo:  Ast.ARRAY,
		Valor: a,
	}

}

func (a Array) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, a.Tipo
}

func (a Array) GetFila() int {
	return a.Fila
}
func (a Array) GetColumna() int {
	return a.Columna
}

func (a Array) GetElement(index int) Ast.TipoRetornado {
	return Ast.TipoRetornado{}
}

/*
func GetTipoVector(vector instrucciones.Vector) Ast.TipoDato {
	if vector.TipoVector == Ast.VECTOR {
		return vector.TipoDelVector
	}
	if vector.TipoVector == Ast.ARRAY {
		return vector.TipoDelArray
	}
	return vector.TipoVector
}

func GetTipoArray(array Array) Ast.TipoDato {
	if array.TipoArray == Ast.VECTOR {
		return array.TipoDelVector
	}
	if array.TipoArray == Ast.ARRAY {
		return array.TipoDelArray
	}
	return array.TipoArray
}
*/

func (a Array) Clonar(scope *Ast.Scope) interface{} {
	var nElemento interface{}
	nElementos := arraylist.New()
	nV := Array{
		Fila:          a.Fila,
		Columna:       a.Columna,
		Size:          a.Size,
		Mutable:       a.Mutable,
		Tipo:          a.Tipo,
		TipoArray:     a.TipoArray,
		TipoDelArray:  a.TipoDelArray,
		TipoDelVector: a.TipoDelVector,
		TipoDelStruct: a.TipoDelStruct,
	}
	for i := 0; i < a.Elementos.Len(); i++ {
		elemento := a.Elementos.GetValue(i).(Ast.TipoRetornado)
		if EsVAS(elemento.Tipo) {
			preElemento := elemento.Valor.(Ast.Clones).Clonar(scope)
			_, tipoParticular := preElemento.(Ast.Abstracto).GetTipo()
			valor := Ast.TipoRetornado{Valor: preElemento}
			switch tipoParticular {
			case Ast.ARRAY:
				valor.Tipo = Ast.ARRAY
			case Ast.VECTOR:
				valor.Tipo = Ast.VECTOR
			case Ast.STRUCT:
				valor.Tipo = Ast.STRUCT
			}
			nElemento = valor
		} else {
			nElemento = elemento
		}
		nElementos.Add(nElemento)
	}
	nV.Elementos = nElementos
	return nV
}

func (v Array) GetMutable() bool {
	return v.Mutable
}
