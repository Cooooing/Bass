# LoginByPasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 

## Methods

### NewLoginByPasswordRequest

`func NewLoginByPasswordRequest() *LoginByPasswordRequest`

NewLoginByPasswordRequest instantiates a new LoginByPasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginByPasswordRequestWithDefaults

`func NewLoginByPasswordRequestWithDefaults() *LoginByPasswordRequest`

NewLoginByPasswordRequestWithDefaults instantiates a new LoginByPasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *LoginByPasswordRequest) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *LoginByPasswordRequest) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *LoginByPasswordRequest) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *LoginByPasswordRequest) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetPassword

`func (o *LoginByPasswordRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *LoginByPasswordRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *LoginByPasswordRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *LoginByPasswordRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


