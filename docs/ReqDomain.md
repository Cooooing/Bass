# ReqDomain

协议消息。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**code** | **string** | 稳定编码 | [default to undefined]
**name** | **string** | 领域名称 | [default to undefined]
**description** | **string** | 描述 | [optional] [default to undefined]
**status** | **string** | 启停状态 | [optional] [default to undefined]
**url** | **string** | 链接 | [optional] [default to undefined]
**icon** | **string** | 图标 | [optional] [default to undefined]
**is_nav** | **boolean** | 是否在导航中展示 | [optional] [default to undefined]
**sort** | **number** | 排序值，越小越靠前 | [optional] [default to undefined]

## Example

```typescript
import { ReqDomain } from '@bass/bbs-sdk-axios';

const instance: ReqDomain = {
    code,
    name,
    description,
    status,
    url,
    icon,
    is_nav,
    sort,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
