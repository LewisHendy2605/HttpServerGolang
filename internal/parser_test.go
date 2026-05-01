package main

import (
	"testing"
)

type HttpVersionTestTable struct {
	name    string
	input   []byte
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

func TestParseRequestLine(t *testing.T) {
	tests := []HttpVersionTestTable{
		{name: "valid", input: []byte("GET http://www.w3.org/pub/WWW/TheProject.html HTTP/1.1"), wantErr: false},
		{name: "missing method", input: []byte("http://www.w3.org/pub/WWW/TheProject.html HTTP/1.1"), wantErr: true},
		{name: "missing uri", input: []byte("GET HTTP/1.1"), wantErr: true},
		{name: "missing version", input: []byte("GET http://www.w3.org/pub/WWW/TheProject.html"), wantErr: true},
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
