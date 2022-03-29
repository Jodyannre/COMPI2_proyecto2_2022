package fn_vectores

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"

	"github.com/colegno/arraylist"
)

type VecWithCapacity struct {
	Tipo     Ast.TipoDato
	Capacity interface{}
	Fila     int
	Columna  int
}

func NewVecWithCapacity(capacity interface{}, fila, columna int) VecWithCapacity {
	//Crear el vector dependiendo de las banderas
	nV := VecWithCapacity{
		Tipo:     Ast.VEC_NEW,
		Fila:     fila,
		Columna:  columna,
		Capacity: capacity,
	}
	return nV
}

func (v VecWithCapacity) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	elementos := arraylist.New()
	capacity := 0
	size := 0
	vacio := true

	//Verificar que capacity sea un USIZE

	//Es un vector vacio, pero tiene que tener el tipo, entonces solo crear la lista y retornarlo

	capacidad := v.Capacity.(Ast.Expresion).GetValue(scope)
	_, tipoParticular := v.Capacity.(Ast.Abstracto).GetTipo()
	//verificar posible error
	if capacidad.Tipo == Ast.ERROR {
		return capacidad
	}
	//Iniciado con With Capacity
	if (tipoParticular == Ast.USIZE || capacidad.Tipo == Ast.USIZE) ||
		tipoParticular == Ast.I64 {
		capacity = capacidad.Valor.(int)

	} else {
		//ERROR NO ES UN USIZE
		fila := v.Capacity.(Ast.Abstracto).GetFila()
		columna := v.Capacity.(Ast.Abstracto).GetColumna()
		msg := "Semantic error, expected usize, found " + Ast.ValorTipoDato[capacidad.Tipo] +
			". -- Line: " + strconv.Itoa(fila) +
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

	vector := expresiones.NewVector(elementos, Ast.TipoRetornado{Tipo: Ast.INDEFINIDO, Valor: true}, size, capacity, vacio, v.Fila, v.Columna)
	return Ast.TipoRetornado{
		Tipo:  Ast.VECTOR,
		Valor: vector,
	}
}

func (v VecWithCapacity) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, v.Tipo
}

func (v VecWithCapacity) GetFila() int {
	return v.Fila
}
func (v VecWithCapacity) GetColumna() int {
	return v.Columna
}
