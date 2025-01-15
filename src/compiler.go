package src

import (
	"fmt"
	"os"
)

func Compile(path string) bool {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	lexer := CreateLexer(file)
	for {
		pos, tok, lit := lexer.Lex()
		if tok == EOF {
			break
		}

		fmt.Printf("%d:%d\t%s\t%s\n", pos.Line, pos.Column, tok, lit)
	}

	return true
}
