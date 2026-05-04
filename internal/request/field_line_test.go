package request

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseFieldLine(t *testing.T) {
	// Valid Field Line
	h := Headers{}
	parsed, err := h.Parse([]byte("Host: example.com"))
	require.NoError(t, err)
	require.NotNil(t, parsed)

	val, ok := h.Get("HOST")
	require.True(t, ok)
	require.Equal(t, "example.com", val)

	// Valid Field Line
	h = Headers{}
	parsed, err = h.Parse([]byte("Content-Type: application/json"))
	require.NoError(t, err)
	require.NotNil(t, parsed)

	val, ok = h.Get("Content-Type")
	require.True(t, ok)
	require.Equal(t, "application/json", val)

	// Missing colon
	h = Headers{}
	_, err = h.Parse([]byte("Host example.com"))
	require.Error(t, err)

	// Space in name
	h = Headers{}
	_, err = h.Parse([]byte("Ho st: value"))
	require.Error(t, err)
}
