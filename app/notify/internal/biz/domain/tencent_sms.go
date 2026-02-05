package domain

import (
	"common/pkg/constant"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"notify/internal/biz/base"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	tencentcloud_errors "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111" // 引入sms
)

type TencentSmsDomain struct {
	*base.BaseDomain
	client *sms.Client
}

func NewTencentSmsDomain(base *base.BaseDomain) *TencentSmsDomain {
	if base.Conf.Sms.Provider != constant.SMSTypeTencent.String() {
		return nil
	}
	// SecretId、SecretKey 查询: https://console.cloud.tencent.com/cam/capi
	credential := common.NewCredential(base.Conf.Sms.Tencent.SecretId, base.Conf.Sms.Tencent.SecretKey)
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

func (d *TencentSmsDomain) SendSms(ctx context.Context, phone []string, params []string) error {
	request := sms.NewSendSmsRequest()
	request.SmsSdkAppId = common.StringPtr(d.Conf.Sms.Tencent.SmsSdkAppId)
	request.SignName = common.StringPtr(d.Conf.Sms.Tencent.SignName)
	request.TemplateId = common.StringPtr(d.Conf.Sms.Tencent.TemplateId)
	request.TemplateParamSet = common.StringPtrs(params)
	request.PhoneNumberSet = common.StringPtrs(phone)

	response, err := d.client.SendSms(request)
	// 处理异常
	if err != nil {
		var sdkErr *tencentcloud_errors.TencentCloudSDKError
		if errors.As(err, &sdkErr) {
			return fmt.Errorf("tencent sms API error has returned: %s: code=%s, message=%s, requestId=%s", err, sdkErr.GetCode(), sdkErr.GetMessage(), sdkErr.GetRequestId())
		}
		return err
	}
	b, _ := json.Marshal(response.Response)
	d.Log.Infof("tencent sms reply: %s", string(b))
	return nil
}
