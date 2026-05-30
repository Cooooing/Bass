# Totp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** | 账号 ID。 | [optional] 
**Enable** | Pointer to **bool** | 是否启用 TOTP。 | [optional] 
**EnableTime** | Pointer to **string** | 启用时间。 | [optional] 

## Methods

### NewTotp

`func NewTotp() *Totp`

NewTotp instantiates a new Totp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTotpWithDefaults

`func NewTotpWithDefaults() *Totp`

NewTotpWithDefaults instantiates a new Totp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *Totp) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Totp) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Totp) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Totp) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetEnable

`func (o *Totp) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *Totp) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *Totp) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *Totp) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEnableTime

`func (o *Totp) GetEnableTime() string`

GetEnableTime returns the EnableTime field if non-nil, zero value otherwise.

### GetEnableTimeOk

`func (o *Totp) GetEnableTimeOk() (*string, bool)`

GetEnableTimeOk returns a tuple with the EnableTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTime

`func (o *Totp) SetEnableTime(v string)`

SetEnableTime sets EnableTime field to given value.

### HasEnableTime

`func (o *Totp) HasEnableTime() bool`

HasEnableTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


