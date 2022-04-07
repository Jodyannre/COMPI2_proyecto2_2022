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
	var codigo3d string
	var obj3d, obj3dValor Ast.O3D
	/*************************************/
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

	//Clonar la lista para evitar la referencia
	nArray := valor.Valor.(Ast.Clones).Clonar(scope)
	//Actualizar la mutabilidad de la instancia
	mnArray := nArray.(expresiones.Array)
	mnArray.Mutable = d.Mutable
	nArray = mnArray
	nSimbolo.Valor = Ast.TipoRetornado{
		Tipo:  valor.Tipo,
		Valor: nArray,
	}
	temp := Ast.GetTemp()
	codigo3d += obj3dValor.Codigo
	codigo3d += "/**************************DECLARACION DE ARRAY*/\n"
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
		scope.Size++
		codigo3d += "heap[(int)" + temp + "] = " + obj3dValor.Referencia + ";\n"
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
