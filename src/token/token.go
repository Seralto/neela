package token

type TokenType string

type Token struct {
	Type  TokenType
	Value string
}

const (
	PLUS     = "PLUS"
	MINUS    = "MINUS"
	DIVIDE   = "DIVIDE"
	MULTIPLY = "MULTIPLY"
	ASSIGN   = "ASSIGN"

	IDENTIFIER = "IDENTIFIER"

	INTEGER = "INTEGER"
	FLOAT   = "FLOAT"
	STRING  = "STRING"

	OUT   = "OUT"
	IF    = "IF"
	TRUE  = "TRUE"
	FALSE = "FALSE"

	BLOCK_CLOSE = "BLOCK_CLOSE"
	LEFT_PAREN  = "LEFT_PARENT"
	RIGHT_PAREN = "RIGHT_PARENT"

	INVALID = "INVALID"
	EOF     = "EOF"
)

var identifiers = map[string]TokenType{
	"out":   OUT,
	"if":    IF,
	"false": FALSE,
}

func New(t TokenType, v string) Token {
	return Token{Type: t, Value: v}
}

func GetIdentifierType(i string) TokenType {
	if _, ok := identifiers[i]; ok {
		return identifiers[i]
	}

	return IDENTIFIER
}
