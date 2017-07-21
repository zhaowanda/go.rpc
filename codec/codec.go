package codec

import (
	context "golang.org/x/net/context"
	"go.rpc/gorpc"
)


type ServerCodec interface {
	ReadRequestHeader(context.Context, *gorpc.Request) error
	ReadRequestBody(context.Context, interface{}) error
	// WriteResponse must be safe for concurrent use by multiple goroutines.
	WriteResponse(context.Context, *gorpc.Response, interface{}) error

	Close() error
}


type ClientCodec interface {
	// WriteRequest must be safe for concurrent use by multiple goroutines.
	WriteRequest(context.Context, *gorpc.Request, interface{}) error
	ReadResponseHeader(context.Context, *gorpc.Response) error
	ReadResponseBody(context.Context, interface{}) error

	Close() error
}