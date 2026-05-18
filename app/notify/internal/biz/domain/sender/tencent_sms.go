package sender

import (
	v1 "common/api/gen/notify/v1"
	"common/pkg/constant"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"notify/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	tencentcloud_errors "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

// TencentSmsSender 腾讯云短信发送器
type TencentSmsSender struct {
	conf   *conf.Bootstrap
	log    *log.Helper
	client *sms.Client
}

func NewTencentSmsSender(conf *conf.Bootstrap, logger log.Logger) *TencentSmsSender {
	if conf.Server.Sms.Provider != constant.SMSTypeTencent.String() {
		return nil
	}
	credential := common.NewCredential(conf.Server.Sms.Tencent.SecretId, conf.Server.Sms.Tencent.SecretKey)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	cpf.HttpProfile.ReqTimeout = 10
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	cpf.SignMethod = "HmacSHA1"
	client, _ := sms.NewClient(credential, "ap-beijing", cpf)
	return &TencentSmsSender{
		conf:   conf,
		log:    log.NewHelper(logger),
		client: client,
	}
}

func (s *TencentSmsSender) Channel() v1.NotificationChannel {
	return v1.NotificationChannel_NOTIFICATION_CHANNEL_SMS
}

func (s *TencentSmsSender) Send(ctx context.Context, req *SendRequest) error {
	phone := req.UserInfo.Phone
	if phone == "" {
		return nil
	}

	request := sms.NewSendSmsRequest()
	request.SmsSdkAppId = new(s.conf.Server.Sms.Tencent.SmsSdkAppId)
	request.SignName = new(s.conf.Server.Sms.Tencent.SignName)
	request.TemplateId = new(s.conf.Server.Sms.Tencent.TemplateId)
	request.TemplateParamSet = common.StringPtrs([]string{req.Title, req.Content})
	request.PhoneNumberSet = common.StringPtrs([]string{phone})

	response, err := s.client.SendSms(request)
	if err != nil {
		if sdkErr, ok := errors.AsType[*tencentcloud_errors.TencentCloudSDKError](err); ok {
			return fmt.Errorf("tencent sms API error: %s: code=%s, message=%s, requestId=%s",
				err, sdkErr.GetCode(), sdkErr.GetMessage(), sdkErr.GetRequestId())
		}
		return err
	}
	b, _ := json.Marshal(response.Response)
	s.log.Infof("tencent sms reply: %s", string(b))
	return nil
}
