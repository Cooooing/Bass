package usecase

import (
	"common/pkg/apperror"
	"common/pkg/constant"
	cerrors "common/proto/gen/common/errors"
	"context"
	"game_idle/internal/biz/model"
	"game_idle/internal/biz/repo"
	"game_idle/internal/config"
	"game_idle/internal/enum"
	"log/slog"
	"time"
)

// CharacterSettlementCommand 表示一次行动完成后要应用到角色状态上的结算。
type CharacterSettlementCommand struct {
	CharacterID int64
	// TaskID 与当前队首 ID 必须一致，否则说明任务已过期，不会产生任何结算。
	TaskID   string
	ActionID string
	// Items 同时包含消耗和产出，负数表示扣减，正数表示增加。
	Items      []*model.CharacterBackpackItemChange
	AbilityID  enum.Ability
	ExpReward  int64
	RemoveHead bool
}

func (c *CharacterSettlementCommand) StateCommandType() enum.StateCommandType {
	return enum.StateCommandTypeActionSettlement
}

func (c *CharacterSettlementCommand) CharacterStateID() int64 {
	return c.CharacterID
}

// CharacterSettlementHandler 在状态机内原子应用行动结算。
type CharacterSettlementHandler struct {
	logger               *slog.Logger
	stateRepo            repo.CharacterStateRepo
	actionQueueRepo      repo.CharacterActionQueueRepo
	backpackRepo         repo.CharacterBackpackRepo
	characterAbilityRepo repo.CharacterAbilityRepo
	backpackPersistLimit int64
	abilityPersistLimit  int64
}

func NewCharacterSettlementHandler(
	logger *slog.Logger,
	conf *config.Bootstrap,
	stateRepo repo.CharacterStateRepo,
	actionQueueRepo repo.CharacterActionQueueRepo,
	backpackRepo repo.CharacterBackpackRepo,
	characterAbilityRepo repo.CharacterAbilityRepo,
) *CharacterSettlementHandler {
	backpackPersistLimit := int64(100)
	if conf.GetGameIdle().GetBackpack().GetPersistThreshold() > 0 {
		backpackPersistLimit = int64(conf.GetGameIdle().GetBackpack().GetPersistThreshold())
	}
	abilityPersistLimit := int64(100)
	if conf.GetGameIdle().GetCharacterAbility().GetPersistThreshold() > 0 {
		abilityPersistLimit = int64(conf.GetGameIdle().GetCharacterAbility().GetPersistThreshold())
	}
	return &CharacterSettlementHandler{
		logger:               logger,
		stateRepo:            stateRepo,
		actionQueueRepo:      actionQueueRepo,
		backpackRepo:         backpackRepo,
		characterAbilityRepo: characterAbilityRepo,
		backpackPersistLimit: backpackPersistLimit,
		abilityPersistLimit:  abilityPersistLimit,
	}
}

func (h *CharacterSettlementHandler) Apply(ctx context.Context, command CharacterStateCommand) (*CharacterStateChangeSet, error) {
	req := command.(*CharacterSettlementCommand)

	// 事务开始前加载需要修改的状态，避免 Begin 拿到不完整草稿。
	if err := h.backpackRepo.LoadItems(ctx, req.CharacterID); err != nil {
		return nil, err
	}
	if _, err := h.actionQueueRepo.Load(ctx, req.CharacterID); err != nil {
		return nil, err
	}
	if req.AbilityID != "" && req.ExpReward > 0 {
		if _, err := h.characterAbilityRepo.Map(ctx, &repo.CharacterAbilityMapReq{
			CharacterID: req.CharacterID,
			AbilityIDs:  []enum.Ability{req.AbilityID},
		}); err != nil {
			return nil, err
		}
	}
	tx, err := h.stateRepo.Begin(ctx, req.CharacterID)
	if err != nil {
		return nil, err
	}
	committed := false
	defer func() {
		// 任何一步失败都回滚本地草稿，避免背包、能力和队列只提交一部分。
		if !committed {
			_ = tx.Rollback(ctx)
		}
	}()

	changeSet := &CharacterStateChangeSet{
		CharacterID:         req.CharacterID,
		QueueCommandApplied: true,
	}
	state := tx.Draft()
	queue := state.Queue
	oldHeadID := ""
	if len(queue.Items) > 0 {
		oldHeadID = queue.Items[0].ID
	}
	if len(queue.Items) == 0 || queue.Items[0].ID != req.TaskID || queue.Items[0].ActionID != req.ActionID {
		changeSet.QueueCommandApplied = false
		return changeSet, nil
	}

	// 同一物品可能同时出现在消耗和产出中，先聚合成净变化量。
	itemDeltas := make(map[string]int64, len(req.Items))
	for _, item := range req.Items {
		itemDeltas[item.ItemID] += item.Quantity
	}
	// 先在草稿上校验所有扣减，任意物品不足时整次结算失败。
	for itemID, quantity := range itemDeltas {
		current := int64(0)
		if item := state.BackpackItems[itemID]; item != nil {
			current = item.Quantity
		}
		if current+quantity < 0 {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_BACKPACK_INSUFFICIENT)
		}
	}
	for itemID, quantity := range itemDeltas {
		if quantity == 0 {
			continue
		}
		item := state.BackpackItems[itemID]
		if item == nil {
			item = &model.CharacterItem{
				CharacterID: req.CharacterID,
				ItemID:      itemID,
			}
			state.BackpackItems[itemID] = item
		}
		item.Quantity += quantity
		if quantity > 0 {
			item.TotalObtained += quantity
		}
		if quantity < 0 {
			item.TotalConsumed -= quantity
		}
		changeSet.ItemChanges = append(changeSet.ItemChanges, &model.ActionCompletedItemChange{
			ItemID:        itemID,
			QuantityDelta: quantity,
			QuantityAfter: item.Quantity,
		})
	}
	state.BackpackOperationCount += int64(len(itemDeltas))

	if req.AbilityID != "" && req.ExpReward > 0 {
		ability := state.Abilities[req.AbilityID]
		if ability == nil {
			ability = &model.CharacterAbility{
				CharacterID:  req.CharacterID,
				AbilityID:    req.AbilityID,
				Level:        1,
				NextLevelExp: 100,
			}
			state.Abilities[req.AbilityID] = ability
		}
		levelBefore := ability.Level
		ability.Exp += req.ExpReward
		// 当前先使用固定等级曲线，后续可替换为公式或配置表。
		for ability.Exp >= ability.NextLevelExp {
			ability.Level++
			ability.NextLevelExp = int64(ability.Level) * int64(ability.Level) * 100
		}
		changeSet.AbilityChanges = append(changeSet.AbilityChanges, &model.ActionCompletedAbilityChange{
			AbilityID: req.AbilityID.String(),
			ExpDelta:  req.ExpReward,
			ExpAfter:  ability.Exp,
		})
		if ability.Level > levelBefore {
			changeSet.AbilityLeveledUp = append(changeSet.AbilityLeveledUp, &model.AbilityLeveledUpEvent{
				CharacterID:  req.CharacterID,
				AbilityID:    req.AbilityID.String(),
				Level:        ability.Level,
				Exp:          ability.Exp,
				NextLevelExp: ability.NextLevelExp,
			})
		}
		state.AbilityOperationCount++
	}

	current := state.Queue.Items[0]
	finishCurrent := req.RemoveHead
	queueChanged := false
	if !finishCurrent && current.Times != -1 {
		current.Times--
		finishCurrent = current.Times <= 0
		queueChanged = true
	}
	changeSet.ActionTimesRemaining = current.Times
	if finishCurrent {
		state.Queue.Items = state.Queue.Items[1:]
		changeSet.ActionTimesRemaining = 0
		queueChanged = true
	}
	if queueChanged {
		// 队列变化事件只描述最终快照，调度器根据新队首决定是否继续进入时间轮。
		newHeadID := ""
		if len(state.Queue.Items) > 0 {
			newHeadID = state.Queue.Items[0].ID
		}
		changeSet.QueueChanged = &CharacterActionQueueChangedStateEvent{
			CharacterID: req.CharacterID,
			Queue:       state.Queue,
			OldHeadID:   oldHeadID,
			NewHeadID:   newHeadID,
			ChangedAt:   time.Now(),
		}
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}
	committed = true

	// 高频结算先写本地热状态，达到阈值或升级时再低频刷库。
	if h.backpackPersistLimit > 0 && state.BackpackOperationCount >= h.backpackPersistLimit {
		if err = h.backpackRepo.PersistItems(ctx, req.CharacterID); err != nil {
			h.logger.ErrorContext(ctx, "game idle persist backpack failed", constant.LogFieldErr, err, "character_id", req.CharacterID)
		}
	}
	if len(changeSet.AbilityLeveledUp) > 0 || h.abilityPersistLimit > 0 && state.AbilityOperationCount >= h.abilityPersistLimit {
		if err = h.characterAbilityRepo.Persist(ctx, req.CharacterID); err != nil {
			h.logger.ErrorContext(ctx, "game idle persist ability failed", constant.LogFieldErr, err, "character_id", req.CharacterID)
		}
	}
	return changeSet, nil
}
