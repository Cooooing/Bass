# GetStatusRelationResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**RespRelationStatus**](RespRelationStatus.md) |  | [optional] 

## Methods

### NewGetStatusRelationResp

`func NewGetStatusRelationResp() *GetStatusRelationResp`

NewGetStatusRelationResp instantiates a new GetStatusRelationResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStatusRelationRespWithDefaults

`func NewGetStatusRelationRespWithDefaults() *GetStatusRelationResp`

NewGetStatusRelationRespWithDefaults instantiates a new GetStatusRelationResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetStatusRelationResp) GetStatus() RespRelationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetStatusRelationResp) GetStatusOk() (*RespRelationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetStatusRelationResp) SetStatus(v RespRelationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetStatusRelationResp) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


