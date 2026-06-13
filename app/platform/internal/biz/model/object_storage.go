package model

import "time"

type ObjectStorage struct {
	ID                 int64
	Provider           string
	Bucket             string
	Key                string
	MimeType           string
	Size               int64
	Hash               string
	UploadBy           int64
	UploadByName       string
	AuditCallbackReply *string
	Blocked            bool
	BlockedReason      *string
	BlockedAt          *time.Time
	BlockedBy          *int64
	BlockedByName      *string
	CreatedAt          *time.Time
	UpdatedAt          *time.Time
}

type UploadToken struct {
	Key   string
	Token string
}
