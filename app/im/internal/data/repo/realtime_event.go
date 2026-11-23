package repo

import (
	"context"
	"fmt"

	nats "common/pkg/client/nats"
	commonenum "common/pkg/enum"
	commonevent "common/proto/gen/common/enums/event/v1"

	"im/internal/biz/model"
	bizrepo "im/internal/biz/repo"
	"im/internal/enum"

	"github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type RealtimeEventPublisher struct {
	nats *nats.Client
}

func (p *RealtimeEventPublisher) Ensure(ctx context.Context) error {
	stream := commonenum.EventStreamConfigFor(commonenum.EventStreamIM)
	return p.nats.EnsureStream(ctx, nats.StreamSpec{
		Name:              stream.Name.String(),
		Subjects:          stream.Subjects,
		Retention:         stream.Retention,
		AllowMsgSchedules: stream.AllowMsgSchedules,
		AllowMsgTTL:       stream.AllowMsgTTL,
		AllowRollup:       stream.AllowRollup,
	})
}

func NewRealtimeEventPublisher(nats *nats.Client) bizrepo.RealtimeEventPublisher {
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
	event := &commonevent.Event{EventId: uuid.NewString(), Subject: commonevent.EventSubject_EVENT_SUBJECT_IM_PRIVATE_CHAT_MESSAGE_CREATED, Timestamp: timestamppb.Now(), Payload: &commonevent.Event_ImPrivateChatMessageCreated{ImPrivateChatMessageCreated: &commonevent.IMPrivateChatMessageCreatedPayload{MessageId: message.ID, SenderId: message.SenderID, ReceiverId: message.ReceiverID, MessageType: messageType, Content: message.Content, CreatedAt: timestamppb.New(*message.CreatedAt)}}}
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
	event := &commonevent.Event{EventId: uuid.NewString(), Subject: commonevent.EventSubject_EVENT_SUBJECT_IM_GLOBAL_CHAT_MESSAGE_CREATED, Timestamp: timestamppb.Now(), Payload: &commonevent.Event_ImGlobalChatMessageCreated{ImGlobalChatMessageCreated: &commonevent.IMGlobalChatMessageCreatedPayload{MessageId: message.ID, SenderId: message.SenderID, MessageType: messageType, Content: message.Content, CreatedAt: timestamppb.New(*message.CreatedAt)}}}
	payload, err := protojson.Marshal(event)
	if err != nil {
		return fmt.Errorf("marshal global chat event: %w", err)
	}
	_, err = p.nats.JSPublish(ctx, commonenum.EventSubjectIMGlobalChatMessageCreated.String(), &nats.Message{ID: event.EventId, Data: payload})
	return err
}

func (p *RealtimeEventPublisher) PublishGlobalChatBarrageCreated(ctx context.Context, barrage *model.GlobalChatBarrage) error {
	if barrage == nil {
		return fmt.Errorf("global chat barrage is required")
	}
	if barrage.CreatedAt == nil {
		return fmt.Errorf("global chat barrage created at is required")
	}
	event := &commonevent.Event{
		EventId:   uuid.NewString(),
		Subject:   commonevent.EventSubject_EVENT_SUBJECT_IM_GLOBAL_CHAT_BARRAGE_CREATED,
		Timestamp: timestamppb.Now(),
		Payload: &commonevent.Event_ImGlobalChatBarrageCreated{
			ImGlobalChatBarrageCreated: &commonevent.IMGlobalChatBarrageCreatedPayload{
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
	event := &commonevent.Event{
		EventId:   uuid.NewString(),
		Subject:   commonevent.EventSubject_EVENT_SUBJECT_IM_GLOBAL_CHAT_TOPIC_UPDATED,
		Timestamp: timestamppb.Now(),
		Payload: &commonevent.Event_ImGlobalChatTopicUpdated{
			ImGlobalChatTopicUpdated: &commonevent.IMGlobalChatTopicUpdatedPayload{
				TopicId:   topic.ID,
				UserId:    topic.UserID,
				Content:   topic.Content,
				CreatedAt: timestamppb.New(*topic.CreatedAt),
			},
		},
	}
	return p.publish(ctx, commonenum.EventSubjectIMGlobalChatTopicUpdated, event)
}

func (p *RealtimeEventPublisher) publish(ctx context.Context, subject commonenum.EventSubject, event *commonevent.Event) error {
	payload, err := protojson.Marshal(event)
	if err != nil {
		return fmt.Errorf("marshal realtime event: %w", err)
	}
	return p.nats.Publish(ctx, subject.String(), &nats.Message{Data: payload})
}
