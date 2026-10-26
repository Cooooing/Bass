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

// GlobalChatTopicUsecase 处理全局聊天室话题的外部设置命令。
type GlobalChatTopicUsecase struct {
	log         *slog.Logger
	topics      repo.GlobalChatTopicRepo
	rateLimiter repo.GlobalChatTopicRateLimiter
	dtm         *commonclient.DTMClient
	branches    *commonclient.TCCBranchCaller
	events      repo.RealtimeEventPublisher
	business    *config.Business
}

func NewGlobalChatTopicUsecase(log *slog.Logger, topics repo.GlobalChatTopicRepo, rateLimiter repo.GlobalChatTopicRateLimiter, dtm *commonclient.DTMClient, branches *commonclient.TCCBranchCaller, events repo.RealtimeEventPublisher, business *config.Business) *GlobalChatTopicUsecase {
	return &GlobalChatTopicUsecase{log: log, topics: topics, rateLimiter: rateLimiter, dtm: dtm, branches: branches, events: events, business: business}
}

func (u *GlobalChatTopicUsecase) Get(ctx context.Context) (*model.GlobalChatTopic, error) {
	return u.topics.GetCurrent(ctx)
}

func (u *GlobalChatTopicUsecase) Set(ctx context.Context, userID int64, content string) (*model.GlobalChatTopic, error) {
	content = strings.TrimSpace(strings.Map(func(r rune) rune {
		if unicode.IsControl(r) {
			return -1
		}
		return r
	}, content))
	if userID <= 0 || content == "" || utf8.RuneCountInString(content) > 64 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	points := int64(16)
	userInterval := 10 * time.Minute
	globalInterval := 30 * time.Second
	if u.business != nil {
		if u.business.GetGlobalChatTopicPoints() > 0 {
			points = u.business.GetGlobalChatTopicPoints()
		}
		if u.business.GetGlobalChatTopicUserInterval() != nil && u.business.GetGlobalChatTopicUserInterval().AsDuration() > 0 {
			userInterval = u.business.GetGlobalChatTopicUserInterval().AsDuration()
		}
		if u.business.GetGlobalChatTopicGlobalInterval() != nil && u.business.GetGlobalChatTopicGlobalInterval().AsDuration() > 0 {
			globalInterval = u.business.GetGlobalChatTopicGlobalInterval().AsDuration()
		}
	}
	limited, err := u.rateLimiter.Limited(ctx, userID)
	if err != nil {
		return nil, err
	}
	if limited {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_IM_GLOBAL_CHAT_TOPIC_RATE_LIMITED)
	}
	topicID, err := u.topics.NextID(ctx)
	if err != nil {
		return nil, err
	}
	topicRequest := &imv1.TCCSetGlobalChatTopic_Req{TopicId: topicID, UserId: userID, Content: content, PointsCost: points}
	pointsRequest := &economyv1.TCCPoints_Req{UserId: userID, Amount: points, RecordType: economyv1enum.EconomyRecordType_ECONOMY_RECORD_TYPE_GLOBAL_CHAT_TOPIC_EXPENSE}
	_, err = u.dtm.InTCC(ctx, func(tcc *dtmgrpc.TccGrpc) error {
		if err = u.branches.Call(ctx, tcc, constant.IMServiceName.String(), topicRequest, imv1.TCCGlobalChatTopicService_Try_FullMethodName, imv1.TCCGlobalChatTopicService_Confirm_FullMethodName, imv1.TCCGlobalChatTopicService_Cancel_FullMethodName, &imv1.TCCSetGlobalChatTopic_Resp{}); err != nil {
			return err
		}
		return u.branches.Call(ctx, tcc, constant.EconomyServiceName.String(), pointsRequest, economyv1.TCCPointsService_Try_FullMethodName, economyv1.TCCPointsService_Confirm_FullMethodName, economyv1.TCCPointsService_Cancel_FullMethodName, &economyv1.TCCPoints_Resp{})
	})
	if err != nil {
		return nil, err
	}
	topic, err := u.topics.Get(ctx, topicID)
	if err != nil {
		return nil, err
	}
	if topic != nil && topic.Status == enum.GlobalChatTopicStatusPublished {
		if err = u.rateLimiter.Mark(context.WithoutCancel(ctx), userID, userInterval, globalInterval); err != nil {
			u.log.WarnContext(ctx, "mark global chat topic rate limit failed", "user_id", userID, "error", err)
		}
		if err = u.events.PublishGlobalChatTopicUpdated(ctx, topic); err != nil {
			u.log.WarnContext(ctx, "publish global chat topic event failed", "topic_id", topic.ID, "error", err)
		}
	}
	return topic, nil
}
