# ArticleViewHistoryItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Article** | Pointer to [**ArticleListItem**](ArticleListItem.md) |  | [optional] 
**ViewedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewArticleViewHistoryItem

`func NewArticleViewHistoryItem() *ArticleViewHistoryItem`

NewArticleViewHistoryItem instantiates a new ArticleViewHistoryItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArticleViewHistoryItemWithDefaults

`func NewArticleViewHistoryItemWithDefaults() *ArticleViewHistoryItem`

NewArticleViewHistoryItemWithDefaults instantiates a new ArticleViewHistoryItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArticle

`func (o *ArticleViewHistoryItem) GetArticle() ArticleListItem`

GetArticle returns the Article field if non-nil, zero value otherwise.

### GetArticleOk

`func (o *ArticleViewHistoryItem) GetArticleOk() (*ArticleListItem, bool)`

GetArticleOk returns a tuple with the Article field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticle

`func (o *ArticleViewHistoryItem) SetArticle(v ArticleListItem)`

SetArticle sets Article field to given value.

### HasArticle

`func (o *ArticleViewHistoryItem) HasArticle() bool`

HasArticle returns a boolean if a field has been set.

### GetViewedAt

`func (o *ArticleViewHistoryItem) GetViewedAt() time.Time`

GetViewedAt returns the ViewedAt field if non-nil, zero value otherwise.

### GetViewedAtOk

`func (o *ArticleViewHistoryItem) GetViewedAtOk() (*time.Time, bool)`

GetViewedAtOk returns a tuple with the ViewedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewedAt

`func (o *ArticleViewHistoryItem) SetViewedAt(v time.Time)`

SetViewedAt sets ViewedAt field to given value.

### HasViewedAt

`func (o *ArticleViewHistoryItem) HasViewedAt() bool`

HasViewedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


