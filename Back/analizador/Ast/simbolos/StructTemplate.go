package simbolos

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"strconv"

	"github.com/colegno/arraylist"
)

type StructTemplate struct {
	Tipo        Ast.TipoDato
	Nombre      string
	Atributos   map[string]*Atributo
	AtributosIn *arraylist.List
	Publico     bool
	Fila        int
	Columna     int
}

func NewStructTemplate(nombre string, atributos *arraylist.List, publico bool, fila, columna int) StructTemplate {
	att := make(map[string]*Atributo)
	//Agregar los elementos al nuevo struct template
	nuevo := StructTemplate{
		Tipo:        Ast.STRUCT_TEMPLATE,
		Nombre:      nombre,
		Atributos:   att,
		Publico:     publico,
		Fila:        fila,
		Columna:     columna,
		AtributosIn: atributos,
	}
	return nuevo
}

func (s StructTemplate) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	sinAtributos := false
	var resultadoFormatoTipo Ast.TipoRetornado
	if s.AtributosIn.Len() == 0 {
		//No tiene atributos, pero no es error
		sinAtributos = true
	}
	if !sinAtributos {
		for i := 0; i < s.AtributosIn.Len(); i++ {
			att_val := s.AtributosIn.GetValue(i).(*Atributo)
			//atributo := att_val.GetValue(scope)
			for key, _ := range s.Atributos {
				if key == att_val.Nombre {
					msg := "Semantic error, field already declared." +
						" type. -- Line: " + strconv.Itoa(att_val.Fila) +
						" Column: " + strconv.Itoa(att_val.Columna)
					nError := errores.NewError(att_val.Fila, att_val.Columna, msg)
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
			resultadoFormatoTipo = att_val.FormatearTipo(scope)
			if resultadoFormatoTipo.Tipo == Ast.ERROR {
				return resultadoFormatoTipo
			}
			s.Atributos[att_val.Nombre] = att_val
		}
	}

	return Ast.TipoRetornado{
		Tipo:  Ast.STRUCT,
		Valor: s,
	}
}
