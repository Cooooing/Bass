# ReqCommentQuery

协议消息。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**comment_id** | **string** | 评论 ID | [optional] [default to undefined]
**article_id** | **string** | 文章 ID | [optional] [default to undefined]
**parent_id** | **string** | 父评论 ID | [optional] [default to undefined]
**reply_id** | **string** | 回复的评论 ID | [optional] [default to undefined]
**order** | **string** | 排序方式 | [optional] [default to undefined]
**user_id** | **string** | 评论账号 ID | [optional] [default to undefined]
**level** | **number** | 评论层级 | [optional] [default to undefined]
**restriction** | **string** | 管理限制 | [optional] [default to undefined]
**restrictions** | **Array&lt;string&gt;** | 管理限制列表 | [optional] [default to undefined]

## Example

```typescript
import { ReqCommentQuery } from '@bass/bbs-sdk-axios';

const instance: ReqCommentQuery = {
    comment_id,
    article_id,
    parent_id,
    reply_id,
    order,
    user_id,
    level,
    restriction,
    restrictions,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
