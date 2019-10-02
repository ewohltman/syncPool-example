package base

import "fmt"

const ObjectJSON = `{"name":"","contents":{"value":0}}`

type Object struct {
	Name     string    `json:"name"`
	Contents *Contents `json:"contents"`
}

type Contents struct {
	Value int `json:"value"`
}

func (object *Object) String() string {
	return fmt.Sprintf(
		"{Name:%s Contents:{Value:%d}}",
		object.Name,
		object.Contents.Value,
	)
}

func (object *Object) Reset() {
	object.Name = ""
	object.Contents.Value = 0
}
