# RespAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Basic** | Pointer to [**RespAccountBasic**](RespAccountBasic.md) |  | [optional] 
**Contact** | Pointer to [**RespAccountContact**](RespAccountContact.md) |  | [optional] 

## Methods

### NewRespAccount

`func NewRespAccount() *RespAccount`

NewRespAccount instantiates a new RespAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRespAccountWithDefaults

`func NewRespAccountWithDefaults() *RespAccount`

NewRespAccountWithDefaults instantiates a new RespAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasic

`func (o *RespAccount) GetBasic() RespAccountBasic`

GetBasic returns the Basic field if non-nil, zero value otherwise.

### GetBasicOk

`func (o *RespAccount) GetBasicOk() (*RespAccountBasic, bool)`

GetBasicOk returns a tuple with the Basic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasic

`func (o *RespAccount) SetBasic(v RespAccountBasic)`

SetBasic sets Basic field to given value.

### HasBasic

`func (o *RespAccount) HasBasic() bool`

HasBasic returns a boolean if a field has been set.

### GetContact

`func (o *RespAccount) GetContact() RespAccountContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *RespAccount) GetContactOk() (*RespAccountContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *RespAccount) SetContact(v RespAccountContact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *RespAccount) HasContact() bool`

HasContact returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


