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
	_, tipoIn := d.Valor.(Ast.Abstracto).GetTipo()

	if tipoIn == Ast.VALOR {
		existe = d.ScopeOriginal.Exist_actual(d.Id)
		valor = d.Valor.(Ast.Expresion).GetValue(d.ScopeOriginal)
	} else {
		existe = scope.Exist_actual(d.Id)
		valor = d.Valor.(Ast.Expresion).GetValue(scope)
	}
	dimension := d.Dimension.(Ast.Expresion).GetValue(scope)

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
			msg := "Semantic error, can't initialize a Vec<" + expresiones.Tipo_String(d.TipoArray) +
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

	//Recuperar el tipo del array que se espera desde dimension
	//d.TipoArray = d.Dimension.(expresiones.DimensionArray).TipoArray

	//Primero que vengan arrays
	if !EsArray(tipoIn) {
		//Error, no se estan asignado arrays al array
		msg := "Semantic error, can't initialize an ARRAY with " + Ast.ValorTipoDato[tipoIn] + " type" +
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
		arrayDimension.Add(dimension.Valor.(*arraylist.List).GetValue(i).(Ast.TipoRetornado).Valor)
	}

	if !fn_array.CompararListas(listaDimensiones, arrayDimension) {
		msg := "Semantic error, ARRAY dimension does not match" +
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
	//expresiones.GetTipoFinal(valor.Valor.(expresiones.Array).TipoDelArray.Tipo)
	//Validar el tipo del array
	if d.TipoArray.Tipo != expresiones.GetTipoFinal(valor.Valor.(expresiones.Array).TipoDelArray).Tipo {
		fila := valor.Valor.(expresiones.Array).GetFila()
		columna := valor.Valor.(expresiones.Array).GetColumna()
		var tipoDelArray string
		if valor.Valor.(expresiones.Array).TipoDelArray.Tipo == Ast.INDEFINIDO {
			tipoDelArray = Ast.ValorTipoDato[d.TipoArray.Tipo]
		} else {
			tipoDelArray = expresiones.Tipo_String(expresiones.GetTipoFinal(d.TipoArray))
		}
		msg := "Semantic error, can't initialize ARRAY[" + tipoDelArray +
			"] with a ARRAY[" +
			Ast.ValorTipoDato[expresiones.GetTipoFinal(valor.Valor.(expresiones.Array).TipoDelArray).Tipo] + "]" +
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

	//Crear el símbolo
	nSimbolo := Ast.Simbolo{
		Identificador: d.Id,
		Valor:         valor,
		Fila:          d.Fila,
		Columna:       d.Columna,
		Tipo:          valor.Tipo,
		Mutable:       d.Mutable,
		Publico:       d.Publico,
	}
	//Actualizar la mutabilidad de la instancia
	nArray := valor.Valor.(expresiones.Array)
	nArray.Mutable = d.Mutable
	nValor := Ast.TipoRetornado{
		Tipo:  Ast.ARRAY,
		Valor: nArray,
	}
	nSimbolo.Valor = nValor
	scope.Add(nSimbolo)
	return Ast.TipoRetornado{
		Tipo:  Ast.EJECUTADO,
		Valor: true,
	}
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
		fila := arreglo.GetFila()
		columna := arreglo.GetColumna()
		msg := "Semantic error, Array dimensions don't match " +
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
		fila := array.Fila
		columna := array.Columna
		msg := "Semantic error, expected  ARRAY[" + expresiones.Tipo_String(dim.TipoArray) + "] " +
			" found ARRAY[" + Ast.ValorTipoDato[array.TipoArray] + "]." +
			" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
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
		fila := array.Fila
		columna := array.Columna
		msg := "Semantic error, ARRAY dimension does not match" +
			" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
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
