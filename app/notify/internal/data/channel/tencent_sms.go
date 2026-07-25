package channel

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	bizchannel "notify/internal/biz/channel"
	"notify/internal/config"
	notifyenum "notify/internal/enum"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	tencenterrors "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

var _ bizchannel.TencentSMSClient = (*TencentSMSClient)(nil)

type TencentSMSClient struct {
	conf   *config.Bootstrap
	client *sms.Client
}

func NewTencentSMSClient(
	conf *config.Bootstrap,
) (*TencentSMSClient, error) {
	if conf == nil || conf.Notify == nil || conf.Notify.Sms == nil || !conf.Notify.Sms.Enable {
		return &TencentSMSClient{
			conf: conf,
		}, nil
	}
	if conf.Notify.Sms.Tencent == nil {
		return nil, errors.New("tencent sms config is required")
	}
	tencentConf := conf.Notify.Sms.Tencent
	credential := common.NewCredential(tencentConf.SecretId, tencentConf.SecretKey)
	clientProfile := profile.NewClientProfile()
	clientProfile.HttpProfile.ReqMethod = "POST"
	clientProfile.HttpProfile.ReqTimeout = 10
	clientProfile.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	clientProfile.SignMethod = "HmacSHA1"

	region := tencentConf.Diyu
	if region == "" {
		region = "ap-beijing"
	}
	client, err := sms.NewClient(credential, region, clientProfile)
	if err != nil {
		return nil, err
	}
	return &TencentSMSClient{
		conf:   conf,
		client: client,
	}, nil
}

func (c *TencentSMSClient) SendTencentSMS(_ context.Context, req *bizchannel.TencentSMSRequest) (*bizchannel.SendResult, error) {
	if req == nil || req.Phone == "" {
		return &bizchannel.SendResult{
			Status: notifyenum.NotificationChannelStatusFailed,
		}, nil
	}
	if c.conf == nil || c.conf.Notify == nil || c.conf.Notify.Sms == nil || !c.conf.Notify.Sms.Enable {
		return &bizchannel.SendResult{
			Status: notifyenum.NotificationChannelStatusSkipped,
		}, nil
	}
	if c.conf.Notify.Sms.Tencent == nil || c.client == nil {
		return nil, errors.New("tencent sms config is required")
	}
	if req.ProviderTemplateID == "" {
		return &bizchannel.SendResult{
			Status: notifyenum.NotificationChannelStatusFailed,
		}, nil
	}
	request := sms.NewSendSmsRequest()
	request.SmsSdkAppId = &req.SMSSDKAppID
	request.SignName = &req.SignName
	request.TemplateId = &req.ProviderTemplateID
	request.TemplateParamSet = common.StringPtrs(req.TemplateParams)
	request.PhoneNumberSet = common.StringPtrs([]string{req.Phone})

	resp, err := c.client.SendSms(request)
	if err != nil {
		if sdkErr, ok := errors.AsType[*tencenterrors.TencentCloudSDKError](err); ok {
			return &bizchannel.SendResult{
				Status:            notifyenum.NotificationChannelStatusFailed,
				ProviderRequestID: new(sdkErr.GetRequestId()),
				ProviderCode:      new(sdkErr.GetCode()),
				ProviderMessage:   new(sdkErr.GetMessage()),
			}, nil
		}
		return &bizchannel.SendResult{
			Status:          notifyenum.NotificationChannelStatusUnknown,
			ProviderMessage: new(fmt.Sprintf("send tencent sms: %v", err)),
		}, nil
	}
	reply, _ := json.Marshal(resp.Response)
	requestID := ""
	if resp.Response != nil && resp.Response.RequestId != nil {
		requestID = *resp.Response.RequestId
	}
	return &bizchannel.SendResult{
		Status:            notifyenum.NotificationChannelStatusSucceeded,
		ProviderRequestID: new(requestID),
		ProviderResp:      new(string(reply)),
	}, nil
}
