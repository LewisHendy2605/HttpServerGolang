package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseHttpVersion(t *testing.T) {
	// Parse valid 1.1 http version
	raw := "HTTP/1.1"
	parsed, err := ParseHttpVersion(raw)
	require.NoError(t, err)
	require.NotNil(t, parsed)
	require.Equal(t, 1, parsed.Major)
	require.Equal(t, 1, parsed.Minor)

	// Parse valid 2.0 http version
	raw = "HTTP/2.0"
	parsed, err = ParseHttpVersion(raw)
	require.NoError(t, err)
	require.NotNil(t, parsed)
	require.Equal(t, 2, parsed.Major)
	require.Equal(t, 0, parsed.Minor)

	// Missing forward slash
	raw = "HTTP1.1"
	parsed, err = ParseHttpVersion(raw)
	require.Error(t, err)

	// Missing period
	raw = "HTTP/11"
	parsed, err = ParseHttpVersion(raw)
	require.Error(t, err)

	// Missing minor
	raw = "HTTP/1."
	parsed, err = ParseHttpVersion(raw)
	require.Error(t, err)

	// Missing major
	raw = "HTTP/.1"
	parsed, err = ParseHttpVersion(raw)
	require.Error(t, err)

	// Missing name
	raw = "/1.1"
	parsed, err = ParseHttpVersion(raw)
	require.Error(t, err)
}

func TestParseRequestLine(t *testing.T) {
	// Valid Request Line
	raw := "GET http://www.w3.org/pub/WWW/TheProject.html HTTP/1.1"
	parsed, err := ParseRequestLine(raw)
	require.NoError(t, err)
	require.Equal(t, MethodGet, parsed.Method)
	require.Equal(t, "http://www.w3.org/pub/WWW/TheProject.html", parsed.RequestTarget)
	require.Equal(t, 1, parsed.HttpVersion.Major)
	require.Equal(t, 1, parsed.HttpVersion.Minor)

	// Valid Request Line
	raw = "GET https:/example.com HTTP/1.1"
	parsed, err = ParseRequestLine(raw)
	require.NoError(t, err)
	require.NotNil(t, parsed)
	require.Equal(t, MethodGet, parsed.Method)
	require.Equal(t, "https:/example.com", parsed.RequestTarget)
	require.Equal(t, 1, parsed.HttpVersion.Major)
	require.Equal(t, 1, parsed.HttpVersion.Minor)

	// Missing Method
	raw = "http://www.w3.org/pub/WWW/TheProject.html HTTP/1.1"
	parsed, err = ParseRequestLine(raw)
	require.Error(t, err)

	// Missing request target
	raw = "GET  HTTP/1.1"
	parsed, err = ParseRequestLine(raw)
	require.Error(t, err)

	// Missing version
	raw = "GET http://www.w3.org/pub/WWW/TheProject.html "
	parsed, err = ParseRequestLine(raw)
	require.Error(t, err)

	// Invalid method
	raw = "GETT http://www.w3.org/pub/WWW/TheProject.html HTTP/1.1"
	parsed, err = ParseRequestLine(raw)
	require.Error(t, err)
}

func TestParseFieldLine(t *testing.T) {
	// Valid Field Line
	raw := "Host: example.com"
	parsed, err := ParseFieldLine(raw)
	require.NoError(t, err)
	require.NotNil(t, parsed)
	require.Equal(t, "Host", parsed.Name)
	require.Equal(t, "example.com", parsed.Value)

	// Valid Field Line
	raw = "Content-Type: application/json"
	parsed, err = ParseFieldLine(raw)
	require.NoError(t, err)
	require.NotNil(t, parsed)
	require.Equal(t, "Content-Type", parsed.Name)
	require.Equal(t, "application/json", parsed.Value)

	// Missing colon
	raw = "Host example.com"
	parsed, err = ParseFieldLine(raw)
	require.Error(t, err)

	// Space in name
	raw = "Ho st: value"
	parsed, err = ParseFieldLine(raw)
	require.Error(t, err)
}

/*
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
*/
