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

func TestObject_Reset(t *testing.T) {
	object := &Object{
		Name: "testObject",
		Contents: &Contents{
			Value: 100,
		},
	}

	object.Reset()

	actualName := object.Name
	expectedName := ""

	if actualName != expectedName {
		t.Errorf(
			"Unexpected *Object.Name. Got: %s, Expected: %s",
			actualName,
			expectedName,
		)
	}

	actualValue := object.Contents.Value
	expectedValue := 0

	if actualValue != expectedValue {
		t.Errorf(
			"Unexpected *Object.Contents.Value. Got: %d, Expected: %d",
			actualValue,
			expectedValue,
		)
	}
}
