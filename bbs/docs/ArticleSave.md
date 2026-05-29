# ArticleSave

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**RewardContent** | Pointer to **string** |  | [optional] 
**RewardPoints** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**BountyPoints** | Pointer to **int32** |  | [optional] 
**Statement** | Pointer to **string** |  | [optional] 
**Commentable** | Pointer to **bool** |  | [optional] 
**Anonymous** | Pointer to **bool** |  | [optional] 
**Listable** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to [**[]TagSave**](TagSave.md) |  | [optional] 

## Methods

### NewArticleSave

`func NewArticleSave() *ArticleSave`

NewArticleSave instantiates a new ArticleSave object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArticleSaveWithDefaults

`func NewArticleSaveWithDefaults() *ArticleSave`

NewArticleSaveWithDefaults instantiates a new ArticleSave object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ArticleSave) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArticleSave) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArticleSave) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ArticleSave) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *ArticleSave) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ArticleSave) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ArticleSave) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ArticleSave) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetContent

`func (o *ArticleSave) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ArticleSave) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ArticleSave) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ArticleSave) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetRewardContent

`func (o *ArticleSave) GetRewardContent() string`

GetRewardContent returns the RewardContent field if non-nil, zero value otherwise.

### GetRewardContentOk

`func (o *ArticleSave) GetRewardContentOk() (*string, bool)`

GetRewardContentOk returns a tuple with the RewardContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardContent

`func (o *ArticleSave) SetRewardContent(v string)`

SetRewardContent sets RewardContent field to given value.

### HasRewardContent

`func (o *ArticleSave) HasRewardContent() bool`

HasRewardContent returns a boolean if a field has been set.

### GetRewardPoints

`func (o *ArticleSave) GetRewardPoints() int32`

GetRewardPoints returns the RewardPoints field if non-nil, zero value otherwise.

### GetRewardPointsOk

`func (o *ArticleSave) GetRewardPointsOk() (*int32, bool)`

GetRewardPointsOk returns a tuple with the RewardPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardPoints

`func (o *ArticleSave) SetRewardPoints(v int32)`

SetRewardPoints sets RewardPoints field to given value.

### HasRewardPoints

`func (o *ArticleSave) HasRewardPoints() bool`

HasRewardPoints returns a boolean if a field has been set.

### GetStatus

`func (o *ArticleSave) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ArticleSave) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ArticleSave) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ArticleSave) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *ArticleSave) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArticleSave) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArticleSave) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ArticleSave) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBountyPoints

`func (o *ArticleSave) GetBountyPoints() int32`

GetBountyPoints returns the BountyPoints field if non-nil, zero value otherwise.

### GetBountyPointsOk

`func (o *ArticleSave) GetBountyPointsOk() (*int32, bool)`

GetBountyPointsOk returns a tuple with the BountyPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBountyPoints

`func (o *ArticleSave) SetBountyPoints(v int32)`

SetBountyPoints sets BountyPoints field to given value.

### HasBountyPoints

`func (o *ArticleSave) HasBountyPoints() bool`

HasBountyPoints returns a boolean if a field has been set.

### GetStatement

`func (o *ArticleSave) GetStatement() string`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *ArticleSave) GetStatementOk() (*string, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *ArticleSave) SetStatement(v string)`

SetStatement sets Statement field to given value.

### HasStatement

`func (o *ArticleSave) HasStatement() bool`

HasStatement returns a boolean if a field has been set.

### GetCommentable

`func (o *ArticleSave) GetCommentable() bool`

GetCommentable returns the Commentable field if non-nil, zero value otherwise.

### GetCommentableOk

`func (o *ArticleSave) GetCommentableOk() (*bool, bool)`

GetCommentableOk returns a tuple with the Commentable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentable

`func (o *ArticleSave) SetCommentable(v bool)`

SetCommentable sets Commentable field to given value.

### HasCommentable

`func (o *ArticleSave) HasCommentable() bool`

HasCommentable returns a boolean if a field has been set.

### GetAnonymous

`func (o *ArticleSave) GetAnonymous() bool`

GetAnonymous returns the Anonymous field if non-nil, zero value otherwise.

### GetAnonymousOk

`func (o *ArticleSave) GetAnonymousOk() (*bool, bool)`

GetAnonymousOk returns a tuple with the Anonymous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymous

`func (o *ArticleSave) SetAnonymous(v bool)`

SetAnonymous sets Anonymous field to given value.

### HasAnonymous

`func (o *ArticleSave) HasAnonymous() bool`

HasAnonymous returns a boolean if a field has been set.

### GetListable

`func (o *ArticleSave) GetListable() bool`

GetListable returns the Listable field if non-nil, zero value otherwise.

### GetListableOk

`func (o *ArticleSave) GetListableOk() (*bool, bool)`

GetListableOk returns a tuple with the Listable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListable

`func (o *ArticleSave) SetListable(v bool)`

SetListable sets Listable field to given value.

### HasListable

`func (o *ArticleSave) HasListable() bool`

HasListable returns a boolean if a field has been set.

### GetTags

`func (o *ArticleSave) GetTags() []TagSave`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ArticleSave) GetTagsOk() (*[]TagSave, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ArticleSave) SetTags(v []TagSave)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ArticleSave) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


