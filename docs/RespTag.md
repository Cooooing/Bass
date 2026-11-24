# RespTag

协议消息。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **string** | 标签 ID | [optional] [default to undefined]
**code** | **string** | 稳定编码 | [optional] [default to undefined]
**name** | **string** | 标签名称 | [optional] [default to undefined]
**description** | **string** | 描述 | [optional] [default to undefined]
**domain_id** | **string** | 领域 ID | [optional] [default to undefined]
**status** | **string** | 启停状态 | [optional] [default to undefined]
**icon** | **string** | 图标 | [optional] [default to undefined]
**sort** | **number** | 排序值，越小越靠前 | [optional] [default to undefined]
**article_count** | **number** | 文章数量 | [optional] [default to undefined]
**created_by** | **string** | 创建用户 ID | [optional] [default to undefined]
**updated_by** | **string** | 更新用户 ID | [optional] [default to undefined]
**created_at** | **string** | 创建时间 | [optional] [default to undefined]
**updated_at** | **string** | 更新时间 | [optional] [default to undefined]

## Example

```typescript
import { RespTag } from '@bass/bbs-sdk-axios';

const instance: RespTag = {
    id,
    code,
    name,
    description,
    domain_id,
    status,
    icon,
    sort,
    article_count,
    created_by,
    updated_by,
    created_at,
    updated_at,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
