# Character

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Slot** | Pointer to **int32** |  | [optional] 
**ActionQueueCapacity** | Pointer to **int32** |  | [optional] 
**MaxOfflineSeconds** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCharacter

`func NewCharacter() *Character`

NewCharacter instantiates a new Character object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharacterWithDefaults

`func NewCharacterWithDefaults() *Character`

NewCharacterWithDefaults instantiates a new Character object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Character) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Character) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Character) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Character) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Character) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Character) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Character) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Character) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *Character) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Character) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Character) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Character) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSlot

`func (o *Character) GetSlot() int32`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *Character) GetSlotOk() (*int32, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *Character) SetSlot(v int32)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *Character) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### GetActionQueueCapacity

`func (o *Character) GetActionQueueCapacity() int32`

GetActionQueueCapacity returns the ActionQueueCapacity field if non-nil, zero value otherwise.

### GetActionQueueCapacityOk

`func (o *Character) GetActionQueueCapacityOk() (*int32, bool)`

GetActionQueueCapacityOk returns a tuple with the ActionQueueCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionQueueCapacity

`func (o *Character) SetActionQueueCapacity(v int32)`

SetActionQueueCapacity sets ActionQueueCapacity field to given value.

### HasActionQueueCapacity

`func (o *Character) HasActionQueueCapacity() bool`

HasActionQueueCapacity returns a boolean if a field has been set.

### GetMaxOfflineSeconds

`func (o *Character) GetMaxOfflineSeconds() string`

GetMaxOfflineSeconds returns the MaxOfflineSeconds field if non-nil, zero value otherwise.

### GetMaxOfflineSecondsOk

`func (o *Character) GetMaxOfflineSecondsOk() (*string, bool)`

GetMaxOfflineSecondsOk returns a tuple with the MaxOfflineSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOfflineSeconds

`func (o *Character) SetMaxOfflineSeconds(v string)`

SetMaxOfflineSeconds sets MaxOfflineSeconds field to given value.

### HasMaxOfflineSeconds

`func (o *Character) HasMaxOfflineSeconds() bool`

HasMaxOfflineSeconds returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Character) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Character) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Character) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Character) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Character) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Character) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Character) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Character) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


