# LikeArticleReq

请求参数。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**article_id** | **string** | 必填，文章 ID | [default to undefined]
**active** | **boolean** | 必填，是否点赞 | [default to undefined]

## Example

```typescript
import { LikeArticleReq } from '@bass/bbs-sdk-axios';

const instance: LikeArticleReq = {
    article_id,
    active,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
