# RespPointRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**TypeLabel** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **string** |  | [optional] 
**BalanceAfter** | Pointer to **string** |  | [optional] 

## Methods

### NewRespPointRecord

`func NewRespPointRecord() *RespPointRecord`

NewRespPointRecord instantiates a new RespPointRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRespPointRecordWithDefaults

`func NewRespPointRecordWithDefaults() *RespPointRecord`

NewRespPointRecordWithDefaults instantiates a new RespPointRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *RespPointRecord) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RespPointRecord) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RespPointRecord) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RespPointRecord) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetTypeLabel

`func (o *RespPointRecord) GetTypeLabel() string`

GetTypeLabel returns the TypeLabel field if non-nil, zero value otherwise.

### GetTypeLabelOk

`func (o *RespPointRecord) GetTypeLabelOk() (*string, bool)`

GetTypeLabelOk returns a tuple with the TypeLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeLabel

`func (o *RespPointRecord) SetTypeLabel(v string)`

SetTypeLabel sets TypeLabel field to given value.

### HasTypeLabel

`func (o *RespPointRecord) HasTypeLabel() bool`

HasTypeLabel returns a boolean if a field has been set.

### GetDirection

`func (o *RespPointRecord) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *RespPointRecord) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *RespPointRecord) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *RespPointRecord) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetAmount

`func (o *RespPointRecord) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RespPointRecord) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RespPointRecord) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *RespPointRecord) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBalanceAfter

`func (o *RespPointRecord) GetBalanceAfter() string`

GetBalanceAfter returns the BalanceAfter field if non-nil, zero value otherwise.

### GetBalanceAfterOk

`func (o *RespPointRecord) GetBalanceAfterOk() (*string, bool)`

GetBalanceAfterOk returns a tuple with the BalanceAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAfter

`func (o *RespPointRecord) SetBalanceAfter(v string)`

SetBalanceAfter sets BalanceAfter field to given value.

### HasBalanceAfter

`func (o *RespPointRecord) HasBalanceAfter() bool`

HasBalanceAfter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


