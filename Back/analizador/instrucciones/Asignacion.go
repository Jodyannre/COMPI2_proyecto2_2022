package instrucciones

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"Back/analizador/fn_array"
	"Back/analizador/fn_vectores"
	"strconv"

	"github.com/colegno/arraylist"
)

type Asignacion struct {
	Id      Ast.Expresion
	Valor   interface{}
	Fila    int
	Columna int
}

func NewAsignacion(id Ast.Expresion, valor interface{}, fila int, columna int) Asignacion {
	na := Asignacion{
		Id:      id,
		Valor:   valor,
		Fila:    fila,
		Columna: columna,
	}
	return na
}

func (a Asignacion) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, Ast.ASIGNACION
}

func (a Asignacion) Run(scope *Ast.Scope) interface{} {
	//Conseguir el valor del id y verificar que sea un id
	var id string
	var resultado Ast.TipoRetornado
	_, tipoParticular := a.Id.(Ast.Abstracto).GetTipo()
	//Verificar que sea un identificador
	if tipoParticular != Ast.IDENTIFICADOR && tipoParticular != Ast.VEC_ACCESO &&
		tipoParticular != Ast.ACCESO_ARRAY && tipoParticular != Ast.ACCESO_STRUCT {
		//Error, se espera un identificador. un acceso a vector o un acceso a un array
		msg := "Semantic error, expected IDENTIFICADOR, found " + Ast.ValorTipoDato[tipoParticular] +
			". -- Line: " + strconv.Itoa(a.Id.(Ast.Abstracto).GetFila()) +
			" Column: " + strconv.Itoa(a.Id.(Ast.Abstracto).GetColumna())
		nError := errores.NewError(a.Id.(Ast.Abstracto).GetFila(), a.Id.(Ast.Abstracto).GetColumna(), msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}

	if tipoParticular == Ast.IDENTIFICADOR {
		id = a.Id.(expresiones.Identificador).Valor
		resultado = a.AsignarVariable(id, scope)

	} else if tipoParticular == Ast.VEC_ACCESO {
		valor := a.Id.(fn_vectores.AccesoVec).GetValue(scope)
		if valor.Tipo == Ast.ERROR {
			return valor
		}
		id = a.Id.(fn_vectores.AccesoVec).Identificador.(expresiones.Identificador).Valor
		simbolo := scope.GetSimbolo(id)
		if simbolo.Tipo == Ast.ARRAY {
			acceso_vector := a.Id.(fn_vectores.AccesoVec)
			nLista := arraylist.New()
			nLista.Add(acceso_vector.Posicion)
			acceso_lista := fn_array.NewAccesoArray(acceso_vector.Identificador, nLista, acceso_vector.Fila, acceso_vector.Columna)
			a.Id = acceso_lista
			resultado = a.AsignarAccesoArray(id, scope)
		} else {
			resultado = a.AsignarAccesoVector(id, scope)
		}

	} else if tipoParticular == Ast.ACCESO_ARRAY {
		_, tipoParticular := a.Id.(fn_array.AccesoArray).Identificador.(Ast.Abstracto).GetTipo()
		if tipoParticular != Ast.IDENTIFICADOR {
			msg := "Semantic error, expected IDENTIFICADOR, found " + Ast.ValorTipoDato[tipoParticular] +
				". -- Line: " + strconv.Itoa(a.Id.(Ast.Abstracto).GetFila()) +
				" Column: " + strconv.Itoa(a.Id.(Ast.Abstracto).GetColumna())
			nError := errores.NewError(a.Id.(Ast.Abstracto).GetFila(), a.Id.(Ast.Abstracto).GetColumna(), msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			nError.Ambito = scope.GetTipoScope()
			scope.Errores.Add(nError)
			scope.Consola += msg + "\n"
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}
		}
		id := a.Id.(fn_array.AccesoArray).Identificador.(expresiones.Identificador).Valor
		resultado = a.AsignarAccesoArray(id, scope)
	}
	return resultado
}

func (op Asignacion) GetFila() int {
	return op.Fila
}
func (op Asignacion) GetColumna() int {
	return op.Columna
}

func (a Asignacion) AsignarAccesoArray(id string, scope *Ast.Scope) Ast.TipoRetornado {

	var posicion interface{}
	var resultadoAsignacion Ast.TipoRetornado
	var valorPosicion Ast.TipoRetornado
	posiciones := arraylist.New()
	var array interface{}

	/*************VARIABLES 3D*************/
	var obj3dPosicion Ast.O3D
	var idExp expresiones.Identificador
	var obj3d, obj3dValor, objtemp Ast.O3D
	var referencia, codigo3d string
	/**************************************/

	//Verificar que el id  exista
	existe := scope.Exist(id)
	if !existe {
		msg := "Semantic error, the element \"" + id + "\" doesn't exist in any scope." +
			" -- Line:" + strconv.Itoa(a.Fila) + " Column: " + strconv.Itoa(a.Columna)
		nError := errores.NewError(a.Fila, a.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}
	//Obtener el valor del id
	simbolo_id := scope.GetSimbolo(id)

	/*********C3D del Simbolo***********/
	codigo3d += "/********************************ACCESO A VECTOR*/\n"
	idExp = expresiones.NewIdentificador(id, Ast.IDENTIFICADOR, 0, 0)
	obj3d = idExp.GetValue(scope).Valor.(Ast.O3D)
	referencia = obj3d.Referencia
	codigo3d += obj3d.Codigo
	/***********************************/

	//Obtener el vector
	if simbolo_id.Tipo == Ast.ARRAY {
		array = simbolo_id.Valor.(Ast.TipoRetornado).Valor.(expresiones.Array)
	} else {
		array = simbolo_id.Valor.(Ast.TipoRetornado).Valor.(expresiones.Vector)
	}

	//Verificar que los tipos sean correctos
	//Primero verificar que no es un if expresion
	_, tipoIn := a.Valor.(Ast.Abstracto).GetTipo()
	var preValor interface{}
	if tipoIn == Ast.IF_EXPRESION || tipoIn == Ast.MATCH_EXPRESION || tipoIn == Ast.LOOP_EXPRESION {
		preValor = a.Valor.(Ast.Instruccion).Run(scope)
	} else {
		preValor = a.Valor.(Ast.Expresion).GetValue(scope)
	}
	valor := preValor.(Ast.TipoRetornado)
	obj3dValor = valor.Valor.(Ast.O3D)
	valor = obj3dValor.Valor

	if existe {
		//Primero verificar si es mutable
		if !simbolo_id.Mutable {
			//No es mutable, error semántico
			msg := "Semantic error, can't modify a non-mutable " + Ast.ValorTipoDato[int(simbolo_id.Tipo)] +
				" type. -- Line: " + strconv.Itoa(a.Fila) +
				" Column: " + strconv.Itoa(a.Columna)
			nError := errores.NewError(a.Fila, a.Columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			nError.Ambito = scope.GetTipoScope()
			scope.Errores.Add(nError)
			scope.Consola += msg + "\n"
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}
		}
		//Primero verificar
		//Existe, ahora verificar los tipos.
		if simbolo_id.Tipo == Ast.ARRAY {
			if valor.Tipo == Ast.ARRAY {
				//Comparar los tipos de vectores
				if !expresiones.CompararTipos(array.(expresiones.Array).TipoDelArray,
					valor.Valor.(expresiones.Array).TipoDelArray) {
					msg := "Semantic error, can't assign " + Ast.ValorTipoDato[int(valor.Tipo)] +
						" type to ARRAY[" + Ast.ValorTipoDato[array.(expresiones.Array).TipoDelArray.Tipo] + "]" +
						" type. -- Line: " + strconv.Itoa(a.Fila) +
						" Column: " + strconv.Itoa(a.Columna)
					nError := errores.NewError(a.Fila, a.Columna, msg)
					nError.Tipo = Ast.ERROR_SEMANTICO
					nError.Ambito = scope.GetTipoScope()
					scope.Errores.Add(nError)
					scope.Consola += msg + "\n"
					return Ast.TipoRetornado{
						Tipo:  Ast.ERROR,
						Valor: nError,
					}

				}
			}

			if expresiones.GetTipoFinal(array.(expresiones.Array).TipoDelArray).Tipo == valor.Tipo {
				//Los tipos son correctos, actualizar el símbolo

				//Revisar si es vector y si es del tipo de vector correcto
				if valor.Tipo == Ast.VECTOR {
					vectorEntrante := valor.Valor.(expresiones.Vector)
					vectorGuardado := valor.Valor.(expresiones.Vector)
					if vectorEntrante.TipoVector != vectorGuardado.TipoVector {
						//Hay varias opciones y una es que la lista que entra es indefinida
						//Y la otra es que si traiga un tipo diferente
						if vectorEntrante.Tipo != Ast.INDEFINIDO {
							//Generar el Error, de lo contrario todo bien
							msg := "Semantic error, can't assign Vector<" + Ast.ValorTipoDato[vectorEntrante.Tipo] + ">" +
								" to Vector<" + Ast.ValorTipoDato[vectorGuardado.Tipo] + ">" +
								" type. -- Line: " + strconv.Itoa(a.Fila) +
								" Column: " + strconv.Itoa(a.Columna)
							nError := errores.NewError(a.Fila, a.Columna, msg)
							nError.Tipo = Ast.ERROR_SEMANTICO
							nError.Ambito = scope.GetTipoScope()
							scope.Errores.Add(nError)
							scope.Consola += msg + "\n"
							return Ast.TipoRetornado{
								Tipo:  Ast.ERROR,
								Valor: nError,
							}
						} else {
							//Copiar los valores del vector guardado al nuevo vector entrante
							CopiarVector(&vectorGuardado, &vectorEntrante, simbolo_id)
							valor = Ast.TipoRetornado{
								Tipo:  Ast.VECTOR,
								Valor: vectorEntrante,
							}
						}
					} else {
						CopiarVector(&vectorGuardado, &vectorEntrante, simbolo_id)
						valor = Ast.TipoRetornado{
							Tipo:  Ast.VECTOR,
							Valor: vectorEntrante,
						}
					}
				}
				//Revisar si el struct es del mismo tipo
				//Primero traer el tipo del símbolo declarado
				if valor.Tipo == Ast.STRUCT {
					tipoSimbolo := expresiones.GetTipoFinal(array.(expresiones.Array).TipoDelArray).Valor.(string)
					tipoValor := valor.Valor.(Ast.Structs).GetPlantilla(scope)

					if tipoSimbolo != tipoValor {
						//Error, los structs no son iguales
						fila := valor.Valor.(Ast.Abstracto).GetFila()
						columna := valor.Valor.(Ast.Abstracto).GetColumna()
						msg := "Semantic error, can't store " + tipoValor +
							" to an ARRAY[" + tipoSimbolo + "]" +
							". -- Line: " + strconv.Itoa(fila) +
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

				}
				//////////////////////////////////////////////////////////////////////////
				//Otro error, estoy copiando el puntero y lo modifico más adelante
				puntero := a.Id.(fn_array.AccesoArray).Posiciones
				prePos := puntero.Clone()
				//Get las posiciones
				for i := 0; i < prePos.Len(); i++ {
					posicion = prePos.GetValue(i)
					/*******************************************************/
					valorPosicion = posicion.(Ast.Expresion).GetValue(scope)
					obj3dPosicion = valorPosicion.Valor.(Ast.O3D)
					valorPosicion = obj3dPosicion.Valor
					/*******************************************************/
					_, tipoParticular := posicion.(Ast.Abstracto).GetTipo()
					if valorPosicion.Tipo == Ast.ERROR {
						return posicion.(Ast.TipoRetornado)
					}
					//Verificar que el número en el acceso sea usize
					resultado := expresiones.EsUsize(valorPosicion, tipoParticular, posicion, scope)
					if resultado.Tipo == Ast.ERROR {
						return resultado
					}
					posiciones.Add(valorPosicion.Valor)
				}

				//Buscar la posición
				resultadoAsignacion = fn_array.UpdateElemento(array.(expresiones.Array), prePos, posiciones, scope, valor)
				if resultadoAsignacion.Tipo == Ast.ERROR {
					return resultadoAsignacion
				}

				simbolo_id.Valor = Ast.TipoRetornado{Tipo: Ast.ARRAY, Valor: array}
				scope.UpdateSimbolo(id, simbolo_id)
			} else {
				//Revisar si el retorno es un error
				if valor.Tipo == Ast.ERROR {
					return valor
				}
				//Error de tipos, generar un error semántico
				//fmt.Println("Erro de tipos")
				msg := "Semantic error, can't assign " + Ast.ValorTipoDato[int(valor.Tipo)] +
					" type to ARRAY[" + Ast.ValorTipoDato[array.(expresiones.Array).TipoDelArray.Tipo] + "]" +
					" type. -- Line: " + strconv.Itoa(a.Fila) +
					" Column: " + strconv.Itoa(a.Columna)
				nError := errores.NewError(a.Fila, a.Columna, msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				nError.Ambito = scope.GetTipoScope()
				scope.Errores.Add(nError)
				scope.Consola += msg + "\n"
				return Ast.TipoRetornado{
					Tipo:  Ast.ERROR,
					Valor: nError,
				}
			}
		} else {
			if valor.Tipo == Ast.VECTOR {
				//Comparar los tipos de vectores
				if !expresiones.CompararTipos(array.(expresiones.Vector).TipoVector,
					valor.Valor.(expresiones.Vector).TipoVector) {
					//Erro de tipos
					msg := "Semantic error, can't assign " + Ast.ValorTipoDato[int(valor.Tipo)] +
						" type to Vector<" + Ast.ValorTipoDato[array.(expresiones.Vector).TipoVector.Tipo] + ">" +
						" type. -- Line: " + strconv.Itoa(a.Fila) +
						" Column: " + strconv.Itoa(a.Columna)
					nError := errores.NewError(a.Fila, a.Columna, msg)
					nError.Tipo = Ast.ERROR_SEMANTICO
					nError.Ambito = scope.GetTipoScope()
					scope.Errores.Add(nError)
					scope.Consola += msg + "\n"
					return Ast.TipoRetornado{
						Tipo:  Ast.ERROR,
						Valor: nError,
					}
				}
			}

			if expresiones.GetTipoFinal(array.(expresiones.Vector).TipoVector).Tipo == valor.Tipo {
				//Los tipos son correctos en el fondo

				//Revisar si es vector y si es del tipo de vector correcto
				if valor.Tipo == Ast.VECTOR {
					vectorEntrante := valor.Valor.(expresiones.Vector)
					vectorGuardado := valor.Valor.(expresiones.Vector)
					if vectorEntrante.TipoVector != vectorGuardado.TipoVector {
						//Hay varias opciones y una es que la lista que entra es indefinida
						//Y la otra es que si traiga un tipo diferente
						if vectorEntrante.Tipo != Ast.INDEFINIDO {
							//Generar el Error, de lo contrario todo bien
							msg := "Semantic error, can't assign Vector<" + Ast.ValorTipoDato[vectorEntrante.Tipo] + ">" +
								" to Vector<" + Ast.ValorTipoDato[vectorGuardado.Tipo] + ">" +
								" type. -- Line: " + strconv.Itoa(a.Fila) +
								" Column: " + strconv.Itoa(a.Columna)
							nError := errores.NewError(a.Fila, a.Columna, msg)
							nError.Tipo = Ast.ERROR_SEMANTICO
							nError.Ambito = scope.GetTipoScope()
							scope.Errores.Add(nError)
							scope.Consola += msg + "\n"
							return Ast.TipoRetornado{
								Tipo:  Ast.ERROR,
								Valor: nError,
							}
						} else {
							//Copiar los valores del vector guardado al nuevo vector entrante
							CopiarVector(&vectorGuardado, &vectorEntrante, simbolo_id)
							valor = Ast.TipoRetornado{
								Tipo:  Ast.VECTOR,
								Valor: vectorEntrante,
							}
						}
					} else {
						CopiarVector(&vectorGuardado, &vectorEntrante, simbolo_id)
						valor = Ast.TipoRetornado{
							Tipo:  Ast.VECTOR,
							Valor: vectorEntrante,
						}
					}
				}
				//Revisar si el struct es del mismo tipo
				//Primero traer el tipo del símbolo declarado
				if valor.Tipo == Ast.STRUCT {
					tipoSimbolo := expresiones.GetTipoFinal(array.(expresiones.Vector).TipoVector).Valor.(string)
					tipoValor := valor.Valor.(Ast.Structs).GetPlantilla(scope)

					if tipoSimbolo != tipoValor {
						//Error, los structs no son iguales
						fila := valor.Valor.(Ast.Abstracto).GetFila()
						columna := valor.Valor.(Ast.Abstracto).GetColumna()
						msg := "Semantic error, can't store " + tipoValor +
							" to an ARRAY[" + tipoSimbolo + "]" +
							". -- Line: " + strconv.Itoa(fila) +
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

				}

				prePos := a.Id.(fn_array.AccesoArray).Posiciones
				//Get las posiciones
				for i := 0; i < prePos.Len(); i++ {
					posicion = prePos.GetValue(i)
					valorPosicion = posicion.(Ast.Expresion).GetValue(scope)
					objtemp = valorPosicion.Valor.(Ast.O3D)
					valorPosicion = objtemp.Valor
					_, tipoParticular := posicion.(Ast.Abstracto).GetTipo()
					if valorPosicion.Tipo == Ast.ERROR {
						return posicion.(Ast.TipoRetornado)
					}
					//Verificar que el número en el acceso sea usize
					resultado := expresiones.EsUsize(valorPosicion, tipoParticular, posicion, scope)
					if resultado.Tipo == Ast.ERROR {
						return resultado
					}
					posiciones.Add(valorPosicion.Valor)
				}

				//Buscar la posición
				//Crear un nuevo valor
				//nuevoValor := valor
				resultadoAsignacion = fn_vectores.UpdateElemento(array.(expresiones.Vector), prePos, posiciones, scope, obj3dValor, referencia)
				objtemp = resultadoAsignacion.Valor.(Ast.O3D)
				resultadoAsignacion = objtemp.Valor
				if resultadoAsignacion.Tipo == Ast.ERROR {
					return resultadoAsignacion
				}
				codigo3d += objtemp.Codigo
				simbolo_id.Valor = Ast.TipoRetornado{Tipo: Ast.VECTOR, Valor: array}
				scope.UpdateSimbolo(id, simbolo_id)
			} else {
				//Revisar si el retorno es un error
				if valor.Tipo == Ast.ERROR {
					return valor
				}
				//Error de tipos, generar un error semántico
				//fmt.Println("Erro de tipos")
				msg := "Semantic error, can't assign " + Ast.ValorTipoDato[int(valor.Tipo)] +
					" type to VECTOR<" + Ast.ValorTipoDato[array.(expresiones.Vector).TipoVector.Tipo] + ">" +
					" type. -- Line: " + strconv.Itoa(a.Fila) +
					" Column: " + strconv.Itoa(a.Columna)
				nError := errores.NewError(a.Fila, a.Columna, msg)
				nError.Tipo = Ast.ERROR_SEMANTICO
				nError.Ambito = scope.GetTipoScope()
				scope.Errores.Add(nError)
				scope.Consola += msg + "\n"
				return Ast.TipoRetornado{
					Tipo:  Ast.ERROR,
					Valor: nError,
				}
			}

		}

	} else {
		//No existe, generar un error semántico
		msg := "Semantic error, the element \"" + id + "\" doesn't exist in any scope." +
			" -- Line:" + strconv.Itoa(a.Fila) + " Column: " + strconv.Itoa(a.Columna)
		nError := errores.NewError(a.Fila, a.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}

	obj3d.Codigo = codigo3d
	obj3d.Valor = Ast.TipoRetornado{
		Tipo:  Ast.EJECUTADO,
		Valor: true,
	}
	return Ast.TipoRetornado{
		Tipo:  Ast.EJECUTADO,
		Valor: obj3d,
	}

}

func (a Asignacion) AsignarVariable(id string, scope *Ast.Scope) Ast.TipoRetornado {
	/*Variables 3d*/
	var obj3d Ast.O3D
	var obj3dValor Ast.O3D
	var codigo3d string
	var referencia string

	//Verificar que el id  exista
	existe := scope.Exist(id)
	//Obtener el valor del id
	simbolo_id := scope.GetSimbolo(id)
	//Obtener el código 3d de conseguir la variable
	obj3d = GenerarC3DGetSimbolo(simbolo_id)
	//Verificar que los tipos sean correctos
	//Primero verificar que no es un if expresion
	_, tipoIn := a.Valor.(Ast.Abstracto).GetTipo()
	var preValor interface{}
	if tipoIn == Ast.IF_EXPRESION || tipoIn == Ast.MATCH_EXPRESION || tipoIn == Ast.LOOP_EXPRESION {
		preValor = a.Valor.(Ast.Instruccion).Run(scope)
	} else {
		preValor = a.Valor.(Ast.Expresion).GetValue(scope)
	}
	obj3dValor = preValor.(Ast.TipoRetornado).Valor.(Ast.O3D)
	valor := obj3dValor.Valor

	if existe {
		//Primero verificar si es mutable
		if !simbolo_id.Mutable {
			//No es mutable, error semántico
			/////////////////////////////ERROR/////////////////////////////////////
			return errores.GenerarError(21, a, a, Ast.ValorTipoDato[int(simbolo_id.Tipo)], "", "", scope)
		}
		//Primero verificar
		//Existe, ahora verificar los tipos
		if simbolo_id.Valor.(Ast.TipoRetornado).Tipo == valor.Tipo {
			//Los tipos son correctos, actualizar el símbolo

			//Revisar si es vector y si es del tipo de vector correcto
			if valor.Tipo == Ast.VECTOR {
				vectorEntrante := valor.Valor.(expresiones.Vector)
				vectorGuardado := simbolo_id.Valor.(Ast.TipoRetornado).Valor.(expresiones.Vector)
				if vectorEntrante.TipoVector != vectorGuardado.TipoVector {
					//Hay varias opciones y una es que la lista que entra es indefinida
					//Y la otra es que si traiga un tipo diferente
					if vectorEntrante.Tipo != Ast.INDEFINIDO {
						//Generar el Error, de lo contrario todo bien
						/////////////////////////////ERROR/////////////////////////////////////
						return errores.GenerarError(22, a, a, "", Ast.ValorTipoDato[vectorEntrante.Tipo],
							Ast.ValorTipoDato[vectorGuardado.Tipo], scope)
					} else {
						//Copiar los valores del vector guardado al nuevo vector entrante
						CopiarVector(&vectorGuardado, &vectorEntrante, simbolo_id)
						valor = Ast.TipoRetornado{
							Tipo:  Ast.VECTOR,
							Valor: vectorEntrante,
						}
					}
				} else if vectorGuardado.Tipo == Ast.VECTOR {
					//Verificar que los tipos de vectores que se guardan son correctos
					if vectorGuardado.Tipo != fn_vectores.GetTipoVector(vectorEntrante) ||
						!fn_vectores.GetNivelesVector(vectorGuardado, vectorEntrante) {
						//Error no se pueden guardar 2 tipos de vectores diferentes
						/////////////////////////////ERROR/////////////////////////////////////
						return errores.GenerarError(25, a, a, id, Ast.ValorTipoDato[fn_vectores.GetTipoVector(vectorEntrante)],
							Ast.ValorTipoDato[vectorGuardado.Tipo], scope)

					} else {
						CopiarVector(&vectorGuardado, &vectorEntrante, simbolo_id)
						valor = Ast.TipoRetornado{
							Tipo:  Ast.VECTOR,
							Valor: vectorEntrante,
						}
					}
				} else {
					CopiarVector(&vectorGuardado, &vectorEntrante, simbolo_id)
					valor = Ast.TipoRetornado{
						Tipo:  Ast.VECTOR,
						Valor: vectorEntrante,
					}
				}
			}

			//Revisar si es array y si es un array del mismo tipo
			if valor.Tipo == Ast.ARRAY {
				arrayEntrante := valor.Valor.(expresiones.Array)
				arrayGuardado := simbolo_id.Valor.(Ast.TipoRetornado).Valor.(expresiones.Array)
				if arrayEntrante.TipoArray != arrayGuardado.TipoArray {
					//Hay varias opciones y una es que la lista que entra es indefinida
					//Y la otra es que si traiga un tipo diferente
					if arrayEntrante.TipoArray != Ast.INDEFINIDO {
						//Generar el Error, de lo contrario todo bien
						/////////////////////////////ERROR/////////////////////////////////////
						return errores.GenerarError(26, a, a, id, Ast.ValorTipoDato[arrayEntrante.TipoArray],
							Ast.ValorTipoDato[arrayGuardado.TipoArray], scope)
					} else {
						//Copiar los valores del vector guardado al nuevo vector entrante
						CopiarArray(arrayGuardado, arrayEntrante, simbolo_id)
						valor = Ast.TipoRetornado{
							Tipo:  Ast.ARRAY,
							Valor: arrayEntrante,
						}
					}
				} else {
					//Comparar las dimensiones del array
					verificarArray := CompararArrays(arrayGuardado, arrayEntrante, scope)
					if !verificarArray {
						/////////////////////////////ERROR/////////////////////////////////////
						return errores.GenerarError(25, a, a, id, "", "", scope)
					}

					CopiarArray(arrayGuardado, arrayEntrante, simbolo_id)
					valor = Ast.TipoRetornado{
						Tipo:  Ast.ARRAY,
						Valor: arrayEntrante,
					}
				}
			}

			/*Trabajar el todo el cod3d desde aqui*/
			codigo3d += obj3d.Codigo
			codigo3d += obj3dValor.Codigo

			/*Si es string verificar el tamaño del entrante y del que esta guardado*/
			/*
				switch simbolo_id.Tipo {
				case Ast.STRING, Ast.STR:
					//Comparar los len de los string
					esMenor := false
					stringGuardado := simbolo_id.Valor.(Ast.TipoRetornado).Valor.(string)
					stringValor := valor.Valor.(string)
					if len(stringGuardado) < len(stringValor) {
						esMenor = true
					}
					//Get posicion, codigo3d y la nueva referencia
					cod, ref := GetCod3DString(obj3d.Referencia, obj3dValor.Referencia, esMenor)
					//Actualizar referencia, codigo y posicion
					referencia = ref
					obj3d.Referencia = referencia
					codigo3d += cod
				case Ast.ARRAY:
				case Ast.VECTOR:

				case Ast.STRUCT:

				}
			*/

			if simbolo_id.TipoDireccion == Ast.STACK {
				//Se va a guardar en el stack
				referencia = obj3d.Referencia
				//codigo3d += temp + " = P + " + strconv.Itoa(simbolo_id.Direccion) + ";\n"
				codigo3d += "/*********************ASIGNACION DE NUEVO VALOR*/\n"
				codigo3d += "stack[(int)" + referencia + "] = " + obj3dValor.Referencia + ";\n"
				codigo3d += "/***********************************************/\n"
			} else {
				//Se va a guardar en el heap
				referencia = obj3d.Referencia
				//codigo3d += temp + " =" + strconv.Itoa(simbolo_id.Direccion) + ";\n"
				codigo3d += "/*********************ASIGNACION DE NUEVO VALOR*/\n"
				codigo3d += "heap[(int)" + referencia + "] = " + obj3dValor.Referencia + ";\n"
				codigo3d += "/***********************************************/\n"
			}

			//Copiar valor
			nuevoValor := valor
			simbolo_id.Valor = nuevoValor
			scope.UpdateSimbolo(id, simbolo_id)
		} else {
			//Revisar si el retorno es un error
			if valor.Tipo == Ast.ERROR {
				return valor
			}
			//Error de tipos, generar un error semántico
			/////////////////////////////ERROR/////////////////////////////////////
			return errores.GenerarError(24, a, a, id, Ast.ValorTipoDato[int(valor.Tipo)],
				Ast.ValorTipoDato[int(simbolo_id.Valor.(Ast.TipoRetornado).Tipo)], scope)
		}
	} else {
		//No existe, generar un error semántico
		/////////////////////////////ERROR/////////////////////////////////////
		return errores.GenerarError(23, a, a, id, "", "", scope)
	}
	obj3d.Codigo = codigo3d
	obj3d.Valor = valor
	return Ast.TipoRetornado{
		Tipo:  Ast.ASIGNACION,
		Valor: obj3d,
	}
}

func (a Asignacion) AsignarAccesoVector(id string, scope *Ast.Scope) Ast.TipoRetornado {
	//Conseguir la posición en donde se quiere agregar el nuevo elemento
	/*************VARIABLES 3D*************/
	var obj3dPosicion Ast.O3D
	var idExp expresiones.Identificador
	var obj3d, obj3dValor Ast.O3D
	var referencia, codigo3d string
	/**************************************/
	agregarElemento := false
	posicion := a.Id.(fn_vectores.AccesoVec).Posicion.(Ast.Expresion).GetValue(scope)
	obj3dPosicion = posicion.Valor.(Ast.O3D)
	posicion = obj3dPosicion.Valor

	//_, tipoParticular := a.Id.(fn_vectores.AccesoVec).Posicion.(Ast.Abstracto).GetTipo()
	//Verificar que sea usize
	if posicion.Tipo != Ast.USIZE && posicion.Tipo != Ast.I64 { //||
		//tipoParticular == Ast.IDENTIFICADOR && posicion.Tipo == Ast.I64 {
		//Error, se espera un usize
		fila := a.Valor.(Ast.Abstracto).GetFila()
		columna := a.Valor.(Ast.Abstracto).GetColumna()
		msg := "Semantic error, expected USIZE, found. " + Ast.ValorTipoDato[posicion.Tipo] +
			". -- Line: " + strconv.Itoa(fila) +
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
	//Conseguir el int de la posicion
	posicionNum := posicion.Valor.(int)
	//Verificar que el id  exista
	existe := scope.Exist(id)
	//Obtener el valor del id
	simbolo_id := scope.GetSimbolo(id)

	/*********C3D del Simbolo***********/
	codigo3d += "/********************************ACCESO A VECTOR*/\n"
	idExp = expresiones.NewIdentificador(id, Ast.IDENTIFICADOR, 0, 0)
	obj3d = idExp.GetValue(scope).Valor.(Ast.O3D)
	referencia = obj3d.Referencia
	codigo3d += obj3d.Codigo
	/***********************************/

	//Verificar que el simbolo sea un vector
	if simbolo_id.Tipo != Ast.VECTOR {
		//Error, se espera un identificador. un acceso a vector o un acceso a un array
		msg := "Semantic error, expected VECTOR, found " + Ast.ValorTipoDato[simbolo_id.Tipo] +
			". -- Line: " + strconv.Itoa(a.Valor.(Ast.Abstracto).GetFila()) +
			" Column: " + strconv.Itoa(a.Valor.(Ast.Abstracto).GetColumna())
		nError := errores.NewError(a.Valor.(Ast.Abstracto).GetFila(), a.Valor.(Ast.Abstracto).GetColumna(), msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}
	//Conseguir el vector
	vector := simbolo_id.Valor.(Ast.TipoRetornado).Valor.(expresiones.Vector)

	//Verificar que los tipos sean correctos
	//Primero verificar que no es un if expresion
	_, tipoIn := a.Valor.(Ast.Abstracto).GetTipo()
	var preValor interface{}
	if tipoIn == Ast.IF_EXPRESION || tipoIn == Ast.MATCH_EXPRESION || tipoIn == Ast.LOOP_EXPRESION {
		preValor = a.Valor.(Ast.Instruccion).Run(scope)
	} else {
		preValor = a.Valor.(Ast.Expresion).GetValue(scope)
	}

	valor := preValor.(Ast.TipoRetornado)
	obj3dValor = valor.Valor.(Ast.O3D)
	valor = obj3dValor.Valor

	if existe {
		//Primero verificar si es mutable
		if !simbolo_id.Mutable {
			//No es mutable, error semántico
			msg := "Semantic error, can't modify a non-mutable " + Ast.ValorTipoDato[int(simbolo_id.Tipo)] +
				" type. -- Line: " + strconv.Itoa(a.Fila) +
				" Column: " + strconv.Itoa(a.Columna)
			nError := errores.NewError(a.Fila, a.Columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			nError.Ambito = scope.GetTipoScope()
			scope.Errores.Add(nError)
			scope.Consola += msg + "\n"
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}
		}
		//Primero verificar
		//Existe, ahora verificar los tipos
		if vector.TipoVector.Tipo == valor.Tipo {
			//Verificar que no sea un struct
			if valor.Tipo == Ast.STRUCT {
				tipoStruct := valor.Valor.(Ast.Structs).GetPlantilla(scope)
				if tipoStruct != expresiones.GetTipoFinal(vector.TipoVector).Valor {
					fila := valor.Valor.(Ast.Abstracto).GetFila()
					columna := valor.Valor.(Ast.Abstracto).GetColumna()
					msg := "Semantic error, can't store \"" + tipoStruct +
						"\" to a VECTOR<" + expresiones.Tipo_String(vector.TipoVector) + ">" +
						". -- Line: " + strconv.Itoa(fila) +
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
				} else {
					agregarElemento = true
				}

			} else
			//Revisar si es vector y si es del tipo de vector correcto
			if valor.Tipo == Ast.VECTOR {
				vectorEntrante := valor.Valor.(expresiones.Vector)
				vectorGuardado := vector
				if vectorEntrante.TipoVector != vectorGuardado.TipoVector {
					//Hay varias opciones y una es que la lista que entra es indefinida
					//Y la otra es que si trae un tipo diferente
					if vectorEntrante.Tipo != Ast.INDEFINIDO {
						//Generar el Error, de lo contrario todo bien
						msg := "Semantic error, can't assign Vector<" + Ast.ValorTipoDato[vectorEntrante.Tipo] + ">" +
							" to Vector<" + Ast.ValorTipoDato[vectorGuardado.Tipo] + ">" +
							" type. -- Line: " + strconv.Itoa(a.Fila) +
							" Column: " + strconv.Itoa(a.Columna)
						nError := errores.NewError(a.Fila, a.Columna, msg)
						nError.Tipo = Ast.ERROR_SEMANTICO
						nError.Ambito = scope.GetTipoScope()
						scope.Errores.Add(nError)
						scope.Consola += msg + "\n"
						return Ast.TipoRetornado{
							Tipo:  Ast.ERROR,
							Valor: nError,
						}
					} else {
						//Copiar los valores del vector guardado al nuevo vector entrante
						agregarElemento = true
					}
				} else if vectorGuardado.Tipo == Ast.VECTOR {
					//Verificar que los tipos de vectores que se guardan son correctos
					if vectorGuardado.Tipo != fn_vectores.GetTipoVector(vectorEntrante) ||
						!fn_vectores.GetNivelesVector(vectorGuardado, vectorEntrante) {
						//Error no se pueden guardar 2 tipos de vectores diferentes
						fila := vectorEntrante.GetFila()
						columna := vectorEntrante.GetColumna()
						msg := "Semantic error, can't store VECTOR<" + Ast.ValorTipoDato[fn_vectores.GetTipoVector(vectorEntrante)] +
							"> to a VECTOR<" + Ast.ValorTipoDato[vectorGuardado.Tipo] + ">" +
							". -- Line: " + strconv.Itoa(fila) +
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
					} else {
						agregarElemento = true
					}
				} else {
					agregarElemento = true
				}
			} else {
				agregarElemento = true
			}

			if agregarElemento {
				//Copiar valor
				nuevoValor := a.Valor
				expresiones.UpdatePosition(&vector, posicionNum, nuevoValor, scope)
				simbolo_id.Valor = Ast.TipoRetornado{
					Tipo:  Ast.VECTOR,
					Valor: vector,
				}
				scope.UpdateSimbolo(id, simbolo_id)

				/*********************CODIGO 3D*****************************/
				codigo3d += GetCod3DAsignacionVector(referencia, obj3dValor.Referencia, posicionNum)
				/***********************************************************/
			}
		} else {
			//No existe, generar un error semántico
			fila := a.Valor.(Ast.Abstracto).GetFila()
			columna := a.Valor.(Ast.Abstracto).GetColumna()
			msg := "Semantic error, can't store " + Ast.ValorTipoDato[valor.Tipo] +
				" to a VECTOR<" + Ast.ValorTipoDato[vector.Tipo] + ">" +
				". -- Line: " + strconv.Itoa(fila) +
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
		return Ast.TipoRetornado{
			Tipo:  Ast.EJECUTADO,
			Valor: true,
		}
	}

	obj3d.Codigo = codigo3d
	obj3d.Valor = Ast.TipoRetornado{
		Tipo:  Ast.EJECUTADO,
		Valor: true,
	}

	return Ast.TipoRetornado{
		Tipo:  Ast.ASIGNACION,
		Valor: obj3d,
	}
}

func CopiarVector(vectorGuardado *expresiones.Vector, vectorEntrante *expresiones.Vector, simbolo Ast.Simbolo) {
	vectorEntrante.Columna = simbolo.Columna
	vectorEntrante.Fila = simbolo.Fila
	vectorEntrante.Mutable = simbolo.Mutable
	vectorEntrante.Tipo = vectorGuardado.Tipo
	vectorEntrante.TipoVector = vectorGuardado.TipoVector
}

func CopiarArray(arrayGuardado expresiones.Array, arrayEntrante expresiones.Array, simbolo Ast.Simbolo) {
	arrayEntrante.Columna = simbolo.Columna
	arrayEntrante.Fila = simbolo.Fila
	arrayEntrante.Mutable = simbolo.Mutable
	arrayEntrante.Tipo = arrayGuardado.Tipo
	arrayEntrante.TipoArray = arrayGuardado.TipoArray
}

func GenerarC3DGetSimbolo(simbolo Ast.Simbolo) Ast.O3D {
	/*Variables para C3D*/
	var obj3d Ast.O3D
	var temp string = Ast.GetTemp()
	//var tempValor string = Ast.GetTemp()
	var codigo3d string = ""
	/*Verificar si la variable viene del heap o del stack*/

	if simbolo.TipoDireccion == Ast.STACK {
		codigo3d = "/***********GET VALOR VARIABLE CON IDENTIFICADOR*/\n"

		codigo3d += temp + " = P + " + strconv.Itoa(simbolo.Direccion) + ";\n"
		//codigo3d += tempValor + " = stack[(int)" + temp + "];\n"
		codigo3d += "/***********************************************/\n"
	} else {
		codigo3d = "/***********GET VALOR VARIABLE CON IDENTIFICADOR*/\n"
		codigo3d += temp + " = " + strconv.Itoa(simbolo.Direccion) + ";\n"
		//codigo3d += tempValor + " = heap[(int)" + temp + "];\n"
		codigo3d += "/***********************************************/\n"
	}

	/*Inicializar el obj3d*/
	obj3d.Referencia = temp
	obj3d.Valor = simbolo.Valor.(Ast.TipoRetornado)
	obj3d.Codigo = codigo3d

	return obj3d
}

func GetCod3DString(posSimbolo, posValor string, nuevaPos bool) (string, string) {
	var temp, temp1, temp2, temp3, temp4, temp5 = "", "", "", "", "", ""
	var lt = ""
	var lf = ""
	var salto = ""
	codigo3d := "/*****************************ASIGNACIÓN STRING*/\n"
	temp = Ast.GetTemp()
	temp1 = Ast.GetTemp()
	temp2 = Ast.GetTemp()
	temp3 = Ast.GetTemp()
	temp4 = Ast.GetTemp()
	lt = Ast.GetLabel()
	lf = Ast.GetLabel()
	salto = Ast.GetLabel()

	if nuevaPos {
		temp5 = Ast.GetTemp()
		codigo3d += temp5 + " = H;\n"
		codigo3d += "H = H + 1;\n"
		Ast.GetH()
		posSimbolo = temp5
	}
	codigo3d += temp + " = " + posSimbolo + "; //Guardar la referencia \n"
	codigo3d += salto + ":\n"
	codigo3d += temp1 + " = " + "heap[(int)" + posSimbolo + "];\n" //Valor string
	codigo3d += temp2 + " = " + "heap[(int)" + posValor + "];\n"   //Valor string
	codigo3d += "if (" + temp2 + " != 0) goto " + lt + ";\n"
	codigo3d += "goto " + lf + ";\n"
	codigo3d += lt + ":\n"
	codigo3d += "heap[(int)" + posSimbolo + "] = " + temp2 + "; //Asignar letra\n"
	codigo3d += temp3 + " = " + posSimbolo + " + 1; //Next pos simbolo\n"
	codigo3d += posSimbolo + " = " + temp3 + ";\n"
	codigo3d += temp4 + " = " + posValor + " + 1; //Next pos valor\n"
	codigo3d += posValor + " = " + temp4 + ";\n"
	codigo3d += "goto " + salto + ";\n"
	codigo3d += lf + ":\n"
	codigo3d += "heap[(int)" + posSimbolo + "] = 0;\n"
	codigo3d += posSimbolo + " = " + temp + "; //Recuperar el valor del inicio de la cadena\n"
	codigo3d += "/***********************************************/\n"
	return codigo3d, posSimbolo
}

func GetCod3DAsignacionVector(refVector, refElemento string, posicion int) string {
	codigo3d := ""
	vector := Ast.GetTemp()
	sizeVec := Ast.GetTemp()
	primeraPos := Ast.GetTemp()
	posVector := Ast.GetTemp()
	lt := Ast.GetLabel()
	lf := Ast.GetLabel()
	salto := Ast.GetLabel()
	codigo3d += vector + " = heap[(int)" + refVector + "]; //Get el vector\n"
	codigo3d += sizeVec + " = heap[(int)" + vector + "] //Size vector\n"
	codigo3d += "if (" + strconv.Itoa(posicion) + " < " + sizeVec + " goto " + lt + ";\n"
	codigo3d += "goto " + lf + ";\n"
	codigo3d += lt + ":\n"
	codigo3d += primeraPos + " = " + vector + " + 1; //Get primera pos de vector \n "
	codigo3d += posVector + " = " + primeraPos + " + " + strconv.Itoa(posicion) + "; //pos para asignar\n "
	codigo3d += "heap[(int)" + posVector + "] = " + refElemento + ";\n"
	codigo3d += "goto " + salto + ";\n"
	codigo3d += fn_vectores.BoundsError(lf)
	return codigo3d
}
