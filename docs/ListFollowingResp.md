# ListFollowingResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**PageResp**](PageResp.md) |  | [optional] 
**Rows** | Pointer to [**[]AccountProfileListItem**](AccountProfileListItem.md) |  | [optional] 

## Methods

### NewListFollowingResp

`func NewListFollowingResp() *ListFollowingResp`

NewListFollowingResp instantiates a new ListFollowingResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFollowingRespWithDefaults

`func NewListFollowingRespWithDefaults() *ListFollowingResp`

NewListFollowingRespWithDefaults instantiates a new ListFollowingResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ListFollowingResp) GetPage() PageResp`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListFollowingResp) GetPageOk() (*PageResp, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListFollowingResp) SetPage(v PageResp)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListFollowingResp) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetRows

`func (o *ListFollowingResp) GetRows() []AccountProfileListItem`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *ListFollowingResp) GetRowsOk() (*[]AccountProfileListItem, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *ListFollowingResp) SetRows(v []AccountProfileListItem)`

SetRows sets Rows field to given value.

### HasRows

`func (o *ListFollowingResp) HasRows() bool`

HasRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


