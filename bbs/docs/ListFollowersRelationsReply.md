# ListFollowersRelationsReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**PageReply**](PageReply.md) |  | [optional] 
**Rows** | Pointer to [**[]Relation**](Relation.md) |  | [optional] 

## Methods

### NewListFollowersRelationsReply

`func NewListFollowersRelationsReply() *ListFollowersRelationsReply`

NewListFollowersRelationsReply instantiates a new ListFollowersRelationsReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFollowersRelationsReplyWithDefaults

`func NewListFollowersRelationsReplyWithDefaults() *ListFollowersRelationsReply`

NewListFollowersRelationsReplyWithDefaults instantiates a new ListFollowersRelationsReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ListFollowersRelationsReply) GetPage() PageReply`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListFollowersRelationsReply) GetPageOk() (*PageReply, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListFollowersRelationsReply) SetPage(v PageReply)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListFollowersRelationsReply) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetRows

`func (o *ListFollowersRelationsReply) GetRows() []Relation`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *ListFollowersRelationsReply) GetRowsOk() (*[]Relation, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *ListFollowersRelationsReply) SetRows(v []Relation)`

SetRows sets Rows field to given value.

### HasRows

`func (o *ListFollowersRelationsReply) HasRows() bool`

HasRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


