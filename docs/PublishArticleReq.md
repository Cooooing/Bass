# PublishArticleReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArticleId** | **string** |  | 
**Visibility** | Pointer to **string** |  | [optional] 

## Methods

### NewPublishArticleReq

`func NewPublishArticleReq(articleId string, ) *PublishArticleReq`

NewPublishArticleReq instantiates a new PublishArticleReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublishArticleReqWithDefaults

`func NewPublishArticleReqWithDefaults() *PublishArticleReq`

NewPublishArticleReqWithDefaults instantiates a new PublishArticleReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArticleId

`func (o *PublishArticleReq) GetArticleId() string`

GetArticleId returns the ArticleId field if non-nil, zero value otherwise.

### GetArticleIdOk

`func (o *PublishArticleReq) GetArticleIdOk() (*string, bool)`

GetArticleIdOk returns a tuple with the ArticleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticleId

`func (o *PublishArticleReq) SetArticleId(v string)`

SetArticleId sets ArticleId field to given value.


### GetVisibility

`func (o *PublishArticleReq) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *PublishArticleReq) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *PublishArticleReq) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *PublishArticleReq) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


