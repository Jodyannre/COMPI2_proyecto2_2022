package expresiones

import (
	"Back/analizador/Ast"
	//"Back/analizador/instrucciones"

	"github.com/colegno/arraylist"
)

type Vector struct {
	Tipo       Ast.TipoDato
	Valor      *arraylist.List
	TipoVector Ast.TipoRetornado
	Fila       int
	Columna    int
	Mutable    bool
	Vacio      bool
	Size       int
	Capacity   int
}

func NewVector(valor *arraylist.List, tipoVector Ast.TipoRetornado, size, capacity int, vacio bool, fila, columna int) Vector {
	//Crear el vector dependiendo de las banderas
	nV := Vector{
		Tipo:       Ast.VECTOR,
		Fila:       fila,
		Columna:    columna,
		Valor:      valor,
		TipoVector: tipoVector,
		Mutable:    false,
		Size:       size,
		Capacity:   capacity,
		Vacio:      vacio,
	}
	return nV
}

func (v Vector) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	//Crear los valores del vector
	return Ast.TipoRetornado{
		Valor: v,
		Tipo:  Ast.VECTOR,
	}
}

func (v Vector) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, v.Tipo
}

func (v Vector) GetFila() int {
	return v.Fila
}
func (v Vector) GetColumna() int {
	return v.Columna
}

func (v Vector) GetTipoVector() Ast.TipoRetornado {
	return v.TipoVector
}

func (v Vector) GetSize() int {
	return v.Size
}

func (v Vector) GetMutable() bool {
	return v.Mutable
}

func (v Vector) Clonar(scope *Ast.Scope) interface{} {
	var nElemento interface{}
	nLista := arraylist.New()
	nV := Vector{
		Fila:       v.Fila,
		Columna:    v.Columna,
		Capacity:   v.Capacity,
		Size:       v.Size,
		Vacio:      v.Vacio,
		Mutable:    v.Mutable,
		Tipo:       v.Tipo,
		TipoVector: v.TipoVector,
	}
	for i := 0; i < v.Valor.Len(); i++ {
		elemento := v.Valor.GetValue(i).(Ast.TipoRetornado)
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
		nLista.Add(nElemento)
	}
	nV.Valor = nLista
	return nV
}

func (v Vector) CalcularCapacity(size int, capacity int) int {
	if size == 1 && capacity == 0 {
		return 4
	}
	if size == 0 && capacity == 0 {
		return 0
	}
	if capacity <= size {
		if capacity == 0 {
			return v.CalcularCapacity(size, capacity+4)
		}
		return v.CalcularCapacity(size, capacity*2)
	} else {
		return capacity
	}
}
