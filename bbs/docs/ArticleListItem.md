# ArticleListItem

文章列表项。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **string** |  | [optional] [default to undefined]
**title** | **string** |  | [optional] [default to undefined]
**content** | **string** |  | [optional] [default to undefined]
**content_render** | **string** |  | [optional] [default to undefined]
**has_postscript** | **boolean** |  | [optional] [default to undefined]
**has_reward** | **boolean** |  | [optional] [default to undefined]
**type** | **string** |  | [optional] [default to undefined]
**statement** | **string** |  | [optional] [default to undefined]
**commentable** | **boolean** |  | [optional] [default to undefined]
**anonymous** | **boolean** |  | [optional] [default to undefined]
**view_count** | **number** |  | [optional] [default to undefined]
**thank_count** | **number** |  | [optional] [default to undefined]
**like_count** | **number** |  | [optional] [default to undefined]
**collect_count** | **number** |  | [optional] [default to undefined]
**watch_count** | **number** |  | [optional] [default to undefined]
**reply_count** | **number** |  | [optional] [default to undefined]
**bounty_points** | **number** |  | [optional] [default to undefined]
**accepted_answer_id** | **string** |  | [optional] [default to undefined]
**author_user** | [**AccountProfile**](AccountProfile.md) |  | [optional] [default to undefined]
**last_reply_user** | [**AccountProfile**](AccountProfile.md) |  | [optional] [default to undefined]
**last_reply_at** | **string** |  | [optional] [default to undefined]
**cover_image_url** | **string** |  | [optional] [default to undefined]
**viewer_action_state** | [**ArticleViewerActionState**](ArticleViewerActionState.md) |  | [optional] [default to undefined]
**published_at** | **string** |  | [optional] [default to undefined]
**publish_status** | **string** |  | [optional] [default to undefined]
**visibility** | **string** |  | [optional] [default to undefined]
**restriction** | **string** |  | [optional] [default to undefined]
**edited_at** | **string** |  | [optional] [default to undefined]
**created_by** | **string** |  | [optional] [default to undefined]
**updated_by** | **string** |  | [optional] [default to undefined]
**created_at** | **string** |  | [optional] [default to undefined]
**updated_at** | **string** |  | [optional] [default to undefined]

## Example

```typescript
import { ArticleListItem } from '@bass/bbs-sdk-axios';

const instance: ArticleListItem = {
    id,
    title,
    content,
    content_render,
    has_postscript,
    has_reward,
    type,
    statement,
    commentable,
    anonymous,
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
    viewer_action_state,
    published_at,
    publish_status,
    visibility,
    restriction,
    edited_at,
    created_by,
    updated_by,
    created_at,
    updated_at,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
