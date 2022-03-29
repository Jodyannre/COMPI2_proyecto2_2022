package simbolos

import (
	"Back/analizador/Ast"
)

type Tipo struct {
	Tipo     Ast.TipoDato
	TipoDato Ast.TipoDato
	Valor    interface{}
	Fila     int
	Columna  int
}

func NewTipo(tipo Ast.TipoDato, valor interface{}, fila, columna int) Tipo {
	nT := Tipo{
		Tipo:     Ast.TIPO,
		TipoDato: tipo,
		Valor:    valor,
		Fila:     fila,
		Columna:  columna,
	}
	return nT
}

func (t Tipo) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	return Ast.TipoRetornado{
		Tipo:  Ast.TIPO,
		Valor: t,
	}
}

func (op Tipo) GetFila() int {
	return op.Fila
}
func (op Tipo) GetColumna() int {
	return op.Columna
}

func (d Tipo) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, Ast.TIPO
}

func (d Tipo) GetTipoFinal() Ast.TipoRetornado {
	return getTipoFinal(d)
}

func getTipoFinal(tipo Tipo) Ast.TipoRetornado {
	if esTipoFinal(tipo.TipoDato) {
		if tipo.Tipo != Ast.STRUCT {
			return Ast.TipoRetornado{
				Tipo:  tipo.TipoDato,
				Valor: true,
			}
		}
		return Ast.TipoRetornado{
			Tipo:  tipo.TipoDato,
			Valor: tipo.Valor,
		}
	} else {
		return getTipoFinal(tipo.Valor.(Tipo))
	}
}

func esTipoFinal(tipo Ast.TipoDato) bool {
	switch tipo {
	case Ast.I64, Ast.F64, Ast.CHAR, Ast.STRING, Ast.STR, Ast.USIZE, Ast.BOOLEAN, Ast.STRUCT, Ast.INDEFINIDO,
		Ast.DIMENSION_ARRAY:
		return true
	default:
		return false
	}
}

func CompararTipos(tipoA Ast.TipoRetornado, tipoB Ast.TipoRetornado) bool {
	if esTipoFinal(tipoA.Tipo) && esTipoFinal(tipoB.Tipo) {
		if tipoA.Tipo == tipoB.Tipo {
			//Verificar si son structs
			if tipoA.Tipo == Ast.STRUCT {
				if tipoA.Valor == tipoB.Valor {
					return true
				} else {
					return false
				}
			}
			return true
		} else {
			return false
		}
	}
	if esTipoFinal(tipoA.Tipo) && !esTipoFinal(tipoB.Tipo) ||
		!esTipoFinal(tipoA.Tipo) && esTipoFinal(tipoB.Tipo) {
		return false
	}
	return CompararTipos(tipoA.Valor.(Ast.TipoRetornado), tipoB.Valor.(Ast.TipoRetornado))
}

func Tipo_String(t Ast.TipoRetornado) string {
	if t.Tipo == Ast.VECTOR {
		return "Vec <" + Tipo_String(t.Valor.(Ast.TipoRetornado)) + ">"
	} else {
		if t.Tipo == Ast.STRUCT {
			return t.Valor.(string)
		} else {
			return Ast.ValorTipoDato[t.Tipo]
		}
	}
}
