package enum

import (
	commonenum "common/pkg/enum"
	imv1enum "common/proto/gen/im/v1/enum"
)

type GlobalChatBarrageStatus string

const (
	GlobalChatBarrageStatusPending   GlobalChatBarrageStatus = "pending"   // 等待事务确认。
	GlobalChatBarrageStatusPublished GlobalChatBarrageStatus = "published" // 已发布。
	GlobalChatBarrageStatusCanceled  GlobalChatBarrageStatus = "canceled"  // 已取消。
)

var GlobalChatBarrageStatusMap = commonenum.NewMapping[GlobalChatBarrageStatus, imv1enum.GlobalChatBarrageStatus](map[GlobalChatBarrageStatus]commonenum.Entry[GlobalChatBarrageStatus, imv1enum.GlobalChatBarrageStatus]{
	GlobalChatBarrageStatusPending:   {Proto: imv1enum.GlobalChatBarrageStatus_GLOBAL_CHAT_BARRAGE_STATUS_PENDING},
	GlobalChatBarrageStatusPublished: {Proto: imv1enum.GlobalChatBarrageStatus_GLOBAL_CHAT_BARRAGE_STATUS_PUBLISHED},
	GlobalChatBarrageStatusCanceled:  {Proto: imv1enum.GlobalChatBarrageStatus_GLOBAL_CHAT_BARRAGE_STATUS_CANCELED},
})

func (s GlobalChatBarrageStatus) String() string { return string(s) }
