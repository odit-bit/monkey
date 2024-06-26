package object

import "fmt"

// INTEGER
var _ Object = (*Integer)(nil)

type Integer struct {
	Value int64
}

// Type implements Object.
func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

// Inspect implements ObjectType.
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}
