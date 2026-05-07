package request

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/LewisHendy2605/HttpServerGolang/internal/field_line"
	"github.com/LewisHendy2605/HttpServerGolang/internal/request_line"
)

type ParserState string

const (
	StateInit        ParserState = "init"
	StateFieldLine   ParserState = "field_line"
	StateMessageBody ParserState = "message_body"
	StateDone        ParserState = "done"
)

/*
2.1. Message Format

An HTTP/1.1 message consists of a start-line followed by a CRLF and a sequence of octets in a
format similar to the Internet Message Format [RFC5322]: zero or more header field lines
(collectively referred to as the "headers" or the "header section"),
an empty line indicating the end of the header section, and an optional message body.

	HTTP-message   = start-line CRLF
	                 *( field-line CRLF )
	                 CRLF
	                 [ message-body ]

A message can be either a request from client to server or a response from server to client.
Syntactically, the two types of messages differ only in the start-line,
which is either a request-line (for requests) or a status-line (for responses),
and in the algorithm for determining the length of the message body (Section 6).

	start-line     = request-line / status-line

In theory, a client could receive requests and a server could receive responses,
distinguishing them by their different start-line formats.

In practice, servers are implemented to only expect a request (a response is interpreted as an unknown or invalid request method),
and clients are implemented to only expect a response.

HTTP makes use of some protocol elements similar to the Multipurpose Internet Mail Extensions (MIME) [RFC2045].
*/

type Request struct {
	RequestLine *request_line.RequestLine
	Headers     *field_line.Headers
	Body        []byte
	state       ParserState
}

// Formats request to a string, primarily for debugging
func (r *Request) String() string {
	return fmt.Sprintf("{ RequestLine: %s, Headers: %s }", r.RequestLine.String(), r.Headers.String())
}

// Returns request body parsed as text
func (r *Request) Text() string {
	return string(r.Body)
}

// Returns request body Unmarshaled to input struct
func (r *Request) Json(v any) error {
	return json.Unmarshal(r.Body, v)
}

// Parses a http request
func (r *Request) parse(data []byte) error {
	bytes_read := 0

outerLoop:
	for {
		current_data := data[bytes_read:]
		switch r.state {
		case StateInit:
			r.RequestLine = &request_line.RequestLine{}
			read, err := r.RequestLine.Parse(current_data)
			if err != nil {
				return err
			}

			bytes_read += read
			r.state = StateFieldLine
		case StateFieldLine:
			r.Headers = &field_line.Headers{}
			read, err := r.Headers.Parse(current_data)
			if err != nil {
				return err
			}

			bytes_read += read
			r.state = StateMessageBody
		case StateMessageBody:
			r.Body = current_data
			r.state = StateDone
		case StateDone:
			break outerLoop
		}
	}

	return nil
}

// Creates and parses a http request from a reader
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

/*
2.2. Message Parsing

The normal procedure for parsing an HTTP message is to read the start-line into a structure,
read each header field line into a hash table by field name until the empty line,
and then use the parsed data to determine if a message body is expected.

If a message body has been indicated,
then it is read as a stream until an amount of octets equal to the message body length is read or the connection is closed.

A recipient MUST parse an HTTP message as a sequence of octets in an encoding that is a superset of US-ASCII [USASCII].

Parsing an HTTP message as a stream of Unicode characters,
without regard for the specific encoding,
creates security vulnerabilities due to the varying ways that string processing libraries handle invalid
multi byte character sequences that contain the octet LF (%x0A).

String-based parsers can only be safely used within protocol elements after the element has been extracted from the message,
such as within a header field line value after message parsing has delineated the individual field lines.

Although the line terminator for the start-line and fields is the sequence CRLF,
a recipient MAY recognize a single LF as a line terminator and ignore any preceding CR.

A sender MUST NOT generate a bare CR (a CR character not immediately followed by LF) within any protocol
elements other than the content.

A recipient of such a bare CR MUST consider that element to be invalid or replace each bare
CR with SP before processing the element or forwarding the message.

Older HTTP/1.0 user agent implementations might send an extra CRLF after a POST request as a
workaround for some early server applications that failed to read message body content that was not
terminated by a line-ending.

An HTTP/1.1 user agent MUST NOT preface or follow a request with an extra CRLF.

If terminating the request message body with a line-ending is desired,
then the user agent MUST count the terminating CRLF octets as part of the message body length.

In the interest of robustness,
a server that is expecting to receive and parse a request-line SHOULD
ignore at least one empty line (CRLF) received prior to the request-line.

A sender MUST NOT send whitespace between the start-line and the first header field.

A recipient that receives whitespace between the start-line and the first header field MUST
either reject the message as invalid or consume each whitespace-preceded line without further processing of it
(i.e., ignore the entire line, along with any subsequent lines preceded by whitespace,
until a properly formed header field is received or the header section is terminated).

Rejection or removal of invalid whitespace-preceded lines is necessary to prevent their misinterpretation
by downstream recipients that might be vulnerable to request smuggling
(Section 11.2) or response splitting (Section 11.1) attacks.

When a server listening only for HTTP request messages,
or processing what appears from the start-line to be an HTTP request message,
receives a sequence of octets that does not match the HTTP-message grammar aside from the
robustness exceptions listed above, the server SHOULD respond with a 400 (Bad Request) response and close the connection.
*/
