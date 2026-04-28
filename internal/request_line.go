package main

import "strings"

type RequestLine struct {
	Method      string
	RequestUri  string
	HttpVersion string
}

func ParseRequestLine(request_line string) (RequestLine, error) {
	parts := strings.Split(request_line, " ")

	return RequestLine{
		Method:      parts[0],
		RequestUri:  parts[1],
		HttpVersion: parts[2],
	}, nil
}
