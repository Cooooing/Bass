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

// UnmarshalPayload 解析持久化的任务 payload。
func UnmarshalPayload(payload string, message proto.Message) error {
	payload = strings.TrimSpace(payload)
	if payload == "" {
		// protojson 将空输入视为语法错误，空 payload 表示未设置任何字段。
		payload = "{}"
	}
	return protojson.Unmarshal([]byte(payload), message)
}
