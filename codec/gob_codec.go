package codec

import (
	"io"
	"github.com/zhaowanda/go.rpc/codec/gob"
)

// NewProtobufServerCodec creates a protobuf ServerCodec by https://github.com/mars9/codec
func NewGobrCodec(conn io.ReadWriteCloser) ServerCodec {
	return gob.NewServerCodec(conn)
}

// NewProtobufClientCodec creates a protobuf ClientCodec by https://github.com/mars9/codec
func NewGobClientCodec(conn io.ReadWriteCloser) ClientCodec {
	return gob.NewClientCodec(conn)
}
