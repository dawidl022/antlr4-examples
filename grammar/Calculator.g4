grammar Calculator;

init: (statement NEWLINE)*;

statement: expr | assignment;

assignment: ID '=' expr;

expr: expr mulOp expr | expr addOp expr | '(' expr ')' | ID | INT;

mulOp: '*' | '/';
addOp: '+' | '-';

ID: [a-zA-Z][0-9a-zA-Z]*;
INT: [0-9]+;
NEWLINE: '\r'? '\n';

WS: [ \t]+ -> skip;
