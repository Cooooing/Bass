# ArticleSave

文章保存内容。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **string** | 文章 ID。 | [optional] [default to undefined]
**title** | **string** | 标题。 | [default to undefined]
**content** | **string** | 原始内容。 | [default to undefined]
**reward_content** | **string** | 打赏后可见内容。 | [optional] [default to undefined]
**reward_points** | **number** | 打赏所需积分。 | [optional] [default to undefined]
**status** | **string** | 文章状态。 | [default to undefined]
**type** | **string** | 文章类型。 | [default to undefined]
**bounty_points** | **number** | 悬赏积分。 | [optional] [default to undefined]
**statement** | **string** | 文章声明。 | [optional] [default to undefined]
**commentable** | **boolean** | 是否允许评论。 | [optional] [default to undefined]
**anonymous** | **boolean** | 是否匿名展示。 | [optional] [default to undefined]
**listable** | **boolean** | 是否在列表中展示。 | [optional] [default to undefined]
**tags** | [**Array&lt;TagSave&gt;**](TagSave.md) | 绑定标签。 | [optional] [default to undefined]

## Example

```typescript
import { ArticleSave } from '@bass/bbs-sdk-axios';

const instance: ArticleSave = {
    id,
    title,
    content,
    reward_content,
    reward_points,
    status,
    type,
    bounty_points,
    statement,
    commentable,
    anonymous,
    listable,
    tags,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
