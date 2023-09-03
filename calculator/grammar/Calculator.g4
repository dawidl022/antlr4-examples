grammar Calculator;
import Lexer;

init: (statement NEWLINE)*;

statement:
	'clear'			# clear
	| expr			# printExpr
	| assignment	# assign;

assignment: ID '=' expr;

expr:
	expr op = (MUL | DIV) expr		# MulDiv
	| expr op = (ADD | SUB) expr	# AddSub
	| '(' expr ')'					# parens
	| ID							# id
	| INT							# int;

MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
