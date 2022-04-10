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
	Referencia string
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
	var codigo3d string
	var referencia string
	var referenciaVectorOriginal string = v.Referencia
	//var contadorVectorOriginal string = Ast.GetTemp()
	var sizeVectorOriginal string = Ast.GetTemp()
	var contadorPosicionVector string = Ast.GetTemp()
	var inicioVectorNuevo string = Ast.GetTemp()
	var contadorVectorClone string = Ast.GetTemp()
	var ElementoActual string = Ast.GetTemp()
	var ReferenciaElementoActual string = Ast.GetTemp()
	var elementoString Ast.TipoRetornado
	var elementoAbstracto interface{}
	var obj3dValor, obj3d, obj3dElementos, obj3dPreElemento Ast.O3D
	nLista := arraylist.New()
	listaFinal := arraylist.New()
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
	codigo3d += "/***************OBTENER VALORES A CLONAR*/\n"

	codigo3d += sizeVectorOriginal + " = heap[(int)" + referenciaVectorOriginal + "]; //Get size de vec \n"
	codigo3d += contadorPosicionVector + " = " + referenciaVectorOriginal + " + 1; //Get primera pos del vec a clonar\n"
	//codigo3d += referenciaVectorOriginal + " = " + contadorPosicionVector + ";\n"
	codigo3d += "/**************************CLONAR VECTOR*/\n"

	codigo3d += inicioVectorNuevo + " = " + " H; //Guardar Inicio del vector clone\n"
	codigo3d += "heap[(int)H] = " + sizeVectorOriginal + "; //Guardar size del vector\n"
	codigo3d += "H = H + 1;\n"
	codigo3d += "H = H + " + sizeVectorOriginal + "; //Apartar el espacio para los elementos del vector\n"
	codigo3d += contadorVectorClone + " = " + inicioVectorNuevo + "; //Iniciar contador para el nuevo vector \n"
	codigo3d += "heap[(int)" + inicioVectorNuevo + "] = " + sizeVectorOriginal + "; //Agregar size al vector clone \n"
	referencia = inicioVectorNuevo
	for i := 0; i < v.Valor.Len(); i++ {
		Ast.GetH()
	}

	for i := 0; i < v.Valor.Len(); i++ {
		elemento := v.Valor.GetValue(i).(Ast.TipoRetornado)
		if EsVAS(elemento.Tipo) {
			//codigo3d += contadorVectorOriginal + " = " + referenciaVectorOriginal + "; //Es vector, retornar a size\n"
			codigo3d += ElementoActual + " = " + "heap[(int)" + contadorPosicionVector + "]; //Get elemento \n"
			preReferencia := elemento.Valor.(Ast.Clones).SetReferencia(ElementoActual)
			preElemento := preReferencia.(Ast.Clones).Clonar(scope)
			obj3dPreElemento = preElemento.(Ast.TipoRetornado).Valor.(Ast.O3D)
			codigo3d += obj3dPreElemento.Codigo
			nElemento = preElemento
			obj3dElementos.Referencia = obj3dPreElemento.Referencia
		} else {
			if elemento.Tipo == Ast.STR || elemento.Tipo == Ast.STRING {
				codigo3d += ElementoActual + " = " + "heap[(int)" + contadorPosicionVector + "]; //Get elemento \n"
			} else {
				codigo3d += ElementoActual + " = " + "heap[(int)" + contadorPosicionVector + "]; //Get elemento \n"
			}
			//nElemento = elemento
			obj3dElementos.Valor = elemento
			if elemento.Tipo == Ast.STRING || elemento.Tipo == Ast.STR {
				elementoAbstracto = elemento
				elementoAbstracto = elementoAbstracto.(Ast.Clones).SetReferencia(ElementoActual)
				elementoString = elementoAbstracto.(Ast.Clones).Clonar(scope).(Ast.TipoRetornado)
				obj3dPreElemento = elementoString.Valor.(Ast.O3D)
				obj3dElementos.Referencia = obj3dPreElemento.Referencia
				codigo3d += obj3dPreElemento.Codigo
			} else {
				obj3dElementos.Referencia = Primitivo_To_String(elemento.Valor, elemento.Tipo)
				obj3dElementos.Referencia = ElementoActual
			}
			nElemento = Ast.TipoRetornado{
				Valor: obj3dElementos,
				Tipo:  elemento.Tipo,
			}
		}
		nLista.Add(nElemento)
		codigo3d += "/*******************ADD ELEMENTOS VECTOR*/\n"

		codigo3d += ReferenciaElementoActual + " = " + obj3dElementos.Referencia + "; //Get referencia del valor\n"
		codigo3d += contadorVectorClone + " = " + contadorVectorClone + " + 1" + "; //sig posicion vec nuevo\n"
		codigo3d += "heap[(int)" + contadorVectorClone + "] = " + ReferenciaElementoActual + "; //Add elemento\n"
		//codigo3d += contadorVectorClone + " = " + contadorVectorClone + " + 1; //sig pos de contador vec clone\n"
		codigo3d += contadorPosicionVector + " = " + contadorPosicionVector + " + 1; // sig pos vec original \n"
		/*
			if elemento.Tipo == Ast.STR || elemento.Tipo == Ast.STRING {
				codigo3d += contadorPosicionVector + " = " + contadorPosicionVector + " + 1; // sig pos vec original \n"
			}
		*/
		codigo3d += "/****************************************/\n"
	}

	codigo3d += "/****************************************/\n"
	codigo3d += "/****************************************/\n"

	for i := 0; i < nLista.Len(); i++ {
		obj3dValor = nLista.GetValue(i).(Ast.TipoRetornado).Valor.(Ast.O3D)
		listaFinal.Add(obj3dValor.Valor)
	}

	nV.Valor = listaFinal
	obj3d.Valor = Ast.TipoRetornado{
		Tipo:  Ast.VECTOR,
		Valor: nV,
	}

	obj3d.Codigo = codigo3d
	obj3d.Referencia = referencia

	return Ast.TipoRetornado{
		Tipo:  Ast.VECTOR,
		Valor: obj3d,
	}

}

func (v Vector) SetReferencia(referencia string) interface{} {
	v.Referencia = referencia
	return v
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
