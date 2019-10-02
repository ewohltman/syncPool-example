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

func UnmarshalObject(jsonBytes []byte) (*base.Object, error) {
	object := objectPool.Get().(*base.Object)

	defer func() {
		object.Reset()
		objectPool.Put(object)
	}()

	err := json.Unmarshal(jsonBytes, object)
	if err != nil {
		return nil, err
	}

	return object, nil
}
