# ArticleLotteryDraft

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prizes** | Pointer to [**[]ArticleLotteryPrizeDraft**](ArticleLotteryPrizeDraft.md) |  | [optional] 
**DrawAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewArticleLotteryDraft

`func NewArticleLotteryDraft() *ArticleLotteryDraft`

NewArticleLotteryDraft instantiates a new ArticleLotteryDraft object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArticleLotteryDraftWithDefaults

`func NewArticleLotteryDraftWithDefaults() *ArticleLotteryDraft`

NewArticleLotteryDraftWithDefaults instantiates a new ArticleLotteryDraft object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrizes

`func (o *ArticleLotteryDraft) GetPrizes() []ArticleLotteryPrizeDraft`

GetPrizes returns the Prizes field if non-nil, zero value otherwise.

### GetPrizesOk

`func (o *ArticleLotteryDraft) GetPrizesOk() (*[]ArticleLotteryPrizeDraft, bool)`

GetPrizesOk returns a tuple with the Prizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrizes

`func (o *ArticleLotteryDraft) SetPrizes(v []ArticleLotteryPrizeDraft)`

SetPrizes sets Prizes field to given value.

### HasPrizes

`func (o *ArticleLotteryDraft) HasPrizes() bool`

HasPrizes returns a boolean if a field has been set.

### GetDrawAt

`func (o *ArticleLotteryDraft) GetDrawAt() time.Time`

GetDrawAt returns the DrawAt field if non-nil, zero value otherwise.

### GetDrawAtOk

`func (o *ArticleLotteryDraft) GetDrawAtOk() (*time.Time, bool)`

GetDrawAtOk returns a tuple with the DrawAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrawAt

`func (o *ArticleLotteryDraft) SetDrawAt(v time.Time)`

SetDrawAt sets DrawAt field to given value.

### HasDrawAt

`func (o *ArticleLotteryDraft) HasDrawAt() bool`

HasDrawAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


