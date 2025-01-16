package src

import (
	"fmt"
)

// ASTNode represents a node in the abstract syntax tree.
type ASTNode struct {
	Token    Token      // Token for this node
	Literal  string     // Literal value of the token
	Children []*ASTNode // Child nodes
}

// Parser reads tokens from the lexer and constructs an abstract syntax tree (AST).
type Parser struct {
	lexer *Lexer   // Lexer instance
	token Token    // Current token
	lit   string   // Current token literal
	pos   Position // Current token position
}

// CreateParser initializes a new parser with the given lexer.
func CreateParser(lexer *Lexer) *Parser {
	p := &Parser{lexer: lexer}
	p.nextToken() // Load the first token
	return p
}

// nextToken advances to the next token from the lexer.
func (p *Parser) nextToken() {
	p.pos, p.token, p.lit = p.lexer.Lex()
}

// expectToken ensures the current token matches the expected token.
func (p *Parser) expectToken(expected Token) {
	if p.token != expected {
		panic(fmt.Sprintf("Expected token %s at %v, but got %s (%q)", expected, p.pos, p.token, p.lit))
	}
	p.nextToken()
}

// Parse parses the input and returns the root of the AST.
func (p *Parser) Parse() *ASTNode {
	root := &ASTNode{Token: ILLEGAL, Literal: "program"}
	for p.token != EOF {
		root.Children = append(root.Children, p.parseStatement())
	}
	return root
}

// parseStatement parses a statement: `expr ';'`.
func (p *Parser) parseStatement() *ASTNode {
	var node *ASTNode
	if p.token == IDENT {
		// Check for assignment
		identNode := &ASTNode{Token: IDENT, Literal: p.lit}
		p.nextToken()
		if p.token == ASSIGN {
			assignNode := &ASTNode{Token: ASSIGN, Literal: p.lit}
			p.nextToken()
			assignNode.Children = append(assignNode.Children, identNode)
			assignNode.Children = append(assignNode.Children, p.parseExpr())
			node = assignNode
		} else {
			// Not an assignment, treat as an expression
			node = identNode
		}
	} else {
		// Parse as an expression
		node = p.parseExpr()
	}

	p.expectToken(SEMI) // Ensure statement ends with a semicolon
	return node
}

// parseExpr parses an expression: `term ((ADD | SUB) term)*`.
func (p *Parser) parseExpr() *ASTNode {
	node := p.parseTerm()
	for p.token == ADD || p.token == SUB {
		op := &ASTNode{Token: p.token, Literal: p.lit}
		p.nextToken()
		op.Children = append(op.Children, node)
		op.Children = append(op.Children, p.parseTerm())
		node = op
	}
	return node
}

// parseTerm parses a term: `factor ((MUL | DIV) factor)*`.
func (p *Parser) parseTerm() *ASTNode {
	node := p.parseFactor()
	for p.token == MUL || p.token == DIV {
		op := &ASTNode{Token: p.token, Literal: p.lit}
		p.nextToken()
		op.Children = append(op.Children, node)
		op.Children = append(op.Children, p.parseFactor())
		node = op
	}
	return node
}

// parseFactor parses a factor: `INT | IDENT | '(' expr ')'`.
func (p *Parser) parseFactor() *ASTNode {
	switch p.token {
	case INT:
		node := &ASTNode{Token: INT, Literal: p.lit}
		p.nextToken()
		return node
	case IDENT:
		node := &ASTNode{Token: IDENT, Literal: p.lit}
		p.nextToken()
		return node
	case '(':
		p.nextToken()
		node := p.parseExpr()
		p.expectToken(')')
		return node
	default:
		panic(fmt.Sprintf("Unexpected token %s (%q) at %v", p.token, p.lit, p.pos))
	}
}
