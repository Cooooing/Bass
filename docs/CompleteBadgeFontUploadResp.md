# CompleteBadgeFontUploadResp

Resp 返回已确认的字体资源与前端本地预览地址。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**asset_id** | **string** | 已确认资源的 ID。 | [optional] [default to undefined]
**font_url** | **string** | 前端本地预览可用的字体公开地址。 | [optional] [default to undefined]

## Example

```typescript
import { CompleteBadgeFontUploadResp } from '@bass/bbs-sdk-axios';

const instance: CompleteBadgeFontUploadResp = {
    asset_id,
    font_url,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
