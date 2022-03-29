// Code generated from Nparser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Nparser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// NparserListener is a complete listener for a parse tree produced by Nparser.
type NparserListener interface {
	antlr.ParseTreeListener

	// EnterInicio is called when entering the inicio production.
	EnterInicio(c *InicioContext)

	// EnterInstruccionesGlobales is called when entering the instruccionesGlobales production.
	EnterInstruccionesGlobales(c *InstruccionesGlobalesContext)

	// EnterInstruccionesModulos is called when entering the instruccionesModulos production.
	EnterInstruccionesModulos(c *InstruccionesModulosContext)

	// EnterInstruccionesControl is called when entering the instruccionesControl production.
	EnterInstruccionesControl(c *InstruccionesControlContext)

	// EnterInstrucciones is called when entering the instrucciones production.
	EnterInstrucciones(c *InstruccionesContext)

	// EnterBloque is called when entering the bloque production.
	EnterBloque(c *BloqueContext)

	// EnterBloque_control is called when entering the bloque_control production.
	EnterBloque_control(c *Bloque_controlContext)

	// EnterBloque_modulo is called when entering the bloque_modulo production.
	EnterBloque_modulo(c *Bloque_moduloContext)

	// EnterInstruccionGlobal is called when entering the instruccionGlobal production.
	EnterInstruccionGlobal(c *InstruccionGlobalContext)

	// EnterInstruccionModulo is called when entering the instruccionModulo production.
	EnterInstruccionModulo(c *InstruccionModuloContext)

	// EnterInstruccion is called when entering the instruccion production.
	EnterInstruccion(c *InstruccionContext)

	// EnterInstruccionControl is called when entering the instruccionControl production.
	EnterInstruccionControl(c *InstruccionControlContext)

	// EnterFuncion_main is called when entering the funcion_main production.
	EnterFuncion_main(c *Funcion_mainContext)

	// EnterDeclaracion is called when entering the declaracion production.
	EnterDeclaracion(c *DeclaracionContext)

	// EnterDeclaracion_struct_template is called when entering the declaracion_struct_template production.
	EnterDeclaracion_struct_template(c *Declaracion_struct_templateContext)

	// EnterAtributos_struct_template is called when entering the atributos_struct_template production.
	EnterAtributos_struct_template(c *Atributos_struct_templateContext)

	// EnterAtributo_struct_template is called when entering the atributo_struct_template production.
	EnterAtributo_struct_template(c *Atributo_struct_templateContext)

	// EnterStruct_instancia is called when entering the struct_instancia production.
	EnterStruct_instancia(c *Struct_instanciaContext)

	// EnterAtributos_struct_instancia is called when entering the atributos_struct_instancia production.
	EnterAtributos_struct_instancia(c *Atributos_struct_instanciaContext)

	// EnterAtributo_struct_instancia is called when entering the atributo_struct_instancia production.
	EnterAtributo_struct_instancia(c *Atributo_struct_instanciaContext)

	// EnterDeclaracion_modulo is called when entering the declaracion_modulo production.
	EnterDeclaracion_modulo(c *Declaracion_moduloContext)

	// EnterDeclaracion_funcion is called when entering the declaracion_funcion production.
	EnterDeclaracion_funcion(c *Declaracion_funcionContext)

	// EnterAsignacion is called when entering the asignacion production.
	EnterAsignacion(c *AsignacionContext)

	// EnterAccesos_vector_array_asignacion is called when entering the accesos_vector_array_asignacion production.
	EnterAccesos_vector_array_asignacion(c *Accesos_vector_array_asignacionContext)

	// EnterExpresion_logica is called when entering the expresion_logica production.
	EnterExpresion_logica(c *Expresion_logicaContext)

	// EnterExpresion_relacional is called when entering the expresion_relacional production.
	EnterExpresion_relacional(c *Expresion_relacionalContext)

	// EnterExpresion is called when entering the expresion production.
	EnterExpresion(c *ExpresionContext)

	// EnterTipo_dato is called when entering the tipo_dato production.
	EnterTipo_dato(c *Tipo_datoContext)

	// EnterControl_if is called when entering the control_if production.
	EnterControl_if(c *Control_ifContext)

	// EnterBloque_else_if is called when entering the bloque_else_if production.
	EnterBloque_else_if(c *Bloque_else_ifContext)

	// EnterElse_if is called when entering the else_if production.
	EnterElse_if(c *Else_ifContext)

	// EnterControl_if_exp is called when entering the control_if_exp production.
	EnterControl_if_exp(c *Control_if_expContext)

	// EnterBloque_else_if_exp is called when entering the bloque_else_if_exp production.
	EnterBloque_else_if_exp(c *Bloque_else_if_expContext)

	// EnterElse_if_exp is called when entering the else_if_exp production.
	EnterElse_if_exp(c *Else_if_expContext)

	// EnterControl_expresion is called when entering the control_expresion production.
	EnterControl_expresion(c *Control_expresionContext)

	// EnterControl_match is called when entering the control_match production.
	EnterControl_match(c *Control_matchContext)

	// EnterControl_case is called when entering the control_case production.
	EnterControl_case(c *Control_caseContext)

	// EnterCases is called when entering the cases production.
	EnterCases(c *CasesContext)

	// EnterCase_match is called when entering the case_match production.
	EnterCase_match(c *Case_matchContext)

	// EnterControl_match_exp is called when entering the control_match_exp production.
	EnterControl_match_exp(c *Control_match_expContext)

	// EnterControl_case_exp is called when entering the control_case_exp production.
	EnterControl_case_exp(c *Control_case_expContext)

	// EnterCases_exp is called when entering the cases_exp production.
	EnterCases_exp(c *Cases_expContext)

	// EnterCase_match_exp is called when entering the case_match_exp production.
	EnterCase_match_exp(c *Case_match_expContext)

	// EnterIreturn is called when entering the ireturn production.
	EnterIreturn(c *IreturnContext)

	// EnterIbreak is called when entering the ibreak production.
	EnterIbreak(c *IbreakContext)

	// EnterIcontinue is called when entering the icontinue production.
	EnterIcontinue(c *IcontinueContext)

	// EnterControl_loop is called when entering the control_loop production.
	EnterControl_loop(c *Control_loopContext)

	// EnterControl_loop_exp is called when entering the control_loop_exp production.
	EnterControl_loop_exp(c *Control_loop_expContext)

	// EnterPrintNormal is called when entering the printNormal production.
	EnterPrintNormal(c *PrintNormalContext)

	// EnterPrintFormato is called when entering the printFormato production.
	EnterPrintFormato(c *PrintFormatoContext)

	// EnterElementosPrint is called when entering the elementosPrint production.
	EnterElementosPrint(c *ElementosPrintContext)

	// EnterControl_while is called when entering the control_while production.
	EnterControl_while(c *Control_whileContext)

	// EnterParametros_funcion is called when entering the parametros_funcion production.
	EnterParametros_funcion(c *Parametros_funcionContext)

	// EnterParametro is called when entering the parametro production.
	EnterParametro(c *ParametroContext)

	// EnterLlamada_funcion is called when entering the llamada_funcion production.
	EnterLlamada_funcion(c *Llamada_funcionContext)

	// EnterParametros_llamada is called when entering the parametros_llamada production.
	EnterParametros_llamada(c *Parametros_llamadaContext)

	// EnterParametro_llamada_referencia is called when entering the parametro_llamada_referencia production.
	EnterParametro_llamada_referencia(c *Parametro_llamada_referenciaContext)

	// EnterElementos_vector is called when entering the elementos_vector production.
	EnterElementos_vector(c *Elementos_vectorContext)

	// EnterMetodos_iniciar_vector is called when entering the metodos_iniciar_vector production.
	EnterMetodos_iniciar_vector(c *Metodos_iniciar_vectorContext)

	// EnterMetodos_vector is called when entering the metodos_vector production.
	EnterMetodos_vector(c *Metodos_vectorContext)

	// EnterPotencia is called when entering the potencia production.
	EnterPotencia(c *PotenciaContext)

	// EnterArray is called when entering the array production.
	EnterArray(c *ArrayContext)

	// EnterDimension_array is called when entering the dimension_array production.
	EnterDimension_array(c *Dimension_arrayContext)

	// EnterDimension_acceso_array is called when entering the dimension_acceso_array production.
	EnterDimension_acceso_array(c *Dimension_acceso_arrayContext)

	// EnterTipo_dato_tipo is called when entering the tipo_dato_tipo production.
	EnterTipo_dato_tipo(c *Tipo_dato_tipoContext)

	// EnterAcceso_modulo is called when entering the acceso_modulo production.
	EnterAcceso_modulo(c *Acceso_moduloContext)

	// EnterAcceso_modulo_elementos is called when entering the acceso_modulo_elementos production.
	EnterAcceso_modulo_elementos(c *Acceso_modulo_elementosContext)

	// EnterAcceso_modulo_elemento_final is called when entering the acceso_modulo_elemento_final production.
	EnterAcceso_modulo_elemento_final(c *Acceso_modulo_elemento_finalContext)

	// EnterControl_for is called when entering the control_for production.
	EnterControl_for(c *Control_forContext)

	// EnterRango_for is called when entering the rango_for production.
	EnterRango_for(c *Rango_forContext)

	// ExitInicio is called when exiting the inicio production.
	ExitInicio(c *InicioContext)

	// ExitInstruccionesGlobales is called when exiting the instruccionesGlobales production.
	ExitInstruccionesGlobales(c *InstruccionesGlobalesContext)

	// ExitInstruccionesModulos is called when exiting the instruccionesModulos production.
	ExitInstruccionesModulos(c *InstruccionesModulosContext)

	// ExitInstruccionesControl is called when exiting the instruccionesControl production.
	ExitInstruccionesControl(c *InstruccionesControlContext)

	// ExitInstrucciones is called when exiting the instrucciones production.
	ExitInstrucciones(c *InstruccionesContext)

	// ExitBloque is called when exiting the bloque production.
	ExitBloque(c *BloqueContext)

	// ExitBloque_control is called when exiting the bloque_control production.
	ExitBloque_control(c *Bloque_controlContext)

	// ExitBloque_modulo is called when exiting the bloque_modulo production.
	ExitBloque_modulo(c *Bloque_moduloContext)

	// ExitInstruccionGlobal is called when exiting the instruccionGlobal production.
	ExitInstruccionGlobal(c *InstruccionGlobalContext)

	// ExitInstruccionModulo is called when exiting the instruccionModulo production.
	ExitInstruccionModulo(c *InstruccionModuloContext)

	// ExitInstruccion is called when exiting the instruccion production.
	ExitInstruccion(c *InstruccionContext)

	// ExitInstruccionControl is called when exiting the instruccionControl production.
	ExitInstruccionControl(c *InstruccionControlContext)

	// ExitFuncion_main is called when exiting the funcion_main production.
	ExitFuncion_main(c *Funcion_mainContext)

	// ExitDeclaracion is called when exiting the declaracion production.
	ExitDeclaracion(c *DeclaracionContext)

	// ExitDeclaracion_struct_template is called when exiting the declaracion_struct_template production.
	ExitDeclaracion_struct_template(c *Declaracion_struct_templateContext)

	// ExitAtributos_struct_template is called when exiting the atributos_struct_template production.
	ExitAtributos_struct_template(c *Atributos_struct_templateContext)

	// ExitAtributo_struct_template is called when exiting the atributo_struct_template production.
	ExitAtributo_struct_template(c *Atributo_struct_templateContext)

	// ExitStruct_instancia is called when exiting the struct_instancia production.
	ExitStruct_instancia(c *Struct_instanciaContext)

	// ExitAtributos_struct_instancia is called when exiting the atributos_struct_instancia production.
	ExitAtributos_struct_instancia(c *Atributos_struct_instanciaContext)

	// ExitAtributo_struct_instancia is called when exiting the atributo_struct_instancia production.
	ExitAtributo_struct_instancia(c *Atributo_struct_instanciaContext)

	// ExitDeclaracion_modulo is called when exiting the declaracion_modulo production.
	ExitDeclaracion_modulo(c *Declaracion_moduloContext)

	// ExitDeclaracion_funcion is called when exiting the declaracion_funcion production.
	ExitDeclaracion_funcion(c *Declaracion_funcionContext)

	// ExitAsignacion is called when exiting the asignacion production.
	ExitAsignacion(c *AsignacionContext)

	// ExitAccesos_vector_array_asignacion is called when exiting the accesos_vector_array_asignacion production.
	ExitAccesos_vector_array_asignacion(c *Accesos_vector_array_asignacionContext)

	// ExitExpresion_logica is called when exiting the expresion_logica production.
	ExitExpresion_logica(c *Expresion_logicaContext)

	// ExitExpresion_relacional is called when exiting the expresion_relacional production.
	ExitExpresion_relacional(c *Expresion_relacionalContext)

	// ExitExpresion is called when exiting the expresion production.
	ExitExpresion(c *ExpresionContext)

	// ExitTipo_dato is called when exiting the tipo_dato production.
	ExitTipo_dato(c *Tipo_datoContext)

	// ExitControl_if is called when exiting the control_if production.
	ExitControl_if(c *Control_ifContext)

	// ExitBloque_else_if is called when exiting the bloque_else_if production.
	ExitBloque_else_if(c *Bloque_else_ifContext)

	// ExitElse_if is called when exiting the else_if production.
	ExitElse_if(c *Else_ifContext)

	// ExitControl_if_exp is called when exiting the control_if_exp production.
	ExitControl_if_exp(c *Control_if_expContext)

	// ExitBloque_else_if_exp is called when exiting the bloque_else_if_exp production.
	ExitBloque_else_if_exp(c *Bloque_else_if_expContext)

	// ExitElse_if_exp is called when exiting the else_if_exp production.
	ExitElse_if_exp(c *Else_if_expContext)

	// ExitControl_expresion is called when exiting the control_expresion production.
	ExitControl_expresion(c *Control_expresionContext)

	// ExitControl_match is called when exiting the control_match production.
	ExitControl_match(c *Control_matchContext)

	// ExitControl_case is called when exiting the control_case production.
	ExitControl_case(c *Control_caseContext)

	// ExitCases is called when exiting the cases production.
	ExitCases(c *CasesContext)

	// ExitCase_match is called when exiting the case_match production.
	ExitCase_match(c *Case_matchContext)

	// ExitControl_match_exp is called when exiting the control_match_exp production.
	ExitControl_match_exp(c *Control_match_expContext)

	// ExitControl_case_exp is called when exiting the control_case_exp production.
	ExitControl_case_exp(c *Control_case_expContext)

	// ExitCases_exp is called when exiting the cases_exp production.
	ExitCases_exp(c *Cases_expContext)

	// ExitCase_match_exp is called when exiting the case_match_exp production.
	ExitCase_match_exp(c *Case_match_expContext)

	// ExitIreturn is called when exiting the ireturn production.
	ExitIreturn(c *IreturnContext)

	// ExitIbreak is called when exiting the ibreak production.
	ExitIbreak(c *IbreakContext)

	// ExitIcontinue is called when exiting the icontinue production.
	ExitIcontinue(c *IcontinueContext)

	// ExitControl_loop is called when exiting the control_loop production.
	ExitControl_loop(c *Control_loopContext)

	// ExitControl_loop_exp is called when exiting the control_loop_exp production.
	ExitControl_loop_exp(c *Control_loop_expContext)

	// ExitPrintNormal is called when exiting the printNormal production.
	ExitPrintNormal(c *PrintNormalContext)

	// ExitPrintFormato is called when exiting the printFormato production.
	ExitPrintFormato(c *PrintFormatoContext)

	// ExitElementosPrint is called when exiting the elementosPrint production.
	ExitElementosPrint(c *ElementosPrintContext)

	// ExitControl_while is called when exiting the control_while production.
	ExitControl_while(c *Control_whileContext)

	// ExitParametros_funcion is called when exiting the parametros_funcion production.
	ExitParametros_funcion(c *Parametros_funcionContext)

	// ExitParametro is called when exiting the parametro production.
	ExitParametro(c *ParametroContext)

	// ExitLlamada_funcion is called when exiting the llamada_funcion production.
	ExitLlamada_funcion(c *Llamada_funcionContext)

	// ExitParametros_llamada is called when exiting the parametros_llamada production.
	ExitParametros_llamada(c *Parametros_llamadaContext)

	// ExitParametro_llamada_referencia is called when exiting the parametro_llamada_referencia production.
	ExitParametro_llamada_referencia(c *Parametro_llamada_referenciaContext)

	// ExitElementos_vector is called when exiting the elementos_vector production.
	ExitElementos_vector(c *Elementos_vectorContext)

	// ExitMetodos_iniciar_vector is called when exiting the metodos_iniciar_vector production.
	ExitMetodos_iniciar_vector(c *Metodos_iniciar_vectorContext)

	// ExitMetodos_vector is called when exiting the metodos_vector production.
	ExitMetodos_vector(c *Metodos_vectorContext)

	// ExitPotencia is called when exiting the potencia production.
	ExitPotencia(c *PotenciaContext)

	// ExitArray is called when exiting the array production.
	ExitArray(c *ArrayContext)

	// ExitDimension_array is called when exiting the dimension_array production.
	ExitDimension_array(c *Dimension_arrayContext)

	// ExitDimension_acceso_array is called when exiting the dimension_acceso_array production.
	ExitDimension_acceso_array(c *Dimension_acceso_arrayContext)

	// ExitTipo_dato_tipo is called when exiting the tipo_dato_tipo production.
	ExitTipo_dato_tipo(c *Tipo_dato_tipoContext)

	// ExitAcceso_modulo is called when exiting the acceso_modulo production.
	ExitAcceso_modulo(c *Acceso_moduloContext)

	// ExitAcceso_modulo_elementos is called when exiting the acceso_modulo_elementos production.
	ExitAcceso_modulo_elementos(c *Acceso_modulo_elementosContext)

	// ExitAcceso_modulo_elemento_final is called when exiting the acceso_modulo_elemento_final production.
	ExitAcceso_modulo_elemento_final(c *Acceso_modulo_elemento_finalContext)

	// ExitControl_for is called when exiting the control_for production.
	ExitControl_for(c *Control_forContext)

	// ExitRango_for is called when exiting the rango_for production.
	ExitRango_for(c *Rango_forContext)
}
