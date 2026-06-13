# RequestTag


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | 标签名称。 | [default to undefined]
**description** | **string** | 标签描述。 | [optional] [default to undefined]
**domain_id** | **string** | 所属板块 ID。 | [optional] [default to undefined]
**status** | **string** | 标签启停状态。 | [optional] [default to undefined]

## Example

```typescript
import { RequestTag } from '@bass/bbs-sdk-axios';

const instance: RequestTag = {
    name,
    description,
    domain_id,
    status,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
