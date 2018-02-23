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
	Connection *Connection

	PathParams  map[string]string
	QueryParams map[string][]string
	Payload     io.Reader

	ctx context.Context
}

func (r *Request) Context() context.Context {
	return r.ctx
}

// WithContext returns a shallow copy of r with its context changed
// to ctx. The provided ctx must be non-nil.
func (r *Request) WithContext(ctx context.Context) *Request {
	if ctx == nil {
		panic("nil context")
	}
	r2 := new(Request)
	*r2 = *r
	r2.ctx = ctx
	return r2
}
