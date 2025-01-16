package src

import (
	"bufio"
	"io"
	"unicode"
)

// Token represents the type of lexical tokens that the lexer can recognize.
type Token int

// Position holds the line and column information for tokens.
type Position struct {
	Line   int
	Column int
}

// Lexer reads and processes the input to generate tokens.
type Lexer struct {
	pos    Position      // Current position in the input
	reader *bufio.Reader // Reader to read input character by character
}

// Token constants representing different types of tokens.
const (
	EOF     = iota // End of file
	ILLEGAL        // Illegal character or unrecognized token
	IDENT          // Identifier (e.g., variable names)
	INT            // Integer literal
	SEMI           // Semicolon (;)
	ADD            // Addition operator (+)
	SUB            // Subtraction operator (-)
	MUL            // Multiplication operator (*)
	DIV            // Division operator (/)
	ASSIGN         // Assignment operator (=)
)

// tokens provides string representations for the token constants.
var tokens = []string{
	EOF:     "EOF",
	ILLEGAL: "ILLEGAL",
	IDENT:   "IDENT",
	INT:     "INT",
	SEMI:    ";",
	ADD:     "+",
	SUB:     "-",
	MUL:     "*",
	DIV:     "/",
	ASSIGN:  "=",
}

// String returns the string representation of a Token.
func (t Token) String() string {
	return tokens[t]
}

// CreateLexer initializes a new Lexer with the given input reader.
func CreateLexer(reader io.Reader) *Lexer {
	return &Lexer{
		pos:    Position{Line: 1, Column: 0},
		reader: bufio.NewReader(reader),
	}
}

// Lex processes the input and returns the position, token, and its literal value.
func (l *Lexer) Lex() (Position, Token, string) {
	for {
		tempRune, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				// Return EOF token when end of file is reached.
				return l.pos, EOF, ""
			}

			// Unexpected error during reading; propagate it.
			panic(err)
		}

		l.pos.Column++

		switch tempRune {
		case '\n':
			// Newline character; reset position to the start of the next line.
			l.resetPosition()
		case ';':
			return l.pos, SEMI, ";"
		case '+':
			return l.pos, ADD, "+"
		case '-':
			return l.pos, SUB, "-"
		case '*':
			return l.pos, MUL, "*"
		case '/':
			return l.pos, DIV, "/"
		case '=':
			return l.pos, ASSIGN, "="
		default:
			if unicode.IsSpace(tempRune) {
				// Skip whitespace characters.
				continue
			} else if unicode.IsDigit(tempRune) {
				// Back up to allow lexInt to handle the integer parsing.
				startPos := l.pos
				l.backup()
				lit := l.lexInt()
				return startPos, INT, lit
			} else if unicode.IsLetter(tempRune) {
				// Back up to allow lexIdent to handle the identifier parsing.
				startPos := l.pos
				l.backup()
				lit := l.lexIdent()
				return startPos, IDENT, lit
			} else {
				// Unrecognized character; return an ILLEGAL token.
				return l.pos, ILLEGAL, string(tempRune)
			}
		}
	}
}

// resetPosition moves the lexer to the start of the next line.
func (l *Lexer) resetPosition() {
	l.pos.Line++
	l.pos.Column = 0
}

// backup moves the reader back by one rune and adjusts the position.
func (l *Lexer) backup() {
	if err := l.reader.UnreadRune(); err != nil {
		panic(err)
	}

	l.pos.Column--
}

// lexInt scans the input to read a complete integer literal.
func (l *Lexer) lexInt() string {
	var lit string
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				// End of the integer literal.
				return lit
			}
		}

		l.pos.Column++
		if unicode.IsDigit(r) {
			// Append digit to the integer literal.
			lit = lit + string(r)
		} else {
			// Non-digit encountered; back up and return the literal.
			l.backup()
			return lit
		}
	}
}

// lexIdent scans the input to read a complete identifier.
func (l *Lexer) lexIdent() string {
	var lit string
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				// End of the identifier.
				return lit
			}
		}

		l.pos.Column++
		if unicode.IsLetter(r) {
			// Append letter to the identifier literal.
			lit = lit + string(r)
		} else {
			// Non-letter encountered; back up and return the literal.
			l.backup()
			return lit
		}
	}
}
