//go:build !go1.24
// +build !go1.24

package testutils

import (
	"context"
	"testing"
)

func getCtx(_ *testing.T) context.Context {
	return context.Background()
}

// Version ...
func Version() string {
	return "<1.24"
}
