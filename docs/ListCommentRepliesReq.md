# ListCommentRepliesReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**PageReq**](PageReq.md) |  | [optional] 
**ArticleId** | **string** |  | 
**ParentId** | **string** |  | 
**Order** | Pointer to **string** |  | [optional] 

## Methods

### NewListCommentRepliesReq

`func NewListCommentRepliesReq(articleId string, parentId string, ) *ListCommentRepliesReq`

NewListCommentRepliesReq instantiates a new ListCommentRepliesReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCommentRepliesReqWithDefaults

`func NewListCommentRepliesReqWithDefaults() *ListCommentRepliesReq`

NewListCommentRepliesReqWithDefaults instantiates a new ListCommentRepliesReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ListCommentRepliesReq) GetPage() PageReq`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListCommentRepliesReq) GetPageOk() (*PageReq, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListCommentRepliesReq) SetPage(v PageReq)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListCommentRepliesReq) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetArticleId

`func (o *ListCommentRepliesReq) GetArticleId() string`

GetArticleId returns the ArticleId field if non-nil, zero value otherwise.

### GetArticleIdOk

`func (o *ListCommentRepliesReq) GetArticleIdOk() (*string, bool)`

GetArticleIdOk returns a tuple with the ArticleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticleId

`func (o *ListCommentRepliesReq) SetArticleId(v string)`

SetArticleId sets ArticleId field to given value.


### GetParentId

`func (o *ListCommentRepliesReq) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ListCommentRepliesReq) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ListCommentRepliesReq) SetParentId(v string)`

SetParentId sets ParentId field to given value.


### GetOrder

`func (o *ListCommentRepliesReq) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ListCommentRepliesReq) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ListCommentRepliesReq) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ListCommentRepliesReq) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


