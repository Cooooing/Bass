# AccountProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Nickname** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**AvatarUrl** | Pointer to **string** |  | [optional] 
**Introduction** | Pointer to **string** |  | [optional] 
**Mbti** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**FollowCount** | Pointer to **int32** |  | [optional] 
**FollowerCount** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewAccountProfile

`func NewAccountProfile() *AccountProfile`

NewAccountProfile instantiates a new AccountProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountProfileWithDefaults

`func NewAccountProfileWithDefaults() *AccountProfile`

NewAccountProfileWithDefaults instantiates a new AccountProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountProfile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AccountProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNickname

`func (o *AccountProfile) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *AccountProfile) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *AccountProfile) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *AccountProfile) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetUrl

`func (o *AccountProfile) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AccountProfile) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AccountProfile) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AccountProfile) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *AccountProfile) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *AccountProfile) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *AccountProfile) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *AccountProfile) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetIntroduction

`func (o *AccountProfile) GetIntroduction() string`

GetIntroduction returns the Introduction field if non-nil, zero value otherwise.

### GetIntroductionOk

`func (o *AccountProfile) GetIntroductionOk() (*string, bool)`

GetIntroductionOk returns a tuple with the Introduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntroduction

`func (o *AccountProfile) SetIntroduction(v string)`

SetIntroduction sets Introduction field to given value.

### HasIntroduction

`func (o *AccountProfile) HasIntroduction() bool`

HasIntroduction returns a boolean if a field has been set.

### GetMbti

`func (o *AccountProfile) GetMbti() string`

GetMbti returns the Mbti field if non-nil, zero value otherwise.

### GetMbtiOk

`func (o *AccountProfile) GetMbtiOk() (*string, bool)`

GetMbtiOk returns a tuple with the Mbti field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbti

`func (o *AccountProfile) SetMbti(v string)`

SetMbti sets Mbti field to given value.

### HasMbti

`func (o *AccountProfile) HasMbti() bool`

HasMbti returns a boolean if a field has been set.

### GetStatus

`func (o *AccountProfile) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountProfile) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountProfile) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccountProfile) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFollowCount

`func (o *AccountProfile) GetFollowCount() int32`

GetFollowCount returns the FollowCount field if non-nil, zero value otherwise.

### GetFollowCountOk

`func (o *AccountProfile) GetFollowCountOk() (*int32, bool)`

GetFollowCountOk returns a tuple with the FollowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowCount

`func (o *AccountProfile) SetFollowCount(v int32)`

SetFollowCount sets FollowCount field to given value.

### HasFollowCount

`func (o *AccountProfile) HasFollowCount() bool`

HasFollowCount returns a boolean if a field has been set.

### GetFollowerCount

`func (o *AccountProfile) GetFollowerCount() int32`

GetFollowerCount returns the FollowerCount field if non-nil, zero value otherwise.

### GetFollowerCountOk

`func (o *AccountProfile) GetFollowerCountOk() (*int32, bool)`

GetFollowerCountOk returns a tuple with the FollowerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerCount

`func (o *AccountProfile) SetFollowerCount(v int32)`

SetFollowerCount sets FollowerCount field to given value.

### HasFollowerCount

`func (o *AccountProfile) HasFollowerCount() bool`

HasFollowerCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AccountProfile) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccountProfile) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccountProfile) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AccountProfile) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AccountProfile) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AccountProfile) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AccountProfile) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AccountProfile) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


