

# AccountProfile

账号展示资料。

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**id** | **String** | 账号 ID。 |  [optional] |
|**name** | **String** | 账号名。 |  [optional] |
|**nickname** | **String** | 昵称。 |  [optional] |
|**url** | **String** | 个人主页 URL。 |  [optional] |
|**avatarUrl** | **String** | 头像 URL。 |  [optional] |
|**introduction** | **String** | 个人简介。 |  [optional] |
|**mbti** | [**MbtiEnum**](#MbtiEnum) | MBTI 类型。 |  [optional] |
|**status** | [**StatusEnum**](#StatusEnum) | 账号状态。 |  [optional] |
|**followCount** | **Integer** | 关注数量。 |  [optional] |
|**followerCount** | **Integer** | 粉丝数量。 |  [optional] |
|**createdAt** | **String** | 创建时间。 |  [optional] |
|**updatedAt** | **String** | 更新时间。 |  [optional] |



## Enum: MbtiEnum

| Name | Value |
|---- | -----|
| MBTI_UNSPECIFIED | &quot;MBTI_UNSPECIFIED&quot; |
| MBTI_INTJ | &quot;MBTI_INTJ&quot; |
| MBTI_INTP | &quot;MBTI_INTP&quot; |
| MBTI_ENTJ | &quot;MBTI_ENTJ&quot; |
| MBTI_ENTP | &quot;MBTI_ENTP&quot; |
| MBTI_INFJ | &quot;MBTI_INFJ&quot; |
| MBTI_INFP | &quot;MBTI_INFP&quot; |
| MBTI_ENFJ | &quot;MBTI_ENFJ&quot; |
| MBTI_ENFP | &quot;MBTI_ENFP&quot; |
| MBTI_ISTJ | &quot;MBTI_ISTJ&quot; |
| MBTI_ISFJ | &quot;MBTI_ISFJ&quot; |
| MBTI_ESTJ | &quot;MBTI_ESTJ&quot; |
| MBTI_ESFJ | &quot;MBTI_ESFJ&quot; |
| MBTI_ISTP | &quot;MBTI_ISTP&quot; |
| MBTI_ISFP | &quot;MBTI_ISFP&quot; |
| MBTI_ESTP | &quot;MBTI_ESTP&quot; |
| MBTI_ESFP | &quot;MBTI_ESFP&quot; |



## Enum: StatusEnum

| Name | Value |
|---- | -----|
| ACCOUNT_STATUS_UNSPECIFIED | &quot;ACCOUNT_STATUS_UNSPECIFIED&quot; |
| ACCOUNT_STATUS_NORMAL | &quot;ACCOUNT_STATUS_NORMAL&quot; |
| ACCOUNT_STATUS_BANNED | &quot;ACCOUNT_STATUS_BANNED&quot; |
| ACCOUNT_STATUS_DELETED | &quot;ACCOUNT_STATUS_DELETED&quot; |



