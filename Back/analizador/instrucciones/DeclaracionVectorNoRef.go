package instrucciones

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"
)

type DeclaracionVectorNoRef struct {
	Id            string
	Tipo          Ast.TipoDato
	TipoVector    Ast.TipoRetornado
	Mutable       bool
	Publico       bool
	Valor         interface{}
	Fila          int
	Columna       int
	ScopeOriginal *Ast.Scope
	Stack         bool
}

func NewDeclaracionVectorNoRef(id string, tipoVector Ast.TipoRetornado, valor interface{}, mutable, publico bool,
	fila int, columna int) DeclaracionVectorNoRef {
	nd := DeclaracionVectorNoRef{
		Id:         id,
		Tipo:       Ast.DECLARACION,
		TipoVector: tipoVector,
		Mutable:    mutable,
		Publico:    publico,
		Valor:      valor,
		Fila:       fila,
		Columna:    columna,
		Stack:      true,
	}
	return nd
}

func (d DeclaracionVectorNoRef) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, Ast.DECLARACION
}

func (d DeclaracionVectorNoRef) Run(scope *Ast.Scope) interface{} {
	//Verificar que no exista
	var existe bool
	var valor Ast.TipoRetornado
	esIndefinido := false
	/**********VARIABLES 3D***************/
	var codigo3d, scopeAnterior string
	var obj3d, obj3dValor, obj3dClone, obj3dTemp Ast.O3D
	/*************************************/
	_, tipoIn := d.Valor.(Ast.Abstracto).GetTipo()
	if tipoIn == Ast.VALOR {
		scopeAnterior = Ast.GetTemp()
		/*********************SCOPE SIMULADO****************************/
		codigo3d += scopeAnterior + " = P; //Guardar el scope anterior \n"
		codigo3d += "P = " + strconv.Itoa(d.ScopeOriginal.Posicion) + "; //Scope de donde proviene el valor\n"
		/***************************************************************/
		existe = d.ScopeOriginal.Exist_actual(d.Id)
		valor = d.Valor.(Ast.Expresion).GetValue(d.ScopeOriginal)
		obj3dTemp = valor.Valor.(Ast.O3D)
		valor = obj3dTemp.Valor
		codigo3d += obj3dTemp.Codigo
		/*********************RETORNO SCOPE ANTERIOR********************/
		codigo3d += "P = " + scopeAnterior + "; //Retornar al scope anterior \n"
		/***************************************************************/
	} else {
		existe = scope.Exist_actual(d.Id)
		valor = d.Valor.(Ast.Expresion).GetValue(scope)
		obj3dValor = valor.Valor.(Ast.O3D)
		valor = obj3dValor.Valor
	}
	tipoIn = valor.Tipo
	//Verificar error en el valor
	if valor.Tipo == Ast.ERROR {
		return valor
	}

	//Si es diferente de vector error
	if tipoIn != Ast.VECTOR {
		msg := "Semantic error, can't initialize a" + Ast.ValorTipoDato[d.Tipo] + "with " + Ast.ValorTipoDato[tipoIn] + " value." +
			" -- Line:" + strconv.Itoa(d.Fila) + " Column: " + strconv.Itoa(d.Columna)
		nError := errores.NewError(d.Fila, d.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}

	//Verificar si ya existe
	if existe {
		msg := "Semantic error, the element \"" + d.Id + "\" already exist in this scope." +
			" -- Line:" + strconv.Itoa(d.Fila) + " Column: " + strconv.Itoa(d.Columna)
		nError := errores.NewError(d.Fila, d.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}
	//Verificar que el tipo del vector no sea un acceso a modulo

	if d.TipoVector.Tipo == Ast.ACCESO_MODULO || EsTipoEspecial(d.TipoVector.Tipo) {
		//Traer el tipo y cambiar el tipo de la declaración
		nTipo := GetTipoEstructura(d.TipoVector, scope)
		if nTipo.Tipo == Ast.ERROR {
			return nTipo
		}
		if nTipo.Tipo == Ast.STRUCT_TEMPLATE {
			nTipo.Tipo = Ast.STRUCT
		}
		d.TipoVector = nTipo
	}

	//Verificar que los tipos de los vectores sean correctos
	if !expresiones.CompararTipos(d.TipoVector, valor.Valor.(expresiones.Vector).TipoVector) {
		if valor.Valor.(expresiones.Vector).TipoVector.Tipo == Ast.INDEFINIDO {
			//Es uno vacio y no hay error, modificar el tipo
			esIndefinido = true
		} else {
			msg := "Semantic error, can't initialize a Vec<" + expresiones.Tipo_String(d.TipoVector) +
				"> with Vec<" + expresiones.Tipo_String(valor.Valor.(expresiones.Vector).TipoVector) + "> value." +
				" -- Line:" + strconv.Itoa(d.Fila) + " Column: " + strconv.Itoa(d.Columna)
			nError := errores.NewError(d.Fila, d.Columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			nError.Ambito = scope.GetTipoScope()
			scope.Errores.Add(nError)
			scope.Consola += msg + "\n"
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}
		}
	}

	//Crear el símbolo y agregarlo al scope
	if esIndefinido {
		temp := Ast.GetTemp()
		vectorTemp := valor.Valor.(expresiones.Vector)
		vectorTemp.Referencia = obj3dTemp.Referencia
		valor.Valor = vectorTemp
		//Clonar la lista para evitar la referencia
		nmVector := valor.Valor.(Ast.Clones).Clonar(scope)
		obj3dClone = nmVector.(Ast.TipoRetornado).Valor.(Ast.O3D)
		nVector := obj3dClone.Valor.Valor.(expresiones.Vector)
		codigo3d += obj3dClone.Codigo
		//Actualizar la mutabilidad de la instancia
		nVector.TipoVector = d.TipoVector
		nVector.Mutable = d.Mutable
		nSimbolo := Ast.Simbolo{
			Identificador: d.Id,
			Valor:         Ast.TipoRetornado{Tipo: Ast.VECTOR, Valor: nVector},
			Fila:          d.Fila,
			Columna:       d.Columna,
			Tipo:          nVector.Tipo,
			Mutable:       d.Mutable,
			Publico:       d.Publico,
			Entorno:       scope,
		}
		//codigo3d += obj3dClone.Codigo
		codigo3d += "/*************************DECLARACION DE VECTOR*/\n"
		if d.Stack {
			codigo3d += temp + " = P + " + strconv.Itoa(scope.Size) + ";\n"
			nSimbolo.Direccion = scope.Size
			nSimbolo.TipoDireccion = Ast.STACK
			scope.Size++
			codigo3d += "stack[(int)" + temp + "] = " + obj3dClone.Referencia + ";\n"
		} else {
			codigo3d += temp + " = P + " + strconv.Itoa(scope.Size) + ";\n"
			nSimbolo.Direccion = scope.Size
			nSimbolo.TipoDireccion = Ast.HEAP
			scope.Size++
			codigo3d += "heap[(int)" + temp + "] = " + obj3dClone.Referencia + ";\n"
		}
		codigo3d += "/***********************************************/\n"
		scope.Add(nSimbolo)
	} else {
		//Temporales
		temp := Ast.GetTemp()
		vectorTemp := valor.Valor.(expresiones.Vector)
		vectorTemp.Referencia = obj3dTemp.Referencia
		valor.Valor = vectorTemp
		//Clonar la lista para evitar la referencia
		nmVector := valor.Valor.(Ast.Clones).Clonar(scope)
		obj3dClone = nmVector.(Ast.TipoRetornado).Valor.(Ast.O3D)
		nVector := obj3dClone.Valor.Valor.(expresiones.Vector)
		codigo3d += obj3dClone.Codigo

		//Actualizar la mutabilidad de la instancia
		nVector.Mutable = d.Mutable
		//Guardar el valor en el obj3d
		obj3d.Valor = Ast.TipoRetornado{Tipo: Ast.VECTOR, Valor: nVector}

		nSimbolo := Ast.Simbolo{
			Identificador: d.Id,
			Valor:         Ast.TipoRetornado{Tipo: Ast.VECTOR, Valor: nVector},
			Fila:          d.Fila,
			Columna:       d.Columna,
			Tipo:          valor.Tipo,
			Mutable:       d.Mutable,
			Publico:       d.Publico,
			Entorno:       scope,
		}

		/*Codigo 3d*/
		//codigo3d += obj3dValor.Codigo
		codigo3d += "/*************************DECLARACION DE VECTOR*/\n"
		if d.Stack {
			codigo3d += temp + " = P + " + strconv.Itoa(scope.Size) + ";\n"
			nSimbolo.Direccion = scope.Size
			nSimbolo.TipoDireccion = Ast.STACK
			scope.Size++
			codigo3d += "stack[(int)" + temp + "] = " + obj3dClone.Referencia + ";\n"
		} else {
			codigo3d += temp + " = P + " + strconv.Itoa(scope.Size) + ";\n"
			nSimbolo.Direccion = scope.Size
			nSimbolo.TipoDireccion = Ast.HEAP
			scope.Size++
			codigo3d += "heap[(int)" + temp + "] = " + obj3dClone.Referencia + ";\n"
		}
		codigo3d += "/***********************************************/\n"
		scope.Add(nSimbolo)
	}
	obj3d.Codigo = codigo3d
	return Ast.TipoRetornado{Valor: obj3d, Tipo: Ast.DECLARACION}
}

func (op DeclaracionVectorNoRef) GetFila() int {
	return op.Fila
}
func (op DeclaracionVectorNoRef) GetColumna() int {
	return op.Columna
}

func ClonarVector3D() (string, string) {

	return "", ""
}
