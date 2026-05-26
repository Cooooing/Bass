package channel

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	bizchannel "notify/internal/biz/channel"
	"notify/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	tencenterrors "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

// TencentSMSClient 调用腾讯云短信接口。
type TencentSMSClient struct {
	conf   *conf.Bootstrap
	log    *log.Helper
	client *sms.Client
}

func NewTencentSMSClient(conf *conf.Bootstrap, logger log.Logger) (*TencentSMSClient, error) {
	if conf == nil || conf.Server == nil || conf.Server.Sms == nil || conf.Server.Sms.Tencent == nil {
		return nil, errors.New("tencent sms config is required")
	}
	tencentConf := conf.Server.Sms.Tencent
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
		log:    log.NewHelper(logger),
		client: client,
	}, nil
}

func (c *TencentSMSClient) Send(_ context.Context, req *bizchannel.SendReq) error {
	if req == nil || req.Target == "" {
		return nil
	}
	tencentConf := c.conf.Server.Sms.Tencent
	templateID := req.TemplateID
	if templateID == "" {
		templateID = tencentConf.TemplateId
	}
	if templateID == "" {
		return errors.New("tencent sms template id is required")
	}
	if len(req.TemplateParams) == 0 {
		return errors.New("tencent sms template params are required")
	}
	request := sms.NewSendSmsRequest()
	request.SmsSdkAppId = new(tencentConf.SmsSdkAppId)
	request.SignName = new(tencentConf.SignName)
	request.TemplateId = new(templateID)
	request.TemplateParamSet = common.StringPtrs(req.TemplateParams)
	request.PhoneNumberSet = common.StringPtrs([]string{req.Target})

	response, err := c.client.SendSms(request)
	if err != nil {
		if sdkErr, ok := errors.AsType[*tencenterrors.TencentCloudSDKError](err); ok {
			return fmt.Errorf("send tencent sms: code=%s message=%s request_id=%s", sdkErr.GetCode(), sdkErr.GetMessage(), sdkErr.GetRequestId())
		}
		return fmt.Errorf("send tencent sms: %w", err)
	}
	reply, _ := json.Marshal(response.Response)
	c.log.Infof("tencent sms sent: target=%s reply=%s", req.Target, string(reply))
	return nil
}
