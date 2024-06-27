package ast

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/odit-bit/monkey/token"
)

// Expression Node Implementation//

var _ Node = (*ExpressionStatement)(nil)
var _ Statement = (*ExpressionStatement)(nil)

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

// statementNode implements Statement.
func (e *ExpressionStatement) statementNode() {
}

// String implements Node.
func (e *ExpressionStatement) String() string {

	if e.Expression != nil {
		return e.Expression.String()
	}
	return ""
}

func (e *ExpressionStatement) TokenLiteral() string {
	return e.Token.Literal
}

func (e *ExpressionStatement) expressionNode() {}

type Prefix struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *Prefix) TokenLiteral() string {
	return pe.Token.Literal
}

func (pe *Prefix) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("(%s%s)", pe.Operator, pe.Right.String()))
	return buf.String()
}

func (pe *Prefix) expressionNode() {}

type Infix struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *Infix) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *Infix) String() string {
	return fmt.Sprintf("(%s %s %s)", ie.Left.String(), ie.Operator, ie.Right.String())
}

func (ip *Infix) expressionNode() {}

// Boolean
type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }

// IF Expression
var _ Expression = (*IF)(nil)

type IF struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (i *IF) String() string {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("if%s %s", i.Condition.String(), i.Consequence.String()))

	if i.Alternative != nil {
		buf.WriteString(fmt.Sprintf("else%s", i.Alternative.String()))
	}

	return buf.String()
}

func (i *IF) TokenLiteral() string {
	return i.Token.Literal
}

func (i *IF) expressionNode() {}

// IF Block
var _ Statement = (*BlockStatement)(nil)

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (b *BlockStatement) String() string {
	var buf bytes.Buffer
	for _, v := range b.Statements {
		buf.WriteString(v.String())
	}
	return buf.String()
}

func (b *BlockStatement) TokenLiteral() string {
	return b.Token.Literal
}
func (b *BlockStatement) statementNode() {}

//Function literal

var _ Expression = (*FunctionLiteral)(nil)

type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (f *FunctionLiteral) String() string {
	var buf bytes.Buffer
	//write syntax ("fn")
	buf.WriteString(f.TokenLiteral())

	//write param
	var params []string
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	buf.WriteString(fmt.Sprintf("(%s)", strings.Join(params, ",")))

	//write block
	buf.WriteString(f.Body.String())

	return buf.String()
}

func (f *FunctionLiteral) TokenLiteral() string {
	return f.Token.Literal
}

func (f *FunctionLiteral) expressionNode() {
}

// Call Expression

var _ Expression = (*CallExpr)(nil)

type CallExpr struct {
	Token    token.Token // The '(' token
	Function Expression  // Identifier or FunctionLiteral
	Args     []Expression
}

func (c *CallExpr) String() string {
	var buf bytes.Buffer

	var args []string
	for _, arg := range c.Args {
		args = append(args, arg.String())
	}
	buf.WriteString(fmt.Sprintf("%s(%s)", c.Function.String(), strings.Join(args, ", ")))
	return buf.String()
}

func (c *CallExpr) TokenLiteral() string {
	return c.Token.Literal
}

func (c *CallExpr) expressionNode() {}
