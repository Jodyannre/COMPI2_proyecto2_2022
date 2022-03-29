package fn_primitivas

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"

	"github.com/colegno/arraylist"
)

type ToChars struct {
	Tipo    Ast.TipoDato
	Valor   interface{}
	Fila    int
	Columna int
}

func NewToChars(valor interface{}, fila, columna int) ToChars {
	nT := ToChars{
		Tipo:    Ast.TO_CHARS,
		Valor:   valor,
		Fila:    fila,
		Columna: columna,
	}
	return nT
}

func (t ToChars) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	listaCaracteres := arraylist.New()
	tipoVector := Ast.TipoRetornado{Tipo: Ast.CHAR, Valor: true}
	var vectorSalida interface{}
	//Primero conseguir el valor a convertir
	valor := t.Valor.(Ast.Expresion).GetValue(scope)

	//Verificar que el valor no sea un error
	if valor.Tipo == Ast.ERROR {
		return valor
	}

	//Verificar el tipo del valor y convertir
	if valor.Tipo != Ast.STR {
		//Error, ese tipo no puede ser convertido a string
		msg := "Semantic error, " + Ast.ValorTipoDato[Ast.STR] + " expected, found " + Ast.ValorTipoDato[valor.Tipo] +
			". -- Line: " + strconv.Itoa(t.Fila) +
			" Column: " + strconv.Itoa(t.Columna)
		nError := errores.NewError(t.Fila, t.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}
	//Crear la lista que contendrá el vector

	for caracter := range valor.Valor.(string) {
		//Ir creando los elementos, guardarlos en la lista y luego crear el vector
		nuevoSimbolo := Ast.TipoRetornado{
			Tipo:  Ast.CHAR,
			Valor: valor.Valor.(string)[caracter],
		}
		listaCaracteres.Add(nuevoSimbolo)
	}

	//Crear el vector de salida

	vectorSalida = expresiones.NewVector(listaCaracteres, tipoVector, listaCaracteres.Len(),
		listaCaracteres.Len(), false, t.Fila, t.Columna)

	return Ast.TipoRetornado{
		Tipo:  Ast.VECTOR,
		Valor: vectorSalida,
	}
}

func (f ToChars) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, f.Tipo
}

func (f ToChars) GetFila() int {
	return f.Fila
}
func (f ToChars) GetColumna() int {
	return f.Columna
}
