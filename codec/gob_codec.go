package codec

import (
	"io"
	"github.com/zhaowanda/go.rpc/core"
	"go.rpc/codec/gob"
)

// NewProtobufServerCodec creates a protobuf ServerCodec by https://github.com/mars9/codec
func NewGobServerCodec(conn io.ReadWriteCloser) core.ServerCodec {
	return gob.NewServerCodec(conn)
}

// NewProtobufClientCodec creates a protobuf ClientCodec by https://github.com/mars9/codec
func NewGobClientCodec(conn io.ReadWriteCloser) core.ClientCodec {
	return gob.NewClientCodec(conn)
}
