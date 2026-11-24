# ReorderEmojiItemsReq

请求参数。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**group_id** | **string** | 可选目标分组 ID；为空时重排未分组表情。 | [optional] [default to undefined]
**item_ids** | **Array&lt;string&gt;** | 按最终展示顺序排列的完整表情项 ID 序列。 | [optional] [default to undefined]

## Example

```typescript
import { ReorderEmojiItemsReq } from '@bass/bbs-sdk-axios';

const instance: ReorderEmojiItemsReq = {
    group_id,
    item_ids,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
