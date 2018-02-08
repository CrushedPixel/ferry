package ferry

import (
	"io"
)

type Request struct {
	Method     string
	RequestURI string
	Payload    io.Reader
}

type ConnectionRequest struct {
	RemoteAddr string
	Header     map[string][]string
}
