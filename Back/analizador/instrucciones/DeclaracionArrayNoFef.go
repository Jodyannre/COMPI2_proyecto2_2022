package instrucciones

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"Back/analizador/fn_array"
	"strconv"
	"strings"

	"github.com/colegno/arraylist"
)

type DeclaracionArrayNoRef struct {
	Id            string
	Tipo          Ast.TipoDato
	TipoArray     Ast.TipoRetornado
	Dimension     interface{}
	Mutable       bool
	Publico       bool
	Valor         interface{}
	Fila          int
	Columna       int
	ScopeOriginal *Ast.Scope
	Stack         bool
}

func NewDeclaracionArrayNoRef(id string, dimension interface{},
	mutable, publico bool, valor interface{}, fila int, columna int) DeclaracionArrayNoRef {
	nd := DeclaracionArrayNoRef{
		Id:        id,
		Tipo:      Ast.DECLARACION_ARRAY,
		Publico:   publico,
		Mutable:   mutable,
		Valor:     valor,
		Fila:      fila,
		Columna:   columna,
		Dimension: dimension,
		TipoArray: dimension.(expresiones.DimensionArray).TipoArray,
		Stack:     true,
	}
	return nd
}

func (d DeclaracionArrayNoRef) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, Ast.DECLARACION
}

func (d DeclaracionArrayNoRef) Run(scope *Ast.Scope) interface{} {
	//Verificar que exista, recuperar los arryas y los tipos
	var validacionDimensiones string
	var existe bool
	var valor Ast.TipoRetornado
	/**********VARIABLES 3D***************/
	var codigo3d string = ""
	var obj3dValor, obj3d, obj3dTemp, obj3dClone, obj3dDimension Ast.O3D /*obj3dClone*/
	var scopeAnterior string
	/*************************************/
	_, tipoIn := d.Valor.(Ast.Abstracto).GetTipo()
	if tipoIn == Ast.VALOR {
		existe = scope.Exist_actual(d.Id)
		//valor = d.Valor.(Ast.Expresion).GetValue(d.ScopeOriginal)
		scopeAnterior = Ast.GetTemp()
		/*********************SCOPE SIMULADO****************************/
		codigo3d += scopeAnterior + " = P; //Guardar el scope anterior \n"
		codigo3d += "P = " + strconv.Itoa(d.ScopeOriginal.Posicion) + "; //Scope de donde proviene el valor\n"
		/***************************************************************/
		valor = d.Valor.(Ast.Expresion).GetValue(d.ScopeOriginal)
		obj3dTemp = valor.Valor.(Ast.O3D)
		valor = obj3dTemp.Valor
		codigo3d += obj3dTemp.Codigo
		/*********************RETORNO SCOPE ANTERIOR********************/
		codigo3d += "P = " + scopeAnterior + "; //Retornar al scope anterior \n"
		/***************************************************************/
		/*********************ACTUALIZAR TIPO IN************************/
		tipoIn = valor.Tipo
		/***************************************************************/
	} else {
		existe = scope.Exist_actual(d.Id)
		valor = d.Valor.(Ast.Expresion).GetValue(scope)
		obj3dValor = valor.Valor.(Ast.O3D)
		valor = obj3dValor.Valor
	}

	dimension := d.Dimension.(Ast.Expresion).GetValue(scope)

	if existe {
		////////////////////////////ERROR//////////////////////////////////
		return errores.GenerarError(12, d, d, d.Id,
			"",
			"",
			scope)
		//////////////////////////////////////////////////////////////////
	}
	//Verificar que no venga ningún error
	if valor.Tipo == Ast.ERROR {
		return valor
	}
	//Verificar error en dimension
	if dimension.Tipo == Ast.ERROR {
		return dimension
	}

	//Verificar los tipos

	//Verificar que el tipo del vector no sea un acceso a modulo

	if d.TipoArray.Tipo == Ast.ACCESO_MODULO || EsTipoEspecial(d.TipoArray.Tipo) {
		//Traer el tipo y cambiar el tipo de la declaración
		nTipo := GetTipoEstructura(d.TipoArray, scope)
		if nTipo.Tipo == Ast.ERROR {
			return nTipo
		}
		if nTipo.Tipo == Ast.STRUCT_TEMPLATE {
			nTipo.Tipo = Ast.STRUCT
		}
		d.TipoArray = nTipo
	}

	//Verificar que los tipos de los vectores sean correctos
	if !expresiones.CompararTipos(d.TipoArray, expresiones.GetTipoFinal(valor.Valor.(expresiones.Array).TipoDelArray)) {

		if valor.Valor.(expresiones.Array).TipoArray == Ast.INDEFINIDO {
			//Es uno vacio y no hay error, modificar el tipo
		} else {
			////////////////////////////ERROR//////////////////////////////////
			return errores.GenerarError(38, d, d, "",
				expresiones.Tipo_String(d.TipoArray),
				expresiones.Tipo_String(valor.Valor.(expresiones.Vector).TipoVector),
				scope)
			//////////////////////////////////////////////////////////////////
		}
	}

	//Recuperar el tipo del array que se espera desde dimension
	//d.TipoArray = d.Dimension.(expresiones.DimensionArray).TipoArray

	//Primero que vengan arrays
	if !EsArray(tipoIn) {
		//Error, no se estan asignado arrays al array
		////////////////////////////ERROR//////////////////////////////////
		return errores.GenerarError(39, d, d, "",
			Ast.ValorTipoDato[tipoIn],
			"",
			scope)
		//////////////////////////////////////////////////////////////////
	}

	//Verificar que las dimensiones concuerda con la lista de arrays
	validacionDimensiones = fn_array.ConcordanciaArray(valor.Valor.(expresiones.Array))
	//Conseguir la lista
	split := strings.Split(validacionDimensiones, ",")
	//Crear la lista con las posiciones
	listaDimensiones := arraylist.New()
	for _, num := range split {
		numero, _ := strconv.Atoi(num)
		listaDimensiones.Add(numero)
	}

	//Comparar las lista de dimensiones
	//Get primitivos del array de dimension
	arrayDimension := arraylist.New()
	for i := 0; i < dimension.Valor.(*arraylist.List).Len(); i++ {
		obj3dDimension = dimension.Valor.(*arraylist.List).GetValue(i).(Ast.TipoRetornado).Valor.(Ast.O3D)
		arrayDimension.Add(obj3dDimension.Valor.Valor)
	}

	if !fn_array.CompararListas(listaDimensiones, arrayDimension) {
		////////////////////////////ERROR//////////////////////////////////
		return errores.GenerarError(25, d, d, "",
			"",
			"",
			scope)
		//////////////////////////////////////////////////////////////////
	}
	//expresiones.GetTipoFinal(valor.Valor.(expresiones.Array).TipoDelArray.Tipo)
	//Validar el tipo del array
	if d.TipoArray.Tipo != expresiones.GetTipoFinal(valor.Valor.(expresiones.Array).TipoDelArray).Tipo {
		var tipoDelArray string
		if valor.Valor.(expresiones.Array).TipoDelArray.Tipo == Ast.INDEFINIDO {
			tipoDelArray = Ast.ValorTipoDato[d.TipoArray.Tipo]
		} else {
			tipoDelArray = expresiones.Tipo_String(expresiones.GetTipoFinal(d.TipoArray))
		}
		////////////////////////////ERROR//////////////////////////////////
		return errores.GenerarError(40, d, d, "",
			tipoDelArray,
			Ast.ValorTipoDato[expresiones.GetTipoFinal(valor.Valor.(expresiones.Array).TipoDelArray).Tipo],
			scope)
		//////////////////////////////////////////////////////////////////
	}

	//Crear el símbolo
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

	//Preparar la referencia del array original
	arrayTemp := valor.Valor.(expresiones.Array)
	arrayTemp.Referencia = obj3dTemp.Referencia
	valor.Valor = arrayTemp
	//Clonar el array
	nArray := valor.Valor.(Ast.Clones).Clonar(scope)
	obj3dClone = nArray.(Ast.TipoRetornado).Valor.(Ast.O3D)
	nArray = obj3dClone.Valor.Valor
	codigo3d += obj3dClone.Codigo
	//Actualizar la mutabilidad de la instancia
	mnArray := nArray.(expresiones.Array)
	mnArray.Mutable = d.Mutable
	nArray = mnArray
	//Agregar el resultado al obj3d
	obj3d.Valor = Ast.TipoRetornado{Tipo: Ast.ARRAY, Valor: nArray}

	nSimbolo.Valor = Ast.TipoRetornado{
		Tipo:  valor.Tipo,
		Valor: nArray,
	}

	temp := Ast.GetTemp()
	codigo3d += obj3dValor.Codigo
	codigo3d += "/**************************DECLARACION DE ARRAY*/\n"
	if d.Stack {
		codigo3d += temp + " = P + " + strconv.Itoa(scope.ContadorDeclaracion) + ";\n"
		nSimbolo.Direccion = scope.ContadorDeclaracion
		nSimbolo.TipoDireccion = Ast.STACK
		scope.ContadorDeclaracion++
		codigo3d += "stack[(int)" + temp + "] = " + obj3dClone.Referencia + ";\n"
		Ast.GetP()
	} else {
		codigo3d += temp + " = P + " + strconv.Itoa(scope.ContadorDeclaracion) + ";\n"
		nSimbolo.Direccion = scope.ContadorDeclaracion
		nSimbolo.TipoDireccion = Ast.HEAP
		scope.ContadorDeclaracion++
		codigo3d += "heap[(int)" + temp + "] = " + obj3dClone.Referencia + ";\n"
		Ast.GetH()
	}
	codigo3d += "/***********************************************/\n"

	scope.Add(nSimbolo)
	obj3d.Codigo = codigo3d
	return Ast.TipoRetornado{Valor: obj3d, Tipo: Ast.DECLARACION}
}

func (op DeclaracionArrayNoRef) GetFila() int {
	return op.Fila
}
func (op DeclaracionArrayNoRef) GetColumna() int {
	return op.Columna
}
