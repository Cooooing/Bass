package repo

import (
	"context"
	"game_idle/internal/biz/model"
	"game_idle/internal/enum"
)

// ActionSettlementRepo 原子完成行动结算产生的库存和经验变更。
type ActionSettlementRepo interface {
	Apply(ctx context.Context, req *ActionSettlementReq) (*model.ActionSettlement, error)
}

// ActionSettlementReq 表示一次行动完成后需要落到角色状态上的变化。
type ActionSettlementReq struct {
	CharacterID int64
	Items       []*model.BackpackItemChange
	AbilityID   enum.Ability
	ExpReward   int64
}
