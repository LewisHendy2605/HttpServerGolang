package response

import "github.com/LewisHendy2605/HttpServerGolang/internal/field_line"

type Response struct {
	StatusLine string
	Headers    *field_line.Headers
	Body       []byte
}

func NewResponse() *Response {
	return &Response{
		Headers: field_line.NewHeaders(),
	}
}
