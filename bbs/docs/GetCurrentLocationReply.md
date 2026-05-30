# GetCurrentLocationReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to [**Location**](Location.md) | 当前账号地理资料。 | [optional] 

## Methods

### NewGetCurrentLocationReply

`func NewGetCurrentLocationReply() *GetCurrentLocationReply`

NewGetCurrentLocationReply instantiates a new GetCurrentLocationReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCurrentLocationReplyWithDefaults

`func NewGetCurrentLocationReplyWithDefaults() *GetCurrentLocationReply`

NewGetCurrentLocationReplyWithDefaults instantiates a new GetCurrentLocationReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *GetCurrentLocationReply) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GetCurrentLocationReply) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GetCurrentLocationReply) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *GetCurrentLocationReply) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


