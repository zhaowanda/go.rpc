package codec

import (
	"io"
	"github.com/zhaowanda/go.rpc/core"
	"go.rpc/codec/protobuf"
)

// NewProtobufServerCodec creates a protobuf ServerCodec by https://github.com/mars9/codec
func NewProtobufServerCodec(conn io.ReadWriteCloser) core.ServerCodec {
	return protobuf.NewServerCodec(conn)
}

// NewProtobufClientCodec creates a protobuf ClientCodec by https://github.com/mars9/codec
func NewProtobufClientCodec(conn io.ReadWriteCloser) core.ClientCodec {
	return protobuf.NewClientCodec(conn)
}
