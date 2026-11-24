# UpdateProfileAccountReq

请求参数。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**nickname** | **string** | 昵称，选填 | [optional] [default to undefined]
**url** | **string** | 个人主页 URL，选填 | [optional] [default to undefined]
**introduction** | **string** | 个人简介，选填 | [optional] [default to undefined]
**mbti** | **string** | MBTI 类型，选填；空字符串表示清除 | [optional] [default to undefined]

## Example

```typescript
import { UpdateProfileAccountReq } from '@bass/bbs-sdk-axios';

const instance: UpdateProfileAccountReq = {
    nickname,
    url,
    introduction,
    mbti,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
