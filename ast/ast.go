// ast/ast.go
//
// defines the abstract syntax tree

package ast

import (
	"../token"
)

// Node is a node in the tree
type Node interface {
	TokenLiteral() string
}

// Statement is implemented by nodes representing statements
type Statement interface {
	Node
	statementNode()
}

// Expression is implemented by nodes representing expressions
type Expression interface {
	Node
	expressionNode()
}

// Program is the root node for the language
type Program struct {
	Statements []Statement
}

// LetStatement defines an assignment
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral returns the token literal
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Identifier represents function and variable names
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral returns the literal
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// TokenLiteral returns the literal for the root
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}
