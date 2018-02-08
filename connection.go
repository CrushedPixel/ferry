package ferry

type Connection struct {
	ferry *Ferry
	data  map[string]interface{}

	RemoteAddr string
}

func (c *Connection) Get(key string) interface{} {
	return c.data[key]
}

func (c *Connection) Set(key string, value interface{}) {
	c.data[key] = value
}

func (c *Connection) Handle(r *Request) *Response {
	return c.ferry.handle(c, r)
}
