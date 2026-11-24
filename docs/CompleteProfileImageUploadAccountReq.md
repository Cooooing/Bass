# CompleteProfileImageUploadAccountReq

请求参数。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**purpose** | **string** | 协议字段。 | [default to undefined]
**hash** | **string** | 文件 SHA-256；回调尚未入库时返回可重试错误 | [default to undefined]

## Example

```typescript
import { CompleteProfileImageUploadAccountReq } from '@bass/bbs-sdk-axios';

const instance: CompleteProfileImageUploadAccountReq = {
    purpose,
    hash,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
