package fn_vectores

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"
)

type ContainsVec struct {
	Identificador interface{}
	Tipo          Ast.TipoDato
	Valor         interface{}
	Fila          int
	Columna       int
}

func NewContainsVec(id interface{}, valor interface{}, tipo Ast.TipoDato, fila, columna int) ContainsVec {
	nP := ContainsVec{
		Identificador: id,
		Tipo:          tipo,
		Valor:         valor,
		Fila:          fila,
		Columna:       columna,
	}
	return nP
}

func (p ContainsVec) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	var simbolo Ast.Simbolo
	var vector expresiones.Vector
	var resultado = false
	var id string
	/*******VARIABLES PARA C3D********/
	var idExp expresiones.Identificador
	var obj3d, obj3Valor Ast.O3D
	var referencia, codigo3d string
	/*********************************/

	//Primero verificar que sea un identificador el id
	_, tipoParticular := p.Identificador.(Ast.Abstracto).GetTipo()
	if tipoParticular != Ast.IDENTIFICADOR {
		//Error se espera un identificador
		msg := "Semantic error, expected VECTOR, found " + Ast.ValorTipoDato[tipoParticular] +
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
	valor := p.Valor.(Ast.Expresion).GetValue(scope)
	obj3Valor = valor.Valor.(Ast.O3D)
	codigo3d += obj3Valor.Codigo
	valor = obj3Valor.Valor

	//Verificar que sean del mismo tipo

	if vector.Vacio {
		resultado = false
	} else if valor.Tipo != vector.TipoVector.Tipo {
		msg := "Semantic error, expected &" + expresiones.Tipo_String(vector.TipoVector) +
			", found &" + Ast.ValorTipoDato[valor.Tipo] + "." +
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
	} else if valor.Tipo == Ast.VECTOR {

		if !expresiones.CompararTipos(valor.Valor.(expresiones.Vector).TipoVector, vector.TipoVector) {
			//Error, no se puede guardar ese tipo de vector en este vector
			msg := "Semantic error, can't store " + expresiones.Tipo_String(valor.Valor.(expresiones.Vector).TipoVector) + " value" +
				" in a VEC< " + expresiones.Tipo_String(vector.TipoVector) + ">." +
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
		for i := 0; i < vector.Valor.Len(); i++ {
			//Primero verificar que sean del mismo tipo
			elemento := vector.Valor.GetValue(i).(Ast.TipoRetornado)
			if elemento.Valor == valor.Valor {
				resultado = true
				break
			}
		}
	} else {
		for i := 0; i < vector.Valor.Len(); i++ {
			//Primero verificar que sean del mismo tipo
			elemento := vector.Valor.GetValue(i).(Ast.TipoRetornado)
			if elemento.Valor == valor.Valor {
				resultado = true
				break
			}
		}
	}

	/*CODIGO 3D PARA AGREGAR EL ELEMENTO AL VECTOR*/
	/* PRIMERO CREAR UN NUEVO VECTOR Y AGREGARLE EL ELEMENTO DE ÚLTIMO */
	if valor.Tipo == Ast.STRING || valor.Tipo == Ast.STR {
		nReferencia, preCodigo3d := GetContainC3DString(referencia, obj3Valor.Referencia)
		codigo3d += preCodigo3d
		referencia = nReferencia
	} else {
		nReferencia, preCodigo3d := GetContainC3D(referencia, obj3Valor.Referencia)
		codigo3d += preCodigo3d
		referencia = nReferencia
	}

	obj3d.Valor = Ast.TipoRetornado{
		Tipo:  Ast.BOOLEAN,
		Valor: resultado,
	}
	obj3d.Codigo = codigo3d
	obj3d.Referencia = referencia

	return Ast.TipoRetornado{
		Tipo:  Ast.VEC_CONTAINS,
		Valor: obj3d,
	}
}

func (v ContainsVec) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, v.Tipo
}

func (v ContainsVec) GetFila() int {
	return v.Fila
}
func (v ContainsVec) GetColumna() int {
	return v.Columna
}

func GetContainC3D(referenciaVector, referenciaElemento string) (string, string) {
	codigo3d := ""
	resultado := Ast.GetTemp()
	posVectorActual := Ast.GetTemp()
	sizeVector := Ast.GetTemp()
	contador := Ast.GetTemp()
	elementoActual := Ast.GetTemp()
	temp1 := Ast.GetTemp()
	temp2 := Ast.GetTemp()
	lt := Ast.GetLabel()
	lf := Ast.GetLabel()
	lt2 := Ast.GetLabel()
	lf2 := Ast.GetLabel()
	salto := Ast.GetLabel()
	codigo3d += resultado + " = 0; //Inicializar el temp que guardará el resultado\n"
	codigo3d += sizeVector + " = heap[(int)" + referenciaVector + "];//Get el size actual\n"
	codigo3d += posVectorActual + " = " + referenciaVector + " + 1;//Get inicio del vector actual\n"
	codigo3d += contador + " = 0; //Inicializo contador\n"
	codigo3d += "/**********VERIFICAR POSICION ACTUAL DEL VECTOR*/\n"
	codigo3d += salto + ":\n"
	codigo3d += "if (" + contador + " < " + sizeVector + ") goto " + lt + ";\n"
	codigo3d += "goto " + lf + ";\n"
	codigo3d += "/******************************COMPARAR VALORES*/\n"
	codigo3d += lt + ":\n"
	codigo3d += elementoActual + " = " + "heap[(int)" + posVectorActual + "]; //Get valor elemento\n"
	codigo3d += "if (" + elementoActual + " != " + referenciaElemento + ") goto " + lt2 + ";\n"
	codigo3d += "goto " + lf2 + ";\n"
	codigo3d += lt2 + ":\n"
	codigo3d += temp1 + " = " + contador + " + 1; //Actualizar contador\n"
	codigo3d += contador + " = " + temp1 + ";\n"
	codigo3d += temp2 + " = " + posVectorActual + " + 1;//Actualizar pos vec actual\n"
	codigo3d += posVectorActual + " = " + temp2 + ";\n"
	codigo3d += "goto " + salto + ";\n"
	codigo3d += lf2 + ":\n"
	codigo3d += resultado + " = 1; //Si es igual\n"
	codigo3d += lf + ":\n"
	codigo3d += "/***********************************************/\n"
	return resultado, codigo3d
}

func GetContainC3DString(referenciaVector, referenciaElemento string) (string, string) {
	codigo3d := ""
	resultado := Ast.GetTemp()
	posVectorActual := Ast.GetTemp()
	sizeVector := Ast.GetTemp()
	contador := Ast.GetTemp()
	temp1 := Ast.GetTemp()
	temp2 := Ast.GetTemp()
	lt := Ast.GetLabel()
	lf := Ast.GetLabel()
	lt2 := Ast.GetLabel()
	lf2 := Ast.GetLabel()
	salto := Ast.GetLabel()
	codigo3d += resultado + " = 0; //Inicializar el temp que guardará el resultado\n"
	codigo3d += sizeVector + " = heap[(int)" + referenciaVector + "];//Get el size actual\n"
	codigo3d += posVectorActual + " = " + referenciaVector + " + 1;//Get inicio del vector actual\n"
	codigo3d += contador + " = 0; //Inicializo contador\n"
	codigo3d += "/**********VERIFICAR POSICION ACTUAL DEL VECTOR*/\n"
	codigo3d += salto + ":\n"
	codigo3d += "if (" + contador + " < " + sizeVector + ") goto " + lt + ";\n"
	codigo3d += "goto " + lf + ";\n"
	codigo3d += "/******************************COMPARAR VALORES*/\n"
	codigo3d += lt + ":\n"
	/*Obtener el codigo para comparar strings*/
	cod, ref := CompararStringC3DVector(posVectorActual, referenciaElemento)
	codigo3d += cod
	codigo3d += "if (" + ref + " == 0 ) goto " + lt2 + ";\n"
	codigo3d += "goto " + lf2 + ";\n"
	codigo3d += lt2 + ":\n"
	codigo3d += temp1 + " = " + contador + " + 1; //Actualizar contador\n"
	codigo3d += contador + " = " + temp1 + ";\n"
	codigo3d += temp2 + " = " + posVectorActual + " + 1;//Actualizar pos vec actual\n"
	codigo3d += posVectorActual + " = " + temp2 + ";\n"
	codigo3d += "goto " + salto + ";\n"
	codigo3d += lf2 + ":\n"
	codigo3d += resultado + " = 1; //Si es igual\n"
	codigo3d += lf + ":\n"
	codigo3d += "/***********************************************/\n"
	return resultado, codigo3d
}

func CompararStringC3DVector(string1, string2 string) (string, string) {
	codigo3d := ""
	resultado := Ast.GetTemp()
	contador1 := Ast.GetTemp()
	contador2 := Ast.GetTemp()
	caracter1 := Ast.GetTemp()
	caracter2 := Ast.GetTemp()
	sigPos1 := Ast.GetTemp()
	sigPos2 := Ast.GetTemp()
	lt := Ast.GetLabel()
	lf := Ast.GetLabel()
	lt2 := Ast.GetLabel()
	lf2 := Ast.GetLabel()
	salto := Ast.GetLabel()
	codigo3d += "/*********************COMPARACION DE STRINGS/STR*/\n"
	codigo3d += resultado + " = 1; //Inicializar temporal para guardar resultado\n"
	codigo3d += contador1 + " = heap[(int)" + string1 + "]; //Guardar pos string 1\n"
	codigo3d += contador2 + " = " + string2 + "; //Guardar pos string 2\n"
	codigo3d += salto + ":\n"
	codigo3d += caracter1 + " = heap[(int)" + contador1 + "]; //Guardar caracter1\n"
	codigo3d += caracter2 + " = heap[(int)" + contador2 + "]; //Guardar caracter2\n"
	codigo3d += "if (" + caracter1 + " != 0) goto " + lt + ";\n"
	codigo3d += "goto " + lf + ";\n"
	codigo3d += lt + ":\n"
	codigo3d += "if (" + caracter1 + " == " + caracter2 + ") goto " + lt2 + ";\n"
	codigo3d += "goto " + lf2 + ";\n"
	codigo3d += lt2 + ":\n"
	codigo3d += "/*****************************SIGUIENTE CARACTER*/\n"
	codigo3d += sigPos1 + " = " + contador1 + " + 1;\n"
	codigo3d += contador1 + " = " + sigPos1 + ";\n"
	codigo3d += sigPos2 + " = " + contador2 + " + 1;\n"
	codigo3d += contador2 + " = " + sigPos2 + ";\n"
	codigo3d += "goto " + salto + ";\n"
	codigo3d += lf2 + ":\n"
	codigo3d += resultado + " = 0;\n"
	codigo3d += lf + ":\n"
	codigo3d += "/***********************************************/\n"
	return codigo3d, resultado
}
