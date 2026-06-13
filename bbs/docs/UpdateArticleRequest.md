# UpdateArticleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArticleId** | **string** |  | 
**Article** | [**RequestArticle**](RequestArticle.md) |  | 

## Methods

### NewUpdateArticleRequest

`func NewUpdateArticleRequest(articleId string, article RequestArticle, ) *UpdateArticleRequest`

NewUpdateArticleRequest instantiates a new UpdateArticleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateArticleRequestWithDefaults

`func NewUpdateArticleRequestWithDefaults() *UpdateArticleRequest`

NewUpdateArticleRequestWithDefaults instantiates a new UpdateArticleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArticleId

`func (o *UpdateArticleRequest) GetArticleId() string`

GetArticleId returns the ArticleId field if non-nil, zero value otherwise.

### GetArticleIdOk

`func (o *UpdateArticleRequest) GetArticleIdOk() (*string, bool)`

GetArticleIdOk returns a tuple with the ArticleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticleId

`func (o *UpdateArticleRequest) SetArticleId(v string)`

SetArticleId sets ArticleId field to given value.


### GetArticle

`func (o *UpdateArticleRequest) GetArticle() RequestArticle`

GetArticle returns the Article field if non-nil, zero value otherwise.

### GetArticleOk

`func (o *UpdateArticleRequest) GetArticleOk() (*RequestArticle, bool)`

GetArticleOk returns a tuple with the Article field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticle

`func (o *UpdateArticleRequest) SetArticle(v RequestArticle)`

SetArticle sets Article field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


