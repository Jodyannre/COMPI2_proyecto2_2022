package fn_primitivas

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"math"
	"strconv"
)

type Sqrt struct {
	Tipo    Ast.TipoDato
	Valor   interface{}
	Fila    int
	Columna int
}

func NewSqrt(tipo Ast.TipoDato, valor interface{}, fila, columna int) Sqrt {
	nS := Sqrt{
		Tipo:    tipo,
		Valor:   valor,
		Fila:    fila,
		Columna: columna,
	}
	return nS
}

func (s Sqrt) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	//Primero conseguir el valor
	valorSalida := Ast.TipoRetornado{}

	//Primero conseguir el valor a convertir
	valor := s.Valor.(Ast.Expresion).GetValue(scope)

	//Verificar que el valor no sea un error
	if valor.Tipo == Ast.ERROR {
		return valor
	}

	//Verificar el tipo del valor y convertir
	if valor.Tipo != Ast.F64 {
		//Error, ese tipo no puede ser convertido a string
		msg := "Semantic error, sqrt method only accepts f64 values." +
			" -- Line: " + strconv.Itoa(s.Fila) +
			" Column: " + strconv.Itoa(s.Columna)
		nError := errores.NewError(s.Fila, s.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		valorSalida.Valor = nError
		valorSalida.Tipo = Ast.ERROR
		return valorSalida
	}

	//Get la raiz cuadrada del valor
	valorSalida.Valor = math.Sqrt(valor.Valor.(float64))
	valorSalida.Tipo = Ast.F64
	return valorSalida

}

func (f Sqrt) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, f.Tipo
}

func (f Sqrt) GetFila() int {
	return f.Fila
}
func (f Sqrt) GetColumna() int {
	return f.Columna
}
