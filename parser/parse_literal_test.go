package parser

import (
	"testing"

	"github.com/odit-bit/monkey/ast"
	"github.com/odit-bit/monkey/lexer"
	"github.com/stretchr/testify/assert"
)

func Test_Parse_integerLiteral(t *testing.T) {
	var input = "5;"

	l := lexer.New(input)
	p := New(l)

	prog := p.ParseProgram()
	if !assert.Len(t, prog.Statements, 1) {
		checkError(t, p)
		t.FailNow()
	}

	assert.IsType(t, &ast.ExpressionStatement{}, prog.Statements[0])
	expr := prog.Statements[0].(*ast.ExpressionStatement)
	testLiteralExpression(t, expr.Expression, 5)
}

func Test_Parse_String_Literal(t *testing.T) {
	var input = `"i am string";`

	l := lexer.New(input)
	p := New(l)

	prog := p.ParseProgram()
	checkError(t, p)

	assert.IsType(t, &ast.ExpressionStatement{}, prog.Statements[0])
	expr := prog.Statements[0].(*ast.ExpressionStatement)

	assert.IsType(t, &ast.String{}, expr.Expression)
	strExpr := expr.Expression.(*ast.String)

	assert.Equal(t, "i am string", strExpr.Value)
}

func testLiteralExpression(t *testing.T, expr ast.Expression, expected any) {

	switch v := expected.(type) {
	case int:
		testIntegerLiteral(t, expr, int64(v))
	case int64:
		testIntegerLiteral(t, expr, v)
	case string:
		testIdentifier(t, expr, v)
	case bool:
		testBoolean(t, expr, v)
	default:
		t.Errorf("invalid type of expr . got= %T", expr)
		return
	}

}
