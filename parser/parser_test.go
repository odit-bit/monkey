package parser

import (
	"fmt"
	"testing"

	"github.com/odit-bit/monkey/ast"
	"github.com/odit-bit/monkey/lexer"
	"github.com/stretchr/testify/assert"
)

func Test_Parse_letStatement(t *testing.T) {

	var input = `let x = 5;
let y = 10;
let foobar = 838383;`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	// assert.Equal(t, nil, program)

	if !assert.Len(t, program.Statements, 3) {
		//check silent error
		checkError(t, p)
		t.FailNow()
	}

	tt := []struct {
		expIdent string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tc := range tt {
		stmt := program.Statements[i]
		testLetStatement(t, stmt, tc.expIdent)
	}

}

func testLetStatement(t *testing.T, stmt ast.Statement, name string) {
	if stmt.TokenLiteral() != "let" {
		t.Fatalf("not let, got %v \n", stmt.TokenLiteral())
	}

	switch stmt := stmt.(type) {
	case *ast.LetStatement:
		assert.Equal(t, name, stmt.Name.Value)
		assert.Equal(t, name, stmt.Name.TokenLiteral())
	default:
		t.Fatalf("unknown type %t \n", stmt)
	}
}

func Test_Parse_returnStatement(t *testing.T) {

	var input = `return 5;
return 10;
return 838383;`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	// assert.Equal(t, nil, program)

	if !assert.Len(t, program.Statements, 3) {
		//check silent error
		checkError(t, p)
		t.FailNow()
	}

	for _, stmt := range program.Statements {
		assert.Equal(t, "return", stmt.TokenLiteral())
	}
}

func Test_Parse_Expression(t *testing.T) {
	var input = `foobar;`

	l := lexer.New(input)
	p := New(l)

	prog := p.ParseProgram()

	assert.Len(t, prog.Statements, 1)

	assert.IsType(t, &ast.ExpressionStatement{}, prog.Statements[0])
	stmt := prog.Statements[0].(*ast.ExpressionStatement)
	testIdentifier(t, stmt.Expression, "foobar")

}

func Test_Parse_PrefixExpression(t *testing.T) {
	var tt = []struct {
		input    string
		operator string
		value    int64
	}{
		{"-15;", "-", 15},
		{"!15;", "!", 15},
	}

	for _, tc := range tt {
		l := lexer.New(tc.input)
		p := New(l)

		prog := p.ParseProgram()

		//stmt should expression statement type
		stmt := prog.Statements[0].(*ast.ExpressionStatement)
		assert.IsType(t, &ast.ExpressionStatement{}, stmt)

		// should prefix expression
		assert.IsType(t, &ast.Prefix{}, stmt.Expression)
		expr := stmt.Expression.(*ast.Prefix)

		//operator value  should same
		assert.Equal(t, tc.operator, expr.Operator)

		//right expression should integerLiteral
		testLiteralExpression(t, expr.Right, tc.value)
	}
}

func Test_Parse_infixExpression(t *testing.T) {
	var tt = []struct {
		input    string
		left     any
		operator string
		right    any
	}{
		{"5 - 15;", 5, "-", 15},
		{"5 == 15;", 5, "==", 15},
		{"5 / 5;", 5, "/", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 != 5;", 5, "!=", 5},
		{"5 + 5;", 5, "+", 5},
		{"true == true", true, "==", true},
		{"true != false", true, "!=", false},
		{"false == false", false, "==", false},
		// {"5 >= 5;", 5, ">=", 5},
	}

	for _, tc := range tt {
		l := lexer.New(tc.input)
		p := New(l)

		prog := p.ParseProgram()
		checkError(t, p)

		// should expressionStatement
		assert.IsType(t, &ast.ExpressionStatement{}, prog.Statements[0])
		stmt := prog.Statements[0].(*ast.ExpressionStatement)

		//should infix
		testInfixExpression(t, stmt.Expression, tc.left, tc.right, tc.operator)
	}

}

func testInfixExpression(t *testing.T, expr ast.Expression, left any, right any, operator string) {
	assert.IsType(t, &ast.Infix{}, expr)
	infixExpr := expr.(*ast.Infix)

	// check expression
	testLiteralExpression(t, infixExpr.Left, left)
	assert.Equal(t, infixExpr.Operator, operator)
	testLiteralExpression(t, infixExpr.Right, right)
}

func Test_precedence(t *testing.T) {
	var tt = []struct {
		input  string
		output string
	}{
		{"5 + 5 * 5", "(5 + (5 * 5))"},
		{"-5 * 5", "((-5) * 5)"},
		{"true", "true"},
		{"false", "false"},
		{"3 > 5 == false", "((3 > 5) == false)"},
		{"1 + (2 + 3) + 4", "((1 + (2 + 3)) + 4)"},
	}

	for _, tc := range tt {
		l := lexer.New(tc.input)
		p := New(l)

		prog := p.ParseProgram()
		checkError(t, p)

		stmt := prog.Statements[0]
		assert.Equal(t, tc.output, stmt.String())
	}

}

func Test_Parse_IF(t *testing.T) {
	var input = "if (x > 5) {x}"

	l := lexer.New(input)
	p := New(l)
	prog := p.ParseProgram()
	checkError(t, p)

	assert.IsType(t, &ast.ExpressionStatement{}, prog.Statements[0])
	exprStmt := prog.Statements[0].(*ast.ExpressionStatement)

	assert.IsType(t, &ast.IF{}, exprStmt.Expression)
	ifExpr := exprStmt.Expression.(*ast.IF)

	assert.Equal(t, "if", ifExpr.TokenLiteral())
	testInfixExpression(t, ifExpr.Condition, "x", 5, ">")

	assert.Len(t, ifExpr.Consequence.Statements, 1)
	assert.IsType(t, &ast.ExpressionStatement{}, ifExpr.Consequence.Statements[0])
	expr := ifExpr.Consequence.Statements[0].(*ast.ExpressionStatement)
	testIdentifier(t, expr.Expression, "x")

}

func Test_Parse_IF_ELSE(t *testing.T) {
	var input = "if (x > 5) {x} else {y}"

	l := lexer.New(input)
	p := New(l)
	prog := p.ParseProgram()
	checkError(t, p)

	assert.IsType(t, &ast.ExpressionStatement{}, prog.Statements[0])
	exprStmt := prog.Statements[0].(*ast.ExpressionStatement)

	assert.IsType(t, &ast.IF{}, exprStmt.Expression)
	ifExpr := exprStmt.Expression.(*ast.IF)

	assert.Equal(t, "if", ifExpr.TokenLiteral())
	testInfixExpression(t, ifExpr.Condition, "x", 5, ">")

	assert.Len(t, ifExpr.Consequence.Statements, 1)
	assert.IsType(t, &ast.ExpressionStatement{}, ifExpr.Consequence.Statements[0])
	consExpr := ifExpr.Consequence.Statements[0].(*ast.ExpressionStatement)
	testIdentifier(t, consExpr.Expression, "x")

	elseExpr := ifExpr.Alternative.Statements[0].(*ast.ExpressionStatement)
	testIdentifier(t, elseExpr.Expression, "y")

}

func Test_Parse_Function(t *testing.T) {
	var input = "fn (x,y){x + y;}"
	l := lexer.New(input)
	p := New(l)
	prog := p.ParseProgram()
	checkError(t, p)

	assert.Len(t, prog.Statements, 1)
	assert.IsType(t, &ast.ExpressionStatement{}, prog.Statements[0])
	expr := prog.Statements[0].(*ast.ExpressionStatement)

	assert.IsType(t, &ast.FunctionLiteral{}, expr.Expression)
	fnExpr := expr.Expression.(*ast.FunctionLiteral)

	//check param
	testLiteralExpression(t, fnExpr.Parameters[0], "x")
	testLiteralExpression(t, fnExpr.Parameters[1], "y")

	//check body
	assert.IsType(t, &ast.ExpressionStatement{}, fnExpr.Body.Statements[0])
	bodyStmt := fnExpr.Body.Statements[0].(*ast.ExpressionStatement)

	assert.IsType(t, &ast.Infix{}, bodyStmt.Expression)
	infixExpr := bodyStmt.Expression.(*ast.Infix)

	testInfixExpression(t, infixExpr, "x", "y", "+")

}

func Test_Parse_Call(t *testing.T) {
	var input = "add(1, 2 * 5, x + 5);"

	l := lexer.New(input)
	p := New(l)
	prog := p.ParseProgram()
	_ = prog
	checkError(t, p)
}

func Test_Parse_comment(t *testing.T) {
	var input = "// this is comment;"

	l := lexer.New(input)
	p := New(l)
	prog := p.ParseProgram()
	_ = prog
	checkError(t, p)
}

func testBoolean(t *testing.T, expr ast.Expression, value bool) {
	assert.IsType(t, &ast.Boolean{}, expr)
	boolExpr := expr.(*ast.Boolean)

	assert.Equal(t, value, boolExpr.Value)
	assert.Equal(t, fmt.Sprintf("%t", value), boolExpr.TokenLiteral())
}

func testIntegerLiteral(t *testing.T, stmt ast.Expression, value int64) {

	assert.IsType(t, &ast.Integer{}, stmt)

	intLiteral := stmt.(*ast.Integer)
	assert.Equal(t, value, intLiteral.Value)
	assert.Equal(t, fmt.Sprintf("%d", value), intLiteral.TokenLiteral())
}

func testIdentifier(t *testing.T, expr ast.Expression, value string) {

	ident := expr.(*ast.Identifier)
	assert.IsType(t, &ast.Identifier{}, ident)
	assert.Equal(t, value, ident.Value)
	assert.Equal(t, value, ident.TokenLiteral())
}

func checkError(t *testing.T, p *Parser) {
	if !assert.Len(t, p.errors, 0) {
		for _, err := range p.errors {
			t.Logf("parserErr: %v \n", err)
		}
		t.FailNow()
	}
}
