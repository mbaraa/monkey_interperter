/*
Package token includes the token related stuff

*/
package token

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

// TokenType defines the different types of tokens
type TokenType string

// TokenType contstants
const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// identifiers and literals
	IDENT TokenType = "IDENT"
	INT   TokenType = "INT"

	// operators
	ASSIGN  TokenType = "="
	PLUS    TokenType = "+"
	MINUS   TokenType = "-"
	BANG    TokenType = "!"
	ASTRISK TokenType = "*"
	SLASH   TokenType = "/"

	LT     TokenType = "<"
	GT     TokenType = ">"
	EQ     TokenType = "=="
	NOT_EQ TokenType = "!="

	// delimeters
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"

	LPAREN TokenType = "("
	RPAREN TokenType = ")"
	LBRACE TokenType = "{"
	RBRACE TokenType = "}"

	// keywords
	FUNCTION TokenType = "FUNCTION"
	LET      TokenType = "LET"
	TRUE     TokenType = "TRUE"
	FALSE    TokenType = "FALSE"
	IF       TokenType = "IF"
	ELSE     TokenType = "ELSE"
	RETURN   TokenType = "RETURN"
)

// Token represents a single token in the monkey language
type Token struct {
	Type    TokenType
	Literal string
}

// New returns a token from the given specs,
// for EOF it assigns "" as literal
func New(tp TokenType, chr byte) (tok Token) {
	tok = Token{Type: tp, Literal: string(chr)}
	if chr == 0 {
		tok.Literal = ""
	}

	return
}
