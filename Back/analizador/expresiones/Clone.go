package expresiones

import (
	"Back/analizador/Ast"
)

type Clone struct {
	Expresion interface{}
	Tipo      Ast.TipoDato
	Fila      int
	Columna   int
}

func NewClone(expresion interface{}, fila, columna int) Clone {
	nC := Clone{
		Expresion: expresion,
		Tipo:      Ast.CLONE,
		Fila:      fila,
		Columna:   columna,
	}
	return nC
}

func (c Clone) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	/****************************VARIABLES 3D******************************/
	var obj3d, obj3dValor, obj3dClone Ast.O3D
	var codigo3d, referencia string
	/**********************************************************************/

	//El nuevo valor
	var nValor interface{}
	//Conseguir el valor
	valor := c.Expresion.(Ast.Expresion).GetValue(scope)
	obj3dValor = valor.Valor.(Ast.O3D)
	valor = obj3dValor.Valor
	codigo3d += obj3dValor.Codigo

	if valor.Tipo == Ast.ERROR {
		return valor
	}
	switch valor.Tipo {
	case Ast.STRUCT, Ast.VECTOR, Ast.ARRAY:
		preValor := valor.Valor.(Ast.Clones).SetReferencia(obj3dValor.Referencia)
		valor.Valor = preValor
		nValor = valor.Valor.(Ast.Clones).Clonar(scope)
		obj3dClone = nValor.(Ast.TipoRetornado).Valor.(Ast.O3D)
		nValor = obj3dClone.Valor
		codigo3d += obj3dClone.Codigo
		referencia = obj3dClone.Referencia
	default:
		nValor = valor
		referencia = Primitivo_To_String(valor.Valor, valor.Tipo)
	}

	obj3d.Codigo = codigo3d
	obj3d.Referencia = referencia
	obj3d.Valor = nValor.(Ast.TipoRetornado)

	return Ast.TipoRetornado{
		Tipo:  valor.Tipo,
		Valor: obj3d,
	}
}

func (c Clone) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, c.Tipo
}

func (c Clone) GetFila() int {
	return c.Fila
}
func (c Clone) GetColumna() int {
	return c.Columna
}
