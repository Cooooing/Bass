# ListEmojiGroupsResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rows** | Pointer to [**[]EmojiGroup**](EmojiGroup.md) |  | [optional] 

## Methods

### NewListEmojiGroupsResp

`func NewListEmojiGroupsResp() *ListEmojiGroupsResp`

NewListEmojiGroupsResp instantiates a new ListEmojiGroupsResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEmojiGroupsRespWithDefaults

`func NewListEmojiGroupsRespWithDefaults() *ListEmojiGroupsResp`

NewListEmojiGroupsRespWithDefaults instantiates a new ListEmojiGroupsResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *ListEmojiGroupsResp) GetRows() []EmojiGroup`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *ListEmojiGroupsResp) GetRowsOk() (*[]EmojiGroup, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *ListEmojiGroupsResp) SetRows(v []EmojiGroup)`

SetRows sets Rows field to given value.

### HasRows

`func (o *ListEmojiGroupsResp) HasRows() bool`

HasRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


