package repo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/url"
	"time"

	commonclient "common/pkg/client"
	platformbizrepo "platform/internal/biz/repo"
	platformconfig "platform/internal/config"

	"github.com/nats-io/nats.go/jetstream"
)

const (
	defaultAssetEventStream          = "PLATFORM_ASSET_UPLOADS"
	defaultAssetEventSubject         = "storage.asset.object-created"
	defaultAssetEventConsumerDurable = "platform-asset-upload-recorder"
)

var _ platformbizrepo.AssetUploadEventConsumer = (*AssetEventNatsRepo)(nil)

// AssetEventNatsRepo owns the durable infrastructure boundary between MinIO
// and the platform asset directory. The stream uses WorkQueue retention: an
// event is retained while Platform is unavailable and removed after ACK.
type AssetEventNatsRepo struct {
	logger          *slog.Logger
	jetStream       jetstream.JetStream
	streamName      string
	subject         string
	consumerDurable string
	ackWait         time.Duration
	maxAckPending   int
	consumer        jetstream.Consumer
	consume         jetstream.ConsumeContext
}

func NewAssetEventNatsRepo(
	logger *slog.Logger,
	conf *platformconfig.Bootstrap,
	natsClient *commonclient.NatsClient,
) (platformbizrepo.AssetUploadEventConsumer, error) {
	if natsClient == nil || natsClient.Conn() == nil {
		return nil, errors.New("platform asset event consumer requires NATS")
	}
	js, err := jetstream.New(natsClient.Conn())
	if err != nil {
		return nil, err
	}
	settings := conf.GetPlatform().GetAssetEvents()
	streamName := settings.GetStreamName()
	if streamName == "" {
		streamName = defaultAssetEventStream
	}
	subject := settings.GetSubject()
	if subject == "" {
		subject = defaultAssetEventSubject
	}
	consumerDurable := settings.GetConsumerDurable()
	if consumerDurable == "" {
		consumerDurable = defaultAssetEventConsumerDurable
	}
	ackWait := time.Minute
	if settings.GetAckWait() != nil && settings.GetAckWait().AsDuration() > 0 {
		ackWait = settings.GetAckWait().AsDuration()
	}
	maxAckPending := int(settings.GetMaxAckPending())
	if maxAckPending <= 0 {
		maxAckPending = 32
	}
	return &AssetEventNatsRepo{
		logger:          logger,
		jetStream:       js,
		streamName:      streamName,
		subject:         subject,
		consumerDurable: consumerDurable,
		ackWait:         ackWait,
		maxAckPending:   maxAckPending,
	}, nil
}

func (r *AssetEventNatsRepo) ensure(ctx context.Context) error {
	if _, err := r.jetStream.CreateOrUpdateStream(ctx, jetstream.StreamConfig{
		Name:      r.streamName,
		Subjects:  []string{r.subject},
		Retention: jetstream.WorkQueuePolicy,
		Storage:   jetstream.FileStorage,
	}); err != nil {
		return err
	}
	consumer, err := r.jetStream.Consumer(ctx, r.streamName, r.consumerDurable)
	if errors.Is(err, jetstream.ErrConsumerNotFound) {
		consumer, err = r.jetStream.CreateConsumer(ctx, r.streamName, jetstream.ConsumerConfig{
			Durable:       r.consumerDurable,
			Name:          r.consumerDurable,
			DeliverPolicy: jetstream.DeliverAllPolicy,
			AckPolicy:     jetstream.AckExplicitPolicy,
			AckWait:       r.ackWait,
			FilterSubject: r.subject,
			MaxAckPending: r.maxAckPending,
		})
	}
	if err != nil {
		return err
	}
	r.consumer = consumer
	return nil
}

func (r *AssetEventNatsRepo) Consume(ctx context.Context, handler platformbizrepo.AssetUploadEventHandler) error {
	if handler == nil {
		return errors.New("platform asset event handler is required")
	}
	if r.consumer == nil {
		if err := r.ensure(ctx); err != nil {
			return err
		}
	}
	consume, err := r.consumer.Consume(func(msg jetstream.Msg) {
		r.consumeMessage(ctx, msg, handler)
	}, jetstream.PullMaxMessages(r.maxAckPending), jetstream.PullExpiry(2*time.Second))
	if err != nil {
		return err
	}
	r.consume = consume
	return nil
}

func (r *AssetEventNatsRepo) Stop(context.Context) error {
	if r.consume != nil {
		r.consume.Stop()
		r.consume = nil
	}
	return nil
}

// consumeMessage owns JetStream acknowledgement policy. Malformed provider
// payloads are terminal because a retry cannot repair them. Domain and storage
// errors are NAKed so JetStream redelivers them after the configured delay.
func (r *AssetEventNatsRepo) consumeMessage(
	ctx context.Context,
	msg jetstream.Msg,
	handler platformbizrepo.AssetUploadEventHandler,
) {
	events, err := decodeMinioAssetUploadEvents(msg.Data())
	if err != nil {
		r.logger.WarnContext(ctx, "discard invalid MinIO asset event", slog.Any("err", err))
		_ = msg.Term()
		return
	}
	for _, event := range events {
		if err := handler(ctx, event); err != nil {
			if errors.Is(err, platformbizrepo.ErrInvalidAssetUploadEvent) {
				r.logger.WarnContext(ctx, "discard invalid asset upload event", slog.Any("err", err))
				_ = msg.Term()
				return
			}
			r.logger.WarnContext(ctx, "asset upload event handling failed", slog.Any("err", err))
			_ = msg.Nak()
			return
		}
	}
	_ = msg.Ack()
}

func decodeMinioAssetUploadEvents(payload []byte) ([]*platformbizrepo.AssetUploadEvent, error) {
	var notification minioEventNotification
	if err := json.Unmarshal(payload, &notification); err != nil {
		return nil, fmt.Errorf("decode minio notification: %w", err)
	}
	events := make([]*platformbizrepo.AssetUploadEvent, 0, len(notification.Records))
	for _, record := range notification.Records {
		if record.EventName != "s3:ObjectCreated:Post" {
			continue
		}
		objectKey, err := url.QueryUnescape(record.S3.Object.Key)
		if err != nil {
			return nil, fmt.Errorf("decode minio object key: %w", err)
		}
		events = append(events, &platformbizrepo.AssetUploadEvent{
			Bucket:    record.S3.Bucket.Name,
			ObjectKey: objectKey,
		})
	}
	return events, nil
}

// minioEventNotification models only the stable portion of MinIO's S3 event
// payload used by the asset directory. Provider-specific fields stay local to
// this adapter and never cross the data-to-biz boundary.
type minioEventNotification struct {
	Records []minioEventRecord `json:"Records"`
}

type minioEventRecord struct {
	EventName string       `json:"eventName"`
	S3        minioS3Event `json:"s3"`
}

type minioS3Event struct {
	Bucket minioBucketEvent `json:"bucket"`
	Object minioObjectEvent `json:"object"`
}

type minioBucketEvent struct {
	Name string `json:"name"`
}

type minioObjectEvent struct {
	Key string `json:"key"`
}
