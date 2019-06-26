package fisttp

type Response struct {
	request *Request
	Payload []byte
}

type IndexResponse struct {
	*Response
}

func NewResponse(payload []byte) *Response {
	return &Response{Payload: payload}
}
