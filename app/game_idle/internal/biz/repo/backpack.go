package repo

import (
	"context"
	"game_idle/internal/biz/model"
)

// BackpackRepo 管理角色背包缓存，数据库只作为低频快照持久化。
type BackpackRepo interface {
	MapItems(ctx context.Context, req *BackpackMapReq) (map[string]*model.CharacterItem, error)
	PersistItems(ctx context.Context, characterID int64) error
	CheckItems(ctx context.Context, req *BackpackCheckReq) error
	ChangeItems(ctx context.Context, req *BackpackChangeReq) (map[string]int64, error)
}

// BackpackMapReq 查询角色背包中指定物品；ItemIDs 为空时返回全部物品。
type BackpackMapReq struct {
	CharacterID int64
	ItemIDs     []string
}

// BackpackChangeReq 表示一次背包批量增减。
type BackpackChangeReq struct {
	CharacterID int64
	Items       []*model.BackpackItemChange
}

// BackpackCheckReq 表示一次行动开始前的背包消耗校验。
type BackpackCheckReq struct {
	CharacterID int64
	Items       map[string]int64
}
