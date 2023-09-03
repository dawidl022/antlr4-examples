lexer grammar Lexer;

ID: [a-zA-Z][0-9a-zA-Z]*;
INT: [0-9]+;
NEWLINE: '\r'? '\n';

WS: [ \t]+ -> skip;
