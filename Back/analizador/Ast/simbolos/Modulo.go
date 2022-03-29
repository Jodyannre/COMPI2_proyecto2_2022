package simbolos

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"strconv"

	"github.com/colegno/arraylist"
)

type Modulo struct {
	Identificador interface{}
	Tipo          Ast.TipoDato
	Instrucciones *arraylist.List
	Fila          int
	Columna       int
	Publico       bool
	Entorno       *Ast.Scope
}

func NewModulo(id interface{}, instrucciones *arraylist.List, publico bool, fila, columna int) Modulo {
	nM := Modulo{
		Tipo:          Ast.MODULO,
		Instrucciones: instrucciones,
		Publico:       publico,
		Fila:          fila,
		Columna:       columna,
		Identificador: id,
	}
	return nM
}

func (m Modulo) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	//Variables de resultado de la ejecución de las instrucciones
	var resultadoInstruccion Ast.TipoRetornado

	//Crear el entorno del módulo
	newScope := Ast.NewScope("Modulo", scope)

	//Declarar todos los módulos y demás atributos del módulo

	//Recorrer las instrucciones y verificar que no sean expresiones
	for i := 0; i < m.Instrucciones.Len(); i++ {
		//Get la instrucción
		instruccion := m.Instrucciones.GetValue(i)
		//Obtener los tipos de la instrucción
		tipoGlobal, tipoParticular := instruccion.(Ast.Abstracto).GetTipo()
		//Verificar que sea una instrucción
		if tipoGlobal == Ast.INSTRUCCION {
			//Verificar que la instrucción sea una declaración o sino ignorarla
			if tipoParticular != Ast.DECLARACION {
				fila := instruccion.(Ast.Abstracto).GetFila()
				columna := instruccion.(Ast.Abstracto).GetColumna()
				msg := "Semantic error, a DECLARACION was expected, found " + Ast.ValorTipoDato[tipoParticular] +
					". -- Line: " + strconv.Itoa(fila) +
					" Column: " + strconv.Itoa(columna)
				nError := errores.NewError(fila, columna, msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				nError.Ambito = scope.GetTipoScope()
				newScope.Errores.Add(nError)
				newScope.Consola += msg + "\n"
				continue
			}

			//Ejecutar la instrucción
			resultadoInstruccion = instruccion.(Ast.Instruccion).Run(&newScope).(Ast.TipoRetornado)

			//Verificar los posibles retornos
			if resultadoInstruccion.Tipo == Ast.ERROR {
				continue
			}

		} else {
			fila := instruccion.(Ast.Abstracto).GetFila()
			columna := instruccion.(Ast.Abstracto).GetColumna()
			msg := "Semantic error, an INSTRUCCION was expected, found " + Ast.ValorTipoDato[tipoParticular] +
				". -- Line: " + strconv.Itoa(fila) +
				" Column: " + strconv.Itoa(columna)
			nError := errores.NewError(fila, columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			nError.Ambito = scope.GetTipoScope()
			newScope.Errores.Add(nError)
			newScope.Consola += msg + "\n"
			continue
		}

	}
	//Agregar el entorno al modulo
	m.Entorno = &newScope

	newScope.UpdateScopeGlobal()

	return Ast.TipoRetornado{
		Tipo:  Ast.MODULO,
		Valor: m,
	}

}

func (a Modulo) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, a.Tipo
}

func (a Modulo) GetFila() int {
	return a.Fila
}
func (a Modulo) GetColumna() int {
	return a.Columna
}

func (a Modulo) GetTablas() int {
	var cantidad int = 0
	tabla := a.Entorno.GetTablaModulos()
	for key, _ := range tabla {
		if key == "nada" {
			println("para que no haya error")
		}
		cantidad++
	}
	return cantidad
}

func (a Modulo) GetEntorno() *Ast.Scope {
	return a.Entorno
}

func (a Modulo) GetNombre() string {
	_, tipoParticular := a.Identificador.(Ast.Abstracto).GetTipo()
	var nombre string
	if tipoParticular == Ast.IDENTIFICADOR {
		nombre = a.Identificador.(Ast.Identificadores).GetNombre()
	} else {
		nombre = a.Identificador.(AccesoModulo).Elementos.GetValue(0).(Ast.Identificadores).GetNombre()
	}
	return nombre
}
