# TagQuery

标签查询条件。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ids** | **Array&lt;string&gt;** | 标签 ID 列表。 | [optional] [default to undefined]
**name** | **string** | 标签名称。 | [optional] [default to undefined]
**names** | **Array&lt;string&gt;** | 标签名称列表。 | [optional] [default to undefined]
**description** | **string** | 标签描述。 | [optional] [default to undefined]
**status** | **string** | 标签状态。 | [optional] [default to undefined]
**domain_id** | **string** | 所属板块 ID。 | [optional] [default to undefined]

## Example

```typescript
import { TagQuery } from '@bass/bbs-sdk-axios';

const instance: TagQuery = {
    ids,
    name,
    names,
    description,
    status,
    domain_id,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
