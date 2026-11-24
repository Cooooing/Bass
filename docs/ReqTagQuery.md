# ReqTagQuery

协议消息。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ids** | **Array&lt;string&gt;** | 标签 ID 列表 | [optional] [default to undefined]
**code** | **string** | 稳定编码 | [optional] [default to undefined]
**name** | **string** | 标签名称 | [optional] [default to undefined]
**names** | **Array&lt;string&gt;** | 标签名称列表 | [optional] [default to undefined]
**description** | **string** | 描述 | [optional] [default to undefined]
**status** | **string** | 启停状态 | [optional] [default to undefined]
**domain_id** | **string** | 领域 ID | [optional] [default to undefined]

## Example

```typescript
import { ReqTagQuery } from '@bass/bbs-sdk-axios';

const instance: ReqTagQuery = {
    ids,
    code,
    name,
    names,
    description,
    status,
    domain_id,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
