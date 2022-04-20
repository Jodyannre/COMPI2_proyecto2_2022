package fn_primitivas

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
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
	/***************************************VARIABLES 3D*****************************************/
	var obj3d, obj3dValor, obj3dConversion Ast.O3D
	var codigo3d string

	/********************************************************************************************/

	valorSalida := Ast.TipoRetornado{}

	//Primero conseguir el valor a convertir
	valor := t.Valor.(Ast.Expresion).GetValue(scope)
	obj3dValor = valor.Valor.(Ast.O3D)
	valor = obj3dValor.Valor

	codigo3d += "/******************************METODO TO STRING*/\n"
	codigo3d += obj3dValor.Codigo

	//Verificar que el valor no sea un error
	if valor.Tipo == Ast.ERROR {
		return valor
	}

	//Verificar el tipo del valor y convertir
	if valor.Tipo > 7 {
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

	//Conseguir el cod3d del nuevo valor string
	obj3dConversion = GetCod3DtoString(Ast.STRING, valorSalida.Valor, scope)
	codigo3d += obj3dConversion.Codigo

	valorSalida.Tipo = Ast.STRING

	obj3d.Valor = valorSalida
	obj3d.Codigo = codigo3d
	obj3d.Referencia = obj3dConversion.Referencia

	codigo3d += "/***********************************************/\n"
	return Ast.TipoRetornado{
		Tipo:  Ast.STRING,
		Valor: obj3d,
	}
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

func GetCod3DtoString(tipo Ast.TipoDato, valor interface{}, scope *Ast.Scope) Ast.O3D {
	var obj3dValor Ast.O3D
	var resultado Ast.TipoRetornado
	primitivo := expresiones.NewPrimitivo(valor, tipo, 0, 0)
	resultado = primitivo.GetValue(scope)
	obj3dValor = resultado.Valor.(Ast.O3D)
	return obj3dValor
}
