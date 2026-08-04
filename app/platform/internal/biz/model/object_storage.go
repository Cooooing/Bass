package model

import (
	"platform/internal/enum"
	"time"
)

type ObjectStorage struct {
	ID                 int64
	Provider           enum.ObjectStorageProvider
	Bucket             string
	Key                string
	MimeType           string
	Size               int64
	Hash               string
	Status             enum.ObjectStorageStatus
	UploadBy           int64
	AuditCallbackReply *string
	Blocked            bool
	BlockedReason      *string
	BlockedAt          *time.Time
	BlockedBy          *int64
	CreatedAt          *time.Time
	UpdatedAt          *time.Time
	URL                string
}

type UploadToken struct {
	Key   string
	Token string
}
