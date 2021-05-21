package lexer

import (
	"testing"

	"github.com/seralto/neela/src/token"
)

func TestLexer_NextToken(t *testing.T) {
	input := `
		myVar = 42
		PI = 3.14

		2 * (3 + 5)

		out "I'm a string"

		if true
			out "You are right"
		.

		lie = false

		1 == 1
		5 != 10
		!true

		[{:,}]
	`

	tests := []struct {
		expectedType  token.TokenType
		expectedValue string
	}{
		{token.IDENTIFIER, "myVar"},
		{token.ASSIGN, "="},
		{token.INTEGER, "42"},
		{token.IDENTIFIER, "PI"},
		{token.ASSIGN, "="},
		{token.FLOAT, "3.14"},
		{token.INTEGER, "2"},
		{token.MULTIPLY, "*"},
		{token.LEFT_PAREN, "("},
		{token.INTEGER, "3"},
		{token.PLUS, "+"},
		{token.INTEGER, "5"},
		{token.RIGHT_PAREN, ")"},
		{token.OUT, "out"},
		{token.STRING, "\"I'm a string\""},
		{token.IF, "if"},
		{token.TRUE, "true"},
		{token.OUT, "out"},
		{token.STRING, "\"You are right\""},
		{token.BLOCK_CLOSE, "."},
		{token.IDENTIFIER, "lie"},
		{token.ASSIGN, "="},
		{token.FALSE, "false"},
		{token.INTEGER, "1"},
		{token.EQUAL, "=="},
		{token.INTEGER, "1"},
		{token.INTEGER, "5"},
		{token.NOT_EQUAL, "!="},
		{token.INTEGER, "10"},
		{token.NOT, "!"},
		{token.TRUE, "true"},
		{token.LEFT_BRACKET, "["},
		{token.LEFT_BRACE, "{"},
		{token.COLON, ":"},
		{token.COMMA, ","},
		{token.RIGHT_BRACE, "}"},
		{token.RIGHT_BRACKET, "]"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Value != tt.expectedValue {
			t.Fatalf("tests[%d] - tokenvalue wrong. expected=%q, got=%q", i, tt.expectedValue, tok.Value)
		}
	}
}
