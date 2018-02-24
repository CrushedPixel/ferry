package ferry

import "context"

type Connection struct {
	remoteAddr string

	ferry *Ferry
	ctx   context.Context
}

func (c *Connection) Handle(r *IncomingRequest) Response {
	return c.ferry.handle(c, r)
}

func (c *Connection) RemoteAddr() string {
	return c.remoteAddr
}

// Context returns the Connection's Context.
func (c *Connection) Context() context.Context {
	return c.ctx
}

// SetContext sets the context to ctx.
// The provided ctx must be non-nil.
func (c *Connection) SetContext(ctx context.Context) {
	if ctx == nil {
		panic("nil context")
	}
	c.ctx = ctx
}
