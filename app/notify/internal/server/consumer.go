package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"slices"

	nats "common/pkg/client/nats"
	commonenum "common/pkg/enum"
	event "common/proto/gen/common/enums/event"

	"notify/internal/biz/usecase"

	"github.com/nats-io/nats.go/jetstream"
	"google.golang.org/protobuf/encoding/protojson"
)

type Consumer struct {
	log             *slog.Logger
	natsClient      *nats.Client
	consumerUsecase *usecase.ConsumerUsecase
	ctx             context.Context
	cancel          context.CancelFunc
	subscriptions   []*nats.Subscription
}

func NewConsumer(
	logger *slog.Logger,
	natsClient *nats.Client,
	consumerUsecase *usecase.ConsumerUsecase,
) *Consumer {
	return &Consumer{
		log:             logger,
		natsClient:      natsClient,
		consumerUsecase: consumerUsecase,
	}
}

func (c *Consumer) Start(ctx context.Context) error {
	c.ctx, c.cancel = context.WithCancel(ctx)
	for _, group := range c.eventGroups() {
		stream := commonenum.EventStreamConfigFor(group.stream)
		if err := c.natsClient.EnsureStream(c.ctx, nats.StreamSpec{
			Name:              stream.Name.String(),
			Subjects:          stream.Subjects,
			Retention:         stream.Retention,
			AllowMsgSchedules: stream.AllowMsgSchedules,
			AllowMsgTTL:       stream.AllowMsgTTL,
			AllowRollup:       stream.AllowRollup,
		}); err != nil {
			return errors.Join(err, c.stopSubscriptions())
		}
		spec := nats.ConsumerSpec{
			Stream:         string(group.stream),
			Durable:        commonenum.EventConsumerNotify.String(),
			FilterSubjects: group.consumerSubjects,
			DeliverPolicy:  jetstream.DeliverAllPolicy,
		}
		subscription, consumeErr := c.natsClient.Consume(c.ctx, spec, c.handle)
		if consumeErr != nil {
			c.log.Error("consume failed", "stream", group.stream, "durable", spec.Durable, "error", consumeErr)
			return errors.Join(consumeErr, c.stopSubscriptions())
		}
		c.subscriptions = append(c.subscriptions, subscription)
	}
	return nil
}

type eventGroup struct {
	stream           commonenum.EventStream
	consumerSubjects []string
}

func (*Consumer) eventGroups() []eventGroup {
	return []eventGroup{
		{
			stream: commonenum.EventStreamUser,
			consumerSubjects: []string{
				commonenum.EventSubjectUserRegister.String(), commonenum.EventSubjectUserFollow.String(),
				commonenum.EventSubjectUserEmailVerificationCode.String(), commonenum.EventSubjectUserPhoneVerificationCode.String(),
			},
		},
		{
			stream: commonenum.EventStreamContent,
			consumerSubjects: []string{
				commonenum.EventSubjectContentArticlePublish.String(), commonenum.EventSubjectContentArticleLike.String(),
				commonenum.EventSubjectContentArticleThank.String(), commonenum.EventSubjectContentArticleReward.String(),
				commonenum.EventSubjectContentArticleCollect.String(), commonenum.EventSubjectContentArticleWatch.String(),
				commonenum.EventSubjectContentArticleAcceptAnswer.String(), commonenum.EventSubjectContentArticlePostscriptAdd.String(),
				commonenum.EventSubjectContentArticleLotteryWon.String(), commonenum.EventSubjectContentArticleLotteryDrawn.String(),
				commonenum.EventSubjectContentArticlePollEnded.String(), commonenum.EventSubjectContentCommentPublish.String(),
				commonenum.EventSubjectContentCommentLike.String(), commonenum.EventSubjectContentCommentThank.String(),
			},
		},
		{
			stream:           commonenum.EventStreamIM,
			consumerSubjects: []string{commonenum.EventSubjectIMGlobalChatMessageCreated.String()},
		},
	}
}

// handle 校验事件契约并交给业务层处理。
func (c *Consumer) handle(ctx context.Context, msg *nats.Message) error {
	if msg == nil {
		return nil
	}
	payload := msg.Data
	var message event.Event
	if err := protojson.Unmarshal(payload, &message); err != nil {
		c.log.Error("unmarshal event failed", "subject", msg.Subject, "error", err)
		return err
	}
	if message.EventId == "" {
		c.log.Error("event id is empty", "subject", msg.Subject)
		return errors.New("event id is required")
	}
	eventSubject := commonenum.EventSubject(msg.Subject)
	if _, ok := commonenum.EventSubjectMap.ToProto(eventSubject); !ok {
		err := fmt.Errorf("unknown event subject: %s", msg.Subject)
		c.log.Error("event subject invalid", "event_id", message.EventId, "subject", msg.Subject, "error", err)
		return err
	}
	if expectedSubject, ok := commonenum.EventSubjectMap.ToEnum(message.Subject); !ok || expectedSubject != eventSubject {
		err := fmt.Errorf("event subject mismatch: subject=%s event_subject=%s", msg.Subject, message.Subject.String())
		c.log.Error("event subject mismatch", "event_id", message.EventId, "subject", msg.Subject, "error", err)
		return err
	}
	err := c.consumerUsecase.Handle(ctx, &usecase.ConsumerHandleReq{
		Event:   &message,
		Subject: eventSubject,
		Payload: string(payload),
	})
	if err != nil {
		c.log.Error("notify event handling failed", "event_id", message.EventId, "subject", msg.Subject, "error", err)
	}
	return err
}

func (c *Consumer) Stop(ctx context.Context) error {
	if c.cancel == nil {
		return nil
	}
	c.cancel()
	return c.stopSubscriptions()
}

func (c *Consumer) stopSubscriptions() error {
	var stopErr error
	for _, subscription := range slices.Backward(c.subscriptions) {
		if err := subscription.Unsubscribe(); err != nil && stopErr == nil {
			stopErr = err
		}
	}
	c.subscriptions = nil
	return stopErr
}
