# Comment

评论。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **string** | 评论 ID。 | [optional] [default to undefined]
**article_id** | **string** | 文章 ID。 | [optional] [default to undefined]
**content** | **string** | 原始内容。 | [optional] [default to undefined]
**content_render** | **string** | 渲染后的内容。 | [optional] [default to undefined]
**level** | **number** | 评论层级。 | [optional] [default to undefined]
**parent_id** | **string** | 父评论 ID。 | [optional] [default to undefined]
**reply_id** | **string** | 回复的评论 ID。 | [optional] [default to undefined]
**status** | **string** | 评论状态。 | [optional] [default to undefined]
**reply_count** | **number** | 回复数量。 | [optional] [default to undefined]
**like_count** | **number** | 点赞数量。 | [optional] [default to undefined]
**thank_count** | **number** | 感谢数量。 | [optional] [default to undefined]
**user** | [**AccountProfile**](AccountProfile.md) | 评论账号展示资料。 | [optional] [default to undefined]
**reply_user** | [**AccountProfile**](AccountProfile.md) | 被回复账号展示资料。 | [optional] [default to undefined]
**article** | [**Article**](Article.md) | 所属文章。 | [optional] [default to undefined]
**created_by** | **string** | 创建账号 ID。 | [optional] [default to undefined]
**updated_by** | **string** | 更新账号 ID。 | [optional] [default to undefined]
**created_at** | **string** | 创建时间。 | [optional] [default to undefined]
**updated_at** | **string** | 更新时间。 | [optional] [default to undefined]

## Example

```typescript
import { Comment } from '@bass/bbs-sdk-axios';

const instance: Comment = {
    id,
    article_id,
    content,
    content_render,
    level,
    parent_id,
    reply_id,
    status,
    reply_count,
    like_count,
    thank_count,
    user,
    reply_user,
    article,
    created_by,
    updated_by,
    created_at,
    updated_at,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
