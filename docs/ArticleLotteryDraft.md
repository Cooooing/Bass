# ArticleLotteryDraft

协议消息。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**prizes** | [**Array&lt;ArticleLotteryPrizeDraft&gt;**](ArticleLotteryPrizeDraft.md) | 奖项配置 | [optional] [default to undefined]
**draw_at** | **string** | 开奖时间 | [optional] [default to undefined]

## Example

```typescript
import { ArticleLotteryDraft } from '@bass/bbs-sdk-axios';

const instance: ArticleLotteryDraft = {
    prizes,
    draw_at,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
