package expresiones3d

import (
	"Back/optimizador/elementos/interfaces3d"
)

type Primitivo3D struct {
	Valor string
	Tipo  interfaces3d.TipoCD3
}

func NewPrimitivo3D(valor string, tipo interfaces3d.TipoCD3) Primitivo3D {
	nP := Primitivo3D{
		Valor: valor,
		Tipo:  tipo,
	}
	return nP
}

func (p Primitivo3D) GetValor() string {
	return p.Valor
}

func (p Primitivo3D) GetTipoParticular() interfaces3d.TipoCD3 {
	return p.Tipo
}

func (p Primitivo3D) GetTiposGeneral() interfaces3d.TipoCD3 {
	return interfaces3d.PRIMITIVO3D
}
