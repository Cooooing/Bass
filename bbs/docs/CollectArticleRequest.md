# CollectArticleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArticleId** | **string** | 文章 ID。 | 
**Active** | **bool** | 是否收藏。 | 

## Methods

### NewCollectArticleRequest

`func NewCollectArticleRequest(articleId string, active bool, ) *CollectArticleRequest`

NewCollectArticleRequest instantiates a new CollectArticleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectArticleRequestWithDefaults

`func NewCollectArticleRequestWithDefaults() *CollectArticleRequest`

NewCollectArticleRequestWithDefaults instantiates a new CollectArticleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArticleId

`func (o *CollectArticleRequest) GetArticleId() string`

GetArticleId returns the ArticleId field if non-nil, zero value otherwise.

### GetArticleIdOk

`func (o *CollectArticleRequest) GetArticleIdOk() (*string, bool)`

GetArticleIdOk returns a tuple with the ArticleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticleId

`func (o *CollectArticleRequest) SetArticleId(v string)`

SetArticleId sets ArticleId field to given value.


### GetActive

`func (o *CollectArticleRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CollectArticleRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CollectArticleRequest) SetActive(v bool)`

SetActive sets Active field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


