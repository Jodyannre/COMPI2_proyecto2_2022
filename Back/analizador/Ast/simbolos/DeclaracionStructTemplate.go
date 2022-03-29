package simbolos

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"strconv"

	"github.com/colegno/arraylist"
)

type DeclaracionStructTemplate struct {
	Id        string
	Tipo      Ast.TipoDato
	Publico   bool
	Atributos *arraylist.List
	Fila      int
	Columna   int
}

func NewDeclaracionStructTemplate(id string, atributos *arraylist.List, publico bool,
	fila int, columna int) DeclaracionStructTemplate {
	nd := DeclaracionStructTemplate{
		Id:        id,
		Tipo:      Ast.DECLARACION,
		Publico:   publico,
		Atributos: atributos,
		Fila:      fila,
		Columna:   columna,
	}
	return nd
}

func (d DeclaracionStructTemplate) Run(scope *Ast.Scope) interface{} {
	//Verificar si existe, devuelve un símbolo
	existe := scope.Exist_fms_declaracion(d.Id)

	if existe.Tipo != Ast.ERROR_NO_EXISTE {
		//Error, ya existe
		msg := "Semantic error, the element \"" + d.Id + "\" already exist." +
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

	//No existe, entonces crearlo

	nuevaPlantilla := NewStructTemplate(d.Id, d.Atributos, d.Publico, d.Fila, d.Columna)
	plantillaCreada := nuevaPlantilla.GetValue(scope)
	if plantillaCreada.Tipo == Ast.ERROR {
		return plantillaCreada
	}

	nSimbolo := Ast.Simbolo{
		Identificador: d.Id,
		Valor:         plantillaCreada,
		Fila:          d.Fila,
		Columna:       d.Columna,
		Tipo:          Ast.STRUCT_TEMPLATE,
		Mutable:       false,
		Publico:       d.Publico,
		Entorno:       scope,
		Referencia:    false,
	}
	scope.Add(nSimbolo)
	scope.Addfms(nSimbolo)

	return Ast.TipoRetornado{
		Tipo:  Ast.EJECUTADO,
		Valor: true,
	}
}

func (d DeclaracionStructTemplate) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, Ast.DECLARACION
}

func (op DeclaracionStructTemplate) GetFila() int {
	return op.Fila
}
func (op DeclaracionStructTemplate) GetColumna() int {
	return op.Columna
}
