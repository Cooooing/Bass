# ListFollowingRelationsResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**PageResp**](PageResp.md) |  | [optional] 
**Rows** | Pointer to [**[]RespRelation**](RespRelation.md) |  | [optional] 

## Methods

### NewListFollowingRelationsResp

`func NewListFollowingRelationsResp() *ListFollowingRelationsResp`

NewListFollowingRelationsResp instantiates a new ListFollowingRelationsResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFollowingRelationsRespWithDefaults

`func NewListFollowingRelationsRespWithDefaults() *ListFollowingRelationsResp`

NewListFollowingRelationsRespWithDefaults instantiates a new ListFollowingRelationsResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ListFollowingRelationsResp) GetPage() PageResp`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListFollowingRelationsResp) GetPageOk() (*PageResp, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListFollowingRelationsResp) SetPage(v PageResp)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListFollowingRelationsResp) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetRows

`func (o *ListFollowingRelationsResp) GetRows() []RespRelation`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *ListFollowingRelationsResp) GetRowsOk() (*[]RespRelation, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *ListFollowingRelationsResp) SetRows(v []RespRelation)`

SetRows sets Rows field to given value.

### HasRows

`func (o *ListFollowingRelationsResp) HasRows() bool`

HasRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


