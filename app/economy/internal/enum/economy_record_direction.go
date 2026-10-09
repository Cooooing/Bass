package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/economy/v1/enum"
)

// EconomyRecordDirection 表示积分流水方向
type EconomyRecordDirection string

const (
	EconomyRecordDirectionIncome  EconomyRecordDirection = "income"  // EconomyRecordDirectionIncome 表示收入
	EconomyRecordDirectionExpense EconomyRecordDirection = "expense" // EconomyRecordDirectionExpense 表示支出
)

// EconomyRecordDirectionMap 维护流水方向映射
var EconomyRecordDirectionMap = commonenum.NewMapping[EconomyRecordDirection, v1.EconomyRecordDirection](map[EconomyRecordDirection]commonenum.Entry[EconomyRecordDirection, v1.EconomyRecordDirection]{
	EconomyRecordDirectionIncome:  {Proto: v1.EconomyRecordDirection_ECONOMY_RECORD_DIRECTION_INCOME},
	EconomyRecordDirectionExpense: {Proto: v1.EconomyRecordDirection_ECONOMY_RECORD_DIRECTION_EXPENSE},
})

func (e EconomyRecordDirection) String() string {
	return string(e)
}
