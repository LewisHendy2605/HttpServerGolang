package main

import (
	"bytes"
	"fmt"
)

/*
1.2. Syntax Notation

This specification uses the Augmented Backus-Naur Form (ABNF) notation of [RFC5234], extended with the notation for case-sensitivity in strings defined in [RFC7405].

It also uses a list extension, defined in Section 5.6.1 of [HTTP], that allows for compact definition of comma-separated lists using a "#" operator (similar to how the "*" operator indicates repetition).

Appendix A shows the collected grammar with all list operators expanded to standard ABNF notation.

As a convention, ABNF rule names prefixed with "obs-" denote obsolete grammar rules that appear for historical reasons.

The following core rules are included by reference, as defined in [RFC5234],
Appendix B.1: ALPHA (letters),
	CR (carriage return),
	CRLF (CR LF),
	CTL (controls),
	DIGIT (decimal 0-9),
	DQUOTE (double quote),
	HEXDIG (hexadecimal 0-9/A-F/a-f),
	HTAB (horizontal tab), LF (line feed),
	OCTET (any 8-bit sequence of data),
	SP (space),
	and VCHAR (any visible [USASCII] character).

The rules below are defined in [HTTP]:

  BWS           = <BWS, see [HTTP], Section 5.6.3>
  OWS           = <OWS, see [HTTP], Section 5.6.3>
  RWS           = <RWS, see [HTTP], Section 5.6.3>
  absolute-path = <absolute-path, see [HTTP], Section 4.1>
  field-name    = <field-name, see [HTTP], Section 5.1>
  field-value   = <field-value, see [HTTP], Section 5.5>
  obs-text      = <obs-text, see [HTTP], Section 5.6.4>
  quoted-string = <quoted-string, see [HTTP], Section 5.6.4>
  token         = <token, see [HTTP], Section 5.6.2>
  transfer-coding =
                  <transfer-coding, see [HTTP], Section 10.1.4>

The rules below are defined in [URI]:

  absolute-URI  = <absolute-URI, see [URI], Section 4.3>
  authority     = <authority, see [URI], Section 3.2>
  uri-host      = <host, see [URI], Section 3.2.2>
  port          = <port, see [URI], Section 3.2.3>
  query         = <query, see [URI], Section 3.4>
*/
var SP = []byte(" ")
var CR = []byte("\r")
var LF = []byte("\n")
var CRLF = []byte("\r\n")
var SLASH = []byte("/")
var DOT = []byte(".")
var COLON = []byte(":")
var SEMICOLON = []byte(";")
var HTAB = []byte("\t")

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
type HttpMessage struct {
	RequestLine *RequestLine
	FieldLines  []*FieldLine
	MessageBody []byte
}

// Parse a http message
func ParseHttpMessage(http_message []byte) (*HttpMessage, error) {
	parts := bytes.SplitN(http_message, CRLF, 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid http message: missing crlf")
	}

	// 'In practice, servers are implemented to only expect a request'
	request_line, err := ParseRequestLine(parts[0])
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

	return &HttpMessage{
		RequestLine: request_line,
		FieldLines:  field_lines,
		MessageBody: message_body,
	}, nil
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

/*
2.3. HTTP Version

HTTP uses a "<major>.<minor>" numbering scheme to indicate versions of the protocol.
This specification defines version "1.1". Section 2.5 of [HTTP] specifies the semantics of HTTP version numbers.

The version of an HTTP/1.x message is indicated by an HTTP-version field in the start-line.
HTTP-version is case-sensitive.

  HTTP-version  = HTTP-name "/" DIGIT "." DIGIT
  HTTP-name     = %s"HTTP"
*/
type HttpVersion struct {
	Name  []byte
	Major []byte
	Minor []byte
}

// Parses http version from start line
func ParseHttpVersion(http_version []byte) (*HttpVersion, error) {
	parts := bytes.Split(http_version, SLASH)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid http version, missing forward slash")
	}

	version := bytes.Split(parts[1], DOT)
	if len(version) != 2 {
		return nil, fmt.Errorf("invalid http version, missing period")
	}

	name := parts[0]
	if len(name) == 0 {
		return nil, fmt.Errorf("invalid http version, missing name")
	}

	major := version[0]
	if len(major) == 0 {
		return nil, fmt.Errorf("invalid http version, missing major")
	}

	minor := version[0]
	if len(minor) == 0 {
		return nil, fmt.Errorf("invalid http version, missing minor")
	}

	return &HttpVersion{
		Name:  name,
		Major: major,
		Minor: minor,
	}, nil
}

/*
3. Request Line

A request-line begins with a method token,
followed by a single space (SP), the request-target,
and another single space (SP), and ends with the protocol version.

request-line   = method SP request-target SP HTTP-version
*/
type RequestLine struct {
	Method        []byte
	RequestTarget []byte
	HttpVersion   *HttpVersion
}

// Parses http request line
func ParseRequestLine(request_line []byte) (*RequestLine, error) {
	parts := bytes.Split(request_line, SP)

	http_version, err := ParseHttpVersion(parts[2])
	if err != nil {
		return nil, err
	}

	return &RequestLine{
		Method:        parts[0],
		RequestTarget: parts[1],
		HttpVersion:   http_version,
	}, nil
}

/*
3.1. Method

The method token indicates the request method to be performed on the target resource.
The request method is case-sensitive.

  method         = token
*/
type HttpMethod string

const (
	HttpMethodGet    HttpMethod = "GET"
	HttpMethodHead   HttpMethod = "HEAD"
	HttpMethodPost   HttpMethod = "POST"
	HttpMethodPut    HttpMethod = "PUT"
	HttpMethodDelete HttpMethod = "OPTIONS"
	HttpMethodTrace  HttpMethod = "TRACE"
)

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
	HttpVersion  []byte
	StatusCode   []byte
	ReasonPhrase []byte
}

func ParseStatusLine(status_line []byte) (*StatusLine, error) {
	parts := bytes.Split(status_line, SP)

	return &StatusLine{
		HttpVersion:  parts[0],
		StatusCode:   parts[1],
		ReasonPhrase: parts[2],
	}, nil
}

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
type FieldLine struct {
	FieldName  []byte
	FieldValue []byte
}

// Parses Field Line
func ParseFieldLine(file_line []byte) (*FieldLine, error) {
	parts := bytes.SplitN(file_line, COLON, 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid field line: missing colon")
	}

	name := parts[0]

	// 5.1. 'No whitespace is allowed between the field name and colon.'
	if bytes.Contains(name, SP) || bytes.Contains(name, HTAB) {
		return nil, fmt.Errorf("Error parsing field line, field name. Found unexpected white space")
	}

	value := parts[1]

	// Trim OWS (space + tab only)
	value = bytes.Trim(value, " \t")

	return &FieldLine{
		FieldName:  name,
		FieldValue: value,
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
