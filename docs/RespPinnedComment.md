# RespPinnedComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Thread** | Pointer to [**RespCommentThread**](RespCommentThread.md) |  | [optional] 

## Methods

### NewRespPinnedComment

`func NewRespPinnedComment() *RespPinnedComment`

NewRespPinnedComment instantiates a new RespPinnedComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRespPinnedCommentWithDefaults

`func NewRespPinnedCommentWithDefaults() *RespPinnedComment`

NewRespPinnedCommentWithDefaults instantiates a new RespPinnedComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThread

`func (o *RespPinnedComment) GetThread() RespCommentThread`

GetThread returns the Thread field if non-nil, zero value otherwise.

### GetThreadOk

`func (o *RespPinnedComment) GetThreadOk() (*RespCommentThread, bool)`

GetThreadOk returns a tuple with the Thread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThread

`func (o *RespPinnedComment) SetThread(v RespCommentThread)`

SetThread sets Thread field to given value.

### HasThread

`func (o *RespPinnedComment) HasThread() bool`

HasThread returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


