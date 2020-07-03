// Package pool provides an example usage of unmarshalling JSON into a struct
// using sync.Pool.
package pool

import (
	"bytes"
	"encoding/json"
	"sync"

	"github.com/ewohltman/syncPool-example/pkg/base"
)

// Pools is a collection of *sync.Pool objects.
type Pools struct {
	Objects  *sync.Pool
	Contents *sync.Pool
	Buffers  *sync.Pool
}

// NewPools returns a new *Pools with the *sync.Pool objects initialized.
func NewPools() *Pools {
	return &Pools{
		Objects: &sync.Pool{
			New: func() interface{} {
				return &base.Object{}
			},
		},
		Contents: &sync.Pool{
			New: func() interface{} {
				return &base.Contents{}
			},
		},
		Buffers: &sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
	}
}

// UnmarshalObject writes jsonString into a *bytes.Buffer using sync.Pool and
// unmarshals the buffer into a *base.Object. It then returns the dereferenced
// base.Object. The *base.Object is also obtained using sync.Pool.
func (pools *Pools) UnmarshalObject(jsonString string) (base.Object, error) {
	object := pools.Objects.Get().(*base.Object)
	defer pools.Objects.Put(object)
	defer object.Reset()

	contents := pools.Contents.Get().(*base.Contents)
	defer pools.Contents.Put(contents)
	defer contents.Reset()

	object.Contents = contents

	buffer := pools.Buffers.Get().(*bytes.Buffer)
	defer pools.Buffers.Put(buffer)
	defer buffer.Reset()

	buffer.WriteString(jsonString)

	err := json.Unmarshal(buffer.Bytes(), object)
	if err != nil {
		return base.Object{}, err
	}

	return *object, nil
}
