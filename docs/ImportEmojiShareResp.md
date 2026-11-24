# ImportEmojiShareResp

响应结果。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**group** | [**EmojiGroup**](EmojiGroup.md) | 新建的导入分组。 | [optional] [default to undefined]
**skipped_count** | **number** | 因资源不可用而跳过的表情数。 | [optional] [default to undefined]

## Example

```typescript
import { ImportEmojiShareResp } from '@bass/bbs-sdk-axios';

const instance: ImportEmojiShareResp = {
    group,
    skipped_count,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
