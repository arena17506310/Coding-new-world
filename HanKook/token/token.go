package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// idntifier, literal
	IDENT  = "IDENT"  // add, foobar, x, y, ...
	INT    = "INT"    // 123456
	STRING = "STRING" // "hello world"

	// operator
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

	// separator
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	LBRACK = "["
	RBRACK = "]"

	// keyword
	START  = "start"
	END    = "end"
	FUNC   = "func"
	IF     = "sparrow"
	ELSE   = "else"
	RETURN = "return"
	PRINT  = "pigeon"

	//type
	BOOL_TYPE     = "bool"
	TRUE          = "true"
	FALSE         = "false"
	P_PLURAL_TYPE = "eyes"
	N_PLURAL_TYPE = "seye"
	P_SINGUL_TYPE = "peye"
	N_SINGUL_TYPE = "meye"
	STRING_TYPE   = "string"
	CHAR_TYPE     = "char"
	ARRAY         = "array"
)

// go ma wu yo chet gee pee tee
