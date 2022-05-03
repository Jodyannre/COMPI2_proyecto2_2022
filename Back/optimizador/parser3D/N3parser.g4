parser grammar N3parser;

options{
    tokenVocab = N3lexer;
    language = Go;
}


@header{
    import "github.com/colegno/arraylist"
    import "Back/optimizador/elementos/bloque3d"
    import "Back/optimizador/elementos/control"
    import "Back/optimizador/elementos/expresiones3d"
    import "Back/optimizador/elementos/funciones3d"
    import "Back/optimizador/elementos/headers3d"
    import "Back/optimizador/elementos/instrucciones3d"
    import "Back/optimizador/elementos/interfaces3d"
    import "strings"
}

/* 
@members{
}
*/


inicio returns[bloque3d.Bloque3dGlobal ex] 
    :   include declaraciones funciones
            {
                $ex = bloque3d.NewBloque3dGlobal($declaraciones.list, 
                $funciones.list , $include.ex)
            }
;


funciones returns [*arraylist.List list]
@init{ $list = arraylist.New()}
:   lista += funcion+ 
        {
            listas := localctx.(*FuncionesContext).GetLista()
            for _, e := range listas {
                $list.Add(e.GetEx())
            }
        }
;

funcion returns [funciones3d.Funcion3D ex]
    :   INT MAIN PAR_IZQ PAR_DER LLAVE_IZQ bloques 
        {
            $ex = funciones3d.NewFuncion3D($MAIN.text,$bloques.list,interfaces3d.MAIN3D)
        }
    |   VOID ID PAR_IZQ PAR_DER LLAVE_IZQ bloques 
        {
            $ex = funciones3d.NewFuncion3D($ID.text,$bloques.list,interfaces3d.FUNCION3D)
        }
;


llamada returns [instrucciones3d.Llamada3D ex]
    :   ID PAR_IZQ PAR_DER
        {
            $ex = instrucciones3d.NewLlamada3D($ID.text + "()")
        }
;


bloques returns [*arraylist.List list]
@init{ $list = arraylist.New()}
:   lista += bloque+ 
        {
            listas := localctx.(*BloquesContext).GetLista()
            for _, e := range listas {
                $list.Add(e.GetEx())
            }
        }
;


bloque returns[bloque3d.Bloque3d ex]
    :   init=bloque_i instrucciones fin=bloque_f
        {
            nLista := arraylist.New()
            if $init.ex != nil{
                nLista.Add($init.ex)
            }
            listas := $instrucciones.list
            for i:=0; i< listas.Len();i++ {
                nLista.Add(listas.GetValue(i))
            }
            if $fin.ex != nil{
                nLista.Add($fin.ex)
            }
            $ex = bloque3d.NewBloque3d(nLista)
        }
    |   instrucciones
        {
            nLista := arraylist.New()
            listas := $instrucciones.list
            for i:=0; i< listas.Len();i++ {
                nLista.Add(listas.GetValue(i))
            } 
            $ex = bloque3d.NewBloque3d(nLista)       
        }
    |   init=bloque_i  fin=bloque_f
        {
            nLista := arraylist.New()
            if $init.ex != nil{
                nLista.Add($init.ex)
            }
            if $fin.ex != nil{
                nLista.Add($fin.ex)
            }
            $ex = bloque3d.NewBloque3d(nLista)
        }
;


bloque_i returns[interfaces3d.Expresion3D ex]
    :   asignacion PUNTOCOMA
        {
            $ex = $asignacion.ex
        }
    |   etiqueta
        {
            $ex = $etiqueta.ex
        }
    |   print3d PUNTOCOMA
        {
            $ex = $print3d.ex
        }
    |   if3d
        {
            $ex = $if3d.ex
        }
    |   salto PUNTOCOMA
        {
            $ex = $salto.ex
        }
;


instrucciones returns [*arraylist.List list]
@init{ $list = arraylist.New()}
:   lista += instruccion+ 
        {
            listas := localctx.(*InstruccionesContext).GetLista()
            for _, e := range listas {
                $list.Add(e.GetEx())
            }
        }
;



instruccion returns[interfaces3d.Expresion3D ex]
    :   asignacion PUNTOCOMA
        {
            $ex = $asignacion.ex
        }
    |   retorno PUNTOCOMA
        {
            $ex = $retorno.ex
        }
    |   print3d PUNTOCOMA
        {   
            $ex = $print3d.ex
        }
    |   llamada PUNTOCOMA
        {
            $ex = $llamada.ex
        }
;


bloque_f returns[interfaces3d.Expresion3D ex]
    :   etiqueta
        {
            $ex = $etiqueta.ex
        }
    |   salto PUNTOCOMA
        {
            $ex = $salto.ex
        }
    |   if3d
        {
            $ex = $if3d.ex
        }
    |   LLAVE_DER
        {
            $ex = expresiones3d.NewPrimitivo3D("",interfaces3d.PUNTERO_STACK)
        }
;


print3d returns[instrucciones3d.Print3D ex]
    :   PRINTF PAR_IZQ CADENA COMA PAR_IZQ INT PAR_DER TEMPORAL PAR_DER
        {
            cad := $CADENA.text
            if strings.Contains(cad,"c"){
                cad = "%%c"
            }else if strings.Contains(cad,"d") {
                cad = "%%d"
            }else{
                cad = "%%f"
            }
            valor := "printf(\""+cad+"\",(int)"+ $TEMPORAL.text +")"
            $ex = instrucciones3d.NewPrint3D(valor)
        }
    |   PRINTF PAR_IZQ CADENA COMA TEMPORAL PAR_DER
        {
            cad := $CADENA.text
            if strings.Contains(cad,"c"){
                cad = "%%c"
            }else if strings.Contains(cad,"d") {
                cad = "%%d"
            }else{
                cad = "%%f"
            }
            valor := "printf(\""+cad+"\","+ $TEMPORAL.text +")"
            $ex = instrucciones3d.NewPrint3D(valor)
        }
    |   PRINTF PAR_IZQ CADENA COMA NUMERO PAR_DER
        {
            cad := $CADENA.text
            if strings.Contains(cad,"c"){
                cad = "%%c"
            }else if strings.Contains(cad,"d") {
                cad = "%%d"
            }else{
                cad = "%%f"
            }
            valor := "printf(\""+cad+"\",(int)"+ $NUMERO.text +")"
            $ex = instrucciones3d.NewPrint3D(valor)
        }
    |   PRINTF PAR_IZQ CADENA COMA PAR_IZQ INT PAR_DER NUMERO PAR_DER
        {
            cad := $CADENA.text
            if strings.Contains(cad,"c"){
                cad = "%%c"
            }else if strings.Contains(cad,"d") {
                cad = "%%d"
            }else{
                cad = "%%f"
            }
            valor := "printf(\""+cad+"\",(int)"+ $NUMERO.text +")"
            $ex = instrucciones3d.NewPrint3D(valor)
        }
    |   PRINTF PAR_IZQ CADENA COMA PAR_IZQ INT PAR_DER DECIMAL PAR_DER
        {
            cad := $CADENA.text
            if strings.Contains(cad,"c"){
                cad = "%%c"
            }else if strings.Contains(cad,"d") {
                cad = "%%d"
            }else{
                cad = "%%f"
            }
            valor := "printf(\""+cad+"\",(int)"+ $DECIMAL.text +")"
            $ex = instrucciones3d.NewPrint3D(valor)
        }
;



if3d returns[control.IF3D ex]
    :   IF PAR_IZQ operacion PAR_DER go1=salto PUNTOCOMA go2=salto PUNTOCOMA etiqueta
        {
            $ex = control.NewIF3D($operacion.ex, $go1.ex, $go2.ex, $etiqueta.ex)
        }
;


asignacion returns[instrucciones3d.Asignacion3D ex]
    :   exp=expresion IGUAL operacion
        {
           $ex = instrucciones3d.NewAsignacion3D($exp.ex, $operacion.ex) 
        }
    |   exp=expresion IGUAL exp2=expresion
        {
            expr := expresiones3d.NewPrimitivo3D("",interfaces3d.NUMERO3D)
            op := expresiones3d.NewOperacion3D($exp2.ex, expr, "", true) 
            $ex = instrucciones3d.NewAsignacion3D($exp.ex, op) 
        }   
;



operacion returns[expresiones3d.Operacion3D ex]
    :   opIzq=expresion operador opDer=expresion 
        {
           $ex = expresiones3d.NewOperacion3D($opIzq.ex, $opDer.ex, $operador.ex, false) 
        }
    |   PAR_IZQ INT PAR_DER opIzq=expresion operador PAR_IZQ INT PAR_DER opDer=expresion 
        {
            expDer := $opDer.ex
            expDer.Valor = "(int)" + expDer.Valor

            expIzq := $opIzq.ex
            expIzq.Valor = "(int)" + expIzq.Valor
           $ex = expresiones3d.NewOperacion3D(expIzq, expDer, $operador.ex, false) 
        }
    |   PAR_IZQ INT PAR_DER opIzq=expresion operador opDer=expresion 
        {
            expIzq := $opIzq.ex
            expIzq.Valor = "(int)" + expIzq.Valor
           $ex = expresiones3d.NewOperacion3D(expIzq, $opDer.ex, $operador.ex, false) 
        }
    |   opIzq=expresion operador PAR_IZQ INT PAR_DER opDer=expresion 
        {
            expDer := $opDer.ex
            expDer.Valor = "(int)" + expDer.Valor
           $ex = expresiones3d.NewOperacion3D($opIzq.ex, expDer, $operador.ex, false) 
        }
;


expresion returns[expresiones3d.Primitivo3D ex]
    :   P
            {
                $ex = expresiones3d.NewPrimitivo3D($P.text,interfaces3d.PUNTERO_STACK)
            }
    |   TEMPORAL
            {
                $ex = expresiones3d.NewPrimitivo3D($TEMPORAL.text,interfaces3d.PUNTERO_STACK)
            }    
    |   H
            {
                $ex = expresiones3d.NewPrimitivo3D($H.text,interfaces3d.PUNTERO_HEAP)
            }
    |   RESTA NUMERO
            {
                $ex = expresiones3d.NewPrimitivo3D("-"+$NUMERO.text,interfaces3d.NUMERO3D)
            }    
    |   NUMERO 
            {
                $ex = expresiones3d.NewPrimitivo3D($NUMERO.text,interfaces3d.NUMERO3D)
            }
    |   DECIMAL
            {
                $ex = expresiones3d.NewPrimitivo3D($DECIMAL.text,interfaces3d.DECIMAL3D)
            }
    |   HEAP CORCHETE_IZQ PAR_IZQ INT PAR_DER TEMPORAL CORCHETE_DER
            {
                valor := "heap [(int)" + $TEMPORAL.text + "]"
                $ex = expresiones3d.NewPrimitivo3D(valor,interfaces3d.ACCESO_HEAP)
            }
    |   HEAP CORCHETE_IZQ PAR_IZQ INT PAR_DER H CORCHETE_DER
            {
                valor := "heap [(int)" + $H.text + "]"
                $ex = expresiones3d.NewPrimitivo3D(valor,interfaces3d.ACCESO_HEAP)
            }
    |   STACK CORCHETE_IZQ PAR_IZQ INT PAR_DER TEMPORAL CORCHETE_DER
            {
                valor := "stack [(int)" + $TEMPORAL.text + "]"
                $ex = expresiones3d.NewPrimitivo3D(valor,interfaces3d.ACCESO_STACK)
            }
;


etiqueta returns[instrucciones3d.Salto3D ex]
    :   LABEL DOSPUNTOS
        {
            $ex = instrucciones3d.NewSalto3D($LABEL.text)
        }
;

salto returns[instrucciones3d.Goto3D ex]
    :   GOTO LABEL
        {
            $ex = instrucciones3d.NewGoto3D("goto "+$LABEL.text)
        }
;

retorno returns[instrucciones3d.Return3D ex]
    :   RETURN
        {
            $ex = instrucciones3d.NewReturn3D($RETURN.text)
        }
    |   RETURN NUMERO
        {
            $ex = instrucciones3d.NewReturn3D($RETURN.text + " " + $NUMERO.text)
        }
;

include returns[headers3d.Include3D ex]
    :   INCLUDE STDIO
        {
            $ex = headers3d.NewInclude3D($INCLUDE.text + " " + $STDIO.text)
        }
;

declaraciones returns [*arraylist.List list]
@init{ $list = arraylist.New()}
:   lista += declaracion+ 
        {
            listas := localctx.(*DeclaracionesContext).GetLista()
            for _, e := range listas {
                $list.Add(e.GetEx())
            }
        }
;

declaracion returns[headers3d.Declaracion3D ex]
    :   FLOAT STACK CORCHETE_IZQ NUMERO CORCHETE_DER PUNTOCOMA
        {
            $ex = headers3d.NewDeclaracion3D("float stack["+$NUMERO.text+"]")
        }
    |   FLOAT HEAP CORCHETE_IZQ NUMERO CORCHETE_DER PUNTOCOMA
        {
            $ex = headers3d.NewDeclaracion3D("float heap["+$NUMERO.text+"]")
        }
    |   FLOAT P PUNTOCOMA
        {
            $ex = headers3d.NewDeclaracion3D("float P")
        }
    |   FLOAT H PUNTOCOMA
        {
            $ex = headers3d.NewDeclaracion3D("float H")
        }
    |   FLOAT temporalesTexto PUNTOCOMA
        {
            $ex = headers3d.NewDeclaracion3D("float "+$temporalesTexto.ex)
        }
;

temporalesTexto returns [string ex]
    :   temporalesLista
        {
            temporales := ""
            lista := $temporalesLista.list
            for i:=0; i< lista.Len();i++ {
                if i != lista.Len()-1{ 
                    temporales += lista.GetValue(i).(string) + ","
                }else{
                    temporales += lista.GetValue(i).(string)
                }
                
            } 
            $ex = temporales
        }
;


temporalesLista returns [*arraylist.List list]
@init{$list = arraylist.New()}
    : lista_elementos = temporalesLista COMA temp=TEMPORAL 
        {
            $lista_elementos.list.Add($temp.text)
            $list = $lista_elementos.list
        }
    | temp=TEMPORAL
        {
            $list.Add($temp.text)
        }
;

operador returns[string ex]
    :   MAYOR_I 
        {
            $ex = $MAYOR_I.text
        }         
    |   MAYOR 
        {
            $ex = $MAYOR.text
        }           
    |   MENOR_I 
        {
            $ex = $MENOR_I.text
        }         
    |   MENOR 
        {
            $ex = $MENOR.text
        }            
    |   IGUALDAD 
        {
            $ex = $IGUALDAD.text
        }        
    |   DISTINTO 
        {
            $ex = $DISTINTO.text
        }       
    |   MODULO 
        {
            $ex = "%%"
        }          
    |   MULTIPLICACION 
        {
            $ex = $MULTIPLICACION.text
        }  
    |   DIVISION 
        {
            $ex = $DIVISION.text
        }        
    |   RESTA 
        {
            $ex = $RESTA.text
        }           
    |   SUMA 
        {
            $ex = $SUMA.text
        }            
;


















