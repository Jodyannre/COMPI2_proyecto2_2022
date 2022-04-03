package Ast

import "github.com/colegno/arraylist"

type Expresion interface {
	GetValue(entorno *Scope) TipoRetornado
}

type Instruccion interface {
	Run(entorno *Scope) interface{}
}

type Abstracto interface {
	GetTipo() (TipoDato, TipoDato)
	GetFila() int
	GetColumna() int
}

type Structs interface {
	GetPlantilla(scope *Scope) string
	SetMutabilidad(mutable bool) interface{}
}

type Modulos interface {
	GetFila() int
	GetTablas() int
	GetNombre() string
	GetEntorno() *Scope
}

type Funciones interface {
	GetTipoRetornado(scope *Scope) string
}

type Identificadores interface {
	GetNombre() string
}

type AbstractoM interface {
	GetMutable() bool
}

type AccesosM interface {
	GetTipoFromAccesoModulo(tipo TipoRetornado, scope *Scope) TipoRetornado
	GetTipoEstructura(tipo TipoRetornado, scope *Scope) TipoRetornado
}

type AccesosStruct interface {
	GetStruct(scope *Scope) TipoRetornado
	GetNombreAtributo() string
}

type AccesoVectorAbstracto interface {
	GetIdentificador() string
}

type Clones interface {
	Clonar(scope *Scope) interface{}
}

type CrearStruct interface {
	CrearStructInstancia(plantilla TipoRetornado, atributos *arraylist.List, mutable bool, fila, columna int) interface{}
}

type DeclaracionLugar interface {
	SetHeap() interface{}
}

type Error interface {
	GetFila() int
	GetColumna() int
	GetAmbito() string
	GetDescripcion() string
	GetFecha() string
}
