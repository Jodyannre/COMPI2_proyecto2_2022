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
	Referencia  string
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
	/*****************VARIABLES 3D***********************/
	var codigo3d string
	var obj3dValorAtt, obj3d Ast.O3D
	var inicioStruct string
	var referenciaStruct string
	atributos := arraylist.New()
	/****************************************************/

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
		obj3dValorAtt = valorAtt.Valor.(Ast.O3D)
		codigo3d += obj3dValorAtt.Codigo
		valorAtt = obj3dValorAtt.Valor
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

		//Todo bien ,entonces crear el atributo y agregarlo al struct
		nuevoSimbolo := Ast.NewSimbolo(atributoActual.Nombre, attActual.Valor,
			atributoActual.Fila, atributoActual.Columna,
			attActual.TipoAtributo.Tipo, atributoActual.Mutable, atributoActual.Publico)
		nuevoSimbolo.Publico = attPlantilla.Publico
		nuevoSimbolo.Direccion = newScope.Size
		newScope.Size++
		nuevoSimbolo.TipoDireccion = Ast.HEAP
		//Agregar las referencias
		/*****************************************/
		atributos.Add(obj3dValorAtt.Referencia)
		/*****************************************/
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

	/***************************************/
	inicioStruct = Ast.GetTemp()
	codigo3d += "/************************************NEW STRUCT*/ \n"
	codigo3d += inicioStruct + " = H; //Guardar referencia del inicio del struct\n"
	referenciaStruct = inicioStruct
	codigo3d += "/***********************************************/ \n"

	/***************AGREGAR ATRIBUTOS********************/
	for i := 0; i < atributos.Len(); i++ {
		referenciaValor := atributos.GetValue(i).(string)
		/***************************COD 3D PARA CREAR EL ATRIBUTO*************************/
		codigo3d += "/**************************CREAR VALOR ATRIBUTO*/ \n"
		codigo3d += "heap[(int)H] = " + referenciaValor + ";//Agregar atributo\n"
		codigo3d += "H = H + 1;\n"
		Ast.GetH()
		codigo3d += "/***********************************************/ \n"
		/*********************************************************************************/
	}
	/****************************************************/
	/***************************************/

	obj3d.Valor = Ast.TipoRetornado{
		Tipo:  Ast.STRUCT,
		Valor: s,
	}
	obj3d.Codigo = codigo3d
	obj3d.Referencia = referenciaStruct

	return Ast.TipoRetornado{
		Tipo:  Ast.STRUCT,
		Valor: obj3d,
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
	direccionStructOriginal := s.Referencia
	posicionActual := Ast.GetTemp()
	nuevaDireccion := Ast.GetTemp()
	contadorNuevoStruct := Ast.GetTemp()
	elementoActual := Ast.GetTemp()
	tempContadorNuevoStruct := Ast.GetTemp()
	tempPosicionActual := Ast.GetTemp()
	codigo3d := ""
	var obj3d, obj3dPreElemento Ast.O3D

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
	codigo3d += "/**************************CLONAR STRUCT*/\n"
	/*********************INICIALIZAR VARIABLES DEL VIEJO STRUCT****************/
	codigo3d += posicionActual + " = " + direccionStructOriginal + "; //Dir del struct a copiar \n"
	/***************************************************************************/

	/********************CREAR EL NUEVO STRUCT EN 3D****************************/
	codigo3d += "/*************VARIABLES DEL NUEVO STRUCT*/\n"
	codigo3d += nuevaDireccion + " = H; //Nueva direccion \n"
	codigo3d += contadorNuevoStruct + " = H; //Contador para el nuevo struct \n"
	codigo3d += "H = H + " + strconv.Itoa(s.AtributosIn.Len()) + "; //Apartar espacio del struct \n"
	for i := 0; i < s.AtributosIn.Len(); i++ {
		Ast.GetH()
	}
	codigo3d += "/****************************************/\n"
	/***************************************************************************/

	for i := 0; i < s.AtributosIn.Len(); i++ {
		elemento := s.AtributosIn.GetValue(i).(*Atributo)
		tipoElemento := elemento.Tipo
		if expresiones.EsVAS(tipoElemento) {
			codigo3d += elementoActual + " = heap[(int)" + posicionActual + "]; //Get elemento\n"
			preReferencia := elemento.Valor.(Ast.Clones).SetReferencia(elementoActual)
			preElemento := preReferencia.(Ast.Clones).Clonar(scope)
			obj3dPreElemento = preElemento.(Ast.TipoRetornado).Valor.(Ast.O3D)
			elemento.Valor = obj3dPreElemento.Valor
			nElemento = elemento
			codigo3d += "heap[(int)" + contadorNuevoStruct + "] = " + obj3dPreElemento.Referencia + "; //Guardar elemento\n"
			codigo3d += "/******************ACTUALIZAR CONTADORES*/\n"
			codigo3d += tempContadorNuevoStruct + " = " + contadorNuevoStruct + " + 1; //sig pos\n"
			codigo3d += contadorNuevoStruct + " = " + tempContadorNuevoStruct + "; //sig pos\n"
			codigo3d += tempPosicionActual + " = " + posicionActual + " + 1; //sig pos\n"
			codigo3d += posicionActual + " = " + tempPosicionActual + "; //sig pos\n"
			codigo3d += "/****************************************/\n"
		} else {
			codigo3d += elementoActual + " = heap[(int)" + posicionActual + "]; //Get elemento\n"
			codigo3d += "heap[(int)" + contadorNuevoStruct + "] = " + elementoActual + "; //Guardar elemento\n"
			codigo3d += "/******************ACTUALIZAR CONTADORES*/\n"
			codigo3d += tempContadorNuevoStruct + " = " + contadorNuevoStruct + " + 1; //sig pos\n"
			codigo3d += contadorNuevoStruct + " = " + tempContadorNuevoStruct + "; //sig pos\n"
			codigo3d += tempPosicionActual + " = " + posicionActual + " + 1; //sig pos\n"
			codigo3d += posicionActual + " = " + tempPosicionActual + "; //sig pos\n"
			codigo3d += "/****************************************/\n"
			nElemento = elemento
		}
		nAtributosIn.Add(nElemento)
	}

	codigo3d += "/****************************************/\n"
	nS.AtributosIn = nAtributosIn

	obj3d.Valor = Ast.TipoRetornado{
		Tipo:  Ast.STRUCT,
		Valor: nS,
	}
	obj3d.Codigo = codigo3d
	obj3d.Referencia = nuevaDireccion
	return Ast.TipoRetornado{
		Tipo:  Ast.STRUCT,
		Valor: obj3d,
	}
}

func (s StructInstancia) SetReferencia(referencia string) interface{} {
	s.Referencia = referencia
	return s
}
