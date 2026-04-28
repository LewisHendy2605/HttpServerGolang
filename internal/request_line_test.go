package main

import (
	"fmt"
	"testing"
)

func TestParseRequestLine(t *testing.T) {
	method := "GET"
	uri := "http://www.w3.org/pub/WWW/TheProject.html"
	version := "HTTP/1.1"

	requestLine, _ := ParseRequestLine(fmt.Sprintf("%s %s %s", method, uri, version))

	if requestLine.Method != method {
		t.Fatalf("unexpected http method. expected: %s, got %s", method, requestLine.Method)
	}
	if requestLine.RequestUri != uri {
		t.Fatalf("unexpected http method. expected: %s, got %s", uri, requestLine.RequestUri)
	}
	if requestLine.HttpVersion != version {
		t.Fatalf("unexpected http method. expected: %s, got %s", version, requestLine.HttpVersion)
	}
}
