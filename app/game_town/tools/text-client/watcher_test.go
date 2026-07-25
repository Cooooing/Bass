package main

import (
	"errors"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestIsRetryableWatcherError(
	t *testing.T,
) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "server max age",
			err:  status.Error(codes.Unavailable, "received prior goaway: max_age"),
			want: true,
		},
		{
			name: "deadline",
			err:  status.Error(codes.DeadlineExceeded, "deadline"),
			want: true,
		},
		{
			name: "not found",
			err:  status.Error(codes.NotFound, "world"),
			want: false,
		},
		{
			name: "plain error",
			err:  errors.New("broken"),
			want: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := isRetryableWatcherError(test.err); got != test.want {
				t.Fatalf("isRetryableWatcherError() = %v, want %v", got, test.want)
			}
		})
	}
}
