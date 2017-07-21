package codec

import (
	"io"
	"github.com/zhaowanda/go.rpc/core"
	"github.com/zhaowanda/go.rpc/codec/protobuf"
	"github.com/zhaowanda/go.rpc/codec"
)

// NewProtobufServerCodec creates a protobuf ServerCodec by https://github.com/mars9/codec
func NewProtobufServerCodec(conn io.ReadWriteCloser) codec.ServerCodec {
	return protobuf.NewServerCodec(conn)
}

// NewProtobufClientCodec creates a protobuf ClientCodec by https://github.com/mars9/codec
func NewProtobufClientCodec(conn io.ReadWriteCloser) codec.ClientCodec {
	return protobuf.NewClientCodec(conn)
}
