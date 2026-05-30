# GetCurrentTotpReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Totp** | Pointer to [**Totp**](Totp.md) | TOTP 认证状态。 | [optional] 

## Methods

### NewGetCurrentTotpReply

`func NewGetCurrentTotpReply() *GetCurrentTotpReply`

NewGetCurrentTotpReply instantiates a new GetCurrentTotpReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCurrentTotpReplyWithDefaults

`func NewGetCurrentTotpReplyWithDefaults() *GetCurrentTotpReply`

NewGetCurrentTotpReplyWithDefaults instantiates a new GetCurrentTotpReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotp

`func (o *GetCurrentTotpReply) GetTotp() Totp`

GetTotp returns the Totp field if non-nil, zero value otherwise.

### GetTotpOk

`func (o *GetCurrentTotpReply) GetTotpOk() (*Totp, bool)`

GetTotpOk returns a tuple with the Totp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotp

`func (o *GetCurrentTotpReply) SetTotp(v Totp)`

SetTotp sets Totp field to given value.

### HasTotp

`func (o *GetCurrentTotpReply) HasTotp() bool`

HasTotp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


