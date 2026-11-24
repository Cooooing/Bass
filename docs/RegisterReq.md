# RegisterReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**r#type** | **Type** | 注册方式 (enum: REGISTER_TYPE_UNSPECIFIED, REGISTER_TYPE_EMAIL, REGISTER_TYPE_PHONE) | 
**name** | **String** | 账号名 | 
**password** | **String** | 账号密码 | 
**nickname** | Option<**String**> | 昵称 | [optional]
**email_credential** | Option<[**models::ReqEmailCredential**](ReqEmailCredential.md)> | 邮箱注册凭证 | [optional]
**phone_credential** | Option<[**models::ReqPhoneCredential**](ReqPhoneCredential.md)> | 手机注册凭证 | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


