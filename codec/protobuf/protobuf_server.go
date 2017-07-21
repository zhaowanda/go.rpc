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

const defaultBufferSize = 4 * 1024


type serverCodec struct {
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
func NewServerCodec(rwc io.ReadWriteCloser) codec.ServerCodec {
	w := bufio.NewWriterSize(rwc, defaultBufferSize)
	r := bufio.NewReaderSize(rwc, defaultBufferSize)
	return &serverCodec{
		enc: NewEncoder(w),
		w:   w,
		dec: NewDecoder(r),
		c:   rwc,
	}
}


func (c *serverCodec) WriteResponse(context context.Context, resp *gorpc.Response, body interface{}) error {
	c.mu.Lock()
	c.resp.Method = resp.ServiceMethod
	c.resp.Seq = resp.Seq
	c.resp.Error = resp.Error

	err := encode(c.enc, &c.resp)
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

func (c *serverCodec) ReadRequestHeader(context context.Context, req *gorpc.Request) error {
	c.req.Reset()
	if err := c.dec.Decode(&c.req); err != nil {
		return err
	}

	req.ServiceMethod = c.req.Method
	req.Seq = c.req.Seq
	return nil
}

func (c *serverCodec) ReadRequestBody(context context.Context, body interface{}) error {
	if pb, ok := body.(proto.Message); ok {
		return c.dec.Decode(pb)
	}
	return fmt.Errorf("%T does not implement proto.Message", body)
}

func (c *serverCodec) Close() error { return c.c.Close() }


