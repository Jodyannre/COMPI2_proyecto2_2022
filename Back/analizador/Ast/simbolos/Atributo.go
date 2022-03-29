package simbolos

import (
	"Back/analizador/Ast"
	"Back/analizador/expresiones"
	"reflect"
)

type Atributo struct {
	Nombre       string
	Tipo         Ast.TipoDato
	Valor        interface{}
	TipoAtributo Ast.TipoRetornado
	Fila         int
	Columna      int
	Publico      bool
	Mutable      bool
}

func NewAtributoTemplate(nombre string, tipo Ast.TipoRetornado, publico bool, fila, columna int) *Atributo {
	nuevo := &Atributo{
		Nombre:       nombre,
		Tipo:         Ast.ATRIBUTO,
		TipoAtributo: tipo,
		Fila:         fila,
		Columna:      columna,
		Publico:      publico,
	}
	return nuevo
}

func NewAtributo(nombre string, valor interface{}, mutable bool, fila, columna int) *Atributo {
	nuevo := &Atributo{
		Tipo:    Ast.ATRIBUTO,
		Nombre:  nombre,
		Fila:    fila,
		Valor:   valor,
		Columna: columna,
		Mutable: mutable,
	}
	return nuevo
}

func (a *Atributo) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	var valor Ast.TipoRetornado
	///////////////////////////////////////////////////////
	//Clono el puntero para no modificarlo luego
	clone := a.Clonar()
	//calcular el valor y los tipos del atributo
	if reflect.TypeOf(a.Valor) == reflect.TypeOf(Ast.TipoRetornado{}) {
		valor = a.Valor.(Ast.TipoRetornado)
	} else {
		valor = a.Valor.(Ast.Expresion).GetValue(scope)
	}

	if valor.Tipo == Ast.ERROR {
		return valor
	}

	if esTipoFinal(valor.Tipo) && valor.Tipo != Ast.STRUCT {
		clone.TipoAtributo = Ast.TipoRetornado{Tipo: valor.Tipo, Valor: true}
	} else if valor.Tipo == Ast.STRUCT {
		nombreStruct := valor.Valor.(Ast.Structs).GetPlantilla(scope)
		clone.TipoAtributo = Ast.TipoRetornado{Tipo: valor.Tipo, Valor: nombreStruct}
	} else if valor.Tipo == Ast.DIMENSION_ARRAY {
		clone.TipoAtributo = Ast.TipoRetornado{Tipo: Ast.ARRAY, Valor: a.Valor.(expresiones.DimensionArray).TipoArray}
	} else if valor.Tipo == Ast.ARRAY {
		clone.TipoAtributo = Ast.TipoRetornado{Tipo: Ast.ARRAY, Valor: valor.Valor.(expresiones.Array).TipoDelArray}
	} else if valor.Tipo == Ast.VECTOR {
		clone.TipoAtributo = Ast.TipoRetornado{Tipo: Ast.VECTOR, Valor: valor.Valor.(expresiones.Vector).TipoVector}
	}
	clone.Valor = valor

	return Ast.TipoRetornado{
		Tipo:  Ast.ATRIBUTO,
		Valor: clone,
	}
}

func (a *Atributo) FormatearTipo(scope *Ast.Scope) Ast.TipoRetornado {
	if EsPosibleReferencia(a.TipoAtributo.Tipo) {
		nTipo := GetTipoEstructura(a.TipoAtributo, scope, a)
		return nTipo
	} else {
		return Ast.TipoRetornado{
			Tipo:  Ast.BOOLEAN,
			Valor: true,
		}
	}
}

/*
func (p Atributo) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, p.Tipo
}

func (p Atributo) GetFila() int {
	return p.Fila
}
func (p Atributo) GetColumna() int {
	return p.Columna
}
*/

func (p *Atributo) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, p.Tipo
}

func (p *Atributo) GetFila() int {
	return p.Fila
}
func (p *Atributo) GetColumna() int {
	return p.Columna
}

func (a *Atributo) Clonar() *Atributo {
	nuevoAtributo := &Atributo{
		Nombre:       a.Nombre,
		Valor:        a.Valor,
		Tipo:         a.Tipo,
		Fila:         a.Fila,
		Columna:      a.Columna,
		Publico:      a.Publico,
		Mutable:      a.Mutable,
		TipoAtributo: a.TipoAtributo,
	}
	return nuevoAtributo
}
