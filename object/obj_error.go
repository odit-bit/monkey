package object

var _ Object = (*ErrorObj)(nil)

type ErrorObj struct {
	Value string
}

// Inspect implements Object.
func (e *ErrorObj) Inspect() string {
	return "ERROR: " + e.Value
}

// Type implements Object.
func (e *ErrorObj) Type() ObjectType {
	return ERROR_OBJ
}
