package ast

import (
	"testing"

	"github.com/odit-bit/monkey/token"
	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	prog := Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{
					Type:    token.LET,
					Literal: "let",
				},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "varValue"},
					Value: "varValue",
				},
			},
		},
	}

	assert.Equal(t, "let myVar = varValue;", prog.String())
}
