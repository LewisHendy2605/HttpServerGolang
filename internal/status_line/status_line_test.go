package status_line

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatusLine(t *testing.T) {
	// Good
	assert.Equal(t, StatusOK.String(), "HTTP/1.1 200 OK")
}
