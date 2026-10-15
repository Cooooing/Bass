# ArticleLottery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DrawAt** | Pointer to **time.Time** |  | [optional] 
**DrawnAt** | Pointer to **time.Time** |  | [optional] 
**EntryCount** | Pointer to **string** |  | [optional] 
**Entered** | Pointer to **bool** |  | [optional] 
**Prizes** | Pointer to [**[]ArticleLotteryPrize**](ArticleLotteryPrize.md) |  | [optional] 

## Methods

### NewArticleLottery

`func NewArticleLottery() *ArticleLottery`

NewArticleLottery instantiates a new ArticleLottery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArticleLotteryWithDefaults

`func NewArticleLotteryWithDefaults() *ArticleLottery`

NewArticleLotteryWithDefaults instantiates a new ArticleLottery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDrawAt

`func (o *ArticleLottery) GetDrawAt() time.Time`

GetDrawAt returns the DrawAt field if non-nil, zero value otherwise.

### GetDrawAtOk

`func (o *ArticleLottery) GetDrawAtOk() (*time.Time, bool)`

GetDrawAtOk returns a tuple with the DrawAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrawAt

`func (o *ArticleLottery) SetDrawAt(v time.Time)`

SetDrawAt sets DrawAt field to given value.

### HasDrawAt

`func (o *ArticleLottery) HasDrawAt() bool`

HasDrawAt returns a boolean if a field has been set.

### GetDrawnAt

`func (o *ArticleLottery) GetDrawnAt() time.Time`

GetDrawnAt returns the DrawnAt field if non-nil, zero value otherwise.

### GetDrawnAtOk

`func (o *ArticleLottery) GetDrawnAtOk() (*time.Time, bool)`

GetDrawnAtOk returns a tuple with the DrawnAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrawnAt

`func (o *ArticleLottery) SetDrawnAt(v time.Time)`

SetDrawnAt sets DrawnAt field to given value.

### HasDrawnAt

`func (o *ArticleLottery) HasDrawnAt() bool`

HasDrawnAt returns a boolean if a field has been set.

### GetEntryCount

`func (o *ArticleLottery) GetEntryCount() string`

GetEntryCount returns the EntryCount field if non-nil, zero value otherwise.

### GetEntryCountOk

`func (o *ArticleLottery) GetEntryCountOk() (*string, bool)`

GetEntryCountOk returns a tuple with the EntryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryCount

`func (o *ArticleLottery) SetEntryCount(v string)`

SetEntryCount sets EntryCount field to given value.

### HasEntryCount

`func (o *ArticleLottery) HasEntryCount() bool`

HasEntryCount returns a boolean if a field has been set.

### GetEntered

`func (o *ArticleLottery) GetEntered() bool`

GetEntered returns the Entered field if non-nil, zero value otherwise.

### GetEnteredOk

`func (o *ArticleLottery) GetEnteredOk() (*bool, bool)`

GetEnteredOk returns a tuple with the Entered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntered

`func (o *ArticleLottery) SetEntered(v bool)`

SetEntered sets Entered field to given value.

### HasEntered

`func (o *ArticleLottery) HasEntered() bool`

HasEntered returns a boolean if a field has been set.

### GetPrizes

`func (o *ArticleLottery) GetPrizes() []ArticleLotteryPrize`

GetPrizes returns the Prizes field if non-nil, zero value otherwise.

### GetPrizesOk

`func (o *ArticleLottery) GetPrizesOk() (*[]ArticleLotteryPrize, bool)`

GetPrizesOk returns a tuple with the Prizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrizes

`func (o *ArticleLottery) SetPrizes(v []ArticleLotteryPrize)`

SetPrizes sets Prizes field to given value.

### HasPrizes

`func (o *ArticleLottery) HasPrizes() bool`

HasPrizes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


