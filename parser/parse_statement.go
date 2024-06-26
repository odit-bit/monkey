package parser

import (
	"fmt"

	"github.com/odit-bit/monkey/ast"
	"github.com/odit-bit/monkey/token"
)

func (p *Parser) parseLetStatement() ast.Statement {
	letStmt := ast.LetStatement{
		Token: p.currToken,
	}

	//peek and read ident
	if p.peekToken.Type == token.IDENT {
		p.nextToken()
	} else {
		err := fmt.Errorf("expected peek type %s, got %s ", token.IDENT, p.peekToken.Type)
		p.addError(err)
		return nil
	}

	//assign identifier
	letStmt.Name = &ast.Identifier{
		Token: p.currToken,
		Value: p.currToken.Literal,
	}

	// peek and read assign
	if p.peekToken.Type == token.ASSIGN {
		p.nextToken()
	} else {
		err := fmt.Errorf("expected peek type %s, got %s ", token.ASSIGN, p.peekToken.Type)
		p.addError(err)
		return nil
	}

	p.nextToken()
	letStmt.Value = p.parseExpression(LOWEST)

	if p.peekToken.Type == token.SEMICOLON {
		p.nextToken()
	}

	return &letStmt
	///....
}

func (p *Parser) parseReturnStatement() ast.Statement {
	//return <expression>
	retStmt := ast.ReturnStatement{
		Token: p.currToken,
	}
	p.nextToken()

	//
	retStmt.ReturnValue = p.parseExpression(LOWEST)
	if p.peekToken.Type == token.SEMICOLON {
		p.nextToken()
	}

	return &retStmt
}

// func (p *Parser) parseComment() ast.Statement {
// 	cmt := ast.Comment{
// 		Token:  p.currToken,
// 		Values: []ast.Expression{},
// 	}

// 	p.nextToken()
// 	for p.currToken.Type != token.SEMICOLON && p.currToken.Type != token.EOF {
// 		stmt := p.parseExpression(LOWEST)
// 		cmt.Values = append(cmt.Values, stmt)
// 		p.nextToken()
// 	}
// 	return &cmt
// }