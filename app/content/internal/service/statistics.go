package service

import (
	"context"

	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/content/v1"

	"content/internal/biz/usecase"
	"content/internal/enum"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

// StatisticsService 适配内容领域内部批量统计协议。
type StatisticsService struct {
	v1.UnimplementedContentStatisticsServiceServer
	usecase *usecase.StatisticsUsecase
}

// NewStatisticsService 创建内容统计服务。
func NewStatisticsService(usecase *usecase.StatisticsUsecase) *StatisticsService {
	return &StatisticsService{usecase: usecase}
}
func (s *StatisticsService) RegisterGrpc(server *grpc.Server) {
	v1.RegisterContentStatisticsServiceServer(server, s)
}
func (*StatisticsService) RegisterHTTP(*http.Server) {}

func (s *StatisticsService) BatchCountUserBehaviors(ctx context.Context, req *v1.BatchCountContentUserBehaviors_Req) (*v1.BatchCountContentUserBehaviors_Resp, error) {
	statuses := make([]enum.ContentActionStatus, 0, len(req.GetActionStatuses()))
	for _, status := range req.GetActionStatuses() {
		value, ok := enum.ContentActionStatusMap.ToEnum(status)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		statuses = append(statuses, value)
	}
	rows, err := s.usecase.BatchCountUserBehaviors(ctx, req.GetUserIds(), req.GetStart().AsTime(), req.GetEnd().AsTime(), statuses)
	if err != nil {
		return nil, err
	}
	response := make([]*v1.BatchCountContentUserBehaviors_UserBehaviorCount, 0, len(rows))
	for _, row := range rows {
		response = append(response, &v1.BatchCountContentUserBehaviors_UserBehaviorCount{UserId: row.UserID, ArticleCount: row.ArticleCount, CommentCount: row.CommentCount, ArticleLikeCount: row.ArticleLikeCount, CommentLikeCount: row.CommentLikeCount, ArticleThankCount: row.ArticleThankCount, CommentThankCount: row.CommentThankCount, ArticleRewardCount: row.ArticleRewardCount, ViewCount: row.ViewCount})
	}
	return &v1.BatchCountContentUserBehaviors_Resp{Rows: response}, nil
}
