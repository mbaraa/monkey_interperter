/*
Package lexer contains the lexical analyzer(also called lexer or scanner) stuff,
that is it converts the program to a seris of tokens.

*/
package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int  // current char position
	readPosition int  // next char position
	ch           byte // current char
}

func New(input string) (l *Lexer) {
	l = &Lexer{input: input}
	l.readChar()

	return
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespaces()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			tok.Literal = "=="
			tok.Type = token.EQ
		} else {
			tok = token.New(token.ASSIGN, l.ch)
		}
	case ';':
		tok = token.New(token.SEMICOLON, l.ch)
	case '(':
		tok = token.New(token.LPAREN, l.ch)
	case ')':
		tok = token.New(token.RPAREN, l.ch)
	case '{':
		tok = token.New(token.LBRACE, l.ch)
	case '}':
		tok = token.New(token.RBRACE, l.ch)
	case '-':
		tok = token.New(token.MINUS, l.ch)
	case '+':
		tok = token.New(token.PLUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			tok.Literal = "!="
			tok.Type = token.NOT_EQ
		} else {
			tok = token.New(token.BANG, l.ch)
		}
	case '*':
		tok = token.New(token.ASTRISK, l.ch)
	case '/':
		tok = token.New(token.SLASH, l.ch)
	case '<':
		tok = token.New(token.LT, l.ch)
	case '>':
		tok = token.New(token.GT, l.ch)
	case ',':
		tok = token.New(token.COMMA, l.ch)
	case 0:
		tok = token.New(token.EOF, 0)
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdent()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = token.New(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // reached EOF
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) readIdent() string {
	initPosition := l.position
	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[initPosition:l.position]
}

func (l *Lexer) readNumber() string {
	initPosition := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[initPosition:l.position]
}

func (l *Lexer) skipWhitespaces() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\r' || l.ch == '\t' {
		l.readChar()
	}
}

func isLetter(chr byte) bool {
	return chr >= 'a' && chr <= 'z' || chr >= 'A' && chr <= 'Z' || chr == '_'
}

func isDigit(chr byte) bool {
	return chr >= '0' && chr <= '9'
}
