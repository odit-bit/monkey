package object

import (
	"fmt"
	"strings"

	"github.com/odit-bit/monkey/ast"
)

// Function

var _ Object = (*Function)(nil)

type Function struct {
	Parameter []*ast.Identifier
	Body      *ast.BlockStatement
	Env       *Environment
}

func (f *Function) Inspect() string {
	var params []string
	for _, ident := range f.Parameter {
		params = append(params, ident.String())
	}

	return fmt.Sprintf("fn(%s){\n%s\n}", strings.Join(params, ", "), f.Body.String())
}

func (f *Function) Type() ObjectType {
	return FUNCTION_OBJ
}

//BUILTIN OBJECT

type BuiltinFunc func(a ...Object) Object

var _ Object = (*Builtin)(nil)

type Builtin struct {
	Fn BuiltinFunc
}

// Inspect implements Object.
func (b *Builtin) Inspect() string {
	return "BUILTIN FUNCTION"
}

// Type implements Object.
func (b *Builtin) Type() ObjectType {
	return BUILTIN_OBJ
}
