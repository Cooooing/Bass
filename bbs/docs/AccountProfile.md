

# AccountProfile


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**id** | **String** |  |  [optional] |
|**name** | **String** |  |  [optional] |
|**nickname** | **String** |  |  [optional] |
|**url** | **String** |  |  [optional] |
|**avatarUrl** | **String** |  |  [optional] |
|**introduction** | **String** |  |  [optional] |
|**mbti** | [**MbtiEnum**](#MbtiEnum) |  |  [optional] |
|**status** | [**StatusEnum**](#StatusEnum) |  |  [optional] |
|**followCount** | **Integer** |  |  [optional] |
|**followerCount** | **Integer** |  |  [optional] |
|**createdAt** | **OffsetDateTime** |  |  [optional] |
|**updatedAt** | **OffsetDateTime** |  |  [optional] |



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
| ACCOUNT_STATUS_CANCELLED | &quot;ACCOUNT_STATUS_CANCELLED&quot; |



