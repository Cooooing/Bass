# CommentDetail

评论详情。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **string** |  | [optional] [default to undefined]
**article_id** | **string** |  | [optional] [default to undefined]
**content** | **string** |  | [optional] [default to undefined]
**content_render** | **string** |  | [optional] [default to undefined]
**level** | **number** |  | [optional] [default to undefined]
**parent_id** | **string** |  | [optional] [default to undefined]
**reply_id** | **string** |  | [optional] [default to undefined]
**reply_count** | **number** |  | [optional] [default to undefined]
**like_count** | **number** |  | [optional] [default to undefined]
**thank_count** | **number** |  | [optional] [default to undefined]
**user** | [**AccountProfile**](AccountProfile.md) |  | [optional] [default to undefined]
**reply_user** | [**AccountProfile**](AccountProfile.md) |  | [optional] [default to undefined]
**article** | [**ArticleBrief**](ArticleBrief.md) |  | [optional] [default to undefined]
**viewer_action_state** | [**CommentViewerActionState**](CommentViewerActionState.md) |  | [optional] [default to undefined]
**restriction** | **string** |  | [optional] [default to undefined]
**deleted_at** | **string** |  | [optional] [default to undefined]
**created_by** | **string** |  | [optional] [default to undefined]
**updated_by** | **string** |  | [optional] [default to undefined]
**created_at** | **string** |  | [optional] [default to undefined]
**updated_at** | **string** |  | [optional] [default to undefined]

## Example

```typescript
import { CommentDetail } from '@bass/bbs-sdk-axios';

const instance: CommentDetail = {
    id,
    article_id,
    content,
    content_render,
    level,
    parent_id,
    reply_id,
    reply_count,
    like_count,
    thank_count,
    user,
    reply_user,
    article,
    viewer_action_state,
    restriction,
    deleted_at,
    created_by,
    updated_by,
    created_at,
    updated_at,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
