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

	/********************VARIABLES 3D*******************************/
	var obj3d, obj3dValor, obj3dCantidad Ast.O3D
	var codigo3d string
	var inicioVector string = Ast.GetTemp()
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

	sizeVector := 0
	tipoDelVector := Ast.TipoRetornado{Valor: true, Tipo: Ast.INDEFINIDO}
	vacio := true

	codigo3d += obj3dCantidad.Codigo
	codigo3d += obj3dValor.Codigo
	codigo3d += "/*****************************CREACION DE ARRAY*/\n"
	codigo3d += "/***********************************************/\n"
	codigo3d += inicioVector + " = H; //Guardar inicio del array \n"
	codigo3d += "H = H + 1;\n"
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

	/******************CODIGO 3D***********************/
	lt := Ast.GetLabel()
	lf := Ast.GetLabel()
	salto := Ast.GetLabel()
	preContadorArray := Ast.GetTemp()
	codigo3d += "/********************GUARDAR EL SIZE DEL VECTOR*/\n"
	codigo3d += "heap[(int)" + inicioVector + "] = " + strconv.Itoa(vector.Size) + ";\n"
	codigo3d += "/*************************GUARDAR LOS ELEMENTOS*/\n"
	codigo3d += salto + ":\n"
	codigo3d += "if (" + contadorArray + " < " + strconv.Itoa(vector.Size) + ") goto " + lt + ";\n"
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
	obj3d.Referencia = inicioVector
	obj3d.Valor = Ast.TipoRetornado{
		Tipo:  Ast.VECTOR,
		Valor: vector,
	}

	return Ast.TipoRetornado{
		Tipo:  Ast.VECTOR,
		Valor: obj3d,
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
