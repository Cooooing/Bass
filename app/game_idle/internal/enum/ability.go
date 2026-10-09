package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_idle/v1/enum"
)

// Ability 表示角色可成长能力。
type Ability string

const (
	AbilityWoodcutting Ability = "woodcutting" // 伐木
	AbilityForaging    Ability = "foraging"    // 采集
	AbilityMining      Ability = "mining"      // 采矿
	AbilityFishing     Ability = "fishing"     // 钓鱼
	AbilityCrafting    Ability = "crafting"    // 制作
	AbilitySewing      Ability = "sewing"      // 缝纫
	AbilitySmithing    Ability = "smithing"    // 锻造
	AbilityCooking     Ability = "cooking"     // 烹饪
	AbilityEnhancing   Ability = "enhancing"   // 强化
	AbilityAlchemy     Ability = "alchemy"     // 炼金
)

var AbilityMap = commonenum.NewMapping[Ability, v1.Ability](map[Ability]commonenum.Entry[Ability, v1.Ability]{
	AbilityWoodcutting: {Proto: v1.Ability_ABILITY_WOODCUTTING}, AbilityForaging: {Proto: v1.Ability_ABILITY_FORAGING}, AbilityMining: {Proto: v1.Ability_ABILITY_MINING}, AbilityFishing: {Proto: v1.Ability_ABILITY_FISHING}, AbilityCrafting: {Proto: v1.Ability_ABILITY_CRAFTING}, AbilitySewing: {Proto: v1.Ability_ABILITY_SEWING}, AbilitySmithing: {Proto: v1.Ability_ABILITY_SMITHING}, AbilityCooking: {Proto: v1.Ability_ABILITY_COOKING}, AbilityEnhancing: {Proto: v1.Ability_ABILITY_ENHANCING}, AbilityAlchemy: {Proto: v1.Ability_ABILITY_ALCHEMY},
})

func (e Ability) String() string {
	return string(e)
}
