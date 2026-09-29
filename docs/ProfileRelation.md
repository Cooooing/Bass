# ProfileRelation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Following** | Pointer to **bool** |  | [optional] 
**FollowedBy** | Pointer to **bool** |  | [optional] 
**Blocking** | Pointer to **bool** |  | [optional] 
**BlockedBy** | Pointer to **bool** |  | [optional] 

## Methods

### NewProfileRelation

`func NewProfileRelation() *ProfileRelation`

NewProfileRelation instantiates a new ProfileRelation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileRelationWithDefaults

`func NewProfileRelationWithDefaults() *ProfileRelation`

NewProfileRelationWithDefaults instantiates a new ProfileRelation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFollowing

`func (o *ProfileRelation) GetFollowing() bool`

GetFollowing returns the Following field if non-nil, zero value otherwise.

### GetFollowingOk

`func (o *ProfileRelation) GetFollowingOk() (*bool, bool)`

GetFollowingOk returns a tuple with the Following field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowing

`func (o *ProfileRelation) SetFollowing(v bool)`

SetFollowing sets Following field to given value.

### HasFollowing

`func (o *ProfileRelation) HasFollowing() bool`

HasFollowing returns a boolean if a field has been set.

### GetFollowedBy

`func (o *ProfileRelation) GetFollowedBy() bool`

GetFollowedBy returns the FollowedBy field if non-nil, zero value otherwise.

### GetFollowedByOk

`func (o *ProfileRelation) GetFollowedByOk() (*bool, bool)`

GetFollowedByOk returns a tuple with the FollowedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowedBy

`func (o *ProfileRelation) SetFollowedBy(v bool)`

SetFollowedBy sets FollowedBy field to given value.

### HasFollowedBy

`func (o *ProfileRelation) HasFollowedBy() bool`

HasFollowedBy returns a boolean if a field has been set.

### GetBlocking

`func (o *ProfileRelation) GetBlocking() bool`

GetBlocking returns the Blocking field if non-nil, zero value otherwise.

### GetBlockingOk

`func (o *ProfileRelation) GetBlockingOk() (*bool, bool)`

GetBlockingOk returns a tuple with the Blocking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocking

`func (o *ProfileRelation) SetBlocking(v bool)`

SetBlocking sets Blocking field to given value.

### HasBlocking

`func (o *ProfileRelation) HasBlocking() bool`

HasBlocking returns a boolean if a field has been set.

### GetBlockedBy

`func (o *ProfileRelation) GetBlockedBy() bool`

GetBlockedBy returns the BlockedBy field if non-nil, zero value otherwise.

### GetBlockedByOk

`func (o *ProfileRelation) GetBlockedByOk() (*bool, bool)`

GetBlockedByOk returns a tuple with the BlockedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedBy

`func (o *ProfileRelation) SetBlockedBy(v bool)`

SetBlockedBy sets BlockedBy field to given value.

### HasBlockedBy

`func (o *ProfileRelation) HasBlockedBy() bool`

HasBlockedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


