# CreateCommentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArticleId** | **string** |  | 
**Content** | **string** |  | 
**ReplyId** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateCommentRequest

`func NewCreateCommentRequest(articleId string, content string, ) *CreateCommentRequest`

NewCreateCommentRequest instantiates a new CreateCommentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCommentRequestWithDefaults

`func NewCreateCommentRequestWithDefaults() *CreateCommentRequest`

NewCreateCommentRequestWithDefaults instantiates a new CreateCommentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArticleId

`func (o *CreateCommentRequest) GetArticleId() string`

GetArticleId returns the ArticleId field if non-nil, zero value otherwise.

### GetArticleIdOk

`func (o *CreateCommentRequest) GetArticleIdOk() (*string, bool)`

GetArticleIdOk returns a tuple with the ArticleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticleId

`func (o *CreateCommentRequest) SetArticleId(v string)`

SetArticleId sets ArticleId field to given value.


### GetContent

`func (o *CreateCommentRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CreateCommentRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CreateCommentRequest) SetContent(v string)`

SetContent sets Content field to given value.


### GetReplyId

`func (o *CreateCommentRequest) GetReplyId() string`

GetReplyId returns the ReplyId field if non-nil, zero value otherwise.

### GetReplyIdOk

`func (o *CreateCommentRequest) GetReplyIdOk() (*string, bool)`

GetReplyIdOk returns a tuple with the ReplyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyId

`func (o *CreateCommentRequest) SetReplyId(v string)`

SetReplyId sets ReplyId field to given value.

### HasReplyId

`func (o *CreateCommentRequest) HasReplyId() bool`

HasReplyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


