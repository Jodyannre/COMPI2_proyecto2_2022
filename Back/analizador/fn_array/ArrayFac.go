package fn_array

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"

	"github.com/colegno/arraylist"
)

type ArrayFactorial struct {
	Tipo      Ast.TipoDato
	Elementos *arraylist.List
	TipoArray Ast.TipoDato
	Fila      int
	Columna   int
}

func NewArrayFactorial(elementos *arraylist.List, fila, columna int) ArrayFactorial {
	//Crear el vector dependiendo de las banderas
	nV := ArrayFactorial{
		Tipo:      Ast.ARRAY_FAC,
		Fila:      fila,
		Columna:   columna,
		TipoArray: Ast.INDEFINIDO,
		Elementos: elementos,
	}
	return nV
}

func (v ArrayFactorial) GetValue(scope *Ast.Scope) Ast.TipoRetornado {

	/********************VARIABLES 3D*******************************/
	var obj3d, obj3dValor, obj3dCantidad Ast.O3D
	var codigo3d string
	var inicioArray string = Ast.GetTemp()
	var contadorArray string = Ast.GetTemp()
	/***************************************************************/
	//Se crea como factorial
	//Crear la cantidad de elementos que se solicita
	//conseguir la cantidad de veces que se va a repetir el valor
	elementoVeces := v.Elementos.GetValue(1)
	_, tipoParticular := elementoVeces.(Ast.Abstracto).GetTipo()
	cantidad := elementoVeces.(Ast.Expresion).GetValue(scope)
	obj3dCantidad = cantidad.Valor.(Ast.O3D)
	cantidad = obj3dCantidad.Valor

	elemento := v.Elementos.GetValue(0).(Ast.Expresion).GetValue(scope)
	obj3dValor = elemento.Valor.(Ast.O3D)
	elemento = obj3dValor.Valor

	sizeArray := 0
	tipoDelVector := Ast.INDEFINIDO
	tipoDelArray := Ast.TipoRetornado{Tipo: Ast.INDEFINIDO, Valor: true}
	vacio := true
	var tipoArray Ast.TipoDato

	codigo3d += obj3dCantidad.Codigo
	codigo3d += obj3dValor.Codigo
	codigo3d += "/*****************************CREACION DE ARRAY*/\n"
	codigo3d += "/***********************************************/\n"
	codigo3d += inicioArray + " = H; //Guardar inicio del array \n"
	codigo3d += contadorArray + " = 0; //Inicio contador \n"

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
	//Reiniciamos el valor del array
	elementos := arraylist.New()
	for i := 0; i < cantidad.Valor.(int); i++ {
		nElemento := elemento
		if nElemento.Tipo == Ast.ERROR {
			return nElemento
		}
		if nElemento.Tipo == Ast.ARRAY {
			tipoDelArray = nElemento.Valor.(expresiones.Array).TipoDelArray
		}
		if nElemento.Tipo == Ast.VECTOR {
			tipoDelVector = GetTipoVector(nElemento.Valor.(expresiones.Vector))
		}
		elementos.Add(nElemento)
		sizeArray++
		if vacio {
			vacio = false
			tipoArray = elemento.Tipo
		}
	}

	newArray := expresiones.NewArray(elementos, tipoArray, sizeArray, v.Fila, v.Columna)
	newArray.TipoDelVector = tipoDelVector
	if tipoDelArray.Tipo == Ast.ARRAY {
		newArray.TipoDelArray = Ast.TipoRetornado{Tipo: Ast.ARRAY, Valor: tipoDelArray}
	} else {
		newArray.TipoDelArray = tipoDelArray
	}
	if newArray.TipoDelArray.Tipo == Ast.INDEFINIDO {
		newArray.TipoDelArray = Ast.TipoRetornado{
			Tipo:  tipoArray,
			Valor: true,
		}
	}

	/******************CODIGO 3D***********************/
	lt := Ast.GetLabel()
	lf := Ast.GetLabel()
	salto := Ast.GetLabel()
	preContadorArray := Ast.GetTemp()
	codigo3d += "/*********************GUARDAR EL SIZE DEL ARRAY*/\n"
	codigo3d += "heap[(int)" + inicioArray + "] = " + strconv.Itoa(newArray.Size) + ";\n"
	codigo3d += "/*************************GUARDAR LOS ELEMENTOS*/\n"
	codigo3d += salto + ":\n"
	codigo3d += "if (" + contadorArray + " <= " + strconv.Itoa(newArray.Size) + ") goto " + lt + ";\n"
	codigo3d += "goto " + lf + ";\n"
	codigo3d += lt + ":\n"
	codigo3d += "heap[(int)H] = " + obj3dValor.Referencia + "; //Copiar elemento\n"
	codigo3d += "H = H + 1;"
	Ast.GetH()
	codigo3d += preContadorArray + " = " + contadorArray + " + 1;\n"
	codigo3d += contadorArray + " = " + preContadorArray + "; //Actualizar contador \n"
	codigo3d += "goto " + salto + ";\n"
	codigo3d += lf + ":\n"
	codigo3d += "/***********************************************/\n"
	codigo3d += "/***********************************************/\n"
	/**************************************************/

	obj3d.Codigo = codigo3d
	obj3d.Referencia = inicioArray
	obj3d.Valor = Ast.TipoRetornado{
		Tipo:  Ast.ARRAY,
		Valor: newArray,
	}

	return Ast.TipoRetornado{
		Tipo:  Ast.ARRAY,
		Valor: obj3d,
	}
}

func (v ArrayFactorial) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, v.Tipo
}

func (v ArrayFactorial) GetFila() int {
	return v.Fila
}
func (v ArrayFactorial) GetColumna() int {
	return v.Columna
}
