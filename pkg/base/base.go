// Package base provides common items used by other packages to standardize the
// tests and benchmarks.
package base

import "fmt"

// ObjectJSON is the zero-value JSON representation of Object.
const ObjectJSON = `{"name":"","contents":{"value":0}}`

// Object is an example struct containing a string and pointer to another
// struct.
type Object struct {
	Name     string    `json:"name"`
	Contents *Contents `json:"contents"`
}

// String satisfies the fmt.Stringer interface for Object.
func (object *Object) String() string {
	return fmt.Sprintf(
		"{Name:%s %s}",
		object.Name,
		object.Contents,
	)
}

// Reset resets object to default values.
func (object *Object) Reset() {
	object.Name = ""
	object.Contents = nil
}

// Contents is an example struct embedded within the Object struct.
type Contents struct {
	Value int `json:"value"`
}

// String satisfies the fmt.Stringer interface for Contents.
func (contents *Contents) String() string {
	return fmt.Sprintf(
		"Contents:{Value:%d}",
		contents.Value,
	)
}

// Reset resets contents to the default value.
func (contents *Contents) Reset() {
	contents.Value = 0
}
