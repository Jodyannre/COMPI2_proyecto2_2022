package headers3d

import "Back/optimizador/elementos/interfaces3d"

type Declaracion3D struct {
	Valor string
	Tipo  interfaces3d.TipoCD3
}

func NewDeclaracion3D(valor string) Declaracion3D {
	nO := Declaracion3D{
		Valor: valor,
		Tipo:  interfaces3d.DECLARACION3D,
	}
	return nO
}

func (p Declaracion3D) GetValor() string {
	return p.Valor + "; \n"
}

func (p Declaracion3D) GetTipoParticular() interfaces3d.TipoCD3 {
	return p.Tipo
}

func (p Declaracion3D) GetTiposGeneral() interfaces3d.TipoCD3 {
	return interfaces3d.DECLARACION3D
}
