package grpc

import (
	"encoding/json"

	"github.com/golang/protobuf/proto"
	"github.com/micro/go-micro/codec"
	"github.com/micro/go-micro/codec/jsonrpc"
	"github.com/micro/go-micro/codec/protorpc"
	"github.com/micro/grpc-go/encoding"
)

type jsonCodec struct{}
type protoCodec struct{}
type bytesCodec struct{}

var (
	defaultGRPCCodecs = map[string]encoding.Codec{
		"application/json":         jsonCodec{},
		"application/proto":        protoCodec{},
		"application/protobuf":     protoCodec{},
		"application/octet-stream": protoCodec{},
		"application/grpc":         protoCodec{},
		"application/grpc+json":    jsonCodec{},
		"application/grpc+proto":   protoCodec{},
		"application/grpc+bytes":   protoCodec{},
	}

	defaultRPCCodecs = map[string]codec.NewCodec{
		"application/json":         jsonrpc.NewCodec,
		"application/json-rpc":     jsonrpc.NewCodec,
		"application/protobuf":     protorpc.NewCodec,
		"application/proto-rpc":    protorpc.NewCodec,
		"application/octet-stream": protorpc.NewCodec,
	}
)

func (protoCodec) Marshal(v interface{}) ([]byte, error) {
	return proto.Marshal(v.(proto.Message))
}

func (protoCodec) Unmarshal(data []byte, v interface{}) error {
	return proto.Unmarshal(data, v.(proto.Message))
}

func (protoCodec) Name() string {
	return "proto"
}

func (jsonCodec) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (jsonCodec) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (jsonCodec) Name() string {
	return "json"
}
