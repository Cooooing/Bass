# ArticleLottery

协议消息。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**draw_at** | **string** | 开奖时间 | [optional] [default to undefined]
**drawn_at** | **string** | 实际开奖时间 | [optional] [default to undefined]
**entry_count** | **string** | 有效报名人数 | [optional] [default to undefined]
**entered** | **boolean** | 当前账号是否已报名 | [optional] [default to undefined]
**prizes** | [**Array&lt;ArticleLotteryPrize&gt;**](ArticleLotteryPrize.md) | 奖项和中奖结果 | [optional] [default to undefined]

## Example

```typescript
import { ArticleLottery } from '@bass/bbs-sdk-axios';

const instance: ArticleLottery = {
    draw_at,
    drawn_at,
    entry_count,
    entered,
    prizes,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
