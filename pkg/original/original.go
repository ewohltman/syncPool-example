package original

import (
	"encoding/json"

	"github.com/ewohltman/syncPool-example/pkg/base"
)

func UnmarshalObject(jsonBytes []byte) (*base.Object, error) {
	object := &base.Object{}

	err := json.Unmarshal(jsonBytes, object)
	if err != nil {
		return nil, err
	}

	return object, nil
}
