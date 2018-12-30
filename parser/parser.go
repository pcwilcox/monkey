// parser/parser.go
//
// Definitions for the parser

package parser

import (
	"../ast"
	"../lexer"
	"../token"
)

// Parser is the parser object
type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

// New creates a new parser
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram doesn't do anything at the moment
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
