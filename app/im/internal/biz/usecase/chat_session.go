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
		sessionResp, err := u.chatSessionRepo.Get(ctx, &repo.ChatSessionGetReq{ChatSessionQuery: repo.ChatSessionQuery{IDs: []int64{id}}})
		if err != nil {
			return err
		}
		latestMsgResp, err := u.chatMessageRepo.Get(ctx, &repo.ChatMessageGetReq{ChatMessageQuery: repo.ChatMessageQuery{SessionID: &id}})
		if err != nil {
			continue
		}
		var readDelta int32
		if sessionResp.ChatSession.LastReadMessageID != nil {
			readDelta = int32(latestMsgResp.ChatMessage.ID - *sessionResp.ChatSession.LastReadMessageID)
		} else {
			readDelta = int32(latestMsgResp.ChatMessage.ID)
		}
		if readDelta <= 0 {
			continue
		}
		_, err = u.chatSessionRepo.UpdateLastReadMessage(ctx, &repo.ChatSessionUpdateLastReadMessageReq{
			ChatSessionID:      id,
			MessageID:          latestMsgResp.ChatMessage.ID,
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

type ChatSessionPageResponse struct {
	List []*model.ChatSession
	Page *base.PageResponse
}

func (u *ChatSessionUsecase) Page(ctx context.Context, req *ChatSessionPageReq) (*ChatSessionPageResponse, error) {
	pageResponse, err := u.chatSessionRepo.Page(ctx, &repo.ChatSessionPageReq{ChatSessionQuery: repo.ChatSessionQuery{
		Page:      req.Page,
		IDs:       req.QueryIDs,
		CreatedBy: &req.UserID,
	}})
	if err != nil {
		return nil, err
	}
	return &ChatSessionPageResponse{List: pageResponse.Rows, Page: pageResponse.Page}, nil
}
