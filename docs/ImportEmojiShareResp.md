# ImportEmojiShareResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to [**EmojiGroup**](EmojiGroup.md) |  | [optional] 
**SkippedCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewImportEmojiShareResp

`func NewImportEmojiShareResp() *ImportEmojiShareResp`

NewImportEmojiShareResp instantiates a new ImportEmojiShareResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportEmojiShareRespWithDefaults

`func NewImportEmojiShareRespWithDefaults() *ImportEmojiShareResp`

NewImportEmojiShareRespWithDefaults instantiates a new ImportEmojiShareResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *ImportEmojiShareResp) GetGroup() EmojiGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ImportEmojiShareResp) GetGroupOk() (*EmojiGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ImportEmojiShareResp) SetGroup(v EmojiGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ImportEmojiShareResp) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetSkippedCount

`func (o *ImportEmojiShareResp) GetSkippedCount() int32`

GetSkippedCount returns the SkippedCount field if non-nil, zero value otherwise.

### GetSkippedCountOk

`func (o *ImportEmojiShareResp) GetSkippedCountOk() (*int32, bool)`

GetSkippedCountOk returns a tuple with the SkippedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippedCount

`func (o *ImportEmojiShareResp) SetSkippedCount(v int32)`

SetSkippedCount sets SkippedCount field to given value.

### HasSkippedCount

`func (o *ImportEmojiShareResp) HasSkippedCount() bool`

HasSkippedCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


