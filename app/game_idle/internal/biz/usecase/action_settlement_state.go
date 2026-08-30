package usecase

import (
	"common/pkg/constant"
	"context"
	"game_idle/internal/biz/model"
	"game_idle/internal/biz/repo"
	"game_idle/internal/enum"
	"log/slog"
)

// ActionSettlementCommand 表示一次行动完成后要应用到角色状态上的结算。
type ActionSettlementCommand struct {
	CharacterID int64
	ActionID    string
	Items       []*model.BackpackItemChange
	AbilityID   enum.Ability
	ExpReward   int64
}

func (c *ActionSettlementCommand) StateCommandType() StateCommandType {
	return StateCommandTypeActionSettlement
}

// ActionSettlementHandler 原子应用行动结算，并把结果转换为状态机领域事件。
type ActionSettlementHandler struct {
	logger         *slog.Logger
	settlementRepo repo.ActionSettlementRepo
}

func NewActionSettlementHandler(logger *slog.Logger, settlementRepo repo.ActionSettlementRepo) *ActionSettlementHandler {
	return &ActionSettlementHandler{
		logger:         logger,
		settlementRepo: settlementRepo,
	}
}

func (r *ActionSettlementHandler) Apply(ctx context.Context, command StateCommand) (*StateChangeSet, error) {
	req := command.(*ActionSettlementCommand)
	settlement, err := r.settlementRepo.Apply(ctx, &repo.ActionSettlementReq{
		CharacterID: req.CharacterID,
		Items:       req.Items,
		AbilityID:   req.AbilityID,
		ExpReward:   req.ExpReward,
	})
	if err != nil && settlement == nil {
		return nil, err
	}
	if err != nil {
		r.logger.ErrorContext(ctx, "game idle action settlement persist failed", constant.LogFieldErr, err, "character_id", req.CharacterID)
	}
	return r.changeSet(req.CharacterID, settlement), nil
}

func (r *ActionSettlementHandler) changeSet(characterID int64, settlement *model.ActionSettlement) *StateChangeSet {
	changeSet := &StateChangeSet{
		CharacterID:    characterID,
		ItemChanges:    settlement.ItemChanges,
		AbilityChanges: settlement.AbilityChanges,
	}
	if len(settlement.ItemChanges) > 0 {
		changeSet.Events = append(changeSet.Events, &StateEvent{
			Type: StateEventTypeItemsChanged,
			ItemsChanged: &ItemsChangedStateEvent{
				CharacterID: characterID,
				Changes:     settlement.ItemChanges,
			},
		})
	}
	if len(settlement.AbilityChanges) > 0 {
		changeSet.Events = append(changeSet.Events, &StateEvent{
			Type: StateEventTypeAbilityExpGained,
			AbilityExpGained: &AbilityExpGainedStateEvent{
				CharacterID: characterID,
				Changes:     settlement.AbilityChanges,
			},
		})
	}
	if settlement.AbilityLeveledUp != nil {
		changeSet.Events = append(changeSet.Events, &StateEvent{
			Type:             StateEventTypeAbilityLeveledUp,
			AbilityLeveledUp: settlement.AbilityLeveledUp,
		})
	}
	return changeSet
}
