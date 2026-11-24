# Profile

协议消息。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**account** | [**AccountProfile**](AccountProfile.md) | 协议字段。 | [optional] [default to undefined]
**background_url** | **string** | 协议字段。 | [optional] [default to undefined]
**location** | [**ProfileLocation**](ProfileLocation.md) | 协议字段。 | [optional] [default to undefined]
**last_success_login_at** | **string** | 协议字段。 | [optional] [default to undefined]
**visibility** | [**ProfileVisibility**](ProfileVisibility.md) | 协议字段。 | [optional] [default to undefined]
**viewer_relation** | [**ProfileRelation**](ProfileRelation.md) | 协议字段。 | [optional] [default to undefined]
**online** | **boolean** | 仅本人或允许公开在线状态的账号返回。 | [optional] [default to undefined]
**points** | **string** | 仅本人或允许公开积分的账号返回。 | [optional] [default to undefined]

## Example

```typescript
import { Profile } from '@bass/bbs-sdk-axios';

const instance: Profile = {
    account,
    background_url,
    location,
    last_success_login_at,
    visibility,
    viewer_relation,
    online,
    points,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
