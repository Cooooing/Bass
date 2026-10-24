package service

import (
	"context"
	"errors"
	"fmt"
	"io"

	imv1enum "common/proto/gen/im/v1/enum"
	realtimehubv1 "common/proto/gen/realtime_hub/v1"
	"realtime_hub/internal/biz/model"
	"realtime_hub/internal/biz/usecase"
	"realtime_hub/internal/enum"

	kratosgrpc "github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type RealtimeService struct {
	realtimehubv1.UnimplementedRealtimeHubServiceServer
	realtime *usecase.RealtimeUsecase
}

func NewRealtimeService(realtime *usecase.RealtimeUsecase) *RealtimeService {
	return &RealtimeService{realtime: realtime}
}

func (s *RealtimeService) RegisterGrpc(server *kratosgrpc.Server) {
	realtimehubv1.RegisterRealtimeHubServiceServer(server, s)
}

func (*RealtimeService) RegisterHTTP(*http.Server) {}

func (s *RealtimeService) QueryOnline(ctx context.Context, req *realtimehubv1.QueryRealtimeOnline_Req) (*realtimehubv1.QueryRealtimeOnline_Resp, error) {
	onlineStatuses, err := s.realtime.QueryOnline(ctx, req.GetRealm(), req.GetUserIds())
	if err != nil {
		return nil, err
	}
	return &realtimehubv1.QueryRealtimeOnline_Resp{OnlineStatuses: onlineStatuses}, nil
}

func (s *RealtimeService) ListOnlineUsers(ctx context.Context, req *realtimehubv1.ListRealtimeOnlineUsers_Req) (*realtimehubv1.ListRealtimeOnlineUsers_Resp, error) {
	userIDs, err := s.realtime.ListOnlineUsers(ctx, req.GetRealm())
	if err != nil {
		return nil, err
	}
	return &realtimehubv1.ListRealtimeOnlineUsers_Resp{UserIds: userIDs}, nil
}

func (s *RealtimeService) Connect(stream grpc.BidiStreamingServer[realtimehubv1.NodeToHub, realtimehubv1.HubToNode]) error {
	first, err := stream.Recv()
	if err != nil {
		return err
	}
	register := first.GetRegisterNode()
	if register == nil {
		return fmt.Errorf("first realtime node message must register node")
	}
	group, ok := enum.NodeGroupMap.ToEnum(register.GetGroup())
	if !ok {
		return fmt.Errorf("node group is invalid")
	}
	messages, err := s.realtime.OpenSession(stream.Context(), register.GetNodeId(), group, register.GetRealms())
	if err != nil {
		return err
	}
	defer s.realtime.CloseSession(context.WithoutCancel(stream.Context()), register.GetNodeId(), register.GetRealms())
	errs := make(chan error, 1)
	go s.receiveNodeMessages(stream, register.GetNodeId(), group, register.GetRealms(), errs)
	for {
		select {
		case <-stream.Context().Done():
			return nil
		case err := <-errs:
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		case message := <-messages:
			response, err := s.hubMessageProto(message)
			if err != nil {
				return err
			}
			if err := stream.Send(response); err != nil {
				return err
			}
		}
	}
}

func (s *RealtimeService) receiveNodeMessages(stream grpc.BidiStreamingServer[realtimehubv1.NodeToHub, realtimehubv1.HubToNode], nodeID string, group enum.NodeGroup, realms []string, errs chan<- error) {
	for {
		message, err := stream.Recv()
		if err != nil {
			errs <- err
			return
		}
		if err := s.handleNodeMessage(stream.Context(), nodeID, group, realms, message); err != nil {
			errs <- err
			return
		}
	}
}

func (s *RealtimeService) handleNodeMessage(ctx context.Context, nodeID string, group enum.NodeGroup, realms []string, message *realtimehubv1.NodeToHub) error {
	if bind := message.GetBindConnection(); bind != nil {
		return s.realtime.BindConnection(ctx, nodeID, bind.GetRealm(), bind.GetUserId(), bind.GetConnectionId())
	}
	if unbind := message.GetUnbindConnection(); unbind != nil {
		return s.realtime.UnbindConnection(ctx, nodeID, unbind.GetRealm(), unbind.GetUserId(), unbind.GetConnectionId())
	}
	if message.GetHeartbeat() != nil {
		return s.realtime.RenewSession(ctx, nodeID, group, realms)
	}
	return fmt.Errorf("node message payload is invalid")
}

func (s *RealtimeService) hubMessageProto(message *model.NodeMessage) (*realtimehubv1.HubToNode, error) {
	if message == nil {
		return nil, fmt.Errorf("hub message is required")
	}
	if message.Delivery != nil {
		delivery, err := s.deliveryProto(message.Delivery)
		if err != nil {
			return nil, err
		}
		return &realtimehubv1.HubToNode{Payload: &realtimehubv1.HubToNode_Delivery{Delivery: delivery}}, nil
	}
	if message.CloseConnection != nil {
		return &realtimehubv1.HubToNode{Payload: &realtimehubv1.HubToNode_Control{Control: &realtimehubv1.Control{Payload: &realtimehubv1.Control_CloseConnection{CloseConnection: &realtimehubv1.CloseConnection{ConnectionId: message.CloseConnection.ConnectionID, Reason: message.CloseConnection.Reason}}}}}, nil
	}
	if message.ConnectionBound != nil {
		return &realtimehubv1.HubToNode{Payload: &realtimehubv1.HubToNode_Control{Control: &realtimehubv1.Control{Payload: &realtimehubv1.Control_ConnectionBound{ConnectionBound: &realtimehubv1.ConnectionBound{ConnectionId: message.ConnectionBound.ConnectionID}}}}}, nil
	}
	return nil, fmt.Errorf("hub message payload is invalid")
}

func (*RealtimeService) deliveryProto(delivery *model.Delivery) (*realtimehubv1.Delivery, error) {
	if delivery == nil {
		return nil, fmt.Errorf("delivery is required")
	}
	kind, ok := enum.DeliveryKindMap.ToProto(delivery.Kind)
	if !ok {
		return nil, fmt.Errorf("delivery kind is invalid")
	}
	response := &realtimehubv1.Delivery{UserId: delivery.UserID, Kind: kind}
	switch payload := delivery.Payload.(type) {
	case model.NotificationDelivery:
		response.Payload = &realtimehubv1.Delivery_Notification{Notification: &realtimehubv1.NotificationDelivery{NotificationId: payload.NotificationID}}
	case model.PrivateChatMessageDelivery:
		response.Payload = &realtimehubv1.Delivery_PrivateChatMessage{PrivateChatMessage: &realtimehubv1.PrivateChatMessageDelivery{MessageId: payload.MessageID, SenderId: payload.SenderID, ReceiverId: payload.ReceiverID, MessageType: imv1enum.IMMessageType(payload.MessageType), Content: payload.Content, CreatedAt: timestamppb.New(payload.CreatedAt)}}
	case model.GlobalChatMessageDelivery:
		response.Payload = &realtimehubv1.Delivery_GlobalChatMessage{GlobalChatMessage: &realtimehubv1.GlobalChatMessageDelivery{MessageId: payload.MessageID, SenderId: payload.SenderID, MessageType: imv1enum.IMMessageType(payload.MessageType), Content: payload.Content, CreatedAt: timestamppb.New(payload.CreatedAt)}}
	default:
		return nil, fmt.Errorf("delivery payload does not match kind")
	}
	return response, nil
}

func (s *RealtimeService) IssueTicket(ctx context.Context, req *realtimehubv1.IssueRealtimeTicket_Req) (*realtimehubv1.IssueRealtimeTicket_Resp, error) {
	ticket, err := s.realtime.IssueTicket(ctx, req.GetRealm(), req.GetUserId())
	if err != nil {
		return nil, err
	}
	return &realtimehubv1.IssueRealtimeTicket_Resp{Ticket: ticket, ExpiresInSeconds: 60}, nil
}

func (s *RealtimeService) ConsumeTicket(ctx context.Context, req *realtimehubv1.ConsumeRealtimeTicket_Req) (*realtimehubv1.ConsumeRealtimeTicket_Resp, error) {
	userID, err := s.realtime.ConsumeTicket(ctx, req.GetRealm(), req.GetTicket())
	if err != nil {
		return nil, err
	}
	return &realtimehubv1.ConsumeRealtimeTicket_Resp{UserId: userID}, nil
}
