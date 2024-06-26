// implement recursive descent parser
package parser

import (
	"github.com/odit-bit/monkey/ast"
	"github.com/odit-bit/monkey/lexer"
	"github.com/odit-bit/monkey/token"
)

type Parser struct {
	l *lexer.Lexer

	currToken token.Token //point to readed token
	peekToken token.Token //point to next token (without advance)

	prefixParse map[token.TokenType]prefixParseFn
	infixParse  map[token.TokenType]infixParseFn

	errors []error
}

func New(l *lexer.Lexer) *Parser {
	p := Parser{
		l:           l,
		prefixParse: map[token.TokenType]prefixParseFn{},
		infixParse:  map[token.TokenType]infixParseFn{},
		errors:      []error{},
	}

	// prefix
	p.registerPrefixFunc(token.IDENT, p.parseIdentifier)
	p.registerPrefixFunc(token.INT, p.parseIntegerLiteral)
	p.registerPrefixFunc(token.BANG, p.parsePrefixExpression)
	p.registerPrefixFunc(token.MINUS, p.parsePrefixExpression)
	p.registerPrefixFunc(token.TRUE, p.parseBoolean)
	p.registerPrefixFunc(token.FALSE, p.parseBoolean)
	p.registerPrefixFunc(token.LPAREN, p.parseGroupExpression)
	p.registerPrefixFunc(token.IF, p.parseIF)
	p.registerPrefixFunc(token.FUNCTION, p.parseFunction)

	// infix
	p.registerInfixFunc(token.PLUS, p.parseInfixExpression)
	p.registerInfixFunc(token.MINUS, p.parseInfixExpression)
	p.registerInfixFunc(token.SLASH, p.parseInfixExpression)
	p.registerInfixFunc(token.ASTERISK, p.parseInfixExpression)
	p.registerInfixFunc(token.EQ, p.parseInfixExpression)
	p.registerInfixFunc(token.NOT_EQ, p.parseInfixExpression)
	p.registerInfixFunc(token.LT, p.parseInfixExpression)
	p.registerInfixFunc(token.GT, p.parseInfixExpression)
	p.registerInfixFunc(token.LPAREN, p.parseCallExpr)

	//set currToken & peekToken
	p.nextToken()
	p.nextToken()
	return &p
}
func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()

}

func (p *Parser) ParseProgram() *ast.Program {
	program := ast.Program{
		Statements: []ast.Statement{},
	}

	for p.currToken.Type != token.EOF {
		stmt := p.parseStatement()
		//push non-nil statement
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		p.nextToken()
	}

	return &program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) Errors() []error {
	return p.errors
}

func (p *Parser) addError(err error) {
	p.errors = append(p.errors, err)
}
