# ArticlePoll

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndsAt** | Pointer to **time.Time** |  | [optional] 
**MaxChoices** | Pointer to **int32** |  | [optional] 
**Options** | Pointer to [**[]ArticlePollOption**](ArticlePollOption.md) |  | [optional] 
**SelectedOptionIds** | Pointer to **[]string** |  | [optional] 
**ResultVisible** | Pointer to **bool** |  | [optional] 
**ParticipantCount** | Pointer to **string** |  | [optional] 
**TotalVoteCount** | Pointer to **string** |  | [optional] 

## Methods

### NewArticlePoll

`func NewArticlePoll() *ArticlePoll`

NewArticlePoll instantiates a new ArticlePoll object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArticlePollWithDefaults

`func NewArticlePollWithDefaults() *ArticlePoll`

NewArticlePollWithDefaults instantiates a new ArticlePoll object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndsAt

`func (o *ArticlePoll) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *ArticlePoll) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *ArticlePoll) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.

### HasEndsAt

`func (o *ArticlePoll) HasEndsAt() bool`

HasEndsAt returns a boolean if a field has been set.

### GetMaxChoices

`func (o *ArticlePoll) GetMaxChoices() int32`

GetMaxChoices returns the MaxChoices field if non-nil, zero value otherwise.

### GetMaxChoicesOk

`func (o *ArticlePoll) GetMaxChoicesOk() (*int32, bool)`

GetMaxChoicesOk returns a tuple with the MaxChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxChoices

`func (o *ArticlePoll) SetMaxChoices(v int32)`

SetMaxChoices sets MaxChoices field to given value.

### HasMaxChoices

`func (o *ArticlePoll) HasMaxChoices() bool`

HasMaxChoices returns a boolean if a field has been set.

### GetOptions

`func (o *ArticlePoll) GetOptions() []ArticlePollOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ArticlePoll) GetOptionsOk() (*[]ArticlePollOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ArticlePoll) SetOptions(v []ArticlePollOption)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ArticlePoll) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetSelectedOptionIds

`func (o *ArticlePoll) GetSelectedOptionIds() []string`

GetSelectedOptionIds returns the SelectedOptionIds field if non-nil, zero value otherwise.

### GetSelectedOptionIdsOk

`func (o *ArticlePoll) GetSelectedOptionIdsOk() (*[]string, bool)`

GetSelectedOptionIdsOk returns a tuple with the SelectedOptionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedOptionIds

`func (o *ArticlePoll) SetSelectedOptionIds(v []string)`

SetSelectedOptionIds sets SelectedOptionIds field to given value.

### HasSelectedOptionIds

`func (o *ArticlePoll) HasSelectedOptionIds() bool`

HasSelectedOptionIds returns a boolean if a field has been set.

### GetResultVisible

`func (o *ArticlePoll) GetResultVisible() bool`

GetResultVisible returns the ResultVisible field if non-nil, zero value otherwise.

### GetResultVisibleOk

`func (o *ArticlePoll) GetResultVisibleOk() (*bool, bool)`

GetResultVisibleOk returns a tuple with the ResultVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultVisible

`func (o *ArticlePoll) SetResultVisible(v bool)`

SetResultVisible sets ResultVisible field to given value.

### HasResultVisible

`func (o *ArticlePoll) HasResultVisible() bool`

HasResultVisible returns a boolean if a field has been set.

### GetParticipantCount

`func (o *ArticlePoll) GetParticipantCount() string`

GetParticipantCount returns the ParticipantCount field if non-nil, zero value otherwise.

### GetParticipantCountOk

`func (o *ArticlePoll) GetParticipantCountOk() (*string, bool)`

GetParticipantCountOk returns a tuple with the ParticipantCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantCount

`func (o *ArticlePoll) SetParticipantCount(v string)`

SetParticipantCount sets ParticipantCount field to given value.

### HasParticipantCount

`func (o *ArticlePoll) HasParticipantCount() bool`

HasParticipantCount returns a boolean if a field has been set.

### GetTotalVoteCount

`func (o *ArticlePoll) GetTotalVoteCount() string`

GetTotalVoteCount returns the TotalVoteCount field if non-nil, zero value otherwise.

### GetTotalVoteCountOk

`func (o *ArticlePoll) GetTotalVoteCountOk() (*string, bool)`

GetTotalVoteCountOk returns a tuple with the TotalVoteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVoteCount

`func (o *ArticlePoll) SetTotalVoteCount(v string)`

SetTotalVoteCount sets TotalVoteCount field to given value.

### HasTotalVoteCount

`func (o *ArticlePoll) HasTotalVoteCount() bool`

HasTotalVoteCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


