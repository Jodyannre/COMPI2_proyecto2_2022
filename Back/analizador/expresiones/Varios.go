package expresiones

import (
	"Back/analizador/Ast"

	"github.com/colegno/arraylist"
)

func UpdatePosition(v *Vector, posicion int, valorEntrante interface{}, scope *Ast.Scope) Ast.TipoRetornado {
	valor := valorEntrante.(Ast.Expresion).GetValue(scope)
	listaNueva := *arraylist.New()
	if valor.Tipo == Ast.ERROR {
		return valor
	}

	for i := 0; i < v.Valor.Len(); i++ {
		if i == posicion {
			listaNueva.Add(valor)
			continue
		}
		listaNueva.Add(v.Valor.GetValue(i))
	}
	v.Valor = &listaNueva
	return Ast.TipoRetornado{
		Tipo:  Ast.EJECUTADO,
		Valor: true,
	}
}

func GetTipoFinal(tipo Ast.TipoRetornado) Ast.TipoRetornado {
	if EsTipoFinal(tipo.Tipo) {
		if tipo.Tipo != Ast.STRUCT {
			return Ast.TipoRetornado{
				Tipo:  tipo.Tipo,
				Valor: true,
			}
		}
		return Ast.TipoRetornado{
			Tipo:  tipo.Tipo,
			Valor: tipo.Valor,
		}
	} else {
		return GetTipoFinal(tipo.Valor.(Ast.TipoRetornado))
	}
}

func EsTipoFinal(tipo Ast.TipoDato) bool {
	switch tipo {
	case Ast.I64, Ast.F64, Ast.CHAR, Ast.STRING, Ast.STR, Ast.USIZE, Ast.BOOLEAN, Ast.STRUCT, Ast.INDEFINIDO,
		Ast.DIMENSION_ARRAY:
		return true
	default:
		return false
	}
}

func CompararTipos(tipoA Ast.TipoRetornado, tipoB Ast.TipoRetornado) bool {
	if EsTipoFinal(tipoA.Tipo) && EsTipoFinal(tipoB.Tipo) {
		if tipoA.Tipo == tipoB.Tipo {
			//Verificar si son structs
			if tipoA.Tipo == Ast.STRUCT {
				if tipoA.Valor == tipoB.Valor {
					return true
				} else {
					return false
				}
			}
			return true
		} else {
			return false
		}
	}
	if EsTipoFinal(tipoA.Tipo) && !EsTipoFinal(tipoB.Tipo) ||
		!EsTipoFinal(tipoA.Tipo) && EsTipoFinal(tipoB.Tipo) {
		return false
	}
	return CompararTipos(tipoA.Valor.(Ast.TipoRetornado), tipoB.Valor.(Ast.TipoRetornado))
}

func Tipo_String(t Ast.TipoRetornado) string {
	if t.Tipo == Ast.VECTOR {
		return "Vec <" + Tipo_String(t.Valor.(Ast.TipoRetornado)) + ">"
	} else {
		if t.Tipo == Ast.STRUCT {
			return t.Valor.(string)
		} else {
			return Ast.ValorTipoDato[t.Tipo]
		}
	}
}

func EsVector(tipo Ast.TipoDato) Ast.TipoDato {
	switch tipo {
	case Ast.VEC_ELEMENTOS, Ast.VEC_FAC, Ast.VEC_WITH_CAPACITY, Ast.VEC_NEW:
		return Ast.VECTOR
	default:
		return tipo
	}
}

func GetTipoEstructura(tipo Ast.TipoRetornado, scope *Ast.Scope, elemento interface{}) Ast.TipoRetornado {

	if tipo.Tipo == Ast.VECTOR {
		return Ast.TipoRetornado{
			Tipo:  Ast.VECTOR,
			Valor: GetTipoEstructura(tipo.Valor.(Ast.TipoRetornado), scope, elemento),
		}
	}
	if tipo.Tipo == Ast.ARRAY {
		return Ast.TipoRetornado{
			Tipo:  Ast.ARRAY,
			Valor: GetTipoEstructura(tipo.Valor.(Ast.TipoRetornado), scope, elemento),
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
		//tipoNuevo := simboloStruct.Valor.(Ast.Simbolo).Identificador
		return Ast.TipoRetornado{
			Valor: simboloStruct.Valor,
			Tipo:  Ast.STRUCT,
		}
	}
	if tipo.Tipo == Ast.STRUCT {
		//Verificar que el struct exista
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
	if EsTipoFinal(tipo.Tipo) {
		return Ast.TipoRetornado{
			Tipo:  Ast.BOOLEAN,
			Valor: true,
		}
	}
	return ErrorEnTipo(tipo.Valor.(Ast.TipoRetornado))
}

func EsVAS(tipo Ast.TipoDato) bool {
	switch tipo {
	case Ast.VECTOR, Ast.ARRAY, Ast.STRUCT:
		return true
	default:
		return false

	}
}
