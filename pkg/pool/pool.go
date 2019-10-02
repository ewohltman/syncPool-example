// Package pool provides an example usage of unmarshalling JSON into a struct
// using sync.Pool.
package pool

import (
	"encoding/json"
	"sync"

	"github.com/ewohltman/syncPool-example/pkg/base"
)

//nolint:gochecknoglobals
var objectPool = sync.Pool{
	New: func() interface{} {
		return &base.Object{
			Name: "",
			Contents: &base.Contents{
				Value: 0,
			},
		}
	},
}

// UnmarshalObject unmarshals jsonBytes into a *base.Object and returns the
// dereferenced base.Object. The *base.Object is obtained using sync.Pool.
func UnmarshalObject(jsonBytes []byte) (base.Object, error) {
	object := objectPool.Get().(*base.Object)

	defer func() {
		object.Reset()
		objectPool.Put(object)
	}()

	err := json.Unmarshal(jsonBytes, object)
	if err != nil {
		return base.Object{}, err
	}

	return *object, nil
}
