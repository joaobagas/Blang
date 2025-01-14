package main

import (
	"Blang/src"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("text.txt")
	if err != nil {
		panic(err)
	}

	lexer := src.CreateLexer(file)
	for {
		pos, tok, lit := lexer.Lex()
		if tok == src.EOF {
			break
		}

		fmt.Printf("%d:%d\t%s\t%s\n", pos.Line, pos.Column, tok, lit)
	}
}
