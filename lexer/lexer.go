package lexer

import (
	"github.com/odit-bit/monkey/token"
)

type Lexer struct {
	input   string
	pos     int
	readPos int
	ch      byte
}

func New(input string) *Lexer {
	l := Lexer{
		input:   input,
		pos:     0,
		readPos: 0,
		ch:      0,
	}
	l.readChar()
	return &l
}

func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPos]
	}
	l.pos = l.readPos
	l.readPos++
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

	switch l.ch {
	//operator
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			eq := string(ch) + string(l.ch)
			tok = token.Token{Literal: eq, Type: token.EQ}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}

	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		if l.peekChar() == '/' {
			l.readChar()
			l.skipComment()
			return l.NextToken()
		} else {
			tok = newToken(token.SLASH, l.ch)
		}

	case '-':
		tok = newToken(token.MINUS, l.ch)

	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			lit := string(ch) + string(l.ch)
			tok = token.Token{Literal: lit, Type: token.NOT_EQ}
		} else {
			tok = newToken(token.BANG, l.ch)
		}

	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)

		//delimiter
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)

	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

	//Data type
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()

	default:
		//this default will advanced the lexer by readIdentifier or readNumber
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func newToken(t token.TokenType, char byte) token.Token {
	return token.Token{
		Type:    t,
		Literal: string(char),
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPos > len(l.input) {
		return 0
	} else {
		return l.input[l.readPos]
	}
}

func (l *Lexer) readIdentifier() string {
	pos := l.pos

	for isLetter(l.ch) {
		l.readChar()
	}

	return string(l.input[pos:l.pos])
}

func (l *Lexer) readString() string {
	position := l.pos + 1 // pointed to char next of char ["]
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}

	return l.input[position:l.pos]
}

func (l *Lexer) readNumber() string {
	pos := l.pos

	for isDigit((l.ch)) {
		l.readChar()
	}

	return string(l.input[pos:l.pos])
}

func (l *Lexer) skipWhiteSpace() {
	// continue read until get non whitespace charater
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// advance pointer until found newline
func (l *Lexer) skipComment() {
	for l.ch != '\n' && l.ch != '\r' {
		l.readChar()
	}
}

//helper function

func isLetter(ch byte) bool {
	return 'A' <= ch && ch <= 'Z' || 'a' <= ch && ch <= 'z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
