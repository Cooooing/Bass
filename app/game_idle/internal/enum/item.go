package enum

// ItemID 表示有特殊业务语义的物品编码。
type ItemID string

const (
	ItemIDEmpty ItemID = "empty" // 空产出
)

func (e ItemID) String() string {
	return string(e)
}

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

func (e ItemType) String() string {
	return string(e)
}

func ItemTypeValues() []string {
	return []string{
		ItemTypeCurrency.String(),
		ItemTypeLoot.String(),
		ItemTypeConsumable.String(),
		ItemTypeSkillBook.String(),
		ItemTypeEquipment.String(),
		ItemTypeResource.String(),
	}
}
