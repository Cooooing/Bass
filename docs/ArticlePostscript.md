# ArticlePostscript

文章附言。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **string** | 附言 ID。 | [optional] [default to undefined]
**article_id** | **string** | 文章 ID。 | [optional] [default to undefined]
**content** | **string** | 原始内容。 | [optional] [default to undefined]
**content_render** | **string** | 渲染后的内容。 | [optional] [default to undefined]
**created_by** | **string** | 创建账号 ID。 | [optional] [default to undefined]
**updated_by** | **string** | 更新账号 ID。 | [optional] [default to undefined]
**created_at** | **string** | 创建时间。 | [optional] [default to undefined]
**updated_at** | **string** | 更新时间。 | [optional] [default to undefined]

## Example

```typescript
import { ArticlePostscript } from '@bass/bbs-sdk-axios';

const instance: ArticlePostscript = {
    id,
    article_id,
    content,
    content_render,
    created_by,
    updated_by,
    created_at,
    updated_at,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
