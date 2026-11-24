# ArticlePollDraft

协议消息。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**_options** | **Array&lt;string&gt;** | 投票选项 | [optional] [default to undefined]
**ends_at** | **string** | 截止时间 | [optional] [default to undefined]
**max_choices** | **number** | 每人最多选择数 | [optional] [default to undefined]

## Example

```typescript
import { ArticlePollDraft } from '@bass/bbs-sdk-axios';

const instance: ArticlePollDraft = {
    _options,
    ends_at,
    max_choices,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
