package ast

import (
	"bytes"
	"fmt"

	"github.com/odit-bit/monkey/token"
)

// STATEMENT //

// LET
var _ Node = (*LetStatement)(nil)

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (lstate *LetStatement) String() string {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%s %s = ", lstate.TokenLiteral(), lstate.Name.String()))
	if lstate.Value != nil {
		buf.WriteString(lstate.Value.String())
	}

	buf.WriteString(";")
	return buf.String()
}

func (lstate *LetStatement) statementNode() {}
func (lstate *LetStatement) TokenLiteral() string {
	return lstate.Token.Literal
}

// Return
var _ Node = (*ReturnStatement)(nil)

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (reState *ReturnStatement) String() string {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%s \n", reState.TokenLiteral()))
	if reState.ReturnValue != nil {
		buf.WriteString(fmt.Sprintf("%s \n", reState.ReturnValue.String()))
	}
	return buf.String()
}

func (reState *ReturnStatement) statementNode() {}
func (reState *ReturnStatement) TokenLiteral() string {
	return reState.Token.Literal
}

// IDENT
var _ Node = (*Identifier)(nil)

type Identifier struct {
	Token token.Token
	Value string
}

func (ident *Identifier) String() string {
	return ident.Value
}

func (ident *Identifier) expressionNode() {}
func (ident *Identifier) TokenLiteral() string {
	return ident.Token.Literal
}

// // COMMENT
// var _ Statement = (*Comment)(nil)

// type Comment struct {
// 	Token  token.Token
// 	Values []Expression
// }

// // String implements Statement.
// func (c *Comment) String() string {
// 	args := []string{}
// 	for _, v := range c.Values {
// 		args = append(args, v.String())
// 	}

// 	return strings.Join(args, " ")
// }

// // TokenLiteral implements Statement.
// func (c *Comment) TokenLiteral() string {
// 	return c.Token.Literal
// }

// // statementNode implements Statement.
// func (c *Comment) statementNode() {}
