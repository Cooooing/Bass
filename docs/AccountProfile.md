# AccountProfile

账号展示资料。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **string** | 账号 ID。 | [optional] [default to undefined]
**name** | **string** | 账号名。 | [optional] [default to undefined]
**nickname** | **string** | 昵称。 | [optional] [default to undefined]
**url** | **string** | 个人主页 URL。 | [optional] [default to undefined]
**avatar_url** | **string** | 头像 URL。 | [optional] [default to undefined]
**introduction** | **string** | 个人简介。 | [optional] [default to undefined]
**mbti** | **string** | MBTI 类型。 | [optional] [default to undefined]
**status** | **string** | 账号状态。 | [optional] [default to undefined]
**follow_count** | **number** | 关注数量。 | [optional] [default to undefined]
**follower_count** | **number** | 粉丝数量。 | [optional] [default to undefined]
**created_at** | **string** | 创建时间。 | [optional] [default to undefined]
**updated_at** | **string** | 更新时间。 | [optional] [default to undefined]

## Example

```typescript
import { AccountProfile } from '@bass/bbs-sdk-axios';

const instance: AccountProfile = {
    id,
    name,
    nickname,
    url,
    avatar_url,
    introduction,
    mbti,
    status,
    follow_count,
    follower_count,
    created_at,
    updated_at,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
