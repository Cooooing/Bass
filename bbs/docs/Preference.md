# Preference

账号偏好设置。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**user_id** | **string** | 账号 ID。 | [optional] [default to undefined]
**timezone** | **string** | 时区。 | [optional] [default to undefined]
**theme** | **string** | 桌面端主题。 | [optional] [default to undefined]
**mobile_theme** | **string** | 移动端主题。 | [optional] [default to undefined]
**language** | **string** | 界面语言。 | [optional] [default to undefined]

## Example

```typescript
import { Preference } from '@bass/bbs-sdk-axios';

const instance: Preference = {
    user_id,
    timezone,
    theme,
    mobile_theme,
    language,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
