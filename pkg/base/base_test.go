package base

import "testing"

func TestObject_String(t *testing.T) {
	object := &Object{
		Name: "",
		Contents: &Contents{
			Value: 0,
		},
	}

	actual := object.String()
	expected := "{Name: Contents:{Value:0}}"

	if actual != expected {
		t.Errorf(
			"Unexpected Object string. Got: %s, Expected: %s",
			actual,
			expected,
		)
	}
}
