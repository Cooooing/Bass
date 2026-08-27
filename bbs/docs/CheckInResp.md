# CheckInResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordId** | Pointer to **string** |  | [optional] 
**Date** | Pointer to **time.Time** |  | [optional] 
**CurrentStreak** | Pointer to **int32** |  | [optional] 
**LongestStreak** | Pointer to **int32** |  | [optional] 

## Methods

### NewCheckInResp

`func NewCheckInResp() *CheckInResp`

NewCheckInResp instantiates a new CheckInResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckInRespWithDefaults

`func NewCheckInRespWithDefaults() *CheckInResp`

NewCheckInRespWithDefaults instantiates a new CheckInResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordId

`func (o *CheckInResp) GetRecordId() string`

GetRecordId returns the RecordId field if non-nil, zero value otherwise.

### GetRecordIdOk

`func (o *CheckInResp) GetRecordIdOk() (*string, bool)`

GetRecordIdOk returns a tuple with the RecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordId

`func (o *CheckInResp) SetRecordId(v string)`

SetRecordId sets RecordId field to given value.

### HasRecordId

`func (o *CheckInResp) HasRecordId() bool`

HasRecordId returns a boolean if a field has been set.

### GetDate

`func (o *CheckInResp) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CheckInResp) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CheckInResp) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *CheckInResp) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetCurrentStreak

`func (o *CheckInResp) GetCurrentStreak() int32`

GetCurrentStreak returns the CurrentStreak field if non-nil, zero value otherwise.

### GetCurrentStreakOk

`func (o *CheckInResp) GetCurrentStreakOk() (*int32, bool)`

GetCurrentStreakOk returns a tuple with the CurrentStreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStreak

`func (o *CheckInResp) SetCurrentStreak(v int32)`

SetCurrentStreak sets CurrentStreak field to given value.

### HasCurrentStreak

`func (o *CheckInResp) HasCurrentStreak() bool`

HasCurrentStreak returns a boolean if a field has been set.

### GetLongestStreak

`func (o *CheckInResp) GetLongestStreak() int32`

GetLongestStreak returns the LongestStreak field if non-nil, zero value otherwise.

### GetLongestStreakOk

`func (o *CheckInResp) GetLongestStreakOk() (*int32, bool)`

GetLongestStreakOk returns a tuple with the LongestStreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongestStreak

`func (o *CheckInResp) SetLongestStreak(v int32)`

SetLongestStreak sets LongestStreak field to given value.

### HasLongestStreak

`func (o *CheckInResp) HasLongestStreak() bool`

HasLongestStreak returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


