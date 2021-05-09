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
	case '(':
		t = token.New(token.LEFT_PAREN, "(")
	case ')':
		t = token.New(token.RIGHT_PAREN, ")")
	case '{':
		t = token.New(token.LEFT_BRACE, "{")
	case '}':
		t = token.New(token.RIGHT_BRACE, "}")
	case '[':
		t = token.New(token.LEFT_BRACKET, "[")
	case ']':
		t = token.New(token.RIGHT_BRACKET, "]")
	case '.':
		t = token.New(token.BLOCK_CLOSE, ".")
	case ',':
		t = token.New(token.COMMA, ",")
	case ':':
		t = token.New(token.COLON, ":")
	case 0:
		t = token.New(token.EOF, "")
	case '"':
		str := l.getString()
		t = token.New(token.STRING, str)
	case '=':
		if l.input[l.nexPosition] == '=' {
			t = token.New(token.EQUAL, "==")
			l.readChar()
		} else {
			t = token.New(token.ASSIGN, "=")
		}
	case '!':
		if l.input[l.nexPosition] == '=' {
			t = token.New(token.NOT_EQUAL, "!=")
			l.readChar()
		} else {
			t = token.New(token.NOT, "!")
		}
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

func (l *Lexer) getString() string {
	start := l.position
	l.readChar()

	for l.char != '"' {
		l.readChar()
	}

	return l.input[start:l.nexPosition]
}

func (l *Lexer) isIdentifier() bool {
	return l.char >= 'a' && l.char <= 'z' || l.char >= 'A' && l.char <= 'Z' || l.char == '_'
}

func (l *Lexer) getIdentifier() string {
	start := l.position

	for l.isIdentifier() {
		l.readChar()
	}

	l.nexPosition = l.position

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

	l.nexPosition = l.position

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
