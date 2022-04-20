package fn_array

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"

	"github.com/colegno/arraylist"
)

type AccesoArray struct {
	Identificador interface{}
	Posiciones    *arraylist.List
	Tipo          Ast.TipoDato
	Fila          int
	Columna       int
}

func NewAccesoArray(id interface{}, posiciones *arraylist.List, fila, columna int) AccesoArray {
	nA := AccesoArray{
		Identificador: id,
		Tipo:          Ast.ACCESO_ARRAY,
		Posiciones:    posiciones,
		Fila:          fila,
		Columna:       columna,
	}
	return nA
}

func (p AccesoArray) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	var simbolo Ast.Simbolo
	var array interface{}
	var resultado Ast.TipoRetornado
	var id string
	var posicion interface{}
	var copiaLista *arraylist.List
	var valorPosicion Ast.TipoRetornado
	var idExp expresiones.Identificador
	var obj3d, obj3dValor Ast.O3D
	var referencia, codigo3d string
	posiciones := arraylist.New()
	posiciones3D := arraylist.New()
	//Primero verificar que sea un identificador el id
	_, tipoParticular := p.Identificador.(Ast.Abstracto).GetTipo()
	if tipoParticular != Ast.IDENTIFICADOR {
		return p.AccesoPorExpresion(scope).(Ast.TipoRetornado)
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

	/*Codigo 3d para conseguir el elemento del stack o del heap*/
	codigo3d += "/*********************************ACCESO A ARRAY*/\n"
	idExp = expresiones.NewIdentificador(id, Ast.IDENTIFICADOR, 0, 0)
	obj3dValor = idExp.GetValue(scope).Valor.(Ast.O3D)
	referencia = obj3dValor.Referencia
	codigo3d += obj3dValor.Codigo
	/************************************************************/

	//Verificar que sea un vector
	if simbolo.Tipo != Ast.ARRAY && simbolo.Tipo != Ast.VECTOR {
		msg := "Semantic error, expected " + Ast.ValorTipoDato[simbolo.Tipo] + ", found " + Ast.ValorTipoDato[simbolo.Tipo] + "." +
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
	if simbolo.Tipo == Ast.ARRAY {
		array = simbolo.Valor.(Ast.TipoRetornado).Valor.(expresiones.Array)
	} else {
		array = simbolo.Valor.(Ast.TipoRetornado).Valor.(expresiones.Vector)
	}
	codigo3d += "/*********************************GET POSICIONES*/\n"
	//Get las posiciones
	for i := 0; i < p.Posiciones.Len(); i++ {
		posicion = p.Posiciones.GetValue(i)
		valorPosicion = posicion.(Ast.Expresion).GetValue(scope)
		_, tipoParticular := posicion.(Ast.Abstracto).GetTipo()
		if valorPosicion.Tipo == Ast.ERROR {
			return posicion.(Ast.TipoRetornado)
		}
		//Verificar que el número en el acceso sea usize
		resultado := expresiones.EsUsize(valorPosicion, tipoParticular, posicion, scope)
		if resultado.Tipo == Ast.ERROR {
			return resultado
		}
		codigo3d += valorPosicion.Valor.(Ast.O3D).Codigo

		posiciones.Add(valorPosicion.Valor.(Ast.O3D).Valor.Valor)
		posiciones3D.Add(valorPosicion.Valor.(Ast.O3D))
	}
	codigo3d += "/************************************************/\n"
	//Buscar la posición
	//Copiar lista
	copiaLista = p.Posiciones.Clone()
	if simbolo.Tipo == Ast.ARRAY {
		resultado = GetElemento(array.(expresiones.Array), copiaLista, posiciones, scope, referencia, posiciones3D)
	} else {
		resultado = GetElementoVector(array.(expresiones.Vector), copiaLista, posiciones, scope, referencia, posiciones3D)
	}
	obj3d = resultado.Valor.(Ast.O3D)
	codigo3d += obj3d.Codigo
	codigo3d += "/***********************************************/\n"
	obj3d.Codigo = codigo3d
	return Ast.TipoRetornado{Tipo: Ast.ACCESO_ARRAY, Valor: obj3d}
}

func (p AccesoArray) AccesoPorExpresion(scope *Ast.Scope) interface{} {
	var resultado Ast.TipoRetornado
	var posicion interface{}
	var valorPosicion Ast.TipoRetornado
	var copiaLista *arraylist.List
	/*******************VARIABLES 3D ******************/
	var obj3d, obj3dValor, objTemp Ast.O3D
	var codigo3d string
	/*************************************************/

	posiciones := arraylist.New()
	posiciones3D := arraylist.New()
	//Conseguir el array
	array := p.Identificador.(Ast.Expresion).GetValue(scope)
	obj3dValor = array.Valor.(Ast.O3D)
	array = obj3dValor.Valor

	//Verificar que sea un array
	if array.Tipo != Ast.ARRAY {
		fila := p.Identificador.(Ast.Abstracto).GetFila()
		columna := p.Identificador.(Ast.Abstracto).GetColumna()
		msg := "Semantic error, expected ARRAY, found " + Ast.ValorTipoDato[array.Tipo] + "." +
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

	//Get las posiciones
	for i := 0; i < p.Posiciones.Len(); i++ {
		posicion = p.Posiciones.GetValue(i)
		valorPosicion = posicion.(Ast.Expresion).GetValue(scope)
		objTemp = valorPosicion.Valor.(Ast.O3D)
		valorPosicion = objTemp.Valor
		_, tipoParticular := posicion.(Ast.Abstracto).GetTipo()
		if valorPosicion.Tipo == Ast.ERROR {
			return posicion.(Ast.TipoRetornado)
		}
		//Verificar que el número en el acceso sea usize
		resultado := expresiones.EsUsize(valorPosicion, tipoParticular, posicion, scope)
		if resultado.Tipo == Ast.ERROR {
			return resultado
		}
		posiciones.Add(valorPosicion.Valor)
		posiciones3D.Add(objTemp)
	}

	//Buscar la posición
	//Copiar lista
	copiaLista = p.Posiciones.Clone()
	resultado = GetElemento(array.Valor.(expresiones.Array), copiaLista, posiciones, scope, obj3dValor.Referencia, posiciones3D)
	codigo3d += obj3dValor.Codigo
	codigo3d += resultado.Valor.(Ast.O3D).Codigo
	obj3d.Codigo = codigo3d
	obj3d.Valor = resultado.Valor.(Ast.O3D).Valor
	obj3d.Referencia = resultado.Valor.(Ast.O3D).Referencia
	return Ast.TipoRetornado{
		Tipo:  Ast.ACCESO_ARRAY,
		Valor: obj3d,
	}
}

func (v AccesoArray) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, v.Tipo
}

func (v AccesoArray) GetFila() int {
	return v.Fila
}
func (v AccesoArray) GetColumna() int {
	return v.Columna
}

//p.Posiciones
func GetElemento(array expresiones.Array, elementos *arraylist.List, posiciones *arraylist.List,
	scope *Ast.Scope, ref string, posiciones3D *arraylist.List) Ast.TipoRetornado {
	/***********VARIABLES 3D********/
	var obj3d, obj3dValor Ast.O3D
	var referencia string
	var codigo3d string
	var lf string
	referencia = ref
	/*******************************/

	posicion := posiciones.GetValue(0).(int)
	elemento := elementos.GetValue(0)
	posicion3D := posiciones3D.GetValue(0)
	posiciones.RemoveAtIndex(0)
	elementos.RemoveAtIndex(0)
	posiciones3D.RemoveAtIndex(0)
	if posicion >= array.Size || posicion < 0 {
		//Error, out of bounds
		fila := elemento.(Ast.Abstracto).GetFila()
		columna := elemento.(Ast.Abstracto).GetColumna()
		msg := "Semantic error, index (" + strconv.Itoa(posicion) + ") out of bounds." +
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
	if posiciones.Len() == 0 {
		//Si es 0, entonces retornar la posición actual
		valor := array.Elementos.GetValue(posicion).(Ast.TipoRetornado)
		if valor.Tipo == Ast.VECTOR {
			obj3dValor = GetCod3dAccesoArray(referencia, posicion, true, posicion3D.(Ast.O3D))
		} else {
			obj3dValor = GetCod3dAccesoArray(referencia, posicion, false, posicion3D.(Ast.O3D))
		}
		codigo3d += obj3dValor.Codigo
		codigo3d += obj3dValor.Lt + ":\n"
		codigo3d += "/***********************************************/\n"
		obj3d.Codigo = codigo3d
		obj3d.Valor = valor
		obj3d.Referencia = obj3dValor.Referencia
		return Ast.TipoRetornado{Valor: obj3d, Tipo: obj3d.Valor.Tipo}
	}
	if posiciones.Len() > 0 && array.TipoArray != Ast.ARRAY {
		//Error, no hay más dimensiones
		fila := elemento.(Ast.Abstracto).GetFila()
		columna := elemento.(Ast.Abstracto).GetColumna()
		msg := "Semantic error, index (" + strconv.Itoa(posicion) + ") out of bounds." +
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
	next := array.Elementos.GetValue(posicion).(Ast.TipoRetornado)
	/*Codigo 3d para el siguiente elemento vector*/
	obj3dValor = GetCod3dAccesoVector(referencia, posicion, true, posicion3D.(Ast.O3D))
	codigo3d += obj3dValor.Codigo
	lf = Ast.GetLabel()
	codigo3d += "goto " + lf + ";\n"
	codigo3d += obj3dValor.Lt + ":\n"
	//Validar que el siguiente sea un array y que todavía existan posiciones que buscar
	obj3d = GetElemento(next.Valor.(expresiones.Array), elementos, posiciones, scope, obj3dValor.Referencia, posiciones3D).Valor.(Ast.O3D)
	codigo3d += obj3d.Codigo
	codigo3d += lf + ":\n"
	codigo3d += "/***********************************************/\n"
	obj3d.Codigo = codigo3d
	return Ast.TipoRetornado{Tipo: Ast.ACCESO_ARRAY, Valor: obj3d}
}

//Get elemento vector

func GetElementoVector(array expresiones.Vector, elementos *arraylist.List, posiciones *arraylist.List,
	scope *Ast.Scope, ref string, posiciones3D *arraylist.List) Ast.TipoRetornado {
	var obj3d, obj3dValor Ast.O3D
	var referencia string
	var codigo3d string
	var lf string
	referencia = ref
	posicion := posiciones.GetValue(0).(int)
	elemento := elementos.GetValue(0)
	posicion3D := posiciones3D.GetValue(0)
	posiciones.RemoveAtIndex(0)
	elementos.RemoveAtIndex(0)
	posiciones3D.RemoveAtIndex(0)
	if posicion >= array.Size || posicion < 0 {
		//Error, out of bounds
		fila := elemento.(Ast.Abstracto).GetFila()
		columna := elemento.(Ast.Abstracto).GetColumna()
		msg := "Semantic error, index (" + strconv.Itoa(posicion) + ") out of bounds." +
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
	if posiciones.Len() == 0 {
		//Si es 0, entonces retornar la posición actual
		valor := array.Valor.GetValue(posicion).(Ast.TipoRetornado)
		if valor.Tipo == Ast.VECTOR {
			obj3dValor = GetCod3dAccesoVector(referencia, posicion, true, posicion3D.(Ast.O3D))
		} else {
			obj3dValor = GetCod3dAccesoVector(referencia, posicion, false, posicion3D.(Ast.O3D))
		}
		codigo3d += obj3dValor.Codigo
		codigo3d += obj3dValor.Lt + ":\n"
		codigo3d += "/***********************************************/\n"
		obj3d.Codigo = codigo3d
		obj3d.Valor = valor
		obj3d.Referencia = obj3dValor.Referencia

		return Ast.TipoRetornado{Valor: obj3d, Tipo: obj3d.Valor.Tipo}
	}
	if posiciones.Len() > 0 && array.TipoVector.Tipo != Ast.VECTOR {
		//Error, no hay más dimensiones
		fila := elemento.(Ast.Abstracto).GetFila()
		columna := elemento.(Ast.Abstracto).GetColumna()
		msg := "Semantic error, index (" + strconv.Itoa(posicion) + ") out of bounds." +
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
	/*Codigo 3d para el siguiente elemento vector*/

	next := array.Valor.GetValue(posicion).(Ast.TipoRetornado)
	obj3dValor = GetCod3dAccesoVector(referencia, posicion, true, posicion3D.(Ast.O3D))
	codigo3d += obj3dValor.Codigo
	lf = Ast.GetLabel()
	codigo3d += "goto " + lf + ";\n"
	codigo3d += obj3dValor.Lt + ":\n"
	//Validar que el siguiente sea un array y que todavía existan posiciones que buscar
	obj3d = GetElementoVector(next.Valor.(expresiones.Vector), elementos, posiciones, scope, obj3dValor.Referencia, posiciones3D).Valor.(Ast.O3D)
	codigo3d += obj3d.Codigo
	codigo3d += lf + ":\n"
	codigo3d += "/***********************************************/\n"
	obj3d.Codigo = codigo3d
	return Ast.TipoRetornado{Tipo: Ast.ACCESO_ARRAY, Valor: obj3d}
}

func GetCod3dAccesoVector(ref string, pos int, estructura bool, posRef Ast.O3D) Ast.O3D {
	var obj3d Ast.O3D
	var codigo3d string
	temp := ""
	temp2 := ""
	temp3 := ""
	temp4 := ""
	temp5 := ""
	lt := ""
	lf := ""
	salto := ""
	referencia := ref
	referenciaPosicion := ""
	//posicion := pos
	/*Trabajar código 3d*/
	lt = Ast.GetLabel()
	lf = Ast.GetLabel()
	salto = Ast.GetLabel()

	/***************CODIGO 3D******************/
	if posRef.DireccionID != "" {
		referenciaPosicion = posRef.DireccionID
	} else {
		referenciaPosicion = posRef.Referencia
	}

	temp = Ast.GetTemp()  //Temporal para guardar el size del vector
	temp2 = Ast.GetTemp() //Temporal que va a guardar el size del vector
	temp3 = Ast.GetTemp()
	temp4 = Ast.GetTemp()
	codigo3d += "/************************GET ELEMENTO DE VECTOR*/\n"
	codigo3d += temp + " = heap[(int)" + referencia + "]; //Get size\n"
	codigo3d += "if (" + referenciaPosicion + " < " + temp + ") goto " + lt + ";\n"
	codigo3d += "goto " + lf + ";\n"
	codigo3d += lt + ":\n"
	codigo3d += temp2 + " = " + referencia + " + 1; //Get inicio del vector\n"
	codigo3d += temp3 + " = " + referenciaPosicion + "; //Get posicion exacta\n"
	codigo3d += temp4 + " = " + temp2 + " + " + temp3 + "; //Get elemento\n"
	/*Si es una estructura, obtener su direccion*/
	if estructura {
		temp5 = Ast.GetTemp()
		codigo3d += temp5 + " = " + "heap[(int)" + temp4 + "];//Get posicion exacta de objeto\n"
		obj3d.Referencia = temp5
	} else {
		temp5 = Ast.GetTemp()
		codigo3d += temp5 + " = " + "heap[(int)" + temp4 + "];//Get posicion exacta de objeto\n"
		obj3d.Referencia = temp5
	}
	codigo3d += "goto " + salto + ";\n"
	codigo3d += BoundsError(lf)
	//codigo3d += salto + ":\n"
	/******************************************/
	obj3d.Lt = salto
	obj3d.Codigo = codigo3d
	return obj3d
}

func GetCod3dAccesoArray(ref string, pos int, estructura bool, posRef Ast.O3D) Ast.O3D {
	var obj3d Ast.O3D
	var codigo3d string
	temp := ""
	temp2 := ""
	temp3 := ""
	temp4 := ""
	temp5 := ""
	lt := ""
	lf := ""
	salto := ""
	referencia := ref
	referenciaPosicion := ""
	//posicion := pos
	/*Trabajar código 3d*/
	lt = Ast.GetLabel()
	lf = Ast.GetLabel()
	salto = Ast.GetLabel()

	/***************CODIGO 3D******************/
	if posRef.DireccionID != "" {
		referenciaPosicion = posRef.DireccionID
	} else {
		referenciaPosicion = posRef.Referencia
	}

	temp = Ast.GetTemp()  //Temporal para guardar el size del vector
	temp2 = Ast.GetTemp() //Temporal que va a guardar el size del vector
	temp3 = Ast.GetTemp()
	temp4 = Ast.GetTemp()
	codigo3d += "/*************************GET ELEMENTO DE ARRAY*/\n"
	codigo3d += temp + " = heap[(int)" + referencia + "]; //Get size\n"
	codigo3d += "if (" + referenciaPosicion + " < " + temp + ") goto " + lt + ";\n"
	codigo3d += "goto " + lf + ";\n"
	codigo3d += lt + ":\n"
	codigo3d += temp2 + " = " + referencia + " + 1; //Get inicio del array\n"
	codigo3d += temp3 + " = " + referenciaPosicion + "; //Get posicion exacta\n"
	codigo3d += temp4 + " = " + temp2 + " + " + temp3 + "; //Get elemento\n"
	/*Si es una estructura, obtener su direccion*/
	if estructura {
		temp5 = Ast.GetTemp()
		codigo3d += temp5 + " = " + "heap[(int)" + temp4 + "];//Get posicion exacta de objeto\n"
		obj3d.Referencia = temp5
	} else {
		temp5 = Ast.GetTemp()
		codigo3d += temp5 + " = " + "heap[(int)" + temp4 + "];//Get posicion exacta de objeto\n"
		obj3d.Referencia = temp5
	}
	codigo3d += "goto " + salto + ";\n"
	codigo3d += BoundsError(lf)
	//codigo3d += salto + ":\n"
	/******************************************/
	obj3d.Lt = salto
	obj3d.Codigo = codigo3d
	return obj3d
}

func UpdateElemento(array expresiones.Array, elementos *arraylist.List,
	posiciones *arraylist.List, scope *Ast.Scope, objeto interface{},
	refVector string, posiciones3D *arraylist.List) Ast.TipoRetornado {
	var obj3dValor, obj3d Ast.O3D
	var resultado Ast.TipoRetornado
	codigo3d := ""
	posVector := refVector
	sizeVector := ""
	primeraPos := ""
	sigVectorPos := ""
	posAsignar := ""
	sigVector := ""
	var lt, lf, salto string

	posicion := posiciones.GetValue(0).(int)
	posicion3D := posiciones3D.GetValue(0).(string)
	elemento := elementos.GetValue(0)
	posiciones.RemoveAtIndex(0)
	elementos.RemoveAtIndex(0)
	posiciones3D.RemoveAtIndex(0)
	if posicion >= array.Size || posicion < 0 {
		//Error, out of bounds
		fila := elemento.(Ast.Abstracto).GetFila()
		columna := elemento.(Ast.Abstracto).GetColumna()
		msg := "Semantic error, index (" + strconv.Itoa(posicion) + ") out of bounds." +
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
	if posiciones.Len() == 0 {
		//Actualizar el elemento simpre y cuando no este en un tipo array en la siguiente columna
		/*********************************************************/
		next := array.Elementos.GetValue(posicion).(Ast.TipoRetornado)
		//obj3dValor = next.Valor.(Ast.O3D)
		//next = obj3dValor.Valor
		/*********************************************************/

		if next.Tipo == Ast.ARRAY {
			//No se puede guardar en esa posición porque es posición de array
			fila := elemento.(Ast.Abstracto).GetFila()
			columna := elemento.(Ast.Abstracto).GetColumna()
			msg := "Semantic error, index (" + strconv.Itoa(posicion) + "). Can't access to that position." +
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

		nuevaLista := *arraylist.New()
		for i := 0; i < array.Elementos.Len(); i++ {
			if i == posicion {
				//Reemplazar el elemento
				/*******************************************/
				obj3dValor = objeto.(Ast.O3D)
				codigo3d += obj3dValor.Codigo
				valor := obj3dValor.Valor
				nuevaLista.Add(valor)
				/*******************************************/
				continue
			}
			nuevaLista.Add(array.Elementos.GetValue(i))
		}

		sizeVector = Ast.GetTemp()
		primeraPos = Ast.GetTemp()
		posAsignar = Ast.GetTemp()
		lt = Ast.GetLabel()
		lf = Ast.GetLabel()
		salto = Ast.GetLabel()
		codigo3d += sizeVector + " = heap[(int)" + posVector + "]; //Get size array\n"
		codigo3d += primeraPos + " = " + posVector + " + 1; //Primera pos \n"
		codigo3d += posAsignar + " = " + primeraPos + " + " + posicion3D + "; //Pos asignar\n"
		codigo3d += "if (" + posicion3D + " < " + sizeVector + ") goto " + lt + ";\n"
		codigo3d += "goto " + lf + ";\n"
		codigo3d += lt + ":\n"
		codigo3d += "heap[(int)" + posAsignar + "] = " + obj3dValor.Referencia + "; //Add nuevo valor\n"
		codigo3d += "goto " + salto + ";\n"
		codigo3d += BoundsError(lf)
		codigo3d += salto + ":\n"
		array.Elementos.Clear()
		for i := 0; i < nuevaLista.Len(); i++ {
			array.Elementos.Add(nuevaLista.GetValue(i))
		}
		obj3d.Valor = Ast.TipoRetornado{Valor: true, Tipo: Ast.EJECUTADO}
		obj3d.Codigo = codigo3d
		return Ast.TipoRetornado{Valor: obj3d, Tipo: Ast.EJECUTADO}
		//return Ast.TipoRetornado{Valor: true, Tipo: Ast.EJECUTADO}
	}
	if posiciones.Len() > 0 && array.TipoArray != Ast.ARRAY {
		//Error, no hay más dimensiones
		fila := elemento.(Ast.Abstracto).GetFila()
		columna := elemento.(Ast.Abstracto).GetColumna()
		msg := "Semantic error, index (" + strconv.Itoa(posicion) + ") out of bounds." +
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
	next := array.Elementos.GetValue(posicion).(Ast.TipoRetornado)
	valorNext := next.Valor.(expresiones.Array)
	//Validar que el siguiente sea un array y que todavía existan posiciones que buscar

	/***************************************************/
	sizeVector = Ast.GetTemp()
	primeraPos = Ast.GetTemp()
	sigVectorPos = Ast.GetTemp()
	sigVector = Ast.GetTemp()
	lt = Ast.GetLabel()
	lf = Ast.GetLabel()
	salto = Ast.GetLabel()
	codigo3d += sizeVector + " = heap[(int)" + posVector + "]; //Get size array\n"
	codigo3d += primeraPos + " = " + posVector + " + 1; //Primera pos \n"
	codigo3d += sigVectorPos + " = " + primeraPos + " + " + posicion3D + "; //Get posicion proxima dimension\n"
	codigo3d += sigVector + " = " + "heap[(int)" + sigVectorPos + "]; //Get proxima dimension \n"
	codigo3d += "if (" + posicion3D + " < " + sizeVector + ") goto " + lt + ";\n"
	codigo3d += "goto " + lf + ";\n"
	codigo3d += lt + ":\n"
	/***************************************************/
	//Validar que el siguiente sea un array y que todavía existan posiciones que buscar
	resultado = UpdateElemento(valorNext, elementos, posiciones, scope, objeto, sigVector, posiciones3D)
	/***************************************************/
	codigo3d += resultado.Valor.(Ast.O3D).Codigo
	codigo3d += "goto " + salto + ";\n"
	codigo3d += BoundsError(lf)
	codigo3d += salto + ":\n"
	obj3d.Valor = Ast.TipoRetornado{Valor: true, Tipo: Ast.EJECUTADO}
	obj3d.Codigo = codigo3d
	/***************************************************/
	return Ast.TipoRetornado{Valor: obj3d, Tipo: Ast.EJECUTADO}
}

func BoundsError(lf string) string {
	codigo3d := ""
	codigo3d += lf + ":\n"
	codigo3d += "printf(\"%%c\", 66);\n"
	codigo3d += "printf(\"%%c\", 111);\n"
	codigo3d += "printf(\"%%c\", 117);\n"
	codigo3d += "printf(\"%%c\", 110);\n"
	codigo3d += "printf(\"%%c\", 100);\n"
	codigo3d += "printf(\"%%c\", 115);\n"
	codigo3d += "printf(\"%%c\", 64);\n"
	codigo3d += "printf(\"%%c\", 114);\n"
	codigo3d += "printf(\"%%c\", 114);\n"
	codigo3d += "printf(\"%%c\", 111);\n"
	codigo3d += "printf(\"%%c\", 114);\n"
	codigo3d += "printf(\"\\n\");\n"
	return codigo3d
}
