package usecase

import (
	"context"
	"log/slog"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"common/pkg/apperror"
	commonclient "common/pkg/client"
	"common/pkg/constant"
	commonmodel "common/pkg/model"
	cerrors "common/proto/gen/common/errors"
	economyv1 "common/proto/gen/economy/v1"
	economyv1enum "common/proto/gen/economy/v1/enum"
	imv1 "common/proto/gen/im/v1"

	"im/internal/biz/model"
	"im/internal/biz/repo"
	"im/internal/config"
	"im/internal/enum"

	"github.com/dtm-labs/client/dtmgrpc"
)

type GlobalChatBarrageUsecase struct {
	log         *slog.Logger
	barrages    repo.GlobalChatBarrageRepo
	rateLimiter repo.GlobalChatBarrageRateLimiter
	dtm         *commonclient.DTMClient
	branches    *commonclient.TCCBranchCaller
	events      repo.RealtimeEventPublisher
	business    *config.Business
}

func NewGlobalChatBarrageUsecase(log *slog.Logger, barrages repo.GlobalChatBarrageRepo, rateLimiter repo.GlobalChatBarrageRateLimiter, dtm *commonclient.DTMClient, branches *commonclient.TCCBranchCaller, events repo.RealtimeEventPublisher, business *config.Business) *GlobalChatBarrageUsecase {
	return &GlobalChatBarrageUsecase{log: log, barrages: barrages, rateLimiter: rateLimiter, dtm: dtm, branches: branches, events: events, business: business}
}

func (u *GlobalChatBarrageUsecase) Send(ctx context.Context, userID int64, clientRequestID, content string) (*model.GlobalChatBarrage, error) {
	content = strings.TrimSpace(strings.Map(func(r rune) rune {
		if unicode.IsControl(r) {
			return -1
		}
		return r
	}, content))
	if userID <= 0 || clientRequestID == "" || content == "" || utf8.RuneCountInString(content) > 32 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	existing, err := u.barrages.GetByClientRequestID(ctx, clientRequestID)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		if existing.Status == enum.GlobalChatBarrageStatusPublished {
			return existing, nil
		}
		if existing.Status == enum.GlobalChatBarrageStatusCanceled || existing.UserID != userID || existing.Content != content {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_CONFLICT)
		}
	}
	if existing == nil {
		limited, err := u.rateLimiter.Limited(ctx, userID)
		if err != nil {
			return nil, err
		}
		if limited {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_IM_GLOBAL_CHAT_BARRAGE_RATE_LIMITED)
		}
	}
	points := int64(5)
	interval := 12 * time.Second
	if u.business != nil {
		if u.business.GetGlobalChatBarragePoints() > 0 {
			points = u.business.GetGlobalChatBarragePoints()
		}
		if u.business.GetGlobalChatBarrageUserInterval() != nil && u.business.GetGlobalChatBarrageUserInterval().AsDuration() > 0 {
			interval = u.business.GetGlobalChatBarrageUserInterval().AsDuration()
		}
	}
	if existing != nil {
		points = existing.PointsCost
	}
	barrageRequest := &imv1.TCCSendGlobalChatBarrage_Req{UserId: userID, ClientRequestId: clientRequestID, Content: content, PointsCost: points}
	pointsRequest := &economyv1.TCCPoints_Req{UserId: userID, Amount: points, RecordType: economyv1enum.EconomyRecordType_ECONOMY_RECORD_TYPE_GLOBAL_CHAT_BARRAGE_EXPENSE}
	_, err = u.dtm.InTCC(ctx, func(tcc *dtmgrpc.TccGrpc) error {
		if err = u.branches.Call(ctx, tcc, constant.IMServiceName.String(), barrageRequest, imv1.TCCGlobalChatBarrageService_Try_FullMethodName, imv1.TCCGlobalChatBarrageService_Confirm_FullMethodName, imv1.TCCGlobalChatBarrageService_Cancel_FullMethodName, &imv1.TCCSendGlobalChatBarrage_Resp{}); err != nil {
			return err
		}
		return u.branches.Call(ctx, tcc, constant.EconomyServiceName.String(), pointsRequest, economyv1.TCCPointsService_Try_FullMethodName, economyv1.TCCPointsService_Confirm_FullMethodName, economyv1.TCCPointsService_Cancel_FullMethodName, &economyv1.TCCPoints_Resp{})
	})
	if err != nil {
		return nil, err
	}
	barrage, err := u.barrages.GetByClientRequestID(ctx, clientRequestID)
	if err != nil || barrage == nil {
		return barrage, err
	}
	if existing == nil && barrage.Status == enum.GlobalChatBarrageStatusPublished {
		if err = u.rateLimiter.Mark(context.WithoutCancel(ctx), userID, interval); err != nil {
			u.log.WarnContext(ctx, "mark global chat barrage rate limit failed", "user_id", userID, "error", err)
		}
		if err = u.events.PublishGlobalChatBarrageCreated(ctx, barrage); err != nil {
			u.log.WarnContext(ctx, "publish global chat barrage event failed", "barrage_id", barrage.ID, "error", err)
		}
	}
	return barrage, nil
}

func (u *GlobalChatBarrageUsecase) List(ctx context.Context, userID *int64, page *commonmodel.PageReq) ([]*model.GlobalChatBarrage, *commonmodel.PageResp, error) {
	p := commonmodel.NormalizePage(page)
	rows, total, err := u.barrages.List(ctx, userID, p.Limit(), p.Offset())
	if err != nil {
		return nil, nil, err
	}
	return rows, &commonmodel.PageResp{Page: p.Page, Size: p.Size, Total: total}, nil
}
