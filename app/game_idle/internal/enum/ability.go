package enum

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

func (e Ability) String() string {
	return string(e)
}

func AbilityValues() []string {
	return []string{
		AbilityWoodcutting.String(),
		AbilityForaging.String(),
		AbilityMining.String(),
		AbilityFishing.String(),
		AbilityCrafting.String(),
		AbilitySewing.String(),
		AbilitySmithing.String(),
		AbilityCooking.String(),
		AbilityEnhancing.String(),
		AbilityAlchemy.String(),
	}
}
