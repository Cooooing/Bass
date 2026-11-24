# ArticleViewHistoryItem

协议消息。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**article** | [**ArticleListItem**](ArticleListItem.md) | 当前仍可展示的文章 | [optional] [default to undefined]
**viewed_at** | **string** | 最近浏览时间 | [optional] [default to undefined]

## Example

```typescript
import { ArticleViewHistoryItem } from '@bass/bbs-sdk-axios';

const instance: ArticleViewHistoryItem = {
    article,
    viewed_at,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
