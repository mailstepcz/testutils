package testcond

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Equal compares the two arguments for equality.
func Equal[T any](t *testing.T, x, y T) {
	require.Equal(t, x, y)
}
