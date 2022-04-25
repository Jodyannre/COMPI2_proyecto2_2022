package expresiones

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"strconv"
)

type Cast struct {
	Valor        Ast.Expresion
	Tipo         Ast.TipoDato
	TipoObjetivo Ast.TipoDato
	Fila         int
	Columna      int
}

func NewCast(valor Ast.Expresion, tipo Ast.TipoDato, tipoOb Ast.TipoDato, fila, columna int) Cast {
	nC := Cast{
		Valor:        valor,
		Tipo:         tipo,
		TipoObjetivo: tipoOb,
		Fila:         fila,
		Columna:      columna,
	}
	return nC
}

func (c Cast) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	/*******************************VARIABLES 3D*****************************/
	var obj3d, obj3dValor Ast.O3D
	var codigo3d string

	/************************************************************************/

	var nuevoTipo Ast.TipoDato
	var nuevoValor Ast.TipoRetornado
	//Primero conseguir el valor a convertir
	valor := c.Valor.GetValue(scope)
	obj3dValor = valor.Valor.(Ast.O3D)
	valor = obj3dValor.Valor
	codigo3d += obj3dValor.Codigo

	if valor.Tipo == Ast.ERROR {
		return valor
	}
	//Verificar que el tipo se puede convertir implicitamente
	if !CanConvert(valor.Tipo, c.TipoObjetivo) {
		//Error el tipo no se puede convertir implicitamente
		msg := "Semantic error, can't implicitly convert " + Ast.ValorTipoDato[valor.Tipo] +
			" type to " + Ast.ValorTipoDato[c.TipoObjetivo] + " type." +
			" -- Line: " + strconv.Itoa(c.Fila) +
			" Column: " + strconv.Itoa(c.Columna)
		nError := errores.NewError(c.Fila, c.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}

	//Fuera de errores, convertir
	nuevoTipo = conversion[valor.Tipo][c.TipoObjetivo]
	if nuevoTipo == Ast.NULL {
		//Erro no se pueden convertir esos tipos
		msg := "Semantic error, can't implicitly convert " + Ast.ValorTipoDato[valor.Tipo] +
			" type to " + Ast.ValorTipoDato[c.TipoObjetivo] + " type." +
			" -- Line: " + strconv.Itoa(c.Fila) +
			" Column: " + strconv.Itoa(c.Columna)
		nError := errores.NewError(c.Fila, c.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}
	nuevoValor = c.convertir(nuevoTipo, valor, scope)

	obj3d.Valor = nuevoValor
	obj3d.Codigo = codigo3d
	//obj3d.Referencia = Primitivo_To_String(nuevoValor.Valor, nuevoValor.Tipo)
	obj3d.Referencia = obj3dValor.Referencia
	return Ast.TipoRetornado{
		Tipo:  nuevoValor.Tipo,
		Valor: obj3d,
	}
}

func (c Cast) convertir(nuevoTipo Ast.TipoDato, valor Ast.TipoRetornado, scope *Ast.Scope) Ast.TipoRetornado {
	var nuevoValor Ast.TipoRetornado
	switch nuevoTipo {
	case Ast.USIZE:
		switch valor.Tipo {
		////////////////////////////////////////
		//Se Me olvido agregar el caso del usize
		case Ast.I64, Ast.USIZE:
			nuevoValor.Valor = valor.Valor.(int)
		}
	case Ast.I64:
		// Esperamos un i64, f64, bool o char
		switch valor.Tipo {
		case Ast.I64, Ast.USIZE:
			nuevoValor.Valor = valor.Valor.(int)
		case Ast.F64:
			nuevoValor.Valor = int(valor.Valor.(float64))
		case Ast.BOOLEAN:
			if valor.Valor.(bool) {
				nuevoValor.Valor = 1
			} else {
				nuevoValor.Valor = 0
			}
		case Ast.CHAR:
			chars := []rune(valor.Valor.(string))
			nuevoValor.Valor = int(chars[0])
		}
	case Ast.F64:
		// Esperamos un i64 o un f64
		switch valor.Tipo {
		case Ast.I64:
			nuevoValor.Valor = float64(valor.Valor.(int))
		case Ast.F64:
			nuevoValor.Valor = valor.Valor.(float64)
		}
	case Ast.BOOLEAN:
		// Esperamos un booleano o un i64
		switch valor.Tipo {
		case Ast.I64:
			if valor.Valor.(bool) {
				nuevoValor.Valor = int(1)
			} else {
				nuevoValor.Valor = int(0)
			}
		case Ast.BOOLEAN:
			nuevoValor.Valor = valor.Valor.(bool)
		}

	case Ast.CHAR:
		// Esperamos un int o un char
		switch valor.Tipo {
		case Ast.I64:
			if valor.Valor.(int) > 255 || valor.Valor.(int) < 0 {
				//Error, ASCII fuera de rango
				//literal out of range to convert to char
				msg := "Semantic error, literal out of range to convert to char " +
					" -- Line: " + strconv.Itoa(c.Fila) +
					" Column: " + strconv.Itoa(c.Columna)
				nError := errores.NewError(c.Fila, c.Columna, msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				nError.Ambito = scope.GetTipoScope()
				scope.Errores.Add(nError)
				scope.Consola += msg + "\n"
				return Ast.TipoRetornado{
					Tipo:  Ast.ERROR,
					Valor: nError,
				}
			}
			nuevoValor.Valor = strconv.Itoa((valor.Valor.(int)))
		case Ast.CHAR:
			nuevoValor.Valor = valor.Valor.(string)
		}
	}
	nuevoValor.Tipo = nuevoTipo
	return nuevoValor
}

// El primero es el valor inicial y el segundo es el tipo objetivo
var conversion = [8][10]Ast.TipoDato{
	{Ast.I64, Ast.F64, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.USIZE, Ast.CHAR, Ast.NULL, Ast.NULL},
	{Ast.I64, Ast.F64, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL},
	{Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL},
	{Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL},
	{Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL},
	{Ast.I64, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.BOOLEAN, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL},
	{Ast.I64, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.USIZE, Ast.NULL, Ast.NULL, Ast.NULL},
	{Ast.I64, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.NULL, Ast.CHAR, Ast.NULL, Ast.NULL},
}

/* TIpo
0	I64
1	F64
2	STRING_OWNED
3	STRING
4	STR
5	BOOLEAN
6	CHAR
7	ARRAY
8	VECTOR
9	STRUCT
*/

func CanConvert(tipo Ast.TipoDato, tipoObjetivo Ast.TipoDato) bool {
	if tipo > 7 || tipoObjetivo > 7 {
		return false
	}
	if tipo == 2 || tipo == 3 || tipo == 4 ||
		tipoObjetivo == 2 || tipoObjetivo == 3 || tipoObjetivo == 4 {
		return false
	}
	return true
}

func (c Cast) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, c.Tipo
}

func (c Cast) GetFila() int {
	return c.Fila
}
func (c Cast) GetColumna() int {
	return c.Columna
}
