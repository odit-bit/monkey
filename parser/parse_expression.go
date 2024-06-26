package parser

import (
	"fmt"
	"strconv"

	"github.com/odit-bit/monkey/ast"
	"github.com/odit-bit/monkey/token"
)

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := ast.ExpressionStatement{
		Token: p.currToken,
	}
	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekToken.Type == token.SEMICOLON {
		p.nextToken()
	}

	return &stmt
}

func (p *Parser) parseExpression(preced int) ast.Expression {
	leftExpFunc := p.prefixParse[p.currToken.Type]
	if leftExpFunc == nil {
		p.addError(fmt.Errorf("no prefix parse function for %s found", p.currToken.Type))
		return nil
	}
	leftExpr := leftExpFunc()

	for p.peekToken.Type != token.SEMICOLON && preced < p.peekPrecedence() {
		inflixFunc := p.infixParse[p.peekToken.Type]
		if inflixFunc == nil {
			return leftExpr
			// p.addError(fmt.Errorf("no inflix parse function for %s found", p.currToken.Type))
		}

		p.nextToken()
		leftExpr = inflixFunc(leftExpr)
	}

	return leftExpr
}

func (p *Parser) parseIdentifier() ast.Expression {
	ident := ast.Identifier{}
	ident.Token = p.currToken
	ident.Value = p.currToken.Literal
	return &ident
}

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

/////////////////////////////////////

// prefix and infix expression
type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

func (p *Parser) registerPrefixFunc(typ token.TokenType, fn prefixParseFn) {
	p.prefixParse[typ] = fn
}

func (p *Parser) registerInfixFunc(typ token.TokenType, fn infixParseFn) {
	p.infixParse[typ] = fn
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expr := ast.Prefix{
		Token:    p.currToken,
		Operator: p.currToken.Literal,
		Right:    nil,
	}

	p.nextToken()

	expr.Right = p.parseExpression(PREFIX)
	return &expr
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expr := ast.Infix{
		Token:    token.Token{},
		Left:     left,
		Operator: p.currToken.Literal,
		Right:    nil,
	}

	preced := p.currPrecedence()
	p.nextToken()
	expr.Right = p.parseExpression(preced)
	return &expr
}

func (p *Parser) parseBoolean() ast.Expression {

	boolVal := p.currToken.Type == token.TRUE

	expr := ast.Boolean{
		Token: p.currToken,
		Value: boolVal,
	}

	return &expr
}

// inspect this method , too many magic
func (p *Parser) parseGroupExpression() ast.Expression {
	p.nextToken()

	expr := p.parseExpression(LOWEST)
	if p.peekToken.Type == token.RPAREN {
		p.nextToken()
	} else {
		return nil
	}

	return expr
}

func (p *Parser) parseIF() ast.Expression {
	// "if (x > 5) {x + 1}"

	ifExpr := ast.IF{
		Token:       p.currToken,
		Condition:   nil,
		Consequence: &ast.BlockStatement{},
		Alternative: &ast.BlockStatement{},
	}

	// peek and read "("
	if p.peekToken.Type == token.LPAREN {
		p.nextToken()
	} else {
		return nil
	}

	p.nextToken()
	ifExpr.Condition = p.parseExpression(LOWEST)

	// peek and read ")"
	if p.peekToken.Type == token.RPAREN {
		p.nextToken()
	} else {
		return nil
	}

	// peek and read "{"
	if p.peekToken.Type == token.LBRACE {
		p.nextToken()
	} else {
		return nil
	}

	ifExpr.Consequence = p.parseBlockStatement()

	if p.peekToken.Type == token.ELSE {
		p.nextToken()

		// peek and read "{"
		if p.peekToken.Type == token.LBRACE {
			p.nextToken()
		} else {
			return nil
		}
		ifExpr.Alternative = p.parseBlockStatement()
	}
	return &ifExpr

}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := ast.BlockStatement{
		Token:      p.currToken,
		Statements: []ast.Statement{},
	}

	p.nextToken()
	for p.currToken.Type != token.RBRACE && p.currToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}

	return &block
}

func (p *Parser) parseFunction() ast.Expression {
	fnExpr := ast.FunctionLiteral{
		Token:      p.currToken,
		Parameters: []*ast.Identifier{},
		Body:       &ast.BlockStatement{},
	}

	// read and peek '('
	if p.peekToken.Type == token.LPAREN {
		p.nextToken()
	} else {
		return nil
	}

	fnExpr.Parameters = p.parseFuncParam()

	//expecting '{'
	if p.peekToken.Type == token.LBRACE {
		p.nextToken()
	} else {
		return nil
	}

	fnExpr.Body = p.parseBlockStatement()

	return &fnExpr
}

func (p *Parser) parseFuncParam() []*ast.Identifier {
	idents := []*ast.Identifier{}

	// peek  ')'
	if p.peekToken.Type == token.RPAREN {
		// read and got no parameter
		p.nextToken()
		return idents
	}

	// got identifier
	p.nextToken()
	ident := ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}
	idents = append(idents, &ident)

	//parse next identifiers
	for p.peekToken.Type == token.COMMA {
		p.nextToken() // point to comma
		p.nextToken() // point next to comma
		ident := ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}
		idents = append(idents, &ident)
	}

	// read and peek ")"
	if p.peekToken.Type == token.RPAREN {
		//found the closing parent , parameter complete
		p.nextToken()
	} else {
		// incomplete parameter closing syntax
		return nil
	}

	return idents
}

// infix func
func (p *Parser) parseCallExpr(expr ast.Expression) ast.Expression {
	call := ast.CallExpr{
		Token:    p.currToken,
		Function: expr,
		Args:     p.parseCallArgs(),
	}

	return &call
}

func (p *Parser) parseCallArgs() []ast.Expression {
	args := []ast.Expression{}

	// case for empty argument
	if p.peekToken.Type == token.RPAREN {
		p.nextToken()
		return args
	}

	p.nextToken()
	args = append(args, p.parseExpression(LOWEST))

	for p.peekToken.Type == token.COMMA {
		p.nextToken()
		p.nextToken()
		args = append(args, p.parseExpression(LOWEST))
	}

	if p.peekToken.Type == token.RPAREN {
		p.nextToken()
	} else {
		return nil
	}

	return args
}
