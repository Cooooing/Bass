package service

import (
	"context"

	imv1 "common/proto/gen/im/v1"

	"im/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

// StatisticsService 适配即时消息领域内部批量统计协议。
type StatisticsService struct {
	imv1.UnimplementedIMStatisticsServiceServer
	usecase *usecase.StatisticsUsecase
}

// NewStatisticsService 创建即时消息统计服务。
func NewStatisticsService(usecase *usecase.StatisticsUsecase) *StatisticsService {
	return &StatisticsService{usecase: usecase}
}
func (s *StatisticsService) RegisterGrpc(server *grpc.Server) {
	imv1.RegisterIMStatisticsServiceServer(server, s)
}
func (*StatisticsService) RegisterHTTP(*http.Server) {}
func (s *StatisticsService) BatchCountGlobalChatMessages(ctx context.Context, req *imv1.BatchCountGlobalChatMessages_Req) (*imv1.BatchCountGlobalChatMessages_Resp, error) {
	rows, err := s.usecase.BatchCountGlobalChatMessages(ctx, req.GetUserIds(), req.GetStart().AsTime(), req.GetEnd().AsTime())
	if err != nil {
		return nil, err
	}
	response := make([]*imv1.BatchCountGlobalChatMessages_Resp_UserMessageCount, 0, len(rows))
	for _, row := range rows {
		response = append(response, &imv1.BatchCountGlobalChatMessages_Resp_UserMessageCount{UserId: row.UserID, Count: row.Count})
	}
	return &imv1.BatchCountGlobalChatMessages_Resp{Rows: response}, nil
}
