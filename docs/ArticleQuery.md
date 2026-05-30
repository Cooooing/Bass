# ArticleQuery

文章查询条件。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**tag_id** | **string** | 标签 ID。 | [optional] [default to undefined]
**domain_id** | **string** | 板块 ID。 | [optional] [default to undefined]
**status** | **string** | 文章状态。 | [optional] [default to undefined]
**type** | **string** | 文章类型。 | [optional] [default to undefined]
**order** | **string** | 排序方式。 | [optional] [default to undefined]
**keyword** | **string** | 搜索关键词。 | [optional] [default to undefined]
**author_id** | **string** | 作者账号 ID。 | [optional] [default to undefined]
**listable** | **boolean** | 是否在列表中展示。 | [optional] [default to undefined]

## Example

```typescript
import { ArticleQuery } from '@bass/bbs-sdk-axios';

const instance: ArticleQuery = {
    tag_id,
    domain_id,
    status,
    type,
    order,
    keyword,
    author_id,
    listable,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
