package usecase

import (
	"context"
	"fmt"

	"common/pkg/apperror"
	commonclient "common/pkg/client"
	"common/pkg/constant"
	cerrors "common/proto/gen/common/errors"
	contentv1 "common/proto/gen/content/v1"
	economyv1 "common/proto/gen/economy/v1"
	economyv1enum "common/proto/gen/economy/v1/enum"

	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/enum"

	"github.com/dtm-labs/client/dtmgrpc"
)

// ArticleRewardUsecase 编排文章打赏的积分转移。
type ArticleRewardUsecase struct {
	articles repo.ArticleRepo
	actions  repo.ArticleActionRecordRepo
	dtm      *commonclient.DTMClient
	branches *commonclient.TCCBranchCaller
}

func NewArticleRewardUsecase(articles repo.ArticleRepo, actions repo.ArticleActionRecordRepo, dtm *commonclient.DTMClient, branches *commonclient.TCCBranchCaller) *ArticleRewardUsecase {
	return &ArticleRewardUsecase{articles: articles, actions: actions, dtm: dtm, branches: branches}
}

func (u *ArticleRewardUsecase) Reward(ctx context.Context, access *model.ContentAccess, articleID int64) (bool, error) {
	article, err := u.articles.Get(ctx, &repo.ArticleGetReq{Filter: &model.ArticleFilter{ArticleID: &articleID}})
	if err != nil {
		return false, err
	}
	if err = (&ArticleAccessUsecase{}).CanInteract(access, article); err != nil {
		return false, err
	}
	if article.CreatedBy == nil || article.RewardContent == nil || article.RewardPoints == nil || *article.RewardPoints <= 0 {
		return false, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_CONFLICT)
	}
	actor := access.ActorUserID
	action := enum.ArticleActionReward
	status := enum.ContentActionStatusPublished
	exists, err := u.actions.Exist(ctx, &repo.ArticleActionRecordReq{ArticleID: &articleID, UserID: &actor, Type: &action, Status: &status})
	if err != nil || exists {
		return exists, err
	}
	points := int64(*article.RewardPoints)
	reward := &contentv1.TCCArticleReward_Req{ArticleId: articleID, SenderUserId: actor, ReceiverUserId: *article.CreatedBy, Points: points}
	transfer := &economyv1.TCCPointsTransfer_Req{
		FromUserId:     actor,
		ToUserId:       *article.CreatedBy,
		Amount:         points,
		OutRecordType:  economyv1enum.EconomyRecordType_ECONOMY_RECORD_TYPE_ARTICLE_REWARD_OUT,
		InRecordType:   economyv1enum.EconomyRecordType_ECONOMY_RECORD_TYPE_ARTICLE_REWARD_IN,
		IdempotencyKey: new(fmt.Sprintf("article reward:%d", articleID)),
	}
	_, err = u.dtm.InTCC(ctx, func(tcc *dtmgrpc.TccGrpc) error {
		if err = u.branches.Call(ctx, tcc, constant.ContentServiceName.String(), reward, contentv1.TCCArticleRewardService_Try_FullMethodName, contentv1.TCCArticleRewardService_Confirm_FullMethodName, contentv1.TCCArticleRewardService_Cancel_FullMethodName, &contentv1.TCCArticleReward_Resp{}); err != nil {
			return err
		}
		return u.branches.Call(ctx, tcc, constant.EconomyServiceName.String(), transfer, economyv1.TCCPointsTransferService_Try_FullMethodName, economyv1.TCCPointsTransferService_Confirm_FullMethodName, economyv1.TCCPointsTransferService_Cancel_FullMethodName, &economyv1.TCCPointsTransfer_Resp{})
	})
	return err == nil, err
}
