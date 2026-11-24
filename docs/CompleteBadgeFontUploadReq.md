# CompleteBadgeFontUploadReq

Req 指定已上传字体的内容哈希。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**hash** | **string** | 文件 SHA-256；回调尚未入库时返回可重试的未找到错误。 | [default to undefined]

## Example

```typescript
import { CompleteBadgeFontUploadReq } from '@bass/bbs-sdk-axios';

const instance: CompleteBadgeFontUploadReq = {
    hash,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
