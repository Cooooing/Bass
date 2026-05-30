# WatchArticleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArticleId** | **string** | 文章 ID。 | 
**Active** | **bool** | 是否关注。 | 

## Methods

### NewWatchArticleRequest

`func NewWatchArticleRequest(articleId string, active bool, ) *WatchArticleRequest`

NewWatchArticleRequest instantiates a new WatchArticleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWatchArticleRequestWithDefaults

`func NewWatchArticleRequestWithDefaults() *WatchArticleRequest`

NewWatchArticleRequestWithDefaults instantiates a new WatchArticleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArticleId

`func (o *WatchArticleRequest) GetArticleId() string`

GetArticleId returns the ArticleId field if non-nil, zero value otherwise.

### GetArticleIdOk

`func (o *WatchArticleRequest) GetArticleIdOk() (*string, bool)`

GetArticleIdOk returns a tuple with the ArticleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticleId

`func (o *WatchArticleRequest) SetArticleId(v string)`

SetArticleId sets ArticleId field to given value.


### GetActive

`func (o *WatchArticleRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *WatchArticleRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *WatchArticleRequest) SetActive(v bool)`

SetActive sets Active field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


