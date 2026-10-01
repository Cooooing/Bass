# PageResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **string** | 总数 | [optional] 
**Page** | Pointer to **string** | 页码 | [optional] 
**Size** | Pointer to **string** | 页大小 | [optional] 

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

`func (o *PageResp) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PageResp) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PageResp) SetTotal(v string)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PageResp) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetPage

`func (o *PageResp) GetPage() string`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PageResp) GetPageOk() (*string, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PageResp) SetPage(v string)`

SetPage sets Page field to given value.

### HasPage

`func (o *PageResp) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetSize

`func (o *PageResp) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PageResp) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PageResp) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *PageResp) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


