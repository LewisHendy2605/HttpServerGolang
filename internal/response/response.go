package response

import (
	"fmt"

	"github.com/LewisHendy2605/HttpServerGolang/internal/field_line"
	"github.com/LewisHendy2605/HttpServerGolang/internal/status_line"
)

type Response struct {
	StatusLine status_line.StatusLine
	Headers    *field_line.Headers
	Body       []byte
}

func NewResponse() *Response {
	return &Response{
		StatusLine: status_line.NewStatusLine(),
		Headers:    field_line.NewHeaders(),
	}
}

func (r *Response) Ok() {
	r.StatusLine = status_line.StatusOK
}

func (r *Response) String() string {
	return fmt.Sprintf("%s\r\n%s\r\n%s", r.StatusLine.String(), r.Headers.String(), string(r.Body))
}

func (r *Response) Byte() []byte {
	return []byte(r.String())
}
