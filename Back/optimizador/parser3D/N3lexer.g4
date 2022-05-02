lexer grammar N3lexer;

options {
    language = Go;
}




/* *********************Tipos primitivos********************* */
INT         : 'int';
FLOAT       : 'float';



/* *******************Palabraws Reservadas******************* */

MAIN            : 'main';
IF              : 'if';
RETURN          : 'return';
VOID            : 'void';
PRINTF          : 'printf';
GOTO            : 'goto';
HEAP            : 'heap';
STACK           : 'stack';
P               : 'P';
H               : 'H';
INCLUDE         : '#include';
STDIO           : '<stdio.h>'; 



NUMERO          : [0-9]+;
DECIMAL         : [0-9]+'.'[0-9]+;
TEMPORAL        : 'T' [0-9]+;
LABEL           : 'L' [0-9]+;
ID_CAMEL        : [A-Z]([a-zA-Z]|[0-9])*; 
ID              : ([a-zA-Z]|[0-9])([a-zA-Z]|[0-9]|'_')*;
DOSPUNTOS       : ':';
PUNTO           : '.';
COMA            : ',';
PUNTOCOMA       : ';';
MAYOR_I         : '>=';
MAYOR           : '>';
MENOR_I         : '<=';
MENOR           : '<';
IGUALDAD        : '==';
CASE            : '=>';
DISTINTO        : '!=';
IGUAL           : '=';
MODULO          : '%';
MULTIPLICACION  : '*' ;
DIVISION        : '/' ;
RESTA           : '-' ;
SUMA            : '+' ;
NOT             : '!';
PAR_IZQ         : '(';
PAR_DER         : ')';
LLAVE_IZQ       : '{';
LLAVE_DER       : '}';
CORCHETE_IZQ    : '[';
CORCHETE_DER    : ']'; 
CADENA          : '"' .*? '"';
//PRINTC          : '"' '%%c' '"';
//PRINTD          : '"' '%%d' '"';
//ASCII           : [#$%&()*+,/:;<>=?@^_{}!-]|CORCHETE_DER|CORCHETE_IZQ;
//CARACTER        : '\'' (ASCII|[0-9]|[a-zA-Z]) '\'';
WHITESPACE      : [ \\\r\n\t]+ -> skip;
COMMENT         : '/*' .*? '*/' -> skip;
LINE_COMMENT    : '//' ~[\r\n]* -> skip;




