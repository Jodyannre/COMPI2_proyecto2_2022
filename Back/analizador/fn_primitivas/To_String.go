package fn_primitivas

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"strconv"
)

type ToString struct {
	Tipo    Ast.TipoDato
	Valor   interface{}
	Fila    int
	Columna int
}

func NewToString(tipo Ast.TipoDato, valor interface{}, fila, columna int) ToString {
	nT := ToString{
		Tipo:    tipo,
		Valor:   valor,
		Fila:    fila,
		Columna: columna,
	}
	return nT
}

func (t ToString) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	valorSalida := Ast.TipoRetornado{}

	//Primero conseguir el valor a convertir
	valor := t.Valor.(Ast.Expresion).GetValue(scope)

	//Verificar que el valor no sea un error
	if valor.Tipo == Ast.ERROR {
		return valor
	}

	//Verificar el tipo del valor y convertir
	if valor.Tipo > 6 {
		//Error, ese tipo no puede ser convertido a string
		msg := "Semantic error, " + Ast.ValorTipoDato[valor.Tipo] + " type can't be converted to string." +
			" -- Line: " + strconv.Itoa(t.Fila) +
			" Column: " + strconv.Itoa(t.Columna)
		nError := errores.NewError(t.Fila, t.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		valorSalida.Valor = nError
		valorSalida.Tipo = Ast.ERROR
		return valorSalida
	}
	//Salida
	switch valor.Tipo {
	case Ast.I64, Ast.USIZE:
		valorSalida.Valor = strconv.Itoa(valor.Valor.(int))
	case Ast.F64:
		valorSalida.Valor = strconv.FormatFloat(valor.Valor.(float64), 'f', -1, 64)
	case Ast.STR:
		valorSalida.Valor = valor.Valor.(string)
	case Ast.STRING:
		valorSalida.Valor = valor.Valor.(string)
	case Ast.STRING_OWNED:
		valorSalida.Valor = valor.Valor.(string)
	case Ast.BOOLEAN:
		valorSalida.Valor = strconv.FormatBool(valor.Valor.(bool))
	case Ast.CHAR:
		valorSalida.Valor = valor.Valor.(string)
	}
	valorSalida.Tipo = Ast.STRING
	return valorSalida
}

func (f ToString) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, f.Tipo
}

func (f ToString) GetFila() int {
	return f.Fila
}
func (f ToString) GetColumna() int {
	return f.Columna
}
