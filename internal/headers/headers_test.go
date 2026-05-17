package headers

import (
	"testing"

	"github.com/LewisHendy2605/HttpServerGolang/internal/syntax_notation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseFieldLine(t *testing.T) {
	// Test single header
	h := Headers{}
	index, err := h.Parse([]byte("Host: example.com" + syntax_notation.CRLF))
	require.NoError(t, err)
	assert.Equal(t, 19, index)
	assert.Equal(t, 1, h.Len())
	assert.Equal(t, "host: example.com", h.String())

	val, ok := h.Get("HOST")
	require.True(t, ok)
	assert.Equal(t, "example.com", val)

	// Test dashed header name
	h = Headers{}
	index, err = h.Parse([]byte("Content-Type: application/json" + syntax_notation.CRLF))
	require.NoError(t, err)
	assert.Equal(t, 32, index)

	val, ok = h.Get("Content-Type")
	require.True(t, ok)
	assert.Equal(t, "application/json", val)

	// Test repeated headers
	h = Headers{}
	index, err = h.Parse([]byte("Host: foo, bar\r\nHost: baz\r\n"))
	require.NoError(t, err)
	assert.Equal(t, 27, index)
	assert.Equal(t, 1, h.Len())

	val, ok = h.Get("host")
	require.True(t, ok)
	assert.Equal(t, "foo, bar, baz", val)

	// Test two different headers
	h = Headers{}
	index, err = h.Parse([]byte("Content-Type: application/json\r\nHost: example.com\r\n"))
	require.NoError(t, err)
	assert.Equal(t, 51, index)
	assert.Equal(t, 2, h.Len())
	assert.Equal(t, "content-type: application/json\r\nhost: example.com", h.String())

	val, ok = h.Get("host")
	require.True(t, ok)
	assert.Equal(t, "example.com", val)

	val, ok = h.Get("content-type")
	require.True(t, ok)
	assert.Equal(t, "application/json", val)

	// Missing colon
	h = Headers{}
	index, err = h.Parse([]byte("Host example.com" + syntax_notation.CRLF))
	require.Error(t, err)
	assert.Equal(t, 0, index)

	// Space in name
	h = Headers{}
	index, err = h.Parse([]byte("Ho st: value" + syntax_notation.CRLF))
	require.Error(t, err)
	assert.Equal(t, 0, index)

	// Missing required space in front of value
	h = Headers{}
	index, err = h.Parse([]byte("Host:value" + syntax_notation.CRLF))
	require.Error(t, err)
	assert.Equal(t, 0, index)

	// Invalid token in name
	h = Headers{}
	index, err = h.Parse([]byte("Ho\tst:value" + syntax_notation.CRLF))
	require.Error(t, err)
	assert.Equal(t, 0, index)

	// Invalid token in name
	h = Headers{}
	index, err = h.Parse([]byte("Ho\rst:value" + syntax_notation.CRLF))
	require.Error(t, err)
	assert.Equal(t, 0, index)

	// Invalid token in name
	h = Headers{}
	index, err = h.Parse([]byte("\x00:value\r\n"))
	require.Error(t, err)
	assert.Equal(t, 0, index)

}
