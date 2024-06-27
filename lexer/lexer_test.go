package lexer

import (
	"testing"

	"github.com/odit-bit/monkey/token"
	"github.com/stretchr/testify/assert"
)

func Test_read_comment(t *testing.T) {
	input := "// this is comment;"

	l := New(input)
	tkn := l.NextToken()
	assert.Equal(t, token.TokenType(token.EOF), tkn.Type)
}

func Test_nextToken(t *testing.T) {
	var input = `let five = 5;
	let ten = 10;

	let add = fn(x,y) {
		x + y;
	};

	let result = add(five, ten);

	// this comment
	!-/*<>==!=;
	// this comment


	"this is STRING"

	[1, "one"];

	`

	test := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},

		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.LT, "<"},
		{token.GT, ">"},
		{token.EQ, "=="},
		{token.NOT_EQ, "!="},
		{token.SEMICOLON, ";"},

		{token.STRING, "this is STRING"},

		{token.LBRACKET, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.STRING, "one"},
		{token.RBRACKET, "]"},
		{token.SEMICOLON, ";"},

		{token.EOF, ""},
	}

	l := New(input)

	for _, tc := range test {
		act := l.NextToken()

		assert.Equal(t, tc.expectedLiteral, act.Literal, "literal")
		assert.Equal(t, tc.expectedType, act.Type, "type")

	}
}
