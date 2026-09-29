# GetArticleReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArticleId** | **string** |  | 
**PublishStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewGetArticleReq

`func NewGetArticleReq(articleId string, ) *GetArticleReq`

NewGetArticleReq instantiates a new GetArticleReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetArticleReqWithDefaults

`func NewGetArticleReqWithDefaults() *GetArticleReq`

NewGetArticleReqWithDefaults instantiates a new GetArticleReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArticleId

`func (o *GetArticleReq) GetArticleId() string`

GetArticleId returns the ArticleId field if non-nil, zero value otherwise.

### GetArticleIdOk

`func (o *GetArticleReq) GetArticleIdOk() (*string, bool)`

GetArticleIdOk returns a tuple with the ArticleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticleId

`func (o *GetArticleReq) SetArticleId(v string)`

SetArticleId sets ArticleId field to given value.


### GetPublishStatus

`func (o *GetArticleReq) GetPublishStatus() string`

GetPublishStatus returns the PublishStatus field if non-nil, zero value otherwise.

### GetPublishStatusOk

`func (o *GetArticleReq) GetPublishStatusOk() (*string, bool)`

GetPublishStatusOk returns a tuple with the PublishStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishStatus

`func (o *GetArticleReq) SetPublishStatus(v string)`

SetPublishStatus sets PublishStatus field to given value.

### HasPublishStatus

`func (o *GetArticleReq) HasPublishStatus() bool`

HasPublishStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


