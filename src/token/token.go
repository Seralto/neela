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

	OUT = "OUT"

	INVALID = "INVALID"
	EOF     = "EOF"
)

var identifiers = map[string]TokenType{
	"out": OUT,
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
