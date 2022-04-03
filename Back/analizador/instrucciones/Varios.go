package instrucciones

import (
	"Back/analizador/Ast"
	"Back/analizador/expresiones"
	"strconv"
)

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

func ValorPorDefecto(tipo Ast.TipoDato, scope *Ast.Scope) Ast.O3D {
	var retorno Ast.TipoRetornado
	var direccion int = scope.Size
	var obj3D Ast.O3D
	var temp = Ast.GetTemp()
	var temp2 string
	var codigo string
	switch tipo {
	case Ast.I64:
		codigo += "/*************************DECLARACION DEFAULT*/\n"
		codigo += temp + " = P + " + strconv.Itoa(direccion) + ";\n"
		codigo += "stack[(int)" + temp + "] = 0;\n"
		codigo += "/********************************************/\n"
		scope.Size++
		obj3D.Codigo = codigo
		obj3D.Referencia = temp
		retorno.Tipo = Ast.I64
		retorno.Valor = 0
	case Ast.F64:
		codigo += "/*************************DECLARACION DEFAULT*/\n"
		codigo += temp + " = P + " + strconv.Itoa(direccion) + ";\n"
		codigo += "stack[(int)" + temp + "] = 0.0;\n"
		codigo += "/********************************************/\n"
		scope.Size++
		obj3D.Codigo = codigo
		obj3D.Referencia = temp
		retorno.Tipo = Ast.F64
		retorno.Valor = float64(0)
	case Ast.CHAR:
		codigo += "/*************************DECLARACION DEFAULT*/\n"
		codigo += temp + " = P + " + strconv.Itoa(direccion) + ";\n"
		codigo += "stack[(int)" + temp + "] = 0;\n"
		codigo += "/********************************************/\n"
		scope.Size++
		obj3D.Codigo = codigo
		obj3D.Referencia = temp
		retorno.Tipo = Ast.CHAR
		retorno.Valor = ""
	case Ast.STRING, Ast.STR:
		temp2 = Ast.GetTemp()
		codigo += "/************AGREGANDO UN STRING/STR AL HEAP*/\n"
		codigo += temp + " = " + "H;\n"
		codigo += "heap[(int)H] = 0;\n"
		codigo += "H = H + 1;\n"
		codigo += "/********************************************/\n"
		codigo += "/*************************DECLARACION DEFAULT*/\n"
		codigo += temp2 + " = P + " + strconv.Itoa(direccion) + ";\n"
		codigo += "stack[(int)" + temp2 + "] = " + temp + ";\n"
		codigo += "/********************************************/\n"
		scope.Size++
		obj3D.Codigo = codigo
		obj3D.Referencia = temp2
		retorno.Tipo = tipo
		retorno.Valor = ""
	case Ast.BOOLEAN:
		codigo += "/*************************DECLARACION DEFAULT*/\n"
		codigo += temp + " = P + " + strconv.Itoa(direccion) + ";\n"
		codigo += "stack[(int)" + temp + "] = 0;\n"
		codigo += "/********************************************/\n"
		scope.Size++
		obj3D.Codigo = codigo
		obj3D.Referencia = temp
		retorno.Tipo = Ast.BOOLEAN
		retorno.Valor = false
	case Ast.USIZE:
		codigo += "/************************DECLARACION DEFAULT*/\n"
		codigo += temp + " = P + " + strconv.Itoa(direccion) + ";\n"
		codigo += "stack[(int)" + temp + "] = 0;\n"
		codigo += "/********************************************/\n"
		scope.Size++
		obj3D.Codigo = codigo
		obj3D.Referencia = temp
		retorno.Tipo = Ast.USIZE
		retorno.Valor = 0
	}
	obj3D.Valor = retorno
	return obj3D
}
