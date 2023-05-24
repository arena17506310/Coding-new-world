package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// idntifier, literal
	IDENT  = "IDENT" // add, foobar, x, y, ...
	INT    = "INT"   // 123456
	STRING = "beak"  // "hello world"

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
	START  = "goose"
	END    = "owl"
	FUNC   = "duck"
	IF     = "sparrow"
	ELSE   = "esparrow"
	RETURN = "crane"
	PRINT  = "pigeon"

	//type
	BOOL_TYPE     = "bool"
	TRUE          = "feat"
	FALSE         = "her"
	P_PLURAL_TYPE = "eyes"
	N_PLURAL_TYPE = "seye"
	P_SINGUL_TYPE = "peye"
	N_SINGUL_TYPE = "meye"
	STRING_TYPE   = "beaks"
	CHAR_TYPE     = "beak"
	ARRAY         = "nest"
)

// go ma wu yo chet gee pee tee
