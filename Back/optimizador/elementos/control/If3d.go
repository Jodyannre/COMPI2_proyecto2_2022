package control

import (
	"Back/optimizador/elementos/expresiones3d"
	"Back/optimizador/elementos/instrucciones3d"
	"Back/optimizador/elementos/interfaces3d"
)

type IF3D struct {
	Valor     string
	Operacion expresiones3d.Operacion3D
	GotoT     instrucciones3d.Goto3D
	GotoF     instrucciones3d.Goto3D
	LabelT    instrucciones3d.Salto3D
	Tipo      interfaces3d.TipoCD3
}

func NewIF3D(operacion expresiones3d.Operacion3D, gotoT, gotoF instrucciones3d.Goto3D,
	labelT instrucciones3d.Salto3D) IF3D {
	nO := IF3D{
		Operacion: operacion,
		GotoT:     gotoT,
		GotoF:     gotoF,
		LabelT:    labelT,
		Tipo:      interfaces3d.IF3D,
	}
	/*contruir el valor*/
	nValor := "if (" + operacion.GetValor() + ") " + gotoT.GetValor()
	nValor += gotoF.GetValor()
	nValor += labelT.GetValor()
	nO.Valor = nValor
	return nO
}

func (p IF3D) GetValor() string {
	return p.Valor
}

func (p IF3D) GetTipoParticular() interfaces3d.TipoCD3 {
	return p.Tipo
}

func (p IF3D) GetTiposGeneral() interfaces3d.TipoCD3 {
	return interfaces3d.IF3D
}
