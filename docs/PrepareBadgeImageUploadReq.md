# PrepareBadgeImageUploadReq

请求参数。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**hash** | **string** | 文件 SHA-256，小写十六进制。 | [default to undefined]
**mime_type** | **string** | 上传图片 MIME 类型。 | [default to undefined]
**size** | **string** | 文件大小，最大 5 MiB。 | [default to undefined]

## Example

```typescript
import { PrepareBadgeImageUploadReq } from '@bass/bbs-sdk-axios';

const instance: PrepareBadgeImageUploadReq = {
    hash,
    mime_type,
    size,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
