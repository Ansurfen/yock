lexer grammar ylangLexer;

TYPE: 'type';
CLASS: 'class';
ENUM: 'enum';
FN: 'fn';
RETURN: 'return';
NEW: 'new';
DELETE: 'delete';
CONSTRUCT: 'construct';
DESOTRY: '~construct';
IMPORT: 'import';
AS: 'as';
PUBLIC: 'public';
PRIVATE: 'private';
THIS: 'this';
BREAK: 'break';
CONTINUE: 'continue';
TRY: 'try';
CATCH: 'catch';
THROW: 'thorw';
GOTO: 'goto';
FOR: 'for';
IF: 'if';
ELSE: 'else';
ELIF: 'elif';
DO: 'do';
WHILE: 'while';
SWITCH: 'switch';
CASE: 'case';
DEFAULT: 'default';
INTERFACE: 'interface';
SUPER: 'super';
EXTENDS: 'extends';
IMPLEMENTS: 'implements';
OPERATIR: 'operator';
IN: 'in';
OF: 'of';
LUA: 'lua.';
STATIC: 'static';
FROM: 'from';
EXPORT: 'export';
GO: 'go';
AWAIT: 'await';
ASYNC: 'async';

VOID: 'void';
STRING: 'string';
ANY: 'any';
NUMBER: 'number';
BOOL: 'bool';

// Literals

DECIMAL_LITERAL: ('0' | [1-9] (Digits? | '_'+ Digits)) [lL]?;
HEX_LITERAL:
	'0' [xX] [0-9a-fA-F] ([0-9a-fA-F_]* [0-9a-fA-F])? [lL]?;
OCT_LITERAL: '0' '_'* [0-7] ([0-7_]* [0-7])? [lL]?;
BINARY_LITERAL: '0' [bB] [01] ([01_]* [01])? [lL]?;

FLOAT_LITERAL: (Digits '.' Digits? | '.' Digits) ExponentPart? [fFdD]?
	| Digits (ExponentPart [fFdD]? | [fFdD]);

HEX_FLOAT_LITERAL:
	'0' [xX] (HexDigits '.'? | HexDigits? '.' HexDigits) [pP] [+-]? Digits [fFdD]?;

BOOL_LITERAL: 'true' | 'false';

CHAR_LITERAL: '\'' (~['\\\r\n] | EscapeSequence) '\'';

STRING_LITERAL: '"' (~["\\\r\n] | EscapeSequence)* '"';

TEXT_BLOCK: '"""' [ \t]* [\r\n] (. | EscapeSequence)*? '"""';

NULL_LITERAL: 'nil';

// Separators

LPAREN: '(';
RPAREN: ')';
LBRACE: '{';
RBRACE: '}';
LBRACK: '[';
RBRACK: ']';
SEMI: ';';
COMMA: ',';
DOT: '.';

// Operators

ASSIGN: '=';
GT: '>';
LT: '<';
BANG: '!';
TILDE: '~';
QUESTION: '?';
COLON: ':';
EQUAL: '==';
LE: '<=';
GE: '>=';
NOTEQUAL: '!=';
AND: '&&';
OR: '||';
INC: '++';
DEC: '--';
ADD: '+';
SUB: '-';
MUL: '*';
DIV: '/';
BITAND: '&';
BITOR: '|';
CARET: '^';
MOD: '%';

ADD_ASSIGN: '+=';
SUB_ASSIGN: '-=';
MUL_ASSIGN: '*=';
DIV_ASSIGN: '/=';
AND_ASSIGN: '&=';
OR_ASSIGN: '|=';
XOR_ASSIGN: '^=';
MOD_ASSIGN: '%=';
LSHIFT_ASSIGN: '<<=';
RSHIFT_ASSIGN: '>>=';
URSHIFT_ASSIGN: '>>>=';

ARROW: '->';
COLONCOLON: '::';

AT: '@';
ELLIPSIS: '...';

// Identifiers

IDENTIFIER: Letter LetterOrDigit*;

// Fragment rules

fragment ExponentPart: [eE] [+-]? Digits;

fragment EscapeSequence:
	'\\' [btnfr"'\\]
	| '\\' ([0-3]? [0-7])? [0-7]
	| '\\' 'u'+ HexDigit HexDigit HexDigit HexDigit;

fragment HexDigits: HexDigit ((HexDigit | '_')* HexDigit)?;

fragment HexDigit: [0-9a-fA-F];

fragment Digits: [0-9] ([0-9_]* [0-9])?;

fragment LetterOrDigit: Letter | [0-9];

fragment Letter:
	[a-zA-Z$_]
	| ~[\u0000-\u007F\uD800-\uDBFF] // covers all characters above 0x7F which are not a surrogate
	| [\uD800-\uDBFF] [\uDC00-\uDFFF];
	// covers UTF-16 surrogate pairs encodings for U+10000 to U+10FFFF

// Whitespace and comments

WS: [ \t\r\n\u000C]+ -> channel(HIDDEN);
COMMENT: '/*' .*? '*/' -> channel(HIDDEN);
LINE_COMMENT: '//' ~[\r\n]* -> channel(HIDDEN);