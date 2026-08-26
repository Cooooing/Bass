package repo

import (
	"common/pkg/client"
	commonenum "common/pkg/enum"
	commonenums "common/proto/gen/common/enums"
	"context"
	"game_idle/internal/biz/model"
	bizrepo "game_idle/internal/biz/repo"
	"strconv"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ bizrepo.ChatMessageEventRepo = (*ChatMessageEventRepo)(nil)

// ChatMessageEventRepo 发布聊天消息事件。
type ChatMessageEventRepo struct {
	natsClient *client.NatsClient
}

func NewChatMessageEventRepo(natsClient *client.NatsClient) bizrepo.ChatMessageEventRepo {
	return &ChatMessageEventRepo{
		natsClient: natsClient,
	}
}

func (r *ChatMessageEventRepo) PublishWorldMessage(ctx context.Context, message *model.ChatMessage) error {
	createdAt := time.Now()
	if message.CreatedAt != nil {
		createdAt = *message.CreatedAt
	}
	eventID := uuid.NewString()
	event := &commonenums.Event{
		EventId:   eventID,
		Type:      commonenums.EventType_EVENT_TYPE_GAME_IDLE_CHAT_WORLD_MESSAGE,
		Subject:   commonenums.EventSubject_EVENT_SUBJECT_GAME_IDLE_CHAT_WORLD_MESSAGE,
		Timestamp: timestamppb.New(createdAt),
		Payload: &commonenums.Event_GameIdleChatWorldMessage{
			GameIdleChatWorldMessage: &commonenums.GameIdleChatWorldMessagePayload{
				MessageId:         message.ID,
				ChannelId:         message.ChannelID,
				SenderCharacterId: message.SenderCharacterID,
				Content:           message.Content,
			},
		},
	}
	payload, err := proto.Marshal(event)
	if err != nil {
		return err
	}
	return r.natsClient.Publish(
		ctx,
		commonenum.EventSubjectGameIdleChatWorldMessage.String(),
		&client.Message{
			Subject: commonenum.EventSubjectGameIdleChatWorldMessage.String(),
			Data:    payload,
			Header: map[string]string{
				"event_id":   eventID,
				"message_id": strconv.FormatInt(message.ID, 10),
			},
		},
	)
}
