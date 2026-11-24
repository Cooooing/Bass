# RefreshTokenResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**access_token** | Option<**String**> | 新的访问令牌 | [optional]
**refresh_token** | Option<**String**> | 新的刷新令牌 | [optional]
**access_token_expires_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> | 访问令牌过期时间 | [optional]
**refresh_token_expires_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> | 刷新令牌空闲过期时间 | [optional]
**session_expires_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> | session 绝对过期时间 | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


