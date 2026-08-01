# ReqCommentQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommentId** | Pointer to **string** |  | [optional] 
**ArticleId** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**ReplyId** | Pointer to **string** |  | [optional] 
**Order** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Level** | Pointer to **int32** |  | [optional] 
**Restriction** | Pointer to **string** |  | [optional] 
**Restrictions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewReqCommentQuery

`func NewReqCommentQuery() *ReqCommentQuery`

NewReqCommentQuery instantiates a new ReqCommentQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReqCommentQueryWithDefaults

`func NewReqCommentQueryWithDefaults() *ReqCommentQuery`

NewReqCommentQueryWithDefaults instantiates a new ReqCommentQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommentId

`func (o *ReqCommentQuery) GetCommentId() string`

GetCommentId returns the CommentId field if non-nil, zero value otherwise.

### GetCommentIdOk

`func (o *ReqCommentQuery) GetCommentIdOk() (*string, bool)`

GetCommentIdOk returns a tuple with the CommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentId

`func (o *ReqCommentQuery) SetCommentId(v string)`

SetCommentId sets CommentId field to given value.

### HasCommentId

`func (o *ReqCommentQuery) HasCommentId() bool`

HasCommentId returns a boolean if a field has been set.

### GetArticleId

`func (o *ReqCommentQuery) GetArticleId() string`

GetArticleId returns the ArticleId field if non-nil, zero value otherwise.

### GetArticleIdOk

`func (o *ReqCommentQuery) GetArticleIdOk() (*string, bool)`

GetArticleIdOk returns a tuple with the ArticleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticleId

`func (o *ReqCommentQuery) SetArticleId(v string)`

SetArticleId sets ArticleId field to given value.

### HasArticleId

`func (o *ReqCommentQuery) HasArticleId() bool`

HasArticleId returns a boolean if a field has been set.

### GetParentId

`func (o *ReqCommentQuery) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ReqCommentQuery) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ReqCommentQuery) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ReqCommentQuery) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetReplyId

`func (o *ReqCommentQuery) GetReplyId() string`

GetReplyId returns the ReplyId field if non-nil, zero value otherwise.

### GetReplyIdOk

`func (o *ReqCommentQuery) GetReplyIdOk() (*string, bool)`

GetReplyIdOk returns a tuple with the ReplyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyId

`func (o *ReqCommentQuery) SetReplyId(v string)`

SetReplyId sets ReplyId field to given value.

### HasReplyId

`func (o *ReqCommentQuery) HasReplyId() bool`

HasReplyId returns a boolean if a field has been set.

### GetOrder

`func (o *ReqCommentQuery) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ReqCommentQuery) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ReqCommentQuery) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ReqCommentQuery) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetUserId

`func (o *ReqCommentQuery) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ReqCommentQuery) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ReqCommentQuery) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ReqCommentQuery) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetLevel

`func (o *ReqCommentQuery) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ReqCommentQuery) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ReqCommentQuery) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *ReqCommentQuery) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetRestriction

`func (o *ReqCommentQuery) GetRestriction() string`

GetRestriction returns the Restriction field if non-nil, zero value otherwise.

### GetRestrictionOk

`func (o *ReqCommentQuery) GetRestrictionOk() (*string, bool)`

GetRestrictionOk returns a tuple with the Restriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestriction

`func (o *ReqCommentQuery) SetRestriction(v string)`

SetRestriction sets Restriction field to given value.

### HasRestriction

`func (o *ReqCommentQuery) HasRestriction() bool`

HasRestriction returns a boolean if a field has been set.

### GetRestrictions

`func (o *ReqCommentQuery) GetRestrictions() []string`

GetRestrictions returns the Restrictions field if non-nil, zero value otherwise.

### GetRestrictionsOk

`func (o *ReqCommentQuery) GetRestrictionsOk() (*[]string, bool)`

GetRestrictionsOk returns a tuple with the Restrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictions

`func (o *ReqCommentQuery) SetRestrictions(v []string)`

SetRestrictions sets Restrictions field to given value.

### HasRestrictions

`func (o *ReqCommentQuery) HasRestrictions() bool`

HasRestrictions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


