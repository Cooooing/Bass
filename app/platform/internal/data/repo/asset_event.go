package repo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/url"

	nats "common/pkg/client/nats"
	commonenum "common/pkg/enum"

	platformbizrepo "platform/internal/biz/repo"

	"github.com/nats-io/nats.go/jetstream"
)

var _ platformbizrepo.AssetUploadEventConsumer = (*AssetEventRepo)(nil)

type AssetEventRepo struct {
	logger          *slog.Logger
	natsClient      *nats.Client
	streamName      string
	subject         string
	consumerDurable string
	subscription    *nats.Subscription
}

func NewAssetEventRepo(
	logger *slog.Logger,
	natsClient *nats.Client,
) (platformbizrepo.AssetUploadEventConsumer, error) {
	if natsClient == nil || natsClient.Conn() == nil {
		return nil, errors.New("platform asset event consumer requires NATS")
	}
	stream := commonenum.EventStreamConfigFor(commonenum.EventStreamPlatformAssetUploads)
	return &AssetEventRepo{
		logger:          logger,
		natsClient:      natsClient,
		streamName:      stream.Name.String(),
		subject:         commonenum.PlatformAssetUploadSubject,
		consumerDurable: commonenum.EventConsumerPlatformAssetUpload.String(),
	}, nil
}

func (r *AssetEventRepo) Consume(ctx context.Context, handler platformbizrepo.AssetUploadEventHandler) error {
	if handler == nil {
		return errors.New("platform asset event handler is required")
	}
	stream := commonenum.EventStreamConfigFor(commonenum.EventStreamPlatformAssetUploads)
	if err := r.natsClient.EnsureStream(ctx, nats.StreamSpec{
		Name:              stream.Name.String(),
		Subjects:          stream.Subjects,
		Retention:         stream.Retention,
		AllowMsgSchedules: stream.AllowMsgSchedules,
		AllowMsgTTL:       stream.AllowMsgTTL,
		AllowRollup:       stream.AllowRollup,
	}); err != nil {
		return err
	}
	subscription, err := r.natsClient.Consume(ctx, nats.ConsumerSpec{
		Stream:        r.streamName,
		Durable:       r.consumerDurable,
		FilterSubject: r.subject,
		DeliverPolicy: jetstream.DeliverAllPolicy,
	}, func(ctx context.Context, msg *nats.Message) error {
		return r.consumeMessage(ctx, msg, handler)
	})
	if err != nil {
		return err
	}
	r.subscription = subscription
	return nil
}

func (r *AssetEventRepo) Stop(context.Context) error {
	if r.subscription != nil {
		if err := r.subscription.Unsubscribe(); err != nil {
			return err
		}
		r.subscription = nil
	}
	return nil
}

func (r *AssetEventRepo) consumeMessage(
	ctx context.Context,
	msg *nats.Message,
	handler platformbizrepo.AssetUploadEventHandler,
) error {
	events, err := r.decodeMinioAssetUploadEvents(msg.Data)
	if err != nil {
		r.logger.WarnContext(ctx, "discard invalid MinIO asset event", slog.Any("err", err))
		return msg.Term()
	}
	for _, event := range events {
		if err := handler(ctx, event); err != nil {
			if _, invalid := errors.AsType[*platformbizrepo.InvalidAssetUploadEventError](err); invalid {
				r.logger.WarnContext(ctx, "discard invalid asset upload event", slog.Any("err", err))
				return msg.Term()
			}
			r.logger.WarnContext(ctx, "asset upload event handling failed", slog.Any("err", err))
			return err
		}
	}
	return nil
}

func (*AssetEventRepo) decodeMinioAssetUploadEvents(payload []byte) ([]*platformbizrepo.AssetUploadEvent, error) {
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
