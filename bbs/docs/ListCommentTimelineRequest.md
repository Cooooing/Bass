# ListCommentTimelineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**PageRequest**](PageRequest.md) |  | [optional] 
**ArticleId** | **string** |  | 
**Order** | Pointer to **string** |  | [optional] 

## Methods

### NewListCommentTimelineRequest

`func NewListCommentTimelineRequest(articleId string, ) *ListCommentTimelineRequest`

NewListCommentTimelineRequest instantiates a new ListCommentTimelineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCommentTimelineRequestWithDefaults

`func NewListCommentTimelineRequestWithDefaults() *ListCommentTimelineRequest`

NewListCommentTimelineRequestWithDefaults instantiates a new ListCommentTimelineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ListCommentTimelineRequest) GetPage() PageRequest`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListCommentTimelineRequest) GetPageOk() (*PageRequest, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListCommentTimelineRequest) SetPage(v PageRequest)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListCommentTimelineRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetArticleId

`func (o *ListCommentTimelineRequest) GetArticleId() string`

GetArticleId returns the ArticleId field if non-nil, zero value otherwise.

### GetArticleIdOk

`func (o *ListCommentTimelineRequest) GetArticleIdOk() (*string, bool)`

GetArticleIdOk returns a tuple with the ArticleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticleId

`func (o *ListCommentTimelineRequest) SetArticleId(v string)`

SetArticleId sets ArticleId field to given value.


### GetOrder

`func (o *ListCommentTimelineRequest) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ListCommentTimelineRequest) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ListCommentTimelineRequest) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ListCommentTimelineRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


