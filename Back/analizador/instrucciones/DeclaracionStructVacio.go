package instrucciones

/*
import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"
)

type DeclaracionVSVacio struct {
	Id      string
	Tipo    Ast.TipoRetornado
	Mutable bool
	Fila    int
	Columna int
}

func NewDeclaracionVSVacio(id string, tipo Ast.TipoRetornado, mutable bool, fila int, columna int) DeclaracionVSVacio {
	nd := DeclaracionVSVacio{
		Id:      id,
		Mutable: mutable,
		Fila:    fila,
		Columna: columna,
		Tipo:    tipo,
	}
	return nd
}

func (d DeclaracionVSVacio) Run(scope *Ast.Scope) interface{} {
	tipoVariable := d.Tipo.Tipo
	valorTemporal := expresiones.NewPrimitivo(nil, Ast.NULL,d.Fila,d.Columna)
	var existe bool

	//Verificar que no exista el elemento
	existe = scope.Exist_actual(d.Id)

	if !existe {
		fila := d.Fila
		columna := d.Columna
		msg := "Semantic error, the element \"" + d.Id + "\" already exist in this scope." +
			" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
		nError := errores.NewError(fila, columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}

	if tipoVariable == Ast.STRUCT{
		nSimbolo := Ast.Simbolo{
			Identificador: d.Id,
			Valor:         Ast.TipoRetornado{
				Valor: valorTemporal,
				Tipo: d.Tipo.Tipo,
			},
			Fila:          d.Fila,
			Columna:       d.Columna,
			Tipo:          Ast.STRUCT,
			Mutable:       d.Mutable,
			Publico:       d.Publico,
			Entorno:       scope,
		}

	}else{

	}

	return Ast.TipoRetornado{
		Tipo:  Ast.EJECUTADO,
		Valor: true,
	}
}

func (op DeclaracionVSVacio) GetFila() int {
	return op.Fila
}
func (op DeclaracionVSVacio) GetColumna() int {
	return op.Columna
}

func (d DeclaracionVSVacio) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, Ast.DECLARACION
}
*/
