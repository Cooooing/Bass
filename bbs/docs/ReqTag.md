# ReqTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**DomainId** | **string** |  | 
**Status** | Pointer to **string** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**Sort** | Pointer to **int32** |  | [optional] 

## Methods

### NewReqTag

`func NewReqTag(code string, name string, domainId string, ) *ReqTag`

NewReqTag instantiates a new ReqTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReqTagWithDefaults

`func NewReqTagWithDefaults() *ReqTag`

NewReqTagWithDefaults instantiates a new ReqTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ReqTag) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ReqTag) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ReqTag) SetCode(v string)`

SetCode sets Code field to given value.


### GetName

`func (o *ReqTag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReqTag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReqTag) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ReqTag) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReqTag) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReqTag) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReqTag) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDomainId

`func (o *ReqTag) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *ReqTag) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *ReqTag) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.


### GetStatus

`func (o *ReqTag) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReqTag) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReqTag) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReqTag) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIcon

`func (o *ReqTag) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *ReqTag) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *ReqTag) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *ReqTag) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetSort

`func (o *ReqTag) GetSort() int32`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *ReqTag) GetSortOk() (*int32, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *ReqTag) SetSort(v int32)`

SetSort sets Sort field to given value.

### HasSort

`func (o *ReqTag) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


