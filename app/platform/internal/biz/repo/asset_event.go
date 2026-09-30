package repo

import (
	"context"
	"errors"
)

// ErrInvalidAssetUploadEvent marks an event that can never produce an asset
// record. The consumer terminates these messages instead of retrying them.
var ErrInvalidAssetUploadEvent = errors.New("invalid asset upload event")

// AssetUploadEvent is the provider notification normalized at the data
// boundary. It intentionally contains only the object identity needed by the
// asset directory; provider-specific notification fields remain in data.
type AssetUploadEvent struct {
	Bucket    string
	ObjectKey string
}

type AssetUploadEventHandler func(context.Context, *AssetUploadEvent) error

// AssetUploadEventConsumer provides durable object-created notifications.
// Implementations must ACK only completed or explicitly discarded events.
type AssetUploadEventConsumer interface {
	Consume(ctx context.Context, handler AssetUploadEventHandler) error
	Stop(ctx context.Context) error
}
