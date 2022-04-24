package simbolos

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"
)

type Valor struct {
	Valor      Ast.Expresion
	Tipo       Ast.TipoDato
	Fila       int
	Columna    int
	Referencia bool
	Mutable    bool
}

func NewValor(valor Ast.Expresion, tipo Ast.TipoDato, referencia bool, mutable bool, fila, columna int) Valor {
	nV := Valor{
		Tipo:       tipo,
		Valor:      valor,
		Referencia: referencia,
		Mutable:    mutable,
		Fila:       fila,
		Columna:    columna,
	}
	return nV
}

func (v Valor) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	/***********************VARIABLES 3D***********************/
	var obj3dValor, obj3d Ast.O3D
	var codigo3d string
	var referencia string
	var abstracto interface{}
	/**********************************************************/

	//Verificar que los valores de referencia sean correctos
	valor := v.Valor.GetValue(scope)
	obj3dValor = valor.Valor.(Ast.O3D)
	valor = obj3dValor.Valor
	codigo3d += obj3dValor.Codigo
	referencia = obj3dValor.Referencia
	obj3d.Valor = valor
	obj3d.Codigo = codigo3d
	obj3d.Referencia = referencia

	/**********PARA REFERENCIAS*****************/
	abstracto = v.Valor
	_, tipoParticular := abstracto.(Ast.Abstracto).GetTipo()
	if tipoParticular == Ast.IDENTIFICADOR && valor.Tipo == Ast.VECTOR {
		obj3d.EsReferencia = v.Valor.(expresiones.Identificador).Valor
	}
	/*******************************************/

	if !EsPosibleReferencia(valor.Tipo) && v.Referencia {
		//Error, ese tipo de valor no se puede enviar como referencia
		fila := v.Fila
		columna := v.Columna
		msg := "Semantic error, " + Ast.ValorTipoDato[valor.Tipo] + " can't be send as a reference." +
			" -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
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
	return Ast.TipoRetornado{
		Valor: obj3d,
		Tipo:  Ast.VALOR,
	}
}
func (p Valor) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, p.Tipo
}

func (p Valor) GetFila() int {
	return p.Fila
}
func (p Valor) GetColumna() int {
	return p.Columna
}
