package ferry

type HandlerFunc func(req *Request) Response

type ConnectionHandler interface {
	Handle(req *ConnectionRequest, conn *Connection) Response
}

type ConnectionHandlerFunc func(req *ConnectionRequest, conn *Connection) Response

func (h ConnectionHandlerFunc) Handle(req *ConnectionRequest, conn *Connection) Response {
	return h(req, conn)
}

type connectionHandlerChain []ConnectionHandler

func (c connectionHandlerChain) Handle(req *ConnectionRequest, conn *Connection) Response {
	for _, h := range c {
		if res := h.Handle(req, conn); res != nil {
			return res
		}
	}
	return nil
}
