# DomainQuery

内容板块查询条件。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ids** | **Array&lt;string&gt;** | 板块 ID 列表。 | [optional] [default to undefined]
**name** | **string** | 板块名称。 | [optional] [default to undefined]
**description** | **string** | 板块描述。 | [optional] [default to undefined]
**status** | **string** | 板块状态。 | [optional] [default to undefined]
**url** | **string** | 板块 URL。 | [optional] [default to undefined]
**icon** | **string** | 板块图标。 | [optional] [default to undefined]
**is_nav** | **boolean** | 是否在导航中展示。 | [optional] [default to undefined]

## Example

```typescript
import { DomainQuery } from '@bass/bbs-sdk-axios';

const instance: DomainQuery = {
    ids,
    name,
    description,
    status,
    url,
    icon,
    is_nav,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
