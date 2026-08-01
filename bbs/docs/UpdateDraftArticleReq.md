# UpdateDraftArticleReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArticleId** | **string** |  | 
**Article** | [**ReqArticle**](ReqArticle.md) |  | 

## Methods

### NewUpdateDraftArticleReq

`func NewUpdateDraftArticleReq(articleId string, article ReqArticle, ) *UpdateDraftArticleReq`

NewUpdateDraftArticleReq instantiates a new UpdateDraftArticleReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDraftArticleReqWithDefaults

`func NewUpdateDraftArticleReqWithDefaults() *UpdateDraftArticleReq`

NewUpdateDraftArticleReqWithDefaults instantiates a new UpdateDraftArticleReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArticleId

`func (o *UpdateDraftArticleReq) GetArticleId() string`

GetArticleId returns the ArticleId field if non-nil, zero value otherwise.

### GetArticleIdOk

`func (o *UpdateDraftArticleReq) GetArticleIdOk() (*string, bool)`

GetArticleIdOk returns a tuple with the ArticleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticleId

`func (o *UpdateDraftArticleReq) SetArticleId(v string)`

SetArticleId sets ArticleId field to given value.


### GetArticle

`func (o *UpdateDraftArticleReq) GetArticle() ReqArticle`

GetArticle returns the Article field if non-nil, zero value otherwise.

### GetArticleOk

`func (o *UpdateDraftArticleReq) GetArticleOk() (*ReqArticle, bool)`

GetArticleOk returns a tuple with the Article field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticle

`func (o *UpdateDraftArticleReq) SetArticle(v ReqArticle)`

SetArticle sets Article field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


