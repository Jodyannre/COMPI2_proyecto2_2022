package headers3d

import "Back/optimizador/elementos/interfaces3d"

type Include3D struct {
	Valor string
	Tipo  interfaces3d.TipoCD3
}

func NewInclude3D(valor string) Include3D {
	nO := Include3D{
		Valor: valor,
		Tipo:  interfaces3d.INCLUDE,
	}
	return nO
}

func (p Include3D) GetValor() string {
	return p.Valor + "\n"
}

func (p Include3D) GetTipoParticular() interfaces3d.TipoCD3 {
	return p.Tipo
}

func (p Include3D) GetTiposGeneral() interfaces3d.TipoCD3 {
	return interfaces3d.INCLUDE
}
