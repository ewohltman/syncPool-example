package pool_test

import (
	"testing"

	"github.com/ewohltman/syncPool-example/pkg/base"
	"github.com/ewohltman/syncPool-example/pkg/pool"
)

func TestUnmarshalObject(t *testing.T) {
	object, err := pool.UnmarshalObject(base.ObjectJSON)
	if err != nil {
		t.Fatal(err)
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
	for n := 0; n < b.N; n++ {
		_, _ = pool.UnmarshalObject(base.ObjectJSON)
	}
}
