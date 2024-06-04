package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"

	EOF = "EOF"

	IDENT = "IDENT" //foobar,x,y, you_name_it

	INT = "INT"

	//operators
	ASSIGN = "="

	PLUS = "+"

	COMMA = ","

	SEMICOLON = ";"

	LPAREN = "("

	RPAREN = ")"

	LBRACE = "{"

	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
