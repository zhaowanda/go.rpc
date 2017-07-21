package codec

import (
	"io"
	"github.com/zhaowanda/go.rpc/codec/codecx"
	"github.com/zhaowanda/go.rpc/core"
)

// NewProtobufServerCodec creates a protobuf ServerCodec by https://github.com/mars9/codec
func NewProtobufServerCodec(conn io.ReadWriteCloser) core.ServerCodec {
	return codecx.NewServerCodec(conn)
}

// NewProtobufClientCodec creates a protobuf ClientCodec by https://github.com/mars9/codec
func NewProtobufClientCodec(conn io.ReadWriteCloser) core.ClientCodec {
	return codecx.NewClientCodec(conn)
}
