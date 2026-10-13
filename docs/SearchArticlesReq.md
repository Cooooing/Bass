# SearchArticlesReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** |  | 
**Cursor** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 

## Methods

### NewSearchArticlesReq

`func NewSearchArticlesReq(query string, ) *SearchArticlesReq`

NewSearchArticlesReq instantiates a new SearchArticlesReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchArticlesReqWithDefaults

`func NewSearchArticlesReqWithDefaults() *SearchArticlesReq`

NewSearchArticlesReqWithDefaults instantiates a new SearchArticlesReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *SearchArticlesReq) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SearchArticlesReq) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SearchArticlesReq) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetCursor

`func (o *SearchArticlesReq) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *SearchArticlesReq) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *SearchArticlesReq) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *SearchArticlesReq) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetSize

`func (o *SearchArticlesReq) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SearchArticlesReq) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SearchArticlesReq) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *SearchArticlesReq) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


