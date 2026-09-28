# LoginAccountResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** |  | [optional] 
**RefreshToken** | Pointer to **string** |  | [optional] 
**AccessTokenExpiresAt** | Pointer to **time.Time** |  | [optional] 
**RefreshTokenExpiresAt** | Pointer to **time.Time** |  | [optional] 
**SessionExpiresAt** | Pointer to **time.Time** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Nickname** | Pointer to **string** |  | [optional] 

## Methods

### NewLoginAccountResp

`func NewLoginAccountResp() *LoginAccountResp`

NewLoginAccountResp instantiates a new LoginAccountResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginAccountRespWithDefaults

`func NewLoginAccountRespWithDefaults() *LoginAccountResp`

NewLoginAccountRespWithDefaults instantiates a new LoginAccountResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *LoginAccountResp) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *LoginAccountResp) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *LoginAccountResp) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *LoginAccountResp) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetRefreshToken

`func (o *LoginAccountResp) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *LoginAccountResp) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *LoginAccountResp) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *LoginAccountResp) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetAccessTokenExpiresAt

`func (o *LoginAccountResp) GetAccessTokenExpiresAt() time.Time`

GetAccessTokenExpiresAt returns the AccessTokenExpiresAt field if non-nil, zero value otherwise.

### GetAccessTokenExpiresAtOk

`func (o *LoginAccountResp) GetAccessTokenExpiresAtOk() (*time.Time, bool)`

GetAccessTokenExpiresAtOk returns a tuple with the AccessTokenExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenExpiresAt

`func (o *LoginAccountResp) SetAccessTokenExpiresAt(v time.Time)`

SetAccessTokenExpiresAt sets AccessTokenExpiresAt field to given value.

### HasAccessTokenExpiresAt

`func (o *LoginAccountResp) HasAccessTokenExpiresAt() bool`

HasAccessTokenExpiresAt returns a boolean if a field has been set.

### GetRefreshTokenExpiresAt

`func (o *LoginAccountResp) GetRefreshTokenExpiresAt() time.Time`

GetRefreshTokenExpiresAt returns the RefreshTokenExpiresAt field if non-nil, zero value otherwise.

### GetRefreshTokenExpiresAtOk

`func (o *LoginAccountResp) GetRefreshTokenExpiresAtOk() (*time.Time, bool)`

GetRefreshTokenExpiresAtOk returns a tuple with the RefreshTokenExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenExpiresAt

`func (o *LoginAccountResp) SetRefreshTokenExpiresAt(v time.Time)`

SetRefreshTokenExpiresAt sets RefreshTokenExpiresAt field to given value.

### HasRefreshTokenExpiresAt

`func (o *LoginAccountResp) HasRefreshTokenExpiresAt() bool`

HasRefreshTokenExpiresAt returns a boolean if a field has been set.

### GetSessionExpiresAt

`func (o *LoginAccountResp) GetSessionExpiresAt() time.Time`

GetSessionExpiresAt returns the SessionExpiresAt field if non-nil, zero value otherwise.

### GetSessionExpiresAtOk

`func (o *LoginAccountResp) GetSessionExpiresAtOk() (*time.Time, bool)`

GetSessionExpiresAtOk returns a tuple with the SessionExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionExpiresAt

`func (o *LoginAccountResp) SetSessionExpiresAt(v time.Time)`

SetSessionExpiresAt sets SessionExpiresAt field to given value.

### HasSessionExpiresAt

`func (o *LoginAccountResp) HasSessionExpiresAt() bool`

HasSessionExpiresAt returns a boolean if a field has been set.

### GetUserId

`func (o *LoginAccountResp) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *LoginAccountResp) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *LoginAccountResp) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *LoginAccountResp) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetName

`func (o *LoginAccountResp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoginAccountResp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoginAccountResp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LoginAccountResp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNickname

`func (o *LoginAccountResp) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *LoginAccountResp) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *LoginAccountResp) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *LoginAccountResp) HasNickname() bool`

HasNickname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


