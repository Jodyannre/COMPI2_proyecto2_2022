package fn_vectores

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"

	"github.com/colegno/arraylist"
)

type RemoveVec struct {
	Identificador interface{}
	Posicion      interface{}
	Tipo          Ast.TipoDato
	Fila          int
	Columna       int
}

func NewRemoveVec(id interface{}, posicion interface{}, tipo Ast.TipoDato, fila, columna int) RemoveVec {
	nP := RemoveVec{
		Identificador: id,
		Posicion:      posicion,
		Tipo:          tipo,
		Fila:          fila,
		Columna:       columna,
	}
	return nP
}

func (p RemoveVec) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	var simbolo Ast.Simbolo
	var vector expresiones.Vector
	var posicion Ast.TipoRetornado
	var removido Ast.TipoRetornado
	var id string
	/*******VARIABLES PARA C3D********/
	var idExp expresiones.Identificador
	var obj3d, obj3Temp Ast.O3D
	var referencia, codigo3d string
	/*********************************/

	//Primero verificar que sea un identificador el id
	_, tipoParticular := p.Identificador.(Ast.Abstracto).GetTipo()
	if tipoParticular != Ast.IDENTIFICADOR {
		//Error se espera un identificador
		msg := "Semantic error, expected IDENTIFICADOR, found. " + Ast.ValorTipoDato[tipoParticular] +
			". -- Line: " + strconv.Itoa(p.Fila) +
			" Column: " + strconv.Itoa(p.Columna)
		nError := errores.NewError(p.Fila, p.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}
	//Recuperar el id del identificador
	id = p.Identificador.(expresiones.Identificador).Valor

	//Verificar que el id exista
	if !scope.Exist(id) {
		//Error la variable no existe
		msg := "Semantic error, the element \"" + id + "\" doesn't exist in any scope." +
			" -- Line:" + strconv.Itoa(p.Fila) + " Column: " + strconv.Itoa(p.Columna)
		nError := errores.NewError(p.Fila, p.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
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
		msg := "Semantic error, expected Vector, found " + Ast.ValorTipoDato[simbolo.Tipo] + "." +
			" -- Line:" + strconv.Itoa(p.Fila) + " Column: " + strconv.Itoa(p.Columna)
		nError := errores.NewError(p.Fila, p.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}
	vector = simbolo.Valor.(Ast.TipoRetornado).Valor.(expresiones.Vector)

	//Verificar que el vector sea mutable
	if !simbolo.Mutable {
		msg := "Semantic error, can't remove an elment from an inmutable Vector<" +
			expresiones.Tipo_String(vector.TipoVector) + ">." +
			" -- Line: " + strconv.Itoa(p.Fila) +
			" Column: " + strconv.Itoa(p.Columna)
		nError := errores.NewError(p.Fila, p.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}

	//Get la posición de donde se va a extraer el elemento
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
		fila := p.Posicion.(Ast.Abstracto).GetFila()
		columna := p.Posicion.(Ast.Abstracto).GetColumna()
		msg := "Semantic error, expected USIZE, found. " + Ast.ValorTipoDato[posicion.Tipo] +
			". -- Line: " + strconv.Itoa(fila) +
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
	//Verificar que la posición exista en el vector
	if posicion.Valor.(int) > vector.Size {
		//Error, fuera de rango
		fila := p.Posicion.(Ast.Abstracto).GetFila()
		columna := p.Posicion.(Ast.Abstracto).GetColumna()
		msg := "Semantic error, index (" + strconv.Itoa(posicion.Valor.(int)) + ") out of bounds." +
			". -- Line: " + strconv.Itoa(fila) +
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

	//Paso todas las pruebas, entonces guardar el elemento
	//Crear la nueva lista que contendrá los valores después de la eliminación
	nLista := arraylist.New()

	for i := 0; i <= vector.Valor.Len(); i++ {
		if i == posicion.Valor.(int) {
			removido = vector.Valor.GetValue(i).(Ast.TipoRetornado)
			continue
		}
		if i < vector.Valor.Len() {
			nLista.Add(vector.Valor.GetValue(i))
		}
	}
	//Limpiar la lista
	vector.Valor.Clear()

	//Regresar los valores a la lista
	for i := 0; i <= nLista.Len(); i++ {
		vector.Valor.Add(nLista.GetValue(i))
	}

	//Agregar la nueva lista
	vector.Valor = nil
	vector.Valor = nLista
	//Disminuir el tamaño
	vector.Size--
	//vector.Capacity = vector.CalcularCapacity(vector.Size, vector.Capacity)
	//Cambiar el estado de vacio
	if vector.Size == 0 {
		vector.Vacio = true
	}
	simbolo.Valor = Ast.TipoRetornado{Tipo: Ast.VECTOR, Valor: vector}

	/*CODIGO 3D PARA AGREGAR EL ELEMENTO AL VECTOR*/
	/* PRIMERO CREAR UN NUEVO VECTOR Y AGREGARLE EL ELEMENTO DE ÚLTIMO */
	nReferencia, preCodigo3d, referenciaEliminado := EliminarElemento3D(referencia, posicion.Valor.(int), obj3Temp.Referencia)
	codigo3d += preCodigo3d
	obj3d.Referencia = referenciaEliminado
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
		tempRef := Ast.GetTemp()
		temp := Ast.GetTemp()
		codigo3d += "/*********************REGISTRAR EL NUEVO VECTOR*/\n"
		codigo3d += temp + " = " + strconv.Itoa(simbolo.Direccion) + "];\n"

		if simbolo.Referencia {
			codigo3d += tempRef + " = stack[(int)" + temp + "];\n"
			referencia = tempRef
		} else {
			referencia = temp
		}

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
	obj3d.Valor = removido

	return Ast.TipoRetornado{
		Tipo:  Ast.VEC_REMOVE,
		Valor: obj3d,
	}
}

func (v RemoveVec) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, v.Tipo
}

func (v RemoveVec) GetFila() int {
	return v.Fila
}
func (v RemoveVec) GetColumna() int {
	return v.Columna
}

func EliminarElemento3D(referencia string, posicion int, refPos string) (string, string, string) {
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
	referenciaEliminado := Ast.GetTemp()
	lt := Ast.GetLabel()
	lf := Ast.GetLabel()
	lt2 := Ast.GetLabel()
	lf2 := Ast.GetLabel()
	salto := Ast.GetLabel()
	salto2 := Ast.GetLabel()
	codigo3d += inicioNuevoVec + " = H; //Guardo la posición del nuevo vector \n"
	codigo3d += "H = H + 1;\n"
	Ast.GetH()
	codigo3d += posVectorNuevo + " = H; //Inicializar contador del nuevo vector\n"
	codigo3d += sizeAnterior + " = heap[(int)" + referencia + "];//Get el size actual\n"
	codigo3d += nuevoSize + " = " + sizeAnterior + " - 1;//Nuevo size del vector\n"
	codigo3d += "heap[(int)" + inicioNuevoVec + "] = " + nuevoSize + "; //Agregar nuevo size\n"
	codigo3d += posVectorActual + " = " + referencia + " + 1;//Get inicio del vector actual\n"
	codigo3d += contador + " = 0; //Inicializo contador\n"
	codigo3d += "/*****************VERIFICAR ELEMENTO A ELIMINAR*/\n"
	codigo3d += salto2 + ":\n"
	codigo3d += salto + ":\n"
	codigo3d += "if (" + contador + " < " + sizeAnterior + ") goto " + lt + ";\n"
	codigo3d += "goto " + lf + ";\n"
	codigo3d += "/**************VERIFICAR EL ELEMENTO A ELIMINAR*/\n"
	codigo3d += lt + ":\n"
	//codigo3d += "if (" + contador + " != " + strconv.Itoa(posicion) + ") goto " + lt2 + ";\n"
	codigo3d += "if (" + contador + " != " + refPos + ") goto " + lt2 + ";\n"
	codigo3d += "goto " + lf2 + ";\n"
	codigo3d += "/*******************COPIAR ELEMENTOS DEL VECTOR*/\n"
	codigo3d += lt2 + ":\n"
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
	codigo3d += "/*****************SIGUIENTE POSICION DEL VECTOR*/\n"
	codigo3d += temp1 + " = " + contador + " + 1; //Actualizar contador\n"
	codigo3d += contador + " = " + temp1 + ";\n"
	codigo3d += referenciaEliminado + " = heap[(int)" + posVectorActual + "]; //Guardar el valor del elemento a eliminar\n"
	codigo3d += temp2 + " = " + posVectorActual + " + 1;//Actualizar pos vec actual\n"
	codigo3d += posVectorActual + " = " + temp2 + ";\n"
	codigo3d += "goto " + salto2 + ";\n"
	codigo3d += lf + ":\n"
	codigo3d += "/***********************************************/\n"
	return inicioNuevoVec, codigo3d, referenciaEliminado
}
