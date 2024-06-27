package object

var _ Object = (*String)(nil)

type String struct {
	Value string
}

// Inspect implements Object.
func (s *String) Inspect() string {
	return s.Value
}

// Type implements Object.
func (s *String) Type() ObjectType {
	return STRING_OBJ
}
