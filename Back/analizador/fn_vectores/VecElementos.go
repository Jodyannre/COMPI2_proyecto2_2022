package fn_vectores

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"

	"github.com/colegno/arraylist"
)

type VecElementos struct {
	Elementos *arraylist.List
	Tipo      Ast.TipoDato
	Fila      int
	Columna   int
}

func NewVecElementos(elementos *arraylist.List, fila, columna int) VecElementos {
	//Crear el vector dependiendo de las banderas
	nV := VecElementos{
		Tipo:      Ast.VEC_ELEMENTOS,
		Fila:      fila,
		Columna:   columna,
		Elementos: elementos,
	}
	return nV
}

func (v VecElementos) GetValue(scope *Ast.Scope) Ast.TipoRetornado {

	elementos := arraylist.New()
	var elemento interface{}
	tipoVector := Ast.TipoRetornado{Tipo: Ast.INDEFINIDO, Valor: true}
	valorElemento := Ast.TipoRetornado{}
	tipoAnterior := Ast.TipoRetornado{Tipo: Ast.INDEFINIDO, Valor: true}
	tipoDelVectorAnterior := Ast.TipoRetornado{Tipo: Ast.INDEFINIDO, Valor: true}
	tipoDelArrayAnterior := Ast.TipoRetornado{Tipo: Ast.INDEFINIDO, Valor: true}
	var vector expresiones.Vector
	var array expresiones.Array
	size := 0
	capacity := 0
	vacio := true

	for i := 0; i < v.Elementos.Len(); i++ {
		elemento = v.Elementos.GetValue(i)
		//Calcular el valor del elemento
		valorElemento = elemento.(Ast.Expresion).GetValue(scope)
		//Si hay error solo lo retorno
		if valorElemento.Tipo == Ast.ERROR {
			return valorElemento
		}
		//Calcular los tipos del elemento
		tipoGeneral, tipoParticular := elemento.(Ast.Abstracto).GetTipo()
		tipoParticular = expresiones.EsVector(tipoParticular)

		//Verificar si la variable proviene de un identificador y cambiar los tipos
		if tipoParticular == Ast.IDENTIFICADOR {
			tipoParticular = valorElemento.Tipo
		}

		if tipoGeneral != Ast.EXPRESION {
			//Error, no se puede guardar algo que no sea una expresión en el vector
			fila := elemento.(Ast.Abstracto).GetFila()
			columna := elemento.(Ast.Abstracto).GetColumna()
			msg := "Semantic error, can't store " + Ast.ValorTipoDato[tipoParticular] + " to a vector" +
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
		if tipoAnterior.Tipo != Ast.INDEFINIDO {
			if tipoAnterior.Tipo != tipoParticular {
				//Los tipos son diferentes, error
				fila := elemento.(Ast.Abstracto).GetFila()
				columna := elemento.(Ast.Abstracto).GetColumna()
				msg := "Semantic error, can't store " + Ast.ValorTipoDato[tipoParticular] + " value" +
					" in a Vec<" + Ast.ValorTipoDato[tipoAnterior.Tipo] + ">." +
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
		}

		//Si es un array o un vector, verificar que los tipos sean correctos
		if tipoParticular == Ast.VECTOR && tipoAnterior.Tipo != Ast.INDEFINIDO {
			vector = valorElemento.Valor.(expresiones.Vector)
			if tipoAnterior.Tipo == Ast.INDEFINIDO {
				tipoDelVectorAnterior.Tipo = GetTipoVector(vector)
			} else {
				if !expresiones.CompararTipos(tipoAnterior.Valor.(Ast.TipoRetornado), vector.TipoVector) {
					//Error no se pueden guardar 2 tipos de vectores diferentes
					fila := elemento.(Ast.Abstracto).GetFila()
					columna := elemento.(Ast.Abstracto).GetColumna()
					msg := "Semantic error, can't store VECTOR<" + Ast.ValorTipoDato[GetTipoVector(vector)] +
						"> to a VECTOR< VECTOR<" + Ast.ValorTipoDato[tipoDelVectorAnterior.Tipo] + "> >" +
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
			}
		}

		//Si es un array, verificar que los arrays sean del mismo tipo
		if tipoParticular == Ast.ARRAY && tipoAnterior.Tipo != Ast.INDEFINIDO {
			array = valorElemento.Valor.(expresiones.Array)
			if tipoDelArrayAnterior.Tipo == Ast.INDEFINIDO {
				tipoDelArrayAnterior.Tipo = array.TipoDelArray.Tipo
			} else {
				if tipoDelArrayAnterior.Tipo != array.TipoDelArray.Tipo {
					//Error no se pueden guardar 2 tipos de vectores diferentes
					fila := elemento.(Ast.Abstracto).GetFila()
					columna := elemento.(Ast.Abstracto).GetColumna()
					msg := "Semantic error, can't store ARRAY[" + Ast.ValorTipoDato[array.TipoDelArray.Tipo] +
						"] to a VECTOR<ARRAY[" + Ast.ValorTipoDato[tipoDelArrayAnterior.Tipo] + "]>" +
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
			}
		}
		if tipoParticular == Ast.STRUCT && tipoAnterior.Tipo != Ast.INDEFINIDO {
			if valorElemento.Valor.(Ast.Structs).GetPlantilla(scope) != tipoAnterior.Valor {
				fila := elemento.(Ast.Abstracto).GetFila()
				columna := elemento.(Ast.Abstracto).GetColumna()
				msg := "Semantic error, can't store" + valorElemento.Valor.(Ast.Structs).GetPlantilla(scope) +
					"to a VECTOR<" + tipoAnterior.Valor.(string) + ">" +
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

		}

		if tipoParticular == Ast.STRUCT {
			tipoAnterior.Tipo = tipoParticular
			tipoAnterior.Valor = valorElemento.Valor.(Ast.Structs).GetPlantilla(scope)
		} else if tipoParticular != Ast.IDENTIFICADOR && !EsFuncion(tipoParticular) && tipoParticular != Ast.VECTOR {
			tipoAnterior.Tipo = tipoParticular

		} else if tipoParticular == Ast.VECTOR {
			tipoAnterior = Ast.TipoRetornado{
				Tipo:  Ast.VECTOR,
				Valor: valorElemento.Valor.(expresiones.Vector).TipoVector,
			}

		} else {
			tipoAnterior.Tipo = valorElemento.Tipo
		}

		//Todo bien, entonces agregar el elemento a la lista del vector
		elementos.Add(valorElemento)
		size++
		//tipoAnterior.Tipo = valorElemento.Tipo
		if vacio {
			vacio = false
			tipoVector = tipoAnterior
		}
	}
	//Actualizar capacity
	capacity = size

	newVector := expresiones.NewVector(elementos, tipoVector, size, capacity, vacio, v.Fila, v.Columna)
	newVector.Tipo = Ast.VECTOR

	return Ast.TipoRetornado{
		Tipo:  Ast.VECTOR,
		Valor: newVector,
	}

}

func (v VecElementos) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, v.Tipo
}

func (v VecElementos) GetFila() int {
	return v.Fila
}
func (v VecElementos) GetColumna() int {
	return v.Columna
}

func GetTipoVector(vector expresiones.Vector) Ast.TipoDato {
	if vector.Tipo == Ast.VECTOR {
		return vector.Tipo
	}
	if vector.Tipo == Ast.ARRAY {
		return vector.Tipo
	}
	return vector.Tipo
}

func GetNivelesVector(vectorGuardado expresiones.Vector, vectorEntrante expresiones.Vector) bool {
	if vectorGuardado.Tipo == Ast.VECTOR && vectorEntrante.Tipo != Ast.VECTOR ||
		vectorGuardado.Tipo != Ast.VECTOR && vectorEntrante.Tipo == Ast.VECTOR {
		return false
	}
	if vectorGuardado.Tipo == Ast.VECTOR && vectorEntrante.Tipo == Ast.VECTOR {
		//Verificar si tiene elementos ambos
		if vectorGuardado.Valor.Len() > 0 && vectorEntrante.Valor.Len() > 0 {
			if vectorGuardado.TipoVector.Tipo == Ast.VECTOR && vectorEntrante.TipoVector.Tipo == Ast.VECTOR {
				return GetNivelesVector(vectorGuardado.Valor.GetValue(0).(Ast.TipoRetornado).Valor.(expresiones.Vector),
					vectorEntrante.Valor.GetValue(0).(Ast.TipoRetornado).Valor.(expresiones.Vector))
			}
			if vectorGuardado.TipoVector.Tipo != Ast.VECTOR && vectorEntrante.TipoVector.Tipo == Ast.VECTOR {
				return false
			}
			if vectorGuardado.TipoVector.Tipo == Ast.VECTOR && vectorEntrante.TipoVector.Tipo != Ast.VECTOR {
				return false
			}

		}
		return true
	}
	return true
}

func GetTipoArray(array expresiones.Array) Ast.TipoRetornado {
	return array.TipoDelArray

}

func EsFuncion(tipo interface{}) bool {
	validador := false

	switch tipo {
	case Ast.FUNCION, Ast.VEC_NEW,
		Ast.VEC_LEN, Ast.VEC_CONTAINS,
		Ast.VEC_CAPACITY, Ast.VEC_REMOVE, Ast.VEC_FAC,
		Ast.VEC_WITH_CAPACITY, Ast.VEC_ELEMENTOS:
		validador = true
	default:
		validador = false
	}

	return validador
}
