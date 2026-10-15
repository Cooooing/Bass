# ArticleLotteryPrize

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**WinnerLimit** | Pointer to **int32** |  | [optional] 
**WinnerUsers** | Pointer to [**[]AccountProfile**](AccountProfile.md) |  | [optional] 

## Methods

### NewArticleLotteryPrize

`func NewArticleLotteryPrize() *ArticleLotteryPrize`

NewArticleLotteryPrize instantiates a new ArticleLotteryPrize object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArticleLotteryPrizeWithDefaults

`func NewArticleLotteryPrizeWithDefaults() *ArticleLotteryPrize`

NewArticleLotteryPrizeWithDefaults instantiates a new ArticleLotteryPrize object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ArticleLotteryPrize) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArticleLotteryPrize) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArticleLotteryPrize) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ArticleLotteryPrize) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ArticleLotteryPrize) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArticleLotteryPrize) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArticleLotteryPrize) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ArticleLotteryPrize) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ArticleLotteryPrize) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ArticleLotteryPrize) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ArticleLotteryPrize) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ArticleLotteryPrize) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetWinnerLimit

`func (o *ArticleLotteryPrize) GetWinnerLimit() int32`

GetWinnerLimit returns the WinnerLimit field if non-nil, zero value otherwise.

### GetWinnerLimitOk

`func (o *ArticleLotteryPrize) GetWinnerLimitOk() (*int32, bool)`

GetWinnerLimitOk returns a tuple with the WinnerLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinnerLimit

`func (o *ArticleLotteryPrize) SetWinnerLimit(v int32)`

SetWinnerLimit sets WinnerLimit field to given value.

### HasWinnerLimit

`func (o *ArticleLotteryPrize) HasWinnerLimit() bool`

HasWinnerLimit returns a boolean if a field has been set.

### GetWinnerUsers

`func (o *ArticleLotteryPrize) GetWinnerUsers() []AccountProfile`

GetWinnerUsers returns the WinnerUsers field if non-nil, zero value otherwise.

### GetWinnerUsersOk

`func (o *ArticleLotteryPrize) GetWinnerUsersOk() (*[]AccountProfile, bool)`

GetWinnerUsersOk returns a tuple with the WinnerUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinnerUsers

`func (o *ArticleLotteryPrize) SetWinnerUsers(v []AccountProfile)`

SetWinnerUsers sets WinnerUsers field to given value.

### HasWinnerUsers

`func (o *ArticleLotteryPrize) HasWinnerUsers() bool`

HasWinnerUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


