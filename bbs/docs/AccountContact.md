# AccountContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** | 账号 ID。 | [optional] 
**Email** | Pointer to **string** | 邮箱地址。 | [optional] 
**Phone** | Pointer to **string** | 手机号。 | [optional] 

## Methods

### NewAccountContact

`func NewAccountContact() *AccountContact`

NewAccountContact instantiates a new AccountContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountContactWithDefaults

`func NewAccountContactWithDefaults() *AccountContact`

NewAccountContactWithDefaults instantiates a new AccountContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *AccountContact) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AccountContact) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AccountContact) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AccountContact) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetEmail

`func (o *AccountContact) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AccountContact) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AccountContact) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AccountContact) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *AccountContact) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *AccountContact) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *AccountContact) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *AccountContact) HasPhone() bool`

HasPhone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


