lexer grammar Nlexer;

options {
    language = Go;
}




/* *********************Tipos primitivos********************* */
BOOL        : 'bool';
CHAR        : 'char';
F64         : 'f64';
I64         : 'i64';
STR         : '&str';
STRING      : 'String';
USIZE       : 'usize';



/* *******************Palabraws Reservadas******************* */

MAIN            : 'main';
POWF            : 'powf';
POW             : 'pow';
AS              : 'as';
VEC             : 'Vec';
VEC_M           : 'vec';
MUT             : 'mut';
LET             : 'let';
STRUCT          : 'struct';
TO_STRING       : 'to_string';
TO_OWNED        : 'to_owned';
TRUE            : 'true';
FALSE           : 'false';
PRINT           : 'println!';
FN              : 'fn';
ABS             : 'abs';
SQRT            : 'sqrt';
CLONE           : 'clone';
CHARS           : 'chars';
NEW             : 'new';
LEN             : 'len';
PUSH            : 'push';
REMOVE          : 'remove';
CONTAINS        : 'contains';
INSERT          : 'insert';
CAPACITY        : 'capacity';
WITH_CAPACITY   : 'with_capacity';
IF              : 'if';
ELSE            : 'else';
MATCH           : 'match';
LOOP            : 'loop';
WHILE           : 'while';
FOR             : 'for';
IN              : 'in';
RETURN          : 'return';
BREAK           : 'break';
CONTINUE        : 'continue';
MOD             : 'mod';
PUB             : 'pub';




NUMERO          : [0-9]+;
DECIMAL         : [0-9]+'.'[0-9]+;
ID_CAMEL        : [A-Z]([a-zA-Z]|[0-9])*; //Para struct por ejemplo
ID              : ([a-zA-Z]|[0-9])([a-zA-Z]|[0-9]|'_')*;
DEFAULT         : '_';
O               : '|';
OR              : '||';
AMPERSAND       : '&';
AND             : '&&';
PRINT_OP_DEBUG  : ':?';
DOBLE_DOSPUNTOS : '::';
DOSPUNTOS       : ':';
RANGO           : '..';
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
FN_TIPO_RETORNO : '->';
RESTA           : '-' ;
SUMA            : '+' ;
NOT             : '!';
PREGUNTA        : '?';
PAR_IZQ         : '(';
PAR_DER         : ')';
LLAVE_IZQ       : '{';
LLAVE_DER       : '}';
CORCHETE_IZQ    : '[';
CORCHETE_DER    : ']'; 
CADENA          : '"' .*? '"';
ASCII           : [#$%&()*+,/:;<>=?@^_{}!-]|CORCHETE_DER|CORCHETE_IZQ;
CARACTER        : '\'' (ASCII|[0-9]|[a-zA-Z]) '\'';
WHITESPACE      : [ \\\r\n\t]+ -> skip;
COMMENT         : '/*' .*? '*/' -> skip;
LINE_COMMENT    : '//' ~[\r\n]* -> skip;




