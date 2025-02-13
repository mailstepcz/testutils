package grpctest

import (
	"github.com/mailstepcz/grpcerr"
	"github.com/mailstepcz/httperr"
)

// IsHTTPError checks whether the error is an HTTP error.
func IsHTTPError(err error) bool {
	_, ok := err.(httperr.HTTPError)
	return ok
}

// IsGRPCError checks whether the error is a gRPC error.
func IsGRPCError(err error) bool {
	_, ok := err.(grpcerr.Convertible)
	return ok
}
