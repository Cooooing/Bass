package channel

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	bizchannel "notify/internal/biz/channel"
	"notify/internal/conf"
	notifyenum "notify/internal/enum"
	"strings"
)

var _ bizchannel.LarkWebhookClient = (*LarkWebhookClient)(nil)

type LarkWebhookClient struct {
	conf       *conf.Bootstrap
	httpClient *http.Client
}

func NewLarkWebhookClient(conf *conf.Bootstrap, httpClient *http.Client) *LarkWebhookClient {
	return &LarkWebhookClient{
		conf:       conf,
		httpClient: httpClient,
	}
}

func (c *LarkWebhookClient) SendLarkWebhook(ctx context.Context, req *bizchannel.LarkWebhookRequest) (*bizchannel.SendResult, error) {
	if req == nil || req.Token == "" || req.RequestBody == "" {
		return &bizchannel.SendResult{Status: notifyenum.NotificationChannelStatusFailed}, nil
	}
	if c.httpClient == nil {
		return nil, errors.New("lark webhook client is not configured")
	}
	baseURL := "https://open.larksuite.com/open-apis/bot/v2/hook/"
	if c.conf != nil && c.conf.Server != nil && c.conf.Server.LarkWebhook != nil && c.conf.Server.LarkWebhook.BaseUrl != "" {
		baseURL = c.conf.Server.LarkWebhook.BaseUrl
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, strings.TrimRight(baseURL, "/")+"/"+req.Token, bytes.NewReader([]byte(req.RequestBody)))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	reply, err := c.httpClient.Do(httpReq)
	if err != nil {
		return &bizchannel.SendResult{Status: notifyenum.NotificationChannelStatusUnknown, ProviderResponse: new(fmt.Sprintf("send lark webhook: %v", err))}, nil
	}
	defer reply.Body.Close()
	body, _ := io.ReadAll(reply.Body)
	status := notifyenum.NotificationChannelStatusSucceeded
	if reply.StatusCode < http.StatusOK || reply.StatusCode >= http.StatusMultipleChoices {
		status = notifyenum.NotificationChannelStatusFailed
	}
	return &bizchannel.SendResult{
		Status:       status,
		HTTPStatus:   new(reply.StatusCode),
		ResponseBody: new(string(body)),
	}, nil
}
