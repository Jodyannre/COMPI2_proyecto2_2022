package instrucciones

import (
	"Back/analizador/Ast"
	"Back/analizador/expresiones"
)

func EsTipoEspecial(tipo Ast.TipoDato) bool {
	switch tipo {
	case Ast.VECTOR, Ast.ARRAY, Ast.STRUCT:
		return true
	default:
		return false
	}
}

func GetTipoEspecial(tipo Ast.TipoDato, elemento interface{}, scope *Ast.Scope) Ast.TipoRetornado {
	switch tipo {
	case Ast.VECTOR:
		//Get el tipo en estructura
		tipoN := GetTipoEstructura(elemento.(expresiones.Vector).TipoVector, scope)
		errores := ErrorEnTipo(tipoN)
		if errores.Tipo == Ast.ERROR {
			return errores
		} else {
			return Ast.TipoRetornado{
				Tipo:  Ast.VECTOR,
				Valor: tipoN,
			}
		}
	case Ast.ARRAY:
		tipoN := GetTipoEstructura(elemento.(expresiones.Array).TipoDelArray, scope)
		errores := ErrorEnTipo(tipoN)
		if errores.Tipo == Ast.ERROR {
			return errores
		} else {
			return Ast.TipoRetornado{
				Tipo:  Ast.ARRAY,
				Valor: tipoN,
			}
		}
	case Ast.STRUCT:
		plantilla := elemento.(Ast.Structs).GetPlantilla(scope)
		return Ast.TipoRetornado{
			Tipo:  Ast.STRUCT,
			Valor: plantilla,
		}
	default:
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: true,
		}

	}
}

func GetTipoEstructura(tipo Ast.TipoRetornado, scope *Ast.Scope) Ast.TipoRetornado {

	if tipo.Tipo == Ast.VECTOR {
		return Ast.TipoRetornado{
			Tipo:  Ast.VECTOR,
			Valor: GetTipoEstructura(tipo.Valor.(Ast.TipoRetornado), scope),
		}
	}
	if tipo.Tipo == Ast.ARRAY {
		return Ast.TipoRetornado{
			Tipo:  Ast.ARRAY,
			Valor: GetTipoEstructura(tipo.Valor.(Ast.TipoRetornado), scope),
		}
	}
	if tipo.Tipo == Ast.ACCESO_MODULO {
		//Verificar que se pueda acceder o que exista
		simboloStruct := tipo.Valor.(Ast.AccesosM).GetTipoFromAccesoModulo(tipo, scope)
		if simboloStruct.Tipo == Ast.ERROR {
			return simboloStruct
		}
		if simboloStruct.Tipo != Ast.STRUCT_TEMPLATE {
			//No es un struct
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: true,
			}
		}
		//De lo contrario devolvio el simbolo
		tipoNuevo := simboloStruct.Valor
		return Ast.TipoRetornado{
			Valor: tipoNuevo,
			Tipo:  Ast.STRUCT,
		}
	}
	if tipo.Tipo == Ast.STRUCT {
		//Verificar que el struct exista
		return tipo
		/*
			nombreStruct := tipo.Valor.(string)
			if scope.Exist(nombreStruct) {
				return tipo
			} else {
				//No existe el struct
				return Ast.TipoRetornado{
					Tipo:  Ast.ERROR,
					Valor: true,
				}
			}
		*/

	}
	return Ast.TipoRetornado{
		Tipo:  tipo.Tipo,
		Valor: true,
	}
}

func ErrorEnTipo(tipo Ast.TipoRetornado) Ast.TipoRetornado {
	if tipo.Tipo == Ast.ERROR {
		return tipo
	}
	if expresiones.EsTipoFinal(tipo.Tipo) {
		return Ast.TipoRetornado{
			Tipo:  Ast.BOOLEAN,
			Valor: true,
		}
	}
	return ErrorEnTipo(tipo.Valor.(Ast.TipoRetornado))
}

func EsAVelementos(tipo Ast.TipoDato) bool {
	switch tipo {
	case Ast.ARRAY_ELEMENTOS, Ast.ARRAY_FAC, Ast.VEC_ELEMENTOS,
		Ast.VEC_FAC:
		return true
	default:
		return false
	}
}
