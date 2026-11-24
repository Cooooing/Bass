# LoginReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**r#type** | **Type** | 登录方式 (enum: LOGIN_TYPE_UNSPECIFIED, LOGIN_TYPE_PASSWORD, LOGIN_TYPE_EMAIL, LOGIN_TYPE_PHONE) | 
**password_credential** | Option<[**models::ReqPasswordCredential**](ReqPasswordCredential.md)> | 密码登录凭证 | [optional]
**email_credential** | Option<[**models::ReqEmailCredential**](ReqEmailCredential.md)> | 邮箱登录凭证 | [optional]
**phone_credential** | Option<[**models::ReqPhoneCredential**](ReqPhoneCredential.md)> | 手机登录凭证 | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


