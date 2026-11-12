package server

import (
	"context"
	"strings"

	"github.com/go-kratos/kratos/v3/transport"
	"google.golang.org/grpc/metadata"
)

func GetHeader(ctx context.Context, key string) string {
	var v string
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if vals := md.Get(strings.ToLower(key)); len(vals) > 0 {
			v = vals[0]
		}
	}

	if tr, ok := transport.FromServerContext(ctx); ok {
		if h := tr.RequestHeader(); h != nil {
			if s := h.Get(key); s != "" {
				v = s
			}
		}
	}
	return v
}
