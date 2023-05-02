package token

type TokenType string

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF               = "EOF"

	// 식별자 및 리터럴
	IDENT  = "IDENT"  // add, foobar, x, y, ...
	INT    = "INT"    // 123456
	STRING = "STRING" // "hello world"

	// 연산자
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	BANG   = "!"
	ASTER  = "*"
	SLASH  = "/"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	// 구분자
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	LBRACK = "["
	RBRACK = "]"

	// 키워드
	FUNC        = "func"
	LET         = "let"
	IF          = "if"
	ELSE        = "else"
	RETURN      = "return"
	TRUE        = "true"
	FALSE       = "false"
	BOOL        = "bool"
	INT_TYPE    = "int"
	STRING_TYPE = "string"
)
