# PublishArticleReq

请求参数。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**article_id** | **string** | 必填，文章 ID | [default to undefined]
**scheduled_at** | **string** | 可选，定时发布时间 | [optional] [default to undefined]

## Example

```typescript
import { PublishArticleReq } from '@bass/bbs-sdk-axios';

const instance: PublishArticleReq = {
    article_id,
    scheduled_at,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
