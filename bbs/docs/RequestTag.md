# RequestTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 标签名称。 | 
**Description** | Pointer to **string** | 标签描述。 | [optional] 
**DomainId** | Pointer to **string** | 所属板块 ID。 | [optional] 
**Status** | Pointer to **string** | 标签启停状态。 | [optional] 

## Methods

### NewRequestTag

`func NewRequestTag(name string, ) *RequestTag`

NewRequestTag instantiates a new RequestTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestTagWithDefaults

`func NewRequestTagWithDefaults() *RequestTag`

NewRequestTagWithDefaults instantiates a new RequestTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RequestTag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestTag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestTag) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RequestTag) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RequestTag) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RequestTag) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RequestTag) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDomainId

`func (o *RequestTag) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *RequestTag) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *RequestTag) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *RequestTag) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetStatus

`func (o *RequestTag) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RequestTag) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RequestTag) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RequestTag) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


