package fn_vectores

import (
	"Back/analizador/Ast"
	"Back/analizador/expresiones"

	"github.com/colegno/arraylist"
)

type VecNew struct {
	Tipo    Ast.TipoDato
	Fila    int
	Columna int
}

func NewVecNew(fila, columna int) VecNew {
	//Crear el vector dependiendo de las banderas
	nV := VecNew{
		Tipo:    Ast.VEC_NEW,
		Fila:    fila,
		Columna: columna,
	}
	return nV
}

func (w VecNew) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	/***********VARIABLES C3D*************/
	var obj3d Ast.O3D
	var codigo3d string
	var referencia string = Ast.GetTemp()
	/*************************************/
	elementos := arraylist.New()
	vector := expresiones.NewVector(elementos, Ast.TipoRetornado{Tipo: Ast.INDEFINIDO, Valor: true}, 0, 0, true, w.Fila, w.Columna)
	/*****************CODIGO 3D************************/
	codigo3d += "/****************************CREACION DE VECTOR*/\n"
	codigo3d += referencia + " = H;//Guardar la referencia\n"
	codigo3d += "heap[(int)H] = 0; //Guardar el size del vector\n"
	codigo3d += "H = H + 1;\n"
	codigo3d += "/***********************************************/\n"
	Ast.GetH()
	obj3d.Valor = Ast.TipoRetornado{
		Tipo:  Ast.VECTOR,
		Valor: vector,
	}
	obj3d.Codigo = codigo3d
	obj3d.Referencia = referencia
	/**************************************************/
	return Ast.TipoRetornado{
		Tipo:  Ast.VECTOR,
		Valor: obj3d,
	}
}

func (v VecNew) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, v.Tipo
}

func (v VecNew) GetFila() int {
	return v.Fila
}
func (v VecNew) GetColumna() int {
	return v.Columna
}
