# UpdateTagReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TagId** | **string** |  | 
**Tag** | [**ReqTag**](ReqTag.md) |  | 

## Methods

### NewUpdateTagReq

`func NewUpdateTagReq(tagId string, tag ReqTag, ) *UpdateTagReq`

NewUpdateTagReq instantiates a new UpdateTagReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTagReqWithDefaults

`func NewUpdateTagReqWithDefaults() *UpdateTagReq`

NewUpdateTagReqWithDefaults instantiates a new UpdateTagReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagId

`func (o *UpdateTagReq) GetTagId() string`

GetTagId returns the TagId field if non-nil, zero value otherwise.

### GetTagIdOk

`func (o *UpdateTagReq) GetTagIdOk() (*string, bool)`

GetTagIdOk returns a tuple with the TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagId

`func (o *UpdateTagReq) SetTagId(v string)`

SetTagId sets TagId field to given value.


### GetTag

`func (o *UpdateTagReq) GetTag() ReqTag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *UpdateTagReq) GetTagOk() (*ReqTag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *UpdateTagReq) SetTag(v ReqTag)`

SetTag sets Tag field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


