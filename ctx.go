package testutils

import (
	"context"
	"testing"
)

// Context returns the context for unit testing.
func Context(t *testing.T) context.Context {
	return getCtx(t)
}
