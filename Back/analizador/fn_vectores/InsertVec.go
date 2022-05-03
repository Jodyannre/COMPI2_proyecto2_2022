package fn_vectores

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"

	"github.com/colegno/arraylist"
)

type InsertVec struct {
	Identificador interface{}
	Posicion      interface{}
	Valor         interface{}
	Tipo          Ast.TipoDato
	Fila          int
	Columna       int
}

func NewInsertVec(id interface{}, valor interface{}, posicion interface{}, tipo Ast.TipoDato, fila, columna int) InsertVec {
	nP := InsertVec{
		Valor:         valor,
		Identificador: id,
		Posicion:      posicion,
		Tipo:          tipo,
		Fila:          fila,
		Columna:       columna,
	}
	return nP
}

func (p InsertVec) Run(scope *Ast.Scope) interface{} {
	var simbolo Ast.Simbolo
	var vector expresiones.Vector
	var valor Ast.TipoRetornado
	var posicion Ast.TipoRetornado
	var id string
	/*******VARIABLES PARA C3D********/
	var idExp expresiones.Identificador
	var obj3d, obj3Temp, obj3dValor Ast.O3D
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
			Ast.ValorTipoDato[vector.Tipo],
			scope)
		//////////////////////////////////////////////////////////////////
	}

	if valor.Tipo != vector.TipoVector.Tipo {
		//Error de tipos dentro del vector
		////////////////////////////ERROR//////////////////////////////////
		return errores.GenerarError(27, p, p, "",
			Ast.ValorTipoDato[valor.Tipo],
			Ast.ValorTipoDato[vector.Tipo],
			scope)
		//////////////////////////////////////////////////////////////////
	}

	//Verificar si es vector el que se va a agregar y el tipo del vector
	if valor.Tipo == Ast.VECTOR {
		if !expresiones.CompararTipos(valor.Valor.(expresiones.Vector).TipoVector, vector.TipoVector) {
			//Error, no se puede guardar ese tipo de vector en este vector
			////////////////////////////ERROR//////////////////////////////////
			return errores.GenerarError(27, p, p, "",
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
			return errores.GenerarError(27, p, p, "",
				plantilla,
				expresiones.Tipo_String(vector.TipoVector),
				scope)
			//////////////////////////////////////////////////////////////////
		}
	}

	//Get la posición en donde se quiere agregar el nuevo valor
	posicion = p.Posicion.(Ast.Expresion).GetValue(scope)
	obj3Temp = posicion.Valor.(Ast.O3D)
	posicion = obj3Temp.Valor
	codigo3d += obj3Temp.Codigo
	_, tipoParticular = p.Posicion.(Ast.Abstracto).GetTipo()
	if posicion.Tipo == Ast.ERROR {
		return posicion
	}
	//Verificar que el número en el acceso sea usize
	if (posicion.Tipo != Ast.USIZE && posicion.Tipo != Ast.I64) ||
		tipoParticular == Ast.IDENTIFICADOR && posicion.Tipo == Ast.I64 {
		//Error, se espera un usize
		////////////////////////////ERROR//////////////////////////////////
		return errores.GenerarError(33, p.Posicion, p.Posicion, "",
			Ast.ValorTipoDato[posicion.Tipo],
			"",
			scope)
		//////////////////////////////////////////////////////////////////
	}
	//Verificar que la posición exista en el vector
	if posicion.Valor.(int) > vector.Size {
		//Error, fuera de rango
		////////////////////////////ERROR//////////////////////////////////
		return errores.GenerarError(47, p.Posicion, p.Posicion, "",
			strconv.Itoa(posicion.Valor.(int)),
			"",
			scope)
		//////////////////////////////////////////////////////////////////
	}

	//Paso todas las pruebas, entonces guardar el elemento
	//Crear la nueva lista que contendrá los valores
	nLista := arraylist.New()

	for i := 0; i <= vector.Valor.Len(); i++ {
		if i == posicion.Valor.(int) {
			nLista.Add(valor)
		}
		if i < vector.Valor.Len() {
			nLista.Add(vector.Valor.GetValue(i))
		}
	}
	//Agregar la nueva lista
	vector.Valor = nil
	vector.Valor = nLista
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
	nReferencia, preCodigo3d := InsertarElemento3D(referencia, obj3dValor.Referencia, posicion.Valor.(int), obj3Temp.Referencia)
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

	/*Actualizar el simbolo en la tabla de simbolos*/
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

func (v InsertVec) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, v.Tipo
}

func (v InsertVec) GetFila() int {
	return v.Fila
}
func (v InsertVec) GetColumna() int {
	return v.Columna
}

func InsertarElemento3D(referencia, referenciaNuevoValor string, posicion int, pos3d string) (string, string) {
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
	lt2 := Ast.GetLabel()
	lf2 := Ast.GetLabel()
	lt3 := Ast.GetLabel()
	lf3 := Ast.GetLabel()
	salto := Ast.GetLabel()
	salto2 := Ast.GetLabel()
	codigo3d += inicioNuevoVec + " = H; //Guardo la posición del nuevo vector \n"
	codigo3d += "H = H + 1;\n"
	Ast.GetH()
	codigo3d += posVectorNuevo + " = H; //Inicializar contador del nuevo vector\n"
	codigo3d += sizeAnterior + " = heap[(int)" + referencia + "];//Get el size actual\n"
	codigo3d += nuevoSize + " = " + sizeAnterior + " + 1;//Nuevo size del vector\n"
	codigo3d += "heap[(int)" + inicioNuevoVec + "] = " + nuevoSize + "; //Agregar nuevo size\n"
	codigo3d += posVectorActual + " = " + referencia + " + 1;//Get inicio del vector actual\n"
	codigo3d += contador + " = 0; //Inicializo contador\n"
	codigo3d += "/******************VERIFICAR ELEMENTO A AGREGAR*/\n"
	codigo3d += salto + ":\n"
	codigo3d += "if (" + contador + " <= " + sizeAnterior + ") goto " + lt + ";\n"
	codigo3d += "goto " + lf + ";\n"
	codigo3d += "/*********VERIFICAR POSICION DEL NUEVO ELEMENTO*/\n"
	codigo3d += lt + ":\n"
	//codigo3d += "if (" + contador + " != " + strconv.Itoa(posicion) + ") goto " + lt2 + ";\n"
	codigo3d += "if (" + contador + " != " + pos3d + ") goto " + lt2 + ";\n"
	codigo3d += "goto " + lf2 + ";\n"
	codigo3d += "/*******************COPIAR ELEMENTOS DEL VECTOR*/\n"
	codigo3d += salto2 + ":\n"
	codigo3d += lt2 + ":\n"
	codigo3d += "if (" + contador + " == " + sizeAnterior + ") goto " + lf3 + ";\n"
	codigo3d += "goto " + lt3 + ";\n"
	codigo3d += lt3 + ":\n"
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
	codigo3d += lf2 + ":\n"
	codigo3d += "/**************AGREGAR EL NUEVO ELEMENTO********/\n"
	codigo3d += "heap[(int)" + posVectorNuevo + "] = " + referenciaNuevoValor + "; //Add nuevo elemento\n"
	codigo3d += "H = H + 1;\n"
	Ast.GetH()
	codigo3d += posVectorNuevo + " = H ;//Actualizra pos vec nuevo\n"
	codigo3d += "goto " + salto2 + ";\n"
	codigo3d += lf3 + ":\n"
	codigo3d += lf + ":\n"
	codigo3d += "/***********************************************/\n"
	return inicioNuevoVec, codigo3d
}
