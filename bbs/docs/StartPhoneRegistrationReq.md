# StartPhoneRegistrationReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | **string** |  | 
**Password** | **string** |  | 
**Name** | **string** |  | 
**Nickname** | Pointer to **string** |  | [optional] 

## Methods

### NewStartPhoneRegistrationReq

`func NewStartPhoneRegistrationReq(phone string, password string, name string, ) *StartPhoneRegistrationReq`

NewStartPhoneRegistrationReq instantiates a new StartPhoneRegistrationReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartPhoneRegistrationReqWithDefaults

`func NewStartPhoneRegistrationReqWithDefaults() *StartPhoneRegistrationReq`

NewStartPhoneRegistrationReqWithDefaults instantiates a new StartPhoneRegistrationReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *StartPhoneRegistrationReq) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *StartPhoneRegistrationReq) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *StartPhoneRegistrationReq) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetPassword

`func (o *StartPhoneRegistrationReq) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *StartPhoneRegistrationReq) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *StartPhoneRegistrationReq) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetName

`func (o *StartPhoneRegistrationReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StartPhoneRegistrationReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StartPhoneRegistrationReq) SetName(v string)`

SetName sets Name field to given value.


### GetNickname

`func (o *StartPhoneRegistrationReq) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *StartPhoneRegistrationReq) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *StartPhoneRegistrationReq) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *StartPhoneRegistrationReq) HasNickname() bool`

HasNickname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


