package pool

import (
	"testing"

	"github.com/ewohltman/syncPool-example/pkg/base"
)

func TestUnmarshalObject(t *testing.T) {
	jsonBytes := []byte(base.ObjectJSON)

	object, err := UnmarshalObject(jsonBytes)
	if err != nil {
		t.Fatal(err)
	}

	if object == nil {
		t.Fatal("Unexpected nil *base.Object")
	}

	if object.Contents == nil {
		t.Fatal("Unexpected nil *base.Object.Contents")
	}

	actualName := object.Name
	expectedName := ""

	if actualName != expectedName {
		t.Errorf(
			"Unexpected *base.Object.Name. Got: %s, Expected: %s",
			actualName,
			expectedName,
		)
	}

	actualValue := object.Contents.Value
	expectedValue := 0

	if actualValue != expectedValue {
		t.Errorf(
			"Unexpected *base.Object.Contents.Value. Got: %d, Expected: %d",
			actualValue,
			expectedValue,
		)
	}
}

func BenchmarkUnmarshalObject(b *testing.B) {
	jsonBytes := []byte(base.ObjectJSON)

	for n := 0; n < b.N; n++ {
		_, _ = UnmarshalObject(jsonBytes)
	}
}
