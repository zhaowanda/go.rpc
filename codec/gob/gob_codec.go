package gob

import (
	"io"
	"github.com/zhaowanda/go.rpc/codec/gob"
	"github.com/zhaowanda/go.rpc/codec"
)

// NewProtobufServerCodec creates a protobuf ServerCodec by https://github.com/mars9/codec
func NewGobrCodec(conn io.ReadWriteCloser) codec.ServerCodec {
	return gob.NewServerCodec(conn)
}

// NewProtobufClientCodec creates a protobuf ClientCodec by https://github.com/mars9/codec
func NewGobClientCodec(conn io.ReadWriteCloser) codec.ClientCodec {
	return gob.NewClientCodec(conn)
}
