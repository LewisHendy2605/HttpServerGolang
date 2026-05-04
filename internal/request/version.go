package request

import (
	"bytes"
	"fmt"
	"strconv"
)

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
	Major int
	Minor int
}

// Parses http version from start line
func (v *HttpVersion) Parse(data []byte) error {
	parts := bytes.Split(data, []byte(SLASH))
	if len(parts) != 2 {
		return fmt.Errorf("invalid http version, missing forward slash")
	}

	name := bytes.ToLower(parts[0])
	if len(name) == 0 || !bytes.Equal(name, []byte("http")) {
		return fmt.Errorf("invalid http version, invalid or missing name")
	}

	version := bytes.Split(parts[1], []byte(DOT))
	if len(version) != 2 {
		return fmt.Errorf("invalid http version, missing period")
	}

	major := string(version[0])
	if len(major) == 0 {
		return fmt.Errorf("invalid http version, missing major")
	}

	minor := string(version[1])
	if len(minor) == 0 {
		return fmt.Errorf("invalid http version, missing minor")
	}

	var err error

	v.Major, err = strconv.Atoi(major)
	if err != nil {
		return fmt.Errorf("invalid http version, invalid major int value")
	}

	v.Minor, err = strconv.Atoi(minor)
	if err != nil {
		return fmt.Errorf("invalid http version, invalid minor int value")
	}

	return nil
}
