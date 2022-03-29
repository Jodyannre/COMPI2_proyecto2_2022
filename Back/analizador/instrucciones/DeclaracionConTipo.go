package instrucciones

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"strconv"
)

type DeclaracionConTipo struct {
	Id      string
	Mutable bool
	Publico bool
	Tipo    Ast.TipoRetornado
	Fila    int
	Columna int
}

func NewDeclaracionConTipo(id string, tipo Ast.TipoRetornado, mutable, publico bool,
	fila int, columna int) DeclaracionConTipo {
	nd := DeclaracionConTipo{
		Id:      id,
		Mutable: mutable,
		Publico: publico,
		Tipo:    tipo,
		Fila:    fila,
		Columna: columna,
	}
	return nd
}

func (d DeclaracionConTipo) Run(scope *Ast.Scope) interface{} {
	var nSimbolo Ast.Simbolo

	//Verificar que el id no exista
	existe := scope.Exist_actual(d.Id)

	//No trae ningún valor

	//Verificar si el tipo es un acceso a un módulo
	if d.Tipo.Tipo == Ast.ACCESO_MODULO {
		//Traer el tipo y cambiar el tipo de la declaración
		nTipo := d.Tipo.Valor.(Ast.AccesosM).GetTipoFromAccesoModulo(d.Tipo, scope)
		if nTipo.Tipo == Ast.ERROR {
			return nTipo
		}
		if nTipo.Tipo == Ast.STRUCT_TEMPLATE {
			nTipo.Tipo = Ast.STRUCT
		}
		d.Tipo = nTipo
	}

	if !existe {
		//No existe, entonces agregarla
		//Crear símbolo y agregarlo a la tabla del entorno actual

		nSimbolo = Ast.Simbolo{
			Identificador: d.Id,
			Valor:         Ast.TipoRetornado{Valor: nil, Tipo: Ast.NULL},
			Fila:          d.Fila,
			Columna:       d.Columna,
			Tipo:          d.Tipo.Tipo,
			Mutable:       d.Mutable,
			Publico:       d.Publico,
			Entorno:       scope,
		}

		//Verificar si no es vector, array o struct para agregar el tipo especial
		if EsTipoEspecial(d.Tipo.Tipo) {
			nSimbolo.TipoEspecial = d.Tipo
		}
		scope.Add(nSimbolo)
	} else {
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

	return Ast.TipoRetornado{
		Tipo:  Ast.EJECUTADO,
		Valor: true,
	}
}

func (op DeclaracionConTipo) GetFila() int {
	return op.Fila
}
func (op DeclaracionConTipo) GetColumna() int {
	return op.Columna
}

func (d DeclaracionConTipo) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, Ast.DECLARACION
}

/*
func ValorPorDefecto(tipo Ast.TipoRetornado)Ast.TipoRetornado{
	var retorno Ast.TipoRetornado
	switch tipo.Tipo{
	case Ast.STRUCT:
		estructura :=
		retorno.Tipo = Ast.STRUCT

	case Ast.I64:
		retorno.Tipo = Ast.I64
		retorno.Valor = 0
	case Ast.F64:
		retorno.Tipo = Ast.F64
		retorno.Valor = float64(0)
	case Ast.CHAR:
		retorno.Tipo = Ast.CHAR
		retorno.Valor = ""
	case Ast.STRING:
		retorno.Tipo = Ast.STRING
		retorno.Valor = ""
	case Ast.STR:
		retorno.Tipo = Ast.STR
		retorno.Valor = ""
	case Ast.BOOLEAN:
		retorno.Tipo = Ast.BOOLEAN
		retorno.Valor = false
	case Ast.USIZE:
		retorno.Tipo = Ast.USIZE
		retorno.Valor = 0
	}
}
*/
