# PageReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | 总数 | [optional] 
**Page** | Pointer to **int32** | 页码 | [optional] 
**Size** | Pointer to **int32** | 页大小 | [optional] 

## Methods

### NewPageReply

`func NewPageReply() *PageReply`

NewPageReply instantiates a new PageReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageReplyWithDefaults

`func NewPageReplyWithDefaults() *PageReply`

NewPageReplyWithDefaults instantiates a new PageReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *PageReply) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PageReply) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PageReply) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PageReply) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetPage

`func (o *PageReply) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PageReply) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PageReply) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *PageReply) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetSize

`func (o *PageReply) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PageReply) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PageReply) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *PageReply) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


