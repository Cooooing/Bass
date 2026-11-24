# PrepareProfileImageUploadAccountResp

响应结果。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**asset_id** | **string** | 已存在时直接返回 | [optional] [default to undefined]
**upload_url** | **string** | 新资源的 MinIO multipart POST 地址 | [optional] [default to undefined]
**form_fields** | **{ [key: string]: string; }** | 协议字段。 | [optional] [default to undefined]

## Example

```typescript
import { PrepareProfileImageUploadAccountResp } from '@bass/bbs-sdk-axios';

const instance: PrepareProfileImageUploadAccountResp = {
    asset_id,
    upload_url,
    form_fields,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
