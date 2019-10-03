// Package pool provides an example usage of unmarshalling JSON into a struct
// using sync.Pool.
package pool

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/ewohltman/syncPool-example/pkg/base"
)

//nolint:gochecknoglobals
var bufferPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

//nolint:gochecknoglobals
var objectPool = sync.Pool{
	New: func() interface{} {
		return &base.Object{
			Name:     "",
			Contents: nil,
		}
	},
}

//nolint:gochecknoglobals
var contentPool = sync.Pool{
	New: func() interface{} {
		return &base.Contents{
			Value: 0,
		}
	},
}

// UnmarshalObject writes jsonString into a *bytes.Buffer using sync.Pool and
// unmarshals the buffer into a *base.Object. It then returns the dereferenced
// base.Object. The *base.Object is also obtained using sync.Pool.
func UnmarshalObject(jsonString string) (base.Object, error) {
	buffer := bufferPool.Get().(*bytes.Buffer)
	buffer.WriteString(jsonString)

	object, ok := objectPool.Get().(*base.Object)
	if !ok {
		return base.Object{}, fmt.Errorf("unable to type assert *base.Object from sync.Pool")
	}

	contents, ok := contentPool.Get().(*base.Contents)
	if !ok {
		return base.Object{}, fmt.Errorf("unable to type assert *base.Contents from sync.Pool")
	}

	object.Contents = contents

	defer func() {
		buffer.Reset()
		object.Reset()
		contents.Reset()

		bufferPool.Put(buffer)
		objectPool.Put(object)
		contentPool.Put(contents)
	}()

	err := json.Unmarshal(buffer.Bytes(), object)
	if err != nil {
		return base.Object{}, err
	}

	return *object, nil
}
