package data

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func formatProtoTime(ts *timestamppb.Timestamp) string {
	if ts == nil {
		return ""
	}
	return ts.AsTime().Format(time.RFC3339Nano)
}
