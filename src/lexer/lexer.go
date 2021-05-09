package lexer

import (
	"strings"

	"github.com/seralto/neela/src/token"
)

type Lexer struct {
	input       string
	position    int
	nexPosition int
	char        byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()

	return l
}

func (l *Lexer) NextToken() token.Token {
	t := token.Token{}

	l.skipWhitespace()

	switch l.char {
	case '+':
		t = token.New(token.PLUS, "+")
	case '-':
		t = token.New(token.MINUS, "-")
	case '*':
		t = token.New(token.MULTIPLY, "*")
	case '/':
		t = token.New(token.DIVIDE, "/")
	case '=':
		t = token.New(token.PLUS, "=")
	case 0:
		t = token.New(token.EOF, "")
	default:
		if l.isIdentifier() {
			identifier := l.getIdentifier()
			identifierType := token.GetIdentifierType(identifier)
			t = token.New(identifierType, identifier)
		} else if l.isNumber() {
			number := l.getNumber()
			tokenType := getNumberType(number)
			t = token.New(tokenType, number)
		} else {
			t = token.New(token.INVALID, "")
		}
	}

	l.readChar()
	return t
}

func (l *Lexer) readChar() {
	if l.nexPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.nexPosition]
	}

	l.position = l.nexPosition
	l.nexPosition += 1
}

func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\n' || l.char == '\t' || l.char == '\r' {
		l.readChar()
	}
}

func (l *Lexer) isIdentifier() bool {
	return l.char >= 'a' && l.char <= 'z' || l.char >= 'A' && l.char <= 'Z' || l.char == '_'
}

func (l *Lexer) getIdentifier() string {
	start := l.position

	for l.isIdentifier() {
		l.readChar()
	}

	return l.input[start:l.position]
}

func (l *Lexer) isNumber() bool {
	return l.char >= '0' && l.char <= '9'
}

func (l *Lexer) getNumber() string {
	start := l.position

	for l.isNumber() || l.char == '.' {
		l.readChar()
	}

	return l.input[start:l.position]
}

func getNumberType(num string) token.TokenType {
	split := strings.Split(num, ".")
	dotCount := len(split) - 1

	if dotCount > 1 {
		// INVALID NUMBER, SHOULD PANIC
		return token.INVALID
	} else if dotCount == 1 {
		return token.FLOAT
	}

	return token.INTEGER
}