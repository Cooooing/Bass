# UpdateCurrentPrivacySettingRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**public_points** | **boolean** | 是否公开积分。 | [optional] [default to undefined]
**public_followers** | **boolean** | 是否公开粉丝列表。 | [optional] [default to undefined]
**public_articles** | **boolean** | 是否公开文章列表。 | [optional] [default to undefined]
**public_comments** | **boolean** | 是否公开评论列表。 | [optional] [default to undefined]
**public_online_status** | **boolean** | 是否公开在线状态。 | [optional] [default to undefined]
**public_location** | **boolean** | 是否公开地理资料。 | [optional] [default to undefined]

## Example

```typescript
import { UpdateCurrentPrivacySettingRequest } from '@bass/bbs-sdk-axios';

const instance: UpdateCurrentPrivacySettingRequest = {
    public_points,
    public_followers,
    public_articles,
    public_comments,
    public_online_status,
    public_location,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
