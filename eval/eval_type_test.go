package eval

import (
	"testing"

	"github.com/odit-bit/monkey/object"
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

func Test_EvalString(t *testing.T) {
	tt := []struct {
		input    string
		expected string
	}{
		{`"i am string 12345";`, "i am string 12345"},
		{`"hello" + " " + "world"; `, "hello world"},
	}

	for _, tc := range tt {
		obj := testEval(tc.input)
		testLiteralObject(t, obj, tc.expected)
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

func Test_Eval_Array(t *testing.T) {
	tt := []struct {
		input    string
		expected string
	}{
		{`["hello", "world", 2000 + 24];`, "[hello, world, 2024]"},
	}

	for _, tc := range tt {
		obj := testEval(tc.input)
		assert.Equal(t, tc.expected, obj.Inspect())
	}
}

func testLiteralObject(t *testing.T, obj object.Object, expect any) {
	switch val := expect.(type) {
	case string:
		testStringObject(t, obj, val)
	case bool:
		testBooleanObject(t, obj, val)
	case int64:
		testIntegerObject(t, obj, val)
	default:
		t.Fatalf("invalide expected value type %T ", val)
	}
}

func testStringObject(t *testing.T, obj object.Object, expect string) {
	if !assert.IsType(t, &object.String{}, obj) {
		t.Log(obj.Inspect())
		t.Fail()
		return
	}

	strObj := obj.(*object.String)
	if !assert.Equal(t, expect, strObj.Value) {
		t.Log(obj.Inspect())
		t.Fail()
	}

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
	boolObj := obj.(*object.Boolean)
	assert.Equal(t, expect, boolObj.Value)
}
