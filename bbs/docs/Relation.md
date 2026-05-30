# Relation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | 关系记录 ID。 | [optional] 
**Type** | Pointer to **int32** | 关系类型。 | [optional] 
**ActorId** | Pointer to **string** | 发起账号 ID。 | [optional] 
**TargetId** | Pointer to **string** | 目标账号 ID。 | [optional] 
**CreatedAt** | Pointer to **string** | 创建时间。 | [optional] 
**UpdatedAt** | Pointer to **string** | 更新时间。 | [optional] 

## Methods

### NewRelation

`func NewRelation() *Relation`

NewRelation instantiates a new Relation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationWithDefaults

`func NewRelationWithDefaults() *Relation`

NewRelationWithDefaults instantiates a new Relation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Relation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Relation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Relation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Relation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Relation) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Relation) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Relation) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *Relation) HasType() bool`

HasType returns a boolean if a field has been set.

### GetActorId

`func (o *Relation) GetActorId() string`

GetActorId returns the ActorId field if non-nil, zero value otherwise.

### GetActorIdOk

`func (o *Relation) GetActorIdOk() (*string, bool)`

GetActorIdOk returns a tuple with the ActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorId

`func (o *Relation) SetActorId(v string)`

SetActorId sets ActorId field to given value.

### HasActorId

`func (o *Relation) HasActorId() bool`

HasActorId returns a boolean if a field has been set.

### GetTargetId

`func (o *Relation) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *Relation) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *Relation) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *Relation) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Relation) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Relation) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Relation) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Relation) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Relation) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Relation) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Relation) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Relation) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


