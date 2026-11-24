# ArticleDetail

协议消息。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **string** | 文章 ID | [optional] [default to undefined]
**title** | **string** | 标题 | [optional] [default to undefined]
**content** | **string** | 原始内容 | [optional] [default to undefined]
**content_render** | **string** | 渲染内容 | [optional] [default to undefined]
**has_postscript** | **boolean** | 是否有附言 | [optional] [default to undefined]
**has_reward** | **boolean** | 是否有打赏区 | [optional] [default to undefined]
**reward_content** | **string** | 打赏后可见内容 | [optional] [default to undefined]
**reward_content_render** | **string** | 渲染打赏内容 | [optional] [default to undefined]
**reward_points** | **number** | 打赏积分 | [optional] [default to undefined]
**type** | **string** | 文章类型 | [optional] [default to undefined]
**statement** | **string** | 创作声明 | [optional] [default to undefined]
**commentable** | **boolean** | 是否允许评论 | [optional] [default to undefined]
**view_count** | **number** | 浏览数 | [optional] [default to undefined]
**thank_count** | **number** | 感谢数 | [optional] [default to undefined]
**like_count** | **number** | 点赞数 | [optional] [default to undefined]
**collect_count** | **number** | 收藏数 | [optional] [default to undefined]
**watch_count** | **number** | 关注数 | [optional] [default to undefined]
**reward_count** | **number** | 打赏数 | [optional] [default to undefined]
**reply_count** | **number** | 回复数 | [optional] [default to undefined]
**author_user** | [**AccountProfile**](AccountProfile.md) | 作者资料 | [optional] [default to undefined]
**last_reply_user** | [**AccountProfile**](AccountProfile.md) | 最后回复用户 | [optional] [default to undefined]
**last_reply_at** | **string** | 最后回复时间 | [optional] [default to undefined]
**cover_image_url** | **string** | 封面图片 URL | [optional] [default to undefined]
**viewer_action_state** | [**ArticleViewerActionState**](ArticleViewerActionState.md) | 当前用户互动状态 | [optional] [default to undefined]
**published_at** | **string** | 发布时间 | [optional] [default to undefined]
**postscripts** | [**Array&lt;ArticlePostscript&gt;**](ArticlePostscript.md) | 附言列表 | [optional] [default to undefined]
**publish_status** | **string** | 发布状态 | [optional] [default to undefined]
**visibility** | **string** | 可见范围 | [optional] [default to undefined]
**restriction** | **string** | 治理限制 | [optional] [default to undefined]
**edited_at** | **string** | 编辑时间 | [optional] [default to undefined]
**tags** | [**Array&lt;ArticleTag&gt;**](ArticleTag.md) | 文章标签 | [optional] [default to undefined]
**domains** | [**Array&lt;ArticleDomain&gt;**](ArticleDomain.md) | 文章所属领域 | [optional] [default to undefined]
**city** | **string** | 发布时城市快照 | [optional] [default to undefined]
**accepted_comment_id** | **string** | 问答已采纳评论 ID | [optional] [default to undefined]
**poll** | [**ArticlePoll**](ArticlePoll.md) | 投票详情 | [optional] [default to undefined]
**lottery** | [**ArticleLottery**](ArticleLottery.md) | 抽奖详情 | [optional] [default to undefined]
**created_by** | **string** | 创建账号 ID | [optional] [default to undefined]
**updated_by** | **string** | 更新账号 ID | [optional] [default to undefined]
**created_at** | **string** | 创建时间 | [optional] [default to undefined]
**updated_at** | **string** | 更新时间 | [optional] [default to undefined]

## Example

```typescript
import { ArticleDetail } from '@bass/bbs-sdk-axios';

const instance: ArticleDetail = {
    id,
    title,
    content,
    content_render,
    has_postscript,
    has_reward,
    reward_content,
    reward_content_render,
    reward_points,
    type,
    statement,
    commentable,
    view_count,
    thank_count,
    like_count,
    collect_count,
    watch_count,
    reward_count,
    reply_count,
    author_user,
    last_reply_user,
    last_reply_at,
    cover_image_url,
    viewer_action_state,
    published_at,
    postscripts,
    publish_status,
    visibility,
    restriction,
    edited_at,
    tags,
    domains,
    city,
    accepted_comment_id,
    poll,
    lottery,
    created_by,
    updated_by,
    created_at,
    updated_at,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
