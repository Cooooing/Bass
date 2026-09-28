# GetCurrentTotpResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Totp** | Pointer to [**RespTotp**](RespTotp.md) |  | [optional] 

## Methods

### NewGetCurrentTotpResp

`func NewGetCurrentTotpResp() *GetCurrentTotpResp`

NewGetCurrentTotpResp instantiates a new GetCurrentTotpResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCurrentTotpRespWithDefaults

`func NewGetCurrentTotpRespWithDefaults() *GetCurrentTotpResp`

NewGetCurrentTotpRespWithDefaults instantiates a new GetCurrentTotpResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotp

`func (o *GetCurrentTotpResp) GetTotp() RespTotp`

GetTotp returns the Totp field if non-nil, zero value otherwise.

### GetTotpOk

`func (o *GetCurrentTotpResp) GetTotpOk() (*RespTotp, bool)`

GetTotpOk returns a tuple with the Totp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotp

`func (o *GetCurrentTotpResp) SetTotp(v RespTotp)`

SetTotp sets Totp field to given value.

### HasTotp

`func (o *GetCurrentTotpResp) HasTotp() bool`

HasTotp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


