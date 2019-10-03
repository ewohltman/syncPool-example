// Package original provides an example usage of unmarshalling JSON into a struct
// without using sync.Pool and allocating a new struct every time.
package original

import (
	"bytes"
	"encoding/json"

	"github.com/ewohltman/syncPool-example/pkg/base"
)

// UnmarshalObject writes jsonString into a *bytes.Buffer allocated each time
// UnmarshalObject is called and unmarshals the buffer into a *base.Object. It
// then returns the dereferenced base.Object. The *base.Object used is also
// allocated each time UnmarshalObject is called.
func UnmarshalObject(jsonString string) (base.Object, error) {
	buffer := new(bytes.Buffer)
	buffer.WriteString(jsonString)

	object := &base.Object{}

	err := json.Unmarshal(buffer.Bytes(), object)
	if err != nil {
		return base.Object{}, err
	}

	return *object, nil
}
