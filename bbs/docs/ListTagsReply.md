# ListTagsReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**PageReply**](PageReply.md) | 分页结果。 | [optional] 
**Rows** | Pointer to [**[]Tag**](Tag.md) | 标签列表。 | [optional] 

## Methods

### NewListTagsReply

`func NewListTagsReply() *ListTagsReply`

NewListTagsReply instantiates a new ListTagsReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTagsReplyWithDefaults

`func NewListTagsReplyWithDefaults() *ListTagsReply`

NewListTagsReplyWithDefaults instantiates a new ListTagsReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ListTagsReply) GetPage() PageReply`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListTagsReply) GetPageOk() (*PageReply, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListTagsReply) SetPage(v PageReply)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListTagsReply) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetRows

`func (o *ListTagsReply) GetRows() []Tag`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *ListTagsReply) GetRowsOk() (*[]Tag, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *ListTagsReply) SetRows(v []Tag)`

SetRows sets Rows field to given value.

### HasRows

`func (o *ListTagsReply) HasRows() bool`

HasRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


