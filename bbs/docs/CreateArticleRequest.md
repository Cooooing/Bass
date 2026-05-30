# CreateArticleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Article** | [**ArticleSave**](ArticleSave.md) | 文章保存内容。 | 

## Methods

### NewCreateArticleRequest

`func NewCreateArticleRequest(article ArticleSave, ) *CreateArticleRequest`

NewCreateArticleRequest instantiates a new CreateArticleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateArticleRequestWithDefaults

`func NewCreateArticleRequestWithDefaults() *CreateArticleRequest`

NewCreateArticleRequestWithDefaults instantiates a new CreateArticleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArticle

`func (o *CreateArticleRequest) GetArticle() ArticleSave`

GetArticle returns the Article field if non-nil, zero value otherwise.

### GetArticleOk

`func (o *CreateArticleRequest) GetArticleOk() (*ArticleSave, bool)`

GetArticleOk returns a tuple with the Article field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticle

`func (o *CreateArticleRequest) SetArticle(v ArticleSave)`

SetArticle sets Article field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


