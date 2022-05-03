package fn_vectores

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"
)

type Push struct {
	Identificador interface{}
	Valor         interface{}
	Tipo          Ast.TipoDato
	Fila          int
	Columna       int
}

func NewPush(id interface{}, valor interface{}, tipo Ast.TipoDato, fila, columna int) Push {
	nP := Push{
		Valor:         valor,
		Identificador: id,
		Tipo:          tipo,
		Fila:          fila,
		Columna:       columna,
	}
	return nP
}

func (p Push) Run(scope *Ast.Scope) interface{} {
	var simbolo Ast.Simbolo
	var vector expresiones.Vector
	var valor Ast.TipoRetornado
	var id string
	/*******VARIABLES PARA C3D********/
	var idExp expresiones.Identificador
	var obj3d, obj3dValor Ast.O3D
	var referencia, codigo3d string
	/*********************************/

	//Primero verificar que sea un identificador el id
	_, tipoParticular := p.Identificador.(Ast.Abstracto).GetTipo()
	if tipoParticular != Ast.IDENTIFICADOR {
		//Error se espera un identificador
		////////////////////////////ERROR//////////////////////////////////
		return errores.GenerarError(36, p, p, "",
			Ast.ValorTipoDato[tipoParticular],
			"",
			scope)
		//////////////////////////////////////////////////////////////////
	}
	//Recuperar el id del identificador
	id = p.Identificador.(expresiones.Identificador).Valor

	//Verificar que el id exista
	if !scope.Exist(id) {
		//Error la variable no existe
		////////////////////////////ERROR//////////////////////////////////
		return errores.GenerarError(15, p, p, id,
			"",
			"",
			scope)
		//////////////////////////////////////////////////////////////////
	}
	//Conseguir el simbolo y el vector
	simbolo = scope.GetSimbolo(id)

	/*********C3D del Simbolo***********/
	codigo3d += "/********************************ACCESO A VECTOR*/\n"
	idExp = expresiones.NewIdentificador(id, Ast.IDENTIFICADOR, 0, 0)
	obj3d = idExp.GetValue(scope).Valor.(Ast.O3D)
	referencia = obj3d.Referencia
	codigo3d += obj3d.Codigo
	/***********************************/
	//Verificar que sea un vector
	if simbolo.Tipo != Ast.VECTOR {
		////////////////////////////ERROR//////////////////////////////////
		return errores.GenerarError(34, p, p, "",
			Ast.ValorTipoDato[simbolo.Tipo],
			"",
			scope)
		//////////////////////////////////////////////////////////////////
	}
	vector = simbolo.Valor.(Ast.TipoRetornado).Valor.(expresiones.Vector)

	//Verificar que el elemento que se va a agregar sea del mismo tipo que el que guarda el vector
	//Primero calcular el valor
	valor = p.Valor.(Ast.Expresion).GetValue(scope)
	obj3dValor = valor.Valor.(Ast.O3D)
	valor = obj3dValor.Valor
	codigo3d += obj3dValor.Codigo
	if valor.Tipo == Ast.ERROR {
		return valor
	}

	//Verificar que el vector sea mutable
	if !simbolo.Mutable {
		////////////////////////////ERROR//////////////////////////////////
		return errores.GenerarError(51, p, p, "",
			Ast.ValorTipoDato[valor.Tipo],
			expresiones.Tipo_String(vector.TipoVector),
			scope)
		//////////////////////////////////////////////////////////////////
	}

	if valor.Tipo != vector.TipoVector.Tipo {
		//Error de tipos dentro del vector
		////////////////////////////ERROR//////////////////////////////////
		return errores.GenerarError(35, p, p, "",
			Ast.ValorTipoDato[valor.Tipo],
			Ast.ValorTipoDato[vector.Tipo],
			scope)
		//////////////////////////////////////////////////////////////////
	}

	//Verificar si es vector el que se va a agregar y el tipo del vector
	if valor.Tipo == Ast.VECTOR {
		//if !expresiones.CompararTipos(valor.Valor.(expresiones.Vector).TipoVector, vector.TipoVector) {
		if valor.Tipo != vector.TipoVector.Tipo {
			//Error, no se puede guardar ese tipo de vector en este vector
			////////////////////////////ERROR//////////////////////////////////
			return errores.GenerarError(35, p, p, "",
				expresiones.Tipo_String(valor.Valor.(expresiones.Vector).TipoVector),
				expresiones.Tipo_String(vector.TipoVector),
				scope)
			//////////////////////////////////////////////////////////////////
		}
	}
	//Verificar si es un struct el que se va a agregar
	if valor.Tipo == Ast.STRUCT {
		plantilla := valor.Valor.(Ast.Structs).GetPlantilla(scope)
		tipoStruct := Ast.TipoRetornado{Valor: plantilla, Tipo: Ast.STRUCT}
		if !expresiones.CompararTipos(tipoStruct, vector.TipoVector) {
			//Error, no se puede guardar ese tipo de vector en este vector
			////////////////////////////ERROR//////////////////////////////////
			return errores.GenerarError(35, p, p, "",
				plantilla,
				expresiones.Tipo_String(vector.TipoVector),
				scope)
			//////////////////////////////////////////////////////////////////
		}
	}

	//Paso todas las pruebas, entonces guardar el elemento
	vector.Valor.Add(valor)
	//Aumentar el tamaño
	vector.Size++
	vector.Capacity = vector.CalcularCapacity(vector.Size, vector.Capacity)
	//Cambiar el estado de vacio
	if vector.Vacio {
		vector.Vacio = false
	}
	simbolo.Valor = Ast.TipoRetornado{Tipo: Ast.VECTOR, Valor: vector}

	/*CODIGO 3D PARA AGREGAR EL ELEMENTO AL VECTOR*/
	/* PRIMERO CREAR UN NUEVO VECTOR Y AGREGARLE EL ELEMENTO DE ÚLTIMO */
	nReferencia, preCodigo3d := GetNuevoVector3D(referencia, obj3dValor.Referencia)
	codigo3d += preCodigo3d
	/*ACTUALIZAR LA POSICIÖN EN LA TABLA DE SÍMBOLOS*/
	if simbolo.TipoDireccion == Ast.STACK {
		/*ESTA GUARDADO EN EL STACK*/
		temp := Ast.GetTemp()
		tempRef := Ast.GetTemp()
		codigo3d += "/*********************REGISTRAR EL NUEVO VECTOR*/\n"
		codigo3d += temp + " = P + " + strconv.Itoa(simbolo.Direccion) + ";\n"

		if simbolo.Referencia {
			codigo3d += tempRef + " = stack[(int)" + temp + "];\n"
			referencia = tempRef
		} else {
			referencia = temp
		}

		codigo3d += "stack[(int)" + referencia + "] = " + nReferencia + ";\n"
		codigo3d += "/***********************************************/\n"
		//nuevaDireccion, _ := strconv.Atoi(nReferencia)
		//simbolo.Direccion = nuevaDireccion
	} else {
		/*ESTA GUARDANDO EN EL HEAP*/
		temp := Ast.GetTemp()
		codigo3d += "/*********************REGISTRAR EL NUEVO VECTOR*/\n"
		codigo3d += temp + " = " + strconv.Itoa(simbolo.Direccion) + "];\n"
		referencia = temp
		codigo3d += "heap[(int)" + referencia + "] = " + nReferencia + ";\n"
		codigo3d += "/***********************************************/\n"
		//nuevaDireccion, _ := strconv.Atoi(nReferencia)
		//simbolo.Direccion = nuevaDireccion
	}
	/*Actualizar el simbolo en la tabla de símbolos*/
	scope.UpdateSimbolo(id, simbolo)
	if simbolo.Referencia {
		id := simbolo.Referencia_puntero.Identificador
		simbolo.Entorno.UpdateValor(id, simbolo.Valor)
	}

	obj3d.Codigo = codigo3d
	obj3d.Valor = Ast.TipoRetornado{
		Tipo:  Ast.EJECUTADO,
		Valor: true,
	}

	return Ast.TipoRetornado{
		Tipo:  Ast.VEC_PUSH,
		Valor: obj3d,
	}
}

func (v Push) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, v.Tipo
}

func (v Push) GetFila() int {
	return v.Fila
}
func (v Push) GetColumna() int {
	return v.Columna
}

func GetNuevoVector3D(referencia, referenciaNuevoValor string) (string, string) {
	codigo3d := ""
	inicioNuevoVec := Ast.GetTemp()
	posVectorNuevo := Ast.GetTemp()
	posVectorActual := Ast.GetTemp()
	sizeAnterior := Ast.GetTemp()
	nuevoSize := Ast.GetTemp()
	contador := Ast.GetTemp()
	elementoActual := Ast.GetTemp()
	temp1 := Ast.GetTemp()
	temp2 := Ast.GetTemp()
	lt := Ast.GetLabel()
	lf := Ast.GetLabel()
	salto := Ast.GetLabel()
	codigo3d += inicioNuevoVec + " = H; //Guardo la posición del nuevo vector \n"
	codigo3d += "H = H + 1;\n"
	Ast.GetH()
	codigo3d += posVectorNuevo + " = H; //Inicializar contador del nuevo vector\n"
	codigo3d += sizeAnterior + " = heap[(int)" + referencia + "];//Get el size actual\n"
	codigo3d += nuevoSize + " = " + sizeAnterior + " + 1;//Nuevo size del vector\n"
	codigo3d += "heap[(int)" + inicioNuevoVec + "] = " + nuevoSize + "; //Agregar nuevo size\n"
	codigo3d += posVectorActual + " = " + referencia + " + 1;//Get inicio del vector actual\n"
	codigo3d += contador + " = 0; //Inicializo contador\n"
	codigo3d += "/*******************COPIAR ELEMENTOS DEL VECTOR*/\n"
	codigo3d += salto + ":\n"
	codigo3d += "if (" + contador + " < " + sizeAnterior + ") goto " + lt + ";\n"
	codigo3d += "goto " + lf + ";\n"
	codigo3d += lt + ":\n"
	codigo3d += elementoActual + " = heap[(int)" + posVectorActual + "]; //Valor del elemento\n"
	codigo3d += "heap[(int)" + posVectorNuevo + "] = " + elementoActual + "; //Add nuevo elemento\n"
	codigo3d += temp1 + " = " + contador + " + 1; //Actualizar contador\n"
	codigo3d += contador + " = " + temp1 + ";\n"
	codigo3d += temp2 + " = " + posVectorActual + " + 1;//Actualizar pos vec actual\n"
	codigo3d += posVectorActual + " = " + temp2 + ";\n"
	codigo3d += "H = H + 1;\n"
	Ast.GetH()
	codigo3d += posVectorNuevo + " = H ;//Actualizra pos vec nuevo\n"
	codigo3d += "goto " + salto + ";\n"
	codigo3d += lf + ":\n"
	codigo3d += "/************************AGREGAR EL NUEVO VALOR*/\n"
	codigo3d += "heap[(int)" + posVectorNuevo + "] = " + referenciaNuevoValor + ";\n"
	codigo3d += "H = H + 1;\n"
	Ast.GetH()
	codigo3d += "/***********************************************/\n"
	return inicioNuevoVec, codigo3d
}

func VerificarTipoEntrante(elemento interface{}, tipoVector Ast.TipoRetornado) {

}
