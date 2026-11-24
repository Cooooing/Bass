# AccountProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | Option<**String**> | 账号 ID | [optional]
**name** | Option<**String**> | 账号名 | [optional]
**nickname** | Option<**String**> | 昵称 | [optional]
**url** | Option<**String**> | 个人主页 | [optional]
**avatar_url** | Option<**String**> | 头像 URL | [optional]
**introduction** | Option<**String**> | 个人简介 | [optional]
**mbti** | Option<**String**> | MBTI 类型，规范化为大写 | [optional]
**status** | Option<**Status**> | 账号状态 (enum: ACCOUNT_STATUS_UNSPECIFIED, ACCOUNT_STATUS_NORMAL, ACCOUNT_STATUS_BANNED, ACCOUNT_STATUS_CANCELLED) | [optional]
**follow_count** | Option<**i32**> | 关注数 | [optional]
**follower_count** | Option<**i32**> | 粉丝数 | [optional]
**created_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> | 创建时间 | [optional]
**updated_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> | 更新时间 | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


