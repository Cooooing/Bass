# UpsertBadgeDraftReq

Req 缺失 draft_id 时创建，存在时更新。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**draft_id** | **string** |  | [optional] [default to undefined]
**name** | **string** |  | [optional] [default to undefined]
**description** | **string** |  | [optional] [default to undefined]
**render_spec** | [**RenderSpec**](RenderSpec.md) |  | [optional] [default to undefined]

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
