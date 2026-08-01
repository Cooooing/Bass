# ListPostscriptsResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rows** | Pointer to [**[]ArticlePostscript**](ArticlePostscript.md) |  | [optional] 

## Methods

### NewListPostscriptsResp

`func NewListPostscriptsResp() *ListPostscriptsResp`

NewListPostscriptsResp instantiates a new ListPostscriptsResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPostscriptsRespWithDefaults

`func NewListPostscriptsRespWithDefaults() *ListPostscriptsResp`

NewListPostscriptsRespWithDefaults instantiates a new ListPostscriptsResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *ListPostscriptsResp) GetRows() []ArticlePostscript`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *ListPostscriptsResp) GetRowsOk() (*[]ArticlePostscript, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *ListPostscriptsResp) SetRows(v []ArticlePostscript)`

SetRows sets Rows field to given value.

### HasRows

`func (o *ListPostscriptsResp) HasRows() bool`

HasRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


