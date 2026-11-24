# ArticleLotteryPrize

协议消息。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **string** | 奖项 ID | [optional] [default to undefined]
**name** | **string** | 奖项名称 | [optional] [default to undefined]
**description** | **string** | 奖项说明 | [optional] [default to undefined]
**winner_limit** | **number** | 中奖名额 | [optional] [default to undefined]
**winner_users** | [**Array&lt;AccountProfile&gt;**](AccountProfile.md) | 按抽取顺序的中奖账号资料 | [optional] [default to undefined]

## Example

```typescript
import { ArticleLotteryPrize } from '@bass/bbs-sdk-axios';

const instance: ArticleLotteryPrize = {
    id,
    name,
    description,
    winner_limit,
    winner_users,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
