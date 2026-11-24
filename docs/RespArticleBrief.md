# RespArticleBrief

协议消息。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **string** | 文章 ID | [optional] [default to undefined]
**title** | **string** | 标题 | [optional] [default to undefined]
**type** | **string** | 文章类型 | [optional] [default to undefined]
**author_user** | [**RespAccountProfile**](RespAccountProfile.md) | 作者展示资料 | [optional] [default to undefined]
**cover_image_url** | **string** | 封面图片 URL | [optional] [default to undefined]
**publish_status** | **string** | 发布状态 | [optional] [default to undefined]
**visibility** | **string** | 可见范围 | [optional] [default to undefined]
**restriction** | **string** | 管理限制 | [optional] [default to undefined]
**created_by** | **string** | 创建账号 ID | [optional] [default to undefined]
**updated_by** | **string** | 更新账号 ID | [optional] [default to undefined]
**created_at** | **string** | 创建时间 | [optional] [default to undefined]
**updated_at** | **string** | 更新时间 | [optional] [default to undefined]

## Example

```typescript
import { RespArticleBrief } from '@bass/bbs-sdk-axios';

const instance: RespArticleBrief = {
    id,
    title,
    type,
    author_user,
    cover_image_url,
    publish_status,
    visibility,
    restriction,
    created_by,
    updated_by,
    created_at,
    updated_at,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
