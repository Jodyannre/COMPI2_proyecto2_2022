package instrucciones3d

import "Back/optimizador/elementos/interfaces3d"

type Goto3D struct {
	Valor string
	Tipo  interfaces3d.TipoCD3
}

func NewGoto3D(valor string) Goto3D {
	nO := Goto3D{
		Valor: valor,
		Tipo:  interfaces3d.GOTO3D,
	}
	return nO
}

func (p Goto3D) GetValor() string {
	return p.Valor + "; \n"
}

func (p Goto3D) GetTipoParticular() interfaces3d.TipoCD3 {
	return p.Tipo
}

func (p Goto3D) GetTiposGeneral() interfaces3d.TipoCD3 {
	return interfaces3d.GOTO3D
}
