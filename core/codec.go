package core

import context "golang.org/x/net/context"

type ServerCodec interface {
	ReadRequestHeader(context.Context, *Request) error
	ReadRequestBody(context.Context, interface{}) error
	// WriteResponse must be safe for concurrent use by multiple goroutines.
	WriteResponse(context.Context, *Response, interface{}) error

	Close() error
}
