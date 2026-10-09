package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_idle/v1/enum"
)

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

var ActionKindMap = commonenum.NewMapping[ActionKind, v1.ActionKind](map[ActionKind]commonenum.Entry[ActionKind, v1.ActionKind]{
	ActionKindWoodcutting: {Proto: v1.ActionKind_ACTION_KIND_WOODCUTTING}, ActionKindForaging: {Proto: v1.ActionKind_ACTION_KIND_FORAGING}, ActionKindMining: {Proto: v1.ActionKind_ACTION_KIND_MINING}, ActionKindFishing: {Proto: v1.ActionKind_ACTION_KIND_FISHING}, ActionKindCrafting: {Proto: v1.ActionKind_ACTION_KIND_CRAFTING}, ActionKindSewing: {Proto: v1.ActionKind_ACTION_KIND_SEWING}, ActionKindSmithing: {Proto: v1.ActionKind_ACTION_KIND_SMITHING}, ActionKindCooking: {Proto: v1.ActionKind_ACTION_KIND_COOKING}, ActionKindEnhancing: {Proto: v1.ActionKind_ACTION_KIND_ENHANCING}, ActionKindAlchemy: {Proto: v1.ActionKind_ACTION_KIND_ALCHEMY},
})

func (e ActionKind) String() string {
	return string(e)
}
