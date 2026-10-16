# GetCurrentHolidayResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **time.Time** | 当前日期。 | [optional] 
**Type** | Pointer to **string** | 当天类型。 | [optional] 
**DayName** | Pointer to **string** | 当天名称。 | [optional] 
**HolidayName** | Pointer to **string** | 下一个假期名称。 | [optional] 
**DaysToHoliday** | Pointer to **int32** | 距离下一个假期的天数。 | [optional] 
**DaysRemaining** | Pointer to **int32** | 当前假期剩余天数。 | [optional] 

## Methods

### NewGetCurrentHolidayResp

`func NewGetCurrentHolidayResp() *GetCurrentHolidayResp`

NewGetCurrentHolidayResp instantiates a new GetCurrentHolidayResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCurrentHolidayRespWithDefaults

`func NewGetCurrentHolidayRespWithDefaults() *GetCurrentHolidayResp`

NewGetCurrentHolidayRespWithDefaults instantiates a new GetCurrentHolidayResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *GetCurrentHolidayResp) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GetCurrentHolidayResp) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GetCurrentHolidayResp) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *GetCurrentHolidayResp) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetType

`func (o *GetCurrentHolidayResp) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetCurrentHolidayResp) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetCurrentHolidayResp) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetCurrentHolidayResp) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDayName

`func (o *GetCurrentHolidayResp) GetDayName() string`

GetDayName returns the DayName field if non-nil, zero value otherwise.

### GetDayNameOk

`func (o *GetCurrentHolidayResp) GetDayNameOk() (*string, bool)`

GetDayNameOk returns a tuple with the DayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayName

`func (o *GetCurrentHolidayResp) SetDayName(v string)`

SetDayName sets DayName field to given value.

### HasDayName

`func (o *GetCurrentHolidayResp) HasDayName() bool`

HasDayName returns a boolean if a field has been set.

### GetHolidayName

`func (o *GetCurrentHolidayResp) GetHolidayName() string`

GetHolidayName returns the HolidayName field if non-nil, zero value otherwise.

### GetHolidayNameOk

`func (o *GetCurrentHolidayResp) GetHolidayNameOk() (*string, bool)`

GetHolidayNameOk returns a tuple with the HolidayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolidayName

`func (o *GetCurrentHolidayResp) SetHolidayName(v string)`

SetHolidayName sets HolidayName field to given value.

### HasHolidayName

`func (o *GetCurrentHolidayResp) HasHolidayName() bool`

HasHolidayName returns a boolean if a field has been set.

### GetDaysToHoliday

`func (o *GetCurrentHolidayResp) GetDaysToHoliday() int32`

GetDaysToHoliday returns the DaysToHoliday field if non-nil, zero value otherwise.

### GetDaysToHolidayOk

`func (o *GetCurrentHolidayResp) GetDaysToHolidayOk() (*int32, bool)`

GetDaysToHolidayOk returns a tuple with the DaysToHoliday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysToHoliday

`func (o *GetCurrentHolidayResp) SetDaysToHoliday(v int32)`

SetDaysToHoliday sets DaysToHoliday field to given value.

### HasDaysToHoliday

`func (o *GetCurrentHolidayResp) HasDaysToHoliday() bool`

HasDaysToHoliday returns a boolean if a field has been set.

### GetDaysRemaining

`func (o *GetCurrentHolidayResp) GetDaysRemaining() int32`

GetDaysRemaining returns the DaysRemaining field if non-nil, zero value otherwise.

### GetDaysRemainingOk

`func (o *GetCurrentHolidayResp) GetDaysRemainingOk() (*int32, bool)`

GetDaysRemainingOk returns a tuple with the DaysRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysRemaining

`func (o *GetCurrentHolidayResp) SetDaysRemaining(v int32)`

SetDaysRemaining sets DaysRemaining field to given value.

### HasDaysRemaining

`func (o *GetCurrentHolidayResp) HasDaysRemaining() bool`

HasDaysRemaining returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


