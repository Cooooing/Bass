package model

import (
	"notify/internal/data/ent/gen"
)

type NotificationRecord struct {
	*gen.NotificationRecord

	meta *NotificationMeta
}
