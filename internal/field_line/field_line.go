package field_line

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/LewisHendy2605/HttpServerGolang/internal/syntax_notation"
)

/*
5. Field Syntax

Each field line consists of a case-insensitive field name followed by a colon (":"),
optional leading whitespace, the field line value, and optional trailing whitespace.

	field-line   = field-name ":" OWS field-value OWS

Rules for parsing within field values are defined in Section 5.5 of [HTTP].

This section covers the generic syntax for header field inclusion within, and extraction from, HTTP/1.1 messages.

5.1. Field Line Parsing

Messages are parsed using a generic algorithm, independent of the individual field names.

The contents within a given field line value are not parsed until a later stage of message interpretation (usually after the message's entire field section has been processed).

No whitespace is allowed between the field name and colon.

In the past, differences in the handling of such whitespace have led to security vulnerabilities in request routing and response handling.

A server MUST reject, with a response status code of 400 (Bad Request), any received request message that contains whitespace between a header field name and colon.

A proxy MUST remove any such whitespace from a response message before forwarding the message downstream.

A field line value might be preceded and/or followed by optional whitespace (OWS); a single SP preceding the field line value is preferred for consistent readability by humans.

The field line value does not include that leading or trailing whitespace: OWS occurring before the first non-whitespace octet of the field line value,
or after the last non-whitespace octet of the field line value, is excluded by parsers when extracting the field line value from a field line.
*/
type Headers struct {
	headers map[string]string
}

// Setter for header values
func (h *Headers) Set(name string, value string) {
	name_lower := strings.ToLower(name)

	// Append any repeating values to a comma separated string
	if _, ok := h.Get(name_lower); ok {
		h.headers[name_lower] = strings.Join([]string{h.headers[name_lower], value}, ", ")
		return
	}

	h.headers[name_lower] = value
}

// Getter for header values
func (h *Headers) Get(name string) (string, bool) {
	val, ok := h.headers[strings.ToLower(name)]
	return val, ok

}

// Formats headers to comma separated list
func (h *Headers) String() string {
	headers := make([]string, len(h.headers))

	for k, v := range h.headers {
		headers = append(headers, fmt.Sprintf("%s: %s", k, v))
	}

	return strings.Join(headers, ", ")
}

// Parses Field Line
func (h *Headers) Parse(data []byte) (int, error) {
	h.headers = make(map[string]string)
	bytes_read := 0

	for {
		index := bytes.Index(data[bytes_read:], []byte(syntax_notation.CRLF))
		if index == -1 {
			break
		}

		field_line := data[bytes_read : bytes_read+index]
		bytes_read += index + len(syntax_notation.CRLF)

		if len(field_line) == 0 {
			break
		}

		parts := bytes.SplitN(field_line, []byte(syntax_notation.COLON), 2)
		if len(parts) != 2 {
			return 0, fmt.Errorf("invalid field line: missing colon")
		}

		name := parts[0]
		if bytes.Contains(name, []byte(syntax_notation.SP)) {
			return 0, fmt.Errorf("Error parsing field line, found unexpected white space in name")
		}
		if !isToken(name) {
			return 0, fmt.Errorf("Error parsing field line, name contained invalid token")
		}

		value := parts[1]
		if bytes.Index(value, []byte(syntax_notation.SP)) != 0 {
			return 0, fmt.Errorf("invalid field line: missing required white space at start of value")
		}
		value = bytes.TrimSpace(value)

		h.Set(string(name), string(value))
	}

	return bytes_read, nil
}

func NewHeaders() *Headers {
	return &Headers{
		headers: make(map[string]string),
	}
}

func isToken(data []byte) bool {
	for _, b := range data {
		switch {
		case b >= 'a' && b <= 'z':
			continue
		case b >= 'A' && b <= 'Z':
			continue
		case b >= '0' && b <= '9':
			continue
		case b == '!' || b == '#' || b == '$' || b == '%' ||
			b == '&' || b == '\'' || b == '*' || b == '+' ||
			b == '-' || b == '.' || b == '^' || b == '_' ||
			b == '`' || b == '|' || b == '~':
			continue
		default:
			return false

		}
	}

	return true
}
