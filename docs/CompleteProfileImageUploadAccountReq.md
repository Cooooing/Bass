

# CompleteProfileImageUploadAccountReq

请求参数。

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**purpose** | [**PurposeEnum**](#PurposeEnum) | 协议字段。 |  |
|**hash** | **String** | 文件 SHA-256；回调尚未入库时返回可重试错误 |  |



## Enum: PurposeEnum

| Name | Value |
|---- | -----|
| PROFILE_IMAGE_PURPOSE_UNSPECIFIED | &quot;PROFILE_IMAGE_PURPOSE_UNSPECIFIED&quot; |
| PROFILE_IMAGE_PURPOSE_AVATAR | &quot;PROFILE_IMAGE_PURPOSE_AVATAR&quot; |
| PROFILE_IMAGE_PURPOSE_BACKGROUND | &quot;PROFILE_IMAGE_PURPOSE_BACKGROUND&quot; |



