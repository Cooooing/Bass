package repo

import (
	"context"
	"time"

	"im/internal/biz/model"
	bizrepo "im/internal/biz/repo"
	"im/internal/data/gen"
	"im/internal/data/gen/globalchatmessage"
	"im/internal/enum"
)

var _ bizrepo.StatisticsRepo = (*StatisticsRepo)(nil)

// StatisticsRepo 以分组查询批量统计即时消息事实表。
type StatisticsRepo struct{ db *gen.Client }

// NewStatisticsRepo 创建即时消息统计仓储。
func NewStatisticsRepo(db *gen.Client) bizrepo.StatisticsRepo { return &StatisticsRepo{db: db} }

// BatchCountGlobalChatMessages 按发送账号聚合未撤回普通文本消息数量。
func (r *StatisticsRepo) BatchCountGlobalChatMessages(ctx context.Context, userIDs []int64, start, end time.Time) ([]*model.GlobalChatMessageCount, error) {
	type row struct {
		UserID int64 `json:"sender_id"`
		Count  int   `json:"count"`
	}
	var rows []row
	err := r.db.GlobalChatMessage.Query().Where(globalchatmessage.SenderIDIn(userIDs...), globalchatmessage.MessageTypeEQ(globalchatmessage.MessageType(enum.MessageTypeText)), globalchatmessage.IsRevokedEQ(false), globalchatmessage.CreatedAtGTE(start), globalchatmessage.CreatedAtLT(end)).GroupBy(globalchatmessage.FieldSenderID).Aggregate(gen.Count()).Scan(ctx, &rows)
	if err != nil {
		return nil, err
	}
	result := make([]*model.GlobalChatMessageCount, 0, len(rows))
	for _, item := range rows {
		result = append(result, &model.GlobalChatMessageCount{UserID: item.UserID, Count: int32(item.Count)})
	}
	return result, nil
}
