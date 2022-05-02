package instrucciones3d

import "Back/optimizador/elementos/interfaces3d"

type Salto3D struct {
	Valor string
	Tipo  interfaces3d.TipoCD3
}

func NewSalto3D(valor string) Salto3D {
	nO := Salto3D{
		Valor: valor,
		Tipo:  interfaces3d.LABEL3D,
	}
	return nO
}

func (p Salto3D) GetValor() string {
	return p.Valor + ": \n"
}

func (p Salto3D) GetTipoParticular() interfaces3d.TipoCD3 {
	return p.Tipo
}

func (p Salto3D) GetTiposGeneral() interfaces3d.TipoCD3 {
	return interfaces3d.LABEL3D
}
