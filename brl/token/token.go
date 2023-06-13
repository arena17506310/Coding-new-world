package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // 잘못된 토큰
	EOF     = "EOF"     // 파일의 끝 토큰

	// 식별자와 리터럴
	IDENT  = "IDENT"  // 식별자 토큰
	INT    = "INT"    // 정수 리터럴 토큰
	STRING = "STRING" // 문자열 리터럴 토큰

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

	// type
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

type Object struct {
	Type  TokenType
	Value interface{}
}

func NewObject(tokenType TokenType, value interface{}) Object {
	return Object{
		Type:  tokenType,
		Value: value,
	}
}
