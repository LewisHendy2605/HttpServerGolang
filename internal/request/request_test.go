package request

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/google/uuid"
)

type TestJsonStruct struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}

func TestParseRequest(t *testing.T) {
	// Test good request
	reader := strings.NewReader("GET / HTTP/1.1\r\nHost: example.com\r\nUser-Agent: golang_application\r\nContent-Length: 5\r\n\r\nHello")
	req, err := RequestFromReader(reader)
	require.NoError(t, err)
	require.NotNil(t, req)
	require.Equal(t, "GET", req.RequestLine.Method)
	require.Equal(t, "/", req.RequestLine.RequestTarget)
	require.Equal(t, 1, req.RequestLine.HttpVersion.Major)
	require.Equal(t, 1, req.RequestLine.HttpVersion.Minor)
	require.Equal(t, []byte("Hello"), req.Body)
	require.Equal(t, "Hello", req.Text())

	val, ok := req.Headers.Get("HOST")
	require.True(t, ok)
	require.Equal(t, "example.com", val)

	val, ok = req.Headers.Get("user-agent")
	require.True(t, ok)
	require.Equal(t, "golang_application", val)

	cl, ok := req.Headers.Get("content-length")
	require.True(t, ok)
	require.Equal(t, "5", cl)

	// Test good json request
	jsonStruct := &TestJsonStruct{
		Id:      uuid.NewString(),
		Message: "Hello",
	}
	jsonString, err := json.Marshal(jsonStruct)
	require.NoError(t, err)

	reader = strings.NewReader(fmt.Sprintf("GET / HTTP/1.1\r\nContent-Type: application/json\r\nContent-Length: %d\r\n\r\n%s", len(jsonString), string(jsonString)))
	req, err = RequestFromReader(reader)
	require.NoError(t, err)
	require.NotNil(t, req)
	require.Equal(t, "GET", req.RequestLine.Method)
	require.Equal(t, "/", req.RequestLine.RequestTarget)
	require.Equal(t, 1, req.RequestLine.HttpVersion.Major)
	require.Equal(t, 1, req.RequestLine.HttpVersion.Minor)

	val, ok = req.Headers.Get("content-type")
	require.True(t, ok)
	require.Equal(t, "application/json", val)

	jsonOutput := TestJsonStruct{}
	err = req.Json(&jsonOutput)
	require.NoError(t, err)
	require.Equal(t, jsonStruct.Id, jsonOutput.Id)
	require.Equal(t, jsonStruct.Message, jsonOutput.Message)

	reader = strings.NewReader("GeT / HTTP/1.1\r\nContent-Type : application/json\r\n\r\n")
	req, err = RequestFromReader(reader)
	require.Error(t, err)
	require.Nil(t, req)
}
