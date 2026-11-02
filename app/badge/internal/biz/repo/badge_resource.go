package repo

import "context"

// BadgeResourceReader retrieves bytes from Platform-confirmed asset IDs. The
// badge renderer consumes normalized bytes and never performs network I/O.
type BadgeResourceReader interface {
	Read(context.Context, int64) ([]byte, error)
}
