package src

import (
	"fmt"
	"os"
	"strings"
)

func Compile(path string) bool {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	lexer := CreateLexer(file)
	parser := CreateParser(lexer)
	ast := parser.Parse()

	printAST(ast, 0)
	return true
}

// Helper function to print the AST.
func printAST(node *ASTNode, level int) {
	if node == nil {
		return
	}
	fmt.Printf("%s%s (%q)\n", strings.Repeat("  ", level), node.Token, node.Literal)
	for _, child := range node.Children {
		printAST(child, level+1)
	}
}
