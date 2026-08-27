# RespRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**TransactionNo** | Pointer to **string** |  | [optional] 
**RecordType** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **string** |  | [optional] 
**BalanceBefore** | Pointer to **string** |  | [optional] 
**BalanceAfter** | Pointer to **string** |  | [optional] 
**Remark** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRespRecord

`func NewRespRecord() *RespRecord`

NewRespRecord instantiates a new RespRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRespRecordWithDefaults

`func NewRespRecordWithDefaults() *RespRecord`

NewRespRecordWithDefaults instantiates a new RespRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RespRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RespRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RespRecord) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RespRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTransactionNo

`func (o *RespRecord) GetTransactionNo() string`

GetTransactionNo returns the TransactionNo field if non-nil, zero value otherwise.

### GetTransactionNoOk

`func (o *RespRecord) GetTransactionNoOk() (*string, bool)`

GetTransactionNoOk returns a tuple with the TransactionNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionNo

`func (o *RespRecord) SetTransactionNo(v string)`

SetTransactionNo sets TransactionNo field to given value.

### HasTransactionNo

`func (o *RespRecord) HasTransactionNo() bool`

HasTransactionNo returns a boolean if a field has been set.

### GetRecordType

`func (o *RespRecord) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *RespRecord) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *RespRecord) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *RespRecord) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetDirection

`func (o *RespRecord) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *RespRecord) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *RespRecord) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *RespRecord) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetAmount

`func (o *RespRecord) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RespRecord) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RespRecord) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *RespRecord) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBalanceBefore

`func (o *RespRecord) GetBalanceBefore() string`

GetBalanceBefore returns the BalanceBefore field if non-nil, zero value otherwise.

### GetBalanceBeforeOk

`func (o *RespRecord) GetBalanceBeforeOk() (*string, bool)`

GetBalanceBeforeOk returns a tuple with the BalanceBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceBefore

`func (o *RespRecord) SetBalanceBefore(v string)`

SetBalanceBefore sets BalanceBefore field to given value.

### HasBalanceBefore

`func (o *RespRecord) HasBalanceBefore() bool`

HasBalanceBefore returns a boolean if a field has been set.

### GetBalanceAfter

`func (o *RespRecord) GetBalanceAfter() string`

GetBalanceAfter returns the BalanceAfter field if non-nil, zero value otherwise.

### GetBalanceAfterOk

`func (o *RespRecord) GetBalanceAfterOk() (*string, bool)`

GetBalanceAfterOk returns a tuple with the BalanceAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAfter

`func (o *RespRecord) SetBalanceAfter(v string)`

SetBalanceAfter sets BalanceAfter field to given value.

### HasBalanceAfter

`func (o *RespRecord) HasBalanceAfter() bool`

HasBalanceAfter returns a boolean if a field has been set.

### GetRemark

`func (o *RespRecord) GetRemark() string`

GetRemark returns the Remark field if non-nil, zero value otherwise.

### GetRemarkOk

`func (o *RespRecord) GetRemarkOk() (*string, bool)`

GetRemarkOk returns a tuple with the Remark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemark

`func (o *RespRecord) SetRemark(v string)`

SetRemark sets Remark field to given value.

### HasRemark

`func (o *RespRecord) HasRemark() bool`

HasRemark returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RespRecord) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RespRecord) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RespRecord) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RespRecord) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


