package eval

import (
	"fmt"

	"github.com/odit-bit/monkey/object"
)

var builtinMap = map[string]*object.Builtin{
	"puts": {
		Fn: func(a ...object.Object) object.Object {
			for _, arg := range a {
				fmt.Println(arg.Inspect())
			}
			return NULL
		},
	},

	"len": {
		Fn: func(a ...object.Object) object.Object {
			if len(a) != 1 {
				return newError("to many argument got %d, want 1", len(a))
			}

			switch arg := a[0].(type) {
			case *object.String:
				return &object.Integer{
					Value: int64(len(arg.Value)),
				}

			default:
				return newError("argument typs is not string, got %s ", arg.Type())
			}
		},
	},
}
