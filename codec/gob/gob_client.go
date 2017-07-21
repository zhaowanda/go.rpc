package gob

import (
	"io"
	"encoding/gob"
	"bufio"
	"github.com/zhaowanda/go.rpc/core"
	"golang.org/x/net/context"
)

type gobClientCodec struct {
	rwc    io.ReadWriteCloser
	dec    *gob.Decoder
	enc    *gob.Encoder
	encBuf *bufio.Writer
}


func NewClientCodec(conn io.ReadWriteCloser) core.ClientCodec {
	encBuf := bufio.NewWriter(conn)
	client := &gobClientCodec{
		rwc: 	conn,
		dec: 	gob.NewDecoder(conn),
		enc: 	gob.NewEncoder(encBuf),
		encBuf: encBuf,
	}
	return client
}

func (c *gobClientCodec) WriteRequest(ctx context.Context, r *core.Request, body interface{}) (err error) {
	if err = c.enc.Encode(r); err != nil {
		return
	}
	if err = c.enc.Encode(body); err != nil {
		return
	}
	return c.encBuf.Flush()
}

func (c *gobClientCodec) ReadResponseHeader(ctx context.Context, r *core.Response) error {
	return c.dec.Decode(r)
}

func (c *gobClientCodec) ReadResponseBody(ctx context.Context, body interface{}) error {
	return c.dec.Decode(body)
}

func (c *gobClientCodec) Close() error {
	return c.rwc.Close()
}
