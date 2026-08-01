# RespTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DomainId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**Sort** | Pointer to **int32** |  | [optional] 
**ArticleCount** | Pointer to **int32** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRespTag

`func NewRespTag() *RespTag`

NewRespTag instantiates a new RespTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRespTagWithDefaults

`func NewRespTagWithDefaults() *RespTag`

NewRespTagWithDefaults instantiates a new RespTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RespTag) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RespTag) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RespTag) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RespTag) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *RespTag) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *RespTag) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *RespTag) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *RespTag) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *RespTag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RespTag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RespTag) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RespTag) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *RespTag) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RespTag) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RespTag) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RespTag) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDomainId

`func (o *RespTag) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *RespTag) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *RespTag) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *RespTag) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetStatus

`func (o *RespTag) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RespTag) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RespTag) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RespTag) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIcon

`func (o *RespTag) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *RespTag) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *RespTag) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *RespTag) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetSort

`func (o *RespTag) GetSort() int32`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *RespTag) GetSortOk() (*int32, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *RespTag) SetSort(v int32)`

SetSort sets Sort field to given value.

### HasSort

`func (o *RespTag) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetArticleCount

`func (o *RespTag) GetArticleCount() int32`

GetArticleCount returns the ArticleCount field if non-nil, zero value otherwise.

### GetArticleCountOk

`func (o *RespTag) GetArticleCountOk() (*int32, bool)`

GetArticleCountOk returns a tuple with the ArticleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticleCount

`func (o *RespTag) SetArticleCount(v int32)`

SetArticleCount sets ArticleCount field to given value.

### HasArticleCount

`func (o *RespTag) HasArticleCount() bool`

HasArticleCount returns a boolean if a field has been set.

### GetCreatedBy

`func (o *RespTag) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RespTag) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RespTag) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *RespTag) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *RespTag) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *RespTag) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *RespTag) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *RespTag) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RespTag) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RespTag) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RespTag) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RespTag) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RespTag) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RespTag) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RespTag) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RespTag) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


