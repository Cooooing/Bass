package service

import (
	"context"

	v1 "common/proto/gen/user/v1"
	"user/internal/biz/model"
	"user/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type PrivacySettingService struct {
	v1.UnimplementedPrivacySettingServiceServer
	privacySettingUsecase *usecase.PrivacySettingUsecase
}

func NewPrivacySettingService(
	privacySettingUsecase *usecase.PrivacySettingUsecase,
) *PrivacySettingService {
	return &PrivacySettingService{
		privacySettingUsecase: privacySettingUsecase,
	}
}

func (s *PrivacySettingService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterPrivacySettingServiceServer(gs, s)
}

func (s *PrivacySettingService) RegisterHttp(hs *http.Server) {
}

func (s *PrivacySettingService) Get(ctx context.Context, req *v1.GetPrivacySetting_Req) (*v1.GetPrivacySetting_Resp, error) {
	res, err := s.privacySettingUsecase.GetByUserID(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}
	if res == nil {
		res = &model.PrivacySetting{UserID: req.GetUserId()}
	}
	return &v1.GetPrivacySetting_Resp{PrivacySetting: s.privacySettingReply(res)}, nil
}

func (s *PrivacySettingService) Map(ctx context.Context, req *v1.MapPrivacySettings_Req) (*v1.MapPrivacySettings_Resp, error) {
	settings, err := s.privacySettingUsecase.MapByUserIDs(ctx, req.GetUserIds())
	if err != nil {
		return nil, err
	}
	reply := make(map[int64]*v1.PrivacySetting, len(settings))
	for userID, setting := range settings {
		reply[userID] = s.privacySettingReply(setting)
	}
	return &v1.MapPrivacySettings_Resp{PrivacySettings: reply}, nil
}

func (s *PrivacySettingService) Update(ctx context.Context, req *v1.UpdatePrivacySetting_Req) (*v1.UpdatePrivacySetting_Resp, error) {
	res, err := s.privacySettingUsecase.UpsertByUserID(ctx, &model.PrivacySetting{
		UserID:               req.GetUserId(),
		PublicPoints:         req.PublicPoints,
		PublicFollowerList:   req.PublicFollowerList,
		PublicFollowingList:  req.PublicFollowingList,
		PublicArticleList:    req.PublicArticleList,
		PublicCommentList:    req.PublicCommentList,
		PublicOnlineStatus:   req.PublicOnlineStatus,
		PublicLocation:       req.PublicLocation,
		PublicMoonbreezeList: req.PublicMoonbreezeList,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UpdatePrivacySetting_Resp{PrivacySetting: s.privacySettingReply(res)}, nil
}

func (*PrivacySettingService) privacySettingReply(setting *model.PrivacySetting) *v1.PrivacySetting {
	if setting == nil {
		return nil
	}
	return &v1.PrivacySetting{
		UserId:               setting.UserID,
		PublicPoints:         setting.PublicPoints,
		PublicFollowerList:   setting.PublicFollowerList,
		PublicFollowingList:  setting.PublicFollowingList,
		PublicArticleList:    setting.PublicArticleList,
		PublicCommentList:    setting.PublicCommentList,
		PublicOnlineStatus:   setting.PublicOnlineStatus,
		PublicLocation:       setting.PublicLocation,
		PublicMoonbreezeList: setting.PublicMoonbreezeList,
	}
}
