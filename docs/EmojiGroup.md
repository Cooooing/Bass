# EmojiGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]EmojiItem**](EmojiItem.md) |  | [optional] 

## Methods

### NewEmojiGroup

`func NewEmojiGroup() *EmojiGroup`

NewEmojiGroup instantiates a new EmojiGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmojiGroupWithDefaults

`func NewEmojiGroupWithDefaults() *EmojiGroup`

NewEmojiGroupWithDefaults instantiates a new EmojiGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmojiGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmojiGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmojiGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EmojiGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EmojiGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmojiGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmojiGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EmojiGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetItems

`func (o *EmojiGroup) GetItems() []EmojiItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EmojiGroup) GetItemsOk() (*[]EmojiItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EmojiGroup) SetItems(v []EmojiItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *EmojiGroup) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


