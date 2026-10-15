# VoteArticlePollReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArticleId** | **string** |  | 
**OptionIds** | **[]string** |  | 

## Methods

### NewVoteArticlePollReq

`func NewVoteArticlePollReq(articleId string, optionIds []string, ) *VoteArticlePollReq`

NewVoteArticlePollReq instantiates a new VoteArticlePollReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoteArticlePollReqWithDefaults

`func NewVoteArticlePollReqWithDefaults() *VoteArticlePollReq`

NewVoteArticlePollReqWithDefaults instantiates a new VoteArticlePollReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArticleId

`func (o *VoteArticlePollReq) GetArticleId() string`

GetArticleId returns the ArticleId field if non-nil, zero value otherwise.

### GetArticleIdOk

`func (o *VoteArticlePollReq) GetArticleIdOk() (*string, bool)`

GetArticleIdOk returns a tuple with the ArticleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticleId

`func (o *VoteArticlePollReq) SetArticleId(v string)`

SetArticleId sets ArticleId field to given value.


### GetOptionIds

`func (o *VoteArticlePollReq) GetOptionIds() []string`

GetOptionIds returns the OptionIds field if non-nil, zero value otherwise.

### GetOptionIdsOk

`func (o *VoteArticlePollReq) GetOptionIdsOk() (*[]string, bool)`

GetOptionIdsOk returns a tuple with the OptionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionIds

`func (o *VoteArticlePollReq) SetOptionIds(v []string)`

SetOptionIds sets OptionIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


