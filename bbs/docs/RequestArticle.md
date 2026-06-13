# RequestArticle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Content** | **string** |  | 
**RewardContent** | Pointer to **string** |  | [optional] 
**RewardPoints** | Pointer to **int32** |  | [optional] 
**Type** | **string** |  | 
**BountyPoints** | Pointer to **int32** |  | [optional] 
**Statement** | Pointer to **string** |  | [optional] 
**Commentable** | Pointer to **bool** |  | [optional] 
**Anonymous** | Pointer to **bool** |  | [optional] 
**TagIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewRequestArticle

`func NewRequestArticle(title string, content string, type_ string, ) *RequestArticle`

NewRequestArticle instantiates a new RequestArticle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestArticleWithDefaults

`func NewRequestArticleWithDefaults() *RequestArticle`

NewRequestArticleWithDefaults instantiates a new RequestArticle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *RequestArticle) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RequestArticle) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RequestArticle) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetContent

`func (o *RequestArticle) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *RequestArticle) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *RequestArticle) SetContent(v string)`

SetContent sets Content field to given value.


### GetRewardContent

`func (o *RequestArticle) GetRewardContent() string`

GetRewardContent returns the RewardContent field if non-nil, zero value otherwise.

### GetRewardContentOk

`func (o *RequestArticle) GetRewardContentOk() (*string, bool)`

GetRewardContentOk returns a tuple with the RewardContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardContent

`func (o *RequestArticle) SetRewardContent(v string)`

SetRewardContent sets RewardContent field to given value.

### HasRewardContent

`func (o *RequestArticle) HasRewardContent() bool`

HasRewardContent returns a boolean if a field has been set.

### GetRewardPoints

`func (o *RequestArticle) GetRewardPoints() int32`

GetRewardPoints returns the RewardPoints field if non-nil, zero value otherwise.

### GetRewardPointsOk

`func (o *RequestArticle) GetRewardPointsOk() (*int32, bool)`

GetRewardPointsOk returns a tuple with the RewardPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardPoints

`func (o *RequestArticle) SetRewardPoints(v int32)`

SetRewardPoints sets RewardPoints field to given value.

### HasRewardPoints

`func (o *RequestArticle) HasRewardPoints() bool`

HasRewardPoints returns a boolean if a field has been set.

### GetType

`func (o *RequestArticle) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RequestArticle) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RequestArticle) SetType(v string)`

SetType sets Type field to given value.


### GetBountyPoints

`func (o *RequestArticle) GetBountyPoints() int32`

GetBountyPoints returns the BountyPoints field if non-nil, zero value otherwise.

### GetBountyPointsOk

`func (o *RequestArticle) GetBountyPointsOk() (*int32, bool)`

GetBountyPointsOk returns a tuple with the BountyPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBountyPoints

`func (o *RequestArticle) SetBountyPoints(v int32)`

SetBountyPoints sets BountyPoints field to given value.

### HasBountyPoints

`func (o *RequestArticle) HasBountyPoints() bool`

HasBountyPoints returns a boolean if a field has been set.

### GetStatement

`func (o *RequestArticle) GetStatement() string`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *RequestArticle) GetStatementOk() (*string, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *RequestArticle) SetStatement(v string)`

SetStatement sets Statement field to given value.

### HasStatement

`func (o *RequestArticle) HasStatement() bool`

HasStatement returns a boolean if a field has been set.

### GetCommentable

`func (o *RequestArticle) GetCommentable() bool`

GetCommentable returns the Commentable field if non-nil, zero value otherwise.

### GetCommentableOk

`func (o *RequestArticle) GetCommentableOk() (*bool, bool)`

GetCommentableOk returns a tuple with the Commentable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentable

`func (o *RequestArticle) SetCommentable(v bool)`

SetCommentable sets Commentable field to given value.

### HasCommentable

`func (o *RequestArticle) HasCommentable() bool`

HasCommentable returns a boolean if a field has been set.

### GetAnonymous

`func (o *RequestArticle) GetAnonymous() bool`

GetAnonymous returns the Anonymous field if non-nil, zero value otherwise.

### GetAnonymousOk

`func (o *RequestArticle) GetAnonymousOk() (*bool, bool)`

GetAnonymousOk returns a tuple with the Anonymous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymous

`func (o *RequestArticle) SetAnonymous(v bool)`

SetAnonymous sets Anonymous field to given value.

### HasAnonymous

`func (o *RequestArticle) HasAnonymous() bool`

HasAnonymous returns a boolean if a field has been set.

### GetTagIds

`func (o *RequestArticle) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *RequestArticle) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *RequestArticle) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *RequestArticle) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


