package repo

import (
	commonclient "common/pkg/client"
	"context"
	"fmt"
	"game_idle/internal/biz/model"
	bizrepo "game_idle/internal/biz/repo"
	"game_idle/internal/config"
	"game_idle/internal/enum"
	"strconv"
	"strings"

	"github.com/redis/go-redis/v9"
)

var _ bizrepo.ActionSettlementRepo = (*ActionSettlementRepo)(nil)

type ActionSettlementRepo struct {
	redisClient                         *commonclient.RedisClient
	backpackRepo                        bizrepo.BackpackRepo
	characterAbilityRepo                bizrepo.CharacterAbilityRepo
	backpackPersistThreshold            int64
	abilityPersistThreshold             int64
	backpackQuantityRedisKeyFormat      string
	backpackObtainedRedisKeyFormat      string
	backpackConsumedRedisKeyFormat      string
	backpackOperationRedisKeyFormat     string
	abilityExpRedisKeyFormat            string
	abilityLevelRedisKeyFormat          string
	abilityNextLevelExpRedisKeyFormat   string
	abilityOperationCountRedisKeyFormat string
	settlementScript                    *redis.Script
}

func NewActionSettlementRepo(
	conf *config.Bootstrap,
	redisClient *commonclient.RedisClient,
	backpackRepo bizrepo.BackpackRepo,
	characterAbilityRepo bizrepo.CharacterAbilityRepo,
) bizrepo.ActionSettlementRepo {
	backpackPersistThreshold := int64(100)
	if conf.GetGameIdle().GetBackpack().GetPersistThreshold() > 0 {
		backpackPersistThreshold = int64(conf.GetGameIdle().GetBackpack().GetPersistThreshold())
	}
	abilityPersistThreshold := int64(100)
	if conf.GetGameIdle().GetCharacterAbility().GetPersistThreshold() > 0 {
		abilityPersistThreshold = int64(conf.GetGameIdle().GetCharacterAbility().GetPersistThreshold())
	}
	return &ActionSettlementRepo{
		redisClient:                         redisClient,
		backpackRepo:                        backpackRepo,
		characterAbilityRepo:                characterAbilityRepo,
		backpackPersistThreshold:            backpackPersistThreshold,
		abilityPersistThreshold:             abilityPersistThreshold,
		backpackQuantityRedisKeyFormat:      "game_idle:character_backpack:%d:quantity",
		backpackObtainedRedisKeyFormat:      "game_idle:character_backpack:%d:total_obtained",
		backpackConsumedRedisKeyFormat:      "game_idle:character_backpack:%d:total_consumed",
		backpackOperationRedisKeyFormat:     "game_idle:character_backpack:%d:operation_count",
		abilityExpRedisKeyFormat:            "game_idle:character_ability:%d:exp",
		abilityLevelRedisKeyFormat:          "game_idle:character_ability:%d:level",
		abilityNextLevelExpRedisKeyFormat:   "game_idle:character_ability:%d:next_level_exp",
		abilityOperationCountRedisKeyFormat: "game_idle:character_ability:%d:operation_count",
		settlementScript: redis.NewScript(`
local function next_level_exp(level)
  return level * level * 100
end

local item_count = tonumber(ARGV[1])
local index = 2
local item_deltas = {}
local item_order = {}
for item_index = 1, item_count do
  local item_id = ARGV[index]
  local quantity = tonumber(ARGV[index + 1])
  if item_deltas[item_id] == nil then
    item_deltas[item_id] = 0
    table.insert(item_order, item_id)
  end
  item_deltas[item_id] = item_deltas[item_id] + quantity
  index = index + 2
end
for _, item_id in ipairs(item_order) do
  local current = tonumber(redis.call('HGET', KEYS[1], item_id) or '0')
  if current + item_deltas[item_id] < 0 then
    return redis.error_reply('backpack_insufficient')
  end
end
for _, item_id in ipairs(item_order) do
  local quantity = item_deltas[item_id]
  if quantity ~= 0 then
    redis.call('HINCRBY', KEYS[1], item_id, quantity)
    if quantity > 0 then
      redis.call('HINCRBY', KEYS[2], item_id, quantity)
    else
      redis.call('HINCRBY', KEYS[3], item_id, -quantity)
    end
  end
end
local ability_id = ARGV[index]
local exp_reward = tonumber(ARGV[index + 1])
local result = {}
table.insert(result, redis.call('INCRBY', KEYS[4], item_count))
if ability_id ~= '' and exp_reward > 0 then
  local current_level = tonumber(redis.call('HGET', KEYS[6], ability_id) or '1')
  local next_exp_text = redis.call('HGET', KEYS[7], ability_id)
  local next_exp_target = tonumber(next_exp_text or next_level_exp(current_level))
  local exp_after = tonumber(redis.call('HINCRBY', KEYS[5], ability_id, exp_reward))
  local new_level = current_level
  local next_exp_after = next_exp_target
  if exp_after >= next_exp_target then
    while exp_after >= next_exp_after do
      new_level = new_level + 1
      next_exp_after = next_level_exp(new_level)
    end
    redis.call('HSET', KEYS[6], ability_id, new_level)
    redis.call('HSET', KEYS[7], ability_id, next_exp_after)
  elseif next_exp_text == false then
    redis.call('HSET', KEYS[7], ability_id, next_exp_after)
  end
  local ability_count = redis.call('INCR', KEYS[8])
  table.insert(result, ability_id)
  table.insert(result, exp_reward)
  table.insert(result, exp_after)
  table.insert(result, ability_count)
  table.insert(result, new_level)
  table.insert(result, next_exp_after)
  if new_level > current_level then
    table.insert(result, 1)
  else
    table.insert(result, 0)
  end
else
  table.insert(result, '')
  table.insert(result, 0)
  table.insert(result, 0)
  table.insert(result, 0)
  table.insert(result, 0)
  table.insert(result, 0)
  table.insert(result, 0)
end
for _, item_id in ipairs(item_order) do
  table.insert(result, item_id)
  table.insert(result, item_deltas[item_id])
  table.insert(result, redis.call('HGET', KEYS[1], item_id) or '0')
end
return result
`),
	}
}

func (r *ActionSettlementRepo) Apply(
	ctx context.Context,
	req *bizrepo.ActionSettlementReq,
) (*model.ActionSettlement, error) {
	if err := r.backpackRepo.EnsureLoaded(ctx, req.CharacterID); err != nil {
		return nil, err
	}
	if req.AbilityID != "" && req.ExpReward > 0 {
		if _, err := r.characterAbilityRepo.Map(ctx, &bizrepo.CharacterAbilityMapReq{
			CharacterID: req.CharacterID,
			AbilityIDs:  []enum.Ability{req.AbilityID},
		}); err != nil {
			return nil, err
		}
	}
	args := make([]any, 0, len(req.Items)*2+3)
	args = append(args, len(req.Items))
	for _, item := range req.Items {
		args = append(args, item.ItemID, item.Quantity)
	}
	args = append(args, req.AbilityID.String(), req.ExpReward)
	result, err := r.settlementScript.Run(
		ctx,
		r.redisClient.Client,
		[]string{
			r.backpackQuantityRedisKey(req.CharacterID),
			r.backpackObtainedRedisKey(req.CharacterID),
			r.backpackConsumedRedisKey(req.CharacterID),
			r.backpackOperationRedisKey(req.CharacterID),
			r.abilityExpRedisKey(req.CharacterID),
			r.abilityLevelRedisKey(req.CharacterID),
			r.abilityNextLevelExpRedisKey(req.CharacterID),
			r.abilityOperationCountRedisKey(req.CharacterID),
		},
		args...,
	).Result()
	if err != nil {
		if strings.Contains(err.Error(), "backpack_insufficient") {
			return nil, model.ErrBackpackInsufficient
		}
		return nil, err
	}
	values, ok := result.([]any)
	if !ok || len(values) == 0 {
		return nil, fmt.Errorf("game idle action settlement result invalid")
	}
	backpackOperateCount, err := strconv.ParseInt(fmt.Sprint(values[0]), 10, 64)
	if err != nil {
		return nil, err
	}
	itemStartIndex := 8
	settlement := &model.ActionSettlement{}
	if fmt.Sprint(values[1]) != "" {
		expDelta, err := strconv.ParseInt(fmt.Sprint(values[2]), 10, 64)
		if err != nil {
			return nil, err
		}
		expAfter, err := strconv.ParseInt(fmt.Sprint(values[3]), 10, 64)
		if err != nil {
			return nil, err
		}
		levelAfter, err := strconv.ParseInt(fmt.Sprint(values[5]), 10, 32)
		if err != nil {
			return nil, err
		}
		nextLevelExp, err := strconv.ParseInt(fmt.Sprint(values[6]), 10, 64)
		if err != nil {
			return nil, err
		}
		settlement.AbilityChanges = append(settlement.AbilityChanges, &model.ActionCompletedAbilityChange{
			AbilityID: fmt.Sprint(values[1]),
			ExpDelta:  expDelta,
			ExpAfter:  expAfter,
		})
		if fmt.Sprint(values[7]) == "1" {
			settlement.AbilityLeveledUp = &model.AbilityLeveledUpEvent{
				CharacterID:  req.CharacterID,
				AbilityID:    fmt.Sprint(values[1]),
				Level:        int32(levelAfter),
				Exp:          expAfter,
				NextLevelExp: nextLevelExp,
			}
		}
	}
	for index := itemStartIndex; index+2 < len(values); index += 3 {
		quantityDelta, err := strconv.ParseInt(fmt.Sprint(values[index+1]), 10, 64)
		if err != nil {
			return nil, err
		}
		quantityAfter, err := strconv.ParseInt(fmt.Sprint(values[index+2]), 10, 64)
		if err != nil {
			return nil, err
		}
		settlement.ItemChanges = append(settlement.ItemChanges, &model.ActionCompletedItemChange{
			ItemID:        fmt.Sprint(values[index]),
			QuantityDelta: quantityDelta,
			QuantityAfter: quantityAfter,
		})
	}
	if r.backpackPersistThreshold > 0 && backpackOperateCount >= r.backpackPersistThreshold {
		if err = r.backpackRepo.PersistItems(ctx, req.CharacterID); err != nil {
			return settlement, err
		}
	}
	abilityOperateCount, err := strconv.ParseInt(fmt.Sprint(values[4]), 10, 64)
	if err != nil {
		return nil, err
	}
	if settlement.AbilityLeveledUp != nil || r.abilityPersistThreshold > 0 && abilityOperateCount >= r.abilityPersistThreshold {
		if err = r.characterAbilityRepo.Persist(ctx, req.CharacterID); err != nil {
			return settlement, err
		}
	}
	return settlement, nil
}

func (r *ActionSettlementRepo) backpackQuantityRedisKey(characterID int64) string {
	return fmt.Sprintf(r.backpackQuantityRedisKeyFormat, characterID)
}

func (r *ActionSettlementRepo) backpackObtainedRedisKey(characterID int64) string {
	return fmt.Sprintf(r.backpackObtainedRedisKeyFormat, characterID)
}

func (r *ActionSettlementRepo) backpackConsumedRedisKey(characterID int64) string {
	return fmt.Sprintf(r.backpackConsumedRedisKeyFormat, characterID)
}

func (r *ActionSettlementRepo) backpackOperationRedisKey(characterID int64) string {
	return fmt.Sprintf(r.backpackOperationRedisKeyFormat, characterID)
}

func (r *ActionSettlementRepo) abilityExpRedisKey(characterID int64) string {
	return fmt.Sprintf(r.abilityExpRedisKeyFormat, characterID)
}

func (r *ActionSettlementRepo) abilityLevelRedisKey(characterID int64) string {
	return fmt.Sprintf(r.abilityLevelRedisKeyFormat, characterID)
}

func (r *ActionSettlementRepo) abilityNextLevelExpRedisKey(characterID int64) string {
	return fmt.Sprintf(r.abilityNextLevelExpRedisKeyFormat, characterID)
}

func (r *ActionSettlementRepo) abilityOperationCountRedisKey(characterID int64) string {
	return fmt.Sprintf(r.abilityOperationCountRedisKeyFormat, characterID)
}
