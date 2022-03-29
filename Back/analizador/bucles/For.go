package bucles

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"

	"github.com/colegno/arraylist"
)

type For struct {
	Tipo          Ast.TipoDato
	Variable      interface{}
	Range         interface{}
	Instrucciones *arraylist.List
	Fila          int
	Columna       int
}

func NewFor(variable interface{}, condicion interface{}, instrucciones *arraylist.List,
	fila, columna int) For {
	nF := For{
		Tipo:          Ast.FOR,
		Instrucciones: instrucciones,
		Range:         condicion,
		Fila:          fila,
		Columna:       columna,
		Variable:      variable,
	}
	return nF
}

func (f For) Run(scope *Ast.Scope) interface{} {
	var variable expresiones.Identificador
	var nombreVariable string
	var tipoGeneral Ast.TipoDato
	var primerValor Ast.TipoRetornado
	var rango Ast.TipoRetornado
	var nSimbolo Ast.Simbolo
	var vector interface{}
	//var simboloTemp Ast.Simbolo
	//var variableTemp Ast.TipoRetornado
	var valorActual Ast.TipoRetornado
	var instruccion interface{}
	var tipoParticular Ast.TipoDato
	var resultadoInstruccion Ast.TipoRetornado
	//var primeraIteracion bool
	newScope := Ast.NewScope("For", scope)

	//Verificar que la expresión sea un identificador o error
	_, tipoParticular = f.Variable.(Ast.Abstracto).GetTipo()

	if tipoParticular != Ast.IDENTIFICADOR {
		//Error se espera un identificador
		fila := f.Variable.(Ast.Abstracto).GetFila()
		columna := f.Variable.(Ast.Abstracto).GetColumna()
		msg := "Semantic error, IDENTIFICADOR expected, found " + Ast.ValorTipoDato[tipoParticular] +
			". -- Line:" + strconv.Itoa(fila) + " Column: " +
			strconv.Itoa(columna)
		nError := errores.NewError(fila, columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		scope.UpdateScopeGlobal()
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}

	//Recuperar la variable y sus datos para crearla
	variable = f.Variable.(expresiones.Identificador)
	nombreVariable = variable.Valor
	//error, identifier expected

	//Ejecutar el range para obtener el vector que se va a iterar
	rango = f.Range.(Ast.Expresion).GetValue(&newScope)
	//Verificar error
	if rango.Tipo == Ast.ERROR {
		return rango
	}
	if rango.Tipo == Ast.VECTOR {
		vector = rango.Valor.(expresiones.Vector)
	} else {
		vector = rango.Valor.(expresiones.Array)
	}

	//Get el primer elemento
	if rango.Tipo == Ast.VECTOR {
		primerValor = rango.Valor.(expresiones.Vector).Valor.GetValue(0).(Ast.TipoRetornado)
	} else {
		primerValor = rango.Valor.(expresiones.Array).Elementos.GetValue(0).(Ast.TipoRetornado)
	}

	//Crear el símbolo de la variable que se va a utilizar en el for
	nSimbolo = Ast.Simbolo{
		Identificador: nombreVariable,
		Valor:         primerValor,
		Fila:          variable.Fila,
		Columna:       variable.Columna,
		Tipo:          primerValor.Tipo,
		Mutable:       true,
		Publico:       false,
		Entorno:       &newScope,
	}
	//Agregar el nuevo simbolo al scope del for
	newScope.Add(nSimbolo)
	//primeraIteracion = true
	if rango.Tipo == Ast.VECTOR {

		for i := 0; i < vector.(expresiones.Vector).Valor.Len(); i++ {
			//Hasta que sea verdadero y termine de iterar toda la lista
			//Actualizar el valor de la variable al siguiente elemento
			//Verificar la variable por si fue modifica en la iteración anterior

			//Recupero la variable tal como esta luego de la iteracion
			valorActual = vector.(expresiones.Vector).Valor.GetValue(i).(Ast.TipoRetornado)
			//simboloTemp = newScope.GetSimbolo(nombreVariable)
			//variableTemp = simboloTemp.Valor.(Ast.TipoRetornado)
			nSimbolo.Valor = vector.(expresiones.Vector).Valor.GetValue(i)
			//Verifico el valor antes de actualizar
			nSimbolo.Valor = valorActual
			newScope.UpdateSimbolo(nombreVariable, nSimbolo)

			//Ejectuar todas las instrucciones dentro del for en la n iteración
			for j := 0; j < f.Instrucciones.Len(); j++ {
				instruccion = f.Instrucciones.GetValue(j)

				//Verificar los tipos para saber que comportamiento tiene que tener
				tipoGeneral, _ = instruccion.(Ast.Abstracto).GetTipo()

				if tipoGeneral == Ast.INSTRUCCION {
					//Ejecutar run
					resultadoInstruccion = instruccion.(Ast.Instruccion).Run(&newScope).(Ast.TipoRetornado)

				} else if tipoGeneral == Ast.EXPRESION {
					//Ejecutar getvalue
					resultadoInstruccion = instruccion.(Ast.Expresion).GetValue(&newScope)

				}

				//Verificar las instrucciones de transferencia
				if Ast.EsTransferencia(resultadoInstruccion.Tipo) {

					//Primero verificar que no sea un return normal, el cual si es permitido
					if resultadoInstruccion.Tipo == Ast.CONTINUE {
						//Rompemos y vamos a la siguiente iteración del for
						break
					}
					switch resultadoInstruccion.Tipo {
					case Ast.BREAK_EXPRESION, Ast.RETURN_EXPRESION, Ast.RETURN:
						return resultadoInstruccion
					case Ast.BREAK:
						return Ast.TipoRetornado{
							Tipo:  Ast.EJECUTADO,
							Valor: true,
						}
					}
				}

			}
			//primeraIteracion = false
		}

	} else {
		for i := 0; i < vector.(expresiones.Array).Elementos.Len(); i++ {
			//Hasta que sea verdadero y termine de iterar toda la lista
			//Actualizar el valor de la variable al siguiente elemento
			//Verificar la variable por si fue modifica en la iteración anterior

			//Recupero la variable tal como esta luego de la iteracion
			valorActual = vector.(expresiones.Array).Elementos.GetValue(i).(Ast.TipoRetornado)
			//simboloTemp = newScope.GetSimbolo(nombreVariable)
			//variableTemp = simboloTemp.Valor.(Ast.TipoRetornado)
			nSimbolo.Valor = vector.(expresiones.Array).Elementos.GetValue(i)
			//Verifico el valor antes de actualizar
			nSimbolo.Valor = valorActual
			newScope.UpdateSimbolo(nombreVariable, nSimbolo)

			//Ejectuar todas las instrucciones dentro del for en la n iteración
			for j := 0; j < f.Instrucciones.Len(); j++ {
				instruccion = f.Instrucciones.GetValue(j)

				//Verificar los tipos para saber que comportamiento tiene que tener
				tipoGeneral, _ = instruccion.(Ast.Abstracto).GetTipo()

				if tipoGeneral == Ast.INSTRUCCION {
					//Ejecutar run
					resultadoInstruccion = instruccion.(Ast.Instruccion).Run(&newScope).(Ast.TipoRetornado)

				} else if tipoGeneral == Ast.EXPRESION {
					//Ejecutar getvalue
					resultadoInstruccion = instruccion.(Ast.Expresion).GetValue(&newScope)

				}
				//Verificar las instrucciones de transferencia
				if Ast.EsTransferencia(resultadoInstruccion.Tipo) {

					//Primero verificar que no sea un return normal, el cual si es permitido
					if resultadoInstruccion.Tipo == Ast.CONTINUE {
						//Rompemos y vamos a la siguiente iteración del for
						break
					}
					switch resultadoInstruccion.Tipo {
					case Ast.BREAK_EXPRESION, Ast.RETURN_EXPRESION, Ast.RETURN:
						return resultadoInstruccion
					case Ast.BREAK:
						return Ast.TipoRetornado{
							Tipo:  Ast.EJECUTADO,
							Valor: true,
						}
					}
				}

			}
			//primeraIteracion = false
		}
	}
	newScope.UpdateScopeGlobal()
	return Ast.TipoRetornado{
		Tipo:  Ast.EJECUTADO,
		Valor: true,
	}
}

func (op For) GetFila() int {
	return op.Fila
}
func (op For) GetColumna() int {
	return op.Columna
}
func (f For) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, f.Tipo
}
