package main

import "strings"

type RequestLine struct {
	Method      string
	RequestUri  string
	HttpVersion string
}

type HttpVersion struct {
}

func ParseRequestLine(request_line string) (RequestLine, error) {
	parts := strings.Split(request_line, " ")

	return RequestLine{
		Method:      parts[0],
		RequestUri:  parts[1],
		HttpVersion: parts[2],
	}, nil
}

// Parse http version from start line
// HTTP-version  = HTTP-name "/" DIGIT "." DIGIT
// HTTP-name     = %s"HTTP
/*
func ParseHttpVersion(http_version string) (HttpVersion, error) {
	parts :=
}
*/
