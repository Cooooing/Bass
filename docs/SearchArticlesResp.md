# SearchArticlesResp

响应结果。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**rows** | [**Array&lt;ArticleListItem&gt;**](ArticleListItem.md) | 文章列表 | [optional] [default to undefined]
**next_cursor** | **string** | 下一页游标 | [optional] [default to undefined]

## Example

```typescript
import { SearchArticlesResp } from '@bass/bbs-sdk-axios';

const instance: SearchArticlesResp = {
    rows,
    next_cursor,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
