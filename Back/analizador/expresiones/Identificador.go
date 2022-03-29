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

func (p Identificador) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	//Buscar el símbolo en la tabla de símbolos y retornar el valor
	//Verificar que el id no exista
	if scope.Exist(p.Valor) {
		//Existe el identificar y retornar el valor
		simbolo := scope.GetSimbolo(p.Valor)
		return simbolo.Valor.(Ast.TipoRetornado)
	} else {
		//No existe el identificador, retornar error semantico
		msg := "Semantic error, \"" + p.Valor + "\" variable doesn't not exist." +
			" -- Line: " + strconv.Itoa(p.Fila) +
			" Column: " + strconv.Itoa(p.Columna)
		nError := errores.NewError(p.Fila, p.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
		//return Ast.TipoRetornado{Valor: nil, Tipo: Ast.NULL}
	}
}

func NewIdentificador(val string, tipo Ast.TipoDato, fila, columna int) Identificador {
	return Identificador{Tipo: tipo, Valor: val, Fila: fila, Columna: columna}
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
