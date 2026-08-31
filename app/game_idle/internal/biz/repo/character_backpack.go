package repo

import (
	"context"
	"game_idle/internal/biz/model"
)

// CharacterBackpackRepo 管理角色背包缓存，数据库只作为低频快照持久化。
type CharacterBackpackRepo interface {
	LoadItems(ctx context.Context, characterID int64) error
	MapItems(ctx context.Context, req *BackpackMapReq) (map[string]*model.CharacterItem, error)
	PersistItems(ctx context.Context, characterID int64) error
	CheckItems(ctx context.Context, req *BackpackCheckReq) (bool, error)
}

// BackpackMapReq 查询角色背包中指定物品；ItemIDs 为空时返回全部物品。
type BackpackMapReq struct {
	CharacterID int64
	ItemIDs     []string
}

// BackpackCheckReq 表示一次行动开始前的背包消耗校验。
type BackpackCheckReq struct {
	CharacterID int64
	Items       map[string]int64
}
