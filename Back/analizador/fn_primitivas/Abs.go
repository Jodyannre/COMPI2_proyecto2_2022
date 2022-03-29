package fn_primitivas

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"math"
	"strconv"
)

type Abs struct {
	Tipo    Ast.TipoDato
	Valor   interface{}
	Fila    int
	Columna int
}

func NewAbs(tipo Ast.TipoDato, valor interface{}, fila, columna int) Abs {
	nS := Abs{
		Tipo:    tipo,
		Valor:   valor,
		Fila:    fila,
		Columna: columna,
	}
	return nS
}

func (a Abs) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	//Primero conseguir el valor
	valorSalida := Ast.TipoRetornado{}

	//Primero conseguir el valor a convertir
	valor := a.Valor.(Ast.Expresion).GetValue(scope)

	//Verificar que el valor no sea un error
	if valor.Tipo == Ast.ERROR {
		return valor
	}

	//Verificar el tipo del valor y convertir
	if valor.Tipo != Ast.F64 && valor.Tipo != Ast.I64 {
		//Error, ese tipo no puede ser convertido a string
		msg := "Semantic error, abs method only accepts numeric values." +
			" -- Line: " + strconv.Itoa(a.Fila) +
			" Column: " + strconv.Itoa(a.Columna)
		nError := errores.NewError(a.Fila, a.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		valorSalida.Valor = nError
		valorSalida.Tipo = Ast.ERROR
		return valorSalida
	}

	//Get la raiz cuadrada del valor
	if valor.Tipo == Ast.I64 {
		if valor.Valor.(int) < 0 {
			valorSalida.Valor = valor.Valor.(int) * -1
		} else {
			valorSalida.Valor = valor.Valor.(int)
		}
		valorSalida.Tipo = Ast.I64
	}
	if valor.Tipo == Ast.F64 {
		valorSalida.Valor = math.Abs(valor.Valor.(float64))
		valorSalida.Tipo = Ast.F64
	}

	return valorSalida

}

func (f Abs) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, f.Tipo
}

func (f Abs) GetFila() int {
	return f.Fila
}
func (f Abs) GetColumna() int {
	return f.Columna
}
