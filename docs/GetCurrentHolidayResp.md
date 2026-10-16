# GetCurrentHolidayResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**date** | Option<**chrono::DateTime<chrono::FixedOffset>**> | 当前日期。 | [optional]
**r#type** | Option<**Type**> | 当天类型。 (enum: HOLIDAY_TYPE_UNSPECIFIED, HOLIDAY_TYPE_WORKDAY, HOLIDAY_TYPE_WEEKEND, HOLIDAY_TYPE_FESTIVAL, HOLIDAY_TYPE_MAKEUP_WORKDAY) | [optional]
**day_name** | Option<**String**> | 当天名称。 | [optional]
**holiday_name** | Option<**String**> | 下一个假期名称。 | [optional]
**days_to_holiday** | Option<**i32**> | 距离下一个假期的天数。 | [optional]
**days_remaining** | Option<**i32**> | 当前假期剩余天数。 | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


