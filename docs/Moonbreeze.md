# Moonbreeze

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**Author** | Pointer to [**AccountProfile**](AccountProfile.md) |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewMoonbreeze

`func NewMoonbreeze() *Moonbreeze`

NewMoonbreeze instantiates a new Moonbreeze object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoonbreezeWithDefaults

`func NewMoonbreezeWithDefaults() *Moonbreeze`

NewMoonbreezeWithDefaults instantiates a new Moonbreeze object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Moonbreeze) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Moonbreeze) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Moonbreeze) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Moonbreeze) HasId() bool`

HasId returns a boolean if a field has been set.

### GetContent

`func (o *Moonbreeze) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Moonbreeze) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Moonbreeze) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *Moonbreeze) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetAuthor

`func (o *Moonbreeze) GetAuthor() AccountProfile`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *Moonbreeze) GetAuthorOk() (*AccountProfile, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *Moonbreeze) SetAuthor(v AccountProfile)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *Moonbreeze) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetCity

`func (o *Moonbreeze) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Moonbreeze) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Moonbreeze) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Moonbreeze) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Moonbreeze) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Moonbreeze) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Moonbreeze) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Moonbreeze) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


