package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int
	readposition int
	char         byte
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.char {
	case '=':
		tok = newToken(token.ASSIGN, l.char)

	case ';':
		tok = newToken(token.SEMICOLON, l.char)
	case '(':
		tok = newToken(token.LPAREN, l.char)
	case ')':
		tok = newToken(token.RPAREN, l.char)
	case ',':
		tok = newToken(token.COMMA, l.char)
	case '+':
		tok = newToken(token.PLUS, l.char)
	case '{':
		tok = newToken(token.LBRACE, l.char)
	case '}':
		tok = newToken(token.RBRACE, l.char)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

	default:
		if isLetter(l.char) {
			tok.Literal = l.readIdentifier()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.char)
		}

	}

	l.readChar()

	return tok
}

func (l *Lexer) readChar() {
	if l.readposition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readposition]

	}
	l.position = l.readposition
	l.readposition += 1
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.char) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func New(input string) *Lexer {
	l := &Lexer{input: input}

	l.readChar()

	return l
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}

func isLetter(ch byte) bool { return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' }
