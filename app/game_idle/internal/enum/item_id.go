package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_idle/v1/enum"
)

// ItemID 表示有特殊业务语义的物品编码。
type ItemID string

const (
	ItemIDEmpty ItemID = "empty" // 空产出
)

var ItemIDMap = commonenum.NewMapping[ItemID, v1.ItemID](map[ItemID]commonenum.Entry[ItemID, v1.ItemID]{
	ItemIDEmpty: {Proto: v1.ItemID_ITEM_ID_EMPTY},
})

func (e ItemID) String() string {
	return string(e)
}
