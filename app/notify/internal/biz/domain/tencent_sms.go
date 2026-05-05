package domain

import (
	"common/pkg/constant"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	domainbase "notify/internal/biz/base"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	tencentcloud_errors "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

type TencentSmsDomain struct {
	*domainbase.BaseDomain
	client *sms.Client
}

func NewTencentSmsDomain(base *domainbase.BaseDomain) *TencentSmsDomain {
	if base.Conf.Server.Sms.Provider != constant.SMSTypeTencent.String() {
		return nil
	}
	credential := common.NewCredential(base.Conf.Server.Sms.Tencent.SecretId, base.Conf.Server.Sms.Tencent.SecretKey)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	cpf.HttpProfile.ReqTimeout = 10
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	cpf.SignMethod = "HmacSHA1"
	client, _ := sms.NewClient(credential, "ap-beijing", cpf)
	return &TencentSmsDomain{
		BaseDomain: base,
		client:     client,
	}
}

func (d *TencentSmsDomain) Send(ctx context.Context, phone []string, params []string) error {
	request := sms.NewSendSmsRequest()
	request.SmsSdkAppId = new(d.Conf.Server.Sms.Tencent.SmsSdkAppId)
	request.SignName = new(d.Conf.Server.Sms.Tencent.SignName)
	request.TemplateId = new(d.Conf.Server.Sms.Tencent.TemplateId)
	request.TemplateParamSet = common.StringPtrs(params)
	request.PhoneNumberSet = common.StringPtrs(phone)

	response, err := d.client.SendSms(request)
	if err != nil {
		if sdkErr, ok := errors.AsType[*tencentcloud_errors.TencentCloudSDKError](err); ok {
			return fmt.Errorf("tencent sms API error: %s: code=%s, message=%s, requestId=%s", err, sdkErr.GetCode(), sdkErr.GetMessage(), sdkErr.GetRequestId())
		}
		return err
	}
	b, _ := json.Marshal(response.Response)
	d.Log.Infof("tencent sms reply: %s", string(b))
	return nil
}
