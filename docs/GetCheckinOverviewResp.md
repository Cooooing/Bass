# GetCheckinOverviewResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Records** | Pointer to [**[]RespRecord**](RespRecord.md) |  | [optional] 
**CurrentStreak** | Pointer to **int32** |  | [optional] 
**LongestStreak** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetCheckinOverviewResp

`func NewGetCheckinOverviewResp() *GetCheckinOverviewResp`

NewGetCheckinOverviewResp instantiates a new GetCheckinOverviewResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCheckinOverviewRespWithDefaults

`func NewGetCheckinOverviewRespWithDefaults() *GetCheckinOverviewResp`

NewGetCheckinOverviewRespWithDefaults instantiates a new GetCheckinOverviewResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecords

`func (o *GetCheckinOverviewResp) GetRecords() []RespRecord`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *GetCheckinOverviewResp) GetRecordsOk() (*[]RespRecord, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *GetCheckinOverviewResp) SetRecords(v []RespRecord)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *GetCheckinOverviewResp) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetCurrentStreak

`func (o *GetCheckinOverviewResp) GetCurrentStreak() int32`

GetCurrentStreak returns the CurrentStreak field if non-nil, zero value otherwise.

### GetCurrentStreakOk

`func (o *GetCheckinOverviewResp) GetCurrentStreakOk() (*int32, bool)`

GetCurrentStreakOk returns a tuple with the CurrentStreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStreak

`func (o *GetCheckinOverviewResp) SetCurrentStreak(v int32)`

SetCurrentStreak sets CurrentStreak field to given value.

### HasCurrentStreak

`func (o *GetCheckinOverviewResp) HasCurrentStreak() bool`

HasCurrentStreak returns a boolean if a field has been set.

### GetLongestStreak

`func (o *GetCheckinOverviewResp) GetLongestStreak() int32`

GetLongestStreak returns the LongestStreak field if non-nil, zero value otherwise.

### GetLongestStreakOk

`func (o *GetCheckinOverviewResp) GetLongestStreakOk() (*int32, bool)`

GetLongestStreakOk returns a tuple with the LongestStreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongestStreak

`func (o *GetCheckinOverviewResp) SetLongestStreak(v int32)`

SetLongestStreak sets LongestStreak field to given value.

### HasLongestStreak

`func (o *GetCheckinOverviewResp) HasLongestStreak() bool`

HasLongestStreak returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


