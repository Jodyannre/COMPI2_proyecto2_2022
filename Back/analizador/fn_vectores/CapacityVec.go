package fn_vectores

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"
)

type CapacityVec struct {
	Identificador interface{}
	Tipo          Ast.TipoDato
	Fila          int
	Columna       int
}

func NewCapacityVec(id interface{}, tipo Ast.TipoDato, fila, columna int) CapacityVec {
	nP := CapacityVec{
		Identificador: id,
		Tipo:          tipo,
		Fila:          fila,
		Columna:       columna,
	}
	return nP
}

func (p CapacityVec) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	var simbolo Ast.Simbolo
	var vector expresiones.Vector
	var id string
	/***********VARIABLES C3D*************/
	var idExp expresiones.Identificador
	var obj3d, obj3dValor Ast.O3D
	var codigo3d string
	/*************************************/

	//Primero verificar que sea un identificador el id
	_, tipoParticular := p.Identificador.(Ast.Abstracto).GetTipo()
	if tipoParticular != Ast.IDENTIFICADOR {
		//Error se espera un identificador
		////////////////////////////ERROR//////////////////////////////////
		return errores.GenerarError(39, p, p, "",
			Ast.ValorTipoDato[tipoParticular],
			"",
			scope)
		//////////////////////////////////////////////////////////////////
	}
	//Recuperar el id del identificador
	id = p.Identificador.(expresiones.Identificador).Valor

	//Verificar que el id exista
	if !scope.Exist(id) {
		//Error la variable no existe
		////////////////////////////ERROR//////////////////////////////////
		return errores.GenerarError(15, p, p, id,
			"",
			"",
			scope)
		//////////////////////////////////////////////////////////////////
	}
	//Conseguir el simbolo y el vector
	simbolo = scope.GetSimbolo(id)

	/***************************CODIGO 3D*************************/
	codigo3d += "/********************************ACCESO A VECTOR*/\n"
	idExp = expresiones.NewIdentificador(id, Ast.IDENTIFICADOR, 0, 0)
	obj3dValor = idExp.GetValue(scope).Valor.(Ast.O3D)
	codigo3d += obj3dValor.Codigo
	/************************************************************/

	//Verificar que sea un vector
	if simbolo.Tipo != Ast.VECTOR {
		////////////////////////////ERROR//////////////////////////////////
		return errores.GenerarError(34, p, p, "",
			Ast.ValorTipoDato[simbolo.Tipo],
			"",
			scope)
		//////////////////////////////////////////////////////////////////
	}
	vector = simbolo.Valor.(Ast.TipoRetornado).Valor.(expresiones.Vector)

	/*Codigo 3D*/
	obj3d.Codigo = codigo3d
	obj3d.Referencia = strconv.Itoa(vector.Capacity)
	obj3d.Valor = Ast.TipoRetornado{
		Tipo:  Ast.I64,
		Valor: vector.Capacity,
	}
	return Ast.TipoRetornado{
		Tipo:  Ast.VEC_CAPACITY,
		Valor: obj3d,
	}
}

func (v CapacityVec) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, v.Tipo
}

func (v CapacityVec) GetFila() int {
	return v.Fila
}
func (v CapacityVec) GetColumna() int {
	return v.Columna
}
