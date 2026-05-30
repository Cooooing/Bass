# ListBlockedRelationsReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**PageReply**](PageReply.md) | 分页结果。 | [optional] 
**Rows** | Pointer to [**[]Relation**](Relation.md) | 账号关系列表。 | [optional] 

## Methods

### NewListBlockedRelationsReply

`func NewListBlockedRelationsReply() *ListBlockedRelationsReply`

NewListBlockedRelationsReply instantiates a new ListBlockedRelationsReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBlockedRelationsReplyWithDefaults

`func NewListBlockedRelationsReplyWithDefaults() *ListBlockedRelationsReply`

NewListBlockedRelationsReplyWithDefaults instantiates a new ListBlockedRelationsReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ListBlockedRelationsReply) GetPage() PageReply`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListBlockedRelationsReply) GetPageOk() (*PageReply, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListBlockedRelationsReply) SetPage(v PageReply)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListBlockedRelationsReply) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetRows

`func (o *ListBlockedRelationsReply) GetRows() []Relation`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *ListBlockedRelationsReply) GetRowsOk() (*[]Relation, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *ListBlockedRelationsReply) SetRows(v []Relation)`

SetRows sets Rows field to given value.

### HasRows

`func (o *ListBlockedRelationsReply) HasRows() bool`

HasRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


