grammar ylang;

program: (statement)*;

NUMBER: '-'? [0-9]+;
STRING: '"' ~["]* '"';
BOOLEAN: 'true' | 'false';
ID: [a-zA-Z][a-zA-Z0-9]*;

TYPE: 'number' | 'string' | 'bool' | ID;
literal: NUMBER | STRING | BOOLEAN;

classStatement:
	'class' className = ID '{' fieldDeclaration* methodDeclaration* '}';

fieldDeclaration: ID ':' TYPE ';';
methodDeclaration: TYPE ID '(' parameterList? ')' block;

parameterList: parameter (',' parameter)*;
parameter: TYPE ID;

block: '{' statement* '}';
statement:
	expressionStatement
	| variableDeclaration
	| ifStatement
	| whileStatement
	| returnStatement
	| classStatement
	| streamStatement;

expressionStatement: expression ';';
variableDeclaration: TYPE ID '=' expression ';';
ifStatement: 'if' '(' expression ')' block ('else' block)?;
whileStatement: 'while' '(' expression ')' block;
returnStatement: 'return' expression? ';';

expression: literal | invokeMethod | variable | assignExpr;
invokeMethod: variable '.' ID '(' argumentList? ')';
argumentList: expression (',' expression)*;
variable: ID;

assignExpr: ID '=' literal;
objectExpr: ID '{' (assignExpr (',' assignExpr)*)? '}';
lambdaExpr: '(' parameterList? ')' '=>' block;
streamFuncCallExpr: (ID ('.' ID)*)+;
streamExpr: objectExpr | streamFuncCallExpr | lambdaExpr;
streamStatement: (streamExpr ('>>' streamExpr)*)+;

WS: [ \t\n\r]+ -> skip;