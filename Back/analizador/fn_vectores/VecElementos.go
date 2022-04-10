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
	/********************VARIABLES 3D*******************************/
	var obj3d, obj3dValor Ast.O3D
	var preCodigo3d, codigo3d, referencia, referenciaRetorno string
	var primeraPos bool = true
	var vectores = arraylist.New()
	var estructuraBase string
	/***************************************************************/
	size := 0
	capacity := 0
	vacio := true
	codigo3d += "/****************************CREACION DE VECTOR*/\n"

	for i := 0; i < v.Elementos.Len(); i++ {
		elemento = v.Elementos.GetValue(i)
		//Calcular el valor del elemento
		valorElemento = elemento.(Ast.Expresion).GetValue(scope)
		obj3dValor = valorElemento.Valor.(Ast.O3D)
		valorElemento = obj3dValor.Valor
		/*Actualizar codigo 3d que viene de la expresión*/
		codigo3d += obj3dValor.Codigo
		referencia = obj3dValor.Referencia
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
		/*Get el codigo3d de agregar ese elemento al heap*/
		if tipoParticular == Ast.VECTOR || tipoParticular == Ast.STRING || tipoParticular == Ast.STR {
			//referencia, preCodigo3d = GetCod3dElemento(referencia, primeraPos, true)
			if primeraPos {
				estructuraBase = referencia
			}
		} else {
			referencia, preCodigo3d = GetCod3dElemento(referencia, primeraPos, false)
		}
		if tipoParticular == Ast.VECTOR || tipoParticular == Ast.STRING || tipoParticular == Ast.STR {
			vectores.Add(referencia)
		}

		codigo3d += preCodigo3d
		//tipoAnterior.Tipo = valorElemento.Tipo
		if vacio {
			primeraPos = false
			vacio = false
			tipoVector = tipoAnterior
			referenciaRetorno = referencia
		}
	}

	if tipoAnterior.Tipo == Ast.VECTOR || tipoAnterior.Tipo == Ast.STRING || tipoAnterior.Tipo == Ast.STR {
		//Agregar los elementos al vector
		//Crear el vector padre
		ref, cod := GetCod3dElemento(estructuraBase, true, false)
		codigo3d += cod
		/*Agregar los elementos al vector*/
		for i := 0; i < vectores.Len(); i++ {
			if i != 0 {
				elemento := vectores.GetValue(i).(string)
				_, cod := GetCod3dElemento(elemento, false, false)
				codigo3d += cod
			}
		}
		referenciaRetorno = ref
	}

	//Actualizar capacity
	capacity = size

	newVector := expresiones.NewVector(elementos, tipoVector, size, capacity, vacio, v.Fila, v.Columna)
	newVector.Tipo = Ast.VECTOR
	codigo3d += "/********************GUARDAR EL SIZE DEL VECTOR*/\n"
	codigo3d += "heap[(int)" + referenciaRetorno + "] = " + strconv.Itoa(newVector.Size) + ";\n"
	codigo3d += "/***********************************************/\n"
	codigo3d += "/***********************************************/\n"
	/*Actualizar datos del obj3d a retornar*/
	obj3d.Codigo = codigo3d
	/*Agregar el tamaño al vector*/
	obj3d.Referencia = referenciaRetorno

	obj3d.Valor = Ast.TipoRetornado{
		Tipo:  Ast.VECTOR,
		Valor: newVector,
	}
	return Ast.TipoRetornado{
		Tipo:  Ast.VECTOR,
		Valor: obj3d,
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

func GetCod3dElemento(referencia string, primeraPos bool, estructura bool) (string, string) {
	codigo3d := ""
	ref := ""
	if primeraPos {
		codigo3d += "/*****************************INICIO DEL VECTOR*/\n"
		temp := Ast.GetTemp()
		codigo3d += temp + " = H; //Guardar la referencia\n"
		codigo3d += "H = H + 1; //La primera posicion guardara el tamaño del vector\n"
		ref = temp
		Ast.GetH()
		codigo3d += "/***********************************************/\n"
	}
	if !estructura {
		codigo3d += "/********************AGREGAR ELEMENTO AL VECTOR*/\n"
		codigo3d += "heap[(int)H] = " + referencia + "; //Agregando elemento al vector\n"
		codigo3d += "H = H + 1;\n"
		Ast.GetH()
		codigo3d += "/***********************************************/\n"
	}

	return ref, codigo3d
}
