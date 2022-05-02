package funciones3d

import (
	"Back/optimizador/elementos/bloque3d"
	"Back/optimizador/elementos/interfaces3d"
	"Back/optimizador/elementos/reglas3d"

	"github.com/colegno/arraylist"
)

type Funcion3D struct {
	Nombre  string
	Tipo    interfaces3d.TipoCD3
	Bloques *arraylist.List
	Valor   string
}

func NewFuncion3D(nombre string, bloques *arraylist.List, tipo interfaces3d.TipoCD3) Funcion3D {
	nO := Funcion3D{
		Nombre:  nombre,
		Tipo:    tipo,
		Bloques: bloques,
	}
	if tipo == interfaces3d.MAIN3D {
		nO.Valor = "int " + nombre + "() { \n"
	} else {
		nO.Valor = "void " + nombre + "() { \n"
	}
	return nO
}

func (p Funcion3D) GetValor() string {
	nValor := p.Valor

	/*RECORRER LA LISTA DE BLOQUES*/
	var bloqueActual bloque3d.Bloque3d
	//var elemento interfaces3d.Expresion3D
	for i := 0; i < p.Bloques.Len(); i++ {
		bloqueActual = p.Bloques.GetValue(i).(bloque3d.Bloque3d)
		nValor += reglas3d.ComprobarReglas(bloqueActual)
		/*
			for i := 0; i < bloqueActual.Instrucciones.Len(); i++ {
				elemento = bloqueActual.Instrucciones.GetValue(i).(interfaces3d.Expresion3D)
				nValor += elemento.GetValor()
			}
		*/
	}
	nValor += "} \n"
	return nValor
}

func (p Funcion3D) GetTipoParticular() interfaces3d.TipoCD3 {
	return p.Tipo
}

func (p Funcion3D) GetTiposGeneral() interfaces3d.TipoCD3 {
	return interfaces3d.FUNCION3D
}
