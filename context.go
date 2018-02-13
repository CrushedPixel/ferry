package ferry

import (
	"fmt"
	"io"
)

type Context struct {
	Connection *Connection

	PathParams  map[string]string
	QueryParams map[string][]string
	Payload     io.Reader

	data map[string]interface{}
}

func (c *Context) MustGet(key string) interface{} {
	val, ok := c.Get(key)
	if !ok {
		panic(fmt.Sprintf(`key "%s" does not exist`, key))
	}
	return val
}

func (c *Context) Get(key string) (interface{}, bool) {
	val, ok := c.data[key]
	return val, ok
}

func (c *Context) Set(key string, value interface{}) {
	c.data[key] = value
}
