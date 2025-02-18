package testutils

import (
	"context"
	"testing"
)

func getCtx(t *testing.T) context.Context {
	return t.Context()
}
