# GetStatusRelationReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**RelationStatus**](RelationStatus.md) | 关系状态。 | [optional] 

## Methods

### NewGetStatusRelationReply

`func NewGetStatusRelationReply() *GetStatusRelationReply`

NewGetStatusRelationReply instantiates a new GetStatusRelationReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStatusRelationReplyWithDefaults

`func NewGetStatusRelationReplyWithDefaults() *GetStatusRelationReply`

NewGetStatusRelationReplyWithDefaults instantiates a new GetStatusRelationReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetStatusRelationReply) GetStatus() RelationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetStatusRelationReply) GetStatusOk() (*RelationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetStatusRelationReply) SetStatus(v RelationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetStatusRelationReply) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


