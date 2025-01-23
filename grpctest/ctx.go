package grpctest

import (
	"context"
	"testing"

	"github.com/mailstepcz/testutils"
	"google.golang.org/grpc/metadata"
)

// IncomingContext returns an incoming gRPC context for unit testing.
func IncomingContext(t *testing.T, m map[string]string) context.Context {
	ctx := testutils.Context(t)
	if m != nil {
		ctx = metadata.NewIncomingContext(ctx, metadata.New(m))
	}
	return ctx
}
