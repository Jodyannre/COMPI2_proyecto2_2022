package instrucciones

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
)

type Declaracion struct {
	Id          string
	Tipo        Ast.TipoDato
	TipoRetorno Ast.TipoDato
	Mutable     bool
	Publico     bool
	Valor       interface{}
	Fila        int
	Columna     int
}

func NewDeclaracion(id string, tipo Ast.TipoDato, mutable, publico bool, tipoRetorno Ast.TipoDato,
	valor interface{}, fila int, columna int) Declaracion {
	nd := Declaracion{
		Id:          id,
		Tipo:        tipo,
		TipoRetorno: tipoRetorno,
		Mutable:     mutable,
		Publico:     publico,
		Valor:       valor,
		Fila:        fila,
		Columna:     columna,
	}
	return nd
}

func (d Declaracion) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, Ast.DECLARACION
}

func (d Declaracion) Run(scope *Ast.Scope) interface{} {

	//Verificar que el id no exista

	existe := scope.Exist_actual(d.Id)

	//Primero verificar que no es un if expresion
	_, tipoIn := d.Valor.(Ast.Abstracto).GetTipo()

	//Verificar si es un struct y si el tipo de la variable es indefinido o error
	if d.Tipo != Ast.INDEFINIDO && tipoIn == Ast.STRUCT {
		////////////////////////////ERROR//////////////////////////////////
		return errores.GenerarError(13, d, d, "",
			Ast.ValorTipoDato[d.Tipo],
			Ast.ValorTipoDato[tipoIn],
			scope)
		//////////////////////////////////////////////////////////////////
	}

	//Verificar que sea un primitivo i64 y la declaración sea usize

	var preValor interface{}
	if tipoIn == Ast.IF_EXPRESION || tipoIn == Ast.MATCH_EXPRESION || tipoIn == Ast.LOOP_EXPRESION {
		preValor = d.Valor.(Ast.Instruccion).Run(scope)
	} else if tipoIn == Ast.FUNCION {
		preValor = Ast.TipoRetornado{
			Valor: d.Valor,
			Tipo:  Ast.FUNCION,
		}
	} else {

		preValor = d.Valor.(Ast.Expresion).GetValue(scope)
	}
	valor := preValor.(Ast.TipoRetornado)

	//Cambiar valor de i64 a usize si la declaración es usize y el valor que viene es un i64
	if d.Tipo == Ast.USIZE && tipoIn == Ast.I64 {
		valor.Tipo = Ast.USIZE
	}

	//Revisar si el retorno es un error
	if valor.Tipo == Ast.ERROR {
		return valor
	}

	//Revisar si se declara un vec o array y la expresion es diferente
	if d.Tipo == Ast.VECTOR && valor.Tipo != Ast.VECTOR ||
		d.Tipo == Ast.ARRAY && valor.Tipo != Ast.ARRAY {
		//Error, no se puede inicializar un vector con un valor
		////////////////////////////ERROR//////////////////////////////////
		return errores.GenerarError(13, d, d, "",
			Ast.ValorTipoDato[d.Tipo],
			Ast.ValorTipoDato[valor.Tipo],
			scope)
		//////////////////////////////////////////////////////////////////
	}

	if valor.Tipo == d.Tipo && !existe {
		//El tipo es correcto y no existe en el entorno actual
		//Crear símbolo y agregarlo a la tabla del entorno actual
		nSimbolo := Ast.Simbolo{
			Identificador: d.Id,
			Valor:         valor,
			Fila:          d.Fila,
			Columna:       d.Columna,
			Tipo:          d.Tipo,
			Mutable:       d.Mutable,
			Publico:       d.Publico,
			Entorno:       scope,
		}
		//Si es vector o array verificar si es referencia o no
		if valor.Tipo == Ast.VECTOR {
			//Verificar que el tipo del vector a agregar es correcto con el vector esperado
			nValor := valor.Valor.(expresiones.Vector)
			vectorCorrecto := TipoVectorCorrecto(d.TipoRetorno, nValor.Tipo)
			if vectorCorrecto.Tipo == Ast.ERROR {
				if vectorCorrecto.Valor == 1 {
					//No tiene ningún tipo
					////////////////////////////ERROR//////////////////////////////////
					return errores.GenerarError(37, d, d, "",
						Ast.ValorTipoDato[nValor.Tipo],
						"",
						scope)
					//////////////////////////////////////////////////////////////////
				}
				if vectorCorrecto.Valor == 2 {
					//Tipos diferentes de declaración y creación
					////////////////////////////ERROR//////////////////////////////////
					return errores.GenerarError(38, d, d, "",
						Ast.ValorTipoDato[d.TipoRetorno],
						Ast.ValorTipoDato[nValor.Tipo],
						scope)
					//////////////////////////////////////////////////////////////////
				}
			}

			if nValor.Tipo == Ast.INDEFINIDO {
				nValor.Tipo = d.TipoRetorno
			}
			nSimbolo.Valor = Ast.TipoRetornado{Tipo: Ast.VECTOR, Valor: nValor}
		}

		//Si es un array

		if valor.Tipo == Ast.ARRAY {
			//Verificar que el tipo del vector a agregar es correcto con el vector esperado
			nValor := valor.Valor.(expresiones.Array)
			vectorCorrecto := TipoVectorCorrecto(d.TipoRetorno, nValor.TipoArray)
			if vectorCorrecto.Tipo == Ast.ERROR {
				if vectorCorrecto.Valor == 1 {
					//No tiene ningún tipo
					////////////////////////////ERROR//////////////////////////////////
					return errores.GenerarError(39, d, d, "",
						Ast.ValorTipoDato[nValor.TipoDelArray.Tipo],
						"",
						scope)
					//////////////////////////////////////////////////////////////////

				}
				if vectorCorrecto.Valor == 2 {
					//Tipos diferentes de declaración y creación
					////////////////////////////ERROR//////////////////////////////////
					return errores.GenerarError(40, d, d, "",
						Ast.ValorTipoDato[d.TipoRetorno],
						Ast.ValorTipoDato[nValor.TipoDelArray.Tipo],
						scope)
					//////////////////////////////////////////////////////////////////
				}
			}

			nSimbolo.Valor = Ast.TipoRetornado{Tipo: Ast.ARRAY, Valor: nValor}
		}
		//Si es función, módulo o struct, agregarlos a las listas globales
		scope.Add(nSimbolo)
		if valor.Tipo == Ast.FUNCION ||
			valor.Tipo == Ast.MODULO ||
			valor.Tipo == Ast.STRUCT {
			scope.Addfms(nSimbolo)
		}

	} else if d.Tipo == Ast.INDEFINIDO && !existe {
		//Es una declaración sin valor asignado
		nSimbolo := Ast.Simbolo{
			Identificador: d.Id,
			Valor:         valor,
			Fila:          d.Fila,
			Columna:       d.Columna,
			Tipo:          valor.Tipo,
			Mutable:       d.Mutable,
			Publico:       d.Publico,
			Entorno:       scope,
		}
		if valor.Tipo == Ast.VECTOR {
			nValor := valor.Valor.(expresiones.Vector)

			vectorCorrecto := TipoVectorCorrecto(d.TipoRetorno, nValor.Tipo)
			if vectorCorrecto.Tipo == Ast.ERROR {
				if vectorCorrecto.Valor == 1 {
					//No tiene ningún tipo
					////////////////////////////ERROR//////////////////////////////////
					return errores.GenerarError(37, d, d, "",
						Ast.ValorTipoDato[nValor.Tipo],
						"",
						scope)
					//////////////////////////////////////////////////////////////////
				}
				if vectorCorrecto.Valor == 2 {
					//Tipos diferentes de declaración y creación
					////////////////////////////ERROR//////////////////////////////////
					return errores.GenerarError(38, d, d, "",
						Ast.ValorTipoDato[d.TipoRetorno],
						Ast.ValorTipoDato[nValor.Tipo],
						scope)
					//////////////////////////////////////////////////////////////////
				}
			}

			if nValor.Tipo == Ast.INDEFINIDO {
				nValor.Tipo = d.Tipo
			}
			nSimbolo.Valor = Ast.TipoRetornado{Tipo: Ast.VECTOR, Valor: nValor}
		}
		if valor.Tipo == Ast.ARRAY {
			nValor := valor.Valor.(expresiones.Array)

			arrayCorrecto := TipoVectorCorrecto(d.TipoRetorno, nValor.TipoDelArray.Tipo)
			if arrayCorrecto.Tipo == Ast.ERROR {
				if arrayCorrecto.Valor == 1 {
					//No tiene ningún tipo
					////////////////////////////ERROR//////////////////////////////////
					return errores.GenerarError(39, d, d, "",
						Ast.ValorTipoDato[nValor.TipoArray],
						"",
						scope)
					//////////////////////////////////////////////////////////////////
				}
				if arrayCorrecto.Valor == 2 {
					//Tipos diferentes de declaración y creación
					////////////////////////////ERROR//////////////////////////////////
					return errores.GenerarError(40, d, d, "",
						Ast.ValorTipoDato[nValor.TipoDelArray.Tipo],
						Ast.ValorTipoDato[nValor.TipoArray],
						scope)
					//////////////////////////////////////////////////////////////////
				}
			}

			nSimbolo.Valor = Ast.TipoRetornado{Tipo: Ast.ARRAY, Valor: nValor}
		}

		scope.Add(nSimbolo)
	} else if d.Tipo != Ast.INDEFINIDO && !existe && valor.Tipo == Ast.NULL {
		//Es una declaración sin valor asignado
		nSimbolo := Ast.Simbolo{
			Identificador: d.Id,
			Valor:         valor,
			Fila:          d.Fila,
			Columna:       d.Columna,
			Tipo:          valor.Tipo,
			Mutable:       d.Mutable,
			Publico:       d.Publico,
			Entorno:       scope,
		}
		scope.Add(nSimbolo)
	} else if existe {
		//Ya existe y generar error semántico
		////////////////////////////ERROR//////////////////////////////////
		return errores.GenerarError(12, d, d, d.Id,
			"",
			"",
			scope)
		//////////////////////////////////////////////////////////////////
	} else {
		//Error de tipos, generar error semántico
		////////////////////////////ERROR//////////////////////////////////
		return errores.GenerarError(24, d, d, "",
			Ast.ValorTipoDato[int(valor.Tipo)],
			Ast.ValorTipoDato[int(d.Tipo)],
			scope)
		//////////////////////////////////////////////////////////////////
	}
	return Ast.TipoRetornado{
		Tipo:  Ast.EJECUTADO,
		Valor: true,
	}
}

func (op Declaracion) GetFila() int {
	return op.Fila
}
func (op Declaracion) GetColumna() int {
	return op.Columna
}

func TipoVectorCorrecto(tipoDeclaracion Ast.TipoDato, tipoRetornado Ast.TipoDato) Ast.TipoRetornado {
	if tipoDeclaracion == Ast.VOID && tipoRetornado != Ast.INDEFINIDO {
		//Retornar correcto
		return Ast.TipoRetornado{
			Tipo:  Ast.BOOLEAN,
			Valor: true,
		}
	}
	if tipoDeclaracion == Ast.VOID && tipoRetornado == Ast.INDEFINIDO {
		//Error, el vector no tiene ningún tipo
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: 1, // No tiene ningún tipo
		}
	}
	if tipoDeclaracion != Ast.VOID && tipoRetornado == Ast.INDEFINIDO {
		return Ast.TipoRetornado{
			Tipo:  Ast.BOOLEAN,
			Valor: true,
		}
	}
	if tipoDeclaracion != Ast.VOID && tipoRetornado != Ast.INDEFINIDO {
		if tipoDeclaracion != tipoRetornado {
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: 2, // Los tipos son diferentes
			}
		} else {
			return Ast.TipoRetornado{
				Tipo:  Ast.BOOLEAN,
				Valor: true,
			}
		}
	}
	return Ast.TipoRetornado{
		Tipo:  Ast.BOOLEAN,
		Valor: true,
	}
}
