package field_line

import (
	"testing"

	"github.com/LewisHendy2605/HttpServerGolang/internal/syntax_notation"
	"github.com/stretchr/testify/require"
)

func TestParseFieldLine(t *testing.T) {
	// Valid Field Line
	h := Headers{}
	index, err := h.Parse([]byte("Host: example.com" + syntax_notation.CRLF))
	require.NoError(t, err)
	require.Equal(t, 19, index)

	val, ok := h.Get("HOST")
	require.True(t, ok)
	require.Equal(t, "example.com", val)

	// Valid Field Line
	h = Headers{}
	index, err = h.Parse([]byte("Content-Type: application/json" + syntax_notation.CRLF))
	require.NoError(t, err)
	require.Equal(t, 32, index)

	val, ok = h.Get("Content-Type")
	require.True(t, ok)
	require.Equal(t, "application/json", val)

	// Valid Field Line
	h = Headers{}
	index, err = h.Parse([]byte("Host: foo, bar\r\nHost: baz\r\n"))
	require.NoError(t, err)
	require.Equal(t, 27, index)

	val, ok = h.Get("host")
	require.True(t, ok)
	require.Equal(t, "foo, bar, baz", val)

	// Missing colon
	h = Headers{}
	index, err = h.Parse([]byte("Host example.com" + syntax_notation.CRLF))
	require.Error(t, err)
	require.Equal(t, 0, index)

	// Space in name
	h = Headers{}
	index, err = h.Parse([]byte("Ho st: value" + syntax_notation.CRLF))
	require.Error(t, err)
	require.Equal(t, 0, index)

	// Missing required space in front of value
	h = Headers{}
	index, err = h.Parse([]byte("Host:value" + syntax_notation.CRLF))
	require.Error(t, err)
	require.Equal(t, 0, index)

	// Invalid token in name
	h = Headers{}
	index, err = h.Parse([]byte("Ho\tst:value" + syntax_notation.CRLF))
	require.Error(t, err)
	require.Equal(t, 0, index)

	// Invalid token in name
	h = Headers{}
	index, err = h.Parse([]byte("Ho\rst:value" + syntax_notation.CRLF))
	require.Error(t, err)
	require.Equal(t, 0, index)

	// Invalid token in name
	h = Headers{}
	index, err = h.Parse([]byte("\x00:value\r\n"))
	require.Error(t, err)
	require.Equal(t, 0, index)

}
