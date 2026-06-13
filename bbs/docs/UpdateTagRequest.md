# UpdateTagRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TagId** | **string** | 标签 ID。 | 
**Tag** | [**RequestTag**](RequestTag.md) | 标签保存内容。 | 

## Methods

### NewUpdateTagRequest

`func NewUpdateTagRequest(tagId string, tag RequestTag, ) *UpdateTagRequest`

NewUpdateTagRequest instantiates a new UpdateTagRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTagRequestWithDefaults

`func NewUpdateTagRequestWithDefaults() *UpdateTagRequest`

NewUpdateTagRequestWithDefaults instantiates a new UpdateTagRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagId

`func (o *UpdateTagRequest) GetTagId() string`

GetTagId returns the TagId field if non-nil, zero value otherwise.

### GetTagIdOk

`func (o *UpdateTagRequest) GetTagIdOk() (*string, bool)`

GetTagIdOk returns a tuple with the TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagId

`func (o *UpdateTagRequest) SetTagId(v string)`

SetTagId sets TagId field to given value.


### GetTag

`func (o *UpdateTagRequest) GetTag() RequestTag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *UpdateTagRequest) GetTagOk() (*RequestTag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *UpdateTagRequest) SetTag(v RequestTag)`

SetTag sets Tag field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


