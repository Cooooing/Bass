# StartPhoneRegistrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | **string** | 手机号。 | 
**Password** | **string** | 账号密码。 | 
**Name** | **string** | 账号名。 | 
**Nickname** | Pointer to **string** | 昵称。 | [optional] 

## Methods

### NewStartPhoneRegistrationRequest

`func NewStartPhoneRegistrationRequest(phone string, password string, name string, ) *StartPhoneRegistrationRequest`

NewStartPhoneRegistrationRequest instantiates a new StartPhoneRegistrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartPhoneRegistrationRequestWithDefaults

`func NewStartPhoneRegistrationRequestWithDefaults() *StartPhoneRegistrationRequest`

NewStartPhoneRegistrationRequestWithDefaults instantiates a new StartPhoneRegistrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *StartPhoneRegistrationRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *StartPhoneRegistrationRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *StartPhoneRegistrationRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetPassword

`func (o *StartPhoneRegistrationRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *StartPhoneRegistrationRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *StartPhoneRegistrationRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetName

`func (o *StartPhoneRegistrationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StartPhoneRegistrationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StartPhoneRegistrationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetNickname

`func (o *StartPhoneRegistrationRequest) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *StartPhoneRegistrationRequest) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *StartPhoneRegistrationRequest) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *StartPhoneRegistrationRequest) HasNickname() bool`

HasNickname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


