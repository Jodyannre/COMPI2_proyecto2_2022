package expresiones

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"strconv"
)

type Identificador struct {
	Tipo    Ast.TipoDato
	Valor   string
	Fila    int
	Columna int
}

func NewIdentificador(val string, tipo Ast.TipoDato, fila, columna int) Identificador {
	return Identificador{Tipo: tipo, Valor: val, Fila: fila, Columna: columna}
}

func (p Identificador) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	/*Variables para C3D*/
	var temp string = Ast.GetTemp()
	var tempValor string = Ast.GetTemp()
	var codigo3d string = ""
	var obj3D Ast.O3D

	//Buscar el símbolo en la tabla de símbolos y retornar el valor

	if !scope.Exist(p.Valor) {
		//No existe el identificador, retornar error semantico
		/////////////////////////////ERROR/////////////////////////////////////
		return errores.GenerarError(20, p, p, p.Valor, "", "", scope)
	}

	//Existe el identificar y retornar el valor
	simbolo := scope.GetSimbolo(p.Valor)

	/*Generar codigo 3d*/

	/*Verificar si la variable viene del heap o del stack*/

	if simbolo.TipoDireccion == Ast.STACK {
		codigo3d = "/****************GET VARIABLE CON IDENTIFICADOR*/\n"
		codigo3d += temp + " = P + " + strconv.Itoa(simbolo.Direccion) + ";\n"
		codigo3d += tempValor + " = stack[(int)" + temp + "];\n"
		codigo3d += "/***********************************************/\n"
	} else {
		codigo3d = "/****************GET VARIABLE CON IDENTIFICADOR*/\n"
		codigo3d += temp + " = " + strconv.Itoa(simbolo.Direccion) + ";\n"
		codigo3d += tempValor + " = heap[(int)" + temp + "];\n"
		codigo3d += "/***********************************************/\n"
	}

	/*Inicializar el obj3d*/
	obj3D.Referencia = tempValor
	obj3D.Valor = simbolo.Valor.(Ast.TipoRetornado)
	obj3D.Codigo = codigo3d

	return Ast.TipoRetornado{
		Tipo:  Ast.IDENTIFICADOR,
		Valor: obj3D,
	}

}

func (p Identificador) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, Ast.IDENTIFICADOR
}
func (op Identificador) GetFila() int {
	return op.Fila
}
func (op Identificador) GetColumna() int {
	return op.Columna
}

func (op Identificador) GetNombre() string {
	return op.Valor
}
