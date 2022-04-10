package simbolos

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"
)

type DeclaracionFuncion struct {
	Id      string
	Mutable bool
	Publico bool
	Tipo    Ast.TipoDato
	Retorno Ast.TipoRetornado
	Valor   interface{}
	Fila    int
	Columna int
}

func NewDeclaracionFuncion(id string, valor interface{}, tipo Ast.TipoRetornado, mutable, publico bool,
	fila int, columna int) DeclaracionFuncion {
	nd := DeclaracionFuncion{
		Id:      id,
		Mutable: mutable,
		Publico: publico,
		Valor:   valor,
		Tipo:    Ast.DECLARACION,
		Fila:    fila,
		Columna: columna,
		Retorno: tipo,
	}
	return nd
}

func (d DeclaracionFuncion) Run(scope *Ast.Scope) interface{} {
	/**************VARIABLES 3D*****************/
	//var direccion int = scope.Size
	var codigo3d string
	var obj3d Ast.O3D
	//var sizeFuncion int
	/*******************************************/

	//Verificar que el id no exista

	existe := scope.Exist_actual(d.Id)

	if existe {
		//Ya existe y generar error semántico
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
	//Verificar el tipo del retorno
	if d.Retorno.Tipo != Ast.VOID {
		if !expresiones.EsTipoFinal(d.Retorno.Tipo) || d.Retorno.Tipo == Ast.ACCESO_MODULO {
			nTipo := GetTipoEstructura(d.Retorno, scope, d)
			errors := ErrorEnTipo(nTipo)
			if errors.Tipo == Ast.ERROR {
				msg := "Semantic error, type error." +
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
			//De lo contrario todo bien y solo actualizar el tipo de la función
			funcion := d.Valor.(Funcion)
			funcion.Retorno = nTipo
			d.Valor = funcion
		}
	}

	//Conseguir el valor de la función en formato de tipo retornado
	valor := Ast.TipoRetornado{
		Valor: d.Valor,
		Tipo:  Ast.FUNCION,
	}

	/*******************APARTAR ESPACIO EN EL STACK PARA SUS PARAMETROS**************************/
	//funcion := valor.Valor.(Funcion)
	/*
		for i := 0; i < funcion.Parametros.Len(); i++ {
			Ast.GetP()
			scope.Size++
			sizeFuncion++
		}
		//Uno último para los returns
		Ast.GetP()
		scope.Size++
		sizeFuncion++
	*/
	/********************************************************************************************/

	//Todo bien crear y agregar el símbolo

	nSimbolo := Ast.Simbolo{
		Identificador: d.Id,
		Valor:         valor,
		Fila:          d.Fila,
		Columna:       d.Columna,
		Tipo:          valor.Tipo,
		Mutable:       d.Mutable,
		Publico:       d.Publico,
		Entorno:       scope,
		Size:          1,
	}
	nSimbolo.Direccion = 0
	nSimbolo.TipoDireccion = Ast.STACK

	scope.Add(nSimbolo)
	scope.Addfms(nSimbolo)
	obj3d.Codigo = codigo3d

	obj3d.Valor = Ast.TipoRetornado{
		Valor: true,
		Tipo:  Ast.EJECUTADO,
	}

	return Ast.TipoRetornado{
		Valor: obj3d,
		Tipo:  Ast.EJECUTADO,
	}

}

func (op DeclaracionFuncion) GetFila() int {
	return op.Fila
}
func (op DeclaracionFuncion) GetColumna() int {
	return op.Columna
}

func (d DeclaracionFuncion) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, Ast.DECLARACION
}
