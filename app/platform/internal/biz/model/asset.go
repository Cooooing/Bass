package model

import (
	"platform/internal/enum"
	"time"
)

// Asset records one completed physical object. Hash is the canonical SHA-256
// content identity and is the only deduplication key.
type Asset struct {
	ID            int64
	Provider      enum.AssetProvider
	Bucket        string
	ObjectKey     string
	Hash          *string
	UploadByID    *int64
	ProviderETag  *string
	MimeType      string
	Size          int64
	Status        enum.AssetStatus
	BlockedReason *string
	BlockedAt     *time.Time
	BlockedBy     *int64
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
}
