package ast

import (
	"bytes"
)

type Node interface {
	TokenLiteral() string // primarly use for testing
	String() string
}

type Statement interface {
	Node
	statementNode() //dummy method for distinguish the node type
}

type Expression interface {
	Node
	expressionNode() //dummy method for distinguish the node type
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {

	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, v := range p.Statements {
		out.WriteString(v.String())
	}

	return out.String()
}
