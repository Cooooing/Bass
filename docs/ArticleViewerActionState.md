# ArticleViewerActionState

协议消息。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**liked** | **boolean** | 是否已点赞 | [optional] [default to undefined]
**thanked** | **boolean** | 是否已感谢 | [optional] [default to undefined]
**collected** | **boolean** | 是否已收藏 | [optional] [default to undefined]
**rewarded** | **boolean** | 是否已打赏 | [optional] [default to undefined]
**watched** | **boolean** | 是否已关注 | [optional] [default to undefined]

## Example

```typescript
import { ArticleViewerActionState } from '@bass/bbs-sdk-axios';

const instance: ArticleViewerActionState = {
    liked,
    thanked,
    collected,
    rewarded,
    watched,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
