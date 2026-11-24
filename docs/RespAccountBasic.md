

# RespAccountBasic

协议消息。

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**id** | **String** | 账号 ID |  [optional] |
|**name** | **String** | 账号名 |  [optional] |
|**nickname** | **String** | 昵称 |  [optional] |
|**url** | **String** | 个人主页 URL |  [optional] |
|**avatarUrl** | **String** | 头像 URL |  [optional] |
|**introduction** | **String** | 个人简介 |  [optional] |
|**mbti** | **String** | MBTI 类型，规范化为大写 |  [optional] |
|**status** | [**StatusEnum**](#StatusEnum) | 账号状态 |  [optional] |
|**followCount** | **Integer** | 关注数 |  [optional] |
|**followerCount** | **Integer** | 粉丝数 |  [optional] |
|**createdAt** | **OffsetDateTime** | 创建时间 |  [optional] |
|**updatedAt** | **OffsetDateTime** | 更新时间 |  [optional] |



## Enum: StatusEnum

| Name | Value |
|---- | -----|
| ACCOUNT_STATUS_UNSPECIFIED | &quot;ACCOUNT_STATUS_UNSPECIFIED&quot; |
| ACCOUNT_STATUS_NORMAL | &quot;ACCOUNT_STATUS_NORMAL&quot; |
| ACCOUNT_STATUS_BANNED | &quot;ACCOUNT_STATUS_BANNED&quot; |
| ACCOUNT_STATUS_CANCELLED | &quot;ACCOUNT_STATUS_CANCELLED&quot; |



