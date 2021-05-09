package main

import (
	"fmt"

	"github.com/seralto/neela/src/lexer"
	"github.com/seralto/neela/src/token"
)

func main() {
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

	l := lexer.New(input)

	for range input {
		t := l.NextToken()
		fmt.Println(t)
		if t.Type == token.EOF {
			break
		}
	}
}
