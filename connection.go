package ferry

import "context"

type Connection struct {
	RemoteAddr string

	ferry *Ferry
	ctx   context.Context
}

func (c *Connection) Handle(r *IncomingRequest) Response {
	return c.ferry.handle(c, r)
}

func (c *Connection) Context() context.Context {
	return c.ctx
}

// WithContext returns a shallow copy of c with its context changed
// to ctx. The provided ctx must be non-nil.
func (c *Connection) WithContext(ctx context.Context) *Connection {
	if ctx == nil {
		panic("nil context")
	}
	c2 := new(Connection)
	*c2 = *c
	c2.ctx = ctx
	return c2
}
