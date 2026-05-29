# LoginByPasswordReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** |  | [optional] 
**Account** | Pointer to [**Account**](Account.md) |  | [optional] 

## Methods

### NewLoginByPasswordReply

`func NewLoginByPasswordReply() *LoginByPasswordReply`

NewLoginByPasswordReply instantiates a new LoginByPasswordReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginByPasswordReplyWithDefaults

`func NewLoginByPasswordReplyWithDefaults() *LoginByPasswordReply`

NewLoginByPasswordReplyWithDefaults instantiates a new LoginByPasswordReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *LoginByPasswordReply) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LoginByPasswordReply) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LoginByPasswordReply) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *LoginByPasswordReply) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetAccount

`func (o *LoginByPasswordReply) GetAccount() Account`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *LoginByPasswordReply) GetAccountOk() (*Account, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *LoginByPasswordReply) SetAccount(v Account)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *LoginByPasswordReply) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


