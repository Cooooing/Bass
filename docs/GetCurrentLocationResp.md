# GetCurrentLocationResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to [**RespLocation**](RespLocation.md) |  | [optional] 

## Methods

### NewGetCurrentLocationResp

`func NewGetCurrentLocationResp() *GetCurrentLocationResp`

NewGetCurrentLocationResp instantiates a new GetCurrentLocationResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCurrentLocationRespWithDefaults

`func NewGetCurrentLocationRespWithDefaults() *GetCurrentLocationResp`

NewGetCurrentLocationRespWithDefaults instantiates a new GetCurrentLocationResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *GetCurrentLocationResp) GetLocation() RespLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GetCurrentLocationResp) GetLocationOk() (*RespLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GetCurrentLocationResp) SetLocation(v RespLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *GetCurrentLocationResp) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


