package instrucciones3d

import "Back/optimizador/elementos/interfaces3d"

type Return3D struct {
	Valor string
	Tipo  interfaces3d.TipoCD3
}

func NewReturn3D(valor string) Return3D {
	nO := Return3D{
		Valor: valor,
		Tipo:  interfaces3d.RETURN3D,
	}
	return nO
}

func (p Return3D) GetValor() string {
	return p.Valor + "; \n"
}

func (p Return3D) GetTipoParticular() interfaces3d.TipoCD3 {
	return p.Tipo
}

func (p Return3D) GetTiposGeneral() interfaces3d.TipoCD3 {
	return interfaces3d.RETURN3D
}
