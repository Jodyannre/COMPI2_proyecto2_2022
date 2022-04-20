package transferencia

import (
	"Back/analizador/Ast"
	"strconv"
)

type Return struct {
	Tipo      Ast.TipoDato
	Expresion Ast.Expresion
	Valor     Ast.TipoRetornado
	Fila      int
	Columna   int
}

func (r Return) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, r.Tipo
}

func NewReturn(tipo Ast.TipoDato, expresion Ast.Expresion, fila, columna int) Return {
	nR := Return{
		Tipo:      tipo,
		Expresion: expresion,
		Valor:     Ast.TipoRetornado{Valor: false, Tipo: Ast.VOID},
		Fila:      fila,
		Columna:   columna,
	}
	return nR
}

func (r Return) Run(scope *Ast.Scope) interface{} {
	/************************VARIABLES 3D**************************************/
	var obj3d, obj3dValor Ast.O3D
	var salto string = Ast.GetLabel()
	var codigo3d string
	var scopeOrigen *Ast.Scope
	var guardarScope string
	var posicionGuardar string
	/**************************************************************************/
	obj3d.SaltoReturn = salto
	obj3d.Salto = salto
	var valor Ast.TipoRetornado
	if r.Expresion != nil {
		valor = r.Expresion.GetValue(scope)
		obj3dValor = valor.Valor.(Ast.O3D)
		valor = obj3dValor.Valor
		codigo3d += obj3dValor.Codigo
	} else {
		valor = Ast.TipoRetornado{
			Tipo:  Ast.NULL,
			Valor: true,
		}
	}

	if valor.Tipo == Ast.ERROR {
		return valor
	}

	if r.Tipo == Ast.RETURN || valor.Tipo == Ast.NULL {
		obj3d.Valor = Ast.TipoRetornado{
			Tipo:  Ast.RETURN,
			Valor: r,
		}
		return Ast.TipoRetornado{
			Tipo:  Ast.RETURN,
			Valor: obj3d,
		}
	}

	valorRetornar := Ast.TipoRetornado{
		Tipo:  Ast.RETURN,
		Valor: valor.Valor,
	}
	guardarScope = Ast.GetTemp()
	posicionGuardar = Ast.GetTemp()
	obj3d.SaltoReturnExp = salto
	obj3d.Salto = salto
	obj3d.Valor = valorRetornar
	scopeOrigen = scope.GetEntornoPadreReturn()
	codigo3d += "/******************GUARDAR EL VALOR DEL RETURN*/\n"
	codigo3d += guardarScope + " = P; //Guardar scope anterior \n"
	codigo3d += "P = " + strconv.Itoa(scopeOrigen.Posicion) + "; //Cambio a entorno simulado \n"
	codigo3d += posicionGuardar + " = P + 0; //Pos del return \n"
	codigo3d += "stack[(int)" + posicionGuardar + "] = " + obj3dValor.Referencia + "; //Guardar valor \n"
	codigo3d += "P = " + guardarScope + "; //Retornar al entorno anterior \n"
	codigo3d += "/***********************************************/\n"
	obj3d.Codigo = codigo3d
	return Ast.TipoRetornado{
		Tipo:  r.Tipo,
		Valor: obj3d,
	}
}

func (op Return) GetFila() int {
	return op.Fila
}
func (op Return) GetColumna() int {
	return op.Columna
}
