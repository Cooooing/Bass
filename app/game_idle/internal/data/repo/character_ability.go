package repo

import (
	"context"
	"game_idle/internal/biz/model"
	bizrepo "game_idle/internal/biz/repo"
	"game_idle/internal/config"
	"game_idle/internal/data/gen"
	characterabilityent "game_idle/internal/data/gen/characterability"
	"game_idle/internal/enum"
)

var _ bizrepo.CharacterAbilityRepo = (*CharacterAbilityRepo)(nil)

type CharacterAbilityRepo struct {
	db               *gen.Client
	cache            *CharacterStateCache
	persistThreshold int64
}

func NewCharacterAbilityRepo(
	conf *config.Bootstrap,
	db *gen.Client,
	cache *CharacterStateCache,
) bizrepo.CharacterAbilityRepo {
	persistThreshold := int64(100)
	if conf.GetGameIdle().GetCharacterAbility().GetPersistThreshold() > 0 {
		persistThreshold = int64(conf.GetGameIdle().GetCharacterAbility().GetPersistThreshold())
	}
	return &CharacterAbilityRepo{
		db:               db,
		cache:            cache,
		persistThreshold: persistThreshold,
	}
}

func (r *CharacterAbilityRepo) Map(
	ctx context.Context,
	req *bizrepo.CharacterAbilityMapReq,
) (map[enum.Ability]*model.CharacterAbility, error) {
	character := r.cache.character(req.CharacterID)
	character.mutex.RLock()
	abilitiesCache := character.abilities
	character.mutex.RUnlock()
	if abilitiesCache == nil {
		rows, err := r.db.CharacterAbility.Query().
			Where(characterabilityent.CharacterIDEQ(req.CharacterID)).
			All(ctx)
		if err != nil {
			return nil, err
		}
		abilities := make(map[enum.Ability]*model.CharacterAbility, len(rows))
		for _, row := range rows {
			abilityID := enum.Ability(row.AbilityID)
			abilities[abilityID] = &model.CharacterAbility{
				CharacterID: row.CharacterID,
				AbilityID:   abilityID,
				Level:       row.Level,
				Exp:         row.Exp,
			}
		}
		character.mutex.Lock()
		if character.abilities == nil {
			character.abilities = &cachedAbilities{
				items:          abilities,
				operationCount: 0,
			}
		}
		character.mutex.Unlock()
	}
	character.mutex.RLock()
	defer character.mutex.RUnlock()
	source := character.abilities.items
	abilities := make(map[enum.Ability]*model.CharacterAbility)
	if len(req.AbilityIDs) == 0 {
		for abilityID, ability := range source {
			copyAbility := *ability
			abilities[abilityID] = &copyAbility
		}
		return abilities, nil
	}
	for _, abilityID := range req.AbilityIDs {
		ability := source[abilityID]
		if ability != nil {
			copyAbility := *ability
			abilities[abilityID] = &copyAbility
		}
	}
	return abilities, nil
}

func (r *CharacterAbilityRepo) Persist(ctx context.Context, characterID int64) error {
	character := r.cache.character(characterID)
	character.mutex.RLock()
	abilitiesCache := character.abilities
	character.mutex.RUnlock()
	if abilitiesCache == nil {
		rows, err := r.db.CharacterAbility.Query().
			Where(characterabilityent.CharacterIDEQ(characterID)).
			All(ctx)
		if err != nil {
			return err
		}
		abilities := make(map[enum.Ability]*model.CharacterAbility, len(rows))
		for _, row := range rows {
			abilityID := enum.Ability(row.AbilityID)
			abilities[abilityID] = &model.CharacterAbility{
				CharacterID: row.CharacterID,
				AbilityID:   abilityID,
				Level:       row.Level,
				Exp:         row.Exp,
			}
		}
		character.mutex.Lock()
		if character.abilities == nil {
			character.abilities = &cachedAbilities{
				items:          abilities,
				operationCount: 0,
			}
		}
		character.mutex.Unlock()
	}
	character.mutex.RLock()
	source := character.abilities.items
	abilities := make([]*model.CharacterAbility, 0, len(source))
	for _, ability := range source {
		copyAbility := *ability
		abilities = append(abilities, &copyAbility)
	}
	character.mutex.RUnlock()

	creates := make([]*gen.CharacterAbilityCreate, 0, len(abilities))
	for _, ability := range abilities {
		creates = append(creates, r.db.CharacterAbility.Create().
			SetCharacterID(characterID).
			SetAbilityID(characterabilityent.AbilityID(ability.AbilityID)).
			SetLevel(ability.Level).
			SetExp(ability.Exp))
	}
	if len(creates) > 0 {
		if err := r.db.CharacterAbility.CreateBulk(creates...).
			OnConflictColumns(characterabilityent.FieldCharacterID, characterabilityent.FieldAbilityID).
			UpdateLevel().
			UpdateExp().
			UpdateUpdatedAt().
			Exec(ctx); err != nil {
			return err
		}
	}
	character.mutex.Lock()
	character.abilities.operationCount = 0
	character.mutex.Unlock()
	return nil
}
