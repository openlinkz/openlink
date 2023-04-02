package json

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/openlinkz/openlink/pkg/encoding"
)

const Name = "json"

func init() {
	encoding.RegisterCodec(codec{})
}

// codec is a Codec implementation with protobuf. It is the default codec for gRPC.
type codec struct{}

func (codec) Marshal(v interface{}) ([]byte, error) {
	return jsoniter.Marshal(v)
}

func (codec) Unmarshal(data []byte, v interface{}) error {
	return jsoniter.Unmarshal(data, v)
}

func (codec) Name() string {
	return Name
}
