// Code generated from N3parser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // N3parser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// N3parserListener is a complete listener for a parse tree produced by N3parser.
type N3parserListener interface {
	antlr.ParseTreeListener

	// EnterInicio is called when entering the inicio production.
	EnterInicio(c *InicioContext)

	// EnterFunciones is called when entering the funciones production.
	EnterFunciones(c *FuncionesContext)

	// EnterFuncion is called when entering the funcion production.
	EnterFuncion(c *FuncionContext)

	// EnterLlamada is called when entering the llamada production.
	EnterLlamada(c *LlamadaContext)

	// EnterBloques is called when entering the bloques production.
	EnterBloques(c *BloquesContext)

	// EnterBloque is called when entering the bloque production.
	EnterBloque(c *BloqueContext)

	// EnterBloque_i is called when entering the bloque_i production.
	EnterBloque_i(c *Bloque_iContext)

	// EnterInstrucciones is called when entering the instrucciones production.
	EnterInstrucciones(c *InstruccionesContext)

	// EnterInstruccion is called when entering the instruccion production.
	EnterInstruccion(c *InstruccionContext)

	// EnterBloque_f is called when entering the bloque_f production.
	EnterBloque_f(c *Bloque_fContext)

	// EnterPrint3d is called when entering the print3d production.
	EnterPrint3d(c *Print3dContext)

	// EnterIf3d is called when entering the if3d production.
	EnterIf3d(c *If3dContext)

	// EnterAsignacion is called when entering the asignacion production.
	EnterAsignacion(c *AsignacionContext)

	// EnterOperacion is called when entering the operacion production.
	EnterOperacion(c *OperacionContext)

	// EnterExpresion is called when entering the expresion production.
	EnterExpresion(c *ExpresionContext)

	// EnterEtiqueta is called when entering the etiqueta production.
	EnterEtiqueta(c *EtiquetaContext)

	// EnterSalto is called when entering the salto production.
	EnterSalto(c *SaltoContext)

	// EnterRetorno is called when entering the retorno production.
	EnterRetorno(c *RetornoContext)

	// EnterInclude is called when entering the include production.
	EnterInclude(c *IncludeContext)

	// EnterDeclaraciones is called when entering the declaraciones production.
	EnterDeclaraciones(c *DeclaracionesContext)

	// EnterDeclaracion is called when entering the declaracion production.
	EnterDeclaracion(c *DeclaracionContext)

	// EnterTemporalesTexto is called when entering the temporalesTexto production.
	EnterTemporalesTexto(c *TemporalesTextoContext)

	// EnterTemporalesLista is called when entering the temporalesLista production.
	EnterTemporalesLista(c *TemporalesListaContext)

	// EnterOperador is called when entering the operador production.
	EnterOperador(c *OperadorContext)

	// ExitInicio is called when exiting the inicio production.
	ExitInicio(c *InicioContext)

	// ExitFunciones is called when exiting the funciones production.
	ExitFunciones(c *FuncionesContext)

	// ExitFuncion is called when exiting the funcion production.
	ExitFuncion(c *FuncionContext)

	// ExitLlamada is called when exiting the llamada production.
	ExitLlamada(c *LlamadaContext)

	// ExitBloques is called when exiting the bloques production.
	ExitBloques(c *BloquesContext)

	// ExitBloque is called when exiting the bloque production.
	ExitBloque(c *BloqueContext)

	// ExitBloque_i is called when exiting the bloque_i production.
	ExitBloque_i(c *Bloque_iContext)

	// ExitInstrucciones is called when exiting the instrucciones production.
	ExitInstrucciones(c *InstruccionesContext)

	// ExitInstruccion is called when exiting the instruccion production.
	ExitInstruccion(c *InstruccionContext)

	// ExitBloque_f is called when exiting the bloque_f production.
	ExitBloque_f(c *Bloque_fContext)

	// ExitPrint3d is called when exiting the print3d production.
	ExitPrint3d(c *Print3dContext)

	// ExitIf3d is called when exiting the if3d production.
	ExitIf3d(c *If3dContext)

	// ExitAsignacion is called when exiting the asignacion production.
	ExitAsignacion(c *AsignacionContext)

	// ExitOperacion is called when exiting the operacion production.
	ExitOperacion(c *OperacionContext)

	// ExitExpresion is called when exiting the expresion production.
	ExitExpresion(c *ExpresionContext)

	// ExitEtiqueta is called when exiting the etiqueta production.
	ExitEtiqueta(c *EtiquetaContext)

	// ExitSalto is called when exiting the salto production.
	ExitSalto(c *SaltoContext)

	// ExitRetorno is called when exiting the retorno production.
	ExitRetorno(c *RetornoContext)

	// ExitInclude is called when exiting the include production.
	ExitInclude(c *IncludeContext)

	// ExitDeclaraciones is called when exiting the declaraciones production.
	ExitDeclaraciones(c *DeclaracionesContext)

	// ExitDeclaracion is called when exiting the declaracion production.
	ExitDeclaracion(c *DeclaracionContext)

	// ExitTemporalesTexto is called when exiting the temporalesTexto production.
	ExitTemporalesTexto(c *TemporalesTextoContext)

	// ExitTemporalesLista is called when exiting the temporalesLista production.
	ExitTemporalesLista(c *TemporalesListaContext)

	// ExitOperador is called when exiting the operador production.
	ExitOperador(c *OperadorContext)
}
