package simbolos

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"Back/analizador/fn_array"
	"strconv"
	"strings"

	"github.com/colegno/arraylist"
)

type StructInstancia struct {
	Plantilla   Ast.TipoRetornado
	Tipo        Ast.TipoDato
	Mutable     bool
	Entorno     *Ast.Scope
	AtributosIn *arraylist.List
	Fila        int
	Columna     int
}

func NewStructInstancia(plantilla Ast.TipoRetornado, atributos *arraylist.List, mutable bool, fila, columna int) StructInstancia {
	//Variables para la validación de tipos
	nS := StructInstancia{
		Plantilla:   plantilla,
		Tipo:        Ast.STRUCT,
		Mutable:     mutable,
		AtributosIn: atributos,
		Fila:        fila,
		Columna:     columna,
	}
	return nS
}

func (s StructInstancia) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	var plantilla StructTemplate
	var nombreNewScope string
	var simboloPlantilla Ast.Simbolo
	var resultadoAccesoModulo Ast.TipoRetornado
	if s.Plantilla.Tipo != Ast.STRUCT && s.Plantilla.Tipo != Ast.ACCESO_MODULO {
		//Error, porque la plantilla no es struct
		fila := s.Fila
		columna := s.Columna
		msg := "Semantic error, an STRUCT was expected." +
			" -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
		nError := errores.NewError(fila, columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}
	if s.Plantilla.Tipo == Ast.STRUCT {
		//simboloPlantilla = scope.Exist_fms(s.Plantilla.Valor.(string))
		simboloPlantilla = scope.Exist_fms_local(s.Plantilla.Valor.(string))
		nombreNewScope = s.Plantilla.Valor.(string)
	} else {
		//Es un acceso a modulo, ejecutarlo
		resultadoAccesoModulo = s.Plantilla.Valor.(AccesoModulo).GetValue(scope)
		if resultadoAccesoModulo.Tipo == Ast.ERROR {
			return resultadoAccesoModulo
		} else if resultadoAccesoModulo.Tipo != Ast.STRUCT_TEMPLATE {
			//Error se esperaba un struct
			fila := s.Fila
			columna := s.Columna
			msg := "Semantic error, an STRUCT was expected." +
				" -- Line: " + strconv.Itoa(fila) +
				" Column: " + strconv.Itoa(columna)
			nError := errores.NewError(fila, columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			nError.Ambito = scope.GetTipoScope()
			scope.Errores.Add(nError)
			scope.Consola += msg + "\n"
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}
		}
		simboloPlantilla = resultadoAccesoModulo.Valor.(Ast.Simbolo)
		nombreNewScope = simboloPlantilla.Identificador
	}

	newScope := Ast.NewScope(nombreNewScope, scope)

	//Verificar que la plantilla exista o que no haya algún tipo de error
	if simboloPlantilla.Tipo == Ast.ERROR_NO_EXISTE {
		fila := s.Fila
		columna := s.Columna
		msg := "Semantic error, \"" + s.Plantilla.Valor.(string) + "\" doesn't exist." +
			" -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
		nError := errores.NewError(fila, columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}

	if simboloPlantilla.Tipo == Ast.ERROR_ACCESO_PRIVADO {
		fila := s.Fila
		columna := s.Columna
		msg := "Semantic error, \"" + s.Plantilla.Valor.(string) + "\" is private to this scope." +
			" -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
		nError := errores.NewError(fila, columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}

	//Verificar que sea un tipo struct
	if simboloPlantilla.Valor.(Ast.TipoRetornado).Tipo != Ast.STRUCT {
		fila := s.Fila
		columna := s.Columna
		msg := "Semantic error, \"" + s.Plantilla.Valor.(string) + "\" isn't a STRUCT." +
			" -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
		nError := errores.NewError(fila, columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}

	}
	//Recuperar el struct
	plantilla = simboloPlantilla.Valor.(Ast.TipoRetornado).Valor.(StructTemplate)

	if plantilla.AtributosIn.Len() != s.AtributosIn.Len() {
		//Error la cantidad de atributos no concuerda
		fila := s.Fila
		columna := s.Columna
		msg := "Semantic error, expected (" + strconv.Itoa(plantilla.AtributosIn.Len()) + ") values, found (" +
			strconv.Itoa(s.AtributosIn.Len()) + ") values." +
			" -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
		nError := errores.NewError(fila, columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}

	//Verificar si los atributos existen y si son del mismo tipo
	for i := 0; i < s.AtributosIn.Len(); i++ {
		atributoActual := s.AtributosIn.GetValue(i).(*Atributo)
		_, ok := plantilla.Atributos[atributoActual.Nombre]
		//Verificar que exista el atributo
		if !ok {
			//Error, ese nombre de atributo no existe
			//struct `Personaje` has no field named `vida`
			fila := atributoActual.Fila
			columna := atributoActual.Columna
			msg := "Semantic error, STRUCT " + plantilla.Nombre + " has no field named \"" + atributoActual.Nombre + "\"" +
				" -- Line: " + strconv.Itoa(fila) +
				" Column: " + strconv.Itoa(columna)
			nError := errores.NewError(fila, columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			nError.Ambito = scope.GetTipoScope()
			scope.Errores.Add(nError)
			scope.Consola += msg + "\n"
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}
		}
		//Verificar si los tipos son correctos
		//Variables para validar el tipo

		//Get el valor del atributo
		valorAtt := atributoActual.GetValue(scope)
		if valorAtt.Tipo == Ast.ERROR {
			return valorAtt
		}
		attActual := valorAtt.Valor.(*Atributo)
		if valorAtt.Tipo == Ast.ERROR {
			return valorAtt
		}

		attPlantilla := plantilla.Atributos[atributoActual.Nombre]
		//Verificar que el atributo sea public o error
		if !attPlantilla.Publico {
			fila := attActual.Fila
			columna := attActual.Columna
			msg := "Semantic error, field \"" + attPlantilla.Nombre + "\" of struct " + plantilla.Nombre +
				" is private." +
				" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
			nError := errores.NewError(fila, columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			nError.Ambito = scope.GetTipoScope()
			scope.Errores.Add(nError)
			scope.Consola += msg + "\n"
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}
		}

		validadorTipo := CompararTipos(attActual.TipoAtributo,
			attPlantilla.TipoAtributo)
		if !validadorTipo {
			if attPlantilla.TipoAtributo.Tipo == Ast.DIMENSION_ARRAY {
				listaPrePlantilla := attPlantilla.TipoAtributo.Valor.(expresiones.DimensionArray).GetValue(scope)
				listaEntrante := fn_array.ConcordanciaArray(attActual.Valor.(Ast.TipoRetornado).Valor.(expresiones.Array))
				arrayDimension := arraylist.New()
				for i := 0; i < listaPrePlantilla.Valor.(*arraylist.List).Len(); i++ {
					arrayDimension.Add(listaPrePlantilla.Valor.(*arraylist.List).GetValue(i).(Ast.TipoRetornado).Valor)
				}
				split := strings.Split(listaEntrante, ",")
				//Crear la lista con las posiciones
				listaDimensiones := arraylist.New()
				for _, num := range split {
					numero, _ := strconv.Atoi(num)
					listaDimensiones.Add(numero)
				}

				//Comparar las lista de dimensiones
				if !fn_array.CompararListas(listaDimensiones, arrayDimension) {
					fila := attActual.Valor.(Ast.TipoRetornado).Valor.(Ast.Abstracto).GetFila()
					columna := attActual.Valor.(Ast.TipoRetornado).Valor.(Ast.Abstracto).GetColumna()
					msg := "Semantic error, ARRAY dimension does not match" +
						" -- Line:" + strconv.Itoa(fila) + " Column: " + strconv.Itoa(columna)
					nError := errores.NewError(fila, columna, msg)
					nError.Tipo = Ast.ERROR_SEMANTICO
					nError.Ambito = scope.GetTipoScope()
					scope.Errores.Add(nError)
					scope.Consola += msg + "\n"
					return Ast.TipoRetornado{
						Tipo:  Ast.ERROR,
						Valor: nError,
					}
				}
				//Ahora comparar los tipos
				//Recuperar el tipo de la plantilla
				tipoArrayPlantilla := Ast.TipoRetornado{
					Tipo:  Ast.ARRAY,
					Valor: attPlantilla.TipoAtributo.Valor.(expresiones.DimensionArray).TipoArray,
				}
				validadorTipo := CompararTipos(attActual.TipoAtributo,
					tipoArrayPlantilla)
				if !validadorTipo {
					return GetmsjError(validadorTipo, attActual, attPlantilla, scope)
				}

			} else {
				return GetmsjError(validadorTipo, attActual, attPlantilla, scope)
			}
		}

		//Todo bien ,entonces crear el struct
		nuevoSimbolo := Ast.NewSimbolo(atributoActual.Nombre, attActual.Valor,
			atributoActual.Fila, atributoActual.Columna,
			attActual.TipoAtributo.Tipo, atributoActual.Mutable, atributoActual.Publico)
		nuevoSimbolo.Publico = attPlantilla.Publico
		//Si el tipo es acceso modulo, modificarlo y dejarlo como un struct

		newScope.Add(nuevoSimbolo)
	}
	//Agregar el scope al struct y finalizar retornando el scope
	s.Entorno = &newScope
	if s.Plantilla.Tipo == Ast.ACCESO_MODULO {
		//Ejecutar el acceso y cambiar el tipo de la declaración
		nTipo := GetTipoEstructura(s.Plantilla, scope, s)
		errors := ErrorEnTipo(nTipo)
		if errors.Tipo == Ast.ERROR {
			msg := "Semantic error, type error." +
				" -- Line:" + strconv.Itoa(s.Fila) + " Column: " + strconv.Itoa(s.Columna)
			nError := errores.NewError(s.Fila, s.Columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			nError.Ambito = scope.GetTipoScope()
			scope.Errores.Add(nError)
			scope.Consola += msg + "\n"
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}
		}
		//De lo contrario actualizar el tipo de la declaracion
		s.Plantilla = nTipo
	}

	return Ast.TipoRetornado{
		Tipo:  Ast.STRUCT,
		Valor: s,
	}
}

func GetmsjError(validadorTipo bool, atributo, template *Atributo, scope *Ast.Scope) Ast.TipoRetornado {
	fila := atributo.Fila
	columna := atributo.Columna
	msg := ""
	msg = "Semantic error,  can't assign " + Tipo_String(atributo.TipoAtributo) +
		" to field named \"" + template.Nombre + "\" (" + Tipo_String(template.TipoAtributo) +
		") -- Line: " + strconv.Itoa(fila) +
		" Column: " + strconv.Itoa(columna)
	nError := errores.NewError(fila, columna, msg)
	nError.Tipo = Ast.ERROR_SEMANTICO
	nError.Ambito = scope.GetTipoScope()
	scope.Errores.Add(nError)
	scope.Consola += msg + "\n"
	return Ast.TipoRetornado{
		Tipo:  Ast.ERROR,
		Valor: nError,
	}

}

func (v StructInstancia) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, v.Tipo
}

func (v StructInstancia) GetFila() int {
	return v.Fila
}
func (v StructInstancia) GetColumna() int {
	return v.Columna
}

func (s StructInstancia) GetPlantilla(scope *Ast.Scope) string {
	if s.Plantilla.Tipo == Ast.ACCESO_MODULO {
		resultado := s.Plantilla.Valor.(AccesoModulo).GetTipoFromAccesoModulo(s.Plantilla, scope)
		if resultado.Tipo == Ast.ERROR {
			return "ERROR"
		} else {
			return resultado.Valor.(string)
		}
	}
	return s.Plantilla.Valor.(string)
}

func (s StructInstancia) SetMutabilidad(mutable bool) interface{} {
	s.Mutable = mutable
	return s
}

func (s StructInstancia) GetMutable() bool {
	return s.Mutable
}

func (s StructInstancia) Clonar(scope *Ast.Scope) interface{} {
	nAtributosIn := arraylist.New()
	var nElemento *Atributo

	nS := StructInstancia{
		Plantilla: s.Plantilla,
		Entorno:   s.Entorno.Clonar(scope).(*Ast.Scope),
		Fila:      s.Fila,
		Columna:   s.Columna,
		Mutable:   s.Mutable,
		Tipo:      s.Tipo,
	}
	//Copiar la lista de atributos
	for i := 0; i < s.AtributosIn.Len(); i++ {
		elemento := s.AtributosIn.GetValue(i).(*Atributo)
		tipoElemento := elemento.Tipo
		if expresiones.EsVAS(tipoElemento) {
			preElemento := elemento.Valor.(Ast.Clones).Clonar(scope)
			valor := preElemento
			nElemento = valor.(*Atributo)
		} else {
			nElemento = elemento
		}
		nAtributosIn.Add(nElemento)
	}
	nS.AtributosIn = nAtributosIn

	return nS
}
