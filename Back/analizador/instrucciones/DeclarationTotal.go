package instrucciones

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"
)

type DeclaracionTotal struct {
	Id            string
	Mutable       bool
	Publico       bool
	Tipo          Ast.TipoRetornado
	Valor         interface{}
	Fila          int
	Columna       int
	Stack         bool
	ScopeOriginal *Ast.Scope
}

func NewDeclaracionTotal(id string, valor interface{}, tipo Ast.TipoRetornado, mutable, publico bool,
	fila int, columna int) DeclaracionTotal {
	nd := DeclaracionTotal{
		Id:      id,
		Mutable: mutable,
		Publico: publico,
		Valor:   valor,
		Tipo:    tipo,
		Fila:    fila,
		Columna: columna,
		Stack:   true,
	}
	return nd
}

func (d DeclaracionTotal) Run(scope *Ast.Scope) interface{} {
	//Verificar si es un tipo especial
	var esEspecial bool = false
	/*****************************VARIABLES 3D*****************************/
	var codigo3d string = ""
	var obj3DValor, obj3D, obj3dTemp Ast.O3D
	var temp string
	var direccion int
	var scopeAnterior string
	/*********************************************************************/

	//Verificar que el id no exista

	existe := scope.Exist_actual(d.Id)

	if existe {
		//Ya existe y generar error semántico
		return errores.GenerarError(12, d, d, d.Id, "", "", scope)
	}

	//Verificar que no es un if expresion
	_, tipoIn := d.Valor.(Ast.Abstracto).GetTipo()

	//Verificar que sea un primitivo i64 y la declaración sea usize

	var preValor interface{}
	if tipoIn == Ast.IF_EXPRESION || tipoIn == Ast.MATCH_EXPRESION || tipoIn == Ast.LOOP_EXPRESION {
		preValor = d.Valor.(Ast.Instruccion).Run(scope)
	} else if tipoIn == Ast.VALOR {
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

	obj3DValor = preValor.(Ast.TipoRetornado).Valor.(Ast.O3D)
	valor := obj3DValor.Valor

	//Cambiar valor de i64 a usize si la declaración es usize y el valor que viene es un i64
	if d.Tipo.Tipo == Ast.USIZE && tipoIn == Ast.I64 {
		valor.Tipo = Ast.USIZE
	}

	//Revisar si el retorno es un error
	if valor.Tipo == Ast.ERROR {
		/////////////////////////////ERROR/////////////////////////////////////
		return valor
	}

	//comparar los tipos
	if !EsTipoEspecial(valor.Tipo) {
		//No es struct,vector,array, entonces comparar los tipos normalmente
		if d.Tipo.Tipo != valor.Tipo {
			//Error de tipos
			/////////////////////////////ERROR/////////////////////////////////////
			return errores.GenerarError(13, d, d, d.Id, expresiones.Tipo_String(d.Tipo),
				Ast.ValorTipoDato[valor.Tipo], scope)
		}

	} else {
		//Primero verificar si el tipo de la declaración no es un acceso a modulo

		if d.Tipo.Tipo == Ast.ACCESO_MODULO {
			//Ejecutar el acceso y cambiar el tipo de la declaración
			nTipo := GetTipoEstructura(d.Tipo, scope)
			errors := ErrorEnTipo(nTipo)
			if errors.Tipo == Ast.ERROR {
				/*Generar el error*/
				/////////////////////////////ERROR/////////////////////////////////////
				return errores.GenerarError(14, d, d, "", "", "", scope)
			}
			//De lo contrario actualizar el tipo de la declaracion
			d.Tipo = nTipo
		}

		//Si es un tipo especial, entonces comparar los tipos en profundidad
		tipoEspecial := GetTipoEspecial(valor.Tipo, valor.Valor, scope)
		if tipoEspecial.Tipo == Ast.ERROR {
			//Erro de tipos
			/////////////////////////////ERROR/////////////////////////////////////
			return errores.GenerarError(14, d, d, "", "", "", scope)
		}

		if !expresiones.CompararTipos(d.Tipo, tipoEspecial) {
			//Error, los tipos no son correctos
			/////////////////////////////ERROR/////////////////////////////////////
			return errores.GenerarError(14, d, d, "", "", "", scope)
		}
		esEspecial = true
	}

	//Todo bien crear y agregar el símbolo

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

	//Desde aquí trabajar el C3D

	if d.Stack {
		//Agregar el código anterior
		if tipoIn != Ast.VALOR {
			codigo3d += obj3DValor.Codigo
		}
		if valor.Tipo == Ast.BOOLEAN {
			if obj3DValor.Lt != "" {
				salto := Ast.GetLabel()
				codigo3d += "/*************************CAMBIO VALOR BOOLEANO*/ \n"
				codigo3d += obj3DValor.Lt + ": \n"
				codigo3d += obj3DValor.Referencia + " = 1;\n"
				codigo3d += "goto " + salto + ";\n"
				codigo3d += obj3DValor.Lf + ": \n"
				codigo3d += obj3DValor.Referencia + " = 0;\n"
				codigo3d += salto + ":\n"
				codigo3d += "/***********************************************/ \n"
			}
		}

		codigo3d += "/***********************DECLARACIÓN DE VARIABLE*/ \n"
		/*Get el nuevo temporal*/
		temp = Ast.GetTemp()
		/*Get la dirección donde se guardara */

		direccion = scope.ContadorDeclaracion
		scope.ContadorDeclaracion++
		//Conseguir la dirección del stack donde se va a guardar la nueva variable
		codigo3d += temp + " = " + " P + " + strconv.Itoa(direccion) + ";\n"
		//Se guarda la nueva variable en el stack
		codigo3d += "stack[(int)" + temp + "] = " + obj3DValor.Referencia + ";\n"
		//Aumentar el SP del ambito

		codigo3d += "/***********************************************/\n"
		//Agregar la dirección al símbolo
		nSimbolo.Direccion = direccion
		nSimbolo.TipoDireccion = Ast.STACK
		//Actualizar el obj3d
		obj3D.Codigo = codigo3d
		obj3D.Valor = valor

	} else {
		//Agregar el código anterior
		if tipoIn != Ast.VALOR {
			codigo3d += obj3DValor.Codigo
		}

		if valor.Tipo == Ast.BOOLEAN {
			if obj3DValor.Lt != "" {
				salto := Ast.GetLabel()
				codigo3d += "/*************************CAMBIO VALOR BOOLEANO*/ \n"
				codigo3d += obj3DValor.Lt + ": \n"
				codigo3d += obj3DValor.Referencia + " = 1;\n"
				codigo3d += "goto " + salto + ";\n"
				codigo3d += obj3DValor.Lf + ": \n"
				codigo3d += obj3DValor.Referencia + " = 0;\n"
				codigo3d += salto + ":\n"
				codigo3d += "/***********************************************/ \n"
			}
		}

		codigo3d += "/***********************DECLARACIÓN DE VARIABLE*/ \n"

		/*Get el nuevo temporal*/
		temp = Ast.GetTemp()
		/*Get la dirección donde se guardara */
		direccion = Ast.GetH()
		//Conseguir la dirección del stack donde se va a guardar la nueva variable
		codigo3d += temp + " = " + " H;\n"
		//Se guarda la nueva variable en el stack
		codigo3d += "heap[(int)" + temp + "] = " + obj3DValor.Referencia + ";\n"
		//Aumentar H
		codigo3d += "H = H + 1;\n"
		codigo3d += "/***********************************************/\n"
		//Aumentar el scope
		scope.ContadorDeclaracion++
		//Agregar la dirección al símbolo
		nSimbolo.Direccion = direccion
		nSimbolo.TipoDireccion = Ast.HEAP
		//Actualizar el obj3d
		obj3D.Codigo = codigo3d
		obj3D.Valor = valor
	}

	scope.Add(nSimbolo)

	return Ast.TipoRetornado{
		Valor: obj3D,
		Tipo:  Ast.DECLARACION,
	}

}

func (op DeclaracionTotal) GetFila() int {
	return op.Fila
}
func (op DeclaracionTotal) GetColumna() int {
	return op.Columna
}

func (d DeclaracionTotal) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, Ast.DECLARACION
}

func (d DeclaracionTotal) SetHeap() interface{} {
	Dcopia := d
	Dcopia.Stack = false
	return Dcopia
}
