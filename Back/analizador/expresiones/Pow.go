package expresiones

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"math"
	"strconv"
)

type Pow struct {
	Tipo     Ast.TipoDato
	TipoOp   Ast.TipoDato
	Valor    interface{}
	Potencia interface{}
	Fila     int
	Columna  int
}

func NewPow(tipo Ast.TipoDato, tipoOp Ast.TipoDato, valor interface{}, potencia interface{}, fila, columna int) Pow {
	nP := Pow{
		Tipo:     tipo,
		TipoOp:   tipoOp,
		Valor:    valor,
		Potencia: potencia,
		Fila:     fila,
		Columna:  columna,
	}
	return nP
}

func (p Pow) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	valor := p.Valor.(Ast.Expresion).GetValue(scope)
	potencia := p.Potencia.(Ast.Expresion).GetValue(scope)

	if valor.Tipo == Ast.ERROR {
		return valor
	}

	if potencia.Tipo == Ast.ERROR {
		return potencia
	}

	if valor.Tipo != p.TipoOp {
		//Error, tipos diferentes en la operacion
		fila := p.Valor.(Ast.Abstracto).GetFila()
		columna := p.Valor.(Ast.Abstracto).GetColumna()
		msg := "Semantic error, expected " + Ast.ValorTipoDato[p.TipoOp] + ", found " + Ast.ValorTipoDato[valor.Tipo] + "." +
			" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
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

	if potencia.Tipo != p.TipoOp {
		//Error, tipos diferentes en la operacion
		fila := p.Potencia.(Ast.Abstracto).GetFila()
		columna := p.Potencia.(Ast.Abstracto).GetColumna()
		msg := "Semantic error, expected " + Ast.ValorTipoDato[p.TipoOp] + ", found " + Ast.ValorTipoDato[potencia.Tipo] + "." +
			" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
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

	//Todo bien, operar
	respuesta := Ast.TipoRetornado{
		Tipo:  p.TipoOp,
		Valor: true,
	}
	if p.TipoOp == Ast.I64 {
		respuesta.Valor = int(math.Pow(float64(valor.Valor.(int)), float64(potencia.Valor.(int))))
	} else {
		respuesta.Valor = math.Pow(valor.Valor.(float64), potencia.Valor.(float64))
	}
	return respuesta
}

func (p Pow) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, p.Tipo
}

func (p Pow) GetFila() int {
	return p.Fila
}
func (p Pow) GetColumna() int {
	return p.Columna
}
