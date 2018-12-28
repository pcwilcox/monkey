// token/token.go
//
// Defines the Token struct and TokenType type for the lexer

package token

// TokenType allows any string value to be used as a token type, for better perf should use an int
type TokenType string

// Token maps the token type to the lexical text
type Token struct {
	Type    TokenType
	Literal string
}

// Token codes
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers and literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
