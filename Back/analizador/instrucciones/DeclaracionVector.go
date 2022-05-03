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
	var codigo3d, scopeAnterior string
	var obj3d, obj3dValor, obj3dTemp Ast.O3D
	var simboloVector Ast.Simbolo
	esIndefinido := false
	_, tipoIn := d.Valor.(Ast.Abstracto).GetTipo()
	if tipoIn == Ast.VALOR {
		scopeAnterior = Ast.GetTemp()
		/*********************SCOPE SIMULADO****************************/
		codigo3d += scopeAnterior + " = P; //Guardar el scope anterior \n"
		codigo3d += "P = " + strconv.Itoa(d.ScopeOriginal.Posicion) + "; //Scope de donde proviene el valor\n"
		/***************************************************************/
		existe = scope.Exist_actual(d.Id)
		valor = d.Valor.(Ast.Expresion).GetValue(d.ScopeOriginal)
		obj3dTemp = valor.Valor.(Ast.O3D)
		valor = obj3dTemp.Valor
		codigo3d += obj3dTemp.Codigo
		/*********************RETORNO SCOPE ANTERIOR********************/
		codigo3d += "P = " + scopeAnterior + "; //Retornar al scope anterior \n"
		/***************************************************************/
		if obj3dTemp.EsReferencia != "" {
			simboloVector = d.ScopeOriginal.GetSimbolo(obj3dTemp.EsReferencia)
		} else {
			simboloVector = scope.GetSimbolo(d.Id)
		}

		obj3dValor = obj3dTemp
	} else {
		existe = scope.Exist_actual(d.Id)
		simboloVector = scope.GetSimbolo(d.Id)
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
		////////////////////////////ERROR//////////////////////////////////
		return errores.GenerarError(13, d, d, "",
			Ast.ValorTipoDato[d.Tipo],
			Ast.ValorTipoDato[tipoIn],
			scope)
		//////////////////////////////////////////////////////////////////
	}

	//Verificar si ya existe
	if existe {
		////////////////////////////ERROR//////////////////////////////////
		return errores.GenerarError(12, d, d, d.Id,
			"",
			"",
			scope)
		//////////////////////////////////////////////////////////////////
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
			////////////////////////////ERROR//////////////////////////////////
			return errores.GenerarError(38, d, d, "",
				expresiones.Tipo_String(d.TipoVector),
				expresiones.Tipo_String(valor.Valor.(expresiones.Vector).TipoVector),
				scope)
			//////////////////////////////////////////////////////////////////
		}
	}

	//Crear el símbolo y agregarlo al scope
	if esIndefinido {
		temp := Ast.GetTemp()
		referenciaVector := Ast.GetTemp()
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
			Referencia:    true,
		}
		if simboloVector.Entorno == nil {
			nSimbolo.Referencia = false
		} else {
			nSimbolo.Referencia_puntero = &simboloVector
		}

		/*Codigo 3d*/
		/*Agregar el código de la creación del vector*/
		codigo3d += obj3dValor.Codigo
		if simboloVector.Referencia && simboloVector.Entorno != nil {
			scopeAnterior := Ast.GetTemp()
			//posVectorRef := Ast.GetTemp()
			codigo3d += scopeAnterior + " = P; //Guardar scope anterior \n"
			codigo3d += "P = " + strconv.Itoa(simboloVector.Entorno.Posicion) + "; //Cambio de entorno \n"
			codigo3d += referenciaVector + " = P + " + strconv.Itoa(simboloVector.Direccion) + "; //Pos vec ref \n"
			codigo3d += "P = " + scopeAnterior + "; //regresar al entorno anterior \n"
			//codigo3d += referenciaVector + " = " + strconv.Itoa(simboloVector.Direccion) + "; //Dir referencia\n"
		} else {
			codigo3d += referenciaVector + " = " + obj3dValor.Referencia + "; //Dir vector\n"
		}

		codigo3d += "/*************************DECLARACION DE VECTOR*/\n"
		if d.Stack {
			codigo3d += temp + " = P + " + strconv.Itoa(scope.ContadorDeclaracion) + ";\n"
			nSimbolo.Direccion = scope.ContadorDeclaracion
			nSimbolo.TipoDireccion = Ast.STACK
			scope.ContadorDeclaracion++
			codigo3d += "stack[(int)" + temp + "] = " + referenciaVector + ";\n"
		} else {
			codigo3d += temp + " = P + " + strconv.Itoa(scope.ContadorDeclaracion) + ";\n"
			nSimbolo.Direccion = scope.ContadorDeclaracion
			nSimbolo.TipoDireccion = Ast.HEAP
			scope.ContadorDeclaracion = Ast.GetValorH() - 1
			codigo3d += "heap[(int)" + temp + "] = " + referenciaVector + ";\n"
		}
		codigo3d += "/***********************************************/\n"
		scope.Add(nSimbolo)
	} else {
		temp := Ast.GetTemp()
		referenciaVector := Ast.GetTemp()
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
			Referencia:    true,
		}

		if simboloVector.Entorno == nil {
			nSimbolo.Referencia = false
		} else {
			nSimbolo.Referencia_puntero = &simboloVector
		}
		codigo3d += obj3dValor.Codigo
		if nSimbolo.Referencia && simboloVector.Entorno != nil {
			codigo3d += referenciaVector + " = " + strconv.Itoa(simboloVector.Direccion) + "; //Dir referencia\n"
		} else {
			codigo3d += referenciaVector + " = " + obj3dValor.Referencia + "; //Dir vector\n"
		}

		/*Codigo 3d*/
		/*Agregar el código de la creación del vector*/

		codigo3d += "/*************************DECLARACION DE VECTOR*/\n"
		if d.Stack {
			codigo3d += temp + " = P + " + strconv.Itoa(scope.ContadorDeclaracion) + ";\n"
			nSimbolo.Direccion = scope.ContadorDeclaracion
			nSimbolo.TipoDireccion = Ast.STACK
			scope.ContadorDeclaracion++
			codigo3d += "stack[(int)" + temp + "] = " + referenciaVector + ";\n"
		} else {
			codigo3d += temp + " = P + " + strconv.Itoa(scope.ContadorDeclaracion) + ";\n"
			nSimbolo.Direccion = scope.ContadorDeclaracion
			nSimbolo.TipoDireccion = Ast.HEAP
			scope.ContadorDeclaracion = Ast.GetValorH() - 1
			codigo3d += "heap[(int)" + temp + "] = " + referenciaVector + ";\n"
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
