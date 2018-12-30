// ast/ast.go
//
// defines the abstract syntax tree

package ast

import (
	"bytes"

	"../token"
)

// Node is a node in the tree
type Node interface {
	TokenLiteral() string
	String() string
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

// TokenLiteral returns the literal for the root
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// String is the stringer function
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
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

// String is the stringer
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// ReturnStatement defines a return statement
type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

//TokenLiteral gives the token
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

// String is the stringer
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// ExpressionStatement is all of the other types of statement
type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

// TokenLiteral gives the token
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

// String is the stringer function
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

// Identifier represents function and variable names
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral returns the literal
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// String is the stringer function
func (i *Identifier) String() string { return i.Value }

// IntegerLiteral is an integer literal
type IntegerLiteral struct {
	Token token.Token // the token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

// TokenLiteral returns the token
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }

// String
func (il *IntegerLiteral) String() string { return il.Token.Literal }

// PrefixExpression is a prefix operator expression
type PrefixExpression struct {
	Token    token.Token // the prefix token, e.g. !
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {}

// TokenLiteral is the token
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }

// String
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

// InfixExpression is a binary operator expression
type InfixExpression struct {
	Token    token.Token // the operator token e.g. '+'
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode() {}

// TokenLiteral returns the token
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }

// String
func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

// Boolean ...
type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode() {}

// TokenLiteral ...
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }

// String ...
func (b *Boolean) String() string { return b.Token.Literal }
