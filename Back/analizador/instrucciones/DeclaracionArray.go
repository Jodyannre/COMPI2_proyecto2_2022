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

type DeclaracionArray struct {
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

func NewDeclaracionArray(id string, dimension interface{},
	mutable, publico bool, valor interface{}, fila int, columna int) DeclaracionArray {
	nd := DeclaracionArray{
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

func (d DeclaracionArray) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, Ast.DECLARACION
}

func (d DeclaracionArray) Run(scope *Ast.Scope) interface{} {
	//Verificar que exista, recuperar los arryas y los tipos
	var validacionDimensiones string
	var existe bool
	var valor Ast.TipoRetornado
	var simboloArray Ast.Simbolo
	/**********VARIABLES 3D***************/
	var codigo3d, scopeAnterior string
	var obj3d, obj3dValor, objtemp, obj3dTemp Ast.O3D
	/*************************************/
	_, tipoIn := d.Valor.(Ast.Abstracto).GetTipo()

	if tipoIn == Ast.VALOR {
		//existe = d.ScopeOriginal.Exist_actual(d.Id)
		//valor = d.Valor.(Ast.Expresion).GetValue(d.ScopeOriginal)
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
			simboloArray = d.ScopeOriginal.GetSimbolo(obj3dTemp.EsReferencia)
		} else {
			simboloArray = scope.GetSimbolo(d.Id)
		}

		obj3dValor = obj3dTemp
	} else {
		existe = scope.Exist_actual(d.Id)
		simboloArray = scope.GetSimbolo(d.Id)
		valor = d.Valor.(Ast.Expresion).GetValue(scope)
		obj3dValor = valor.Valor.(Ast.O3D)
		valor = obj3dValor.Valor
	}
	tipoIn = valor.Tipo
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
		objtemp = (dimension.Valor.(*arraylist.List).GetValue(i).(Ast.TipoRetornado).Valor.(Ast.O3D))
		arrayDimension.Add(objtemp.Valor.Valor)
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
		Referencia:    true,
	}

	if simboloArray.Entorno == nil {
		nSimbolo.Referencia = false
	} else {
		nSimbolo.Referencia_puntero = &simboloArray
	}

	//Actualizar la mutabilidad de la instancia
	nArray := valor.Valor.(expresiones.Array)
	nArray.Mutable = d.Mutable
	nValor := Ast.TipoRetornado{
		Tipo:  Ast.ARRAY,
		Valor: nArray,
	}
	nSimbolo.Valor = nValor

	temp := Ast.GetTemp()
	referenciaArray := Ast.GetTemp()
	codigo3d += obj3dValor.Codigo

	if simboloArray.Referencia && simboloArray.Entorno != nil {
		scopeAnterior := Ast.GetTemp()
		//posVectorRef := Ast.GetTemp()
		codigo3d += scopeAnterior + " = P; //Guardar scope anterior \n"
		//codigo3d += "P = " + strconv.Itoa(simboloArray.Entorno.Posicion) + "; //Cambio de entorno \n"
		codigo3d += "P = " + GetReferenciaOriginal(simboloArray) + "; //Cambio de entorno \n"
		//codigo3d += referenciaArray + " = P + " + strconv.Itoa(simboloArray.Direccion) + "; //Pos arr ref \n"
		codigo3d += referenciaArray + " = P + " + GetReferenciaOriginalPos(simboloArray) + "; //Pos arr ref \n"
		codigo3d += "P = " + scopeAnterior + "; //regresar al entorno anterior \n"
		//codigo3d += referenciaVector + " = " + strconv.Itoa(simboloVector.Direccion) + "; //Dir referencia\n"
	} else {
		if obj3dValor.PosId != "" && simboloArray.Entorno != nil {
			codigo3d += referenciaArray + " = " + obj3dValor.PosId + "; //Dir array\n"
		} else {
			codigo3d += referenciaArray + " = " + obj3dValor.Referencia + "; //Dir array\n"
		}

	}

	codigo3d += "/**************************DECLARACION DE ARRAY*/\n"
	if d.Stack {
		codigo3d += temp + " = P + " + strconv.Itoa(scope.ContadorDeclaracion) + ";\n"
		nSimbolo.Direccion = scope.ContadorDeclaracion
		nSimbolo.TipoDireccion = Ast.STACK
		scope.ContadorDeclaracion++
		codigo3d += "stack[(int)" + temp + "] = " + referenciaArray + ";\n"
	} else {
		codigo3d += temp + " = P + " + strconv.Itoa(scope.ContadorDeclaracion) + ";\n"
		nSimbolo.Direccion = scope.ContadorDeclaracion
		nSimbolo.TipoDireccion = Ast.HEAP
		scope.ContadorDeclaracion++
		codigo3d += "heap[(int)" + temp + "] = " + referenciaArray + ";\n"
	}
	codigo3d += "/***********************************************/\n"

	scope.Add(nSimbolo)
	obj3d.Codigo = codigo3d
	return Ast.TipoRetornado{Valor: obj3d, Tipo: Ast.DECLARACION}
}

func (op DeclaracionArray) GetFila() int {
	return op.Fila
}
func (op DeclaracionArray) GetColumna() int {
	return op.Columna
}

func DimensionesCorrectas(dimensiones arraylist.List, array interface{}, scope *Ast.Scope) Ast.TipoRetornado {
	dimension := dimensiones.GetValue(0).(Ast.TipoRetornado)
	arreglo := array.(expresiones.Array)
	//Validar dimensiones
	if arreglo.Size != dimension.Valor.(int) {
		////////////////////////////ERROR//////////////////////////////////
		return errores.GenerarError(25, arreglo, arreglo, "",
			"",
			"",
			scope)
		//////////////////////////////////////////////////////////////////
	}

	for i := 0; i < arreglo.Size; i++ {
		dimensiones.RemoveAtIndex(0)
		if dimensiones.Len() == 0 {
			return Ast.TipoRetornado{
				Valor: true,
				Tipo:  Ast.BOOLEAN,
			}
		}
		validacion := DimensionesCorrectas(dimensiones, arreglo.Elementos.GetValue(i).(Ast.TipoRetornado).Valor, scope)
		if validacion.Tipo == Ast.ERROR {
			return validacion
		}
	}
	return Ast.TipoRetornado{
		Valor: true,
		Tipo:  Ast.BOOLEAN,
	}

}

func EsArray(tipo Ast.TipoDato) bool {
	switch tipo {
	case Ast.ARRAY, Ast.ARRAY_ELEMENTOS, Ast.ARRAY_FAC:
		return true
	default:
		return false
	}
}

func CompararDimensiones(dim expresiones.DimensionArray, array expresiones.Array, scope *Ast.Scope) Ast.TipoRetornado {
	dimension := dim.GetValue(scope)
	//Comparar los tipos, error si no son del mismo tipo final
	if !expresiones.CompararTipos(dim.TipoArray, expresiones.GetTipoFinal(array.TipoDelArray)) {
		////////////////////////////ERROR//////////////////////////////////
		return errores.GenerarError(41, array, array, "",
			expresiones.Tipo_String(dim.TipoArray),
			Ast.ValorTipoDato[array.TipoArray],
			scope)
		//////////////////////////////////////////////////////////////////
	}

	//Verificar que las dimensiones concuerda con la lista de arrays
	validacionDimensiones := fn_array.ConcordanciaArray(array)
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
		arrayDimension.Add(dimension.Valor.(*arraylist.List).GetValue(i).(Ast.TipoRetornado).Valor)
	}

	if !fn_array.CompararListas(listaDimensiones, arrayDimension) {
		////////////////////////////ERROR//////////////////////////////////
		return errores.GenerarError(25, array, array, "",
			"",
			"",
			scope)
		//////////////////////////////////////////////////////////////////
	}

	return Ast.TipoRetornado{
		Tipo:  Ast.BOOLEAN,
		Valor: true,
	}

}

func CompararArrays(array1 expresiones.Array, array2 expresiones.Array, scope *Ast.Scope) bool {
	//Comparar las dimensiones
	if array1.Elementos.Len() != array2.Elementos.Len() {
		return false
	} else {
		elemento1 := array1.Elementos.GetValue(0)
		elemento2 := array2.Elementos.GetValue(0)
		tp1 := elemento1.(Ast.TipoRetornado).Tipo
		tp2 := elemento2.(Ast.TipoRetornado).Tipo
		if expresiones.EsTipoFinal(tp1) && expresiones.EsTipoFinal(tp2) {
			return true
		} else if !expresiones.EsTipoFinal(tp1) && !expresiones.EsTipoFinal(tp2) &&
			tp1 == Ast.ARRAY && tp2 == Ast.ARRAY {
			return CompararArrays(elemento1.(expresiones.Array), elemento2.(expresiones.Array), scope)
		} else {
			return false
		}
	}
}

func GetReferenciaOriginal(simbolo Ast.Simbolo) string {
	simboloActual := simbolo
	if simboloActual.Referencia {
		simboloActual = *simboloActual.Referencia_puntero
	}

	return strconv.Itoa(simboloActual.Entorno.Posicion)
}

func GetReferenciaOriginalPos(simbolo Ast.Simbolo) string {
	simboloActual := simbolo
	if simboloActual.Referencia {
		simboloActual = *simboloActual.Referencia_puntero
	}
	return strconv.Itoa(simboloActual.Direccion)
}
