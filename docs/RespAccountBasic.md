# RespAccountBasic

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
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRespAccountBasic

`func NewRespAccountBasic() *RespAccountBasic`

NewRespAccountBasic instantiates a new RespAccountBasic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRespAccountBasicWithDefaults

`func NewRespAccountBasicWithDefaults() *RespAccountBasic`

NewRespAccountBasicWithDefaults instantiates a new RespAccountBasic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RespAccountBasic) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RespAccountBasic) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RespAccountBasic) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RespAccountBasic) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RespAccountBasic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RespAccountBasic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RespAccountBasic) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RespAccountBasic) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNickname

`func (o *RespAccountBasic) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *RespAccountBasic) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *RespAccountBasic) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *RespAccountBasic) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetUrl

`func (o *RespAccountBasic) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RespAccountBasic) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RespAccountBasic) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *RespAccountBasic) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *RespAccountBasic) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *RespAccountBasic) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *RespAccountBasic) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *RespAccountBasic) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetIntroduction

`func (o *RespAccountBasic) GetIntroduction() string`

GetIntroduction returns the Introduction field if non-nil, zero value otherwise.

### GetIntroductionOk

`func (o *RespAccountBasic) GetIntroductionOk() (*string, bool)`

GetIntroductionOk returns a tuple with the Introduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntroduction

`func (o *RespAccountBasic) SetIntroduction(v string)`

SetIntroduction sets Introduction field to given value.

### HasIntroduction

`func (o *RespAccountBasic) HasIntroduction() bool`

HasIntroduction returns a boolean if a field has been set.

### GetMbti

`func (o *RespAccountBasic) GetMbti() string`

GetMbti returns the Mbti field if non-nil, zero value otherwise.

### GetMbtiOk

`func (o *RespAccountBasic) GetMbtiOk() (*string, bool)`

GetMbtiOk returns a tuple with the Mbti field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbti

`func (o *RespAccountBasic) SetMbti(v string)`

SetMbti sets Mbti field to given value.

### HasMbti

`func (o *RespAccountBasic) HasMbti() bool`

HasMbti returns a boolean if a field has been set.

### GetStatus

`func (o *RespAccountBasic) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RespAccountBasic) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RespAccountBasic) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RespAccountBasic) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFollowCount

`func (o *RespAccountBasic) GetFollowCount() int32`

GetFollowCount returns the FollowCount field if non-nil, zero value otherwise.

### GetFollowCountOk

`func (o *RespAccountBasic) GetFollowCountOk() (*int32, bool)`

GetFollowCountOk returns a tuple with the FollowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowCount

`func (o *RespAccountBasic) SetFollowCount(v int32)`

SetFollowCount sets FollowCount field to given value.

### HasFollowCount

`func (o *RespAccountBasic) HasFollowCount() bool`

HasFollowCount returns a boolean if a field has been set.

### GetFollowerCount

`func (o *RespAccountBasic) GetFollowerCount() int32`

GetFollowerCount returns the FollowerCount field if non-nil, zero value otherwise.

### GetFollowerCountOk

`func (o *RespAccountBasic) GetFollowerCountOk() (*int32, bool)`

GetFollowerCountOk returns a tuple with the FollowerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerCount

`func (o *RespAccountBasic) SetFollowerCount(v int32)`

SetFollowerCount sets FollowerCount field to given value.

### HasFollowerCount

`func (o *RespAccountBasic) HasFollowerCount() bool`

HasFollowerCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RespAccountBasic) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RespAccountBasic) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RespAccountBasic) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RespAccountBasic) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RespAccountBasic) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RespAccountBasic) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RespAccountBasic) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RespAccountBasic) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


