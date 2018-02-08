package ferry

import (
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
)

type Ferry struct {
	OnConnection HandleConnectionFunc
	router       *mux.Router
	routes       map[string][]HandlerFunc
}

type HandlerFunc func(c *Context) *Response
type HandleConnectionFunc func(r *ConnectionRequest, c *Connection) *Response

func defaultOnConnection(r *ConnectionRequest, c *Connection) *Response {
	return nil
}

func New() *Ferry {
	return &Ferry{
		OnConnection: defaultOnConnection,
		router:       mux.NewRouter(),
		routes:       make(map[string][]HandlerFunc),
	}
}

func (f *Ferry) NewConnection(r *ConnectionRequest) (*Connection, *Response) {
	c := &Connection{
		ferry: f,
		data:  make(map[string]interface{}),

		RemoteAddr: r.RemoteAddr,
	}

	if res := f.OnConnection(r, c); res == nil {
		return c, nil
	} else {
		return nil, res
	}
}

func (f *Ferry) handle(c *Connection, r *Request) *Response {
	u, err := url.Parse(r.RequestURI)
	if err != nil {
		panic(err)
	}

	// emulate http request to use gorilla-mux
	// for URL parsing
	req := &http.Request{
		Method: r.Method,
		URL:    u,
	}
	var match mux.RouteMatch
	if f.router.Match(req, &match) {
		handlers := f.routes[match.Route.GetName()]

		// create context
		context := &Context{
			Connection:  c,
			PathParams:  match.Vars,
			QueryParams: u.Query(),
			Payload:     r.Payload,

			data: make(map[string]interface{}),
		}

		// execute handler chain
		for _, h := range handlers {
			res := h(context)
			if res != nil {
				return res
			}
		}

		panic("last handler in chain did not return a value")
	}

	return &Response{
		Status:  http.StatusNotFound,
		Payload: "404 page not found",
	}
}

func (f *Ferry) Handle(method string, path string, handlers ...HandlerFunc) {
	key := string(len(f.routes))
	f.router.Handle(path, nil).Methods(method).Name(key)
	f.routes[key] = handlers
}

func (f *Ferry) GET(path string, handlers ...HandlerFunc) {
	f.Handle(http.MethodGet, path, handlers...)
}

func (f *Ferry) POST(path string, handlers ...HandlerFunc) {
	f.Handle(http.MethodPost, path, handlers...)
}

func (f *Ferry) PATCH(path string, handlers ...HandlerFunc) {
	f.Handle(http.MethodPatch, path, handlers...)
}

func (f *Ferry) PUT(path string, handlers ...HandlerFunc) {
	f.Handle(http.MethodPut, path, handlers...)
}

func (f *Ferry) DELETE(path string, handlers ...HandlerFunc) {
	f.Handle(http.MethodDelete, path, handlers...)
}

func (f *Ferry) OPTIONS(path string, handlers ...HandlerFunc) {
	f.Handle(http.MethodOptions, path, handlers...)
}

func (f *Ferry) TRACE(path string, handlers ...HandlerFunc) {
	f.Handle(http.MethodTrace, path, handlers...)
}

func (f *Ferry) CONNECT(path string, handlers ...HandlerFunc) {
	f.Handle(http.MethodConnect, path, handlers...)
}

func (f *Ferry) HEAD(path string, handlers ...HandlerFunc) {
	f.Handle(http.MethodHead, path, handlers...)
}
