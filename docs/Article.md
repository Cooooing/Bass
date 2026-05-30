# Article

文章。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **string** | 文章 ID。 | [optional] [default to undefined]
**title** | **string** | 标题。 | [optional] [default to undefined]
**content** | **string** | 原始内容。 | [optional] [default to undefined]
**content_render** | **string** | 渲染后的内容。 | [optional] [default to undefined]
**has_postscript** | **boolean** | 是否有附言。 | [optional] [default to undefined]
**has_reward** | **boolean** | 是否有打赏内容。 | [optional] [default to undefined]
**reward_content** | **string** | 打赏后可见内容。 | [optional] [default to undefined]
**reward_content_render** | **string** | 渲染后的打赏内容。 | [optional] [default to undefined]
**reward_points** | **number** | 打赏所需积分。 | [optional] [default to undefined]
**status** | **string** | 文章状态。 | [optional] [default to undefined]
**type** | **string** | 文章类型。 | [optional] [default to undefined]
**statement** | **string** | 文章声明。 | [optional] [default to undefined]
**commentable** | **boolean** | 是否允许评论。 | [optional] [default to undefined]
**anonymous** | **boolean** | 是否匿名展示。 | [optional] [default to undefined]
**listable** | **boolean** | 是否在列表中展示。 | [optional] [default to undefined]
**view_count** | **number** | 浏览数量。 | [optional] [default to undefined]
**thank_count** | **number** | 感谢数量。 | [optional] [default to undefined]
**like_count** | **number** | 点赞数量。 | [optional] [default to undefined]
**collect_count** | **number** | 收藏数量。 | [optional] [default to undefined]
**watch_count** | **number** | 关注数量。 | [optional] [default to undefined]
**reply_count** | **number** | 回复数量。 | [optional] [default to undefined]
**bounty_points** | **number** | 悬赏积分。 | [optional] [default to undefined]
**accepted_answer_id** | **string** | 已采纳答案评论 ID。 | [optional] [default to undefined]
**author_user** | [**AccountProfile**](AccountProfile.md) | 作者账号展示资料。 | [optional] [default to undefined]
**last_reply_user** | [**AccountProfile**](AccountProfile.md) | 最后回复账号展示资料。 | [optional] [default to undefined]
**last_reply_at** | **string** | 最后回复时间。 | [optional] [default to undefined]
**cover_image_url** | **string** | 封面图片 URL。 | [optional] [default to undefined]
**created_by** | **string** | 创建账号 ID。 | [optional] [default to undefined]
**updated_by** | **string** | 更新账号 ID。 | [optional] [default to undefined]
**created_at** | **string** | 创建时间。 | [optional] [default to undefined]
**updated_at** | **string** | 更新时间。 | [optional] [default to undefined]

## Example

```typescript
import { Article } from '@bass/bbs-sdk-axios';

const instance: Article = {
    id,
    title,
    content,
    content_render,
    has_postscript,
    has_reward,
    reward_content,
    reward_content_render,
    reward_points,
    status,
    type,
    statement,
    commentable,
    anonymous,
    listable,
    view_count,
    thank_count,
    like_count,
    collect_count,
    watch_count,
    reply_count,
    bounty_points,
    accepted_answer_id,
    author_user,
    last_reply_user,
    last_reply_at,
    cover_image_url,
    created_by,
    updated_by,
    created_at,
    updated_at,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
