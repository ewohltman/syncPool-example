package base_test

import (
	"testing"

	"github.com/ewohltman/syncPool-example/pkg/base"
)

func TestObject_String(t *testing.T) {
	object := &base.Object{
		Name: "",
		Contents: &base.Contents{
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
	object := &base.Object{
		Name: "testObject",
		Contents: &base.Contents{
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

	if object.Contents != nil {
		t.Errorf(
			"Unexpected non-nil *Object.Contents. Got: %+v, Expected: nil",
			*object.Contents,
		)
	}
}

func TestContents_String(t *testing.T) {
	contents := &base.Contents{
		Value: 0,
	}

	actual := contents.String()
	expected := "Contents:{Value:0}"

	if actual != expected {
		t.Errorf(
			"Unexpected Contents string. Got: %s, Expected: %s",
			actual,
			expected,
		)
	}
}

func TestContents_Reset(t *testing.T) {
	contents := &base.Contents{
		Value: 100,
	}

	contents.Reset()

	actual := contents.Value
	expected := 0

	if actual != expected {
		t.Errorf(
			"Unexpected *Contents.Value. Got: %d, Expected: %d",
			actual,
			expected,
		)
	}
}
