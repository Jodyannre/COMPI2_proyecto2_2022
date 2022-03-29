package simbolos

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"
)

type AsignacionStruct struct {
	AccesoStruct interface{}
	Valor        interface{}
	Fila         int
	Columna      int
	Tipo         Ast.TipoDato
}

func NewAsignacionStruct(acceso, valor interface{}, fila, columna int) AsignacionStruct {
	nA := AsignacionStruct{
		AccesoStruct: acceso,
		Valor:        valor,
		Fila:         fila,
		Columna:      columna,
		Tipo:         Ast.ASIGNACION_STRUCT,
	}
	return nA
}

func (a AsignacionStruct) Run(scope *Ast.Scope) interface{} {
	//Verificar que el acceso este correcto
	accesoCorreto := a.AccesoStruct.(Ast.Expresion).GetValue(scope)
	var nombreStruct string
	var nombreAtributo string
	var structActual StructInstancia
	var simboloAtributo Ast.Simbolo
	var simboloStruct Ast.Simbolo
	var valor Ast.TipoRetornado

	if accesoCorreto.Tipo == Ast.ERROR {
		return accesoCorreto
	}

	//Recuperar los nombres
	nombreStruct = a.AccesoStruct.(AccesoStruct).NombreStruct.(expresiones.Identificador).Valor
	nombreAtributo = a.AccesoStruct.(AccesoStruct).NombreAtributo.(expresiones.Identificador).Valor

	//Recuperar el struct
	structActual = a.AccesoStruct.(AccesoStruct).NombreStruct.(Ast.Expresion).GetValue(scope).Valor.(StructInstancia)

	//Recuperar el símbolo del campo del struct
	simboloAtributo = structActual.Entorno.GetSimbolo(nombreAtributo)

	//Get el valor a asignar
	valor = a.Valor.(Ast.Expresion).GetValue(scope)

	//Verificar posible error en el valor
	if valor.Tipo == Ast.ERROR {
		return valor
	}

	//Recuperar el símbolo del struct
	simboloStruct = scope.GetSimbolo(nombreStruct)

	//Verificar que la instancia del struct sea mutable

	if !simboloStruct.Mutable {
		//Error, la instancia no es mutable
		fila := a.Fila
		columna := a.Columna
		msg := "Semantic error, can't modify a non-mutable " + Ast.ValorTipoDato[int(simboloStruct.Tipo)] +
			" type. -- Line: " + strconv.Itoa(fila) +
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

	//Verificar que los tipos sean correctos

	if simboloAtributo.Tipo != valor.Tipo {
		//Error, los tipos no son correctos
		fila := a.Valor.(Ast.Abstracto).GetFila()
		columna := a.Valor.(Ast.Abstracto).GetColumna()
		msg := "Semantic error, can't assign " + Ast.ValorTipoDato[valor.Tipo] +
			" to " + Ast.ValorTipoDato[simboloAtributo.Tipo] + " field." +
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
	//Los simbolos son del mismo tipo, entonces asignar el nuevo valor al campo del struct
	simboloAtributo.Valor = valor
	structActual.Entorno.UpdateSimbolo(nombreAtributo, simboloAtributo)
	return Ast.TipoRetornado{
		Tipo:  Ast.EJECUTADO,
		Valor: true,
	}
}

func (v AsignacionStruct) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, v.Tipo
}

func (v AsignacionStruct) GetFila() int {
	return v.Fila
}
func (v AsignacionStruct) GetColumna() int {
	return v.Columna
}
