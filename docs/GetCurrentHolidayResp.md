

# GetCurrentHolidayResp


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**date** | **OffsetDateTime** | 当前日期。 |  [optional] |
|**type** | [**TypeEnum**](#TypeEnum) | 当天类型。 |  [optional] |
|**dayName** | **String** | 当天名称。 |  [optional] |
|**holidayName** | **String** | 下一个假期名称。 |  [optional] |
|**daysToHoliday** | **Integer** | 距离下一个假期的天数。 |  [optional] |
|**daysRemaining** | **Integer** | 当前假期剩余天数。 |  [optional] |



## Enum: TypeEnum

| Name | Value |
|---- | -----|
| HOLIDAY_TYPE_UNSPECIFIED | &quot;HOLIDAY_TYPE_UNSPECIFIED&quot; |
| HOLIDAY_TYPE_WORKDAY | &quot;HOLIDAY_TYPE_WORKDAY&quot; |
| HOLIDAY_TYPE_WEEKEND | &quot;HOLIDAY_TYPE_WEEKEND&quot; |
| HOLIDAY_TYPE_FESTIVAL | &quot;HOLIDAY_TYPE_FESTIVAL&quot; |
| HOLIDAY_TYPE_MAKEUP_WORKDAY | &quot;HOLIDAY_TYPE_MAKEUP_WORKDAY&quot; |



