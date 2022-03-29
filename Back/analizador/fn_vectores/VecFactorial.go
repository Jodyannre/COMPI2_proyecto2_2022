package fn_vectores

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"

	"github.com/colegno/arraylist"
)

type VecFactorial struct {
	Tipo       Ast.TipoDato
	Elementos  *arraylist.List
	TipoVector Ast.TipoDato
	Fila       int
	Columna    int
}

func NewVecFactorial(elementos *arraylist.List, fila, columna int) VecFactorial {
	//Crear el vector dependiendo de las banderas
	nV := VecFactorial{
		Tipo:       Ast.VEC_FAC,
		Fila:       fila,
		Columna:    columna,
		TipoVector: Ast.INDEFINIDO,
		Elementos:  elementos,
	}
	return nV
}

func (v VecFactorial) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	//Se crea como factorial
	//Crear la cantidad de elementos que se solicita
	//conseguir la cantidad de veces que se va a repetir el valor
	elementoVeces := v.Elementos.GetValue(1)
	_, tipoParticular := elementoVeces.(Ast.Abstracto).GetTipo()
	cantidad := elementoVeces.(Ast.Expresion).GetValue(scope)
	elemento := v.Elementos.GetValue(0).(Ast.Expresion).GetValue(scope)
	sizeVector := 0
	tipoDelVector := Ast.TipoRetornado{Valor: true, Tipo: Ast.INDEFINIDO}
	vacio := true

	if cantidad.Tipo == Ast.ERROR {
		return cantidad
	}
	if elemento.Tipo == Ast.ERROR {
		return elemento
	}
	if cantidad.Tipo != Ast.USIZE && tipoParticular != Ast.I64 {
		//Error, se esperaba un USIZE
		fila := v.Elementos.GetValue(1).(Ast.Abstracto).GetFila()
		columna := v.Elementos.GetValue(1).(Ast.Abstracto).GetColumna()
		msg := "Semantic error, expected usize, found. " + Ast.ValorTipoDato[cantidad.Tipo] +
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
	//Reiniciamos el valor del vector
	elementos := arraylist.New()
	for i := 0; i < cantidad.Valor.(int); i++ {
		nElemento := elemento
		if nElemento.Tipo == Ast.ERROR {
			return nElemento
		}
		if nElemento.Tipo == Ast.VECTOR {
			tipoDelVector = Ast.TipoRetornado{Tipo: Ast.VECTOR, Valor: nElemento.Valor.(expresiones.Vector).TipoVector}
		} else if tipoDelVector.Tipo == Ast.INDEFINIDO {
			tipoDelVector = Ast.TipoRetornado{Tipo: nElemento.Tipo, Valor: true}
			if tipoDelVector.Tipo == Ast.STRUCT {
				//Agregar el simbolo del struct
				tipoDelVector.Valor = nElemento.Valor.(Ast.Structs).GetPlantilla(scope)
			}
		}
		elementos.Add(nElemento)
		sizeVector++
		if vacio {
			vacio = false
		}
	}
	vector := expresiones.NewVector(elementos, tipoDelVector, sizeVector, sizeVector, vacio, v.Fila, v.Columna)
	vector.TipoVector = tipoDelVector
	return Ast.TipoRetornado{
		Tipo:  Ast.VECTOR,
		Valor: vector,
	}
}

func (v VecFactorial) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, v.Tipo
}

func (v VecFactorial) GetFila() int {
	return v.Fila
}
func (v VecFactorial) GetColumna() int {
	return v.Columna
}
