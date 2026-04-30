package main

import "bytes"

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
func ParseHttpVersion(http_version []byte) HttpVersion {
	parts := bytes.Split(http_version, []byte("/"))
	version := bytes.Split(parts[1], []byte("."))

	return HttpVersion{
		Name:  parts[0],
		Major: version[0],
		Minor: version[1],
	}
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
	HttpVersion   HttpVersion
}

// Parses http request line
func ParseRequestLine(request_line []byte) RequestLine {
	parts := bytes.Split(request_line, []byte(" "))

	return RequestLine{
		Method:        parts[0],
		RequestTarget: parts[1],
		HttpVersion:   ParseHttpVersion(parts[2]),
	}
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
