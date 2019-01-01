// object/object.go
//
// defines object system for evaluating values

package object

import "fmt"

// ObjectType defines the type
type ObjectType string

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
)

// Object is a thing
type Object interface {
	Type() ObjectType
	Inspect() string
}

// Integer are numbers
type Integer struct {
	Value int64
}

// Inspect returns the value
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

// Type returns the type
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

// Boolean are true or false
type Boolean struct {
	Value bool
}

// Inspect ...
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

// Type ...
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

// Null object
type Null struct{}

// Inspect ...
func (n *Null) Inspect() string { return "null" }

// Type ...
func (n *Null) Type() ObjectType { return NULL_OBJ }

// ReturnValue is a return value
type ReturnValue struct {
	Value Object
}

// Type ...
func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }

// Inspect ...
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }

// Error has a message
type Error struct {
	Message string
}

// Type ...
func (e *Error) Type() ObjectType { return ERROR_OBJ }

// Inspect ...
func (e *Error) Inspect() string { return "ERROR: " + e.Message }
