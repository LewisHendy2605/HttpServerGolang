package response

import (
	"fmt"

	"github.com/LewisHendy2605/HttpServerGolang/internal/headers"
)

type Response struct {
	StatusLine StatusLine
	Headers    *headers.Headers
	Body       []byte
}

func NewResponse() *Response {
	return &Response{
		StatusLine: NewStatusLine(),
		Headers:    headers.NewHeaders(),
	}
}

func (r *Response) Ok() {
	r.StatusLine = StatusOK
}

func (r *Response) String() string {
	return fmt.Sprintf("%s\r\n%s\r\n%s", r.StatusLine.String(), r.Headers.String(), string(r.Body))
}

func (r *Response) Byte() []byte {
	return []byte(r.String())
}
