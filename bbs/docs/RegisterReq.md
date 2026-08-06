# RegisterReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Name** | **string** |  | 
**Password** | **string** |  | 
**Nickname** | Pointer to **string** |  | [optional] 
**EmailCredential** | Pointer to [**ReqEmailCredential**](ReqEmailCredential.md) |  | [optional] 
**PhoneCredential** | Pointer to [**ReqPhoneCredential**](ReqPhoneCredential.md) |  | [optional] 

## Methods

### NewRegisterReq

`func NewRegisterReq(type_ string, name string, password string, ) *RegisterReq`

NewRegisterReq instantiates a new RegisterReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterReqWithDefaults

`func NewRegisterReqWithDefaults() *RegisterReq`

NewRegisterReqWithDefaults instantiates a new RegisterReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RegisterReq) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RegisterReq) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RegisterReq) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *RegisterReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegisterReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegisterReq) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *RegisterReq) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RegisterReq) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RegisterReq) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetNickname

`func (o *RegisterReq) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *RegisterReq) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *RegisterReq) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *RegisterReq) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetEmailCredential

`func (o *RegisterReq) GetEmailCredential() ReqEmailCredential`

GetEmailCredential returns the EmailCredential field if non-nil, zero value otherwise.

### GetEmailCredentialOk

`func (o *RegisterReq) GetEmailCredentialOk() (*ReqEmailCredential, bool)`

GetEmailCredentialOk returns a tuple with the EmailCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailCredential

`func (o *RegisterReq) SetEmailCredential(v ReqEmailCredential)`

SetEmailCredential sets EmailCredential field to given value.

### HasEmailCredential

`func (o *RegisterReq) HasEmailCredential() bool`

HasEmailCredential returns a boolean if a field has been set.

### GetPhoneCredential

`func (o *RegisterReq) GetPhoneCredential() ReqPhoneCredential`

GetPhoneCredential returns the PhoneCredential field if non-nil, zero value otherwise.

### GetPhoneCredentialOk

`func (o *RegisterReq) GetPhoneCredentialOk() (*ReqPhoneCredential, bool)`

GetPhoneCredentialOk returns a tuple with the PhoneCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneCredential

`func (o *RegisterReq) SetPhoneCredential(v ReqPhoneCredential)`

SetPhoneCredential sets PhoneCredential field to given value.

### HasPhoneCredential

`func (o *RegisterReq) HasPhoneCredential() bool`

HasPhoneCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


