# RespCurrentAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | Pointer to [**AccountProfile**](AccountProfile.md) |  | [optional] 
**Contact** | Pointer to [**AccountContact**](AccountContact.md) |  | [optional] 

## Methods

### NewRespCurrentAccount

`func NewRespCurrentAccount() *RespCurrentAccount`

NewRespCurrentAccount instantiates a new RespCurrentAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRespCurrentAccountWithDefaults

`func NewRespCurrentAccountWithDefaults() *RespCurrentAccount`

NewRespCurrentAccountWithDefaults instantiates a new RespCurrentAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfile

`func (o *RespCurrentAccount) GetProfile() AccountProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *RespCurrentAccount) GetProfileOk() (*AccountProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *RespCurrentAccount) SetProfile(v AccountProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *RespCurrentAccount) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetContact

`func (o *RespCurrentAccount) GetContact() AccountContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *RespCurrentAccount) GetContactOk() (*AccountContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *RespCurrentAccount) SetContact(v AccountContact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *RespCurrentAccount) HasContact() bool`

HasContact returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


