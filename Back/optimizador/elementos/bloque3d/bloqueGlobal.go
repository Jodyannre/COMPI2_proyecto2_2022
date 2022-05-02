package bloque3d

import (
	"Back/optimizador/elementos/interfaces3d"

	"github.com/colegno/arraylist"
)

type Bloque3dGlobal struct {
	Tipo          interfaces3d.TipoCD3
	Declaraciones *arraylist.List
	Funciones     *arraylist.List
	Include       interfaces3d.Expresion3D
}

func NewBloque3dGlobal(declaraciones, funciones *arraylist.List, include interfaces3d.Expresion3D) Bloque3dGlobal {
	nP := Bloque3dGlobal{
		Declaraciones: declaraciones,
		Funciones:     funciones,
		Include:       include,
		Tipo:          interfaces3d.BLOQUE3DGLOBAL,
	}
	return nP
}

func (p Bloque3dGlobal) GetValor() string {
	return "nada"
}

func (p Bloque3dGlobal) GetTipoParticular() interfaces3d.TipoCD3 {
	return p.Tipo
}

func (p Bloque3dGlobal) GetTiposGeneral() interfaces3d.TipoCD3 {
	return interfaces3d.BLOQUE3DGLOBAL
}
