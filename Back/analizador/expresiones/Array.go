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
	Referencia    string
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
	/******************************VARIABLES 3D*********************************/
	var codigo3d, tempContadorOriginal, tempContadorClone string
	var contadorArrayOriginal, contadorArrayClone string
	var referenciaArrayOriginal, sizeArrayOriginal, posicionNuevoArray string
	var elementoActual string
	var obj3d, obj3dElemento Ast.O3D
	/***************************************************************************/

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

	/*******************INICIALIZACION DE TEMPORALES***********************/
	referenciaArrayOriginal = Ast.GetTemp()
	sizeArrayOriginal = Ast.GetTemp()
	contadorArrayOriginal = Ast.GetTemp()
	posicionNuevoArray = Ast.GetTemp()
	contadorArrayClone = Ast.GetTemp()
	elementoActual = Ast.GetTemp()
	tempContadorOriginal = Ast.GetTemp()
	tempContadorClone = Ast.GetTemp()
	//Referencia del array
	codigo3d += referenciaArrayOriginal + " = " + a.Referencia + "; //Copia de ref del array \n"
	//Size del array
	codigo3d += sizeArrayOriginal + " = heap[(int)" + referenciaArrayOriginal + "] //Get size \n"
	//Iniciar contador del array original
	codigo3d += contadorArrayOriginal + " = " + referenciaArrayOriginal + " + 1; //Get primera pos de array a clonar\n"
	//Guardar posicion del nuevo array
	codigo3d += posicionNuevoArray + " = H; //Guardar pos del array clone\n"
	//Guardar el size en la nueva pos
	codigo3d += "heap[(int)" + posicionNuevoArray + "] = " + sizeArrayOriginal + "; //Agregar el size al nuevo array\n"
	codigo3d += "H = H + 1; \n"
	//Iniciar contador para el array nuevo
	codigo3d += contadorArrayClone + " = H; //Iniciar contador para el array clone \n"
	/*********************************************************************/

	for i := 0; i < a.Elementos.Len(); i++ {
		elemento := a.Elementos.GetValue(i).(Ast.TipoRetornado)
		if EsVAS(elemento.Tipo) {
			codigo3d += elementoActual + " = " + "heap[(int)" + contadorArrayOriginal + "]; //Get elemento \n"
			preReferencia := elemento.Valor.(Ast.Clones).SetReferencia(elementoActual)
			preElemento := preReferencia.(Ast.Clones).Clonar(scope)
			obj3dElemento = preElemento.(Ast.TipoRetornado).Valor.(Ast.O3D)
			codigo3d += obj3dElemento.Codigo
			nElemento = obj3dElemento.Valor
			elementoActual = obj3dElemento.Referencia
		} else {
			codigo3d += elementoActual + " = " + "heap[(int)" + contadorArrayOriginal + "]; //Get elemento \n"
			nElemento = elemento
		}

		/************GET REFERENCIA DEL ELEMENTO ACTUAL SI ES STRING*********************/
		if elemento.Tipo == Ast.STRING || elemento.Tipo == Ast.STR {
			var elementoAbstracto interface{}
			var elementoString Ast.TipoRetornado
			elementoAbstracto = elemento
			elementoAbstracto = elementoAbstracto.(Ast.Clones).SetReferencia(elementoActual)
			elementoString = elementoAbstracto.(Ast.Clones).Clonar(scope).(Ast.TipoRetornado)
			obj3dElemento = elementoString.Valor.(Ast.O3D)
			codigo3d += obj3dElemento.Codigo
			elementoActual = obj3dElemento.Referencia
			nElemento = obj3dElemento.Valor
		}
		/********************************************************************************/
		nElementos.Add(nElemento)
		/**********************AGREGAR ELEMENTO A ARRAY EN 3D****************************/
		codigo3d += "heap[(int)" + contadorArrayClone + "] = " + elementoActual + "; Add nuevo elemento\n"
		/********************************************************************************/

		/***************************ACTUALIZAR CONTADORES********************************/
		codigo3d += tempContadorOriginal + " = " + contadorArrayOriginal + " + 1; Sig pos array original\n"
		codigo3d += contadorArrayOriginal + " = " + tempContadorOriginal + "; \n"
		codigo3d += tempContadorClone + " = " + contadorArrayClone + " + 1; Sig pos array clone \n"
		codigo3d += contadorArrayClone + " = " + tempContadorClone + "; \n"
		/********************************************************************************/
	}
	nV.Elementos = nElementos

	obj3d.Valor = Ast.TipoRetornado{
		Tipo:  Ast.ARRAY,
		Valor: nV,
	}

	obj3d.Codigo = codigo3d
	obj3d.Referencia = posicionNuevoArray

	return Ast.TipoRetornado{
		Tipo:  Ast.ARRAY,
		Valor: obj3d,
	}
}

func (v Array) GetMutable() bool {
	return v.Mutable
}

func (a Array) SetReferencia(referencia string) interface{} {
	a.Referencia = referencia
	return a
}
