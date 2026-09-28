# ImageResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **string** | 图片数据 | [optional] 
**ContentType** | Pointer to **string** | 图片类型 | [optional] 

## Methods

### NewImageResp

`func NewImageResp() *ImageResp`

NewImageResp instantiates a new ImageResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageRespWithDefaults

`func NewImageRespWithDefaults() *ImageResp`

NewImageRespWithDefaults instantiates a new ImageResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ImageResp) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ImageResp) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ImageResp) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *ImageResp) HasData() bool`

HasData returns a boolean if a field has been set.

### GetContentType

`func (o *ImageResp) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ImageResp) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ImageResp) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *ImageResp) HasContentType() bool`

HasContentType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


