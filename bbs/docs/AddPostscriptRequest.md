# AddPostscriptRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArticleId** | **string** | 文章 ID。 | 
**Content** | **string** | 附言内容。 | 

## Methods

### NewAddPostscriptRequest

`func NewAddPostscriptRequest(articleId string, content string, ) *AddPostscriptRequest`

NewAddPostscriptRequest instantiates a new AddPostscriptRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPostscriptRequestWithDefaults

`func NewAddPostscriptRequestWithDefaults() *AddPostscriptRequest`

NewAddPostscriptRequestWithDefaults instantiates a new AddPostscriptRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArticleId

`func (o *AddPostscriptRequest) GetArticleId() string`

GetArticleId returns the ArticleId field if non-nil, zero value otherwise.

### GetArticleIdOk

`func (o *AddPostscriptRequest) GetArticleIdOk() (*string, bool)`

GetArticleIdOk returns a tuple with the ArticleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticleId

`func (o *AddPostscriptRequest) SetArticleId(v string)`

SetArticleId sets ArticleId field to given value.


### GetContent

`func (o *AddPostscriptRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *AddPostscriptRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *AddPostscriptRequest) SetContent(v string)`

SetContent sets Content field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


