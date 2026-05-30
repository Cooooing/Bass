# RelationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetId** | Pointer to **string** | 目标账号 ID。 | [optional] 
**Following** | Pointer to **bool** | 当前账号是否关注目标账号。 | [optional] 
**FollowedBy** | Pointer to **bool** | 目标账号是否关注当前账号。 | [optional] 
**Blocking** | Pointer to **bool** | 当前账号是否拉黑目标账号。 | [optional] 
**BlockedBy** | Pointer to **bool** | 目标账号是否拉黑当前账号。 | [optional] 

## Methods

### NewRelationStatus

`func NewRelationStatus() *RelationStatus`

NewRelationStatus instantiates a new RelationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationStatusWithDefaults

`func NewRelationStatusWithDefaults() *RelationStatus`

NewRelationStatusWithDefaults instantiates a new RelationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetId

`func (o *RelationStatus) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *RelationStatus) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *RelationStatus) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *RelationStatus) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetFollowing

`func (o *RelationStatus) GetFollowing() bool`

GetFollowing returns the Following field if non-nil, zero value otherwise.

### GetFollowingOk

`func (o *RelationStatus) GetFollowingOk() (*bool, bool)`

GetFollowingOk returns a tuple with the Following field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowing

`func (o *RelationStatus) SetFollowing(v bool)`

SetFollowing sets Following field to given value.

### HasFollowing

`func (o *RelationStatus) HasFollowing() bool`

HasFollowing returns a boolean if a field has been set.

### GetFollowedBy

`func (o *RelationStatus) GetFollowedBy() bool`

GetFollowedBy returns the FollowedBy field if non-nil, zero value otherwise.

### GetFollowedByOk

`func (o *RelationStatus) GetFollowedByOk() (*bool, bool)`

GetFollowedByOk returns a tuple with the FollowedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowedBy

`func (o *RelationStatus) SetFollowedBy(v bool)`

SetFollowedBy sets FollowedBy field to given value.

### HasFollowedBy

`func (o *RelationStatus) HasFollowedBy() bool`

HasFollowedBy returns a boolean if a field has been set.

### GetBlocking

`func (o *RelationStatus) GetBlocking() bool`

GetBlocking returns the Blocking field if non-nil, zero value otherwise.

### GetBlockingOk

`func (o *RelationStatus) GetBlockingOk() (*bool, bool)`

GetBlockingOk returns a tuple with the Blocking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocking

`func (o *RelationStatus) SetBlocking(v bool)`

SetBlocking sets Blocking field to given value.

### HasBlocking

`func (o *RelationStatus) HasBlocking() bool`

HasBlocking returns a boolean if a field has been set.

### GetBlockedBy

`func (o *RelationStatus) GetBlockedBy() bool`

GetBlockedBy returns the BlockedBy field if non-nil, zero value otherwise.

### GetBlockedByOk

`func (o *RelationStatus) GetBlockedByOk() (*bool, bool)`

GetBlockedByOk returns a tuple with the BlockedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedBy

`func (o *RelationStatus) SetBlockedBy(v bool)`

SetBlockedBy sets BlockedBy field to given value.

### HasBlockedBy

`func (o *RelationStatus) HasBlockedBy() bool`

HasBlockedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


