# PrepareMarkdownImageUploadResp

响应结果。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**asset_id** | **string** | 已存在资源的 ID。 | [optional] [default to undefined]
**upload_url** | **string** | 新资源的 MinIO multipart POST 地址。 | [optional] [default to undefined]
**form_fields** | **{ [key: string]: string; }** | MinIO multipart POST 表单字段。 | [optional] [default to undefined]

## Example

```typescript
import { PrepareMarkdownImageUploadResp } from '@bass/bbs-sdk-axios';

const instance: PrepareMarkdownImageUploadResp = {
    asset_id,
    upload_url,
    form_fields,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
