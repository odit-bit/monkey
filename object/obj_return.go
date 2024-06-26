package object

var _ Object = (*ReturnObj)(nil)

type ReturnObj struct {
	Value Object
}

// Inspect implements Object.
func (r *ReturnObj) Inspect() string {
	return r.Value.Inspect()
}

// Type implements Object.
func (r *ReturnObj) Type() ObjectType {
	return RETURN_OBJ
}
