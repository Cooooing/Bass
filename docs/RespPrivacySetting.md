# RespPrivacySetting

协议消息。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**user_id** | **string** | 账号 ID | [optional] [default to undefined]
**public_points** | **boolean** | 是否公开积分 | [optional] [default to undefined]
**public_follower_list** | **boolean** | 是否公开粉丝列表 | [optional] [default to undefined]
**public_following_list** | **boolean** | 是否公开关注列表 | [optional] [default to undefined]
**public_article_list** | **boolean** | 是否公开帖子列表 | [optional] [default to undefined]
**public_comment_list** | **boolean** | 是否公开回复列表 | [optional] [default to undefined]
**public_online_status** | **boolean** | 是否公开在线状态 | [optional] [default to undefined]
**public_location** | **boolean** | 是否公开地理资料 | [optional] [default to undefined]
**public_moonbreeze_list** | **boolean** | 是否公开清风明月列表 | [optional] [default to undefined]

## Example

```typescript
import { RespPrivacySetting } from '@bass/bbs-sdk-axios';

const instance: RespPrivacySetting = {
    user_id,
    public_points,
    public_follower_list,
    public_following_list,
    public_article_list,
    public_comment_list,
    public_online_status,
    public_location,
    public_moonbreeze_list,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
