package recodecJson

import (
	"github.com/json-iterator/go"
	"github.com/remicro/api/serialization"
)

func Decoder() serialization.Decoder {
	return coder{}
}

func Encoder() serialization.Encoder {
	return coder{}
}

type coder struct{}

func (coder) Encode(object interface{}) (data []byte, err error) {
	return jsoniter.Marshal(object)
}

func (coder) Decode(object interface{}, data []byte) (err error) {
	return jsoniter.Unmarshal(data, object)
}
