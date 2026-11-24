# GetArticleReq

请求参数。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**article_id** | **string** | 必填，文章 ID | [default to undefined]
**publish_status** | **string** | 可选，预期发布状态；未传按已发布查询 | [optional] [default to undefined]

## Example

```typescript
import { GetArticleReq } from '@bass/bbs-sdk-axios';

const instance: GetArticleReq = {
    article_id,
    publish_status,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
