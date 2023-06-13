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

func LookupIdent(ident string) TokenType {
	switch ident {
	case "goose":
		return "START" // 키워드 "goose"에 해당하는 토큰 타입 "START" 반환
	case "owl":
		return "END" // 키워드 "owl"에 해당하는 토큰 타입 "END" 반환
	case "duck":
		return "FUNC" // 키워드 "duck"에 해당하는 토큰 타입 "FUNC" 반환
	case "sparrow":
		return "IF" // 키워드 "sparrow"에 해당하는 토큰 타입 "IF" 반환
	case "esparrow":
		return "ELSE" // 키워드 "esparrow"에 해당하는 토큰 타입 "ELSE" 반환
	case "crane":
		return "RETURN" // 키워드 "crane"에 해당하는 토큰 타입 "RETURN" 반환
	case "pigeon":
		return "PRINT" // 키워드 "pigeon"에 해당하는 토큰 타입 "PRINT" 반환
	case "bool":
		return "BOOL_TYPE" // 키워드 "bool"에 해당하는 토큰 타입 "BOOL_TYPE" 반환
	case "feat":
		return "TRUE" // 키워드 "feat"에 해당하는 토큰 타입 "TRUE" 반환
	case "her":
		return "FALSE" // 키워드 "her"에 해당하는 토큰 타입 "FALSE" 반환
	case "eyes":
		return "P_PLURAL_TYPE" // 키워드 "eyes"에 해당하는 토큰 타입 "P_PLURAL_TYPE" 반환
	case "seye":
		return "N_PLURAL_TYPE" // 키워드 "seye"에 해당하는 토큰 타입 "N_PLURAL_TYPE" 반환
	case "peye":
		return "P_SINGUL_TYPE" // 키워드 "peye"에 해당하는 토큰 타입 "P_SINGUL_TYPE" 반환
	case "meye":
		return "N_SINGUL_TYPE" // 키워드 "meye"에 해당하는 토큰 타입 "N_SINGUL_TYPE" 반환
	case "beaks":
		return "STRING_TYPE" // 키워드 "beaks"에 해당하는 토큰 타입 "STRING_TYPE" 반환
	case "beak":
		return "CHAR_TYPE" // 키워드 "beak"에 해당하는 토큰 타입 "CHAR_TYPE" 반환
	case "nest":
		return "ARRAY" // 키워드 "nest"에 해당하는 토큰 타입 "ARRAY" 반환
	default:
		return IDENT // 키워드가 아니라면 식별자 토큰 타입 "IDENT" 반환
	}
}
