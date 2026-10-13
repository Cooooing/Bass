# SearchArticlesResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rows** | Pointer to [**[]ArticleListItem**](ArticleListItem.md) |  | [optional] 
**NextCursor** | Pointer to **string** |  | [optional] 

## Methods

### NewSearchArticlesResp

`func NewSearchArticlesResp() *SearchArticlesResp`

NewSearchArticlesResp instantiates a new SearchArticlesResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchArticlesRespWithDefaults

`func NewSearchArticlesRespWithDefaults() *SearchArticlesResp`

NewSearchArticlesRespWithDefaults instantiates a new SearchArticlesResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *SearchArticlesResp) GetRows() []ArticleListItem`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *SearchArticlesResp) GetRowsOk() (*[]ArticleListItem, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *SearchArticlesResp) SetRows(v []ArticleListItem)`

SetRows sets Rows field to given value.

### HasRows

`func (o *SearchArticlesResp) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetNextCursor

`func (o *SearchArticlesResp) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *SearchArticlesResp) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *SearchArticlesResp) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *SearchArticlesResp) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


