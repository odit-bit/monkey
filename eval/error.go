package eval

import (
	"fmt"

	"github.com/odit-bit/monkey/object"
)

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	} else {
		return false
	}
}

func newError(format string, msg ...any) object.Object {
	return &object.ErrorObj{
		Value: fmt.Sprintf(format, msg...),
	}
}
