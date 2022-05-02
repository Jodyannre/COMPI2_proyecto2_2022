package bloque3d

import (
	"Back/optimizador/elementos/interfaces3d"

	"github.com/colegno/arraylist"
)

type Bloque3d struct {
	Tipo          interfaces3d.TipoCD3
	Instrucciones *arraylist.List
}

func NewBloque3d(instrucciones *arraylist.List) Bloque3d {
	nP := Bloque3d{
		Instrucciones: instrucciones,
		Tipo:          interfaces3d.BLOQUE3D,
	}
	return nP
}

func (p Bloque3d) GetInstrucciones() *arraylist.List {
	return p.Instrucciones
}

func (p Bloque3d) GetValor() string {
	return "nada"
}

func (p Bloque3d) GetTipoParticular() interfaces3d.TipoCD3 {
	return p.Tipo
}

func (p Bloque3d) GetTiposGeneral() interfaces3d.TipoCD3 {
	return interfaces3d.BLOQUE3D
}
