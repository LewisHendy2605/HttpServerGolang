package parser

import (
	"io"
	"strings"
	"testing"
)

type HttpVersionTestTable struct {
	name    string
	input   []byte
	wantErr bool
}

type HttpMethodTestTable struct {
	name  string
	input HttpMethod
	valid bool
}

type HttpRequestTestTable struct {
	name    string
	input   io.Reader
	wantErr bool
}

func TestParseHttpVersion(t *testing.T) {
	tests := []HttpVersionTestTable{
		{name: "valid", input: []byte("HTTP/1.1"), wantErr: false},
		{name: "missing slash", input: []byte("HTTP1.1"), wantErr: true},
		{name: "missing period", input: []byte("HTTP/11"), wantErr: true},
		{name: "missing minor", input: []byte("HTTP/1."), wantErr: true},
		{name: "missing major", input: []byte("HTTP/.1"), wantErr: true},
		{name: "missing prefix", input: []byte("/.1"), wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseHttpVersion(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("got err=%v, wantErr=%v", err, tt.wantErr)
			}
		})
	}
}

func TestIsValidHttpMethod(t *testing.T) {
	tests := []HttpMethodTestTable{
		{name: "valid", input: HttpMethod("GET"), valid: true},
		{name: "valid", input: HttpMethod("POST"), valid: true},
		{name: "valid", input: HttpMethod("PUT"), valid: true},
		{name: "valid", input: HttpMethod("PATCH"), valid: true},
		{name: "valid", input: HttpMethod("DELETE"), valid: true},
		{name: "valid", input: HttpMethod("OPTIONS"), valid: true},
		{name: "valid", input: HttpMethod("TRACE"), valid: true},
		{name: "valid", input: HttpMethod("HEAD"), valid: true},
		{name: "valid", input: HttpMethod("Get"), valid: false},
		{name: "valid", input: HttpMethod("get"), valid: false},
		{name: "missing method", input: HttpMethod("GETT"), valid: false},
		{name: "missing uri", input: HttpMethod("TEST"), valid: false},
		{name: "missing uri", input: HttpMethod(""), valid: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid := IsValidHttpMethod(tt.input)
			if (valid) != tt.valid {
				t.Fatalf("got: %v, expected: %v", valid, tt.valid)
			}
		})
	}
}

func TestParseRequestLine(t *testing.T) {
	tests := []HttpVersionTestTable{
		{name: "valid", input: []byte("GET http://www.w3.org/pub/WWW/TheProject.html HTTP/1.1"), wantErr: false},
		{name: "valid", input: []byte("GET https:/example.com HTTP/1.1"), wantErr: false},
		{name: "missing method", input: []byte("http://www.w3.org/pub/WWW/TheProject.html HTTP/1.1"), wantErr: true},
		{name: "missing uri", input: []byte("GET  HTTP/1.1"), wantErr: true},
		{name: "missing version", input: []byte("GET http://www.w3.org/pub/WWW/TheProject.html"), wantErr: true},
		{name: "invalid method", input: []byte("TEST http://www.w3.org/pub/WWW/TheProject.html HTTP/1.1"), wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseRequestLine(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("got err=%v, wantErr=%v", err, tt.wantErr)
			}
		})
	}
}

func TestParseFieldLine(t *testing.T) {
	tests := []HttpVersionTestTable{
		{name: "valid", input: []byte("Host: example.com"), wantErr: false},
		{name: "valid", input: []byte("Content-Type: text/plain"), wantErr: false},
		{name: "missing colon", input: []byte("Host example.com"), wantErr: true},
		{name: "space in name", input: []byte("Ho st: value"), wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseFieldLine(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("got err=%v, wantErr=%v", err, tt.wantErr)
			}
		})
	}
}

func TestParseHttpRequest(t *testing.T) {
	tests := []HttpRequestTestTable{
		{name: "valid", input: strings.NewReader("GET https:/example.com HTTP/1.1\r\nContent-Type: text/plain\r\n\r\n"), wantErr: false},
		{name: "valid", input: strings.NewReader("GET https:/example.com HTTP/1.1\r\nContent-Type: text/plain\r\n\r\nHello"), wantErr: false},
		{name: "valid", input: strings.NewReader("GET https:/example.com HTTP/1.1\r\n\r\n"), wantErr: false},
		{name: "valid", input: strings.NewReader("GET https:/example.com \r\n\r\n"), wantErr: true},
		{name: "valid", input: strings.NewReader(" https:/example.com HTTP/1.1\r\n\r\n"), wantErr: true},
		{name: "valid", input: strings.NewReader("GET  HTTP/1.1\r\n\r\n"), wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseHttpRequest(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("got err=%v, wantErr=%v", err, tt.wantErr)
			}
		})
	}
}
