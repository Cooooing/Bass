# GetCurrentHolidayResp


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**date** | **string** | 当前日期。 | [optional] [default to undefined]
**type** | **string** | 当天类型。 | [optional] [default to undefined]
**day_name** | **string** | 当天名称。 | [optional] [default to undefined]
**holiday_name** | **string** | 下一个假期名称。 | [optional] [default to undefined]
**days_to_holiday** | **number** | 距离下一个假期的天数。 | [optional] [default to undefined]
**days_remaining** | **number** | 当前假期剩余天数。 | [optional] [default to undefined]

## Example

```typescript
import { GetCurrentHolidayResp } from '@bass/bbs-sdk-axios';

const instance: GetCurrentHolidayResp = {
    date,
    type,
    day_name,
    holiday_name,
    days_to_holiday,
    days_remaining,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
