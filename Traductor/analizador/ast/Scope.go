package ast

import (
	"strings"

	"github.com/colegno/arraylist"
)

type Scope struct {
	Nombre               string
	prev                 *Scope
	tablaSimbolos        map[string]interface{}
	tablaModulos         map[string]interface{}
	tablaFunciones       map[string]interface{}
	tablaStructs         map[string]interface{}
	tablaSimbolosReporte *arraylist.List
	Errores              *arraylist.List
	Consola              string
	Global               bool
}

func NewScope(name string, prev *Scope) Scope {

	nuevo := Scope{Nombre: name, prev: prev}
	nuevo.Errores = arraylist.New()
	nuevo.tablaSimbolos = make(map[string]interface{})
	nuevo.tablaModulos = make(map[string]interface{})
	nuevo.tablaFunciones = make(map[string]interface{})
	nuevo.tablaStructs = make(map[string]interface{})
	nuevo.tablaSimbolosReporte = arraylist.New()
	nuevo.Global = false
	return nuevo
}

func (scope *Scope) Add(simbolo Simbolo) {
	var global *Scope = scope
	id := strings.ToUpper(simbolo.Identificador)

	//Agregar el símbolo al scope actual
	scope.tablaSimbolos[id] = simbolo

	//Recuperar el scope global
	for scope_actual := scope; scope_actual.prev != nil; scope_actual = scope_actual.prev {
		global = scope_actual
	}
	//Crear el símbolo para la tabla de reporte de símbolos
	simboloReporte := simbolo.NewSimboloReporte(scope)
	global.tablaSimbolosReporte.Add(simboloReporte)
}

func (scope *Scope) Exist(ident string) bool {
	id := strings.ToUpper(ident)
	for scope_actual := scope; scope_actual != nil; scope_actual = scope_actual.prev {
		for key, _ := range scope_actual.tablaSimbolos {
			if key == id {
				return true
			}
		}
	}
	return false
}

func (scope *Scope) Exist_actual(ident string) bool {
	id := strings.ToUpper(ident)
	for key, _ := range scope.tablaSimbolos {
		if key == id {
			return true
		}
	}
	return false
}

func (scope *Scope) UpdateSimbolo(ident string, valorNuevo Simbolo) {
	id := strings.ToUpper(ident)
	for scope_actual := scope; scope_actual != nil; scope_actual = scope_actual.prev {

		for key, _ := range scope_actual.tablaSimbolos {
			if key == id {
				scope_actual.tablaSimbolos[key] = valorNuevo
				return
			}
		}
	}

}

func (scope *Scope) GetSimbolo(ident string) Simbolo {
	id := strings.ToUpper(ident)

	for scope_actual := scope; scope_actual != nil; scope_actual = scope_actual.prev {
		for key, simboloRetorno := range scope_actual.tablaSimbolos {
			if key == id {
				nsimbolo := simboloRetorno.(Simbolo)
				return nsimbolo
			}
		}
	}
	var simboloNull Simbolo
	simboloNull.Tipo = NULL
	return simboloNull
}

func (scope *Scope) GetSimboloReferencia(ident string) Simbolo {
	id := strings.ToUpper(ident)

	for scope_actual := scope; scope_actual != nil; scope_actual = scope_actual.prev {
		for key, simboloRetorno := range scope_actual.tablaSimbolos {
			if key == id && scope_actual != scope {
				nsimbolo := simboloRetorno.(Simbolo)
				return nsimbolo
			}
		}
	}
	var simboloNull Simbolo
	return simboloNull
}

func (scope *Scope) GetTipoScope() string {
	if strings.ToUpper(scope.Nombre) == "GLOBAL" {
		return "Global"
	} else {
		return "Local"
	}

}

func (s *Scope) UpdateScopeGlobal() {
	//Primero actualizar todas los valores por referencia

	//Obtener el scope global
	var scope_global *Scope
	if s.prev != nil {
		for scope_global = s; scope_global.prev != nil; scope_global = scope_global.prev {
			//Buscando el scope global
		}
	} else {
		scope_global = s
	}
	if s != scope_global {
		scope_global.Consola += s.Consola
		for i := 0; i < s.Errores.Len(); i++ {
			elemento := s.Errores.GetValue(i)
			scope_global.Errores.Add(elemento)
		}
		for i := 0; i < s.tablaSimbolosReporte.Len(); i++ {
			elemento := s.tablaSimbolosReporte.Clone().GetValue(i)
			scope_global.tablaSimbolosReporte.Add(elemento)
		}
	}

}

func (scope *Scope) GetTablaModulos() map[string]interface{} {
	return scope.tablaModulos
}
