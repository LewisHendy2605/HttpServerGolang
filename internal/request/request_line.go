package request

import (
	"bytes"
	"fmt"
	"strings"
)

/*
3. Request Line

A request-line begins with a method token,
followed by a single space (SP), the request-target,
and another single space (SP), and ends with the protocol version.

request-line   = method SP request-target SP HTTP-version
*/
type RequestLine struct {
	Method        string
	RequestTarget string
	HttpVersion   *HttpVersion
}

// Formats Request Line to string for debugging
func (rl *RequestLine) String() string {
	return fmt.Sprintf("%s %s HTTP/%d.%d", rl.Method, rl.RequestTarget, rl.HttpVersion.Major, rl.HttpVersion.Minor)
}

// Parses http request line
func (rl *RequestLine) Parse(data []byte) (int, error) {
	index := bytes.Index(data, []byte(CRLF))
	if index == -1 {
		return 0, nil
	}

	start_line := data[:index]

	parts := bytes.Split(start_line, []byte(SP))
	if len(parts) != 3 {
		return 0, fmt.Errorf("invalid request line, missing space")
	}

	rl.Method = strings.ToUpper(string(parts[0]))
	if !IsMethod(rl.Method) {
		return 0, fmt.Errorf("invalid request line, invalid http method")
	}

	rl.RequestTarget = strings.Trim(string(parts[1]), SP)
	if len(rl.RequestTarget) == 0 {
		return 0, fmt.Errorf("invalid request line, missing request target")
	}

	rl.HttpVersion = &HttpVersion{}
	err := rl.HttpVersion.Parse(parts[2])
	if err != nil {
		return 0, err
	}

	return index + len(CRLF), nil
}
