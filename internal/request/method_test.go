package request

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMethod(t *testing.T) {
	// Test GET
	require.True(t, IsMethod("GET"))

	// Test HEAD
	require.True(t, IsMethod("HEAD"))

	// Test POST
	require.True(t, IsMethod("POST"))

	// Test PUT
	require.True(t, IsMethod("PUT"))

	// Test PATCH
	require.True(t, IsMethod("PATCH"))

	// Test DELETE
	require.True(t, IsMethod("DELETE"))

	// Test OPTIONS
	require.True(t, IsMethod("OPTIONS"))

	// Test TRACE
	require.True(t, IsMethod("TRACE"))

	// Test bad method
	require.False(t, IsMethod("TEST"))

	// Test bad method
	require.False(t, IsMethod(""))

	// Test bad method
	require.False(t, IsMethod("get"))

	// Test bad method
	require.False(t, IsMethod("gEt"))

	// Test bad method
	require.False(t, IsMethod("GEt"))

	// Test bad method
	require.False(t, IsMethod("Get"))
}
