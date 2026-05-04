package request

import (
	"bytes"
	"fmt"
	"strings"
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
type Headers map[string]string

func (h Headers) Set(name string, value string) {
	h[strings.ToLower(name)] = value
}

func (h Headers) Get(name string) (string, bool) {
	val, ok := h[strings.ToLower(name)]
	return val, ok

}

func (h Headers) String() string {
	headers := make([]string, len(h))

	for k, v := range h {
		headers = append(headers, fmt.Sprintf("%s: %s", k, v))
	}

	return strings.Join(headers, ", ")
}

// Parses Field Line
func (h *Headers) Parse(data []byte) (int, error) {
	bytes_read := 0

	for {
		index := bytes.Index(data[bytes_read:], []byte(CRLF))
		if index == -1 {
			break
		}

		field_line := data[bytes_read : bytes_read+index]
		bytes_read += index + len(CRLF)

		if len(field_line) == 0 {
			break
		}

		fmt.Printf("Parsing header: %s\n", string(field_line))

		parts := bytes.SplitN(field_line, []byte(COLON), 2)
		if len(parts) != 2 {
			return 0, fmt.Errorf("invalid field line: missing colon")
		}

		name := string(parts[0])
		if strings.Contains(name, SP) || strings.Contains(name, HTAB) {
			return 0, fmt.Errorf("Error parsing field line, field name. Found unexpected white space")
		}

		value := strings.TrimSpace(strings.Trim(string(parts[1]), HTAB))

		fmt.Printf("Setting new header. Name: %s, Value: %s\n", name, value)
		h.Set(name, value)
	}

	return bytes_read, nil
}
