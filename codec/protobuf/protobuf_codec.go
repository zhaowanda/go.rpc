package protobuf

import (
	"io"
	"github.com/zhaowanda/go.rpc/codec"
)

// NewProtobufServerCodec creates a protobuf ServerCodec by https://github.com/mars9/codec
func NewProtobufServerCodec(conn io.ReadWriteCloser) codec.ServerCodec {
	return NewServerCodec(conn)
}

// NewProtobufClientCodec creates a protobuf ClientCodec by https://github.com/mars9/codec
func NewProtobufClientCodec(conn io.ReadWriteCloser) codec.ClientCodec {
	return NewClientCodec(conn)
}
