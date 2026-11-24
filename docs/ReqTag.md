# ReqTag

协议消息。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**code** | **string** | 稳定编码 | [default to undefined]
**name** | **string** | 标签名称 | [default to undefined]
**description** | **string** | 描述 | [optional] [default to undefined]
**domain_id** | **string** | 领域 ID | [optional] [default to undefined]
**status** | **string** | 启停状态 | [optional] [default to undefined]
**icon** | **string** | 图标 | [optional] [default to undefined]
**sort** | **number** | 排序值，越小越靠前 | [optional] [default to undefined]

## Example

```typescript
import { ReqTag } from '@bass/bbs-sdk-axios';

const instance: ReqTag = {
    code,
    name,
    description,
    domain_id,
    status,
    icon,
    sort,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
