# RespRelationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetId** | Pointer to **string** |  | [optional] 
**Following** | Pointer to **bool** |  | [optional] 
**FollowedBy** | Pointer to **bool** |  | [optional] 
**Blocking** | Pointer to **bool** |  | [optional] 
**BlockedBy** | Pointer to **bool** |  | [optional] 

## Methods

### NewRespRelationStatus

`func NewRespRelationStatus() *RespRelationStatus`

NewRespRelationStatus instantiates a new RespRelationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRespRelationStatusWithDefaults

`func NewRespRelationStatusWithDefaults() *RespRelationStatus`

NewRespRelationStatusWithDefaults instantiates a new RespRelationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetId

`func (o *RespRelationStatus) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *RespRelationStatus) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *RespRelationStatus) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *RespRelationStatus) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetFollowing

`func (o *RespRelationStatus) GetFollowing() bool`

GetFollowing returns the Following field if non-nil, zero value otherwise.

### GetFollowingOk

`func (o *RespRelationStatus) GetFollowingOk() (*bool, bool)`

GetFollowingOk returns a tuple with the Following field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowing

`func (o *RespRelationStatus) SetFollowing(v bool)`

SetFollowing sets Following field to given value.

### HasFollowing

`func (o *RespRelationStatus) HasFollowing() bool`

HasFollowing returns a boolean if a field has been set.

### GetFollowedBy

`func (o *RespRelationStatus) GetFollowedBy() bool`

GetFollowedBy returns the FollowedBy field if non-nil, zero value otherwise.

### GetFollowedByOk

`func (o *RespRelationStatus) GetFollowedByOk() (*bool, bool)`

GetFollowedByOk returns a tuple with the FollowedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowedBy

`func (o *RespRelationStatus) SetFollowedBy(v bool)`

SetFollowedBy sets FollowedBy field to given value.

### HasFollowedBy

`func (o *RespRelationStatus) HasFollowedBy() bool`

HasFollowedBy returns a boolean if a field has been set.

### GetBlocking

`func (o *RespRelationStatus) GetBlocking() bool`

GetBlocking returns the Blocking field if non-nil, zero value otherwise.

### GetBlockingOk

`func (o *RespRelationStatus) GetBlockingOk() (*bool, bool)`

GetBlockingOk returns a tuple with the Blocking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocking

`func (o *RespRelationStatus) SetBlocking(v bool)`

SetBlocking sets Blocking field to given value.

### HasBlocking

`func (o *RespRelationStatus) HasBlocking() bool`

HasBlocking returns a boolean if a field has been set.

### GetBlockedBy

`func (o *RespRelationStatus) GetBlockedBy() bool`

GetBlockedBy returns the BlockedBy field if non-nil, zero value otherwise.

### GetBlockedByOk

`func (o *RespRelationStatus) GetBlockedByOk() (*bool, bool)`

GetBlockedByOk returns a tuple with the BlockedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedBy

`func (o *RespRelationStatus) SetBlockedBy(v bool)`

SetBlockedBy sets BlockedBy field to given value.

### HasBlockedBy

`func (o *RespRelationStatus) HasBlockedBy() bool`

HasBlockedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


