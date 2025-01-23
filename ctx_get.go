//go:build go1.24
// +build go1.24

package testutils

import (
	"context"
	"testing"
)

func getCtx(t *testing.T) context.Context {
	return t.Context()
}
