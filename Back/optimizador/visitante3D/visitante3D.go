package visitante3D

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/optimizador/elementos/bloque3d"
	"Back/optimizador/elementos/funciones3d"
	"Back/optimizador/elementos/headers3d"
	"Back/optimizador/elementos/interfaces3d"
	parser "Back/optimizador/parser3D"
	"fmt"

	"github.com/colegno/arraylist"
)

type Visitador3D struct {
	*parser.BaseN3parserListener
	Consola       string
	Codigo        string
	Errores       arraylist.List
	EntornoGlobal Ast.Scope
}

func NewVisitor() *Visitador3D {
	return new(Visitador3D)
}

func (v *Visitador3D) ExitInicio(ctx *parser.InicioContext) {
	fmt.Println("Ingreso a visitante 3d>>>>>>>>>>>>>>>")
	/**************ELEMENTOS PARA EL BLOQUE GLOBAL*********************/
	var includes interfaces3d.Expresion3D
	funciones := arraylist.New()
	declaraciones := arraylist.New()
	var bloqueGlobal bloque3d.Bloque3dGlobal
	/*****************************************************************/
	instrucciones := ctx.GetEx()
	//Verificar que no existan errores sintácticos o semánticos
	if v.Errores.Len() > 0 {
		EntornoGlobal := Ast.NewScope("global", nil)
		EntornoGlobal.Global = true
		for i := 0; i < v.Errores.Len(); i++ {
			err := v.Errores.GetValue(i).(errores.CustomSyntaxError)
			nError := errores.NewError(err.Fila, err.Columna, err.Msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			nError.Ambito = EntornoGlobal.GetTipoScope()
			EntornoGlobal.Errores.Add(nError)
		}
		EntornoGlobal.GenerarTablaErrores()
		return
	}
	/*GET LOS ELEMENTOS DEL BLOQUE GLOBAL*/
	bloqueGlobal = instrucciones
	funciones = bloqueGlobal.Funciones
	declaraciones = bloqueGlobal.Declaraciones
	includes = bloqueGlobal.Include

	/*AGREGAR EL INCLUDE INICIAL*/
	v.Consola += includes.GetValor()

	/*RECORRER LA LISTA DE DECLARACIONES*/
	var declaracionActual headers3d.Declaracion3D
	for i := 0; i < declaraciones.Len(); i++ {
		declaracionActual = declaraciones.GetValue(i).(headers3d.Declaracion3D)
		v.Consola += declaracionActual.GetValor()
	}

	/*RECORRER LA LISTA DE FUNCIONES*/
	var funcionActual funciones3d.Funcion3D
	for i := 0; i < funciones.Len(); i++ {
		funcionActual = funciones.GetValue(i).(funciones3d.Funcion3D)
		v.Consola += funcionActual.GetValor()
	}

	//v.Consola = instrucciones.GetValor()
}

func (v *Visitador3D) GetConsola() string {
	//return v.Consola
	return v.Consola
}

func (v *Visitador3D) UpdateConsola(entrada string) {
	v.Consola += entrada + "\n"
}

func (v *Visitador3D) GetCodigo3D() string {
	return v.Codigo
}
