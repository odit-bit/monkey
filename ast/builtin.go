package ast

import (
	"bytes"
	"strings"

	"github.com/odit-bit/monkey/token"
)

// built-in data types

// INTEGER Literal
type Integer struct {
	Token token.Token
	Value int64
}

func (il *Integer) TokenLiteral() string {
	return il.Token.Literal
}

func (il *Integer) String() string {
	return il.Token.Literal
}

func (il *Integer) expressionNode() {}

// STRING LITERAL
var _ Expression = (*String)(nil)

type String struct {
	Token token.Token
	Value string
}

// String implements Expression.
func (s *String) String() string {
	return s.Token.Literal
}

// TokenLiteral implements Expression.
func (s *String) TokenLiteral() string {
	return s.Token.Literal
}

// expressionNode implements Expression.
func (s *String) expressionNode() {}

// ARRAY literal
var _ Expression = (*Array)(nil)

type Array struct {
	Token    token.Token
	Elements []Expression
}

// String implements Expression.
func (a *Array) String() string {
	var buf bytes.Buffer

	values := []string{}
	for _, expr := range a.Elements {
		values = append(values, expr.String())
	}
	buf.WriteByte('[')
	buf.WriteString(strings.Join(values, ", "))
	buf.WriteByte(']')

	return buf.String()
}

// TokenLiteral implements Expression.
func (a *Array) TokenLiteral() string {
	return a.Token.Literal
}

// expressionNode implements Expression.
func (a *Array) expressionNode() {}
