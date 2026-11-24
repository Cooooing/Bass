

# LoginReq

请求参数。

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**type** | [**TypeEnum**](#TypeEnum) | 登录方式 |  |
|**passwordCredential** | [**ReqPasswordCredential**](ReqPasswordCredential.md) | 密码登录凭证 |  [optional] |
|**emailCredential** | [**ReqEmailCredential**](ReqEmailCredential.md) | 邮箱登录凭证 |  [optional] |
|**phoneCredential** | [**ReqPhoneCredential**](ReqPhoneCredential.md) | 手机登录凭证 |  [optional] |



## Enum: TypeEnum

| Name | Value |
|---- | -----|
| LOGIN_TYPE_UNSPECIFIED | &quot;LOGIN_TYPE_UNSPECIFIED&quot; |
| LOGIN_TYPE_PASSWORD | &quot;LOGIN_TYPE_PASSWORD&quot; |
| LOGIN_TYPE_EMAIL | &quot;LOGIN_TYPE_EMAIL&quot; |
| LOGIN_TYPE_PHONE | &quot;LOGIN_TYPE_PHONE&quot; |



