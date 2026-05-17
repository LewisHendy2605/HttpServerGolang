package response

import (
	"fmt"
	"io"

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

func (r *Response) String() string {
	return fmt.Sprintf("%s\r\n%s\r\n%s", r.StatusLine.String(), r.Headers.String(), string(r.Body))
}

func (r *Response) Byte() []byte {
	return []byte(r.String())
}

func (r *Response) WriteStatusLine(w io.Writer, code StatusCode) error {
	statusLine := ""
	switch code {
	case StatusOK:
		statusLine = "HTTP/1.1 200 OK"
	case StatusBadRequest:
		statusLine = "HTTP/1.1 400 Bad Request"
	case StatusInternalServerError:
		statusLine = "HTTP/1.1 500 Internal Server Error"
	default:
		return fmt.Errorf("unrecognized status code")
	}

	_, err := w.Write([]byte(statusLine + "\r\n"))
	return err
}

func (r *Response) WriteHeaders(w io.Writer) error {
	_, err := w.Write([]byte(r.Headers.String() + "\r\n\r\n"))
	return err
}
