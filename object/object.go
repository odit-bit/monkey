package object

import (
	"fmt"
	"strings"

	"github.com/odit-bit/monkey/ast"
)

type ObjectType string

const (
	INTEGER_OBJ  = "INTEGER"
	BOOLEAN_OBJ  = "BOOLEAN"
	NULL_OBJ     = "NULL"
	RETURN_OBJ   = "RETURN"
	ERROR_OBJ    = "ERROR"
	FUNCTION_OBJ = "FUNCTION"
	STRING_OBJ   = "STRING"

	BUILTIN_OBJ = "BUILTIN"
	ARRAY_OBJ   = "ARRAY"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Null struct{}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "null" }

//ERROR

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

//RETURN

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

//BUILTIN Function

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

//STRING

var _ Object = (*String)(nil)

type String struct {
	Value string
}

func (s *String) Inspect() string {
	return s.Value
}

func (s *String) Type() ObjectType {
	return STRING_OBJ
}

// INTEGER
var _ Object = (*Integer)(nil)

type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

//BOOLEAN

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

//ARRAY

var _ Object = (*Array)(nil)

type Array struct {
	Elements []Object
}

func (a *Array) Inspect() string {
	el := []string{}
	for _, obj := range a.Elements {
		el = append(el, obj.Inspect())
	}

	return fmt.Sprintf("[%s]", strings.Join(el, ", "))
}

func (a *Array) Type() ObjectType {
	return ARRAY_OBJ
}
