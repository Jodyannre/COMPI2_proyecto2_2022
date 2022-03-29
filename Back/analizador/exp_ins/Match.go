package exp_ins

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"strconv"

	"github.com/colegno/arraylist"
)

type Match struct {
	Expresion Ast.Expresion
	Cases     *arraylist.List
	Tipo      Ast.TipoDato
	Fila      int
	Columna   int
}

type Case struct {
	Expresion     *arraylist.List
	Instrucciones *arraylist.List
	Tipo          Ast.TipoDato
	Default       bool
	Fila          int
	Columna       int
}

func NewMatch(expresion Ast.Expresion, cases *arraylist.List, tipo Ast.TipoDato, fila, columna int) Match {
	nMatch := Match{
		Expresion: expresion,
		Cases:     cases,
		Tipo:      tipo,
		Fila:      fila,
		Columna:   columna,
	}
	return nMatch
}

func NewCase(expresion *arraylist.List, instrucciones *arraylist.List, tipo Ast.TipoDato,
	fila, columna int, Default bool) Case {
	nCase := Case{
		Expresion:     expresion,
		Instrucciones: instrucciones,
		Tipo:          tipo,
		Default:       Default,
		Fila:          fila,
		Columna:       columna,
	}
	return nCase
}

func (m Match) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, m.Tipo
}

func (c Case) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, c.Tipo
}

func (m Match) Run(scope *Ast.Scope) interface{} {
	//Primero conseguir el exp_match de la expresion
	buscar_default := false
	default_econtrado := false
	exp_match := m.Expresion.GetValue(scope)
	pos_expresion_correcta := -1
	var i, j int

	//Validar si es un boolean y si tiene todos los casos
	if exp_match.Tipo == Ast.BOOLEAN && m.Cases.Len() < 2 {
		//Error, no estan todos los casos incluidos
		msg := "Semantic error, a default case was expected." +
			" -- Line:" + strconv.Itoa(m.Fila) + " Column: " + strconv.Itoa(m.Columna)
		nError := errores.NewError(m.Fila, m.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Valor: nError,
			Tipo:  Ast.ERROR,
		}
	}

	//Verificar si no es boolean y se espera un default
	if exp_match.Tipo != Ast.BOOLEAN {
		buscar_default = true
	}

	for i = 0; i < m.Cases.Len(); i++ {
		//Recorrer la lista de cases y verificar que los exp_matches de las expresiones concuerdan
		caso := m.Cases.GetValue(i).(Case)
		listaExpresiones := caso.Expresion

		for j = 0; j < listaExpresiones.Len(); j++ {
			expresion := listaExpresiones.GetValue(j).(Ast.Expresion).GetValue(scope)

			if expresion.Tipo != exp_match.Tipo && !caso.Default {
				//Error de tipos
				msg := "Semantic error, the expression is not a " + Ast.ValorTipoDato[exp_match.Tipo] +
					" value. -- Line:" + strconv.Itoa(caso.Fila) + " Column: " + strconv.Itoa(caso.Columna)
				nError := errores.NewError(caso.Fila, caso.Columna, msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				nError.Ambito = scope.GetTipoScope()
				scope.Errores.Add(nError)
				scope.Consola += msg + "\n"
				return Ast.TipoRetornado{
					Valor: nError,
					Tipo:  Ast.ERROR,
				}
			}

			if exp_match.Valor == expresion.Valor {
				pos_expresion_correcta = i
			}
		}
		//Verificar si viene default
		if buscar_default {
			if caso.Default {
				default_econtrado = true
			}
		}

		if default_econtrado && i != m.Cases.Len()-1 {
			//Error, el default tiene que venir de último
			msg := "Semantic error, the default case was expected in the last position." +
				" -- Line:" + strconv.Itoa(caso.Fila) + " Column: " + strconv.Itoa(caso.Columna)
			nError := errores.NewError(caso.Fila, caso.Columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			nError.Ambito = scope.GetTipoScope()
			scope.Errores.Add(nError)
			scope.Consola += msg + "\n"
			return Ast.TipoRetornado{
				Valor: nError,
				Tipo:  Ast.ERROR,
			}
		}

	}
	//Verificar que se necesite default y que lo traiga
	if buscar_default {
		if !default_econtrado {
			msg := "Semantic error, default case not found." +
				" -- Line:" + strconv.Itoa(m.Fila) + " Column: " + strconv.Itoa(m.Columna)
			nError := errores.NewError(m.Fila, m.Columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			nError.Ambito = scope.GetTipoScope()
			scope.Errores.Add(nError)
			scope.Consola += msg + "\n"
			scope.UpdateScopeGlobal()
			return Ast.TipoRetornado{
				Valor: nError,
				Tipo:  Ast.ERROR,
			}
		}
	}

	//Todos los casos son correctos
	//Recorrer los cases
	var resultado_retornado Ast.TipoRetornado
	//Recuperar el caso en donde se encontró el valor igual
	if pos_expresion_correcta != -1 {
		caso := m.Cases.GetValue(pos_expresion_correcta).(Case)
		resultado_retornado = caso.Run(scope).(Ast.TipoRetornado)
	} else {
		//Ejecutar default
		caso := m.Cases.GetValue(m.Cases.Len() - 1).(Case)
		resultado_retornado = caso.Run(scope).(Ast.TipoRetornado)
	}
	if Ast.EsTransferencia(resultado_retornado.Tipo) &&
		m.Tipo != Ast.MATCH_EXPRESION {
		return resultado_retornado
	}

	if resultado_retornado.Tipo == Ast.ERROR {
		return resultado_retornado
	}

	if resultado_retornado.Tipo != Ast.EJECUTADO && m.Tipo == Ast.MATCH_EXPRESION {
		return resultado_retornado
	}

	if m.Tipo == Ast.MATCH_EXPRESION && resultado_retornado.Tipo == Ast.EJECUTADO {
		//Error, el match no esta retornando nada
		msg := "Semantic error, Match statement is not returning any value." +
			" -- Line:" + strconv.Itoa(m.Fila) + " Column: " + strconv.Itoa(m.Columna)
		nError := errores.NewError(m.Fila, m.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Valor: nError,
			Tipo:  Ast.ERROR,
		}
	}

	return Ast.TipoRetornado{
		Tipo:  Ast.EJECUTADO,
		Valor: true,
	}
}

func (c Case) Run(scope *Ast.Scope) interface{} {
	//Crear un nuevo scope y otras variables
	newScope := Ast.NewScope("Case", scope)
	var expresion = false
	//ultimaExpresion := Ast.TipoRetornado{Valor: nil, Tipo: Ast.NULL}
	var ultimoTipo Ast.TipoDato
	var instruccion interface{}
	var resultado Ast.TipoRetornado
	i := 0

	//Verificar el tipo del caso, si es exp o ins
	if c.Tipo == Ast.CASE_EXPRESION {
		expresion = true
	}

	//Ejecutar las instrucciones

	for i = 0; i < c.Instrucciones.Len(); i++ {
		//Verificar el tipo de entrada
		instruccion = c.Instrucciones.GetValue(i).(Ast.Abstracto)
		tipo_abstracto, _ := instruccion.(Ast.Abstracto).GetTipo()
		ultimoTipo = tipo_abstracto
		if tipo_abstracto == Ast.EXPRESION {
			//Si es un if expresión, tiene que retornar algo
			expresion := c.Instrucciones.GetValue(i).(Ast.Expresion)
			resultado = expresion.GetValue(&newScope)
		} else if tipo_abstracto == Ast.INSTRUCCION {
			instruccion := c.Instrucciones.GetValue(i).(Ast.Instruccion)
			resultado = instruccion.Run(&newScope).(Ast.TipoRetornado)
		}

		//Error si viene un break o un return dentro de un case expresion
		if Ast.EsTransferencia(resultado.Tipo) &&
			c.Tipo == Ast.CASE_EXPRESION {
			temp := instruccion.(Ast.Abstracto)
			msg := "Semantic error, transfer statements are not allowed within a case expression statement." +
				" -- Line:" + strconv.Itoa(temp.GetFila()) + " Column: " + strconv.Itoa(temp.GetColumna())
			nError := errores.NewError(temp.GetFila(), temp.GetColumna(), msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			nError.Ambito = scope.GetTipoScope()
			newScope.Errores.Add(nError)
			newScope.Consola += msg + "\n"
			newScope.UpdateScopeGlobal()
			return Ast.TipoRetornado{
				Valor: nError,
				Tipo:  Ast.ERROR,
			}
		}

		if Ast.EsTransferencia(resultado.Tipo) {
			//Retornar la transferencia
			newScope.UpdateScopeGlobal()
			return resultado
		}

	}

	//Termino el for, retornar la ultima expresion
	//Verificar si hay algun retorno o retornar un error
	if ultimoTipo != Ast.EXPRESION && expresion {
		msg := "Semantic error, the match clause is not returning any value." +
			" -- Line:" + strconv.Itoa(c.Fila) + " Column: " + strconv.Itoa(c.Columna)
		nError := errores.NewError(c.Fila, c.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		newScope.UpdateScopeGlobal()
		return Ast.TipoRetornado{
			Valor: nError,
			Tipo:  Ast.ERROR,
		}
	} else if expresion && ultimoTipo == Ast.EXPRESION {
		//Si esta retornado algún valor
		newScope.UpdateScopeGlobal()
		return resultado
	}

	newScope.UpdateScopeGlobal()
	return Ast.TipoRetornado{
		Tipo:  Ast.EJECUTADO,
		Valor: true,
	}
}

func (op Match) GetFila() int {
	return op.Fila
}
func (op Match) GetColumna() int {
	return op.Columna
}
