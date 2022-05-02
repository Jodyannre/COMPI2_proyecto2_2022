package instrucciones3d

import (
	"Back/optimizador/elementos/interfaces3d"
)

type Llamada3D struct {
	Valor string
	Tipo  interfaces3d.TipoCD3
}

func NewLlamada3D(valor string) Llamada3D {
	nO := Llamada3D{
		Valor: valor,
		Tipo:  interfaces3d.LLAMADA3D,
	}
	return nO
}

func (p Llamada3D) GetValor() string {
	return p.Valor + ";\n"
}

func (p Llamada3D) GetTipoParticular() interfaces3d.TipoCD3 {
	return p.Tipo
}

func (p Llamada3D) GetTiposGeneral() interfaces3d.TipoCD3 {
	return interfaces3d.LLAMADA3D
}
