package client

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"log/slog"
)

const defaultLarkWebhookBaseURL = "https://open.larksuite.com/open-apis/bot/v2/hook/"

// LarkWebhookClient 封装 Lark 自定义机器人 Webhook 调用。
type LarkWebhookClient struct {
	httpClient *http.Client
}

type LarkWebhookRequest struct {
	BaseURL string
	Token   string
	Secret  string
	Timeout time.Duration
	Text    string
}

func NewLarkWebhookClient(
	logger *slog.Logger,
) *LarkWebhookClient {
	return &LarkWebhookClient{
		httpClient: http.DefaultClient,
	}
}

func (c *LarkWebhookClient) SendText(
	ctx context.Context,
	req *LarkWebhookRequest,
) error {
	if req == nil || strings.TrimSpace(req.Token) == "" || strings.TrimSpace(req.Text) == "" {
		return nil
	}
	if req.Timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, req.Timeout)
		defer cancel()
	}
	body := map[string]any{
		"msg_type": "text",
		"content":  map[string]string{"text": req.Text},
	}
	if strings.TrimSpace(req.Secret) != "" {
		timestamp := time.Now().Unix()
		body["timestamp"] = fmt.Sprintf("%d", timestamp)
		body["sign"] = c.sign(req.Secret, timestamp)
	}
	payload, err := json.Marshal(body)
	if err != nil {
		return err
	}
	baseURL := strings.TrimSpace(req.BaseURL)
	if baseURL == "" {
		baseURL = defaultLarkWebhookBaseURL
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, strings.TrimRight(baseURL, "/")+"/"+req.Token, bytes.NewReader(payload))
	if err != nil {
		return err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	reply, err := c.httpClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer reply.Body.Close()
	respBody, _ := io.ReadAll(reply.Body)
	if reply.StatusCode < http.StatusOK || reply.StatusCode >= http.StatusMultipleChoices {
		text := strings.TrimSpace(string(respBody))
		text = strings.ReplaceAll(text, "\r", " ")
		text = strings.ReplaceAll(text, "\n", " ")
		if len(text) > 120 {
			text = text[:120] + "..."
		}
		return fmt.Errorf("send lark webhook failed: status=%d resp_summary=len=%d text=%q", reply.StatusCode, len(respBody), text)
	}
	return nil
}

func (c *LarkWebhookClient) sign(
	secret string,
	timestamp int64,
) string {
	stringToSign := fmt.Sprintf("%d", timestamp) + "\n" + secret
	h := hmac.New(sha256.New, []byte(stringToSign))
	_, _ = h.Write(nil)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
