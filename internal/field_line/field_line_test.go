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

	// Missing colon
	h = Headers{}
	_, err = h.Parse([]byte("Host example.com" + syntax_notation.CRLF))
	require.Error(t, err)

	// Space in name
	h = Headers{}
	_, err = h.Parse([]byte("Ho st: value" + syntax_notation.CRLF))
	require.Error(t, err)
}
