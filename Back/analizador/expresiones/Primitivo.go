package expresiones

import (
	"Back/analizador/Ast"
	"strings"
)

type Primitivo struct {
	Tipo    Ast.TipoDato
	Valor   interface{}
	Fila    int
	Columna int
}

func (p Primitivo) GetValue(entorno *Ast.Scope) Ast.TipoRetornado {
	valor := Ast.TipoRetornado{
		Tipo:  p.Tipo,
		Valor: p.Valor,
	}
	obj := Ast.O3D{
		Lt:         "",
		Lf:         "",
		Valor:      valor,
		Codigo:     "",
		Referencia: Primitivo_To_String(p.Valor, p.Tipo),
	}

	//Verificar que sea un string o un str
	if EsCadena(p.Tipo) {
		temp := Ast.GetTemp()
		cadenaAscii := obj.Referencia
		arrayAscii := strings.Split(cadenaAscii, ",")
		codigo := ""
		//Inicializar la cadena con el valor inicial del H guardado en el temporal
		codigo += "/***************AGREGANDO UN STRING/STR AL HEAP*/\n"
		codigo += temp + " = " + "H;\n"

		for _, valor := range arrayAscii {
			codigo += "heap[(int)H] = " + valor + "; //Letra\n"
			codigo += "H = H + 1;\n"
			Ast.GetH()
		}
		//Agregar caracter para saber que la cadena ha terminado
		codigo += "heap[(int)H] = 0;\n"
		codigo += "H = H + 1;\n"
		Ast.GetH()
		codigo += "/***********************************************/\n"

		obj.Codigo = codigo
		obj.Referencia = temp
	}

	return Ast.TipoRetornado{
		Tipo:  Ast.PRIMITIVO,
		Valor: obj,
	}
}

func (p Primitivo) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, p.Tipo
}

func NewPrimitivo(val interface{}, tipo Ast.TipoDato, fila, columna int) Primitivo {
	nuevo := Primitivo{Tipo: tipo, Valor: val, Fila: fila, Columna: columna}
	return nuevo
}

func (p Primitivo) GetFila() int {
	return p.Fila
}
func (p Primitivo) GetColumna() int {
	return p.Columna
}
