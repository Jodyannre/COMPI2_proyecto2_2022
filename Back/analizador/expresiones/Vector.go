package expresiones

import (
	"Back/analizador/Ast"
	"strconv"

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
	var contadorVectorOriginal string = Ast.GetTemp()
	var sizeVectorOriginal string = Ast.GetTemp()
	var contadorPosicionVector string = Ast.GetTemp()
	var inicioVectorNuevo string = Ast.GetTemp()
	var contadorVectorClone string = Ast.GetTemp()
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
	codigo3d += contadorVectorOriginal + " = 0; //Iniciar contador para el vector\n"
	codigo3d += sizeVectorOriginal + " = heap[(int)" + referenciaVectorOriginal + "]; //Get size de vec \n"
	codigo3d += contadorPosicionVector + " = " + referenciaVectorOriginal + " + 1; //Get primera pos del vec a clonar\n"

	codigo3d += "/**************************CLONAR VECTOR*/\n"
	codigo3d += inicioVectorNuevo + " = " + " H; //Guardar Inicio del vector clone\n"
	codigo3d += "heap[(int)H] = " + sizeVectorOriginal + "; //Guardar size del vector\n"
	codigo3d += "H = H + " + strconv.Itoa(nLista.Len()) + "; //Apartar el espacio para los elementos del vector\n"
	codigo3d += contadorVectorClone + " = 0; //Iniciar contador para el nuevo vector \n"
	for i := 0; i < v.Valor.Len(); i++ {
		Ast.GetH()
	}

	for i := 0; i < v.Valor.Len(); i++ {
		elemento := v.Valor.GetValue(i).(Ast.TipoRetornado)
		if EsVAS(elemento.Tipo) {
			preElemento := elemento.Valor.(Ast.Clones).Clonar(scope)
			/*
				obj3dPreElemento = preElemento.(Ast.TipoRetornado).Valor.(Ast.O3D)
				preElemento = obj3dPreElemento.Valor
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
			*/
			//obj3dPreElemento.Valor = valor
			//nElemento = valor
			obj3dPreElemento = preElemento.(Ast.TipoRetornado).Valor.(Ast.O3D)
			codigo3d += obj3dPreElemento.Codigo
			nElemento = preElemento
		} else {
			//nElemento = elemento
			obj3dElementos.Valor = elemento
			if elemento.Tipo == Ast.STRING || elemento.Tipo == Ast.STR {
				elementoAbstracto = elemento
				elementoString = elementoAbstracto.(Ast.Clones).Clonar(scope).(Ast.TipoRetornado)
				obj3dPreElemento = elementoString.Valor.(Ast.O3D)
				obj3dElementos.Referencia = obj3dPreElemento.Referencia
				codigo3d += obj3dPreElemento.Codigo
			} else {
				obj3dElementos.Referencia = Primitivo_To_String(elemento.Valor, elemento.Tipo)
			}
			nElemento = Ast.TipoRetornado{
				Valor: obj3dElementos,
				Tipo:  elemento.Tipo,
			}
		}
		nLista.Add(nElemento)
	}
	codigo3d += "/****************************************/\n"
	//nV.Valor = nLista

	codigo3d += "/********************ADD ELEMENTOS VECTOR*/\n"
	referencia = inicioVector
	for i := 0; i < nLista.Len(); i++ {
		obj3dValor = nLista.GetValue(i).(Ast.TipoRetornado).Valor.(Ast.O3D)
		codigo3d += "heap[(int)H] = " + obj3dValor.Referencia + ";//Agregar elemento al vector\n"
		codigo3d += "H = H + 1;\n"
		Ast.GetH()
		listaFinal.Add(obj3dValor.Valor)
	}
	codigo3d += "/****************************************/\n"
	codigo3d += "/****************************************/\n"

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
