# PageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **int32** | 页码 | [optional] 
**Size** | Pointer to **int32** | 页大小 | [optional] 

## Methods

### NewPageRequest

`func NewPageRequest() *PageRequest`

NewPageRequest instantiates a new PageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageRequestWithDefaults

`func NewPageRequestWithDefaults() *PageRequest`

NewPageRequestWithDefaults instantiates a new PageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *PageRequest) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PageRequest) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PageRequest) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *PageRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetSize

`func (o *PageRequest) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PageRequest) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PageRequest) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *PageRequest) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


