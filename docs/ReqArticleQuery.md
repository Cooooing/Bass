# ReqArticleQuery

协议消息。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**tag_id** | **string** | 可选，标签 ID | [optional] [default to undefined]
**domain_id** | **string** | 可选，领域 ID | [optional] [default to undefined]
**type** | **string** | 可选，文章类型 | [optional] [default to undefined]
**order** | **string** | 可选，排序方式 | [optional] [default to undefined]
**author_id** | **string** | 可选，作者账号 ID | [optional] [default to undefined]
**publish_status** | **string** | 可选，发布状态 | [optional] [default to undefined]
**publish_statuses** | **Array&lt;string&gt;** | 可选，发布状态列表 | [optional] [default to undefined]
**visibility** | **string** | 可选，可见范围 | [optional] [default to undefined]
**visibilities** | **Array&lt;string&gt;** | 可选，可见范围列表 | [optional] [default to undefined]
**restriction** | **string** | 可选，治理限制 | [optional] [default to undefined]
**restrictions** | **Array&lt;string&gt;** | 可选，治理限制列表 | [optional] [default to undefined]
**scheduled** | **boolean** | 可选，仅本人草稿：是否已设置发布时间 | [optional] [default to undefined]

## Example

```typescript
import { ReqArticleQuery } from '@bass/bbs-sdk-axios';

const instance: ReqArticleQuery = {
    tag_id,
    domain_id,
    type,
    order,
    author_id,
    publish_status,
    publish_statuses,
    visibility,
    visibilities,
    restriction,
    restrictions,
    scheduled,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
