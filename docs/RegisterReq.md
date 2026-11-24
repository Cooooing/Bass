

# RegisterReq

请求参数。

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**type** | [**TypeEnum**](#TypeEnum) | 注册方式 |  |
|**name** | **String** | 账号名 |  |
|**password** | **String** | 账号密码 |  |
|**nickname** | **String** | 昵称 |  [optional] |
|**emailCredential** | [**ReqEmailCredential**](ReqEmailCredential.md) | 邮箱注册凭证 |  [optional] |
|**phoneCredential** | [**ReqPhoneCredential**](ReqPhoneCredential.md) | 手机注册凭证 |  [optional] |



## Enum: TypeEnum

| Name | Value |
|---- | -----|
| REGISTER_TYPE_UNSPECIFIED | &quot;REGISTER_TYPE_UNSPECIFIED&quot; |
| REGISTER_TYPE_EMAIL | &quot;REGISTER_TYPE_EMAIL&quot; |
| REGISTER_TYPE_PHONE | &quot;REGISTER_TYPE_PHONE&quot; |



