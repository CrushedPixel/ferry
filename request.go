package ferry

import (
	"context"
	"io"
)

type ConnectionRequest struct {
	RemoteAddr string
	Header     map[string][]string
}

type IncomingRequest struct {
	Method     string
	RequestURI string
	Payload    io.Reader
}

type Request struct {
	connection *Connection

	pathParams  map[string]string
	queryParams map[string][]string
	payload     io.Reader

	ctx context.Context
}

// Connection returns the Connection making the Request.
func (r *Request) Connection() *Connection {
	return r.connection
}

func (r *Request) PathParams() map[string]string {
	return r.pathParams
}

func (r *Request) QueryParams() map[string][]string {
	return r.queryParams
}

func (r *Request) Payload() io.Reader {
	return r.payload
}

// Context returns the Request's Context.
func (r *Request) Context() context.Context {
	return r.ctx
}

// SetContext sets the context to ctx.
// The provided ctx must be non-nil.
func (r *Request) SetContext(ctx context.Context) {
	if ctx == nil {
		panic("nil context")
	}
	r.ctx = ctx
}
