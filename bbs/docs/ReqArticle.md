# ReqArticle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Content** | **string** |  | 
**RewardContent** | Pointer to **string** |  | [optional] 
**RewardPoints** | Pointer to **int32** |  | [optional] 
**Type** | **string** |  | 
**Statement** | Pointer to **string** |  | [optional] 
**Commentable** | Pointer to **bool** |  | [optional] 

## Methods

### NewReqArticle

`func NewReqArticle(title string, content string, type_ string, ) *ReqArticle`

NewReqArticle instantiates a new ReqArticle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReqArticleWithDefaults

`func NewReqArticleWithDefaults() *ReqArticle`

NewReqArticleWithDefaults instantiates a new ReqArticle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ReqArticle) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ReqArticle) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ReqArticle) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetContent

`func (o *ReqArticle) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ReqArticle) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ReqArticle) SetContent(v string)`

SetContent sets Content field to given value.


### GetRewardContent

`func (o *ReqArticle) GetRewardContent() string`

GetRewardContent returns the RewardContent field if non-nil, zero value otherwise.

### GetRewardContentOk

`func (o *ReqArticle) GetRewardContentOk() (*string, bool)`

GetRewardContentOk returns a tuple with the RewardContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardContent

`func (o *ReqArticle) SetRewardContent(v string)`

SetRewardContent sets RewardContent field to given value.

### HasRewardContent

`func (o *ReqArticle) HasRewardContent() bool`

HasRewardContent returns a boolean if a field has been set.

### GetRewardPoints

`func (o *ReqArticle) GetRewardPoints() int32`

GetRewardPoints returns the RewardPoints field if non-nil, zero value otherwise.

### GetRewardPointsOk

`func (o *ReqArticle) GetRewardPointsOk() (*int32, bool)`

GetRewardPointsOk returns a tuple with the RewardPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardPoints

`func (o *ReqArticle) SetRewardPoints(v int32)`

SetRewardPoints sets RewardPoints field to given value.

### HasRewardPoints

`func (o *ReqArticle) HasRewardPoints() bool`

HasRewardPoints returns a boolean if a field has been set.

### GetType

`func (o *ReqArticle) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReqArticle) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReqArticle) SetType(v string)`

SetType sets Type field to given value.


### GetStatement

`func (o *ReqArticle) GetStatement() string`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *ReqArticle) GetStatementOk() (*string, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *ReqArticle) SetStatement(v string)`

SetStatement sets Statement field to given value.

### HasStatement

`func (o *ReqArticle) HasStatement() bool`

HasStatement returns a boolean if a field has been set.

### GetCommentable

`func (o *ReqArticle) GetCommentable() bool`

GetCommentable returns the Commentable field if non-nil, zero value otherwise.

### GetCommentableOk

`func (o *ReqArticle) GetCommentableOk() (*bool, bool)`

GetCommentableOk returns a tuple with the Commentable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentable

`func (o *ReqArticle) SetCommentable(v bool)`

SetCommentable sets Commentable field to given value.

### HasCommentable

`func (o *ReqArticle) HasCommentable() bool`

HasCommentable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


