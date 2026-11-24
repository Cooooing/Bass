# ArticlePoll

协议消息。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ends_at** | **string** | 截止时间 | [optional] [default to undefined]
**max_choices** | **number** | 每人最多选择数 | [optional] [default to undefined]
**_options** | [**Array&lt;ArticlePollOption&gt;**](ArticlePollOption.md) | 投票选项 | [optional] [default to undefined]
**selected_option_ids** | **Array&lt;string&gt;** | 当前账号已选选项 ID | [optional] [default to undefined]
**result_visible** | **boolean** | 当前账号是否可查看实时结果 | [optional] [default to undefined]
**participant_count** | **string** | 已投票账号数 | [optional] [default to undefined]
**total_vote_count** | **string** | 所有选项累计票数 | [optional] [default to undefined]

## Example

```typescript
import { ArticlePoll } from '@bass/bbs-sdk-axios';

const instance: ArticlePoll = {
    ends_at,
    max_choices,
    _options,
    selected_option_ids,
    result_visible,
    participant_count,
    total_vote_count,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
