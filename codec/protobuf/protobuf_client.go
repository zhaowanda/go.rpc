package protobuf

import (
	"io"
	"bufio"
	"sync"

	"github.com/golang/protobuf/proto"
	"fmt"
	context "golang.org/x/net/context"
	"github.com/zhaowanda/go.rpc/gorpc"
	"github.com/zhaowanda/go.rpc/codec"
)


type clientCodec struct {
	mu   sync.Mutex // exclusive writer lock
	resp ResponseHeader
	enc  *Encoder
	w    *bufio.Writer

	req RequestHeader
	dec *Decoder
	c   io.Closer
}

// NewServerCodec returns a new rpc.ServerCodec.
//
// A ServerCodec implements reading of RPC requests and writing of RPC
// responses for the server side of an RPC session. The server calls
// ReadRequestHeader and ReadRequestBody in pairs to read requests from the
// connection, and it calls WriteResponse to write a response back. The
// server calls Close when finished with the connection. ReadRequestBody
// may be called with a nil argument to force the body of the request to be
// read and discarded.

func NewClientCodec(rwc io.ReadWriteCloser) codec.ClientCodec {
	w := bufio.NewWriterSize(rwc, defaultBufferSize)
	r := bufio.NewReaderSize(rwc, defaultBufferSize)
	return &clientCodec{
		enc: NewEncoder(w),
		w:   w,
		dec: NewDecoder(r),
		c:   rwc,
	}
}


func (c *clientCodec) WriteRequest(context context.Context, req *gorpc.Request, body interface{}) error {
	c.mu.Lock()
	c.req.Method = req.ServiceMethod
	c.req.Seq = req.Seq

	err := encode(c.enc, &c.req)
	if err != nil {
		c.mu.Unlock()
		return err
	}
	if err = encode(c.enc, body); err != nil {
		c.mu.Unlock()
		return err
	}
	err = c.w.Flush()
	c.mu.Unlock()
	return err
}

func (c *clientCodec) ReadResponseHeader(context context.Context, resp *gorpc.Response) error {
	c.req.Reset()
	if err := c.dec.Decode(&c.req); err != nil {
		return err
	}

	resp.ServiceMethod = c.req.Method
	resp.Seq = c.req.Seq
	resp.Error = c.resp.Error
	return nil
}

func (c *clientCodec) ReadResponseBody(context context.Context, body interface{}) error {
	if pb, ok := body.(proto.Message); ok {
		return c.dec.Decode(pb)
	}
	return fmt.Errorf("%T does not implement proto.Message", body)
}

func (c *clientCodec) Close() error { return c.c.Close() }


