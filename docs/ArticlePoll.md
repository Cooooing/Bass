# ArticlePoll


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ends_at** | **string** |  | [optional] [default to undefined]
**max_choices** | **number** |  | [optional] [default to undefined]
**_options** | [**Array&lt;ArticlePollOption&gt;**](ArticlePollOption.md) |  | [optional] [default to undefined]
**selected_option_ids** | **Array&lt;string&gt;** |  | [optional] [default to undefined]
**result_visible** | **boolean** |  | [optional] [default to undefined]
**participant_count** | **string** |  | [optional] [default to undefined]
**total_vote_count** | **string** |  | [optional] [default to undefined]

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
