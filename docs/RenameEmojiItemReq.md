# RenameEmojiItemReq

请求参数。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**item_id** | **string** | 待重命名表情项 ID。 | [optional] [default to undefined]
**name** | **string** | 新表情别名。 | [optional] [default to undefined]

## Example

```typescript
import { RenameEmojiItemReq } from '@bass/bbs-sdk-axios';

const instance: RenameEmojiItemReq = {
    item_id,
    name,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
