# StartEmailRegistrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | 邮箱地址。 | 
**Password** | **string** | 账号密码。 | 
**Name** | **string** | 账号名。 | 
**Nickname** | Pointer to **string** | 昵称。 | [optional] 

## Methods

### NewStartEmailRegistrationRequest

`func NewStartEmailRegistrationRequest(email string, password string, name string, ) *StartEmailRegistrationRequest`

NewStartEmailRegistrationRequest instantiates a new StartEmailRegistrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartEmailRegistrationRequestWithDefaults

`func NewStartEmailRegistrationRequestWithDefaults() *StartEmailRegistrationRequest`

NewStartEmailRegistrationRequestWithDefaults instantiates a new StartEmailRegistrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *StartEmailRegistrationRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *StartEmailRegistrationRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *StartEmailRegistrationRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *StartEmailRegistrationRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *StartEmailRegistrationRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *StartEmailRegistrationRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetName

`func (o *StartEmailRegistrationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StartEmailRegistrationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StartEmailRegistrationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetNickname

`func (o *StartEmailRegistrationRequest) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *StartEmailRegistrationRequest) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *StartEmailRegistrationRequest) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *StartEmailRegistrationRequest) HasNickname() bool`

HasNickname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


