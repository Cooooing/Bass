package domain

import (
	"common/api/gen/common"
	"common/api/gen/common/enums"
	"common/pkg/client"
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const testNatsURL = "nats://192.168.100.10:30083"

func newTestNats(t *testing.T) (*client.NatsClient, func()) {
	t.Helper()
	c, cleanup, err := client.NewNatsClient(log.DefaultLogger, &common.Nats{
		Url:  testNatsURL,
		Name: fmt.Sprintf("test-producer-%d", time.Now().UnixNano()),
	})
	if err != nil {
		t.Fatalf("new nats client: %v", err)
	}
	return c, cleanup
}

func publishEvent(t *testing.T, c *client.NatsClient, subject string, event *enums.Event) {
	t.Helper()
	data, err := proto.Marshal(event)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	err = c.Publish(context.Background(), subject, &client.Message{Data: data})
	if err != nil {
		t.Fatalf("publish to %s: %v", subject, err)
	}
	t.Logf("published %s → %s (receivers=%d)", event.Type, subject, len(event.ReceiverIds))
}

func TestPublishArticlePublished(t *testing.T) {
	c, cleanup := newTestNats(t)
	defer cleanup()

	publishEvent(t, c, "content.article.published", &enums.Event{
		EventId:     uuid.New().String(),
		Type:        enums.EventType_EVENT_TYPE_ARTICLE_PUBLISHED,
		Timestamp:   timestamppb.Now(),
		ReceiverIds: []int64{1, 2, 3},
		Payload: &enums.Event_ArticlePublished{
			ArticlePublished: &enums.ArticlePublishedPayload{
				SenderId:   42,
				SenderName: "test-user",
				ArticleId:  1001,
				Title:      "测试文章",
			},
		},
	})
}

func TestPublishArticleLiked(t *testing.T) {
	c, cleanup := newTestNats(t)
	defer cleanup()

	publishEvent(t, c, "content.article.liked", &enums.Event{
		EventId:     "test-like-001",
		Type:        enums.EventType_EVENT_TYPE_ARTICLE_LIKED,
		Timestamp:   timestamppb.Now(),
		ReceiverIds: []int64{42},
		Payload: &enums.Event_ArticlePublished{
			ArticlePublished: &enums.ArticlePublishedPayload{
				SenderId:   1,
				SenderName: "222",
				ArticleId:  0,
				Title:      "222",
			},
		},
	})
}

func TestPublishUserFollowCreated(t *testing.T) {
	c, cleanup := newTestNats(t)
	defer cleanup()

	publishEvent(t, c, "user.follow.created", &enums.Event{
		EventId:     "test-follow-001",
		Type:        enums.EventType_EVENT_TYPE_USER_FOLLOW_CREATED,
		Timestamp:   timestamppb.Now(),
		ReceiverIds: []int64{42},
		Payload: &enums.Event_UserFollowCreated{
			UserFollowCreated: &enums.UserFollowCreatedPayload{
				SenderId:   99,
				SenderName: "follower",
				FollowedId: 42,
			},
		},
	})
}

func TestPublishCommentPublished(t *testing.T) {
	c, cleanup := newTestNats(t)
	defer cleanup()

	publishEvent(t, c, "content.comment.published", &enums.Event{
		EventId:     "test-comment-001",
		Type:        enums.EventType_EVENT_TYPE_COMMENT_PUBLISHED,
		Timestamp:   timestamppb.Now(),
		ReceiverIds: []int64{42},
		Payload: &enums.Event_CommentPublished{
			CommentPublished: &enums.CommentPublishedPayload{
				SenderId:   7,
				SenderName: "commenter",
				CommentId:  5001,
				ArticleId:  1001,
				AuthorId:   42,
				Content:    "好文章！",
			},
		},
	})
}
