// Code generated from N3parser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // N3parser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseN3parserListener is a complete listener for a parse tree produced by N3parser.
type BaseN3parserListener struct{}

var _ N3parserListener = &BaseN3parserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseN3parserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseN3parserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseN3parserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseN3parserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterInicio is called when production inicio is entered.
func (s *BaseN3parserListener) EnterInicio(ctx *InicioContext) {}

// ExitInicio is called when production inicio is exited.
func (s *BaseN3parserListener) ExitInicio(ctx *InicioContext) {}

// EnterFunciones is called when production funciones is entered.
func (s *BaseN3parserListener) EnterFunciones(ctx *FuncionesContext) {}

// ExitFunciones is called when production funciones is exited.
func (s *BaseN3parserListener) ExitFunciones(ctx *FuncionesContext) {}

// EnterFuncion is called when production funcion is entered.
func (s *BaseN3parserListener) EnterFuncion(ctx *FuncionContext) {}

// ExitFuncion is called when production funcion is exited.
func (s *BaseN3parserListener) ExitFuncion(ctx *FuncionContext) {}

// EnterBloques is called when production bloques is entered.
func (s *BaseN3parserListener) EnterBloques(ctx *BloquesContext) {}

// ExitBloques is called when production bloques is exited.
func (s *BaseN3parserListener) ExitBloques(ctx *BloquesContext) {}

// EnterBloque is called when production bloque is entered.
func (s *BaseN3parserListener) EnterBloque(ctx *BloqueContext) {}

// ExitBloque is called when production bloque is exited.
func (s *BaseN3parserListener) ExitBloque(ctx *BloqueContext) {}

// EnterBloque_i is called when production bloque_i is entered.
func (s *BaseN3parserListener) EnterBloque_i(ctx *Bloque_iContext) {}

// ExitBloque_i is called when production bloque_i is exited.
func (s *BaseN3parserListener) ExitBloque_i(ctx *Bloque_iContext) {}

// EnterInstrucciones is called when production instrucciones is entered.
func (s *BaseN3parserListener) EnterInstrucciones(ctx *InstruccionesContext) {}

// ExitInstrucciones is called when production instrucciones is exited.
func (s *BaseN3parserListener) ExitInstrucciones(ctx *InstruccionesContext) {}

// EnterInstruccion is called when production instruccion is entered.
func (s *BaseN3parserListener) EnterInstruccion(ctx *InstruccionContext) {}

// ExitInstruccion is called when production instruccion is exited.
func (s *BaseN3parserListener) ExitInstruccion(ctx *InstruccionContext) {}

// EnterBloque_f is called when production bloque_f is entered.
func (s *BaseN3parserListener) EnterBloque_f(ctx *Bloque_fContext) {}

// ExitBloque_f is called when production bloque_f is exited.
func (s *BaseN3parserListener) ExitBloque_f(ctx *Bloque_fContext) {}

// EnterPrint3d is called when production print3d is entered.
func (s *BaseN3parserListener) EnterPrint3d(ctx *Print3dContext) {}

// ExitPrint3d is called when production print3d is exited.
func (s *BaseN3parserListener) ExitPrint3d(ctx *Print3dContext) {}

// EnterIf3d is called when production if3d is entered.
func (s *BaseN3parserListener) EnterIf3d(ctx *If3dContext) {}

// ExitIf3d is called when production if3d is exited.
func (s *BaseN3parserListener) ExitIf3d(ctx *If3dContext) {}

// EnterAsignacion is called when production asignacion is entered.
func (s *BaseN3parserListener) EnterAsignacion(ctx *AsignacionContext) {}

// ExitAsignacion is called when production asignacion is exited.
func (s *BaseN3parserListener) ExitAsignacion(ctx *AsignacionContext) {}

// EnterOperacion is called when production operacion is entered.
func (s *BaseN3parserListener) EnterOperacion(ctx *OperacionContext) {}

// ExitOperacion is called when production operacion is exited.
func (s *BaseN3parserListener) ExitOperacion(ctx *OperacionContext) {}

// EnterExpresion is called when production expresion is entered.
func (s *BaseN3parserListener) EnterExpresion(ctx *ExpresionContext) {}

// ExitExpresion is called when production expresion is exited.
func (s *BaseN3parserListener) ExitExpresion(ctx *ExpresionContext) {}

// EnterEtiqueta is called when production etiqueta is entered.
func (s *BaseN3parserListener) EnterEtiqueta(ctx *EtiquetaContext) {}

// ExitEtiqueta is called when production etiqueta is exited.
func (s *BaseN3parserListener) ExitEtiqueta(ctx *EtiquetaContext) {}

// EnterSalto is called when production salto is entered.
func (s *BaseN3parserListener) EnterSalto(ctx *SaltoContext) {}

// ExitSalto is called when production salto is exited.
func (s *BaseN3parserListener) ExitSalto(ctx *SaltoContext) {}

// EnterRetorno is called when production retorno is entered.
func (s *BaseN3parserListener) EnterRetorno(ctx *RetornoContext) {}

// ExitRetorno is called when production retorno is exited.
func (s *BaseN3parserListener) ExitRetorno(ctx *RetornoContext) {}

// EnterInclude is called when production include is entered.
func (s *BaseN3parserListener) EnterInclude(ctx *IncludeContext) {}

// ExitInclude is called when production include is exited.
func (s *BaseN3parserListener) ExitInclude(ctx *IncludeContext) {}

// EnterDeclaraciones is called when production declaraciones is entered.
func (s *BaseN3parserListener) EnterDeclaraciones(ctx *DeclaracionesContext) {}

// ExitDeclaraciones is called when production declaraciones is exited.
func (s *BaseN3parserListener) ExitDeclaraciones(ctx *DeclaracionesContext) {}

// EnterDeclaracion is called when production declaracion is entered.
func (s *BaseN3parserListener) EnterDeclaracion(ctx *DeclaracionContext) {}

// ExitDeclaracion is called when production declaracion is exited.
func (s *BaseN3parserListener) ExitDeclaracion(ctx *DeclaracionContext) {}

// EnterTemporalesTexto is called when production temporalesTexto is entered.
func (s *BaseN3parserListener) EnterTemporalesTexto(ctx *TemporalesTextoContext) {}

// ExitTemporalesTexto is called when production temporalesTexto is exited.
func (s *BaseN3parserListener) ExitTemporalesTexto(ctx *TemporalesTextoContext) {}

// EnterTemporalesLista is called when production temporalesLista is entered.
func (s *BaseN3parserListener) EnterTemporalesLista(ctx *TemporalesListaContext) {}

// ExitTemporalesLista is called when production temporalesLista is exited.
func (s *BaseN3parserListener) ExitTemporalesLista(ctx *TemporalesListaContext) {}

// EnterOperador is called when production operador is entered.
func (s *BaseN3parserListener) EnterOperador(ctx *OperadorContext) {}

// ExitOperador is called when production operador is exited.
func (s *BaseN3parserListener) ExitOperador(ctx *OperadorContext) {}
