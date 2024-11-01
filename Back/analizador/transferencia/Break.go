package transferencia

import (
	"Back/analizador/Ast"
	"strconv"
)

type Break struct {
	Tipo      Ast.TipoDato
	Expresion Ast.Expresion
	Valor     Ast.TipoRetornado
	Fila      int
	Columna   int
}

func (b Break) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, b.Tipo
}

func NewBreak(tipo Ast.TipoDato, expresion Ast.Expresion, fila, columna int) Break {
	nB := Break{
		Tipo:      tipo,
		Expresion: expresion,
		Valor:     Ast.TipoRetornado{Valor: false, Tipo: Ast.VOID},
		Fila:      fila,
		Columna:   columna,
	}
	return nB
}

func (b Break) Run(scope *Ast.Scope) interface{} {

	/**********************VARIABLES 3D*****************************/
	var obj3d, obj3dValor Ast.O3D
	var salto string = Ast.GetLabel()
	var codigo3d string
	var guardarScope, posicionGuardar string
	var scopeOrigen *Ast.Scope
	/***************************************************************/
	obj3d.SaltoTranferencia = salto
	obj3d.SaltoBreak = salto
	obj3d.Valor = Ast.TipoRetornado{
		Tipo:  Ast.BREAK,
		Valor: b,
	}
	obj3d.TranferenciaAgregada = false
	if b.Tipo == Ast.BREAK {

		return Ast.TipoRetornado{
			Tipo:  Ast.BREAK,
			Valor: obj3d,
		}
	}
	valor := b.Expresion.GetValue(scope)
	obj3dValor = valor.Valor.(Ast.O3D)
	codigo3d += obj3dValor.Codigo
	valor = obj3dValor.Valor

	if valor.Tipo == Ast.ERROR {
		return valor
	}

	valorRetornar := Ast.TipoRetornado{
		Tipo:  valor.Tipo,
		Valor: valor.Valor,
	}

	guardarScope = Ast.GetTemp()
	posicionGuardar = Ast.GetTemp()
	scopeOrigen = scope.GetEntornoPadreBreak()
	codigo3d += "/********************GUARDAR EL VALOR DEL BREAK*/\n"
	codigo3d += guardarScope + " = P; //Guardar scope anterior \n"
	codigo3d += "P = " + strconv.Itoa(scopeOrigen.Posicion) + "; //Cambio a entorno simulado \n"
	codigo3d += posicionGuardar + " = P + 0; //Pos del break \n"
	codigo3d += "stack[(int)" + posicionGuardar + "] = " + obj3dValor.Referencia + "; //Guardar valor \n"
	codigo3d += "P = " + guardarScope + "; //Retornar al entorno anterior \n"
	codigo3d += "/***********************************************/\n"

	obj3d.Valor = valorRetornar
	obj3d.Codigo = codigo3d
	obj3d.SaltoTranferencia = salto
	obj3d.SaltoBreak = salto
	obj3d.SaltoBreakExp = salto
	obj3d.Referencia = obj3dValor.Referencia

	return Ast.TipoRetornado{
		Tipo:  b.Tipo,
		Valor: obj3d,
	}
}

func (op Break) GetFila() int {
	return op.Fila
}
func (op Break) GetColumna() int {
	return op.Columna
}
