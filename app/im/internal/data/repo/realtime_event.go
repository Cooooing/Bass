package repo

import (
	"context"
	"fmt"

	"common/pkg/client"
	commonenum "common/pkg/enum"
	commonenums "common/proto/gen/common/enums"

	"im/internal/biz/model"
	bizrepo "im/internal/biz/repo"
	"im/internal/enum"

	"github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type RealtimeEventPublisher struct {
	nats *client.NatsClient
}

func NewRealtimeEventPublisher(nats *client.NatsClient) bizrepo.RealtimeEventPublisher {
	return &RealtimeEventPublisher{nats: nats}
}

func (p *RealtimeEventPublisher) PublishPrivateChatMessageCreated(ctx context.Context, message *model.PrivateChatMessage) error {
	if message == nil {
		return fmt.Errorf("private chat message is required")
	}
	messageType, ok := enum.MessageTypeMap.ToProto(message.Type)
	if !ok {
		return fmt.Errorf("private chat message type is invalid")
	}
	if message.CreatedAt == nil {
		return fmt.Errorf("private chat message created at is required")
	}
	event := &commonenums.Event{EventId: uuid.NewString(), Type: commonenums.EventType_EVENT_TYPE_IM_PRIVATE_CHAT_MESSAGE_CREATED, Subject: commonenums.EventSubject_EVENT_SUBJECT_IM_PRIVATE_CHAT_MESSAGE_CREATED, Timestamp: timestamppb.Now(), Payload: &commonenums.Event_ImPrivateChatMessageCreated{ImPrivateChatMessageCreated: &commonenums.IMPrivateChatMessageCreatedPayload{MessageId: message.ID, SenderId: message.SenderID, ReceiverId: message.ReceiverID, MessageType: messageType, Content: message.Content, CreatedAt: timestamppb.New(*message.CreatedAt)}}}
	return p.publish(ctx, commonenum.EventSubjectIMPrivateChatMessageCreated, event)
}

func (p *RealtimeEventPublisher) PublishGlobalChatMessageCreated(ctx context.Context, message *model.GlobalChatMessage) error {
	if message == nil {
		return fmt.Errorf("global chat message is required")
	}
	messageType, ok := enum.MessageTypeMap.ToProto(message.Type)
	if !ok {
		return fmt.Errorf("global chat message type is invalid")
	}
	if message.CreatedAt == nil {
		return fmt.Errorf("global chat message created at is required")
	}
	event := &commonenums.Event{EventId: uuid.NewString(), Type: commonenums.EventType_EVENT_TYPE_IM_GLOBAL_CHAT_MESSAGE_CREATED, Subject: commonenums.EventSubject_EVENT_SUBJECT_IM_GLOBAL_CHAT_MESSAGE_CREATED, Timestamp: timestamppb.Now(), Payload: &commonenums.Event_ImGlobalChatMessageCreated{ImGlobalChatMessageCreated: &commonenums.IMGlobalChatMessageCreatedPayload{MessageId: message.ID, SenderId: message.SenderID, MessageType: messageType, Content: message.Content, CreatedAt: timestamppb.New(*message.CreatedAt)}}}
	return p.publish(ctx, commonenum.EventSubjectIMGlobalChatMessageCreated, event)
}

func (p *RealtimeEventPublisher) PublishGlobalChatBarrageCreated(ctx context.Context, barrage *model.GlobalChatBarrage) error {
	if barrage == nil {
		return fmt.Errorf("global chat barrage is required")
	}
	if barrage.CreatedAt == nil {
		return fmt.Errorf("global chat barrage created at is required")
	}
	event := &commonenums.Event{
		EventId:   uuid.NewString(),
		Type:      commonenums.EventType_EVENT_TYPE_IM_GLOBAL_CHAT_BARRAGE_CREATED,
		Subject:   commonenums.EventSubject_EVENT_SUBJECT_IM_GLOBAL_CHAT_BARRAGE_CREATED,
		Timestamp: timestamppb.Now(),
		Payload: &commonenums.Event_ImGlobalChatBarrageCreated{
			ImGlobalChatBarrageCreated: &commonenums.IMGlobalChatBarrageCreatedPayload{
				BarrageId: barrage.ID,
				UserId:    barrage.UserID,
				Content:   barrage.Content,
				CreatedAt: timestamppb.New(*barrage.CreatedAt),
			},
		},
	}
	return p.publish(ctx, commonenum.EventSubjectIMGlobalChatBarrageCreated, event)
}

func (p *RealtimeEventPublisher) PublishGlobalChatTopicUpdated(ctx context.Context, topic *model.GlobalChatTopic) error {
	if topic == nil || topic.CreatedAt == nil {
		return fmt.Errorf("global chat topic is required")
	}
	event := &commonenums.Event{
		EventId:   uuid.NewString(),
		Type:      commonenums.EventType_EVENT_TYPE_IM_GLOBAL_CHAT_TOPIC_UPDATED,
		Subject:   commonenums.EventSubject_EVENT_SUBJECT_IM_GLOBAL_CHAT_TOPIC_UPDATED,
		Timestamp: timestamppb.Now(),
		Payload: &commonenums.Event_ImGlobalChatTopicUpdated{
			ImGlobalChatTopicUpdated: &commonenums.IMGlobalChatTopicUpdatedPayload{
				TopicId:   topic.ID,
				UserId:    topic.UserID,
				Content:   topic.Content,
				CreatedAt: timestamppb.New(*topic.CreatedAt),
			},
		},
	}
	return p.publish(ctx, commonenum.EventSubjectIMGlobalChatTopicUpdated, event)
}

func (p *RealtimeEventPublisher) publish(ctx context.Context, subject commonenum.EventSubject, event *commonenums.Event) error {
	payload, err := protojson.Marshal(event)
	if err != nil {
		return fmt.Errorf("marshal realtime event: %w", err)
	}
	return p.nats.Publish(ctx, subject.String(), &client.Message{Data: payload})
}
