

# PrepareProfileImageUploadAccountReq

请求参数。

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**purpose** | [**PurposeEnum**](#PurposeEnum) | 图片用途，必填；仅用于 BFF 业务校验 |  |
|**hash** | **String** | 文件 SHA-256，小写十六进制 |  |
|**mimeType** | **String** | 上传 MIME 类型 |  |
|**size** | **String** | 文件大小，最大 2 MiB |  |



## Enum: PurposeEnum

| Name | Value |
|---- | -----|
| PROFILE_IMAGE_PURPOSE_UNSPECIFIED | &quot;PROFILE_IMAGE_PURPOSE_UNSPECIFIED&quot; |
| PROFILE_IMAGE_PURPOSE_AVATAR | &quot;PROFILE_IMAGE_PURPOSE_AVATAR&quot; |
| PROFILE_IMAGE_PURPOSE_BACKGROUND | &quot;PROFILE_IMAGE_PURPOSE_BACKGROUND&quot; |



