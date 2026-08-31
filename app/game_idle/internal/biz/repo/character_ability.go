package repo

import (
	"context"
	"game_idle/internal/biz/model"
	"game_idle/internal/enum"
)

// CharacterAbilityRepo 管理角色能力缓存，数据库只保存低频快照。
type CharacterAbilityRepo interface {
	Map(ctx context.Context, req *CharacterAbilityMapReq) (map[enum.Ability]*model.CharacterAbility, error)
	Persist(ctx context.Context, characterID int64) error
}

// CharacterAbilityMapReq 查询角色能力；AbilityIDs 为空时返回全部已载入能力。
type CharacterAbilityMapReq struct {
	CharacterID int64
	AbilityIDs  []enum.Ability
}
