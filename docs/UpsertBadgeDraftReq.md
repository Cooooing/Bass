# UpsertBadgeDraftReq

Req 缺失 draft_id 时创建，存在时更新。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**draft_id** | **string** | 待保存草稿 ID。 | [optional] [default to undefined]
**name** | **string** | 草稿名称。 | [optional] [default to undefined]
**description** | **string** | 草稿说明。 | [optional] [default to undefined]
**render_spec** | [**RenderSpec**](RenderSpec.md) | 草稿渲染参数。 | [optional] [default to undefined]

## Example

```typescript
import { UpsertBadgeDraftReq } from '@bass/bbs-sdk-axios';

const instance: UpsertBadgeDraftReq = {
    draft_id,
    name,
    description,
    render_spec,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
