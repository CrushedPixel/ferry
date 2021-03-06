package ferry

import (
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

type Ferry struct {
	connectionHandlers connectionHandlerChain
	router             *mux.Router
	routes             map[string][]HandlerFunc
}

func New() *Ferry {
	return &Ferry{
		router: mux.NewRouter(),
		routes: make(map[string][]HandlerFunc),
	}
}

func (f *Ferry) NewConnection(r *ConnectionRequest) (*Connection, Response) {
	c := &Connection{
		remoteAddr: r.RemoteAddr,

		ferry: f,
		ctx:   context.Background(),
	}

	if res := f.connectionHandlers.Handle(r, c); res == nil {
		return c, nil
	} else {
		return nil, res
	}
}

func (f *Ferry) handle(c *Connection, r *IncomingRequest) Response {
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

		// create Request object
		req := &Request{
			connection:  c,
			pathParams:  match.Vars,
			queryParams: u.Query(),
			payload:     r.Payload,
			ctx:         context.Background(),
		}

		// execute handler chain
		for _, h := range handlers {
			res := h(req)
			if res != nil {
				return res
			}
		}

		panic("last handler in chain did not return a value")
	}

	return NewResponse(http.StatusNotFound, "404 page not found")
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

// OnConnectionHandler adds a ConnectionHandler to be executed
// whenever a connection request is received.
func (f *Ferry) OnConnectionHandler(c ConnectionHandler) {
	f.connectionHandlers = append(f.connectionHandlers, c)
}

// OnConnectionHAndlerFunc adds a ConnectionHandlerFunc to be executed
// whenever a connection request is received.
func (f *Ferry) OnConnectionHandlerFunc(c ConnectionHandlerFunc) {
	f.OnConnectionHandler(c)
}
