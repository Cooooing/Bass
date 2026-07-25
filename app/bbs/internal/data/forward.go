package data

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type protoTimeFormatter struct{}

func (protoTimeFormatter) formatProtoTime(ts *timestamppb.Timestamp) string {
	if ts == nil {
		return ""
	}
	return ts.AsTime().Format(time.RFC3339Nano)
}
