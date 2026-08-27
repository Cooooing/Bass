package repo

import (
	commonclient "common/pkg/client"
	"context"
	"fmt"
	"game_idle/internal/biz/model"
	bizrepo "game_idle/internal/biz/repo"
	"game_idle/internal/config"
	"game_idle/internal/data/gen"
	characterabilityent "game_idle/internal/data/gen/characterability"
	"game_idle/internal/enum"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var _ bizrepo.CharacterAbilityRepo = (*CharacterAbilityRepo)(nil)

type CharacterAbilityRepo struct {
	db                           *gen.Client
	redisClient                  *commonclient.RedisClient
	persistThreshold             int64
	expRedisKeyFormat            string
	levelRedisKeyFormat          string
	nextLevelExpRedisKeyFormat   string
	loadedRedisKeyFormat         string
	operationCountRedisKeyFormat string
}

func NewCharacterAbilityRepo(
	conf *config.Bootstrap,
	db *gen.Client,
	redisClient *commonclient.RedisClient,
) bizrepo.CharacterAbilityRepo {
	persistThreshold := int64(100)
	if conf.GetGameIdle().GetCharacterAbility().GetPersistThreshold() > 0 {
		persistThreshold = int64(conf.GetGameIdle().GetCharacterAbility().GetPersistThreshold())
	}
	return &CharacterAbilityRepo{
		db:                           db,
		redisClient:                  redisClient,
		persistThreshold:             persistThreshold,
		expRedisKeyFormat:            "game_idle:ability:{character_id:%d}:exp",
		levelRedisKeyFormat:          "game_idle:ability:{character_id:%d}:level",
		nextLevelExpRedisKeyFormat:   "game_idle:ability:{character_id:%d}:next_level_exp",
		loadedRedisKeyFormat:         "game_idle:ability:{character_id:%d}:loaded",
		operationCountRedisKeyFormat: "game_idle:ability:{character_id:%d}:operation_count",
	}
}

func (r *CharacterAbilityRepo) load(ctx context.Context, characterID int64) error {
	rows, err := r.db.CharacterAbility.Query().
		Where(characterabilityent.CharacterIDEQ(characterID)).
		All(ctx)
	if err != nil {
		return err
	}
	_, err = r.redisClient.Client.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.Del(
			ctx,
			r.expRedisKey(characterID),
			r.levelRedisKey(characterID),
			r.nextLevelExpRedisKey(characterID),
		)
		for _, row := range rows {
			pipe.HSet(ctx, r.expRedisKey(characterID), row.AbilityID.String(), row.Exp)
			pipe.HSet(ctx, r.levelRedisKey(characterID), row.AbilityID.String(), row.Level)
			pipe.HSet(ctx, r.nextLevelExpRedisKey(characterID), row.AbilityID.String(), r.nextLevelExp(row.Level))
		}
		pipe.Set(ctx, r.loadedRedisKey(characterID), "1", 0)
		pipe.SetNX(ctx, r.operationCountRedisKey(characterID), 0, 0)
		return nil
	})
	return err
}

func (r *CharacterAbilityRepo) Map(
	ctx context.Context,
	req *bizrepo.CharacterAbilityMapReq,
) (map[enum.Ability]*model.CharacterAbility, error) {
	if err := r.ensureLoaded(ctx, req.CharacterID); err != nil {
		return nil, err
	}
	expValues, err := r.redisClient.Client.HGetAll(ctx, r.expRedisKey(req.CharacterID)).Result()
	if err != nil {
		return nil, err
	}
	levelValues, err := r.redisClient.Client.HGetAll(ctx, r.levelRedisKey(req.CharacterID)).Result()
	if err != nil {
		return nil, err
	}
	nextLevelExpValues, err := r.redisClient.Client.HGetAll(ctx, r.nextLevelExpRedisKey(req.CharacterID)).Result()
	if err != nil {
		return nil, err
	}
	abilityIDs := req.AbilityIDs
	if len(abilityIDs) == 0 {
		abilityIDs = make([]enum.Ability, 0, len(expValues))
		for abilityID := range expValues {
			abilityIDs = append(abilityIDs, enum.Ability(abilityID))
		}
	}
	abilities := make(map[enum.Ability]*model.CharacterAbility, len(abilityIDs))
	for _, abilityID := range abilityIDs {
		expText, ok := expValues[abilityID.String()]
		if !ok {
			continue
		}
		exp, err := strconv.ParseInt(expText, 10, 64)
		if err != nil {
			return nil, err
		}
		level := int32(1)
		if levelText, ok := levelValues[abilityID.String()]; ok {
			value, err := strconv.ParseInt(levelText, 10, 32)
			if err != nil {
				return nil, err
			}
			level = int32(value)
		}
		nextLevelExp := r.nextLevelExp(level)
		if nextLevelExpText, ok := nextLevelExpValues[abilityID.String()]; ok {
			value, err := strconv.ParseInt(nextLevelExpText, 10, 64)
			if err != nil {
				return nil, err
			}
			nextLevelExp = value
		}
		abilities[abilityID] = &model.CharacterAbility{
			CharacterID:  req.CharacterID,
			AbilityID:    abilityID,
			Level:        level,
			Exp:          exp,
			NextLevelExp: nextLevelExp,
		}
	}
	return abilities, nil
}

func (r *CharacterAbilityRepo) Persist(ctx context.Context, characterID int64) error {
	if err := r.ensureLoaded(ctx, characterID); err != nil {
		return err
	}
	expValues, err := r.redisClient.Client.HGetAll(ctx, r.expRedisKey(characterID)).Result()
	if err != nil {
		return err
	}
	levelValues, err := r.redisClient.Client.HGetAll(ctx, r.levelRedisKey(characterID)).Result()
	if err != nil {
		return err
	}
	creates := make([]*gen.CharacterAbilityCreate, 0, len(expValues))
	for abilityID, expText := range expValues {
		exp, err := strconv.ParseInt(expText, 10, 64)
		if err != nil {
			return err
		}
		level := int32(1)
		if levelText := levelValues[abilityID]; levelText != "" {
			value, err := strconv.ParseInt(levelText, 10, 32)
			if err != nil {
				return err
			}
			level = int32(value)
		}
		creates = append(creates, r.db.CharacterAbility.Create().
			SetCharacterID(characterID).
			SetAbilityID(characterabilityent.AbilityID(abilityID)).
			SetLevel(level).
			SetExp(exp))
	}
	if len(creates) > 0 {
		err = r.db.CharacterAbility.CreateBulk(creates...).
			OnConflictColumns(characterabilityent.FieldCharacterID, characterabilityent.FieldAbilityID).
			UpdateLevel().
			UpdateExp().
			UpdateUpdatedAt().
			Exec(ctx)
		if err != nil {
			return err
		}
	}
	return r.redisClient.Client.Set(ctx, r.operationCountRedisKey(characterID), 0, 0).Err()
}

func (r *CharacterAbilityRepo) ensureLoaded(ctx context.Context, characterID int64) error {
	loaded, err := r.redisClient.Client.Exists(ctx, r.loadedRedisKey(characterID)).Result()
	if err != nil {
		return err
	}
	if loaded > 0 {
		return nil
	}
	return r.load(ctx, characterID)
}

func (r *CharacterAbilityRepo) expRedisKey(characterID int64) string {
	return fmt.Sprintf(r.expRedisKeyFormat, characterID)
}

func (r *CharacterAbilityRepo) levelRedisKey(characterID int64) string {
	return fmt.Sprintf(r.levelRedisKeyFormat, characterID)
}

func (r *CharacterAbilityRepo) nextLevelExpRedisKey(characterID int64) string {
	return fmt.Sprintf(r.nextLevelExpRedisKeyFormat, characterID)
}

func (r *CharacterAbilityRepo) loadedRedisKey(characterID int64) string {
	return fmt.Sprintf(r.loadedRedisKeyFormat, characterID)
}

func (r *CharacterAbilityRepo) operationCountRedisKey(characterID int64) string {
	return fmt.Sprintf(r.operationCountRedisKeyFormat, characterID)
}

func (r *CharacterAbilityRepo) nextLevelExp(level int32) int64 {
	return int64(level) * int64(level) * 100
}
