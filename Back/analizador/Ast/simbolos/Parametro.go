package simbolos

import "Back/analizador/Ast"

type Parametro struct {
	Identificador   string
	Tipo            Ast.TipoDato
	Fila            int
	Columna         int
	Mutable         bool
	Referencia      bool
	TipoDeclaracion Ast.TipoRetornado
}

func NewParametro(id string, tipo Ast.TipoDato, tipoD Ast.TipoRetornado, mutable bool,
	referencia bool, fila, columna int) Parametro {
	nP := Parametro{
		Identificador:   id,
		Tipo:            tipo,
		Mutable:         mutable,
		Fila:            fila,
		Columna:         columna,
		TipoDeclaracion: tipoD,
		Referencia:      referencia,
	}
	return nP
}

func (p Parametro) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	var obj3d Ast.O3D
	obj3d.Valor = Ast.TipoRetornado{
		Tipo:  p.TipoDeclaracion.Tipo,
		Valor: p.Identificador,
	}
	return Ast.TipoRetornado{
		Tipo:  p.TipoDeclaracion.Tipo,
		Valor: obj3d,
	}
}

func (p Parametro) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, p.Tipo
}

func (p Parametro) GetFila() int {
	return p.Fila
}
func (p Parametro) GetColumna() int {
	return p.Columna
}

func (a Parametro) FormatearTipo(scope *Ast.Scope) Ast.TipoRetornado {
	if EsPosibleReferencia(a.TipoDeclaracion.Tipo) {
		nTipo := GetTipoEstructura(a.TipoDeclaracion, scope, a)
		return nTipo
	} else {
		return a.TipoDeclaracion
	}
}
