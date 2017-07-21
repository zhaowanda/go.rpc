package gob

import (
	"io"
	"encoding/gob"
	"bufio"
	"golang.org/x/net/context"
	"github.com/zhaowanda/go.rpc/log"
	"github.com/zhaowanda/go.rpc/codec"
	"github.com/zhaowanda/go.rpc/gorpc"
)

type gobServerCodec struct {
	rwc    io.ReadWriteCloser
	dec    *gob.Decoder
	enc    *gob.Encoder
	encBuf *bufio.Writer
	closed bool
}

func NewServerCodec(conn io.ReadWriteCloser) codec.ServerCodec {
	buf := bufio.NewWriter(conn)
	srv := &gobServerCodec{
		rwc:    conn,
		dec:    gob.NewDecoder(conn),
		enc:    gob.NewEncoder(buf),
		encBuf: buf,
	}
	return srv
}

func (c *gobServerCodec) ReadRequestHeader(ctx context.Context, r *gorpc.Request) error {
	return c.dec.Decode(r)
}

func (c *gobServerCodec) ReadRequestBody(ctx context.Context, body interface{}) error {
	return c.dec.Decode(body)
}

func (c *gobServerCodec) WriteResponse(ctx context.Context, r *gorpc.Response, body interface{}) (err error) {
	if err = c.enc.Encode(r); err != nil {
		if c.encBuf.Flush() == nil {
			// Gob couldn't encode the header. Should not happen, so if it does,
			// shut down the connection to signal that the connection is broken.
			log.Info("rpc: gob error encoding response:", err)
			c.Close()
		}
		return
	}
	if err = c.enc.Encode(body); err != nil {
		if c.encBuf.Flush() == nil {
			// Was a gob problem encoding the body but the header has been written.
			// Shut down the connection to signal that the connection is broken.
			log.Info("rpc: gob error encoding body:", err)
			c.Close()
		}
		return
	}
	return c.encBuf.Flush()
}

func (c *gobServerCodec) Close() error {
	if c.closed {
		// Only call c.rwc.Close once; otherwise the semantics are undefined.
		return nil
	}
	c.closed = true
	return c.rwc.Close()
}
