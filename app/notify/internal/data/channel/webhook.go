package channel

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	bizchannel "notify/internal/biz/channel"
	"time"
)

// WebhookChannel 通过 HTTP POST 投递通知。
type WebhookChannel struct {
	httpClient *http.Client
}

func NewWebhookChannel() *WebhookChannel {
	return &WebhookChannel{
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *WebhookChannel) Send(ctx context.Context, req *bizchannel.SendReq) error {
	if req == nil || req.Target == "" {
		return nil
	}
	if c.httpClient == nil {
		return errors.New("webhook client is not configured")
	}
	body, err := json.Marshal(map[string]any{
		"event_id":    req.EventID,
		"event_type":  string(req.EventType),
		"receiver_id": req.ReceiverID,
		"channel":     string(req.Channel),
		"title":       req.Title,
		"content":     req.Content,
	})
	if err != nil {
		return err
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, req.Target, bytes.NewReader(body))
	if err != nil {
		return err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	reply, err := c.httpClient.Do(httpReq)
	if err != nil {
		return fmt.Errorf("send webhook: %w", err)
	}
	defer reply.Body.Close()
	if reply.StatusCode < http.StatusOK || reply.StatusCode >= http.StatusMultipleChoices {
		return fmt.Errorf("send webhook: status=%d", reply.StatusCode)
	}
	return nil
}
