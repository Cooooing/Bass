package usecase

import (
	"common/proto/gen/common"
	"context"
	"im/internal/biz/model"
	"im/internal/biz/repo"
)

type ChatSessionUsecase struct {
	chatSessionRepo repo.ChatSessionRepo
	chatMessageRepo repo.ChatMessageRepo
}

func NewChatSessionUsecase(
	chatSessionRepo repo.ChatSessionRepo,
	chatMessageRepo repo.ChatMessageRepo,
) (*ChatSessionUsecase, error) {
	return &ChatSessionUsecase{
		chatSessionRepo: chatSessionRepo,
		chatMessageRepo: chatMessageRepo,
	}, nil
}

// MarkMuted 设置免打扰状态。
func (u *ChatSessionUsecase) MarkMuted(ctx context.Context, ids []int64, disturb bool, userId int64) error {
	for _, id := range ids {
		_, err := u.chatSessionRepo.UpdateMuted(ctx, id, disturb, userId)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarkPinned 设置置顶状态。
func (u *ChatSessionUsecase) MarkPinned(ctx context.Context, ids []int64, top bool, userId int64) error {
	for _, id := range ids {
		_, err := u.chatSessionRepo.UpdatePinned(ctx, id, top, userId)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarkRead 标记已读，将 last_read_message_id 更新为最新消息 ID。
func (u *ChatSessionUsecase) MarkRead(ctx context.Context, ids []int64, userId int64) error {
	for _, id := range ids {
		session, err := u.chatSessionRepo.Get(ctx, &repo.ChatSessionGetReq{
			IDs: []int64{id},
		})
		if err != nil {
			return err
		}
		// 获取该会话最新的消息作为已读位置
		latestMsg, err := u.chatMessageRepo.Get(ctx, &repo.ChatMessageGetReq{
			SessionID: &id,
		})
		if err != nil {
			continue // 会话无消息则跳过
		}
		// 计算新增已读数：latestMsg.ID - lastReadMsgID
		var readDelta int32
		if session.LastReadMessageID != nil {
			readDelta = int32(latestMsg.ID - *session.LastReadMessageID)
		} else {
			readDelta = int32(latestMsg.ID)
		}
		if readDelta <= 0 {
			continue
		}
		_, err = u.chatSessionRepo.UpdateLastReadMessage(ctx, id, latestMsg.ID, readDelta, userId)
		if err != nil {
			return err
		}
	}
	return nil
}

// Page 分页查询会话列表。
func (u *ChatSessionUsecase) Page(ctx context.Context, page *common.PageRequest, queryIds []int64, userId int64) ([]*model.ChatSession, *common.PageReply, error) {
	getReq := &repo.ChatSessionGetReq{
		IDs:       queryIds,
		CreatedBy: &userId,
	}
	return u.chatSessionRepo.Page(ctx, page, getReq)
}
