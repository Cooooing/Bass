# Domain

内容板块。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **string** | 板块 ID。 | [optional] [default to undefined]
**name** | **string** | 板块名称。 | [optional] [default to undefined]
**description** | **string** | 板块描述。 | [optional] [default to undefined]
**status** | **string** | 板块状态。 | [optional] [default to undefined]
**url** | **string** | 板块 URL。 | [optional] [default to undefined]
**icon** | **string** | 板块图标。 | [optional] [default to undefined]
**is_nav** | **boolean** | 是否在导航中展示。 | [optional] [default to undefined]
**created_by** | **string** | 创建账号 ID。 | [optional] [default to undefined]
**updated_by** | **string** | 更新账号 ID。 | [optional] [default to undefined]
**created_at** | **string** | 创建时间。 | [optional] [default to undefined]
**updated_at** | **string** | 更新时间。 | [optional] [default to undefined]

## Example

```typescript
import { Domain } from '@bass/bbs-sdk-axios';

const instance: Domain = {
    id,
    name,
    description,
    status,
    url,
    icon,
    is_nav,
    created_by,
    updated_by,
    created_at,
    updated_at,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
