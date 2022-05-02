package expresiones3d

import (
	"Back/optimizador/elementos/interfaces3d"
)

type Operacion3D struct {
	OpDer    Primitivo3D
	OpIzq    Primitivo3D
	Operador string
	Valor    string
	Unario   bool
	Tipo     interfaces3d.TipoCD3
}

func NewOperacion3D(opIzq, opDer Primitivo3D, operador string, unario bool) Operacion3D {
	nO := Operacion3D{
		Tipo:     interfaces3d.OPERACION3D,
		OpIzq:    opIzq,
		OpDer:    opDer,
		Operador: operador,
		Unario:   unario,
	}

	/*Calcular el valor*/
	if unario {
		nO.Valor = opIzq.GetValor()
	} else {
		nO.Valor = opIzq.GetValor() + " " + operador + " " + opDer.GetValor()
	}
	return nO
}

func (p Operacion3D) GetValor() string {
	return p.Valor
}

func (p Operacion3D) GetTipoParticular() interfaces3d.TipoCD3 {
	return p.Tipo
}

func (p Operacion3D) GetTiposGeneral() interfaces3d.TipoCD3 {
	return interfaces3d.OPERACION3D
}
