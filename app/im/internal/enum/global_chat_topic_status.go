package enum

import (
	commonenum "common/pkg/enum"
	imv1enum "common/proto/gen/im/v1/enum"
)

// GlobalChatTopicStatus 表示全局聊天室话题设置的状态。
type GlobalChatTopicStatus string

const (
	GlobalChatTopicStatusPending   GlobalChatTopicStatus = "pending"   // 等待事务确认。
	GlobalChatTopicStatusPublished GlobalChatTopicStatus = "published" // 已发布并公开。
	GlobalChatTopicStatusCanceled  GlobalChatTopicStatus = "canceled"  // 已取消且不公开。
)

// GlobalChatTopicStatusMap 维护话题设置状态的内部与契约映射。
var GlobalChatTopicStatusMap = commonenum.NewMapping[GlobalChatTopicStatus, imv1enum.GlobalChatTopicStatus](map[GlobalChatTopicStatus]commonenum.Entry[GlobalChatTopicStatus, imv1enum.GlobalChatTopicStatus]{
	GlobalChatTopicStatusPending:   {Proto: imv1enum.GlobalChatTopicStatus_GLOBAL_CHAT_TOPIC_STATUS_PENDING},
	GlobalChatTopicStatusPublished: {Proto: imv1enum.GlobalChatTopicStatus_GLOBAL_CHAT_TOPIC_STATUS_PUBLISHED},
	GlobalChatTopicStatusCanceled:  {Proto: imv1enum.GlobalChatTopicStatus_GLOBAL_CHAT_TOPIC_STATUS_CANCELED},
})

func (s GlobalChatTopicStatus) String() string { return string(s) }
