package request

import (
	"fmt"
	"io"
)

type ParserState string

const (
	StateInit        ParserState = "init"
	StateFieldLine   ParserState = "field_line"
	StateMessageBody ParserState = "message_body"
	StateDone        ParserState = "done"
)

type Request struct {
	RequestLine *RequestLine
	Headers     Headers
	state       ParserState
}

func (r *Request) String() string {
	return fmt.Sprintf("{ RequestLine: %s, Headers: %s }", r.RequestLine.String(), r.Headers.String())
}

// Parses a http request
func (r *Request) parse(data []byte) error {
	bytes_read := 0

outerLoop:
	for {
		current_data := data[bytes_read:]
		switch r.state {
		case StateInit:
			r.RequestLine = &RequestLine{}
			read, err := r.RequestLine.Parse(current_data)
			if err != nil {
				return err
			}

			bytes_read += read
			r.state = StateFieldLine
		case StateFieldLine:
			r.Headers = Headers{}
			read, err := r.Headers.Parse(current_data)
			if err != nil {
				return err
			}

			bytes_read += read
			r.state = StateDone
		case StateDone:
			break outerLoop
		}
	}

	return nil
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	buffer := make([]byte, 1024)
	request := &Request{state: StateInit}

	bytes_read, err := reader.Read(buffer)
	if err != nil {
		return nil, err
	}

	buffer = buffer[:bytes_read]

	err = request.parse(buffer)
	if err != nil {
		return nil, err
	}

	return request, nil
}
