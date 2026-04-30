package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestParseRequestLine(t *testing.T) {
	method := []byte("GET")
	uri := []byte("http://www.w3.org/pub/WWW/TheProject.html")
	major := []byte("1")
	minor := []byte("1")
	version := fmt.Appendf(nil, "HTTP/%s.%s", major, minor)

	requestLine := fmt.Appendf(nil, "%s %s %s", method, uri, version)

	parsed := ParseRequestLine(requestLine)

	if !bytes.Equal(parsed.Method, method) {
		t.Fatalf("unexpected http method. expected: %s, got %s", method, parsed.Method)
	}
	if !bytes.Equal(parsed.RequestTarget, uri) {
		t.Fatalf("unexpected http method. expected: %s, got %s", uri, parsed.RequestTarget)
	}
	if !bytes.Equal(parsed.HttpVersion.Major, major) {
		t.Fatalf("unexpected http method. expected: %s, got %v", major, parsed.HttpVersion.Major)
	}
	if !bytes.Equal(parsed.HttpVersion.Minor, minor) {
		t.Fatalf("unexpected http method. expected: %s, got %v", minor, parsed.HttpVersion.Minor)
	}
}
