// Package original provides an example usage of unmarshalling JSON into a struct
// without using sync.Pool and allocating a new struct every time.
package original

import (
	"encoding/json"

	"github.com/ewohltman/syncPool-example/pkg/base"
)

// UnmarshalObject unmarshals jsonBytes into a *base.Object and returns the
// dereferenced base.Object. The *base.Object used is allocated each time
// UnmarshalObject is called.
func UnmarshalObject(jsonBytes []byte) (base.Object, error) {
	object := &base.Object{}

	err := json.Unmarshal(jsonBytes, object)
	if err != nil {
		return base.Object{}, err
	}

	return *object, nil
}
