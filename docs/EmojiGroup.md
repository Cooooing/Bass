# EmojiGroup

EmojiGroup 是当前账号可操作的表情分组。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **string** | 分组 ID。 | [optional] [default to undefined]
**name** | **string** | 分组名称。 | [optional] [default to undefined]
**items** | [**Array&lt;EmojiItem&gt;**](EmojiItem.md) | 分组内可用表情。 | [optional] [default to undefined]

## Example

```typescript
import { EmojiGroup } from '@bass/bbs-sdk-axios';

const instance: EmojiGroup = {
    id,
    name,
    items,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
