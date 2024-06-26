package eval

import "github.com/odit-bit/monkey/object"

func evalInfixOp(op string, left, right object.Object) object.Object {

	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalInfixInteger(op, left, right)

	case op == "==":
		return evalBool(left == right)
	case op == "!=":
		return evalBool(left != right)
	default:
		return newError("unknown infix operator: %s", op)
	}

}

func evalInfixInteger(op string, left, right object.Object) object.Object {
	l := left.(*object.Integer)
	r := right.(*object.Integer)
	var val int64
	switch op {
	case "-":
		val = int64(l.Value - r.Value)

	case "+":
		val = int64(l.Value + r.Value)

	case "/":
		val = int64(l.Value / r.Value)

	case "*":
		val = int64(l.Value * r.Value)

	case "<":
		return evalBool(l.Value < r.Value)

	case ">":
		return evalBool(l.Value > r.Value)
	case "==":
		return evalBool(l.Value == r.Value)
	case "!=":
		return evalBool(l.Value != r.Value)

	default:
		return NULL
	}

	return &object.Integer{Value: val}
}
