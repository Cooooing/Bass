# AddEmojiItemReq

请求参数。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**group_id** | **string** | 可选目标分组 ID；缺省时加入默认分组。 | [optional] [default to undefined]
**asset_id** | **string** | 已确认可用的平台图片资产 ID。 | [optional] [default to undefined]
**name** | **string** | 可选表情别名。 | [optional] [default to undefined]

## Example

```typescript
import { AddEmojiItemReq } from '@bass/bbs-sdk-axios';

const instance: AddEmojiItemReq = {
    group_id,
    asset_id,
    name,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
