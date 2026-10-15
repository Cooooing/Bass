package scheduler

import (
	"strings"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// MarshalPayload uses proto field names so persisted task JSON stays stable and reviewable.
func MarshalPayload(message proto.Message) (string, error) {
	payload, err := protojson.MarshalOptions{UseProtoNames: true}.Marshal(message)
	return string(payload), err
}

// MustMarshalPayload serializes static code-owned default task payloads.
func MustMarshalPayload(message proto.Message) string {
	payload, err := MarshalPayload(message)
	if err != nil {
		panic(err)
	}
	return payload
}

// UnmarshalPayload rejects unknown fields to keep task producers and handlers in sync.
func UnmarshalPayload(payload string, message proto.Message) error {
	payload = strings.TrimSpace(payload)
	if payload == "" {
		payload = "{}"
	}
	return protojson.Unmarshal([]byte(payload), message)
}
