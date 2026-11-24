# BadgeDraft

BadgeDraft 是 Maker 可编辑的徽章草稿。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **string** | 草稿主键。 | [optional] [default to undefined]
**name** | **string** | 草稿名称。 | [optional] [default to undefined]
**description** | **string** | 草稿说明。 | [optional] [default to undefined]
**render_spec** | [**RenderSpec**](RenderSpec.md) | 草稿渲染参数。 | [optional] [default to undefined]
**designer_user_id** | **string** | 设计者用户 ID。 | [optional] [default to undefined]

## Example

```typescript
import { BadgeDraft } from '@bass/bbs-sdk-axios';

const instance: BadgeDraft = {
    id,
    name,
    description,
    render_spec,
    designer_user_id,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
