# LoginReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**PasswordCredential** | Pointer to [**ReqPasswordCredential**](ReqPasswordCredential.md) |  | [optional] 
**EmailCredential** | Pointer to [**ReqEmailCredential**](ReqEmailCredential.md) |  | [optional] 
**PhoneCredential** | Pointer to [**ReqPhoneCredential**](ReqPhoneCredential.md) |  | [optional] 

## Methods

### NewLoginReq

`func NewLoginReq(type_ string, ) *LoginReq`

NewLoginReq instantiates a new LoginReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginReqWithDefaults

`func NewLoginReqWithDefaults() *LoginReq`

NewLoginReqWithDefaults instantiates a new LoginReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LoginReq) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LoginReq) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LoginReq) SetType(v string)`

SetType sets Type field to given value.


### GetPasswordCredential

`func (o *LoginReq) GetPasswordCredential() ReqPasswordCredential`

GetPasswordCredential returns the PasswordCredential field if non-nil, zero value otherwise.

### GetPasswordCredentialOk

`func (o *LoginReq) GetPasswordCredentialOk() (*ReqPasswordCredential, bool)`

GetPasswordCredentialOk returns a tuple with the PasswordCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordCredential

`func (o *LoginReq) SetPasswordCredential(v ReqPasswordCredential)`

SetPasswordCredential sets PasswordCredential field to given value.

### HasPasswordCredential

`func (o *LoginReq) HasPasswordCredential() bool`

HasPasswordCredential returns a boolean if a field has been set.

### GetEmailCredential

`func (o *LoginReq) GetEmailCredential() ReqEmailCredential`

GetEmailCredential returns the EmailCredential field if non-nil, zero value otherwise.

### GetEmailCredentialOk

`func (o *LoginReq) GetEmailCredentialOk() (*ReqEmailCredential, bool)`

GetEmailCredentialOk returns a tuple with the EmailCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailCredential

`func (o *LoginReq) SetEmailCredential(v ReqEmailCredential)`

SetEmailCredential sets EmailCredential field to given value.

### HasEmailCredential

`func (o *LoginReq) HasEmailCredential() bool`

HasEmailCredential returns a boolean if a field has been set.

### GetPhoneCredential

`func (o *LoginReq) GetPhoneCredential() ReqPhoneCredential`

GetPhoneCredential returns the PhoneCredential field if non-nil, zero value otherwise.

### GetPhoneCredentialOk

`func (o *LoginReq) GetPhoneCredentialOk() (*ReqPhoneCredential, bool)`

GetPhoneCredentialOk returns a tuple with the PhoneCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneCredential

`func (o *LoginReq) SetPhoneCredential(v ReqPhoneCredential)`

SetPhoneCredential sets PhoneCredential field to given value.

### HasPhoneCredential

`func (o *LoginReq) HasPhoneCredential() bool`

HasPhoneCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


