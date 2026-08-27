package repo

import (
	commonclient "common/pkg/client"
	"context"
	"fmt"
	"game_idle/internal/biz/model"
	bizrepo "game_idle/internal/biz/repo"
	"game_idle/internal/config"
	"game_idle/internal/data/gen"
	characteritement "game_idle/internal/data/gen/characteritem"
	"strconv"
	"strings"

	"github.com/redis/go-redis/v9"
)

var _ bizrepo.BackpackRepo = (*BackpackRepo)(nil)

type BackpackRepo struct {
	db                           *gen.Client
	redisClient                  *commonclient.RedisClient
	persistThreshold             int64
	quantityRedisKeyFormat       string
	totalObtainedRedisKeyFormat  string
	totalConsumedRedisKeyFormat  string
	loadedRedisKeyFormat         string
	operationCountRedisKeyFormat string
	checkScript                  *redis.Script
	changeScript                 *redis.Script
}

func NewBackpackRepo(
	conf *config.Bootstrap,
	db *gen.Client,
	redisClient *commonclient.RedisClient,
) (bizrepo.BackpackRepo, error) {
	persistThreshold := int64(100)
	if conf.GetGameIdle().GetBackpack().GetPersistThreshold() > 0 {
		persistThreshold = int64(conf.GetGameIdle().GetBackpack().GetPersistThreshold())
	}
	return &BackpackRepo{
		db:                           db,
		redisClient:                  redisClient,
		persistThreshold:             persistThreshold,
		quantityRedisKeyFormat:       "game_idle:character_backpack:%d:quantity",
		totalObtainedRedisKeyFormat:  "game_idle:character_backpack:%d:total_obtained",
		totalConsumedRedisKeyFormat:  "game_idle:character_backpack:%d:total_consumed",
		loadedRedisKeyFormat:         "game_idle:character_backpack:%d:loaded",
		operationCountRedisKeyFormat: "game_idle:character_backpack:%d:operation_count",
		checkScript: redis.NewScript(`
for index = 1, #ARGV, 2 do
  local item_id = ARGV[index]
  local quantity = tonumber(ARGV[index + 1])
  local current = tonumber(redis.call('HGET', KEYS[1], item_id) or '0')
  if current < quantity then
    return redis.error_reply('backpack_insufficient')
  end
end
return 1
`),
		changeScript: redis.NewScript(`
local deltas = {}
for index = 1, #ARGV, 2 do
  local item_id = ARGV[index]
  local quantity = tonumber(ARGV[index + 1])
  deltas[item_id] = (deltas[item_id] or 0) + quantity
end
for item_id, quantity in pairs(deltas) do
  local current = tonumber(redis.call('HGET', KEYS[1], item_id) or '0')
  if current + quantity < 0 then
    return redis.error_reply('backpack_insufficient')
  end
end
for item_id, quantity in pairs(deltas) do
  if quantity ~= 0 then
    redis.call('HINCRBY', KEYS[1], item_id, quantity)
  end
end
for index = 1, #ARGV, 2 do
  local item_id = ARGV[index]
  local quantity = tonumber(ARGV[index + 1])
  if quantity ~= 0 then
    if quantity > 0 then
      redis.call('HINCRBY', KEYS[2], item_id, quantity)
    else
      redis.call('HINCRBY', KEYS[3], item_id, -quantity)
    end
  end
end
local result = {redis.call('INCRBY', KEYS[4], #ARGV / 2)}
local returned = {}
for index = 1, #ARGV, 2 do
  local item_id = ARGV[index]
  if returned[item_id] == nil then
    returned[item_id] = true
    table.insert(result, item_id)
    table.insert(result, redis.call('HGET', KEYS[1], item_id) or '0')
  end
end
return result
`),
	}, nil
}

func (r *BackpackRepo) LoadItems(ctx context.Context, characterID int64) error {
	rows, err := r.db.CharacterItem.Query().
		Where(characteritement.CharacterIDEQ(characterID)).
		All(ctx)
	if err != nil {
		return err
	}
	_, err = r.redisClient.Client.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.Del(
			ctx,
			r.quantityRedisKey(characterID),
			r.totalObtainedRedisKey(characterID),
			r.totalConsumedRedisKey(characterID),
		)
		for _, row := range rows {
			pipe.HSet(ctx, r.quantityRedisKey(characterID), row.ItemID, row.Quantity)
			pipe.HSet(ctx, r.totalObtainedRedisKey(characterID), row.ItemID, row.TotalObtained)
			pipe.HSet(ctx, r.totalConsumedRedisKey(characterID), row.ItemID, row.TotalConsumed)
		}
		pipe.Set(ctx, r.loadedRedisKey(characterID), "1", 0)
		pipe.SetNX(ctx, r.operationCountRedisKey(characterID), 0, 0)
		return nil
	})
	return err
}

func (r *BackpackRepo) MapItems(ctx context.Context, req *bizrepo.BackpackMapReq) (map[string]*model.CharacterItem, error) {
	if err := r.EnsureLoaded(ctx, req.CharacterID); err != nil {
		return nil, err
	}
	quantityValues, err := r.redisClient.Client.HGetAll(ctx, r.quantityRedisKey(req.CharacterID)).Result()
	if err != nil {
		return nil, err
	}
	obtainedValues, err := r.redisClient.Client.HGetAll(ctx, r.totalObtainedRedisKey(req.CharacterID)).Result()
	if err != nil {
		return nil, err
	}
	consumedValues, err := r.redisClient.Client.HGetAll(ctx, r.totalConsumedRedisKey(req.CharacterID)).Result()
	if err != nil {
		return nil, err
	}
	itemIDs := req.ItemIDs
	if len(itemIDs) == 0 {
		itemIDs = make([]string, 0, len(quantityValues))
		for itemID := range quantityValues {
			itemIDs = append(itemIDs, itemID)
		}
	}
	items := make(map[string]*model.CharacterItem, len(itemIDs))
	for _, itemID := range itemIDs {
		quantityText, ok := quantityValues[itemID]
		if !ok {
			continue
		}
		quantity, err := strconv.ParseInt(quantityText, 10, 64)
		if err != nil {
			return nil, err
		}
		totalObtained := int64(0)
		if value, ok := obtainedValues[itemID]; ok {
			totalObtained, err = strconv.ParseInt(value, 10, 64)
			if err != nil {
				return nil, err
			}
		}
		totalConsumed := int64(0)
		if value, ok := consumedValues[itemID]; ok {
			totalConsumed, err = strconv.ParseInt(value, 10, 64)
			if err != nil {
				return nil, err
			}
		}
		items[itemID] = &model.CharacterItem{
			CharacterID:   req.CharacterID,
			ItemID:        itemID,
			Quantity:      quantity,
			TotalObtained: totalObtained,
			TotalConsumed: totalConsumed,
		}
	}
	return items, nil
}

func (r *BackpackRepo) PersistItems(ctx context.Context, characterID int64) error {
	if err := r.EnsureLoaded(ctx, characterID); err != nil {
		return err
	}
	quantityValues, err := r.redisClient.Client.HGetAll(ctx, r.quantityRedisKey(characterID)).Result()
	if err != nil {
		return err
	}
	obtainedValues, err := r.redisClient.Client.HGetAll(ctx, r.totalObtainedRedisKey(characterID)).Result()
	if err != nil {
		return err
	}
	consumedValues, err := r.redisClient.Client.HGetAll(ctx, r.totalConsumedRedisKey(characterID)).Result()
	if err != nil {
		return err
	}
	itemIDs := make(map[string]struct{})
	for itemID := range quantityValues {
		itemIDs[itemID] = struct{}{}
	}
	for itemID := range obtainedValues {
		itemIDs[itemID] = struct{}{}
	}
	for itemID := range consumedValues {
		itemIDs[itemID] = struct{}{}
	}
	creates := make([]*gen.CharacterItemCreate, 0, len(itemIDs))
	for itemID := range itemIDs {
		quantity := int64(0)
		if value, ok := quantityValues[itemID]; ok {
			quantity, err = strconv.ParseInt(value, 10, 64)
			if err != nil {
				return err
			}
		}
		totalObtained := int64(0)
		if value, ok := obtainedValues[itemID]; ok {
			totalObtained, err = strconv.ParseInt(value, 10, 64)
			if err != nil {
				return err
			}
		}
		totalConsumed := int64(0)
		if value, ok := consumedValues[itemID]; ok {
			totalConsumed, err = strconv.ParseInt(value, 10, 64)
			if err != nil {
				return err
			}
		}
		creates = append(creates, r.db.CharacterItem.Create().
			SetCharacterID(characterID).
			SetItemID(itemID).
			SetQuantity(quantity).
			SetTotalObtained(totalObtained).
			SetTotalConsumed(totalConsumed))
	}
	if len(creates) > 0 {
		err = r.db.CharacterItem.CreateBulk(creates...).
			OnConflictColumns(characteritement.FieldCharacterID, characteritement.FieldItemID).
			UpdateQuantity().
			UpdateTotalObtained().
			UpdateTotalConsumed().
			UpdateUpdatedAt().
			Exec(ctx)
		if err != nil {
			return err
		}
	}
	return r.redisClient.Client.Set(ctx, r.operationCountRedisKey(characterID), 0, 0).Err()
}

func (r *BackpackRepo) CheckItems(ctx context.Context, req *bizrepo.BackpackCheckReq) error {
	args := make([]any, 0, len(req.Items)*2)
	for itemID, quantity := range req.Items {
		args = append(args, itemID, quantity)
	}
	if err := r.EnsureLoaded(ctx, req.CharacterID); err != nil {
		return err
	}
	if err := r.checkScript.Run(
		ctx,
		r.redisClient.Client,
		[]string{r.quantityRedisKey(req.CharacterID)},
		args...,
	).Err(); err != nil {
		if strings.Contains(err.Error(), "backpack_insufficient") {
			return model.ErrBackpackInsufficient
		}
		return err
	}
	return nil
}

func (r *BackpackRepo) ChangeItems(ctx context.Context, req *bizrepo.BackpackChangeReq) (map[string]int64, error) {
	args := make([]any, 0, len(req.Items)*2)
	for _, item := range req.Items {
		args = append(args, item.ItemID, item.Quantity)
	}
	if err := r.EnsureLoaded(ctx, req.CharacterID); err != nil {
		return nil, err
	}
	result, err := r.changeScript.Run(
		ctx,
		r.redisClient.Client,
		[]string{
			r.quantityRedisKey(req.CharacterID),
			r.totalObtainedRedisKey(req.CharacterID),
			r.totalConsumedRedisKey(req.CharacterID),
			r.operationCountRedisKey(req.CharacterID),
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
		return nil, fmt.Errorf("game idle backpack change result invalid")
	}
	count, err := strconv.ParseInt(fmt.Sprint(values[0]), 10, 64)
	if err != nil {
		return nil, err
	}
	quantityAfter := make(map[string]int64, len(values)/2)
	for index := 1; index+1 < len(values); index += 2 {
		itemID := fmt.Sprint(values[index])
		quantity, err := strconv.ParseInt(fmt.Sprint(values[index+1]), 10, 64)
		if err != nil {
			return nil, err
		}
		quantityAfter[itemID] = quantity
	}
	if r.persistThreshold > 0 && count >= r.persistThreshold {
		return quantityAfter, r.PersistItems(ctx, req.CharacterID)
	}
	return quantityAfter, nil
}

func (r *BackpackRepo) EnsureLoaded(ctx context.Context, characterID int64) error {
	loaded, err := r.redisClient.Client.Exists(ctx, r.loadedRedisKey(characterID)).Result()
	if err != nil {
		return err
	}
	if loaded > 0 {
		return nil
	}
	return r.LoadItems(ctx, characterID)
}

func (r *BackpackRepo) quantityRedisKey(characterID int64) string {
	return fmt.Sprintf(r.quantityRedisKeyFormat, characterID)
}

func (r *BackpackRepo) totalObtainedRedisKey(characterID int64) string {
	return fmt.Sprintf(r.totalObtainedRedisKeyFormat, characterID)
}

func (r *BackpackRepo) totalConsumedRedisKey(characterID int64) string {
	return fmt.Sprintf(r.totalConsumedRedisKeyFormat, characterID)
}

func (r *BackpackRepo) loadedRedisKey(characterID int64) string {
	return fmt.Sprintf(r.loadedRedisKeyFormat, characterID)
}

func (r *BackpackRepo) operationCountRedisKey(characterID int64) string {
	return fmt.Sprintf(r.operationCountRedisKeyFormat, characterID)
}
