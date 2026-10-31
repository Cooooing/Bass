# TransferAccountPointsReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetName** | **string** |  | 
**Amount** | **string** |  | 
**Remark** | Pointer to **string** |  | [optional] 
**IdempotencyKey** | **string** |  | 

## Methods

### NewTransferAccountPointsReq

`func NewTransferAccountPointsReq(targetName string, amount string, idempotencyKey string, ) *TransferAccountPointsReq`

NewTransferAccountPointsReq instantiates a new TransferAccountPointsReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferAccountPointsReqWithDefaults

`func NewTransferAccountPointsReqWithDefaults() *TransferAccountPointsReq`

NewTransferAccountPointsReqWithDefaults instantiates a new TransferAccountPointsReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetName

`func (o *TransferAccountPointsReq) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *TransferAccountPointsReq) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *TransferAccountPointsReq) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.


### GetAmount

`func (o *TransferAccountPointsReq) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransferAccountPointsReq) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransferAccountPointsReq) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetRemark

`func (o *TransferAccountPointsReq) GetRemark() string`

GetRemark returns the Remark field if non-nil, zero value otherwise.

### GetRemarkOk

`func (o *TransferAccountPointsReq) GetRemarkOk() (*string, bool)`

GetRemarkOk returns a tuple with the Remark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemark

`func (o *TransferAccountPointsReq) SetRemark(v string)`

SetRemark sets Remark field to given value.

### HasRemark

`func (o *TransferAccountPointsReq) HasRemark() bool`

HasRemark returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *TransferAccountPointsReq) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *TransferAccountPointsReq) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *TransferAccountPointsReq) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


