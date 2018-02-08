package ferry

type Response interface {
	Response() (int, string)
}

// response is a basic implementation of Response
type response struct {
	status int
	body   string
}

func (r *response) Response() (int, string) {
	return r.status, r.body
}

// NewResponse returns a new Response for a status code and response body
func NewResponse(status int, body string) Response {
	return &response{
		status: status,
		body:   body,
	}
}
