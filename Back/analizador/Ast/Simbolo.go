package Ast

import (
	"reflect"
)

type Simbolo struct {
	Identificador      string
	Valor              interface{}
	Fila               int
	Columna            int
	Tipo               TipoDato
	TipoEspecial       TipoRetornado
	Mutable            bool
	Publico            bool
	Entorno            *Scope
	Referencia         bool
	Referencia_puntero *Simbolo
	Direccion          int
	TipoDireccion      TipoDato
	Size               int
	CodigoGenerado     bool
	ReferenciaRetorno  string
}

type SimboloReporte struct {
	Identificador string
	TipoSimbolo   string
	TipoDato      string
	Scope         string
	Fila          int
	Columna       int
}

func NewSimbolo(identificador string, valor interface{}, fila int, columna int,
	tipo TipoDato, mutable bool, publico bool) Simbolo {
	simbolo := Simbolo{
		Identificador:      identificador,
		Valor:              valor,
		Fila:               fila,
		Columna:            columna,
		Tipo:               tipo,
		Mutable:            mutable,
		Publico:            publico,
		Referencia:         false,
		Referencia_puntero: nil,
		Entorno:            nil,
		TipoEspecial:       TipoRetornado{Valor: true, Tipo: INDEFINIDO},
		Size:               0,
	}
	return simbolo
}

func (s Simbolo) NewSimboloReporte(scope *Scope) SimboloReporte {
	var tipo string
	var tipoDato string
	var nombreScope string
	if s.Tipo == FUNCION {
		tipo = ValorTipoDato[FUNCION]
		tipoDato = s.Valor.(TipoRetornado).Valor.(Funciones).GetTipoRetornado(scope)
	} else if s.Tipo == MODULO {
		tipo = ValorTipoDato[MODULO]
		tipoDato = ""
	} else {
		tipo = ValorTipoDato[VARIABLE]
		tipoDato = ValorTipoDato[s.Valor.(TipoRetornado).Tipo]
	}

	if scope.Global {
		nombreScope = "Global"
	} else {
		nombreScope = "Local"
	}
	if reflect.TypeOf(s.Valor) == reflect.TypeOf(TipoRetornado{}) {
		return SimboloReporte{
			Identificador: s.Identificador,
			TipoSimbolo:   tipo,
			TipoDato:      tipoDato,
			Scope:         nombreScope,
			Fila:          s.Fila,
			Columna:       s.Columna,
		}
	} else {
		_, tipoElemento := s.Valor.(Abstracto).GetTipo()

		return SimboloReporte{
			Identificador: s.Identificador,
			TipoSimbolo:   tipo,
			TipoDato:      ValorTipoDato[tipoElemento],
			Scope:         nombreScope,
			Fila:          s.Fila,
			Columna:       s.Columna,
		}
	}

}
