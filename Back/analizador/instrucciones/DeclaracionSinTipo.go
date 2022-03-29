package instrucciones

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"
)

type DeclaracionSinTipo struct {
	Id      string
	Mutable bool
	Publico bool
	Valor   interface{}
	Fila    int
	Columna int
}

func NewDeclaracionSinTipo(id string, valor interface{}, mutable, publico bool,
	fila int, columna int) DeclaracionSinTipo {
	nd := DeclaracionSinTipo{
		Id:      id,
		Mutable: mutable,
		Publico: publico,
		Valor:   valor,
		Fila:    fila,
		Columna: columna,
	}
	return nd
}

func (d DeclaracionSinTipo) Run(scope *Ast.Scope) interface{} {
	var nSimbolo Ast.Simbolo

	//Verificar que el id no exista

	existe := scope.Exist_actual(d.Id)

	//Primero verificar que no es un if expresion
	_, tipoIn := d.Valor.(Ast.Abstracto).GetTipo()

	var preValor interface{}
	if tipoIn == Ast.IF_EXPRESION || tipoIn == Ast.MATCH_EXPRESION || tipoIn == Ast.LOOP_EXPRESION {
		preValor = d.Valor.(Ast.Instruccion).Run(scope)
	} else if tipoIn == Ast.FUNCION {
		preValor = Ast.TipoRetornado{
			Valor: d.Valor,
			Tipo:  Ast.FUNCION,
		}
	} else {
		preValor = d.Valor.(Ast.Expresion).GetValue(scope)
	}
	nuevoValor := preValor.(Ast.TipoRetornado)
	valor := nuevoValor
	//Revisar si el retorno es un error
	if valor.Tipo == Ast.ERROR {
		return valor
	}

	if !existe {
		//No existe, entonces agregarla
		//Verificar la mutabilidad para cambiarla si es un vector o un array o un struct
		//Agregar el tipo al simbolo en el apartado de tipo especial para despues
		if valor.Tipo == Ast.VECTOR {
			elemento := valor.Valor.(expresiones.Vector)
			elemento.Mutable = d.Mutable
			valor.Valor = elemento
		} else if valor.Tipo == Ast.ARRAY {
			elemento := valor.Valor.(expresiones.Array)
			elemento.Mutable = d.Mutable
			valor.Valor = elemento
		} else if valor.Tipo == Ast.STRUCT {
			nValor := valor.Valor.(Ast.Structs).SetMutabilidad(d.Mutable)
			valor.Valor = nValor
		}
		//Crear símbolo y agregarlo a la tabla del entorno actual
		nSimbolo = Ast.Simbolo{
			Identificador: d.Id,
			Valor:         valor,
			Fila:          d.Fila,
			Columna:       d.Columna,
			Tipo:          valor.Tipo,
			Mutable:       d.Mutable,
			Publico:       d.Publico,
			Entorno:       scope,
		}

		//Puede que sea un tipo especial

		if EsTipoEspecial(valor.Tipo) {
			//Agregar el tipo al simbolo en el apartado de tipo especial para despues
			if valor.Tipo == Ast.VECTOR {
				elemento := valor.Valor.(expresiones.Vector)
				nSimbolo.TipoEspecial = elemento.TipoVector
			} else if valor.Tipo == Ast.ARRAY {
				elemento := valor.Valor.(expresiones.Array)
				nSimbolo.TipoEspecial = elemento.TipoDelArray
			} else if valor.Tipo == Ast.STRUCT {
				elemento := valor.Valor.(Ast.Structs).GetPlantilla(scope)
				nSimbolo.TipoEspecial = Ast.TipoRetornado{
					Valor: elemento,
					Tipo:  Ast.STRUCT,
				}
			}
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

func (op DeclaracionSinTipo) GetFila() int {
	return op.Fila
}
func (op DeclaracionSinTipo) GetColumna() int {
	return op.Columna
}

func (d DeclaracionSinTipo) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, Ast.DECLARACION
}
