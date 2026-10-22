package service

import (
	"context"

	"common/pkg/apperror"
	"common/pkg/constant"
	commonmodel "common/pkg/model"
	"common/pkg/util"
	bffbbsv1 "common/proto/gen/bff_bbs/v1"
	cerrors "common/proto/gen/common/errors"
	imv1 "common/proto/gen/im/v1"

	"bff_bbs/internal/biz/model"
	"bff_bbs/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type IMService struct {
	bffbbsv1.UnimplementedIMServiceServer
	im *usecase.IMUsecase
}

func NewIMService(im *usecase.IMUsecase) *IMService {
	return &IMService{im: im}
}

func (*IMService) RegisterGrpc(*grpc.Server) {}

func (s *IMService) RegisterHTTP(server *http.Server) {
	bffbbsv1.RegisterIMServiceHTTPServer(server, s)
}

func (s *IMService) ListPrivateChatSessions(ctx context.Context, req *bffbbsv1.ListPrivateChatSessions_Req) (*bffbbsv1.ListPrivateChatSessions_Resp, error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := s.im.ListPrivateChatSessions(ctx, &imv1.ListPrivateChatSessions_Req{UserId: user.ID, Page: req.GetPage()})
	if err != nil {
		return nil, err
	}
	rows := make([]*bffbbsv1.PrivateChatSession, 0, len(resp.GetRows()))
	userIDs := make([]int64, 0, len(resp.GetRows()))
	for _, row := range resp.GetRows() {
		if row.GetPeerUserId() > 0 {
			userIDs = append(userIDs, row.GetPeerUserId())
		}
	}
	profiles, err := s.im.MapChatProfiles(ctx, userIDs)
	if err != nil {
		return nil, err
	}
	for _, row := range resp.GetRows() {
		rows = append(rows, s.privateChatSession(row, profiles[row.GetPeerUserId()]))
	}
	return &bffbbsv1.ListPrivateChatSessions_Resp{Page: resp.GetPage(), Rows: rows}, nil
}

func (s *IMService) ListPrivateChatMessages(ctx context.Context, req *bffbbsv1.ListPrivateChatMessages_Req) (*bffbbsv1.ListPrivateChatMessages_Resp, error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := s.im.ListPrivateChatMessages(ctx, &imv1.ListPrivateChatMessages_Req{UserId: user.ID, PeerUserId: req.GetPeerUserId(), BeforeId: req.GetBeforeId(), Page: req.GetPage()})
	if err != nil {
		return nil, err
	}
	rows := make([]*bffbbsv1.PrivateChatMessage, 0, len(resp.GetRows()))
	for _, row := range resp.GetRows() {
		rows = append(rows, s.privateChatMessage(row))
	}
	return &bffbbsv1.ListPrivateChatMessages_Resp{Page: resp.GetPage(), Rows: rows}, nil
}

func (s *IMService) CountPrivateChatUnread(ctx context.Context, _ *bffbbsv1.CountPrivateChatUnread_Req) (*bffbbsv1.CountPrivateChatUnread_Resp, error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	count, err := s.im.CountPrivateChatUnread(ctx, user.ID)
	if err != nil {
		return nil, err
	}
	return &bffbbsv1.CountPrivateChatUnread_Resp{Count: count}, nil
}

func (s *IMService) SendPrivateChatMessage(ctx context.Context, req *bffbbsv1.SendPrivateChatMessage_Req) (*bffbbsv1.SendPrivateChatMessage_Resp, error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := s.im.SendPrivateChatMessage(ctx, &imv1.SendPrivateChatMessage_Req{SenderId: user.ID, ReceiverId: req.GetReceiverId(), ClientMessageId: req.GetClientMessageId(), Content: req.GetContent()})
	if err != nil {
		return nil, err
	}
	return &bffbbsv1.SendPrivateChatMessage_Resp{Message: s.privateChatMessage(resp.GetMessage())}, nil
}

func (s *IMService) RevokePrivateChatMessage(ctx context.Context, req *bffbbsv1.RevokePrivateChatMessage_Req) (*bffbbsv1.RevokePrivateChatMessage_Resp, error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := s.im.RevokePrivateChatMessage(ctx, &imv1.RevokePrivateChatMessage_Req{UserId: user.ID, MessageId: req.GetMessageId()})
	if err != nil {
		return nil, err
	}
	return &bffbbsv1.RevokePrivateChatMessage_Resp{Message: s.privateChatMessage(resp.GetMessage())}, nil
}

func (s *IMService) MarkPrivateChatSessionRead(ctx context.Context, req *bffbbsv1.MarkPrivateChatSessionRead_Req) (*bffbbsv1.MarkPrivateChatSessionRead_Resp, error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	err = s.im.MarkPrivateChatSessionRead(ctx, &imv1.MarkPrivateChatSessionRead_Req{UserId: user.ID, PeerUserId: req.GetPeerUserId(), MessageId: req.GetMessageId()})
	return &bffbbsv1.MarkPrivateChatSessionRead_Resp{}, err
}

func (s *IMService) SetPrivateChatSessionMuted(ctx context.Context, req *bffbbsv1.SetPrivateChatSessionMuted_Req) (*bffbbsv1.SetPrivateChatSessionMuted_Resp, error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	err = s.im.SetPrivateChatSessionMuted(ctx, &imv1.SetPrivateChatSessionMuted_Req{UserId: user.ID, PeerUserId: req.GetPeerUserId(), Muted: req.GetMuted()})
	return &bffbbsv1.SetPrivateChatSessionMuted_Resp{}, err
}

func (s *IMService) SetPrivateChatSessionPinned(ctx context.Context, req *bffbbsv1.SetPrivateChatSessionPinned_Req) (*bffbbsv1.SetPrivateChatSessionPinned_Resp, error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	err = s.im.SetPrivateChatSessionPinned(ctx, &imv1.SetPrivateChatSessionPinned_Req{UserId: user.ID, PeerUserId: req.GetPeerUserId(), Pinned: req.GetPinned()})
	return &bffbbsv1.SetPrivateChatSessionPinned_Resp{}, err
}

func (s *IMService) ListGlobalChatMessages(ctx context.Context, req *bffbbsv1.ListGlobalChatMessages_Req) (*bffbbsv1.ListGlobalChatMessages_Resp, error) {
	resp, err := s.im.ListGlobalChatMessages(ctx, &imv1.ListGlobalChatMessages_Req{BeforeId: req.GetBeforeId(), Page: req.GetPage()})
	if err != nil {
		return nil, err
	}
	userIDs := make([]int64, 0, len(resp.GetRows()))
	for _, row := range resp.GetRows() {
		if row.GetSenderId() > 0 {
			userIDs = append(userIDs, row.GetSenderId())
		}
	}
	profiles, err := s.im.MapChatProfiles(ctx, userIDs)
	if err != nil {
		return nil, err
	}
	rows := make([]*bffbbsv1.GlobalChatMessage, 0, len(resp.GetRows()))
	for _, row := range resp.GetRows() {
		rows = append(rows, s.globalChatMessage(row, profiles[row.GetSenderId()]))
	}
	return &bffbbsv1.ListGlobalChatMessages_Resp{Page: resp.GetPage(), Rows: rows}, nil
}

func (s *IMService) SendGlobalChatMessage(ctx context.Context, req *bffbbsv1.SendGlobalChatMessage_Req) (*bffbbsv1.SendGlobalChatMessage_Resp, error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := s.im.SendGlobalChatMessage(ctx, &imv1.SendGlobalChatMessage_Req{SenderId: user.ID, ClientMessageId: req.GetClientMessageId(), Content: req.GetContent()})
	if err != nil {
		return nil, err
	}
	profiles, err := s.im.MapChatProfiles(ctx, []int64{resp.GetMessage().GetSenderId()})
	if err != nil {
		return nil, err
	}
	return &bffbbsv1.SendGlobalChatMessage_Resp{Message: s.globalChatMessage(resp.GetMessage(), profiles[resp.GetMessage().GetSenderId()])}, nil
}

func (s *IMService) RevokeGlobalChatMessage(ctx context.Context, req *bffbbsv1.RevokeGlobalChatMessage_Req) (*bffbbsv1.RevokeGlobalChatMessage_Resp, error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := s.im.RevokeGlobalChatMessage(ctx, &imv1.RevokeGlobalChatMessage_Req{UserId: user.ID, MessageId: req.GetMessageId()})
	if err != nil {
		return nil, err
	}
	profiles, err := s.im.MapChatProfiles(ctx, []int64{resp.GetMessage().GetSenderId()})
	if err != nil {
		return nil, err
	}
	return &bffbbsv1.RevokeGlobalChatMessage_Resp{Message: s.globalChatMessage(resp.GetMessage(), profiles[resp.GetMessage().GetSenderId()])}, nil
}

func (s *IMService) ListGlobalChatOnlineMembers(ctx context.Context, req *bffbbsv1.ListGlobalChatOnlineMembers_Req) (*bffbbsv1.ListGlobalChatOnlineMembers_Resp, error) {
	profiles, err := s.im.ListGlobalChatOnlineMembers(ctx)
	if err != nil {
		return nil, err
	}
	rows := make([]*bffbbsv1.GlobalChatOnlineMember, 0, len(profiles))
	for _, profile := range profiles {
		if profile == nil {
			continue
		}
		rows = append(rows, &bffbbsv1.GlobalChatOnlineMember{UserId: profile.ID, Name: profile.Name, Nickname: profile.Nickname, AvatarUrl: profile.AvatarURL})
	}
	return &bffbbsv1.ListGlobalChatOnlineMembers_Resp{Rows: rows}, nil
}

func (*IMService) currentUser(ctx context.Context) (*commonmodel.User, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil || user.ID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	return user, nil
}

func (*IMService) privateChatMessage(row *imv1.PrivateChatMessage) *bffbbsv1.PrivateChatMessage {
	if row == nil {
		return nil
	}
	return &bffbbsv1.PrivateChatMessage{Id: row.GetId(), SenderId: row.GetSenderId(), ReceiverId: row.GetReceiverId(), Content: row.GetContent(), MessageType: row.GetMessageType(), Revoked: row.GetRevoked(), CreatedAt: row.GetCreatedAt(), RevokedAt: row.GetRevokedAt()}
}

func (s *IMService) privateChatSession(row *imv1.PrivateChatSession, profile *model.AccountProfile) *bffbbsv1.PrivateChatSession {
	if row == nil {
		return nil
	}
	session := &bffbbsv1.PrivateChatSession{PeerUserId: row.GetPeerUserId(), LastReadMessageId: row.GetLastReadMessageId(), UnreadCount: row.GetUnreadCount(), Muted: row.GetMuted(), Pinned: row.GetPinned(), LastMessage: s.privateChatMessage(row.GetLastMessage()), LastMessageAt: row.GetLastMessageAt()}
	if profile != nil {
		session.Peer = &bffbbsv1.PrivateChatMember{UserId: profile.ID, Name: profile.Name, Nickname: profile.Nickname, AvatarUrl: profile.AvatarURL}
	}
	return session
}

func (*IMService) globalChatMessage(row *imv1.GlobalChatMessage, profile *model.AccountProfile) *bffbbsv1.GlobalChatMessage {
	if row == nil {
		return nil
	}
	message := &bffbbsv1.GlobalChatMessage{Id: row.GetId(), SenderId: row.GetSenderId(), Content: row.GetContent(), MessageType: row.GetMessageType(), Revoked: row.GetRevoked(), CreatedAt: row.GetCreatedAt(), RevokedAt: row.GetRevokedAt()}
	if profile != nil {
		message.Sender = &bffbbsv1.GlobalChatMember{UserId: profile.ID, Name: profile.Name, Nickname: profile.Nickname, AvatarUrl: profile.AvatarURL}
	}
	return message
}
