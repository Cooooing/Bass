package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_idle/v1/enum"
)

// ItemType 表示物品类型，货币、资源、装备等资产统一作为物品。
type ItemType string

const (
	ItemTypeCurrency   ItemType = "currency"   // 货币
	ItemTypeLoot       ItemType = "loot"       // 战利品
	ItemTypeConsumable ItemType = "consumable" // 消耗品
	ItemTypeSkillBook  ItemType = "skill_book" // 技能书
	ItemTypeEquipment  ItemType = "equipment"  // 装备
	ItemTypeResource   ItemType = "resource"   // 资源
)

var ItemTypeMap = commonenum.NewMapping[ItemType, v1.ItemType](map[ItemType]commonenum.Entry[ItemType, v1.ItemType]{
	ItemTypeCurrency: {Proto: v1.ItemType_ITEM_TYPE_CURRENCY}, ItemTypeLoot: {Proto: v1.ItemType_ITEM_TYPE_LOOT}, ItemTypeConsumable: {Proto: v1.ItemType_ITEM_TYPE_CONSUMABLE}, ItemTypeSkillBook: {Proto: v1.ItemType_ITEM_TYPE_SKILL_BOOK}, ItemTypeEquipment: {Proto: v1.ItemType_ITEM_TYPE_EQUIPMENT}, ItemTypeResource: {Proto: v1.ItemType_ITEM_TYPE_RESOURCE},
})

func (e ItemType) String() string {
	return string(e)
}
