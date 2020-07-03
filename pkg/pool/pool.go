// Package pool provides an example usage of unmarshalling JSON into a struct
// using sync.Pool.
package pool

import (
	"bytes"
	"encoding/json"
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

	defer bufferPool.Put(buffer)
	defer buffer.Reset()

	object := objectPool.Get().(*base.Object)

	defer objectPool.Put(object)
	defer object.Reset()

	contents := contentPool.Get().(*base.Contents)

	defer contentPool.Put(contents)
	defer contents.Reset()

	object.Contents = contents

	buffer.WriteString(jsonString)

	err := json.Unmarshal(buffer.Bytes(), object)
	if err != nil {
		return base.Object{}, err
	}

	return *object, nil
}
