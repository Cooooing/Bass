# UpdateCurrentPreferencesReq

请求参数。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**timezone** | **string** | 时区 | [optional] [default to undefined]
**theme** | **string** | 桌面端主题 | [optional] [default to undefined]
**mobile_theme** | **string** | 移动端主题 | [optional] [default to undefined]
**language** | **string** | 界面语言 | [optional] [default to undefined]

## Example

```typescript
import { UpdateCurrentPreferencesReq } from '@bass/bbs-sdk-axios';

const instance: UpdateCurrentPreferencesReq = {
    timezone,
    theme,
    mobile_theme,
    language,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
