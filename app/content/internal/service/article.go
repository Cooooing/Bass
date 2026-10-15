package service

import (
	"common/pkg/apperror"
	commonmodel "common/pkg/model"
	"common/pkg/util"
	"common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/content/v1"
	"context"
	"time"

	"content/internal/biz/model"
	"content/internal/biz/usecase"
	"content/internal/enum"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ArticleService struct {
	v1.UnimplementedContentArticleServiceServer

	articleUsecase *usecase.ArticleUsecase
}

func NewArticleService(
	articleUsecase *usecase.ArticleUsecase,
) *ArticleService {
	return &ArticleService{
		articleUsecase: articleUsecase,
	}
}

func (s *ArticleService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterContentArticleServiceServer(gs, s)
}

func (s *ArticleService) RegisterHTTP(hs *http.Server) {
}

func (s *ArticleService) CreateDraft(ctx context.Context, req *v1.CreateDraftArticle_Req) (*v1.CreateDraftArticle_Resp, error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize("")
	if err != nil {
		return nil, err
	}
	articleReq := req.GetArticle()
	if articleReq == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	articleType, ok := enum.ArticleTypeMap.ToEnum(articleReq.GetType())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_TYPE)
	}
	row, err := s.articleUsecase.Add(ctx, &usecase.ArticleAddReq{
		Access: access,
		Article: &model.Article{
			Title:         articleReq.GetTitle(),
			Content:       articleReq.GetContent(),
			RewardContent: articleReq.RewardContent,
			RewardPoints:  articleReq.RewardPoints,
			Type:          articleType,
			Statement:     articleReq.Statement,
			Commentable:   util.DerefOrDefault(articleReq.Commentable, true),
		},
		Poll:    s.articlePollDraft(articleReq.GetPoll()),
		Lottery: s.articleLotteryDraft(articleReq.GetLottery()),
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateDraftArticle_Resp{ArticleId: row.ID}, nil
}

func (s *ArticleService) UpdateDraft(ctx context.Context, req *v1.UpdateDraftArticle_Req) (*v1.UpdateDraftArticle_Resp, error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize("")
	if err != nil {
		return nil, err
	}
	articleReq := req.GetArticle()
	if articleReq == nil || req.GetArticleId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	articleType, ok := enum.ArticleTypeMap.ToEnum(articleReq.GetType())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_TYPE)
	}
	row, err := s.articleUsecase.Update(ctx, &usecase.ArticleUpdateDraftReq{
		Access: access,
		Article: &model.Article{
			ID:            req.GetArticleId(),
			Title:         articleReq.GetTitle(),
			Content:       articleReq.GetContent(),
			RewardContent: articleReq.RewardContent,
			RewardPoints:  articleReq.RewardPoints,
			Type:          articleType,
			Statement:     articleReq.Statement,
			Commentable:   util.DerefOrDefault(articleReq.Commentable, true),
			UpdatedBy:     new(access.ActorUserID),
		},
		Poll:    s.articlePollDraft(articleReq.GetPoll()),
		Lottery: s.articleLotteryDraft(articleReq.GetLottery()),
	})
	if err != nil {
		return nil, err
	}
	return &v1.UpdateDraftArticle_Resp{ArticleId: row.ID}, nil
}

func (s *ArticleService) VotePoll(ctx context.Context, req *v1.VoteArticlePoll_Req) (*v1.VoteArticlePoll_Resp, error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize("")
	if err != nil {
		return nil, err
	}
	if req.GetArticleId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err = s.articleUsecase.VotePoll(ctx, &usecase.ArticleVotePollReq{Access: access, ArticleID: req.GetArticleId(), OptionIDs: req.GetOptionIds()})
	return &v1.VoteArticlePoll_Resp{}, err
}

func (*ArticleService) articlePollDraft(req *v1.ArticlePollDraft) *model.ArticlePollDraft {
	if req == nil || req.GetEndsAt() == nil {
		return nil
	}
	return &model.ArticlePollDraft{Options: req.GetOptions(), EndsAt: req.GetEndsAt().AsTime(), MaxChoices: req.GetMaxChoices()}
}

func (*ArticleService) articleLotteryDraft(req *v1.ArticleLotteryDraft) *model.ArticleLotteryDraft {
	if req == nil || req.GetDrawAt() == nil {
		return nil
	}
	prizes := make([]*model.ArticleLotteryPrizeDraft, 0, len(req.GetPrizes()))
	for _, prize := range req.GetPrizes() {
		if prize != nil {
			prizes = append(prizes, &model.ArticleLotteryPrizeDraft{Name: prize.GetName(), Description: prize.Description, WinnerLimit: prize.GetWinnerLimit()})
		}
	}
	return &model.ArticleLotteryDraft{Prizes: prizes, DrawAt: req.GetDrawAt().AsTime()}
}

func (s *ArticleService) EnterLottery(ctx context.Context, req *v1.EnterArticleLottery_Req) (*v1.EnterArticleLottery_Resp, error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize("")
	if err != nil {
		return nil, err
	}
	if req.GetArticleId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err = s.articleUsecase.EnterLottery(ctx, &usecase.ArticleEnterLotteryReq{Access: access, ArticleID: req.GetArticleId(), Entered: req.GetEntered()})
	return &v1.EnterArticleLottery_Resp{}, err
}

func (s *ArticleService) DrawLottery(ctx context.Context, req *v1.DrawArticleLottery_Req) (*v1.DrawArticleLottery_Resp, error) {
	if req.GetArticleId() <= 0 || req.GetDrawAt() == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err := s.articleUsecase.DrawLottery(ctx, &usecase.ArticleDrawLotteryReq{ArticleID: req.GetArticleId(), DrawAt: req.GetDrawAt().AsTime()})
	return &v1.DrawArticleLottery_Resp{}, err
}

func (s *ArticleService) Publish(ctx context.Context, req *v1.PublishArticle_Req) (*v1.PublishArticle_Resp, error) {
	access := &model.ContentAccess{Scope: enum.ContentAccessScopeInternalTask}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize(enum.ContentAccessScopeInternalTask)
	if err != nil {
		return nil, err
	}
	var scheduledAt *time.Time
	if req.GetScheduledAt() != nil {
		value := req.GetScheduledAt().AsTime()
		scheduledAt = &value
	}
	if err := s.articleUsecase.Publish(ctx, &usecase.ArticlePublishReq{Access: access, ArticleID: req.GetArticleId(), ScheduledAt: scheduledAt, PublishedCity: req.City}); err != nil {
		return nil, err
	}
	return &v1.PublishArticle_Resp{}, nil
}

func (s *ArticleService) CancelPublish(ctx context.Context, req *v1.CancelPublishArticle_Req) (*v1.CancelPublishArticle_Resp, error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize("")
	if err != nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if err := s.articleUsecase.CancelPublish(ctx, &usecase.ArticleCancelPublishReq{Access: access, ArticleID: req.GetArticleId()}); err != nil {
		return nil, err
	}
	return &v1.CancelPublishArticle_Resp{}, nil
}

func (s *ArticleService) DiscardDraft(ctx context.Context, req *v1.DiscardDraftArticle_Req) (*v1.DiscardDraftArticle_Resp, error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize("")
	if err != nil {
		return nil, err
	}
	return &v1.DiscardDraftArticle_Resp{}, s.articleUsecase.DiscardDraft(ctx, &usecase.ArticleDiscardDraftReq{Access: access, ArticleID: req.GetArticleId()})
}

func (s *ArticleService) MakePrivate(ctx context.Context, req *v1.MakePrivateArticle_Req) (*v1.MakePrivateArticle_Resp, error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize("")
	if err != nil {
		return nil, err
	}
	return &v1.MakePrivateArticle_Resp{}, s.articleUsecase.MakePrivate(ctx, &usecase.ArticleMakePrivateReq{Access: access, ArticleID: req.GetArticleId()})
}

func (s *ArticleService) MakePublic(ctx context.Context, req *v1.MakePublicArticle_Req) (*v1.MakePublicArticle_Resp, error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize("")
	if err != nil {
		return nil, err
	}
	return &v1.MakePublicArticle_Resp{}, s.articleUsecase.MakePublic(ctx, &usecase.ArticleMakePublicReq{Access: access, ArticleID: req.GetArticleId()})
}

func (s *ArticleService) Archive(ctx context.Context, req *v1.ArchiveArticle_Req) (*v1.ArchiveArticle_Resp, error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize("")
	if err != nil {
		return nil, err
	}
	return &v1.ArchiveArticle_Resp{}, s.articleUsecase.Archive(ctx, &usecase.ArticleArchiveReq{Access: access, ArticleID: req.GetArticleId(), Reason: req.Reason})
}

func (s *ArticleService) SetPinnedComment(ctx context.Context, req *v1.SetPinnedArticleComment_Req) (*v1.SetPinnedArticleComment_Resp, error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	err := s.articleUsecase.SetPinnedComment(ctx, access, req.GetArticleId(), req.CommentId)
	return &v1.SetPinnedArticleComment_Resp{}, err
}

func (s *ArticleService) AcceptAnswer(ctx context.Context, req *v1.AcceptArticleAnswer_Req) (*v1.AcceptArticleAnswer_Resp, error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	err := s.articleUsecase.AcceptAnswer(ctx, access, req.GetArticleId(), req.GetCommentId())
	return &v1.AcceptArticleAnswer_Resp{}, err
}

func (s *ArticleService) Unarchive(ctx context.Context, req *v1.UnarchiveArticle_Req) (*v1.UnarchiveArticle_Resp, error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize("")
	if err != nil {
		return nil, err
	}
	return &v1.UnarchiveArticle_Resp{}, s.articleUsecase.Unarchive(ctx, &usecase.ArticleUnarchiveReq{Access: access, ArticleID: req.GetArticleId(), Reason: req.Reason})
}

func (s *ArticleService) Hide(ctx context.Context, req *v1.HideArticle_Req) (*v1.HideArticle_Resp, error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize("")
	if err != nil {
		return nil, err
	}
	return &v1.HideArticle_Resp{}, s.articleUsecase.Hide(ctx, &usecase.ArticleHideReq{Access: access, ArticleID: req.GetArticleId(), Reason: req.Reason})
}

func (s *ArticleService) Unhide(ctx context.Context, req *v1.UnhideArticle_Req) (*v1.UnhideArticle_Resp, error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize("")
	if err != nil {
		return nil, err
	}
	return &v1.UnhideArticle_Resp{}, s.articleUsecase.Unhide(ctx, &usecase.ArticleUnhideReq{Access: access, ArticleID: req.GetArticleId(), Reason: req.Reason})
}

func (s *ArticleService) Lock(ctx context.Context, req *v1.LockArticle_Req) (*v1.LockArticle_Resp, error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize("")
	if err != nil {
		return nil, err
	}
	return &v1.LockArticle_Resp{}, s.articleUsecase.Lock(ctx, &usecase.ArticleLockReq{Access: access, ArticleID: req.GetArticleId(), Reason: req.Reason})
}

func (s *ArticleService) Unlock(ctx context.Context, req *v1.UnlockArticle_Req) (*v1.UnlockArticle_Resp, error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize("")
	if err != nil {
		return nil, err
	}
	return &v1.UnlockArticle_Resp{}, s.articleUsecase.Unlock(ctx, &usecase.ArticleUnlockReq{Access: access, ArticleID: req.GetArticleId(), Reason: req.Reason})
}

func (s *ArticleService) Like(ctx context.Context, req *v1.LikeArticle_Req) (*v1.LikeArticle_Resp, error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize("")
	if err != nil {
		return nil, err
	}
	liked, err := s.articleUsecase.Like(ctx, &usecase.ArticleLikeReq{Access: access, ArticleID: req.GetArticleId(), UserID: access.ActorUserID, Active: req.GetLiked()})
	return &v1.LikeArticle_Resp{Liked: liked}, err
}

func (s *ArticleService) Thank(ctx context.Context, req *v1.ThankArticle_Req) (*v1.ThankArticle_Resp, error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize("")
	if err != nil {
		return nil, err
	}
	thanked, err := s.articleUsecase.Thank(ctx, &usecase.ArticleThankReq{Access: access, ArticleID: req.GetArticleId(), UserID: access.ActorUserID, Active: req.GetThanked()})
	return &v1.ThankArticle_Resp{Thanked: thanked}, err
}

func (s *ArticleService) Collect(ctx context.Context, req *v1.CollectArticle_Req) (*v1.CollectArticle_Resp, error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize("")
	if err != nil {
		return nil, err
	}
	collected, err := s.articleUsecase.Collect(ctx, &usecase.ArticleCollectReq{Access: access, ArticleID: req.GetArticleId(), UserID: access.ActorUserID, Active: req.GetCollected()})
	return &v1.CollectArticle_Resp{Collected: collected}, err
}

func (s *ArticleService) Reward(ctx context.Context, req *v1.RewardArticle_Req) (*v1.RewardArticle_Resp, error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize("")
	if err != nil {
		return nil, err
	}
	return &v1.RewardArticle_Resp{}, s.articleUsecase.Reward(ctx, &usecase.ArticleRewardReq{Access: access, ArticleID: req.GetArticleId(), UserID: access.ActorUserID, Points: req.GetPoints()})
}

func (s *ArticleService) View(ctx context.Context, req *v1.ViewArticle_Req) (*v1.ViewArticle_Resp, error) {
	access := &model.ContentAccess{Scope: enum.ContentAccessScopeGuest}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize(enum.ContentAccessScopeGuest)
	if err != nil {
		return nil, err
	}
	return &v1.ViewArticle_Resp{}, s.articleUsecase.View(ctx, &usecase.ArticleViewReq{Access: access, ArticleID: req.GetArticleId(), IP: req.Ip, UserAgent: req.UserAgent, BrowserFingerprint: req.BrowserFingerprint})
}

func (s *ArticleService) PageViewHistory(
	ctx context.Context,
	req *v1.PageArticleViewHistory_Req,
) (*v1.PageArticleViewHistory_Resp, error) {
	access := &model.ContentAccess{Scope: enum.ContentAccessScopeUser}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	pageRequest := commonmodel.PageReqFromProto(req.GetPage())
	result, err := s.articleUsecase.PageViewHistory(ctx, access, pageRequest, req.GetExcludedAuthorIds())
	if err != nil {
		return nil, err
	}
	rows := make([]*v1.ArticleViewHistoryItem, 0, len(result.Rows))
	for _, row := range result.Rows {
		if row == nil || row.Article == nil {
			continue
		}
		item := &v1.ArticleViewHistoryItem{Article: s.articleViewHistoryArticle(row.Article)}
		if row.ViewedAt != nil {
			item.ViewedAt = timestamppb.New(*row.ViewedAt)
		}
		rows = append(rows, item)
	}
	var pageResponse *common.PageResp
	if result.Page != nil {
		pageResponse = result.Page.ToProto()
	}
	return &v1.PageArticleViewHistory_Resp{Rows: rows, Page: pageResponse}, nil
}

func (s *ArticleService) articleViewHistoryArticle(row *model.Article) *v1.Article {
	if row == nil {
		return nil
	}
	article := &v1.Article{
		Id:                row.ID,
		Title:             row.Title,
		Content:           row.Content,
		RewardContent:     row.RewardContent,
		RewardPoints:      row.RewardPoints,
		HasPostscript:     row.HasPostscript,
		HasReward:         row.RewardPoints != nil,
		Type:              enum.ArticleTypeMap.MustToProto(row.Type),
		Statement:         row.Statement,
		Commentable:       row.Commentable,
		PublishStatus:     enum.ArticlePublishStatusMap.MustToProto(row.PublishStatus),
		Visibility:        enum.ArticleVisibilityMap.MustToProto(row.Visibility),
		Restriction:       enum.ContentRestrictionMap.MustToProto(row.Restriction),
		ViewCount:         row.ViewCount,
		ThankCount:        row.ThankCount,
		LikeCount:         row.LikeCount,
		CollectCount:      row.CollectCount,
		RewardCount:       row.RewardCount,
		ReplyCount:        row.ReplyCount,
		CreatedBy:         row.CreatedBy,
		UpdatedBy:         row.UpdatedBy,
		PublishedCity:     row.PublishedCity,
		PinnedCommentId:   row.PinnedID,
		AcceptedCommentId: row.AcceptedCommentID,
		Poll:              s.articlePoll(row.Poll),
		Lottery:           s.articleLottery(row.Lottery),
	}
	if row.CreatedAt != nil {
		article.CreatedAt = timestamppb.New(*row.CreatedAt)
	}
	if row.UpdatedAt != nil {
		article.UpdatedAt = timestamppb.New(*row.UpdatedAt)
	}
	if row.PublishedAt != nil {
		article.PublishedAt = timestamppb.New(*row.PublishedAt)
	}
	if row.EditedAt != nil {
		article.EditedAt = timestamppb.New(*row.EditedAt)
	}
	return article
}

func (*ArticleService) articlePoll(row *model.ArticlePoll) *v1.ArticlePoll {
	if row == nil {
		return nil
	}
	poll := &v1.ArticlePoll{MaxChoices: row.MaxChoices, ParticipantCount: row.ParticipantCount, TotalVoteCount: row.TotalVoteCount, SelectedOptionIds: row.SelectedOptionIDs, ResultVisible: row.ResultVisible, Options: make([]*v1.ArticlePollOption, 0, len(row.Options))}
	if !row.EndsAt.IsZero() {
		poll.EndsAt = timestamppb.New(row.EndsAt)
	}
	for _, option := range row.Options {
		if option == nil {
			continue
		}
		poll.Options = append(poll.Options, &v1.ArticlePollOption{Id: option.ID, Content: option.Content, VoteCount: option.VoteCount})
	}
	return poll
}

func (*ArticleService) articleLottery(row *model.ArticleLottery) *v1.ArticleLottery {
	if row == nil {
		return nil
	}
	lottery := &v1.ArticleLottery{EntryCount: row.EntryCount, Entered: row.Entered, Prizes: make([]*v1.ArticleLotteryPrize, 0, len(row.Prizes))}
	if !row.DrawAt.IsZero() {
		lottery.DrawAt = timestamppb.New(row.DrawAt)
	}
	if row.DrawnAt != nil {
		lottery.DrawnAt = timestamppb.New(*row.DrawnAt)
	}
	for _, prize := range row.Prizes {
		if prize != nil {
			lottery.Prizes = append(lottery.Prizes, &v1.ArticleLotteryPrize{Id: prize.ID, Name: prize.Name, Description: prize.Description, WinnerLimit: prize.WinnerLimit, WinnerUserIds: prize.WinnerUserIDs})
		}
	}
	return lottery
}

func (s *ArticleService) FlushViews(ctx context.Context, req *v1.FlushArticleViews_Req) (*v1.FlushArticleViews_Resp, error) {
	flushed, err := s.articleUsecase.FlushViews(ctx, &usecase.ArticleFlushViewsReq{Limit: req.GetLimit()})
	if err != nil {
		return nil, err
	}
	return &v1.FlushArticleViews_Resp{Flushed: flushed}, nil
}

func (s *ArticleService) Get(ctx context.Context, req *v1.GetArticle_Req) (*v1.GetArticle_Resp, error) {
	access := &model.ContentAccess{Scope: enum.ContentAccessScopeGuest}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize(enum.ContentAccessScopeGuest)
	if err != nil {
		return nil, err
	}
	row, err := s.articleUsecase.GetByAccess(ctx, &usecase.ArticleGetByAccessReq{ArticleID: req.GetArticleId(), Access: access})
	if err != nil {
		return nil, err
	}
	article := &v1.Article{
		Id:                row.ID,
		Title:             row.Title,
		Content:           row.Content,
		RewardContent:     row.RewardContent,
		RewardPoints:      row.RewardPoints,
		HasPostscript:     row.HasPostscript,
		HasReward:         row.RewardPoints != nil,
		Type:              enum.ArticleTypeMap.MustToProto(row.Type),
		Statement:         row.Statement,
		Commentable:       row.Commentable,
		PublishStatus:     enum.ArticlePublishStatusMap.MustToProto(row.PublishStatus),
		Visibility:        enum.ArticleVisibilityMap.MustToProto(row.Visibility),
		Restriction:       enum.ContentRestrictionMap.MustToProto(row.Restriction),
		ViewCount:         row.ViewCount,
		ThankCount:        row.ThankCount,
		LikeCount:         row.LikeCount,
		CollectCount:      row.CollectCount,
		RewardCount:       row.RewardCount,
		ReplyCount:        row.ReplyCount,
		CreatedBy:         row.CreatedBy,
		UpdatedBy:         row.UpdatedBy,
		PublishedCity:     row.PublishedCity,
		PinnedCommentId:   row.PinnedID,
		AcceptedCommentId: row.AcceptedCommentID,
		Poll:              s.articlePoll(row.Poll),
		Lottery:           s.articleLottery(row.Lottery),
	}
	if row.CreatedAt != nil {
		article.CreatedAt = timestamppb.New(*row.CreatedAt)
	}
	if row.UpdatedAt != nil {
		article.UpdatedAt = timestamppb.New(*row.UpdatedAt)
	}
	if row.PublishedAt != nil {
		article.PublishedAt = timestamppb.New(*row.PublishedAt)
	}
	if row.EditedAt != nil {
		article.EditedAt = timestamppb.New(*row.EditedAt)
	}
	return &v1.GetArticle_Resp{Article: article}, nil
}

func (s *ArticleService) List(ctx context.Context, req *v1.ListArticles_Req) (*v1.ListArticles_Resp, error) {
	access := &model.ContentAccess{Scope: enum.ContentAccessScopeGuest}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize(enum.ContentAccessScopeGuest)
	if err != nil {
		return nil, err
	}
	query, err := s.articleQueryFilter(req.GetQuery())
	if err != nil {
		return nil, err
	}
	list, err := s.articleUsecase.ListByAccess(ctx, &usecase.ArticleListByAccessReq{Access: access, Filter: query})
	if err != nil {
		return nil, err
	}
	rows := make([]*v1.Article, 0, len(list))
	for _, row := range list {
		article := &v1.Article{
			Id:            row.ID,
			Title:         row.Title,
			Content:       row.Content,
			RewardContent: row.RewardContent,
			RewardPoints:  row.RewardPoints,
			HasPostscript: row.HasPostscript,
			HasReward:     row.RewardPoints != nil,
			Type:          enum.ArticleTypeMap.MustToProto(row.Type),
			Statement:     row.Statement,
			Commentable:   row.Commentable,
			PublishStatus: enum.ArticlePublishStatusMap.MustToProto(row.PublishStatus),
			Visibility:    enum.ArticleVisibilityMap.MustToProto(row.Visibility),
			Restriction:   enum.ContentRestrictionMap.MustToProto(row.Restriction),
			ViewCount:     row.ViewCount,
			ThankCount:    row.ThankCount,
			LikeCount:     row.LikeCount,
			CollectCount:  row.CollectCount,
			RewardCount:   row.RewardCount,
			ReplyCount:    row.ReplyCount,
			CreatedBy:     row.CreatedBy,
			UpdatedBy:     row.UpdatedBy,
		}
		if row.CreatedAt != nil {
			article.CreatedAt = timestamppb.New(*row.CreatedAt)
		}
		if row.UpdatedAt != nil {
			article.UpdatedAt = timestamppb.New(*row.UpdatedAt)
		}
		if row.PublishedAt != nil {
			article.PublishedAt = timestamppb.New(*row.PublishedAt)
		}
		if row.EditedAt != nil {
			article.EditedAt = timestamppb.New(*row.EditedAt)
		}
		rows = append(rows, article)
	}
	return &v1.ListArticles_Resp{Rows: rows}, nil
}

func (s *ArticleService) Page(ctx context.Context, req *v1.PageArticles_Req) (*v1.PageArticles_Resp, error) {
	access := &model.ContentAccess{Scope: enum.ContentAccessScopeGuest}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize(enum.ContentAccessScopeGuest)
	if err != nil {
		return nil, err
	}
	query, err := s.articleQueryFilter(req.GetQuery())
	if err != nil {
		return nil, err
	}
	pageResp, err := s.articleUsecase.PageByAccess(ctx, &usecase.ArticlePageByAccessReq{Access: access, Filter: query, Page: commonmodel.PageReqFromProto(req.GetPage())})
	if err != nil {
		return nil, err
	}
	rows := make([]*v1.Article, 0, len(pageResp.Rows))
	for _, row := range pageResp.Rows {
		article := &v1.Article{
			Id:            row.ID,
			Title:         row.Title,
			Content:       row.Content,
			RewardContent: row.RewardContent,
			RewardPoints:  row.RewardPoints,
			HasPostscript: row.HasPostscript,
			HasReward:     row.RewardPoints != nil,
			Type:          enum.ArticleTypeMap.MustToProto(row.Type),
			Statement:     row.Statement,
			Commentable:   row.Commentable,
			PublishStatus: enum.ArticlePublishStatusMap.MustToProto(row.PublishStatus),
			Visibility:    enum.ArticleVisibilityMap.MustToProto(row.Visibility),
			Restriction:   enum.ContentRestrictionMap.MustToProto(row.Restriction),
			ViewCount:     row.ViewCount,
			ThankCount:    row.ThankCount,
			LikeCount:     row.LikeCount,
			CollectCount:  row.CollectCount,
			RewardCount:   row.RewardCount,
			ReplyCount:    row.ReplyCount,
			CreatedBy:     row.CreatedBy,
			UpdatedBy:     row.UpdatedBy,
		}
		if row.CreatedAt != nil {
			article.CreatedAt = timestamppb.New(*row.CreatedAt)
		}
		if row.UpdatedAt != nil {
			article.UpdatedAt = timestamppb.New(*row.UpdatedAt)
		}
		if row.PublishedAt != nil {
			article.PublishedAt = timestamppb.New(*row.PublishedAt)
		}
		if row.EditedAt != nil {
			article.EditedAt = timestamppb.New(*row.EditedAt)
		}
		rows = append(rows, article)
	}
	return &v1.PageArticles_Resp{Rows: rows, Page: pageResp.Page.ToProto()}, nil
}

func (s *ArticleService) MapViewerActionStates(ctx context.Context, req *v1.MapArticleViewerActionStates_Req) (*v1.MapArticleViewerActionStates_Resp, error) {
	states, err := s.articleUsecase.MapViewerActionStates(ctx, &usecase.ArticleMapViewerActionStatesReq{ArticleIDs: req.GetArticleIds(), UserID: req.GetUserId()})
	if err != nil {
		return nil, err
	}
	reply := make(map[int64]*v1.MapArticleViewerActionStates_Resp_ArticleViewerActionState, len(states))
	for articleID, state := range states {
		reply[articleID] = &v1.MapArticleViewerActionStates_Resp_ArticleViewerActionState{Liked: state.Liked, Thanked: state.Thanked, Collected: state.Collected, Rewarded: state.Rewarded}
	}
	return &v1.MapArticleViewerActionStates_Resp{States: reply}, nil
}

func (*ArticleService) articleQueryFilter(queryParams *v1.ArticleQueryParams) (*model.ArticleFilter, error) {
	query := &model.ArticleFilter{}
	if queryParams == nil {
		return query, nil
	}
	query.TagID = queryParams.TagId
	query.DomainID = queryParams.DomainId
	query.ArticleIDs = queryParams.GetArticleIds()
	query.AuthorID = queryParams.AuthorId
	query.AuthorIDs = queryParams.AuthorIds
	query.PublishedCity = queryParams.PublishedCity
	query.ExcludedAuthorIDs = queryParams.GetExcludedAuthorIds()
	if queryParams.Type != nil {
		value, ok := enum.ArticleTypeMap.ToEnum(queryParams.GetType())
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_TYPE)
		}
		query.Type = new(value)
	}
	if queryParams.Order != nil {
		value, ok := enum.ArticleOrderMap.ToEnum(queryParams.GetOrder())
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		query.Order = new(value)
	}
	if queryParams.PublishStatus != nil {
		value, ok := enum.ArticlePublishStatusMap.ToEnum(queryParams.GetPublishStatus())
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_STATUS)
		}
		query.PublishStatus = new(value)
	}
	for _, item := range queryParams.GetPublishStatuses() {
		value, ok := enum.ArticlePublishStatusMap.ToEnum(item)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_STATUS)
		}
		query.PublishStatuses = append(query.PublishStatuses, value)
	}
	if queryParams.Visibility != nil {
		value, ok := enum.ArticleVisibilityMap.ToEnum(queryParams.GetVisibility())
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_STATUS)
		}
		query.Visibility = new(value)
	}
	for _, item := range queryParams.GetVisibilities() {
		value, ok := enum.ArticleVisibilityMap.ToEnum(item)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_STATUS)
		}
		query.Visibilities = append(query.Visibilities, value)
	}
	if queryParams.Restriction != nil {
		value, ok := enum.ContentRestrictionMap.ToEnum(queryParams.GetRestriction())
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		query.Restriction = new(value)
	}
	for _, item := range queryParams.GetRestrictions() {
		value, ok := enum.ContentRestrictionMap.ToEnum(item)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		query.Restrictions = append(query.Restrictions, value)
	}
	if queryParams.GetPublishedAtEnd() != nil {
		publishedAtEnd := queryParams.GetPublishedAtEnd().AsTime()
		query.PublishedAtEnd = new(publishedAtEnd)
	}
	query.Scheduled = queryParams.Scheduled
	return query, nil
}
