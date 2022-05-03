package Ast

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
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
	Temporales           *arraylist.List
	Errores              *arraylist.List
	Consola              string
	Codigo               string
	Global               bool
	Posicion             int
	Size                 int
	Stack                bool
	ContadorDeclaracion  int
}

func (scope *Scope) GetTablaModulos() map[string]interface{} {
	return scope.tablaModulos
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
	nuevo.Size = 0
	nuevo.Codigo = ""
	nuevo.Consola = ""
	nuevo.Stack = true
	nuevo.ContadorDeclaracion = 0
	return nuevo
}

func (scope *Scope) Add(simbolo Simbolo) {
	var global *Scope = scope
	id := strings.ToUpper(simbolo.Identificador)

	//Agregar el símbolo al scope actual
	scope.tablaSimbolos[id] = simbolo

	//Recuperar el scope global
	for scope_actual := scope; scope_actual != nil; scope_actual = scope_actual.prev {
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

func (scope *Scope) Exist_fms(ident string) Simbolo {
	//Primero conseguir el scope global
	var scope_global *Scope
	var scope_encontrado *Scope
	var retorno TipoRetornado
	id := strings.ToUpper(ident)
	if scope.prev != nil {
		for scope_global = scope; scope_global.prev != nil; scope_global = scope_global.prev {
			//Buscando el scope global
		}
	} else {
		scope_global = scope
	}

	//Verificar que la fms exista y si puede ser accedida
	retorno = buscarMap(id, MODULO, scope_global, scope)
	if retorno.Tipo != NULL {
		scope_encontrado = retorno.Valor.(Simbolo).Entorno
		if scope_global != scope_encontrado {
			//Los scopes no son iguales, verificar que la función sea pública
			if scope_encontrado != scope {
				if !retorno.Valor.(Simbolo).Publico {
					return NewSimbolo("", nil, -1, -1, ERROR_ACCESO_PRIVADO, false, false)
				}
			}

		}
		return retorno.Valor.(Simbolo)
	}
	retorno = buscarMap(id, STRUCT, scope_global, scope)
	if retorno.Tipo != NULL {
		scope_encontrado = retorno.Valor.(Simbolo).Entorno
		if scope_global != scope_encontrado {
			//Los scopes no son iguales, verificar que el struct sea público
			if scope_encontrado != scope {
				if !retorno.Valor.(Simbolo).Publico {
					return NewSimbolo("", nil, -1, -1, ERROR_ACCESO_PRIVADO, false, false)
				}
			}
		}
		return retorno.Valor.(Simbolo)
	}
	retorno = buscarMap(id, FUNCION, scope_global, scope)
	if retorno.Tipo != NULL {
		scope_encontrado = retorno.Valor.(Simbolo).Entorno
		if scope_global != scope_encontrado {
			//Los scopes no son iguales, verificar que el módulo sea pública
			if scope_encontrado != scope {
				if !retorno.Valor.(Simbolo).Publico {
					return NewSimbolo("", nil, -1, -1, ERROR_ACCESO_PRIVADO, false, false)
				}
			}
		}
		return retorno.Valor.(Simbolo)
	}
	return NewSimbolo("", nil, -1, -1, ERROR_NO_EXISTE, false, false)
}

func (scope *Scope) Exist_fms_declaracion(ident string) Simbolo {
	//Primero conseguir el scope global
	var scope_global *Scope
	var retorno TipoRetornado
	id := strings.ToUpper(ident)
	if scope.prev != nil {
		for scope_global = scope; scope_global.prev != nil; scope_global = scope_global.prev {
			//Buscando el scope global
		}
	} else {
		scope_global = scope
	}

	//Verificar que la fms exista y si puede ser accedida
	retorno = buscarMap(id, MODULO, scope_global, scope)
	if retorno.Tipo != NULL {

		return retorno.Valor.(Simbolo)
	}
	retorno = buscarMap(id, STRUCT, scope_global, scope)
	if retorno.Tipo != NULL {

		return retorno.Valor.(Simbolo)
	}
	retorno = buscarMap(id, FUNCION, scope_global, scope)
	if retorno.Tipo != NULL {
		return retorno.Valor.(Simbolo)
	}
	return NewSimbolo("", nil, -1, -1, ERROR_NO_EXISTE, false, false)
}

func (scope *Scope) Exist_fms_local(ident string) Simbolo {
	//Buscar entodos los scopes desde un local
	//Primero conseguir el scope global
	scope_Actual := scope
	scope_global := scope
	var retorno TipoRetornado
	var simboloRetorno Simbolo
	id := strings.ToUpper(ident)
	//Buscar el scope global
	for scope_Actual = scope; scope_Actual != nil; scope_Actual = scope_Actual.prev {
		scope_global = scope_Actual
	}

	for scope_Actual = scope; scope_Actual != nil; scope_Actual = scope_Actual.prev {

		//Buscar el módulo en los entornos locales
		//Verificar que la fms exista y si puede ser accedida
		retorno = buscarMap(id, MODULO, scope_Actual, scope)
		if retorno.Tipo != NULL {
			simboloRetorno = retorno.Valor.(Simbolo)
			if !simboloRetorno.Publico && simboloRetorno.Entorno != scope_global {
				return NewSimbolo("", nil, -1, -1, ERROR_ACCESO_PRIVADO, false, false)
			}
			return retorno.Valor.(Simbolo)
		}
		retorno = buscarMap(id, STRUCT, scope_Actual, scope)
		if retorno.Tipo != NULL {
			simboloRetorno = retorno.Valor.(Simbolo)
			if !simboloRetorno.Publico && simboloRetorno.Entorno != scope_global {
				return NewSimbolo("", nil, -1, -1, ERROR_ACCESO_PRIVADO, false, false)
			}
			return retorno.Valor.(Simbolo)
		}
		retorno = buscarMap(id, FUNCION, scope_Actual, scope)
		if retorno.Tipo != NULL {
			simboloRetorno = retorno.Valor.(Simbolo)
			if !simboloRetorno.Publico && simboloRetorno.Entorno != scope_global {
				return NewSimbolo("", nil, -1, -1, ERROR_ACCESO_PRIVADO, false, false)
			}
			return retorno.Valor.(Simbolo)
		}

	}
	return NewSimbolo("", nil, -1, -1, ERROR_NO_EXISTE, false, false)
}

func buscarMap(id string, tipo TipoDato, scope *Scope, local *Scope) TipoRetornado {
	var encontrado = false
	var simbolo Simbolo
	switch tipo {
	case FUNCION:
		for key, value := range scope.tablaFunciones {
			if key == id {
				simbolo = value.(Simbolo)
				encontrado = true
				break
			}
		}

	case MODULO:
		for key, value := range scope.tablaModulos {
			if key == id {
				//Verificar que sea publica
				simbolo = value.(Simbolo)
				encontrado = true
				break
			}
		}

	case STRUCT:
		for key, value := range scope.tablaStructs {
			if key == id {
				//Verificar que sea publica
				simbolo = value.(Simbolo)
				encontrado = true
				break
			}
		}
	}
	if encontrado {
		return TipoRetornado{
			Tipo:  SIMBOLO,
			Valor: simbolo,
		}
	}
	return TipoRetornado{
		Valor: true,
		Tipo:  NULL,
	}
}

func (scope *Scope) Update_fms_local(ident string, simbolo Simbolo) {
	//Buscar entodos los scopes desde un local
	//Primero conseguir el scope global
	scope_Actual := scope
	var retorno TipoRetornado
	id := strings.ToUpper(ident)
	//Buscar el scope global

	for scope_Actual = scope; scope_Actual != nil; scope_Actual = scope_Actual.prev {
		//Buscar el módulo en los entornos locales
		//Verificar que la fms exista y si puede ser accedida
		retorno = ActualizarEnMap(id, FUNCION, scope_Actual, scope, simbolo)
		if retorno.Valor == true {
			break
		}
	}
}

func ActualizarEnMap(id string, tipo TipoDato, scope *Scope, local *Scope, simbolo Simbolo) TipoRetornado {
	var encontrado = false

	for key, _ := range scope.tablaFunciones {
		if key == id {

			encontrado = true
			scope.tablaFunciones[key] = simbolo
			break
		}
	}
	if encontrado {
		return TipoRetornado{
			Tipo:  BOOLEAN,
			Valor: true,
		}
	}
	return TipoRetornado{
		Valor: false,
		Tipo:  BOOLEAN,
	}
}

/*fms = funcion modulo struct*/
func (scope *Scope) Addfms(simbolo Simbolo) {
	var scope_global *Scope = scope
	id := strings.ToUpper(simbolo.Identificador)
	//Recuperar el scope global
	if scope.prev != nil {
		for scope_global = scope; scope_global.prev != nil; scope_global = scope_global.prev {
			//Buscando el scope global
		}
	} else {
		scope_global = scope
	}
	switch simbolo.Tipo {
	case FUNCION:
		scope.tablaFunciones[id] = simbolo
	case MODULO:
		scope.tablaModulos[id] = simbolo
	case STRUCT_TEMPLATE:
		scope.tablaStructs[id] = simbolo
	}
}

func (scope *Scope) AddfmsParticular(simbolo Simbolo) {
	scope_global := scope
	id := strings.ToUpper(simbolo.Identificador)

	switch simbolo.Tipo {
	case FUNCION:
		scope_global.tablaFunciones[id] = simbolo
	case MODULO:
		scope_global.tablaModulos[id] = simbolo
	case STRUCT_TEMPLATE:
		scope_global.tablaStructs[id] = simbolo
	}
}

/*
func (scope *Scope) Exist_acceso(lista *arraylist.List) bool {
	var global *Scope
	var validador bool
	var simbolo_actual Simbolo
	//Recuperar el global
	for scope_actual := scope; scope_actual != nil; scope_actual = scope_actual.prev {
		if scope_actual.prev == nil{
			global = scope_actual
		}
	}
	for i:= 0; i< lista.Len();i++{
		//Obtener los padres
		var actual string = lista.GetValue(i).(string)
		for key, simbolo := range scope.tablaSimbolos {
			if key == actual {
				validador = true
				simbolo_actual = simbolo.(Simbolo)
				break
			}
		}
		if validador{
			actual = simbolo.
		}
	}



	return false
}
*/

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

func (scope *Scope) UpdateValor(ident string, valorNuevo interface{}) {
	id := strings.ToUpper(ident)
	var simbolo Simbolo
	for scope_actual := scope; scope_actual != nil; scope_actual = scope_actual.prev {

		for key, value := range scope_actual.tablaSimbolos {
			if key == id {
				simbolo = value.(Simbolo)
				simbolo.Valor = valorNuevo
				scope_actual.tablaSimbolos[key] = simbolo
				return
			}
		}
	}

}

func GetEntornoGlobal(scope *Scope) *Scope {
	var global *Scope = scope
	for scope_actual := scope; scope_actual != nil; scope_actual = scope_actual.prev {
		global = scope_actual
	}
	return global
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

func (scope *Scope) UpdateReferencias() TipoRetornado {
	var valor TipoRetornado
	var scopeReferencia *Scope
	var simboloReferencia Simbolo
	for _, simboloGuardado := range scope.tablaSimbolos {
		if simboloGuardado.(Simbolo).Referencia {
			valor = simboloGuardado.(Simbolo).Valor.(TipoRetornado)
			simboloReferencia = *simboloGuardado.(Simbolo).Referencia_puntero
			scopeReferencia = simboloReferencia.Entorno
			simboloReferencia.Valor = valor
			//Get el símbolo
			scopeReferencia.UpdateSimbolo(simboloReferencia.Identificador, simboloReferencia)
			break
		}
	}
	return TipoRetornado{Valor: true, Tipo: EJECUTADO}
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
		scope_global.Codigo += s.Codigo
		for i := 0; i < s.Errores.Len(); i++ {
			elemento := s.Errores.GetValue(i)
			scope_global.Errores.Add(elemento)
		}
		for i := 0; i < s.tablaSimbolosReporte.Len(); i++ {
			elemento := s.tablaSimbolosReporte.GetValue(i)
			scope_global.tablaSimbolosReporte.Add(elemento)
		}
		/*
			for i := 0; i < s.Temporales.Len(); i++ {
				elemento := s.Temporales.GetValue(i)
				scope_global.Temporales.Add(elemento)
			}
		*/
	}

}

func (s *Scope) GetEntornoPadreReturn() *Scope {
	var scopePadre *Scope
	if s.prev != nil {
		for scopePadre = s; scopePadre.prev != nil; scopePadre = scopePadre.prev {
			//Buscando el scope padre
			if scopePadre.Nombre == "funcion" {
				break
			}
		}
	}

	return scopePadre
}

func (s *Scope) GetEntornoPadreBreak() *Scope {
	var scopePadre *Scope
	scopePadre = s
	if s.prev != nil {
		for scopePadre = s; scopePadre.prev != nil; scopePadre = scopePadre.prev {
			//Buscando el scope padre
			if scopePadre.Nombre == "Case" {
				break
			}
		}
	}

	return scopePadre
}

func (s *Scope) GetEntornoPadreIF() *Scope {
	var scopePadre *Scope
	if s.prev != nil {
		for scopePadre = s; scopePadre.prev != nil; scopePadre = scopePadre.prev {
			//Buscando el scope padre
			if scopePadre.Nombre == "IF_I" {
				break
			}
		}
	}

	return scopePadre
}

func (entorno *Scope) Clonar(scope *Scope) interface{} {
	nTablaFunciones := make(map[string]interface{})
	nTablaSimbolos := make(map[string]interface{})
	nTablaStructs := make(map[string]interface{})
	nTablaModulos := make(map[string]interface{})
	nTablaSimbolosReporte := arraylist.New()
	nPrev := entorno.prev
	nErrores := arraylist.New()
	nGlobal := entorno.Global
	nConsola := entorno.Consola
	nNombre := entorno.Nombre
	nEntorno := NewScope(nNombre, nPrev)
	//Copiar las listas
	for key, value := range entorno.tablaFunciones {
		nTablaFunciones[key] = value
	}
	for key, value := range entorno.tablaSimbolos {
		nTablaSimbolos[key] = value
	}
	for key, value := range entorno.tablaStructs {
		nTablaStructs[key] = value
	}
	for key, value := range entorno.tablaModulos {
		nTablaModulos[key] = value
	}
	for i := 0; i < entorno.Errores.Len()-1; i++ {
		nErrores.Add(entorno.Errores.GetValue(i))
	}
	for i := 0; i < entorno.tablaSimbolosReporte.Len()-1; i++ {
		nTablaSimbolosReporte.Add(entorno.tablaSimbolosReporte.GetValue(i))
	}
	nEntorno.Global = nGlobal
	nEntorno.Consola = nConsola
	nEntorno.Nombre = nNombre
	nEntorno.tablaFunciones = nTablaFunciones
	nEntorno.tablaSimbolos = nTablaSimbolos
	nEntorno.tablaModulos = nTablaModulos
	nEntorno.tablaStructs = nTablaStructs
	nEntorno.tablaSimbolosReporte = nTablaSimbolosReporte
	nEntorno.Errores = nErrores
	nEntorno.prev = nPrev
	return &nEntorno
}

func (entorno *Scope) GenerarTablaSimbolos() {
	var simbolo SimboloReporte
	var ruta, nombre string
	inicio := `
	digraph {
		tablaSimbolos [
		  shape=plaintext
		  label=<
			<table border='0' cellborder='1' color='black' cellspacing='0'>
			  <tr>
				  <td>Id</td>
				  <td>Tipo símbolo</td>
				  <td>Tipo dato</td>
				  <td>Ámbito</td>
				  <td>Fila</td>
				  <td>Columna</td>
			  </tr>
	`
	final := `
		</table>
		>];
	}
	`
	if entorno.tablaSimbolosReporte.Len() > 0 {
		for i := 0; i < entorno.tablaSimbolosReporte.Len(); i++ {
			simbolo = entorno.tablaSimbolosReporte.GetValue(i).(SimboloReporte)
			inicio += GenerarFilaSimboloReporte(simbolo)
		}
		//Dot listo para la conversion
		inicio += final

		//Crear el dot y obtener la ruta donde fue creado
		ruta, nombre = CrearDot(inicio, "tablaSimbolos")
		GenerarGrafica(ruta, nombre)
	}
}

func (entorno *Scope) GenerarTablaErrores() {
	var ruta, nombre string
	inicio := `
	digraph {
		tablaSimbolos [
		  shape=plaintext
		  label=<
			<table border='0' cellborder='1' color='black' cellspacing='0'>
			  <tr>
				  <td>No.</td>
				  <td>Descripción</td>
				  <td>Ámbito</td>
				  <td>Línea</td>
				  <td>Columna</td>
				  <td>Fecha y hora</td>
			  </tr>
	`
	final := `
		</table>
		>];
	}
	`
	if entorno.Errores.Len() > 0 {
		for i := 0; i < entorno.Errores.Len(); i++ {
			err := entorno.Errores.GetValue(i)
			inicio += GenerarFilaErrores(err, i)
		}
		//Dot listo para la conversion
		inicio += final

		//Crear el dot y obtener la ruta donde fue creado
		ruta, nombre = CrearDot(inicio, "tablaErrores")
		GenerarGrafica(ruta, nombre)
	}
}

func (entorno *Scope) GenerarTablaBD() {
	var ruta, nombre string
	var contador = 0
	inicio := `
	digraph {
		tablaSimbolos [
		  shape=plaintext
		  label=<
			<table border='0' cellborder='1' color='black' cellspacing='0'>
			  <tr>
				  <td>No.</td>
				  <td>Nombre</td>
				  <td>No. Tablas</td>
				  <td>Línea</td>
			  </tr>`
	final := `
		</table>
		>];
	}
	`
	for _, value := range entorno.tablaModulos {
		mod := value.(Simbolo).Valor.(TipoRetornado).Valor
		inicio += GenerarFilaBD(mod, contador)
		contador++
	}
	if contador > 0 {
		//Dot listo para la conversion
		inicio += final

		//Crear el dot y obtener la ruta donde fue creado
		ruta, nombre = CrearDot(inicio, "tablaBD")
		GenerarGrafica(ruta, nombre)
	}
}

func (entorno *Scope) GenerarTablaTablas() {
	var ruta, nombre, nombreBD string
	var contador = 0
	inicio := `
	digraph {
		tablaSimbolos [
		  shape=plaintext
		  label=<
			<table border='0' cellborder='1' color='black' cellspacing='0'>
			  <tr>
				  <td>No.</td>
				  <td>Nombre tabla</td>
				  <td>Nombre de la base de datos</td>
				  <td>Línea</td>
			  </tr>`
	final := `
		</table>
		>];
	}
	`
	for _, value := range entorno.tablaModulos {
		mod := value.(Simbolo).Valor.(TipoRetornado).Valor
		nombreBD = mod.(Modulos).GetNombre()
		modulos := mod.(Modulos).GetEntorno().GetTablaModulos()
		for _, value1 := range modulos {
			mod1 := value1.(Simbolo).Valor.(TipoRetornado).Valor
			inicio += GenerarFilaTabla(mod1, contador, nombreBD)
			contador++
		}
	}
	if contador > 0 {
		//Dot listo para la conversion
		inicio += final
		//Crear el dot y obtener la ruta donde fue creado
		ruta, nombre = CrearDot(inicio, "tablas")
		GenerarGrafica(ruta, nombre)
	}
}

func GenerarFilaErrores(err interface{}, num int) string {
	salida := "" + "\n" +
		"<tr>" + "\n" +
		"<td cellpadding='4'>" + "\n" +
		strconv.Itoa(num+1) + "\n" +
		"</td>" + "\n" +
		"<td cellpadding='4'>" +
		err.(Error).GetDescripcion() + "\n" +
		"</td>" + "\n" +
		"<td cellpadding='4'>" +
		err.(Error).GetAmbito() + "\n" +
		"</td>" + "\n" +
		"<td cellpadding='4'>" +
		strconv.Itoa(err.(Error).GetFila()) + "\n" +
		"</td>" + "\n" +
		"<td cellpadding='4'>" + "\n" +
		strconv.Itoa(err.(Error).GetColumna()) + "\n" +
		"</td>" + "\n" +
		"<td cellpadding='4'>" + "\n" +
		err.(Error).GetFecha() + "\n" +
		"</td>" + "\n" +
		"</tr>" + "\n"
	return salida
}

func GenerarFilaSimboloReporte(simbolo SimboloReporte) string {
	tipo := strings.Replace(simbolo.TipoDato, ">", "-", -1)
	tipo = strings.Replace(tipo, "<", "-", -1)

	salida := "" + "\n" +
		"<tr>" + "\n" +
		"<td cellpadding='4'>" + "\n" +
		simbolo.Identificador + "\n" +
		"</td>" + "\n" +
		"<td cellpadding='4'>" +
		simbolo.TipoSimbolo + "\n" +
		"</td>" + "\n" +
		"<td cellpadding='4'>" +
		tipo + "\n" +
		"</td>" + "\n" +
		"<td cellpadding='4'>" +
		simbolo.Scope + "\n" +
		"</td>" + "\n" +
		"<td cellpadding='4'>" + "\n" +
		strconv.Itoa(simbolo.Fila) + "\n" +
		"</td>" + "\n" +
		"<td cellpadding='4'>" + "\n" +
		strconv.Itoa(simbolo.Columna) + "\n" +
		"</td>" + "\n" +
		"</tr>" + "\n"
	return salida
}

func GenerarFilaBD(mod interface{}, num int) string {
	salida := "" + "\n" +
		"<tr>" + "\n" +
		"<td cellpadding='4'>" + "\n" +
		strconv.Itoa(num+1) + "\n" +
		"</td>" + "\n" +
		"<td cellpadding='4'>" +
		mod.(Modulos).GetNombre() + "\n" +
		"</td>" + "\n" +
		"<td cellpadding='4'>" +
		strconv.Itoa(mod.(Modulos).GetTablas()) + "\n" +
		"</td>" + "\n" +
		"<td cellpadding='4'>" +
		strconv.Itoa(mod.(Modulos).GetFila()) + "\n" +
		"</td>" + "\n" +
		"</tr>" + "\n"
	return salida
}

func GenerarFilaTabla(mod interface{}, num int, bd string) string {
	salida := "" + "\n" +
		"<tr>" + "\n" +
		"<td cellpadding='4'>" + "\n" +
		strconv.Itoa(num+1) + "\n" +
		"</td>" + "\n" +
		"<td cellpadding='4'>" +
		mod.(Modulos).GetNombre() + "\n" +
		"</td>" + "\n" +
		"<td cellpadding='4'>" +
		bd + "\n" +
		"</td>" + "\n" +
		"<td cellpadding='4'>" +
		strconv.Itoa(mod.(Modulos).GetFila()) + "\n" +
		"</td>" + "\n" +
		"</tr>" + "\n"
	return salida
}

func CrearDot(contenido string, nombre string) (string, string) {
	//Get el directorio del proyecto
	dir, _ := os.Getwd()
	dir += "\\Web\\tablas"
	extension := ".dot"
	//Crear el archivo
	file, err := os.Create(dir + "\\" + nombre + extension)
	if err != nil {
		fmt.Println(err)
	} else {
		file.WriteString(contenido)
		fmt.Println("Done")
	}
	file.Close()
	//Generar un delay
	for i := 0; i < 2000; i++ {

	}
	return dir, nombre
}

func GenerarGrafica(ruta, nombre string) {
	comando := ruta + "\\" + nombre + ".dot"
	tipo := "-Tsvg"
	path := "C:\\Program Files\\Graphviz\\bin\\dot.exe"
	cmd, _ := exec.Command(path, tipo, comando).Output()
	mode := int(0777)
	ioutil.WriteFile("Web\\tablas\\"+nombre+".svg", cmd, os.FileMode(mode))

}

func (s *Scope) AgregarPrint(cadena string) {
	//Obtener el scope global
	var scope_global *Scope
	if s.prev != nil {
		for scope_global = s; scope_global.prev != nil; scope_global = scope_global.prev {
			//Buscando el scope global
		}
	} else {
		scope_global = s
	}

	scope_global.Consola += cadena
}

func (s *Scope) GetNivel() int {
	salida := 0
	var scope_global *Scope
	if s.prev != nil {
		for scope_global = s; scope_global.prev != nil; scope_global = scope_global.prev {
			//Buscando el scope global
			salida++
		}
	}
	return salida
}
