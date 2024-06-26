package eval

import (
	"github.com/odit-bit/monkey/ast"
	"github.com/odit-bit/monkey/object"
)

func evalFunction(node *ast.FunctionLiteral, env *object.Environment) object.Object {
	return &object.Function{
		Parameter: node.Parameters,
		Body:      node.Body,
		Env:       env,
	}
}

func applyFunction(function object.Object, args []object.Object, env *object.Environment) object.Object {
	fn, ok := function.(*object.Function)
	if !ok {
		return newError("not function type: %s", function.Type())
	}

	encloseEnv := object.NewEnclosed(env)
	for i, v := range fn.Parameter {
		encloseEnv.Set(v.Value, args[i])
	}

	obj := Eval(fn.Body, encloseEnv)

	if ret, ok := obj.(*object.ReturnObj); ok {
		return ret.Value
	}
	return obj
}
