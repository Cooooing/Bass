# ReqDomainQuery

协议消息。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ids** | **Array&lt;string&gt;** | 领域 ID 列表 | [optional] [default to undefined]
**code** | **string** | 稳定编码 | [optional] [default to undefined]
**name** | **string** | 领域名称 | [optional] [default to undefined]
**description** | **string** | 描述 | [optional] [default to undefined]
**status** | **string** | 启停状态 | [optional] [default to undefined]
**url** | **string** | 链接 | [optional] [default to undefined]
**icon** | **string** | 图标 | [optional] [default to undefined]
**is_nav** | **boolean** | 是否在导航中展示 | [optional] [default to undefined]

## Example

```typescript
import { ReqDomainQuery } from '@bass/bbs-sdk-axios';

const instance: ReqDomainQuery = {
    ids,
    code,
    name,
    description,
    status,
    url,
    icon,
    is_nav,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
