# RewardArticleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArticleId** | **string** |  | 
**Points** | Pointer to **int32** |  | [optional] 

## Methods

### NewRewardArticleRequest

`func NewRewardArticleRequest(articleId string, ) *RewardArticleRequest`

NewRewardArticleRequest instantiates a new RewardArticleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardArticleRequestWithDefaults

`func NewRewardArticleRequestWithDefaults() *RewardArticleRequest`

NewRewardArticleRequestWithDefaults instantiates a new RewardArticleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArticleId

`func (o *RewardArticleRequest) GetArticleId() string`

GetArticleId returns the ArticleId field if non-nil, zero value otherwise.

### GetArticleIdOk

`func (o *RewardArticleRequest) GetArticleIdOk() (*string, bool)`

GetArticleIdOk returns a tuple with the ArticleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticleId

`func (o *RewardArticleRequest) SetArticleId(v string)`

SetArticleId sets ArticleId field to given value.


### GetPoints

`func (o *RewardArticleRequest) GetPoints() int32`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *RewardArticleRequest) GetPointsOk() (*int32, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *RewardArticleRequest) SetPoints(v int32)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *RewardArticleRequest) HasPoints() bool`

HasPoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


