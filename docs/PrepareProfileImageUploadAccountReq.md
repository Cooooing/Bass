# PrepareProfileImageUploadAccountReq

请求参数。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**purpose** | **string** | 图片用途，必填；仅用于 BFF 业务校验 | [default to undefined]
**hash** | **string** | 文件 SHA-256，小写十六进制 | [default to undefined]
**mime_type** | **string** | 上传 MIME 类型 | [default to undefined]
**size** | **string** | 文件大小，最大 2 MiB | [default to undefined]

## Example

```typescript
import { PrepareProfileImageUploadAccountReq } from '@bass/bbs-sdk-axios';

const instance: PrepareProfileImageUploadAccountReq = {
    purpose,
    hash,
    mime_type,
    size,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
