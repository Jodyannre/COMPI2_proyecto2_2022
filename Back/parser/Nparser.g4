parser grammar Nparser;

options{
    tokenVocab = Nlexer;
    language = Go;
}


@header{
    import "github.com/colegno/arraylist"
    import "Back/analizador/Ast"
    import "Back/analizador/expresiones"
    import "Back/analizador/instrucciones"
    import "Back/analizador/exp_ins"
    import "Back/analizador/transferencia"
    import "Back/analizador/bucles"
    import "Back/analizador/Ast/simbolos"
    import "Back/analizador/fn_primitivas"
    import "Back/analizador/fn_vectores"
    import "Back/analizador/fn_array"
}

/* 
@members{
}
*/


inicio returns[*arraylist.List lista] 
            : instruccionesGlobales	
            {
                $lista = $instruccionesGlobales.list
            }
;


instruccionesGlobales returns [*arraylist.List list]
			@init{
				$list =  arraylist.New()
			}
			: e += instruccionGlobal*  {
				listInt := localctx.(*InstruccionesGlobalesContext).GetE()
					for _, e := range listInt {
						$list.Add(e.GetEx())
					}
			}
;


instruccionesModulos returns [*arraylist.List list]
			@init{
				$list =  arraylist.New()
			}
			: e += instruccionModulo*  {
				listInt := localctx.(*InstruccionesModulosContext).GetE()
					for _, e := range listInt {
						$list.Add(e.GetEx())
					}
			}
;

instruccionesControl returns [*arraylist.List list]
			@init{
				$list =  arraylist.New()
			}
			: e += instruccionControl*  {
				listInt := localctx.(*InstruccionesControlContext).GetE()
					for _, e := range listInt {
						$list.Add(e.GetEx())
					}
			}
;

/* Para el main, funciones, etc */
instrucciones returns [*arraylist.List list]
			@init{
				$list =  arraylist.New()
			}
			: e += instruccion*  {
				listInt := localctx.(*InstruccionesContext).GetE()
					for _, e := range listInt {
						$list.Add(e.GetEx())
					}
			}
;


bloque returns[*arraylist.List list] 
            : LLAVE_IZQ instrucciones LLAVE_DER	
            {
                $list = $instrucciones.list
            }
;

bloque_control returns [*arraylist.List list]
            : LLAVE_IZQ instruccionesControl LLAVE_DER	
            {
                $list = $instruccionesControl.list
            }
;

bloque_modulo returns[*arraylist.List list]
            : LLAVE_IZQ instruccionesModulos LLAVE_DER	
            {
                $list = $instruccionesModulos.list
            }
;


instruccionGlobal returns[interface{} ex]
            : funcion_main                  {$ex = $funcion_main.ex}
            | declaracion_struct_template   {$ex = $declaracion_struct_template.ex} 
            | declaracion_funcion           {$ex = $declaracion_funcion.ex} 
            | declaracion_modulo            {$ex = $declaracion_modulo.ex} 
;

instruccionModulo returns[interface{} ex]
            : declaracion_struct_template   {$ex = $declaracion_struct_template.ex} 
            | declaracion_funcion           {$ex = $declaracion_funcion.ex} 
            | declaracion_modulo            {$ex = $declaracion_modulo.ex} 

;


instruccion returns[interface{} ex] 
			:llamada_funcion PUNTOCOMA  {$ex = $llamada_funcion.ex}
            |asignacion PUNTOCOMA       {$ex = $asignacion.ex}
            |expresion PUNTOCOMA   	    {$ex = $expresion.ex}	
            //|expresion                  {$ex = $expresion.ex}
            |declaracion PUNTOCOMA      {$ex = $declaracion.ex}	
            //|declaracion_funcion        {$ex = $declaracion_funcion.ex}    
            |control_if                 {$ex = $control_if.ex}	 
            |control_match              {$ex = $control_match.ex}   
            |control_loop               {$ex = $control_loop.ex}
            |control_while              {$ex = $control_while.ex}
            |control_for                {$ex = $control_for.ex}
            |ibreak PUNTOCOMA           {$ex = $ibreak.ex}             
            |icontinue PUNTOCOMA        {$ex = $icontinue.ex} 
            |ireturn PUNTOCOMA          {$ex = $ireturn.ex} 
            |printNormal PUNTOCOMA      {$ex = $printNormal.ex} 
            |printFormato PUNTOCOMA     {$ex = $printFormato.ex} 
            |metodos_vector PUNTOCOMA   {$ex = $metodos_vector.ex} 
            //|declaracion_struct_template {$ex = $declaracion_struct_template.ex}

;



instruccionControl returns[interface{} ex] 
			:llamada_funcion PUNTOCOMA  {$ex = $llamada_funcion.ex}
            |asignacion PUNTOCOMA       {$ex = $asignacion.ex}
            |expresion PUNTOCOMA   	    {$ex = $expresion.ex}	
            |expresion                  {$ex = $expresion.ex}
            |declaracion PUNTOCOMA      {$ex = $declaracion.ex}	
            //|declaracion_funcion        {$ex = $declaracion_funcion.ex}      
            |control_if                 {$ex = $control_if.ex}	 
            |control_match              {$ex = $control_match.ex}   
            |control_loop               {$ex = $control_loop.ex}
            |control_while              {$ex = $control_while.ex}
            |control_for                {$ex = $control_for.ex}
            |ibreak PUNTOCOMA           {$ex = $ibreak.ex}             
            |icontinue PUNTOCOMA        {$ex = $icontinue.ex} 
            |ireturn PUNTOCOMA          {$ex = $ireturn.ex} 
            |printNormal PUNTOCOMA      {$ex = $printNormal.ex} 
            |printFormato PUNTOCOMA     {$ex = $printFormato.ex} 
            |metodos_vector PUNTOCOMA   {$ex = $metodos_vector.ex} 
            //|declaracion_struct_template {$ex = $declaracion_struct_template.ex}

;



funcion_main returns[Ast.Expresion ex]
    : FN MAIN PAR_IZQ PAR_DER bloque 
        {
            fila := $FN.line
            columna := $FN.pos     
            $ex = simbolos.NewFuncionMain($bloque.list,fila,columna)             
        }
    | FN MAIN PAR_IZQ PAR_DER LLAVE_IZQ LLAVE_DER
        {
            fila := $FN.line
            columna := $FN.pos
            instrucciones := arraylist.New()   
            $ex = simbolos.NewFuncionMain(instrucciones,fila,columna)           
        }
;


declaracion returns[Ast.Instruccion ex]
    : LET ID IGUAL expresion_logica
        {
            fila := $LET.line
            columna := $LET.pos
            $ex = instrucciones.NewDeclaracionSinTipo($ID.text,$expresion_logica.ex,
            false,false,fila,columna)
        }
    | LET ID DOSPUNTOS VEC MENOR tipo=tipo_dato_tipo MAYOR IGUAL expresion_logica
        {
            fila := $LET.line
            columna := $LET.pos 
            tipoVector := $tipo.ex
            $ex = instrucciones.NewDeclaracionVector($ID.text,tipoVector,$expresion_logica.ex,false,false,fila,columna)    
    
        }
    | LET MUT ID DOSPUNTOS VEC MENOR tipo=tipo_dato_tipo MAYOR IGUAL expresion_logica
        {
            fila := $LET.line
            columna := $LET.pos 
            tipoVector := $tipo.ex
            $ex = instrucciones.NewDeclaracionVector($ID.text,tipoVector,$expresion_logica.ex,true,false,fila,columna)            
        }
    | LET ID DOSPUNTOS dimension=dimension_array IGUAL expresion_logica
        {
            fila := $LET.line
            columna := $LET.pos 
            $ex = instrucciones.NewDeclaracionArray($ID.text,$dimension.ex,false,false,$expresion_logica.ex,fila,columna)            
        }
    | LET MUT ID DOSPUNTOS dimension=dimension_array IGUAL expresion_logica
        {
            fila := $LET.line
            columna := $LET.pos 
            $ex = instrucciones.NewDeclaracionArray($ID.text,$dimension.ex,true,false,$expresion_logica.ex,fila,columna)            
        }
        
    | LET ID IGUAL control_expresion
        {
            fila := $LET.line
            columna := $LET.pos
            $ex = instrucciones.NewDeclaracionSinTipo($ID.text,$control_expresion.ex,
            false,false,fila,columna)
        }
    | LET ID DOSPUNTOS tipo_dato_tipo IGUAL expresion_logica
        {
            fila := $LET.line
            columna := $LET.pos
            $ex = instrucciones.NewDeclaracionTotal($ID.text,$expresion_logica.ex,
            $tipo_dato_tipo.ex,false,false,fila,columna)             
        }
    | LET ID DOSPUNTOS tipo_dato_tipo IGUAL control_expresion
        {
            fila := $LET.line
            columna := $LET.pos
            $ex = instrucciones.NewDeclaracionTotal($ID.text,$control_expresion.ex,
            $tipo_dato_tipo.ex,false,false,fila,columna)   
        }
    | LET MUT ID IGUAL expresion_logica
        {
            fila := $LET.line
            columna := $LET.pos
            $ex = instrucciones.NewDeclaracionSinTipo($ID.text,$expresion_logica.ex,
            true,false,fila,columna)                 
        }
    | LET MUT ID IGUAL control_expresion
        {
            fila := $ID.line
            columna := $ID.pos
            $ex = instrucciones.NewDeclaracionSinTipo($ID.text,$control_expresion.ex,
            true,false,fila,columna) 
        }
    | LET MUT ID DOSPUNTOS tipo_dato_tipo
        {        
            fila := $LET.line
            columna := $LET.pos
            //valor := expresiones.NewPrimitivo(nil, Ast.NULL,fila,columna)
            $ex = instrucciones.NewDeclaracionConTipo($ID.text,$tipo_dato_tipo.ex,
            true,false,fila,columna)            
        }
        /* 
    | LET MUT ID DOSPUNTOS tipo=tipo_dato_tipo
        {        
            fila := $LET.line
            columna := $LET.pos
            valor := expresiones.NewPrimitivo(nil, Ast.NULL,fila,columna)
            $ex = instrucciones.NewDeclaracion($ID.text,$tipo.ex,
            true,false,Ast.VOID,valor,fila,columna)                
        }
    */
    
    | LET MUT ID DOSPUNTOS tipo_dato_tipo IGUAL expresion_logica
        {
            fila := $LET.line
            columna := $LET.pos
            $ex = instrucciones.NewDeclaracionTotal($ID.text,$expresion_logica.ex,
            $tipo_dato_tipo.ex,true,false,fila,columna)                
        }
    | LET MUT ID DOSPUNTOS tipo_dato_tipo IGUAL control_expresion
        {
            fila := $LET.line
            columna := $LET.pos
            $ex = instrucciones.NewDeclaracionTotal($ID.text,$control_expresion.ex,
            $tipo_dato_tipo.ex,true,false,fila,columna)                
        } 
        /*  
    | LET ID IGUAL expresion
        {
            fila := $LET.line
            columna := $LET.pos 
            $ex = instrucciones.NewDeclaracion($ID.text,Ast.VECTOR,false,false,Ast.VOID,$expresion.ex,fila,columna)
        }
    */
    | LET MUT ID IGUAL expresion
        {
            fila := $LET.line
            columna := $LET.pos 
            $ex = instrucciones.NewDeclaracionSinTipo($ID.text,$expresion.ex,
            true,false,fila,columna)
        }
   


;

declaracion_struct_template returns [Ast.Instruccion ex]
: PUB STRUCT id=ID_CAMEL LLAVE_IZQ att=atributos_struct_template LLAVE_DER
    {
        fila := $id.line
        columna := $id.pos       
        //tipo := simbolos.NewTipo(Ast.STRUCT, $id.text, fila,columna) 
        $ex = simbolos.NewDeclaracionStructTemplate($id.text,$att.list,true,fila,columna)   
    }
| STRUCT id=ID_CAMEL LLAVE_IZQ att=atributos_struct_template LLAVE_DER
    {
        fila := $id.line
        columna := $id.pos       
        //tipo := simbolos.NewTipo(Ast.STRUCT, $id.text, fila,columna) 
        $ex = simbolos.NewDeclaracionStructTemplate($id.text,$att.list,false,fila,columna)        
    }
;




atributos_struct_template returns [*arraylist.List list]
@init{$list = arraylist.New()}
    : lista_elementos = atributos_struct_template COMA att=atributo_struct_template 
        {
            $lista_elementos.list.Add($att.ex)
            $list = $lista_elementos.list
        }
    | att=atributo_struct_template 
        {
            $list.Add($att.ex)
        }
;


atributo_struct_template returns [Ast.Expresion ex]
    : ID DOSPUNTOS tipo=tipo_dato_tipo 
        {
            fila := $ID.line
            columna := $ID.pos
            $ex = simbolos.NewAtributoTemplate($ID.text,$tipo.ex,false,fila,columna)
        }
    | PUB ID DOSPUNTOS tipo=tipo_dato_tipo 
        {
            fila := $ID.line
            columna := $ID.pos
            $ex = simbolos.NewAtributoTemplate($ID.text,$tipo.ex,true,fila,columna)             
        }
;

struct_instancia returns [Ast.Expresion ex]
    :   id=ID_CAMEL LLAVE_IZQ att=atributos_struct_instancia LLAVE_DER
    {
        fila := $id.line
        columna := $id.pos       
        tipo := Ast.TipoRetornado{Tipo:Ast.STRUCT, Valor:$id.text} 
        $ex = simbolos.NewStructInstancia(tipo,$att.list,false,fila,columna)
    }
    |   tipo_dato_tipo LLAVE_IZQ att=atributos_struct_instancia LLAVE_DER
    {
        fila := $LLAVE_IZQ.line
        columna := $LLAVE_IZQ.pos -1      
        //tipo := Ast.TipoRetornado{Tipo:Ast.ACCESO_MODULO, Valor:$tipo_dato_tipo.ex} 
        $ex = simbolos.NewStructInstancia($tipo_dato_tipo.ex,$att.list,false,fila,columna)
    }
;


atributos_struct_instancia returns [*arraylist.List list]
@init{$list = arraylist.New()}
    : lista_elementos = atributos_struct_instancia COMA att1=atributo_struct_instancia 
        {
            $lista_elementos.list.Add($att1.ex)
            $list = $lista_elementos.list
        }
    | att2=atributo_struct_instancia 
        {
            $list.Add($att2.ex)
        }
;


atributo_struct_instancia returns [Ast.Expresion ex]
    : ID DOSPUNTOS expresion
    {
        fila := $ID.line
        columna := $ID.pos
        $ex = simbolos.NewAtributo($ID.text,$expresion.ex,false,fila,columna)
    }
;


declaracion_modulo returns [Ast.Instruccion ex]
    :   MOD ID bloque_modulo
        {
            fila := $MOD.line
            columna := $MOD.pos
            id := expresiones.NewIdentificador($ID.text,Ast.IDENTIFICADOR,fila,columna) 
            modulo := simbolos.NewModulo(id,$bloque_modulo.list,false,fila,columna)
            $ex = simbolos.NewDeclaracionModulo(modulo,false,fila,columna)

        }
    |   PUB MOD ID bloque_modulo
        {
            fila := $PUB.line
            columna := $PUB.pos
            id := expresiones.NewIdentificador($ID.text,Ast.IDENTIFICADOR,fila,columna)          
            modulo := simbolos.NewModulo(id,$bloque_modulo.list,true,fila,columna)
            $ex = simbolos.NewDeclaracionModulo(modulo,true,fila,columna)
        }
;


declaracion_funcion returns [Ast.Instruccion ex]

    : PUB FN ID PAR_IZQ PAR_DER FN_TIPO_RETORNO tipo_dato_tipo bloque 
        {
            parametros := arraylist.New() 
            fila := $FN.line
            columna := $FN.pos
            funcion := simbolos.NewFuncion($ID.text,Ast.FUNCION,$bloque.list,
            parametros,$tipo_dato_tipo.ex,true,fila,columna)
            $ex = simbolos.NewDeclaracionFuncion($ID.text,funcion,$tipo_dato_tipo.ex,
            false,true,fila,columna)            
        } 
    | PUB FN ID PAR_IZQ PAR_DER bloque
        {
            parametros := arraylist.New() 
            fila := $FN.line
            columna := $FN.pos
            funcion := simbolos.NewFuncion($ID.text,Ast.FUNCION,$bloque.list,
            parametros,Ast.TipoRetornado{Tipo:Ast.VOID,Valor:true},true,fila,columna)
            nTipo := Ast.TipoRetornado{Tipo:Ast.VOID,Valor:true}
            $ex = simbolos.NewDeclaracionFuncion($ID.text,funcion,nTipo,
            false,true,fila,columna)            
        }

    | FN ID PAR_IZQ PAR_DER FN_TIPO_RETORNO tipo_dato_tipo bloque
        {
            fila := $FN.line
            columna := $FN.pos
            parametros := arraylist.New()
            funcion := simbolos.NewFuncion($ID.text,Ast.FUNCION,$bloque.list,
            parametros,$tipo_dato_tipo.ex,false,fila,columna)
            $ex = simbolos.NewDeclaracionFuncion($ID.text,funcion,$tipo_dato_tipo.ex,
            false,false,fila,columna)  
        }

    | FN ID PAR_IZQ PAR_DER bloque
        {
            fila := $FN.line
            columna := $FN.pos
            parametros := arraylist.New()
            funcion := simbolos.NewFuncion($ID.text,Ast.FUNCION,$bloque.list,
            parametros,Ast.TipoRetornado{Tipo:Ast.VOID,Valor:true},false,fila,columna)
            nTipo := Ast.TipoRetornado{Tipo:Ast.VOID,Valor:true}
            $ex = simbolos.NewDeclaracionFuncion($ID.text,funcion,nTipo,
            false,false,fila,columna)  
        }

    | PUB FN ID PAR_IZQ parametros_funcion PAR_DER FN_TIPO_RETORNO tipo_dato_tipo bloque
        {
            fila := $FN.line
            columna := $FN.pos
            funcion := simbolos.NewFuncion($ID.text,Ast.FUNCION,$bloque.list,
            $parametros_funcion.list,$tipo_dato_tipo.ex,true,fila,columna)
            $ex = simbolos.NewDeclaracionFuncion($ID.text,funcion,$tipo_dato_tipo.ex,
            false,true,fila,columna)            
        }

    | PUB FN ID PAR_IZQ parametros_funcion PAR_DER bloque
        {
            fila := $FN.line
            columna := $FN.pos
            funcion := simbolos.NewFuncion($ID.text,Ast.FUNCION,$bloque.list,
            $parametros_funcion.list,Ast.TipoRetornado{Tipo:Ast.VOID,Valor:true},true,fila,columna)
            nTipo := Ast.TipoRetornado{Tipo:Ast.VOID,Valor:true}
            $ex = simbolos.NewDeclaracionFuncion($ID.text,funcion,nTipo,
            false,true,fila,columna)            
        }

    | FN ID PAR_IZQ parametros_funcion PAR_DER FN_TIPO_RETORNO tipo_dato_tipo bloque
        {
            fila := $FN.line
            columna := $FN.pos
            funcion := simbolos.NewFuncion($ID.text,Ast.FUNCION,$bloque.list,
            $parametros_funcion.list,$tipo_dato_tipo.ex,true,fila,columna)
            $ex = simbolos.NewDeclaracionFuncion($ID.text,funcion,$tipo_dato_tipo.ex,
            false,false,fila,columna)  
        }

    | FN ID PAR_IZQ parametros_funcion PAR_DER bloque
        {
            fila := $FN.line
            columna := $FN.pos
            funcion := simbolos.NewFuncion($ID.text,Ast.FUNCION,$bloque.list,
            $parametros_funcion.list,Ast.TipoRetornado{Tipo:Ast.VOID,Valor:true},true,fila,columna)
            nTipo := Ast.TipoRetornado{Tipo:Ast.VOID,Valor:true}
            $ex = simbolos.NewDeclaracionFuncion($ID.text,funcion,nTipo,
            false,false,fila,columna)  
        }
;

asignacion returns[Ast.Instruccion ex]
    : id=accesos_vector_array_asignacion IGUAL elemento=expresion_logica
    {
        fila := $IGUAL.line
        columna := $IGUAL.pos-1
        $ex = instrucciones.NewAsignacion($id.ex,$elemento.ex,fila,columna)        
    }
    | id=accesos_vector_array_asignacion IGUAL control_expresion
    {
        fila := $IGUAL.line
        columna := $IGUAL.pos
        $ex = instrucciones.NewAsignacion($id.ex,$control_expresion.ex,fila,columna)
    }
    | ID IGUAL expresion_logica
    {
        fila := $ID.line
        columna := $ID.pos
        id := expresiones.NewIdentificador($ID.text, Ast.IDENTIFICADOR,fila,columna)
        $ex = instrucciones.NewAsignacion(id,$expresion_logica.ex,fila,columna)
    }
    | ID IGUAL control_expresion
    {
        fila := $ID.line
        columna := $ID.pos
        id := expresiones.NewIdentificador($ID.text, Ast.IDENTIFICADOR,fila,columna)
        $ex = instrucciones.NewAsignacion(id,$control_expresion.ex,fila,columna)
    }
    | idExp=expresion IGUAL valor=expresion_logica
    {
        fila := $IGUAL.line
        columna := $IGUAL.pos-1
        elemento := localctx.(*AsignacionContext).GetIdExp().GetEx()
        _,tipoParticular := elemento.(Ast.Abstracto).GetTipo()
        if tipoParticular == Ast.ACCESO_STRUCT{
            $ex = simbolos.NewAsignacionAccesoStruct(elemento,$valor.ex,fila,columna)  
        }else{
            $ex = instrucciones.NewAsignacion($idExp.ex,$valor.ex,fila,columna)  
        }
             
    }

    | ex1=expresion PUNTO atributo=ID IGUAL ex2=expresion_logica
        {
            filaS := $PUNTO.line
            columnaS := $PUNTO.pos-1
            filaA := $atributo.line
            columnaA := $atributo.pos-1
            idAtributo := expresiones.NewIdentificador($atributo.text,Ast.IDENTIFICADOR,filaA,columnaA)
            acceso := simbolos.NewAccesoStruct($ex1.ex,idAtributo,filaS,columnaS)
            elemento := localctx.(*AsignacionContext).GetEx1().GetEx()
            fila := elemento.(Ast.Abstracto).GetFila()
            columna := elemento.(Ast.Abstracto).GetColumna()
            $ex = simbolos.NewAsignacionStruct(acceso,$ex2.ex,fila,columna)

        }
    | ex1=expresion PUNTO atributo=ID IGUAL ex3=control_expresion
        {
            filaS := $PUNTO.line
            columnaS := $PUNTO.pos-1
            filaA := $atributo.line
            columnaA := $atributo.pos-1
            idAtributo := expresiones.NewIdentificador($atributo.text,Ast.IDENTIFICADOR,filaA,columnaA)
            acceso := simbolos.NewAccesoStruct($ex1.ex,idAtributo,filaS,columnaS)
            elemento := localctx.(*AsignacionContext).GetEx1().GetEx()
            fila := elemento.(Ast.Abstracto).GetFila()
            columna := elemento.(Ast.Abstracto).GetColumna()
            $ex = simbolos.NewAsignacionStruct(acceso,$ex3.ex,fila,columna)
        }

;

accesos_vector_array_asignacion returns [Ast.Expresion ex]
    :   ID lista=dimension_acceso_array
        {
            id := $ID.text
            fila := $ID.line
            columna := $ID.pos
            idE := expresiones.NewIdentificador(id,Ast.IDENTIFICADOR,fila,columna)      
            $ex = fn_array.NewAccesoArray(idE,$lista.list,fila,columna)           
        }
        //Acceso a vector
    |   ID CORCHETE_IZQ index=expresion CORCHETE_DER
        {
            id := $ID.text
            fila := $ID.line
            columna := $ID.pos
            idE := expresiones.NewIdentificador (id,Ast.IDENTIFICADOR,fila,columna)     
            $ex = fn_vectores.NewAccesoVec(idE,$index.ex,Ast.VEC_ACCESO,fila,columna)
        }
    |   ID
        {
            id := $ID.text
            fila := $ID.line
            columna := $ID.pos
            $ex = expresiones.NewIdentificador (id,Ast.IDENTIFICADOR,fila,columna)            
        }
;



expresion_logica returns[Ast.Expresion ex]
    :   op_izq=expresion_logica AND op_der=expresion_logica
        {
            fila := $AND.line
            columna := $AND.pos
            $ex = expresiones.NewOperation($op_izq.ex,$AND.text,$op_der.ex,false,fila,columna)
        }
    |   op_izq=expresion_logica OR op_der=expresion_logica
        {
            fila := $OR.line
            columna := $OR.pos
            $ex = expresiones.NewOperation($op_izq.ex,$OR.text,$op_der.ex,false,fila,columna)
        }
    |   expresion_relacional
        {
            $ex = $expresion_relacional.ex   
        }
;


expresion_relacional returns[Ast.Expresion ex]
    :  op_izq=expresion_relacional op=(MAYOR_I|MAYOR|MENOR_I|MENOR|IGUALDAD|DISTINTO) op_der=expresion_relacional
        {
            fila := $op.line
            columna := $op.pos
            $ex = expresiones.NewOperation($op_izq.ex,$op.text,$op_der.ex,false,fila,columna)
        }
    |   op_izq=expresion_relacional op=(IGUALDAD|DISTINTO) op_der=expresion_relacional
        {
            fila := $op.line
            columna := $op.pos
            $ex = expresiones.NewOperation($op_izq.ex,$op.text,$op_der.ex,false,fila,columna)
        }
    |   expresion 
    {
        $ex = $expresion.ex
    }
;

expresion returns[Ast.Expresion ex] 
    :   op=(RESTA|NOT) op_izq= expresion
        {
            fila := $op.line
            columna := $op.pos
            $ex = expresiones.NewOperation($op_izq.ex,$op.text,nil,true,fila,columna)
        }
    |   e=expresion PUNTO TO_STRING PAR_IZQ PAR_DER
        {
            fila := $PUNTO.line
            columna := $PUNTO.pos - 1
            $ex = fn_primitivas.NewToString(Ast.LLAMADA_FUNCION,$e.ex,fila,columna)
        }
    |   e=expresion PUNTO SQRT PAR_IZQ PAR_DER
        {
            fila := $PUNTO.line
            columna := $PUNTO.pos - 1
            $ex = fn_primitivas.NewSqrt(Ast.LLAMADA_FUNCION,$e.ex,fila,columna)
        }
    |   e=expresion PUNTO ABS PAR_IZQ PAR_DER
        {
            fila := $PUNTO.line
            columna := $PUNTO.pos - 1
            $ex = fn_primitivas.NewAbs(Ast.LLAMADA_FUNCION,$e.ex,fila,columna)
        }

        //Metodo para clonar
    |   exp=expresion PUNTO CLONE PAR_IZQ PAR_DER

        {
            fila:= $PUNTO.line
            columna:= $PUNTO.pos
            $ex = expresiones.NewClone($exp.ex,fila,columna)
        }

        //Metodo len de vector
    |   id=expresion PUNTO LEN PAR_IZQ PAR_DER
        {
            fila := $PUNTO.line
            columna := $PUNTO.pos
            $ex = fn_vectores.NewLenVec($id.ex,Ast.VEC_PUSH,fila,columna)
        }   
    | id=expresion PUNTO CAPACITY PAR_IZQ PAR_DER
        {
            fila := $PUNTO.line
            columna := $PUNTO.pos
            $ex = fn_vectores.NewCapacityVec($id.ex,Ast.VEC_CAPACITY,fila,columna)
        }
    | id=expresion PUNTO CONTAINS PAR_IZQ AMPERSAND exp=expresion PAR_DER
        {
            fila := $PUNTO.line
            columna := $PUNTO.pos
            $ex = fn_vectores.NewContainsVec($id.ex,$exp.ex,Ast.VEC_CONTAINS,fila,columna)            
        }
    | id=expresion PUNTO REMOVE PAR_IZQ index=expresion PAR_DER
        {
            fila := $PUNTO.line
            columna := $PUNTO.pos
            $ex = fn_vectores.NewRemoveVec($id.ex,$index.ex,Ast.VEC_REMOVE,fila,columna)

        }
        //Funcion to chars
    |   ex1=expresion PUNTO CHARS PAR_IZQ PAR_DER
        {
            elemento := localctx.(*ExpresionContext).GetEx1().GetEx()
            fila := elemento.(Ast.Abstracto).GetFila()
            columna := elemento.(Ast.Abstracto).GetColumna()
            $ex = fn_primitivas.NewToChars($ex1.ex,fila,columna)
        }

    |   op_izq=expresion op=(MULTIPLICACION|DIVISION|MODULO) op_der=expresion
        {
            fila := $op.line
            columna := $op.pos
            $ex = expresiones.NewOperation($op_izq.ex,$op.text,$op_der.ex,false,fila,columna)
        }    
    
    |   op_izq=expresion op=(SUMA|RESTA) op_der=expresion
        {
            fila := $op.line
            columna := $op.pos
            $ex = expresiones.NewOperation($op_izq.ex,$op.text,$op_der.ex,false,fila,columna)
        }
    |   PAR_IZQ expresion_logica PAR_DER
        {
            $ex = $expresion_logica.ex
        }
    |   PAR_IZQ expresion AS tipo_dato PAR_DER
        {
            //Cast
            fila := $PAR_IZQ.line
            columna := $PAR_IZQ.pos            
            $ex = expresiones.NewCast($expresion.ex, Ast.CAST, $tipo_dato.ex,fila,columna)
        } 
    |   llamada_funcion 
        {
            $ex = $llamada_funcion.ex   
        }
    |   metodos_iniciar_vector
        {
            $ex = $metodos_iniciar_vector.ex
        }
    |   potencia
        {
            $ex = $potencia.ex   
        }
    |   array
        {
            $ex = $array.ex
        }
    |   struct_instancia
        {
            $ex = $struct_instancia.ex   
        }
    |   acceso_modulo LLAVE_IZQ att=atributos_struct_instancia LLAVE_DER
        {
            fila := $LLAVE_IZQ.line
            columna := $LLAVE_IZQ.pos -1   
            acceso:= Ast.TipoRetornado{
                Valor: $acceso_modulo.ex,
                Tipo: Ast.ACCESO_MODULO,
            }
            $ex = simbolos.NewStructInstancia(acceso,$att.list,false,fila,columna)
        }
    |   obj=expresion PUNTO atributo=ID 
    {
            filaS := $PUNTO.line
            columnaS := $PUNTO.pos-1
            filaA := $atributo.line
            columnaA := $atributo.pos-1
            idAtributo := expresiones.NewIdentificador($atributo.text,Ast.IDENTIFICADOR,filaA,columnaA)
            fmt.Println(idAtributo)
            $ex = simbolos.NewAccesoStruct($obj.ex,idAtributo,filaS,columnaS)
    }
        //Acceso a array
    |   id=expresion lista=dimension_acceso_array
        {
            elemento := localctx.(*ExpresionContext).GetId().GetEx()
            fila := elemento.(Ast.Abstracto).GetFila()
            columna := elemento.(Ast.Abstracto).GetColumna() -1 
            $ex = fn_array.NewAccesoArray($id.ex,$lista.list,fila,columna)           
        }
        
        //Acceso a vector
    |   id=expresion CORCHETE_IZQ index=expresion CORCHETE_DER  
        {
            fila := $CORCHETE_IZQ.line
            columna := $CORCHETE_IZQ.pos-1 
            $ex = fn_vectores.NewAccesoVec($id.ex,$index.ex,Ast.VEC_ACCESO,fila,columna)
        }
        //Acceso a array
    |   ID lista=dimension_acceso_array
        {
            id := $ID.text
            fila := $ID.line
            columna := $ID.pos
            idE := expresiones.NewIdentificador(id,Ast.IDENTIFICADOR,fila,columna)      
            $ex = fn_array.NewAccesoArray(idE,$lista.list,fila,columna)           
        }
    |   ID		
        {
            id := $ID.text
            fila := $ID.line
            columna := $ID.pos
            $ex = expresiones.NewIdentificador (id,Ast.IDENTIFICADOR,fila,columna)
        }
    |   acceso_modulo
        {
            $ex = $acceso_modulo.ex
        }
    |   TRUE        
        {
            valor := true
            fila := $TRUE.line
            columna := $TRUE.pos
            $ex = expresiones.NewPrimitivo(valor, Ast.BOOLEAN,fila,columna)
        }         
    |   FALSE
        {
            valor := false
            fila := $FALSE.line
            columna := $FALSE.pos
            $ex = expresiones.NewPrimitivo(valor, Ast.BOOLEAN,fila,columna)
        }     
    |   CARACTER
        {
            valor := $CARACTER.text
            valor = valor[1:len(valor)-1]
            fila := $CARACTER.line
            columna := $CARACTER.pos
            $ex = expresiones.NewPrimitivo(valor, Ast.CHAR,fila,columna)
        }   
    |   DECIMAL
        {
            valor,err := strconv.ParseFloat($DECIMAL.text,64)
            if err!= nil{
                fmt.Println("Hay un error en el get número")
                fmt.Println(err)
            }
            fila := $DECIMAL.line
            columna := $DECIMAL.pos
            $ex = expresiones.NewPrimitivo(valor, Ast.F64,fila,columna)
        }   
    |   NUMERO
        {
            valor,err := strconv.Atoi($NUMERO.text)
            if err!= nil{
                fmt.Println("Hay un error en el get número")
                fmt.Println(err)
            }
            fila := $NUMERO.line
            columna := $NUMERO.pos
            $ex = expresiones.NewPrimitivo(valor, Ast.I64,fila,columna)
        }   
    |   CADENA
        {
            fila := $CADENA.line
            columna := $CADENA.pos
            valor := $CADENA.text
            valor = valor[1:len(valor)-1]
            $ex = expresiones.NewPrimitivo(valor, Ast.STR,fila,columna)
        }             
;

tipo_dato returns[Ast.TipoDato ex]
    :   BOOL    {$ex = Ast.BOOLEAN}
    |   CHAR    {$ex = Ast.CHAR}
    |   I64     {$ex = Ast.I64}
    |   F64     {$ex = Ast.F64}
    |   STR     {$ex = Ast.STR}
    |   STRING  {$ex = Ast.STRING}
    |   USIZE   {$ex = Ast.USIZE}
;



control_if returns[Ast.Instruccion ex]
	:IF expresion_logica bloqueIf = bloque
	{
		fila:= $IF.line
		columna:= $IF.pos
		columna++
		lista_null := arraylist.New()
		$ex = exp_ins.NewIF($expresion_logica.ex,$bloqueIf.list,lista_null,Ast.IF,fila,columna,false)
	}	
	|IF expresion_logica bloqueIf = bloque ELSE bloqueElse = bloque
	  
	{
		fila:= $IF.line
		columna:= $IF.pos
		columna++
        lista_entonces := arraylist.New()
        lista_null := arraylist.New()
        Else := exp_ins.NewIF ($expresion_logica.ex,$bloqueElse.list,lista_null,Ast.ELSE,fila,columna,false)
		lista_entonces.Add(Else)
		$ex = exp_ins.NewIF($expresion_logica.ex,$bloqueIf.list,lista_entonces,Ast.IF,fila,columna,false)	
	}
	|IF expresion_logica bloqueIf = bloque bloque_else_if
	{
		fila:= $IF.line
		columna:= $IF.pos
		columna++
		lista_entonces := $bloque_else_if.list
		$ex = exp_ins.NewIF($expresion_logica.ex,$bloqueIf.list,lista_entonces,Ast.IF,fila,columna,false)		
	}
	|IF expresion_logica bloqueIf = bloque bloque_else_if ELSE bloqueElse = bloque
	{
		fila:= $IF.line
		columna:= $IF.pos
		columna++
        lista_null := arraylist.New()
        Else := exp_ins.NewIF ($expresion_logica.ex,$bloqueElse.list,lista_null,Ast.ELSE,fila,columna,false)
		lista_entonces := $bloque_else_if.list
        lista_entonces.Add(Else)
		$ex = exp_ins.NewIF($expresion_logica.ex,$bloqueIf.list,lista_entonces,Ast.IF,fila,columna,false)		
	}
;


bloque_else_if returns [*arraylist.List list]
@init{ $list = arraylist.New()}
: lista += else_if+ {
                    listas := localctx.(*Bloque_else_ifContext).GetLista()
                    for _, e := range listas {
                        $list.Add(e.GetEx())
                    }
    }
;


else_if returns [Ast.Instruccion ex]
    : ELSE IF expresion_logica bloquec=bloque
    {
        fila:= $ELSE.line
		columna:= $ELSE.pos
		columna++
        lista_null := arraylist.New()
        $ex = exp_ins.NewIF($expresion_logica.ex,$bloquec.list,lista_null,Ast.ELSEIF,fila,columna,false)	
    }
;





control_if_exp returns[Ast.Instruccion ex]
	:IF expresion_logica bloqueIf = bloque_control
	{
		fila:= $IF.line
		columna:= $IF.pos
		columna++
		lista_null := arraylist.New()
		$ex = exp_ins.NewIF($expresion_logica.ex,$bloqueIf.list,lista_null,Ast.IF_EXPRESION,fila,columna,true)
	}	
	|IF expresion_logica bloqueIf = bloque_control ELSE bloqueElse = bloque_control
	  
	{
		fila:= $IF.line
		columna:= $IF.pos
		columna++
        lista_entonces := arraylist.New()
        lista_null := arraylist.New()
        Else := exp_ins.NewIF ($expresion_logica.ex,$bloqueElse.list,lista_null,Ast.ELSE_EXPRESION,fila,columna,true)
		lista_entonces.Add(Else)
		$ex = exp_ins.NewIF($expresion_logica.ex,$bloqueIf.list,lista_entonces,Ast.IF_EXPRESION,fila,columna,true)	
	}
	|IF expresion_logica bloqueIf = bloque_control bloque_else_if_exp
	{
		fila:= $IF.line
		columna:= $IF.pos
		columna++
		lista_entonces := $bloque_else_if_exp.list
		$ex = exp_ins.NewIF($expresion_logica.ex,$bloqueIf.list,lista_entonces,Ast.IF_EXPRESION,fila,columna,true)		
	}
	|IF expresion_logica bloqueIf = bloque_control bloque_else_if_exp ELSE bloqueElse = bloque_control
	{
		fila:= $IF.line
		columna:= $IF.pos
		columna++
        lista_null := arraylist.New()
        Else := exp_ins.NewIF ($expresion_logica.ex,$bloqueElse.list,lista_null,Ast.ELSE_EXPRESION,fila,columna,true)
		lista_entonces := $bloque_else_if_exp.list
        lista_entonces.Add(Else)
		$ex = exp_ins.NewIF($expresion_logica.ex,$bloqueIf.list,lista_entonces,Ast.IF_EXPRESION,fila,columna,true)		
	}
;


bloque_else_if_exp returns [*arraylist.List list]
@init{ $list = arraylist.New()}
: lista += else_if_exp+ {
                    listas := localctx.(*Bloque_else_if_expContext).GetLista()
                    for _, e := range listas {
                        $list.Add(e.GetEx())
                    }
    }
;


else_if_exp returns [Ast.Instruccion ex]
    : ELSE IF expresion_logica bloquec=bloque_control
    {
        fila:= $ELSE.line
		columna:= $ELSE.pos
		columna++
        lista_null := arraylist.New()
        $ex = exp_ins.NewIF($expresion_logica.ex,$bloquec.list,lista_null,Ast.ELSEIF_EXPRESION,fila,columna,true)	
    }
;


control_expresion returns [Ast.Instruccion ex]
    : control_if_exp {$ex = $control_if_exp.ex}
    | control_match_exp {$ex = $control_match_exp.ex}
    | control_loop_exp {$ex = $control_loop_exp.ex}
;


control_match returns[Ast.Instruccion ex]
    : MATCH expresion_logica LLAVE_IZQ control_case LLAVE_DER
    {
        fila := $MATCH.line
        columna := $MATCH.pos 
        $ex = exp_ins.NewMatch($expresion_logica.ex,$control_case.list,Ast.MATCH,fila,columna)
    }
;


control_case returns[*arraylist.List list]
@init{$list = arraylist.New()}
    : lista += cases+ 
        {
            listas := localctx.(*Control_caseContext).GetLista()
            for _, e := range listas {
                $list.Add(e.GetEx())
        }
    }
    
;



cases returns[Ast.Instruccion ex]
    : case_match CASE bloquec=bloque COMA
    {

        fila := $CASE.line
        columna := $CASE.pos
        //Verificar si lo que vienen es un default
        listaTemp := $case_match.list
        _, tipo := listaTemp.GetValue(0).(Ast.Abstracto).GetTipo()
        if tipo == Ast.DEFAULT{
            $ex = exp_ins.NewCase($case_match.list,$bloquec.list, Ast.CASE,fila,columna,true)
        }else{
            $ex = exp_ins.NewCase($case_match.list,$bloquec.list, Ast.CASE,fila,columna,false)
        }      
    }
;


case_match returns[*arraylist.List list]
@init{$list = arraylist.New()}
    : lista_cases = case_match O expresion_logica 
        {
            $lista_cases.list.Add($expresion_logica.ex)
            $list = $lista_cases.list
        }
    | expresion_logica 
        {
            $list.Add($expresion_logica.ex)
        }
    | DEFAULT
        {
            fila := $DEFAULT.line
            columna := $DEFAULT.pos
            expresion := expresiones.NewPrimitivo(false, Ast.DEFAULT, fila, columna)
            $list.Add(expresion)
        }
;





control_match_exp returns[Ast.Instruccion ex]
    : MATCH expresion_logica LLAVE_IZQ control_case_exp LLAVE_DER
    {
        fila := $MATCH.line
        columna := $MATCH.pos 
        $ex = exp_ins.NewMatch($expresion_logica.ex,$control_case_exp.list,Ast.MATCH_EXPRESION,fila,columna)
    }
;


control_case_exp returns[*arraylist.List list]
@init{$list = arraylist.New()}
    : lista += cases_exp+ 
        {
            listas := localctx.(*Control_case_expContext).GetLista()
            for _, e := range listas {
                $list.Add(e.GetEx())
        }
    }
    
;



cases_exp returns[Ast.Instruccion ex]
    : case_match_exp CASE bloquec=bloque_control COMA
    {
        fila := $CASE.line
        columna := $CASE.pos -1
        //Verificar si lo que vienen es un default
        listaTemp := $case_match_exp.list
        _, tipo := listaTemp.GetValue(0).(Ast.Abstracto).GetTipo()
        if tipo == Ast.DEFAULT{
            $ex = exp_ins.NewCase($case_match_exp.list,$bloquec.list, Ast.CASE_EXPRESION,fila,columna,true)
        }else{
            $ex = exp_ins.NewCase($case_match_exp.list,$bloquec.list, Ast.CASE_EXPRESION,fila,columna,false)
        }      
    }
;


case_match_exp returns[*arraylist.List list]
@init{$list = arraylist.New()}
    : lista_cases = case_match_exp O expresion_logica 
        {
            $lista_cases.list.Add($expresion_logica.ex)
            $list = $lista_cases.list
        }
    | expresion_logica 
        {
            $list.Add($expresion_logica.ex)
        }
    | DEFAULT
        {
            fila := $DEFAULT.line
            columna := $DEFAULT.pos
            expresion := expresiones.NewPrimitivo(false, Ast.DEFAULT, fila, columna)
            $list.Add(expresion)
        }
;


ireturn returns[Ast.Instruccion ex]
    : RETURN 
        {
            fila := $RETURN.line
            columna := $RETURN.pos
            $ex = transferencia.NewReturn(Ast.RETURN,nil,fila,columna)
        }
    | RETURN expresion_logica
        {
            fila := $RETURN.line
            columna := $RETURN.pos
            $ex = transferencia.NewReturn(Ast.RETURN_EXPRESION,$expresion_logica.ex,fila,columna)
        }
;

ibreak returns[Ast.Instruccion ex]
    : BREAK 
        {
            fila := $BREAK.line
            columna := $BREAK.pos
            $ex = transferencia.NewBreak(Ast.BREAK,nil,fila,columna)            
        }
    | BREAK expresion_logica
        {
            fila := $BREAK.line
            columna := $BREAK.pos
            $ex = transferencia.NewBreak(Ast.BREAK_EXPRESION,$expresion_logica.ex,fila,columna)                   
        }
;

icontinue returns[Ast.Instruccion ex]
    : CONTINUE 
        {
            fila := $CONTINUE.line
            columna := $CONTINUE.pos
            $ex = transferencia.NewContinue(fila,columna)                 
        }
;

control_loop returns[Ast.Instruccion ex]
    : LOOP bloque 
    {
        fila := $LOOP.line
        columna := $LOOP.pos
        $ex = bucles.NewLoop(Ast.LOOP,$bloque.list,fila,columna)
    }
;


control_loop_exp returns[Ast.Instruccion ex]
    : LOOP bloquec=bloque_control 
    {
        fila := $LOOP.line
        columna := $LOOP.pos
        $ex = bucles.NewLoop(Ast.LOOP_EXPRESION,$bloquec.list,fila,columna)
    }
;

printNormal returns[Ast.Instruccion ex]
    : PRINT PAR_IZQ expresion_logica PAR_DER
        {
            fila := $PRINT.line
            columna := $PRINT.pos
            $ex = instrucciones.NewPrint($expresion_logica.ex,Ast.PRINT,fila,columna)            
        }
;

printFormato returns[Ast.Instruccion ex]
    : PRINT PAR_IZQ CADENA COMA expresiones=elementosPrint PAR_DER
        {
            fila := $PRINT.line
            columna := $PRINT.pos
            valor := $CADENA.text
            valor = valor[1:len(valor)-1]
            $ex = instrucciones.NewPrintF($expresiones.list,valor,Ast.PRINTF,fila,columna)       
        }
;

elementosPrint returns[*arraylist.List list]
@init{$list = arraylist.New()}
    : lista_elementos = elementosPrint COMA expresion_logica 
        {
            $lista_elementos.list.Add($expresion_logica.ex)
            $list = $lista_elementos.list
        }
    | expresion_logica 
        {
            $list.Add($expresion_logica.ex)
        }
;


control_while returns[Ast.Instruccion ex]
    : WHILE expresion_logica bloque 
        {
            fila := $WHILE.line
            columna := $WHILE.pos
            $ex = bucles.NewWhile(Ast.WHILE,$expresion_logica.ex,$bloque.list,fila,columna)
        }
;


parametros_funcion returns [*arraylist.List list]
@init{$list = arraylist.New()}
    : lista_elementos = parametros_funcion COMA parametro
        {
            $lista_elementos.list.Add($parametro.ex)
            $list = $lista_elementos.list
        }
    | parametro 
        {
            $list.Add($parametro.ex)
        }
;

parametro returns [Ast.Expresion ex]
    : MUT ID DOSPUNTOS tipo_dato_tipo
        {
            fila := $MUT.line
            columna := $MUT.pos
            $ex = simbolos.NewParametro($ID.text,Ast.PARAMETRO,$tipo_dato_tipo.ex,true,false,fila,columna)

        }
    | ID DOSPUNTOS tipo_dato_tipo
        {
            fila := $ID.line
            columna := $ID.pos
            $ex = simbolos.NewParametro($ID.text,Ast.PARAMETRO,$tipo_dato_tipo.ex,false,false,fila,columna)
            
        }
    | ID DOSPUNTOS AMPERSAND tipo_dato_tipo
        {
            fila := $ID.line
            columna := $ID.pos
            $ex = simbolos.NewParametro($ID.text,Ast.PARAMETRO,$tipo_dato_tipo.ex,false,true,fila,columna)
            
        }
    /*
    | ID DOSPUNTOS AMPERSAND MUT CORCHETE_IZQ tipo_dato_tipo CORCHETE_DER
    {
        fila := $ID.line
        columna := $ID.pos
        $ex = simbolos.NewParametro($ID.text,Ast.PARAMETRO,$tipo_dato_tipo.ex,true,true,fila,columna)
    }
   
    | ID DOSPUNTOS AMPERSAND MUT VEC MENOR tipo_dato_tipo MAYOR
    {
        fila := $ID.line
        columna := $ID.pos
        $ex = simbolos.NewParametro($ID.text,Ast.PARAMETRO,$tipo_dato_tipo.ex,true,true,fila,columna)
    }
    
    | ID DOSPUNTOS VEC MENOR tipo_dato_tipo MAYOR
    {
        fila := $ID.line
        columna := $ID.pos
        $ex = simbolos.NewParametro($ID.text,Ast.PARAMETRO,$tipo_dato_tipo.ex,false,false,fila,columna)
    }
    */
    | ID DOSPUNTOS AMPERSAND MUT tipo_dato_tipo
    {
        fila := $ID.line
        columna := $ID.pos
        $ex = simbolos.NewParametro($ID.text,Ast.PARAMETRO,$tipo_dato_tipo.ex,true,true,fila,columna)
    }
;

llamada_funcion returns [Ast.Expresion ex]
    :   ID PAR_IZQ parametros_llamada PAR_DER
        {
            fila := $ID.line
            columna := $ID.pos
            id := expresiones.NewIdentificador($ID.text,Ast.IDENTIFICADOR,fila,columna)
            $ex = simbolos.NewLlamadaFuncion(id,$parametros_llamada.list,Ast.LLAMADA_FUNCION,fila,columna)
        }
    |   ID PAR_IZQ PAR_DER
        {
            fila := $ID.line
            columna := $ID.pos
            params := arraylist.New()
            id := expresiones.NewIdentificador($ID.text,Ast.IDENTIFICADOR,fila,columna)        
            $ex = simbolos.NewLlamadaFuncion(id,params,Ast.LLAMADA_FUNCION,fila,columna)   
        }
;

parametros_llamada returns [*arraylist.List list]
@init{$list = arraylist.New()}
:   lista_elementos = parametros_llamada COMA parametro_llamada_referencia
        {
            $lista_elementos.list.Add($parametro_llamada_referencia.ex)
            $list = $lista_elementos.list
        }
|   parametro_llamada_referencia
        {
            $list.Add($parametro_llamada_referencia.ex)
        }
;

parametro_llamada_referencia returns [Ast.Expresion ex]
    :   e = expresion_logica
    {
        temp := localctx.(*Parametro_llamada_referenciaContext).GetE().GetEx()
        fila := temp.(Ast.Abstracto).GetFila()
        columna := temp.(Ast.Abstracto).GetColumna()
        $ex = simbolos.NewValor($e.ex, Ast.VALOR , false, false, fila, columna)
    }
    |   AMPERSAND MUT id=expresion_logica
    {
        fila := $AMPERSAND.line
        columna := $AMPERSAND.pos
        
        $ex = simbolos.NewValor($id.ex, Ast.VALOR , true, true, fila, columna)
    }
    //id := expresiones.NewIdentificador($ID.text,Ast.IDENTIFICADOR,fila,columna)
    |   AMPERSAND id=expresion_logica
    {
        fila := $AMPERSAND.line
        columna := $AMPERSAND.pos
        
        $ex = simbolos.NewValor($id.ex, Ast.VALOR , true, false, fila, columna)        
    }
    //id := expresiones.NewIdentificador($ID.text,Ast.IDENTIFICADOR,fila,columna)
;


elementos_vector returns[*arraylist.List list]
@init{$list = arraylist.New()}
:   lista_elementos = elementos_vector COMA expresion_logica
        {
            $lista_elementos.list.Add($expresion_logica.ex)
            $list = $lista_elementos.list
        }
|   expresion_logica
        {
            $list.Add($expresion_logica.ex)
        }
;


metodos_iniciar_vector returns[Ast.Expresion ex]
    : VEC DOBLE_DOSPUNTOS NEW PAR_IZQ PAR_DER
        {
            fila := $VEC.line
            columna := $VEC.pos 
            $ex = fn_vectores.NewVecNew(fila,columna)           
        }
    | VEC_M NOT CORCHETE_IZQ e=elementos_vector CORCHETE_DER
        {
            fila := $VEC_M.line
            columna := $VEC_M.pos 
            //listaTemp := localctx.(*Metodos_iniciar_vectorContext).GetE().GetList()
            $ex = fn_vectores.NewVecElementos($elementos_vector.list,fila,columna)            
        }
    | VEC_M NOT CORCHETE_IZQ ex1=expresion_logica PUNTOCOMA ex2=expresion_logica CORCHETE_DER
        {
            fila := $VEC_M.line
            columna := $VEC_M.pos 
            listaTemp := arraylist.New()
            listaTemp.Add($ex1.ex)
            listaTemp.Add($ex2.ex)
            $ex = fn_vectores.NewVecFactorial(listaTemp,fila,columna)            
        }
    | VEC DOBLE_DOSPUNTOS WITH_CAPACITY PAR_IZQ capacity=expresion_logica PAR_DER
        {
            fila := $VEC.line
            columna := $VEC.pos 
            $ex = fn_vectores.NewVecWithCapacity($capacity.ex,fila,columna)                   
        }
;

metodos_vector returns[Ast.Instruccion ex]
    : id=expresion_logica PUNTO PUSH PAR_IZQ exp=expresion_logica PAR_DER
        {
            fila := $PUNTO.line
            columna := $PUNTO.pos
            $ex = fn_vectores.NewPush($id.ex,$exp.ex,Ast.VEC_PUSH,fila,columna)
        }
    | id=expresion_logica PUNTO INSERT PAR_IZQ pos=expresion_logica COMA exp=expresion_logica PAR_DER
        {
            fila := $PUNTO.line
            columna := $PUNTO.pos
            $ex = fn_vectores.NewInsertVec($id.ex,$exp.ex,$pos.ex,Ast.VEC_INSERT,fila,columna)            
        }
;

potencia returns[Ast.Expresion ex]
    : I64 DOBLE_DOSPUNTOS POW PAR_IZQ val=expresion_logica COMA pot=expresion_logica PAR_DER 
    {
        fila := $DOBLE_DOSPUNTOS.line
        columna := $DOBLE_DOSPUNTOS.pos-1        
        $ex = expresiones.NewPow(Ast.POW,Ast.I64,$val.ex,$pot.ex,fila,columna)
    }
| F64 DOBLE_DOSPUNTOS POWF PAR_IZQ val=expresion_logica COMA pot=expresion_logica PAR_DER 
    {
        fila := $DOBLE_DOSPUNTOS.line
        columna := $DOBLE_DOSPUNTOS.pos-1        
        $ex = expresiones.NewPow(Ast.POW,Ast.F64,$val.ex,$pot.ex,fila,columna)
    }
;


array returns[Ast.Expresion ex]
    : CORCHETE_IZQ elementos=elementos_vector CORCHETE_DER
        {
            fila := $CORCHETE_IZQ.line
            columna := $CORCHETE_IZQ.pos 
            $ex = fn_array.NewArrayElementos($elementos_vector.list,fila,columna)                
        }
    | CORCHETE_IZQ elemento=expresion_logica PUNTOCOMA serie=expresion_logica CORCHETE_DER
        {
            fila := $CORCHETE_IZQ.line
            columna := $CORCHETE_IZQ.pos
            listaTemp := arraylist.New()
            listaTemp.Add($elemento.ex)
            listaTemp.Add($serie.ex)
            $ex = fn_array.NewArrayFactorial(listaTemp,fila,columna)
        }
;

dimension_array returns[Ast.Expresion ex]
    : CORCHETE_IZQ lista_elementos = dimension_array PUNTOCOMA expresion_logica CORCHETE_DER
        {
            dimension := localctx.(*Dimension_arrayContext).GetLista_elementos().GetEx()
            dimension.(expresiones.DimensionArray).Elementos.Add($expresion_logica.ex)
            $ex = dimension
        }
    | CORCHETE_IZQ tipo=tipo_dato_tipo PUNTOCOMA expresion_logica CORCHETE_DER
        {
            fila := $CORCHETE_IZQ.line
            columna := $CORCHETE_IZQ.pos
            listaD := arraylist.New()
            listaD.Add($expresion_logica.ex)
            $ex = expresiones.NewDimensionArray(listaD, $tipo.ex,fila,columna)
        }

;

dimension_acceso_array returns[*arraylist.List list]
@init{$list = arraylist.New()}
    :   lista_elementos = dimension_acceso_array CORCHETE_IZQ expresion_logica CORCHETE_DER
             {
                $lista_elementos.list.Add($expresion_logica.ex)
                $list = $lista_elementos.list
            }
    |   CORCHETE_IZQ ex1=expresion_logica CORCHETE_DER /*CORCHETE_IZQ ex2=expresion CORCHETE_DER*/
            {
                $list.Add($ex1.ex)
            }
            //$list.Add($ex2.ex)
;


tipo_dato_tipo returns[Ast.TipoRetornado ex]
    :   tipo_dato
        {
            $ex = Ast.TipoRetornado{
                Valor: true,
                Tipo: $tipo_dato.ex,
            }
        }
    |   VEC MENOR tipo=tipo_dato_tipo MAYOR
        {
            $ex = Ast.TipoRetornado{
                Valor: $tipo.ex,
                Tipo: Ast.VECTOR,
            }
        }
    |   ID_CAMEL
        {   
            $ex = Ast.TipoRetornado{
                Valor: $ID_CAMEL.text,
                Tipo: Ast.STRUCT,
            }
        }
    |   dimension_array
        {
            $ex = Ast.TipoRetornado{
            Valor: $dimension_array.ex,
            Tipo: Ast.DIMENSION_ARRAY,
            }   

        }
    |   acceso_modulo
        {
            $ex = Ast.TipoRetornado{
                Valor: $acceso_modulo.ex,
                Tipo: Ast.ACCESO_MODULO,
            }
        }

;

acceso_modulo returns[Ast.Expresion ex]
    : acceso = acceso_modulo_elementos
        {
            lista := localctx.(*Acceso_moduloContext).GetAcceso().GetList()
            elemento := lista.GetValue(lista.Len()-1)
            fila:= elemento.(Ast.Abstracto).GetFila()
            columna:= elemento.(Ast.Abstracto).GetColumna()
            $ex = simbolos.NewAccesoModulo($acceso.list,fila,columna)
        }
;

acceso_modulo_elementos returns[*arraylist.List list]
@init{$list = arraylist.New()}
    :  lista_elementos=acceso_modulo_elementos id=acceso_modulo_elemento_final
        {
            $lista_elementos.list.Add($id.ex)
            $list = $lista_elementos.list          
        }
    |   ID
        {
            fila:= $ID.line
            columna:= $ID.pos      
            id := expresiones.NewIdentificador($ID.text,Ast.IDENTIFICADOR,fila,columna)   
            $list.Add(id)
        }
;


acceso_modulo_elemento_final returns [Ast.Expresion ex]
    : DOBLE_DOSPUNTOS ID_CAMEL
        {
            fila:= $ID_CAMEL.line
            columna:= $ID_CAMEL.pos-1       
            $ex = expresiones.NewIdentificador($ID_CAMEL.text,Ast.IDENTIFICADOR,fila,columna)  

        }
    | DOBLE_DOSPUNTOS ID
        {
            fila:= $ID.line
            columna:= $ID.pos-1       
            $ex = expresiones.NewIdentificador($ID.text,Ast.IDENTIFICADOR,fila,columna)  

        }
    | DOBLE_DOSPUNTOS llamada_funcion
        {
            $ex = $llamada_funcion.ex
        }
;


control_for returns [Ast.Instruccion ex]
    :   FOR ID IN rango_for bloque
        {
            fila:= $FOR.line
            columna:= $FOR.pos 
            id := expresiones.NewIdentificador($ID.text,Ast.IDENTIFICADOR,fila,columna)
            $ex = bucles.NewFor(id,$rango_for.ex,$bloque.list,fila,columna)      
        }
        //Los fors vacios sin instrucciones
    |   FOR ID IN rango_for LLAVE_IZQ LLAVE_DER
        {
            fila:= $FOR.line
            columna:= $FOR.pos 
            id := expresiones.NewIdentificador($ID.text,Ast.IDENTIFICADOR,fila,columna)  
            listaVacia := arraylist.New()
            $ex = bucles.NewFor(id,$rango_for.ex,listaVacia,fila,columna) 
        }
;



rango_for returns [Ast.Expresion ex]
    :   ex1=expresion_logica RANGO ex2=expresion_logica
        {
            elemento := localctx.(*Rango_forContext).GetEx1().GetEx()
            fila:= elemento.(Ast.Abstracto).GetFila()
            columna:= elemento.(Ast.Abstracto).GetColumna()  
            $ex = bucles.NewRange(Ast.RANGE_RANGO,$ex1.ex,$ex2.ex,fila,columna)       
        }
    |   ex3=expresion_logica
    {
            elemento := localctx.(*Rango_forContext).GetEx3().GetEx()
            fila:= elemento.(Ast.Abstracto).GetFila()
            columna:= elemento.(Ast.Abstracto).GetColumna()    
            expresionVacia := expresiones.NewPrimitivo(nil,Ast.NULL,fila,columna) 
            $ex = bucles.NewRange(Ast.RANGE_EXPRESION,$expresion_logica.ex,expresionVacia,fila,columna)      
    }
;



/* 
fn_chars returns [Ast.Expresion ex]
    :   expresion PUNTO CHARS PAR_IZQ PAR_DER
    {
            elemento := localctx.(*Fn_charsContext).GetExpresion().GetEx()
            fila := elemento.(Ast.Abstracto).GetFila()
            columna := elemento.(Ast.Abstracto).GetColumna()
            $ex = fn_primitivas.NewToChars($expresion.ex,fila,columna)
    }
;


acceso_atributo_struct returns[Ast.Expresion ex]
    :
     
    |   estructura=ID PUNTO atributo=ID 
        {
            filaS := $estructura.line
            columnaS := $estructura.pos-1
            filaA := $atributo.line
            columnaA := $atributo.pos-1
            idStruct := expresiones.NewIdentificador($estructura.text,Ast.IDENTIFICADOR,filaS,columnaS)
            idAtributo := expresiones.NewIdentificador($atributo.text,Ast.IDENTIFICADOR,filaA,columnaA)
            $ex = simbolos.NewAccesoStruct(idStruct,idAtributo,filaS,columnaS)
        }
;
*/
