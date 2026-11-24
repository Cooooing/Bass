# EmojiItem

EmojiItem 是可直接插入编辑器的表情图片。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **string** | 表情项 ID。 | [optional] [default to undefined]
**asset_id** | **string** | 平台图片资产 ID。 | [optional] [default to undefined]
**name** | **string** | 可选表情别名。 | [optional] [default to undefined]
**image_url** | **string** | 可插入 Markdown 的公开图片地址。 | [optional] [default to undefined]

## Example

```typescript
import { EmojiItem } from '@bass/bbs-sdk-axios';

const instance: EmojiItem = {
    id,
    asset_id,
    name,
    image_url,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
