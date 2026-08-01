# ListCommentsResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**PageResp**](PageResp.md) |  | [optional] 
**Rows** | Pointer to [**[]RespCommentListItem**](RespCommentListItem.md) |  | [optional] 

## Methods

### NewListCommentsResp

`func NewListCommentsResp() *ListCommentsResp`

NewListCommentsResp instantiates a new ListCommentsResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCommentsRespWithDefaults

`func NewListCommentsRespWithDefaults() *ListCommentsResp`

NewListCommentsRespWithDefaults instantiates a new ListCommentsResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ListCommentsResp) GetPage() PageResp`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListCommentsResp) GetPageOk() (*PageResp, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListCommentsResp) SetPage(v PageResp)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListCommentsResp) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetRows

`func (o *ListCommentsResp) GetRows() []RespCommentListItem`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *ListCommentsResp) GetRowsOk() (*[]RespCommentListItem, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *ListCommentsResp) SetRows(v []RespCommentListItem)`

SetRows sets Rows field to given value.

### HasRows

`func (o *ListCommentsResp) HasRows() bool`

HasRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


