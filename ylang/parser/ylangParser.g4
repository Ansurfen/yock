parser grammar ylangParser;

options {
	tokenVocab = ylangLexer;
}

primitiveType: BOOL | NUMBER | STRING | ANY;

typeType: primitiveType;

typeList: typeType (',' typeType)*;

identifier: IDENTIFIER;

arrayDeclarator: '[' ']';

statement: FOR '(' ')' statement;

classOrInterfaceModifier: PUBLIC | PRIVATE | STATIC;

classDeclaration:
	CLASS identifier (EXTENDS typeType)? (IMPLEMENTS typeType)? classBody;

classBody: '{' classBodyDeclaration* '}';

classBodyDeclaration:
	';'
	| STATIC? block
	| modifier* memberDeclaration;

modifier: classOrInterfaceModifier;

memberDeclaration:
	methodDeclaration
	| fieldDeclaration
	| constructorDeclaration
	| destoryDeclaration;

methodDeclaration:
	FN identifier formalParameters returnParameters? methodBody;

returnParameters:
	typeTypeOrVoid
	| '(' typeTypeOrVoid (',' typeTypeOrVoid)* ')';

fieldDeclaration: identifier ':' typeType ';';

constructorDeclaration:;

destoryDeclaration:;

methodBody: block | ';';

formalParameters:
	'(' receiverParameter?
	| receiverParameter (',') ')';

receiverParameter:;

lastFormalParameter:;

typeTypeOrVoid: typeType | VOID;

block:;