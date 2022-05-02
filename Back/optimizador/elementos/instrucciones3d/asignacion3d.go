package instrucciones3d

import (
	"Back/optimizador/elementos/expresiones3d"
	"Back/optimizador/elementos/interfaces3d"
)

type Asignacion3D struct {
	Elemento  expresiones3d.Primitivo3D
	Operacion expresiones3d.Operacion3D
	Valor     string
	Tipo      interfaces3d.TipoCD3
}

func NewAsignacion3D(elemento expresiones3d.Primitivo3D, operacion expresiones3d.Operacion3D) Asignacion3D {
	nO := Asignacion3D{
		Elemento:  elemento,
		Operacion: operacion,
		Tipo:      interfaces3d.ASIGNACION3D,
	}
	/*Construir el valor*/
	valorOp := operacion.GetValor()
	valorEl := elemento.GetValor()
	nO.Valor = valorEl + " = " + valorOp
	return nO
}

func (p Asignacion3D) GetValor() string {
	return p.Valor + "; \n"
}

func (p Asignacion3D) GetTipoParticular() interfaces3d.TipoCD3 {
	return p.Tipo
}

func (p Asignacion3D) GetTiposGeneral() interfaces3d.TipoCD3 {
	return interfaces3d.ASIGNACION3D
}
