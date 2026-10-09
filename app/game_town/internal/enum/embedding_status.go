package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_town/v1/enum"
)

type EmbeddingStatus string

const (
	EmbeddingStatusPending EmbeddingStatus = "pending" // pending
	EmbeddingStatusReady   EmbeddingStatus = "ready"   // ready
	EmbeddingStatusFailed  EmbeddingStatus = "failed"  // failed
)

var EmbeddingStatusMap = commonenum.NewMapping[EmbeddingStatus, v1.GameTownEmbeddingStatus](map[EmbeddingStatus]commonenum.Entry[EmbeddingStatus, v1.GameTownEmbeddingStatus]{EmbeddingStatusPending: {Proto: v1.GameTownEmbeddingStatus_GAME_TOWN_EMBEDDING_STATUS_PENDING}, EmbeddingStatusReady: {Proto: v1.GameTownEmbeddingStatus_GAME_TOWN_EMBEDDING_STATUS_READY}, EmbeddingStatusFailed: {Proto: v1.GameTownEmbeddingStatus_GAME_TOWN_EMBEDDING_STATUS_FAILED}})

func (e EmbeddingStatus) String() string { return string(e) }
