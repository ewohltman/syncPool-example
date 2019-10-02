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

// Contents is an example struct embedded within the Object struct.
type Contents struct {
	Value int `json:"value"`
}

// String satisfies the fmt.Stringer interface.
func (object *Object) String() string {
	return fmt.Sprintf(
		"{Name:%s Contents:{Value:%d}}",
		object.Name,
		object.Contents.Value,
	)
}

// Reset resets the values of object to default values.
func (object *Object) Reset() {
	object.Name = ""
	object.Contents.Value = 0
}
