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
	switch t := function.(type) {
	case *object.Function:
		encloseEnv := object.NewEnclosed(env)
		for i, v := range t.Parameter {
			encloseEnv.Set(v.Value, args[i])
		}
		result := Eval(t.Body, encloseEnv)
		if ret, ok := result.(*object.ReturnObj); ok {
			return ret.Value
		}
		return result
	case *object.Builtin:
		return t.Fn(args...)

	default:
		return newError("not function type: %s", function.Type())
	}

}
