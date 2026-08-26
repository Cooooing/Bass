package enum

// ActionKind 表示区域活动关联的行动能力类型。
type ActionKind string

const (
	ActionKindWoodcutting ActionKind = "woodcutting" // 伐木
	ActionKindForaging    ActionKind = "foraging"    // 采集
	ActionKindMining      ActionKind = "mining"      // 采矿
	ActionKindFishing     ActionKind = "fishing"     // 钓鱼
	ActionKindCrafting    ActionKind = "crafting"    // 制作
	ActionKindSewing      ActionKind = "sewing"      // 缝纫
	ActionKindSmithing    ActionKind = "smithing"    // 锻造
	ActionKindCooking     ActionKind = "cooking"     // 烹饪
	ActionKindEnhancing   ActionKind = "enhancing"   // 强化
	ActionKindAlchemy     ActionKind = "alchemy"     // 炼金
)

func (e ActionKind) String() string {
	return string(e)
}

func ActionKindValues() []string {
	return []string{
		ActionKindWoodcutting.String(),
		ActionKindForaging.String(),
		ActionKindMining.String(),
		ActionKindFishing.String(),
		ActionKindCrafting.String(),
		ActionKindSewing.String(),
		ActionKindSmithing.String(),
		ActionKindCooking.String(),
		ActionKindEnhancing.String(),
		ActionKindAlchemy.String(),
	}
}
