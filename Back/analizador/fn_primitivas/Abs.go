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
	/********************************VARIABLES 3D*****************************/
	var obj3d, obj3dValor Ast.O3D
	var codigo3d string
	var nuevoValor string
	/************************************************************************/

	//Primero conseguir el valor
	valorSalida := Ast.TipoRetornado{}
	nuevoValor = Ast.GetTemp()

	//Primero conseguir el valor a convertir
	codigo3d += "/************************************METODO ABS*/\n"

	valor := a.Valor.(Ast.Expresion).GetValue(scope)
	obj3dValor = valor.Valor.(Ast.O3D)
	valor = obj3dValor.Valor
	codigo3d += obj3dValor.Codigo

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
			codigo3d += nuevoValor + " = " + obj3dValor.Referencia + " * -1; //Get nuevo valor \n"
		} else {
			codigo3d += nuevoValor + " = " + obj3dValor.Referencia + "; //Get nuevo valor \n"
			valorSalida.Valor = valor.Valor.(int)
		}
		valorSalida.Tipo = Ast.I64
	}
	if valor.Tipo == Ast.F64 {
		valorSalida.Valor = math.Abs(valor.Valor.(float64))
		if valor.Valor.(float64) < 0 {
			codigo3d += nuevoValor + " = " + obj3dValor.Referencia + " * -1; //Get nuevo valor \n"
		} else {
			codigo3d += nuevoValor + " = " + obj3dValor.Referencia + "; //Get nuevo valor \n"
		}
		valorSalida.Tipo = Ast.F64
	}

	codigo3d += "/***********************************************/\n"
	obj3d.Valor = valorSalida
	obj3d.Codigo = codigo3d
	obj3d.Referencia = nuevoValor
	return Ast.TipoRetornado{
		Valor: obj3d,
		Tipo:  valorSalida.Tipo,
	}

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
