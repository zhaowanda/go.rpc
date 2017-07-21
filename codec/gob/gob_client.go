package gob

import (
	"io"
	"encoding/gob"
	"bufio"
	"golang.org/x/net/context"
	"github.com/zhaowanda/go.rpc/codec"
	"github.com/zhaowanda/go.rpc/gorpc"
)

type gobClientCodec struct {
	rwc    io.ReadWriteCloser
	dec    *gob.Decoder
	enc    *gob.Encoder
	encBuf *bufio.Writer
}


func NewClientCodec(conn io.ReadWriteCloser) codec.ClientCodec {
	encBuf := bufio.NewWriter(conn)
	client := &gobClientCodec{
		rwc: 	conn,
		dec: 	gob.NewDecoder(conn),
		enc: 	gob.NewEncoder(encBuf),
		encBuf: encBuf,
	}
	return client
}

func (c *gobClientCodec) WriteRequest(ctx context.Context, r *gorpc.Request, body interface{}) (err error) {
	if err = c.enc.Encode(r); err != nil {
		return
	}
	if err = c.enc.Encode(body); err != nil {
		return
	}
	return c.encBuf.Flush()
}

func (c *gobClientCodec) ReadResponseHeader(ctx context.Context, r *gorpc.Response) error {
	return c.dec.Decode(r)
}

func (c *gobClientCodec) ReadResponseBody(ctx context.Context, body interface{}) error {
	return c.dec.Decode(body)
}

func (c *gobClientCodec) Close() error {
	return c.rwc.Close()
}
