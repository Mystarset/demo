package grpc2fuse

import (
	"testing"

	"github.com/hanwen/go-fuse/v2/fuse"
)

func TestDealGrpcError(t *testing.T) {
	tests := []struct {
		name     string
		method   string
		err      error
		expected fuse.Status
	}{
		{
			name:     "nil error",
			method:   "Release",
			err:      nil,
			expected: fuse.OK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			status := dealGrpcError(tt.method, tt.err)
			if status != tt.expected {
				t.Errorf("dealGrpcError() = %v, want %v", status, tt.expected)
			}
		})
	}
}
