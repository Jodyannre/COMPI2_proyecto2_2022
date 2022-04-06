package instrucciones

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"
)

type DeclaracionVector struct {
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

func NewDeclaracionVector(id string, tipoVector Ast.TipoRetornado, valor interface{}, mutable, publico bool,
	fila int, columna int) DeclaracionVector {
	nd := DeclaracionVector{
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

func (d DeclaracionVector) Run(scope *Ast.Scope) interface{} {
	//Verificar que no exista
	var existe bool
	var valor Ast.TipoRetornado
	var codigo3d string
	var obj3d, obj3dValor Ast.O3D
	esIndefinido := false
	_, tipoIn := d.Valor.(Ast.Abstracto).GetTipo()
	if tipoIn == Ast.VALOR {
		existe = d.ScopeOriginal.Exist_actual(d.Id)
		valor = d.Valor.(Ast.Expresion).GetValue(d.ScopeOriginal)
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
		//Actualizar la mutabilidad de la instancia
		nVector := valor.Valor.(expresiones.Vector)
		nVector.TipoVector = d.TipoVector
		nVector.Mutable = d.Mutable
		//Agregar el valor al obj3d
		obj3d.Valor = Ast.TipoRetornado{Tipo: Ast.VECTOR, Valor: nVector}
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
		/*Codigo 3d*/
		/*Agregar el código de la creación del vector*/
		codigo3d += obj3dValor.Codigo
		codigo3d += "/*************************DECLARACION DE VECTOR*/\n"
		if d.Stack {
			codigo3d += temp + " = P + " + strconv.Itoa(scope.Size) + ";\n"
			nSimbolo.Direccion = scope.Size
			nSimbolo.TipoDireccion = Ast.STACK
			scope.Size++
			codigo3d += "stack[(int)" + temp + "] = " + obj3dValor.Referencia + ";\n"
		} else {
			codigo3d += temp + " = P + " + strconv.Itoa(scope.Size) + ";\n"
			nSimbolo.Direccion = scope.Size
			nSimbolo.TipoDireccion = Ast.HEAP
			scope.Size = Ast.GetValorH() - 1
			codigo3d += "heap[(int)" + temp + "] = " + obj3dValor.Referencia + ";\n"
		}
		codigo3d += "/***********************************************/\n"
		scope.Add(nSimbolo)
	} else {
		temp := Ast.GetTemp()
		//Actualizar la mutabilidad de la instancia
		nVector := valor.Valor.(expresiones.Vector)
		nVector.Mutable = d.Mutable
		//Agregar el valor al obj3d
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
		/*Agregar el código de la creación del vector*/
		codigo3d += obj3dValor.Codigo
		codigo3d += "/*************************DECLARACION DE VECTOR*/\n"
		if d.Stack {
			codigo3d += temp + " = P + " + strconv.Itoa(scope.Size) + ";\n"
			nSimbolo.Direccion = scope.Size
			nSimbolo.TipoDireccion = Ast.STACK
			scope.Size++
			codigo3d += "stack[(int)" + temp + "] = " + obj3dValor.Referencia + ";\n"
		} else {
			codigo3d += temp + " = P + " + strconv.Itoa(scope.Size) + ";\n"
			nSimbolo.Direccion = scope.Size
			nSimbolo.TipoDireccion = Ast.HEAP
			scope.Size = Ast.GetValorH() - 1
			codigo3d += "heap[(int)" + temp + "] = " + obj3dValor.Referencia + ";\n"
		}
		codigo3d += "/***********************************************/\n"

		scope.Add(nSimbolo)
	}

	obj3d.Codigo = codigo3d
	return Ast.TipoRetornado{Valor: obj3d, Tipo: Ast.DECLARACION}
}

func (op DeclaracionVector) GetFila() int {
	return op.Fila
}
func (op DeclaracionVector) GetColumna() int {
	return op.Columna
}

func (d DeclaracionVector) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, Ast.DECLARACION
}
