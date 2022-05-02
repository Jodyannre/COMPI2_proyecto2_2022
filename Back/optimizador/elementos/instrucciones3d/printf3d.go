package instrucciones3d

import "Back/optimizador/elementos/interfaces3d"

type Print3D struct {
	Valor string
	Tipo  interfaces3d.TipoCD3
}

func NewPrint3D(valor string) Print3D {
	nO := Print3D{
		Valor: valor,
		Tipo:  interfaces3d.PRINT3D,
	}
	return nO
}

func (p Print3D) GetValor() string {
	return p.Valor + "; \n"
}

func (p Print3D) GetTipoParticular() interfaces3d.TipoCD3 {
	return p.Tipo
}

func (p Print3D) GetTiposGeneral() interfaces3d.TipoCD3 {
	return interfaces3d.PRINT3D
}
