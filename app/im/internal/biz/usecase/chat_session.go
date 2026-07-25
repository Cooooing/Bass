package usecase

import (
	"context"
	"im/internal/biz/base"
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

type MarkMutedReq struct {
	IDs     []int64
	Disturb bool
	UserID  int64
}

func (u *ChatSessionUsecase) MarkMuted(ctx context.Context, req *MarkMutedReq) error {
	for _, id := range req.IDs {
		_, err := u.chatSessionRepo.UpdateMuted(ctx, &repo.ChatSessionUpdateMutedReq{
			ChatSessionID: id,
			Muted:         req.Disturb,
			UpdatedBy:     req.UserID,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

type MarkPinnedReq struct {
	IDs    []int64
	Top    bool
	UserID int64
}

func (u *ChatSessionUsecase) MarkPinned(ctx context.Context, req *MarkPinnedReq) error {
	for _, id := range req.IDs {
		_, err := u.chatSessionRepo.UpdatePinned(ctx, &repo.ChatSessionUpdatePinnedReq{
			ChatSessionID: id,
			Pinned:        req.Top,
			UpdatedBy:     req.UserID,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

type MarkReadReq struct {
	IDs    []int64
	UserID int64
}

func (u *ChatSessionUsecase) MarkRead(ctx context.Context, req *MarkReadReq) error {
	for _, id := range req.IDs {
		session, err := u.chatSessionRepo.Get(ctx, &repo.ChatSessionQuery{
			IDs: []int64{id},
		})
		if err != nil {
			return err
		}
		latestMsg, err := u.chatMessageRepo.Get(ctx, &repo.ChatMessageQuery{
			SessionID: new(id),
		})
		if err != nil {
			continue
		}
		var readDelta int32
		if session.LastReadMessageID != nil {
			readDelta = int32(latestMsg.ID - *session.LastReadMessageID)
		} else {
			readDelta = int32(latestMsg.ID)
		}
		if readDelta <= 0 {
			continue
		}
		_, err = u.chatSessionRepo.UpdateLastReadMessage(ctx, &repo.ChatSessionUpdateLastReadMessageReq{
			ChatSessionID:      id,
			MessageID:          latestMsg.ID,
			OperationReadCount: readDelta,
			UpdatedBy:          req.UserID,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

type ChatSessionPageReq struct {
	Page     *base.PageRequest
	QueryIDs []int64
	UserID   int64
}

type ChatSessionPageResp struct {
	List []*model.ChatSession
	Page *base.PageResp
}

func (u *ChatSessionUsecase) Page(ctx context.Context, req *ChatSessionPageReq) (*ChatSessionPageResp, error) {
	pageResp, err := u.chatSessionRepo.Page(ctx, &repo.ChatSessionQuery{
		Page:      req.Page,
		IDs:       req.QueryIDs,
		CreatedBy: &req.UserID,
	})
	if err != nil {
		return nil, err
	}
	return &ChatSessionPageResp{
		List: pageResp.Rows,
		Page: pageResp.Page,
	}, nil
}
