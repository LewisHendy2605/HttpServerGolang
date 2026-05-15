package response

import (
	"fmt"

	"github.com/LewisHendy2605/HttpServerGolang/internal/version"
)

var (
	StatusOK = StatusLine{
		HttpVersion: version.HTTP11,
		StatusCode:  200,
		Reason:      "OK",
	}
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
	HttpVersion version.HttpVersion
	StatusCode  uint16
	Reason      string
}

func NewStatusLine() StatusLine {
	return StatusLine{HttpVersion: version.HTTP11}
}

// Formats Request Line to string for debugging
func (sl *StatusLine) String() string {
	return fmt.Sprintf("%s %d %s", sl.HttpVersion.String(), sl.StatusCode, sl.Reason)
}

// Parses http request line
/*
func (sl *StatusLine) Parse(data []byte) (int, error) {
	index := bytes.Index(data, []byte(syntax_notation.CRLF))
	if index == -1 {
		return 0, nil
	}

	start_line := data[:index]

	parts := bytes.Split(start_line, []byte(syntax_notation.SP))
	if len(parts) != 3 {
		return 0, fmt.Errorf("invalid status line, missing space")
	}

	sl.HttpVersion = &version.HttpVersion{}
	err := sl.HttpVersion.Parse(parts[2])
	if err != nil {
		return 0, err
	}

	code, err := strconv.ParseUint(string(parts[0]), 10, 16)
	if err != nil {
		return 0, err
	}

	sl.StatusCode = uint16(code)

	sl.ReasonPhrase = string(parts[1])

	return index + len(syntax_notation.CRLF), nil
}
*/
