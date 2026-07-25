package channel

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	bizchannel "notify/internal/biz/channel"
	"notify/internal/config"
	notifyenum "notify/internal/enum"
	"strings"
	"time"
)

var _ bizchannel.LarkWebhookClient = (*LarkWebhookClient)(nil)

type LarkWebhookClient struct {
	conf       *config.Bootstrap
	httpClient *http.Client
}

func NewLarkWebhookClient(
	conf *config.Bootstrap,
	httpClient *http.Client,
) *LarkWebhookClient {
	return &LarkWebhookClient{
		conf:       conf,
		httpClient: httpClient,
	}
}

func (c *LarkWebhookClient) SendLarkWebhook(ctx context.Context, req *bizchannel.LarkWebhookRequest) (*bizchannel.SendResult, error) {
	if req == nil || req.Token == "" || req.RequestBody == "" {
		return &bizchannel.SendResult{
			Status: notifyenum.NotificationChannelStatusFailed,
		}, nil
	}
	if c.httpClient == nil {
		return nil, errors.New("lark webhook client is not configured")
	}
	baseURL := "https://open.larksuite.com/open-apis/bot/v2/hook/"
	if c.conf != nil && c.conf.Notify != nil && c.conf.Notify.LarkWebhook != nil && c.conf.Notify.LarkWebhook.BaseUrl != "" {
		baseURL = c.conf.Notify.LarkWebhook.BaseUrl
	}
	requestBody := req.RequestBody
	if strings.TrimSpace(req.Secret) != "" {
		var err error
		requestBody, err = c.signLarkWebhookRequestBody(req.RequestBody, req.Secret, time.Now().Unix())
		if err != nil {
			return nil, err
		}
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, strings.TrimRight(baseURL, "/")+"/"+req.Token, bytes.NewReader([]byte(requestBody)))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	reply, err := c.httpClient.Do(httpReq)
	if err != nil {
		return &bizchannel.SendResult{
			Status:       notifyenum.NotificationChannelStatusUnknown,
			ProviderResp: new(fmt.Sprintf("send lark webhook: %v", err)),
		}, nil
	}
	defer reply.Body.Close()
	body, _ := io.ReadAll(reply.Body)
	status := notifyenum.NotificationChannelStatusSucceeded
	if reply.StatusCode < http.StatusOK || reply.StatusCode >= http.StatusMultipleChoices {
		status = notifyenum.NotificationChannelStatusFailed
	}
	return &bizchannel.SendResult{
		Status:     status,
		HTTPStatus: new(reply.StatusCode),
		RespBody:   new(string(body)),
	}, nil
}

func (c *LarkWebhookClient) signLarkWebhookRequestBody(requestBody string, secret string, timestamp int64) (string, error) {
	var payload map[string]any
	if err := json.Unmarshal([]byte(requestBody), &payload); err != nil {
		return "", err
	}
	if payload == nil {
		return "", errors.New("lark webhook request body must be json object")
	}
	sign, err := c.genLarkWebhookSign(secret, timestamp)
	if err != nil {
		return "", err
	}
	payload["timestamp"] = fmt.Sprintf("%d", timestamp)
	payload["sign"] = sign
	signedBody, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	return string(signedBody), nil
}

func (c *LarkWebhookClient) genLarkWebhookSign(secret string, timestamp int64) (string, error) {
	stringToSign := fmt.Sprintf("%d", timestamp) + "\n" + secret
	h := hmac.New(sha256.New, []byte(stringToSign))
	if _, err := h.Write(nil); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}
