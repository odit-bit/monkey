package eval

import "github.com/odit-bit/monkey/object"

func evalPrefix(op string, obj object.Object) object.Object {
	switch op {
	case "!":
		return evalBangOp(obj)
	case "-":
		return evalMinusOp(obj)
	default:
		//return object error
		return newError("unknown prefix operator: %s", op)
	}
}

func evalMinusOp(obj object.Object) object.Object {
	if obj.Type() != object.INTEGER_OBJ {
		return NULL
	}

	val := obj.(*object.Integer).Value
	return &object.Integer{Value: -val}
}

func evalBangOp(obj object.Object) object.Object {
	//bang (!) return opposite
	switch obj {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}
