package request

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseHttpVersion(t *testing.T) {
	// Test good http version
	version := &HttpVersion{}
	err := version.Parse([]byte("HTTP/1.1"))
	require.NoError(t, err)
	require.Equal(t, 1, version.Major)
	require.Equal(t, 1, version.Minor)

	// Missing forward slash
	version = &HttpVersion{}
	err = version.Parse([]byte("HTTP1.1"))
	require.Error(t, err)

	// Missing period
	version = &HttpVersion{}
	err = version.Parse([]byte("HTTP/11"))
	require.Error(t, err)

	// Missing minor
	version = &HttpVersion{}
	err = version.Parse([]byte("HTTP/1."))
	require.Error(t, err)

	// Missing major
	version = &HttpVersion{}
	err = version.Parse([]byte("HTTP/.1"))
	require.Error(t, err)

	// Missing name
	version = &HttpVersion{}
	err = version.Parse([]byte("/1.1"))
	require.Error(t, err)

	// Invalid name
	version = &HttpVersion{}
	err = version.Parse([]byte("htt/1.1"))
	require.Error(t, err)
}
