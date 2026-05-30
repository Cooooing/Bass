# CommentQuery

评论查询条件。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**comment_id** | **string** | 评论 ID。 | [optional] [default to undefined]
**article_id** | **string** | 文章 ID。 | [optional] [default to undefined]
**parent_id** | **string** | 父评论 ID。 | [optional] [default to undefined]
**reply_id** | **string** | 回复的评论 ID。 | [optional] [default to undefined]
**order** | **string** | 排序方式。 | [optional] [default to undefined]
**user_id** | **string** | 评论账号 ID。 | [optional] [default to undefined]
**level** | **number** | 评论层级。 | [optional] [default to undefined]
**status** | **string** | 评论状态。 | [optional] [default to undefined]

## Example

```typescript
import { CommentQuery } from '@bass/bbs-sdk-axios';

const instance: CommentQuery = {
    comment_id,
    article_id,
    parent_id,
    reply_id,
    order,
    user_id,
    level,
    status,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
