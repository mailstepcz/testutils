//go:build !go1.24
// +build !go1.24

package testutils

import (
	"context"
	"testing"

	"google.golang.org/grpc/metadata"
)

// Context returns the context for unit testing.
func Context(t *testing.T) context.Context {
	return context.Background()
}

// IncomingContext returns the incoming gRPC context for unit testing.
func IncomingContext(t *testing.T, m map[string]string) context.Context {
	ctx := context.Background()
	if m != nil {
		ctx = metadata.NewIncomingContext(ctx, metadata.New(m))
	}
	return ctx
}
