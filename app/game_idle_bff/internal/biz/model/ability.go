package model

type CharacterAbility struct {
	AbilityID    string `json:"ability_id"`
	Level        int32  `json:"level"`
	Exp          int64  `json:"exp"`
	NextLevelExp int64  `json:"next_level_exp"`
}
