# PrepareProfileImageUploadAccountReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**purpose** | **Purpose** | 图片用途，必填；仅用于 BFF 业务校验 (enum: PROFILE_IMAGE_PURPOSE_UNSPECIFIED, PROFILE_IMAGE_PURPOSE_AVATAR, PROFILE_IMAGE_PURPOSE_BACKGROUND) | 
**hash** | **String** | 文件 SHA-256，小写十六进制 | 
**mime_type** | **String** | 上传 MIME 类型 | 
**size** | **String** | 文件大小，最大 2 MiB | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


