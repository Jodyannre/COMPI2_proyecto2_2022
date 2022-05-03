package instrucciones

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"
)

type DeclaracionNoRef struct {
	Id            string
	Mutable       bool
	Publico       bool
	Tipo          Ast.TipoRetornado
	Valor         interface{}
	Fila          int
	Columna       int
	ScopeOriginal *Ast.Scope
}

func NewDeclaracionNoRef(id string, valor interface{}, tipo Ast.TipoRetornado, mutable, publico bool,
	fila int, columna int) DeclaracionNoRef {
	nd := DeclaracionNoRef{
		Id:      id,
		Mutable: mutable,
		Publico: publico,
		Valor:   valor,
		Tipo:    tipo,
		Fila:    fila,
		Columna: columna,
	}
	return nd
}

func (d DeclaracionNoRef) Run(scope *Ast.Scope) interface{} {
	//Verificar si es un tipo especial
	var esEspecial bool = false
	/*****************************VARIABLES 3D*****************************/
	var codigo3d string = ""
	var posicionAguardar string = ""
	var obj3DValor, obj3d, obj3dTemp, obj3dClone Ast.O3D
	var scopeAnterior string
	/*********************************************************************/
	//Verificar que el id no exista

	existe := scope.Exist_actual(d.Id)

	if existe {
		//Ya existe y generar error semántico
		////////////////////////////ERROR//////////////////////////////////
		return errores.GenerarError(12, d, d, d.Id,
			"",
			"",
			scope)
		//////////////////////////////////////////////////////////////////
	}

	//Verificar que no es un if expresion
	_, tipoIn := d.Valor.(Ast.Abstracto).GetTipo()

	//Verificar que sea un primitivo i64 y la declaración sea usize

	var preValor interface{}
	if tipoIn == Ast.IF_EXPRESION || tipoIn == Ast.MATCH_EXPRESION || tipoIn == Ast.LOOP_EXPRESION {
		preValor = d.Valor.(Ast.Instruccion).Run(scope)
	} else if tipoIn == Ast.VALOR {
		//preValor = d.Valor.(Ast.Expresion).GetValue(d.ScopeOriginal)
		scopeAnterior = Ast.GetTemp()
		/*********************SCOPE SIMULADO****************************/
		codigo3d += scopeAnterior + " = P; //Guardar el scope anterior \n"
		codigo3d += "P = " + strconv.Itoa(d.ScopeOriginal.Posicion) + "; //Scope de donde proviene el valor\n"
		/***************************************************************/
		preValor = d.Valor.(Ast.Expresion).GetValue(d.ScopeOriginal)
		obj3dTemp = preValor.(Ast.TipoRetornado).Valor.(Ast.O3D)
		codigo3d += obj3dTemp.Codigo
		/*********************RETORNO SCOPE ANTERIOR********************/
		codigo3d += "P = " + scopeAnterior + "; //Retornar al scope anterior \n"
		/***************************************************************/
	} else {
		preValor = d.Valor.(Ast.Expresion).GetValue(scope)
	}
	//valor := preValor.(Ast.TipoRetornado)
	obj3DValor = preValor.(Ast.TipoRetornado).Valor.(Ast.O3D)
	valor := obj3DValor.Valor

	//Cambiar valor de i64 a usize si la declaración es usize y el valor que viene es un i64
	if d.Tipo.Tipo == Ast.USIZE && tipoIn == Ast.I64 {
		valor.Tipo = Ast.USIZE
	}

	//Revisar si el retorno es un error
	if valor.Tipo == Ast.ERROR {
		return valor
	}

	//comparar los tipos
	if !EsTipoEspecial(valor.Tipo) {
		//No es struct,vector,array, entonces comparar los tipos normalmente
		if d.Tipo.Tipo != valor.Tipo {
			//Error de tipos
			////////////////////////////ERROR//////////////////////////////////
			return errores.GenerarError(13, d, d, d.Id,
				expresiones.Tipo_String(d.Tipo),
				Ast.ValorTipoDato[valor.Tipo],
				scope)
			//////////////////////////////////////////////////////////////////

		}

	} else {
		//Primero verificar si el tipo de la declaración no es un acceso a modulo

		if d.Tipo.Tipo == Ast.ACCESO_MODULO {
			//Ejecutar el acceso y cambiar el tipo de la declaración
			nTipo := GetTipoEstructura(d.Tipo, scope)
			errors := ErrorEnTipo(nTipo)
			if errors.Tipo == Ast.ERROR {
				////////////////////////////ERROR//////////////////////////////////
				return errores.GenerarError(14, d, d, "",
					"",
					"",
					scope)
				//////////////////////////////////////////////////////////////////
			}
			//De lo contrario actualizar el tipo de la declaracion
			d.Tipo = nTipo
		}

		//Si es un tipo especial, entonces comparar los tipos en profundidad
		tipoEspecial := GetTipoEspecial(valor.Tipo, valor.Valor, scope)
		if tipoEspecial.Tipo == Ast.ERROR {
			//Erro de tipos
			////////////////////////////ERROR//////////////////////////////////
			return errores.GenerarError(14, d, d, "",
				"",
				"",
				scope)
			//////////////////////////////////////////////////////////////////
		}

		if !expresiones.CompararTipos(d.Tipo, tipoEspecial) {
			//Error, los tipos no son correctos
			////////////////////////////ERROR//////////////////////////////////
			return errores.GenerarError(14, d, d, "",
				"",
				"",
				scope)
			//////////////////////////////////////////////////////////////////
		}
		esEspecial = true
	}

	//Todo bien crear y agregar el símbolo
	//Clonar el valor
	preValorRef := valor.Valor.(Ast.Clones).SetReferencia(obj3DValor.Referencia)
	valorRef := preValorRef.(Ast.Clones).Clonar(scope)
	obj3dClone = valorRef.(Ast.TipoRetornado).Valor.(Ast.O3D)
	codigo3d += obj3dClone.Codigo
	valorRef = obj3dClone.Valor
	valor = valorRef.(Ast.TipoRetornado)
	/*
		valor = Ast.TipoRetornado{
			Tipo:  valor.Tipo,
			Valor: valorRef,
		}
	*/

	/*******************************CODIGO 3D*************************************/
	posicionAguardar = Ast.GetTemp()
	codigo3d += posicionAguardar + " = P + " + strconv.Itoa(scope.ContadorDeclaracion) + "; //Get pos para guardar\n"
	codigo3d += "stack[(int)" + posicionAguardar + "] = " + obj3dClone.Referencia + "; //Guardar nuevo valor\n"

	nSimbolo := Ast.Simbolo{
		Identificador: d.Id,
		Valor:         valor,
		Fila:          d.Fila,
		Columna:       d.Columna,
		Tipo:          valor.Tipo,
		Mutable:       d.Mutable,
		Publico:       d.Publico,
		Entorno:       scope,
	}
	//Agregar el tipo especial
	if esEspecial {
		nSimbolo.TipoEspecial = d.Tipo
	}

	//Agregar la direccion
	nSimbolo.Direccion = scope.ContadorDeclaracion
	scope.ContadorDeclaracion++
	nSimbolo.TipoDireccion = Ast.STACK

	//Verificar si es array, vector o struct, para clonarlos

	scope.Add(nSimbolo)

	obj3d.Valor = Ast.TipoRetornado{
		Valor: true,
		Tipo:  Ast.EJECUTADO,
	}
	obj3d.Codigo = codigo3d

	return Ast.TipoRetornado{
		Valor: obj3d,
		Tipo:  Ast.EJECUTADO,
	}

}

func (op DeclaracionNoRef) GetFila() int {
	return op.Fila
}
func (op DeclaracionNoRef) GetColumna() int {
	return op.Columna
}

func (d DeclaracionNoRef) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, Ast.DECLARACION
}
