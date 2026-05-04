package parser

import (
	"strings"
)

/*
Resources:
	RFC_9112  : https://datatracker.ietf.org/doc/html/rfc9112
	RFC: 9110 : https://www.rfc-editor.org/rfc/rfc9110
	RFC_5789  : https://www.rfc-editor.org/rfc/rfc5789.html
*/

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

type Response struct {
	StatusLine  *StatusLine
	FieldLines  []string
	MessageBody string
}

/*
// Parses a http response
func ParseHttpResponse(http_message []byte) (*HttpResponse, error) {
	parts := bytes.SplitN(http_message, CRLF, 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid http message: missing crlf")
	}

	// 'In practice, servers are implemented to only expect a request'
	status_line, err := ParseStatusLine(parts[0])
	if err != nil {
		return nil, err
	}

	var message_body []byte
	var field_lines []*FieldLine

	lines := bytes.Split(parts[1], CRLF)

	for i, line := range lines {
		if bytes.Equal(line, CRLF) {
			message_body = bytes.Join(lines[i:], CRLF)
		} else {
			fl, err := ParseFieldLine(line)
			if err != nil {
				return nil, err
			}
			field_lines = append(field_lines, fl)
		}
	}

	return &HttpResponse{
		StatusLine:  status_line,
		FieldLines:  field_lines,
		MessageBody: message_body,
	}, nil
}
*/

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

/*
4. Status Line

The first line of a response message is the status-line, consisting of the protocol version, a space (SP),
the status code, and another space and ending with an OPTIONAL textual phrase describing the status code.

	status-line = HTTP-version SP status-code SP [ reason-phrase ]

Although the status-line grammar rule requires that each of the component elements be separated by a single SP octet,
recipients MAY instead parse on whitespace-delimited word boundaries and,
aside from the line terminator,
treat any form of whitespace as the SP separator while ignoring preceding or trailing whitespace;
such whitespace includes one or more of the following octets: SP, HTAB, VT (%x0B), FF (%x0C), or bare CR.

However, lenient parsing can result in response splitting security vulnerabilities if there are multiple recipients of the message and each has its own unique interpretation of robustness (see Section 11.1).

The status-code element is a 3-digit integer code describing the result of the server's attempt to understand and satisfy the client's corresponding request.

A recipient parses and interprets the remainder of the response message in light of the semantics defined for that status code,
if the status code is recognized by that recipient,
or in accordance with the class of that status code when the specific code is unrecognized.

	status-code    = 3DIGIT

HTTP's core status codes are defined in Section 15 of [HTTP], along with the classes of status codes, considerations for the definition of new status codes, and the IANA registry for collecting such definitions.

The reason-phrase element exists for the sole purpose of providing a textual description associated with the numeric status code,
mostly out of deference to earlier Internet application protocols that were more frequently used with interactive text clients.

	reason-phrase  = 1*( HTAB / SP / VCHAR / obs-text )

A client SHOULD ignore the reason-phrase content because it is not a reliable channel for information (it might be translated for a given locale,
overwritten by intermediaries,
or discarded when the message is forwarded via other versions of HTTP).

A server MUST send the space that separates the status-code from the reason-phrase even when the reason-phrase is absent (i.e., the status-line would end with the space).
*/
type StatusLine struct {
	HttpVersion  string
	StatusCode   string
	ReasonPhrase string
}

func ParseStatusLine(status_line string) (*StatusLine, error) {
	parts := strings.Split(status_line, " ")

	return &StatusLine{
		HttpVersion:  parts[0],
		StatusCode:   parts[1],
		ReasonPhrase: parts[2],
	}, nil
}

/*
6. Message Body

The message body (if any) of an HTTP/1.1 message is used to carry content (Section 6.4 of [HTTP]) for the request or response.

The message body is identical to the content unless a transfer coding has been applied, as described in Section 6.1.

	message-body = *OCTET

The rules for determining when a message body is present in an HTTP/1.1 message differ for requests and responses.

The presence of a message body in a request is signaled by a Content-Length or Transfer-Encoding header field.

Request message framing is independent of method semantics.

The presence of a message body in a response, as detailed in Section 6.3,
depends on both the request method to which it is responding and the response status code.

This corresponds to when response content is allowed by HTTP semantics (Section 6.4.1 of [HTTP]).
*/
type MessageBody struct {
}

func ParseMessageBody(message_body []byte) (*MessageBody, error) {
	return &MessageBody{}, nil
}

/*
6.1. Transfer-Encoding

The Transfer-Encoding header field lists the transfer coding names corresponding to the sequence of transfer codings that have been (or will be) applied to the content in order to form the message body.

Transfer codings are defined in Section 7.

  Transfer-Encoding = #transfer-coding
                       ; defined in [HTTP], Section 10.1.4

Transfer-Encoding is analogous to the Content-Transfer-Encoding field of MIME, which was designed to enable safe transport of binary data over a 7-bit transport service ([RFC2045], Section 6).

However, safe transport has a different focus for an 8bit-clean transfer protocol.

In HTTP's case, Transfer-Encoding is primarily intended to accurately delimit dynamically generated content.

It also serves to distinguish encodings that are only applied in transit from the encodings that are a characteristic of the selected representation.

A recipient MUST be able to parse the chunked transfer coding (Section 7.1) because it plays a crucial role in framing messages when the content size is not known in advance.

A sender MUST NOT apply the chunked transfer coding more than once to a message body (i.e., chunking an already chunked message is not allowed).

If any transfer coding other than chunked is applied to a request's content,
the sender MUST apply chunked as the final transfer coding to ensure that the message is properly framed.

If any transfer coding other than chunked is applied to a response's content, the sender MUST either apply chunked as the final transfer coding or terminate the message by closing the connection.
*/
