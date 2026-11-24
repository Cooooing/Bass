# PrivateChatMember

协议消息。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**user_id** | **string** | 账号 ID。 | [optional] [default to undefined]
**name** | **string** | 用户名。 | [optional] [default to undefined]
**nickname** | **string** | 昵称。 | [optional] [default to undefined]
**avatar_url** | **string** | 头像地址。 | [optional] [default to undefined]
**online** | **boolean** | 是否存在有效实时连接。 | [optional] [default to undefined]

## Example

```typescript
import { PrivateChatMember } from '@bass/bbs-sdk-axios';

const instance: PrivateChatMember = {
    user_id,
    name,
    nickname,
    avatar_url,
    online,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
