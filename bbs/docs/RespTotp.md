# RespTotp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** |  | [optional] 
**Enable** | Pointer to **bool** |  | [optional] 
**EnableTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRespTotp

`func NewRespTotp() *RespTotp`

NewRespTotp instantiates a new RespTotp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRespTotpWithDefaults

`func NewRespTotpWithDefaults() *RespTotp`

NewRespTotpWithDefaults instantiates a new RespTotp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *RespTotp) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *RespTotp) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *RespTotp) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *RespTotp) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetEnable

`func (o *RespTotp) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *RespTotp) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *RespTotp) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *RespTotp) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEnableTime

`func (o *RespTotp) GetEnableTime() time.Time`

GetEnableTime returns the EnableTime field if non-nil, zero value otherwise.

### GetEnableTimeOk

`func (o *RespTotp) GetEnableTimeOk() (*time.Time, bool)`

GetEnableTimeOk returns a tuple with the EnableTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTime

`func (o *RespTotp) SetEnableTime(v time.Time)`

SetEnableTime sets EnableTime field to given value.

### HasEnableTime

`func (o *RespTotp) HasEnableTime() bool`

HasEnableTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


