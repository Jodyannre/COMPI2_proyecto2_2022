package fn_vectores

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"

	"github.com/colegno/arraylist"
)

type AccesoVec struct {
	Identificador interface{}
	Posicion      interface{}
	Tipo          Ast.TipoDato
	Fila          int
	Columna       int
}

func NewAccesoVec(id interface{}, posicion interface{}, tipo Ast.TipoDato, fila, columna int) AccesoVec {
	nA := AccesoVec{
		Identificador: id,
		Tipo:          tipo,
		Posicion:      posicion,
		Fila:          fila,
		Columna:       columna,
	}
	return nA
}

func (p AccesoVec) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	var simbolo Ast.Simbolo
	var vector interface{}
	var posicion Ast.TipoRetornado
	var resultado Ast.TipoRetornado
	var id string
	var idExp expresiones.Identificador
	var obj3d, obj3dValor Ast.O3D
	var referencia, codigo3d string
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

	/*Codigo 3d para conseguir el elemento del stack o del heap*/
	codigo3d += "/********************************ACCESO A VECTOR*/\n"
	idExp = expresiones.NewIdentificador(id, Ast.IDENTIFICADOR, 0, 0)
	obj3dValor = idExp.GetValue(scope).Valor.(Ast.O3D)
	codigo3d += obj3dValor.Codigo
	codigo3d += obj3d.Codigo
	/************************************************************/

	//Verificar que sea un vector
	if simbolo.Tipo != Ast.VECTOR && simbolo.Tipo != Ast.ARRAY {
		msg := "Semantic error, expected (VECTOR|ARRAY), found " + Ast.ValorTipoDato[simbolo.Tipo] + "." +
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
	if simbolo.Tipo == Ast.VECTOR {
		vector = simbolo.Valor.(Ast.TipoRetornado).Valor.(expresiones.Vector)
	} else {
		vector = simbolo.Valor.(Ast.TipoRetornado).Valor.(expresiones.Array)
	}

	//Get la posición de donde se va a extraer el elemento
	posicion = p.Posicion.(Ast.Expresion).GetValue(scope)
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
	if simbolo.Tipo == Ast.VECTOR {
		temp := ""
		temp2 := ""
		temp3 := ""
		temp4 := ""
		lt := ""
		lf := ""
		salto := ""
		if posicion.Valor.(int) >= vector.(expresiones.Vector).Size {
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

		//Acceder al elemento
		resultado = vector.(expresiones.Vector).Valor.GetValue(posicion.Valor.(int)).(Ast.TipoRetornado)

		/*Trabajar código 3d*/
		lt = Ast.GetLabel()
		lf = Ast.GetLabel()
		salto = Ast.GetLabel()

		/***************CODIGO 3D******************/
		temp = Ast.GetTemp()  //Temporal para guardar el size del vector
		temp2 = Ast.GetTemp() //Temporal que va a guardar el size del vector
		temp3 = Ast.GetTemp()
		temp4 = Ast.GetTemp()
		referencia = obj3dValor.Referencia
		codigo3d += "/********************************GET SIZE VECTOR*/\n"
		codigo3d += temp + " = heap[(int)" + referencia + "]; //Get size\n"
		codigo3d += "if (" + strconv.Itoa(posicion.Valor.(int)) + " < " + temp + ") goto " + lt + ";\n"
		codigo3d += "goto " + lf + ";\n"
		codigo3d += lt + ":\n"
		codigo3d += temp2 + " = " + referencia + " + 1; //Get inicio del vector\n"
		codigo3d += temp3 + " = " + strconv.Itoa(posicion.Valor.(int)) + "; //Get posicion exacta\n"
		codigo3d += temp4 + " = " + temp2 + " + " + temp3 + "; //Get elemento\n"
		codigo3d += "goto " + salto + ";\n"
		codigo3d += BoundsError(lf)
		codigo3d += salto + ":\n"
		codigo3d += "/***********************************************/\n"
		referencia = temp4
		obj3d.Valor = resultado
		/******************************************/

	}

	if simbolo.Tipo == Ast.ARRAY {
		if posicion.Valor.(int) >= vector.(expresiones.Array).Size {
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

		//Acceder al elemento
		resultado = vector.(expresiones.Array).Elementos.GetValue(posicion.Valor.(int)).(Ast.TipoRetornado)
	}
	/* Retornar el OBJ3D*/
	obj3d.Codigo += codigo3d
	return Ast.TipoRetornado{
		Tipo:  Ast.ACCESO_VECTOR,
		Valor: obj3d,
	}
}

func UpdateElemento(array expresiones.Vector, elementos *arraylist.List, posiciones *arraylist.List,
	scope *Ast.Scope, objeto interface{}, refVector string, posiciones3D *arraylist.List) Ast.TipoRetornado {
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

		//Actualizar el elemento siempre y cuando no sea una capa vector
		/*********************************************************/
		next := array.Valor.GetValue(posicion).(Ast.TipoRetornado)
		//obj3dValor = next.Valor.(Ast.O3D)
		//next = obj3dValor.Valor
		/*********************************************************/
		if next.Tipo == Ast.VECTOR {
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
		for i := 0; i < array.Valor.Len(); i++ {
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
			nuevaLista.Add(array.Valor.GetValue(i))
		}
		sizeVector = Ast.GetTemp()
		primeraPos = Ast.GetTemp()
		posAsignar = Ast.GetTemp()
		lt = Ast.GetLabel()
		lf = Ast.GetLabel()
		salto = Ast.GetLabel()
		codigo3d += sizeVector + " = heap[(int)" + posVector + "]; //Get size vector\n"
		codigo3d += primeraPos + " = " + posVector + " + 1; //Primera pos \n"
		codigo3d += posAsignar + " = " + primeraPos + " + " + posicion3D + "; //Pos asignar\n"
		codigo3d += "if (" + posicion3D + " < " + sizeVector + ") goto " + lt + ";\n"
		codigo3d += "goto " + lf + ";\n"
		codigo3d += lt + ":\n"
		codigo3d += "heap[(int)" + posAsignar + "] = " + obj3dValor.Referencia + "; //Add nuevo valor\n"
		codigo3d += "goto " + salto + ";\n"
		codigo3d += BoundsError(lf)
		codigo3d += salto + ":\n"
		array.Valor.Clear()
		for i := 0; i < nuevaLista.Len(); i++ {
			array.Valor.Add(nuevaLista.GetValue(i))
		}
		obj3d.Valor = Ast.TipoRetornado{Valor: true, Tipo: Ast.EJECUTADO}
		obj3d.Codigo = codigo3d
		return Ast.TipoRetornado{Valor: obj3d, Tipo: Ast.EJECUTADO}
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
	next := array.Valor.GetValue(posicion).(Ast.TipoRetornado)
	valorNext := next.Valor.(expresiones.Vector)

	/***************************************************/
	sizeVector = Ast.GetTemp()
	primeraPos = Ast.GetTemp()
	sigVectorPos = Ast.GetTemp()
	sigVector = Ast.GetTemp()
	lt = Ast.GetLabel()
	lf = Ast.GetLabel()
	salto = Ast.GetLabel()
	codigo3d += sizeVector + " = heap[(int)" + posVector + "]; //Get size vector\n"
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

func (v AccesoVec) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, v.Tipo
}

func (v AccesoVec) GetFila() int {
	return v.Fila
}
func (v AccesoVec) GetColumna() int {
	return v.Columna
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
