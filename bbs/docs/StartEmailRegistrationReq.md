# StartEmailRegistrationReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Password** | **string** |  | 
**Name** | **string** |  | 
**Nickname** | Pointer to **string** |  | [optional] 

## Methods

### NewStartEmailRegistrationReq

`func NewStartEmailRegistrationReq(email string, password string, name string, ) *StartEmailRegistrationReq`

NewStartEmailRegistrationReq instantiates a new StartEmailRegistrationReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartEmailRegistrationReqWithDefaults

`func NewStartEmailRegistrationReqWithDefaults() *StartEmailRegistrationReq`

NewStartEmailRegistrationReqWithDefaults instantiates a new StartEmailRegistrationReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *StartEmailRegistrationReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *StartEmailRegistrationReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *StartEmailRegistrationReq) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *StartEmailRegistrationReq) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *StartEmailRegistrationReq) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *StartEmailRegistrationReq) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetName

`func (o *StartEmailRegistrationReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StartEmailRegistrationReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StartEmailRegistrationReq) SetName(v string)`

SetName sets Name field to given value.


### GetNickname

`func (o *StartEmailRegistrationReq) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *StartEmailRegistrationReq) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *StartEmailRegistrationReq) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *StartEmailRegistrationReq) HasNickname() bool`

HasNickname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


