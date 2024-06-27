package eval

import (
	"github.com/odit-bit/monkey/ast"
	"github.com/odit-bit/monkey/object"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	//statement
	case *ast.Program:
		return evalProgram(node, env)

	case *ast.LetStatement:
		obj := Eval(node.Value, env)
		if isError(obj) {
			return obj
		}
		env.Set(node.Name.Value, obj)

	case *ast.ReturnStatement:
		return evalReturn(node, env)

	// expression
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)

	case *ast.Integer:
		return &object.Integer{Value: node.Value}

	case *ast.Boolean:
		return evalBool(node.Value)

	case *ast.String:
		return &object.String{Value: node.Value}

	case *ast.Prefix:
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalPrefix(node.Operator, right)

	case *ast.Infix:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalInfixOp(node.Operator, left, right)

	case *ast.IF:
		return evalIF(node, env)

	case *ast.BlockStatement:
		return evalBlockStatement(node, env)

	case *ast.Identifier:
		if obj, ok := env.Get(node.Value); !ok {
			return newError("identifier is not found: %s", node.Value)
		} else {
			return obj
		}
	case *ast.FunctionLiteral:
		return evalFunction(node, env)

	case *ast.CallExpr:
		function := Eval(node.Function, env)
		if isError(function) {
			return function
		}
		args := evalExpression(node.Args, env)

		return applyFunction(function, args, env)
	default:
		return newError("invalid node type: %T \n", node)
	}

	return nil
}

func evalExpression(node []ast.Expression, env *object.Environment) []object.Object {
	var result []object.Object
	for _, expr := range node {
		obj := Eval(expr, env)
		result = append(result, obj)
	}

	return result
}

func evalReturn(stmt *ast.ReturnStatement, env *object.Environment) object.Object {

	obj := Eval(stmt.ReturnValue, env)
	if isError(obj) {
		return obj
	}
	return &object.ReturnObj{Value: obj}

}

func evalProgram(prog *ast.Program, env *object.Environment) object.Object {
	var result object.Object
	for _, stmt := range prog.Statements {
		result = Eval(stmt, env)

		switch obj := result.(type) {
		case *object.ErrorObj:
			return obj
		case *object.ReturnObj:
			return obj.Value
		}

	}

	return result
}

func evalBlockStatement(block *ast.BlockStatement, env *object.Environment) object.Object {
	var result object.Object
	for _, stmt := range block.Statements {
		result = Eval(stmt, env)

		if result != nil {
			t := result.Type()
			if t == object.RETURN_OBJ || t == object.ERROR_OBJ {
				return result
			}
		}
	}

	return result
}

func evalBool(input bool) *object.Boolean {
	if input {
		return TRUE
	} else {
		return FALSE
	}
}

func evalIF(expr *ast.IF, env *object.Environment) object.Object {
	cond := Eval(expr.Condition, env)
	if isError(cond) {
		return cond
	}

	var truth bool
	switch cond {
	case NULL:
		truth = false
	case TRUE:
		truth = true
	case FALSE:
		truth = false
	default:
		truth = true
	}

	if truth {
		return Eval(expr.Consequence, env)
	} else if expr.Alternative != nil {
		return Eval(expr.Alternative, env)
	} else {
		return NULL
	}
}
