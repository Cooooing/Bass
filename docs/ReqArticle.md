# ReqArticle

协议消息。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**title** | **string** | 必填，标题 | [default to undefined]
**content** | **string** | 必填，正文 | [default to undefined]
**reward_content** | **string** | 可选，打赏区内容 | [optional] [default to undefined]
**reward_points** | **number** | 可选，打赏积分 | [optional] [default to undefined]
**type** | **string** | 必填，文章类型 | [default to undefined]
**statement** | **string** | 可选，创作声明 | [optional] [default to undefined]
**commentable** | **boolean** | 可选，是否允许评论 | [optional] [default to undefined]
**poll** | [**ArticlePollDraft**](ArticlePollDraft.md) | 可选，投票文章的草稿配置 | [optional] [default to undefined]
**lottery** | [**ArticleLotteryDraft**](ArticleLotteryDraft.md) | 可选，抽奖文章的草稿配置 | [optional] [default to undefined]

## Example

```typescript
import { ReqArticle } from '@bass/bbs-sdk-axios';

const instance: ReqArticle = {
    title,
    content,
    reward_content,
    reward_points,
    type,
    statement,
    commentable,
    poll,
    lottery,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
