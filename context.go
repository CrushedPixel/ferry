package ferry

import "io"

type Context struct {
	Connection *Connection

	PathParams  map[string]string
	QueryParams map[string][]string
	Payload     io.Reader

	data map[string]interface{}
}

func (c *Context) Get(key string) interface{} {
	return c.data[key]
}

func (c *Context) Set(key string, value interface{}) {
	c.data[key] = value
}
