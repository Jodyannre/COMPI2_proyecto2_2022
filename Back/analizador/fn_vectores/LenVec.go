package fn_vectores

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"
)

type LenVec struct {
	Identificador interface{}
	Tipo          Ast.TipoDato
	Fila          int
	Columna       int
}

func NewLenVec(id interface{}, tipo Ast.TipoDato, fila, columna int) LenVec {
	nP := LenVec{
		Identificador: id,
		Tipo:          tipo,
		Fila:          fila,
		Columna:       columna,
	}
	return nP
}

func (p LenVec) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	var simbolo Ast.Simbolo
	var vector interface{}
	var id string
	var idExp expresiones.Identificador
	var obj3d, obj3dValor Ast.O3D
	var temp string
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
	referencia = obj3dValor.Referencia
	codigo3d += obj3dValor.Codigo

	/************************************************************/

	//Verificar que sea un vector
	if simbolo.Tipo != Ast.VECTOR && simbolo.Tipo != Ast.ARRAY {
		msg := "Semantic error, expected (Vector|Array), found " + Ast.ValorTipoDato[simbolo.Tipo] + "." +
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
		/*Codigo 3D*/
		temp = Ast.GetTemp()
		codigo3d += "/********************************LEN DEL VECTOR*/\n"
		codigo3d += temp + " = heap[(int)" + referencia + "]; //Get size\n"
		codigo3d += "/***********************************************/\n"
		codigo3d += "/***********************************************/\n"
		vector = simbolo.Valor.(Ast.TipoRetornado).Valor.(expresiones.Vector)
		obj3d.Codigo = codigo3d
		obj3d.Referencia = temp
		obj3d.Valor = Ast.TipoRetornado{
			Tipo:  Ast.I64,
			Valor: vector.(expresiones.Vector).Size,
		}
		return Ast.TipoRetornado{
			Tipo:  Ast.VEC_LEN,
			Valor: obj3d,
		}
	} else {
		vector = simbolo.Valor.(Ast.TipoRetornado).Valor.(expresiones.Array)
		/*Codigo 3D*/
		temp = Ast.GetTemp()
		codigo3d += "/*********************************LEN DEL ARRAY*/\n"
		codigo3d += temp + " = heap[(int)" + referencia + "]; //Get size\n"
		codigo3d += "/***********************************************/\n"
		codigo3d += "/***********************************************/\n"
		obj3d.Codigo = codigo3d
		obj3d.Referencia = temp
		obj3d.Valor = Ast.TipoRetornado{
			Tipo:  Ast.I64,
			Valor: vector.(expresiones.Array).Size,
		}
		return Ast.TipoRetornado{
			Tipo:  Ast.VEC_LEN,
			Valor: obj3d,
		}
	}

}

func (v LenVec) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, v.Tipo
}

func (v LenVec) GetFila() int {
	return v.Fila
}
func (v LenVec) GetColumna() int {
	return v.Columna
}
