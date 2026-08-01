# RespDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**IsNav** | Pointer to **bool** |  | [optional] 
**Sort** | Pointer to **int32** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRespDomain

`func NewRespDomain() *RespDomain`

NewRespDomain instantiates a new RespDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRespDomainWithDefaults

`func NewRespDomainWithDefaults() *RespDomain`

NewRespDomainWithDefaults instantiates a new RespDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RespDomain) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RespDomain) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RespDomain) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RespDomain) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *RespDomain) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *RespDomain) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *RespDomain) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *RespDomain) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *RespDomain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RespDomain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RespDomain) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RespDomain) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *RespDomain) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RespDomain) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RespDomain) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RespDomain) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *RespDomain) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RespDomain) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RespDomain) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RespDomain) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUrl

`func (o *RespDomain) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RespDomain) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RespDomain) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *RespDomain) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetIcon

`func (o *RespDomain) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *RespDomain) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *RespDomain) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *RespDomain) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetIsNav

`func (o *RespDomain) GetIsNav() bool`

GetIsNav returns the IsNav field if non-nil, zero value otherwise.

### GetIsNavOk

`func (o *RespDomain) GetIsNavOk() (*bool, bool)`

GetIsNavOk returns a tuple with the IsNav field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNav

`func (o *RespDomain) SetIsNav(v bool)`

SetIsNav sets IsNav field to given value.

### HasIsNav

`func (o *RespDomain) HasIsNav() bool`

HasIsNav returns a boolean if a field has been set.

### GetSort

`func (o *RespDomain) GetSort() int32`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *RespDomain) GetSortOk() (*int32, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *RespDomain) SetSort(v int32)`

SetSort sets Sort field to given value.

### HasSort

`func (o *RespDomain) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetCreatedBy

`func (o *RespDomain) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RespDomain) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RespDomain) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *RespDomain) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *RespDomain) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *RespDomain) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *RespDomain) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *RespDomain) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RespDomain) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RespDomain) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RespDomain) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RespDomain) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RespDomain) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RespDomain) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RespDomain) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RespDomain) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


