package parser

import (
	"fmt"
	"strconv"

	"github.com/odit-bit/monkey/ast"
	"github.com/odit-bit/monkey/token"
)

func (p *Parser) parseIntegerLiteral() ast.Expression {
	n, err := strconv.Atoi(p.currToken.Literal)
	if err != nil {
		p.addError(fmt.Errorf("wrong type of integer, got %s", p.currToken.Literal))
		return nil
	}
	stmt := ast.Integer{
		Token: p.currToken,
		Value: int64(n),
	}

	return &stmt
}

func (p *Parser) parseStringLiteral() ast.Expression {
	stmt := ast.String{
		Token: p.currToken,
		Value: p.currToken.Literal,
	}
	return &stmt
}

func (p *Parser) parseArrayLiteral() ast.Expression {
	arr := ast.Array{
		Token:    p.currToken,
		Elements: []ast.Expression{},
	}

	arr.Elements = p.parseExpressionList(token.RBRACKET)
	return &arr
}
