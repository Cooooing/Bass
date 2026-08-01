# ListCommentThreadsReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**PageReq**](PageReq.md) |  | [optional] 
**ArticleId** | **string** |  | 
**Order** | Pointer to **string** |  | [optional] 
**ReplyPreviewLimit** | Pointer to **int32** |  | [optional] 

## Methods

### NewListCommentThreadsReq

`func NewListCommentThreadsReq(articleId string, ) *ListCommentThreadsReq`

NewListCommentThreadsReq instantiates a new ListCommentThreadsReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCommentThreadsReqWithDefaults

`func NewListCommentThreadsReqWithDefaults() *ListCommentThreadsReq`

NewListCommentThreadsReqWithDefaults instantiates a new ListCommentThreadsReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ListCommentThreadsReq) GetPage() PageReq`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListCommentThreadsReq) GetPageOk() (*PageReq, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListCommentThreadsReq) SetPage(v PageReq)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListCommentThreadsReq) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetArticleId

`func (o *ListCommentThreadsReq) GetArticleId() string`

GetArticleId returns the ArticleId field if non-nil, zero value otherwise.

### GetArticleIdOk

`func (o *ListCommentThreadsReq) GetArticleIdOk() (*string, bool)`

GetArticleIdOk returns a tuple with the ArticleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticleId

`func (o *ListCommentThreadsReq) SetArticleId(v string)`

SetArticleId sets ArticleId field to given value.


### GetOrder

`func (o *ListCommentThreadsReq) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ListCommentThreadsReq) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ListCommentThreadsReq) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ListCommentThreadsReq) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetReplyPreviewLimit

`func (o *ListCommentThreadsReq) GetReplyPreviewLimit() int32`

GetReplyPreviewLimit returns the ReplyPreviewLimit field if non-nil, zero value otherwise.

### GetReplyPreviewLimitOk

`func (o *ListCommentThreadsReq) GetReplyPreviewLimitOk() (*int32, bool)`

GetReplyPreviewLimitOk returns a tuple with the ReplyPreviewLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyPreviewLimit

`func (o *ListCommentThreadsReq) SetReplyPreviewLimit(v int32)`

SetReplyPreviewLimit sets ReplyPreviewLimit field to given value.

### HasReplyPreviewLimit

`func (o *ListCommentThreadsReq) HasReplyPreviewLimit() bool`

HasReplyPreviewLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


