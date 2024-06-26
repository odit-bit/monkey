package eval

import (
	"testing"

	"github.com/odit-bit/monkey/lexer"
	"github.com/odit-bit/monkey/object"
	"github.com/odit-bit/monkey/parser"
	"github.com/stretchr/testify/assert"
)

func Test_EvalInteger(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-25", -25},
		{"10 - 5", 5},
		{"10 - 5 * 2", 0},
		{"10 - 5 * 2 + 10", 10},
		{"10 - 5 * 2 + 10/5", 2},
	}

	for _, tc := range tests {
		obj := testEval(tc.input)
		testIntegerObject(t, obj, tc.expected)
	}
}

func Test_EvalBoolean(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
		{"5 == 5", true},
		{"(1 < 2) == true", true},
		{"(1 < 2) == false", false},
	}

	for _, tc := range tests {
		obj := testEval(tc.input)
		testBooleanObject(t, obj, tc.expected)
	}
}

func Test_BangOperator(t *testing.T) {
	tt := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!!true", true},
		{"!!5", true},
	}

	for _, tc := range tt {
		obj := testEval(tc.input)
		testBooleanObject(t, obj, tc.expected)
	}

}

func Test_Eval_IF_Else(t *testing.T) {
	tt := []struct {
		input    string
		expected any
	}{
		{"if (true) { 10 }", 10},
		{"if (false) { 10 }", nil},
		{"if (1) { 10 }", 10},
		{"if (1 < 2) { 10 }", 10},
		{"if (1 > 2) { 10 }", nil},
		{"if (1 > 2) { 10 } else { 20 }", 20},
		{"if (1 < 2) { 10 } else { 20 }", 10},
	}

	for _, tc := range tt {
		obj := testEval(tc.input)
		exp, ok := tc.expected.(int)
		if ok {
			testIntegerObject(t, obj, int64(exp))
		} else {
			testNullObject(t, obj)
		}
	}
}

func Test_Eval_Return(t *testing.T) {
	tt := []struct {
		input    string
		expected int64
	}{
		{"return 10", 10},
		{"10 * 20; return 3 * 5; 9*7;", 15},
		{`if (10 > 1) {
			if (10 > 1) {
				return 10;
			}
			return 1;
		}`,
			10,
		},
	}

	for _, tc := range tt {
		obj := testEval(tc.input)
		testIntegerObject(t, obj, tc.expected)
	}

}

func Test_Eval_Function(t *testing.T) {
	input := "fn (x) {x + 5;};"
	obj := testEval(input)
	assert.IsType(t, &object.Function{}, obj)
	fnObj := obj.(*object.Function)

	assert.Len(t, fnObj.Parameter, 1)

	inspect := "fn(x){\n(x + 5)\n}"
	assert.Equal(t, inspect, fnObj.Inspect())

	assert.Equal(t, "x", fnObj.Parameter[0].Value)
	assert.Equal(t, "(x + 5)", fnObj.Body.String())
}

func Test_Eval_Call_Function(t *testing.T) {
	tt := []struct {
		input    string
		expected int64
	}{
		{"let a = fn(x){x+5;}; a(5);", 10},
		{"let a = fn(x){let b = x+5; return b + x}; a(5);", 15},
	}

	for _, tc := range tt {
		obj := testEval(tc.input)
		testIntegerObject(t, obj, tc.expected)
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	env := object.NewEnvironment()
	prog := p.ParseProgram()
	return Eval(prog, env)

}

func testNullObject(t *testing.T, obj object.Object) {
	truth := obj != NULL
	assert.IsType(t, true, truth)
}

func testIntegerObject(t *testing.T, obj object.Object, expect int64) {
	if !assert.IsType(t, &object.Integer{}, obj) {
		t.Fail()
		return
	}
	intObj := obj.(*object.Integer)
	if !assert.Equal(t, expect, intObj.Value) {
		t.Log(obj.Inspect())
		t.Fail()
	}
}

func testBooleanObject(t *testing.T, obj object.Object, expect bool) {
	assert.IsType(t, &object.Boolean{}, obj)
	intObj := obj.(*object.Boolean)
	assert.Equal(t, expect, intObj.Value)
}
