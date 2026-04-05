package model

import (
	v1 "common/gen/infra/v1"
	"infra/internal/data/ent/gen"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type ObjectStorage struct {
	*gen.ObjectStorage
}

func (o *ObjectStorage) ConvertToRpc() *v1.Oss {
	v := &v1.Oss{
		CreatedAt:          timestamppb.New(*o.CreatedAt),
		UpdatedAt:          timestamppb.New(*o.UpdatedAt),
		Id:                 o.ID,
		Provider:           o.Provider,
		Bucket:             o.Bucket,
		Key:                o.Key,
		MimeType:           o.MimeType,
		Size:               o.Size,
		Hash:               o.Hash,
		AuditCallbackReply: o.AuditCallbackReply,
		Blocked:            o.Blocked,
		BlockedReason:      o.BlockedReason,
		BlockedBy:          o.BlockedBy,
		BlockedByName:      o.BlockedByName,
	}
	if o.BlockedAt != nil {
		v.BlockedAt = timestamppb.New(*o.BlockedAt)
	}
	return v
}

type UploadToken struct {
	Key   string
	Token string
}
