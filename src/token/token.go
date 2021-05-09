package token

type TokenType string

type Token struct {
	Type  TokenType
	Value string
}

const (
	PLUS         = "PLUS"
	MINUS        = "MINUS"
	DIVIDE       = "DIVIDE"
	MULTIPLY     = "MULTIPLY"
	ASSIGN       = "ASSIGN"
	EQUAL        = "EQUAL"
	NOT_EQUAL    = "NOT_EQUAL"
	NOT          = "NOT"
	GREATER_THAN = "GREATER_THAN"
	LESS_THAN    = "LESS_THAN"

	IDENTIFIER = "IDENTIFIER"

	INTEGER = "INTEGER"
	FLOAT   = "FLOAT"
	STRING  = "STRING"

	OUT    = "OUT"
	IF     = "IF"
	ELSE   = "ELSE"
	FUN    = "FUN"
	FOR    = "FOR"
	IN     = "IN"
	RETURN = "RETURN"

	TRUE  = "TRUE"
	FALSE = "FALSE"

	BLOCK_CLOSE = "BLOCK_CLOSE"
	COMMA       = "COMMA"
	COLON       = "COLON"

	LEFT_PAREN    = "LEFT_PARENT"
	RIGHT_PAREN   = "RIGHT_PARENT"
	LEFT_BRACE    = "LEFT_BRACE"
	RIGHT_BRACE   = "RIGHT_BRACE"
	LEFT_BRACKET  = "LEFT_BRACKET"
	RIGHT_BRACKET = "RIGHT_BRACKET"

	INVALID = "INVALID"
	EOF     = "EOF"
)

var keywords = map[string]TokenType{
	"out":    OUT,
	"if":     IF,
	"else":   ELSE,
	"fun":    FUN,
	"for":    FOR,
	"in":     IN,
	"return": RETURN,
	"false":  FALSE,
	"true":   TRUE,
}

func New(t TokenType, v string) Token {
	return Token{Type: t, Value: v}
}

func GetIdentifierType(i string) TokenType {
	if t, ok := keywords[i]; ok {
		return t
	}

	return IDENTIFIER
}
