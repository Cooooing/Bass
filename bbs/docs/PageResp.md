# PageResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | 总数 | [optional] 
**Page** | Pointer to **int32** | 页码 | [optional] 
**Size** | Pointer to **int32** | 页大小 | [optional] 

## Methods

### NewPageResp

`func NewPageResp() *PageResp`

NewPageResp instantiates a new PageResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageRespWithDefaults

`func NewPageRespWithDefaults() *PageResp`

NewPageRespWithDefaults instantiates a new PageResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *PageResp) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PageResp) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PageResp) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PageResp) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetPage

`func (o *PageResp) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PageResp) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PageResp) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *PageResp) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetSize

`func (o *PageResp) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PageResp) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PageResp) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *PageResp) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


