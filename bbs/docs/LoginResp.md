# LoginResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** |  | [optional] 
**RefreshToken** | Pointer to **string** |  | [optional] 
**AccessTokenExpiresAt** | Pointer to **time.Time** |  | [optional] 
**RefreshTokenExpiresAt** | Pointer to **time.Time** |  | [optional] 
**SessionExpiresAt** | Pointer to **time.Time** |  | [optional] 
**Account** | Pointer to [**RespAccount**](RespAccount.md) |  | [optional] 

## Methods

### NewLoginResp

`func NewLoginResp() *LoginResp`

NewLoginResp instantiates a new LoginResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginRespWithDefaults

`func NewLoginRespWithDefaults() *LoginResp`

NewLoginRespWithDefaults instantiates a new LoginResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *LoginResp) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *LoginResp) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *LoginResp) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *LoginResp) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetRefreshToken

`func (o *LoginResp) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *LoginResp) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *LoginResp) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *LoginResp) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetAccessTokenExpiresAt

`func (o *LoginResp) GetAccessTokenExpiresAt() time.Time`

GetAccessTokenExpiresAt returns the AccessTokenExpiresAt field if non-nil, zero value otherwise.

### GetAccessTokenExpiresAtOk

`func (o *LoginResp) GetAccessTokenExpiresAtOk() (*time.Time, bool)`

GetAccessTokenExpiresAtOk returns a tuple with the AccessTokenExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenExpiresAt

`func (o *LoginResp) SetAccessTokenExpiresAt(v time.Time)`

SetAccessTokenExpiresAt sets AccessTokenExpiresAt field to given value.

### HasAccessTokenExpiresAt

`func (o *LoginResp) HasAccessTokenExpiresAt() bool`

HasAccessTokenExpiresAt returns a boolean if a field has been set.

### GetRefreshTokenExpiresAt

`func (o *LoginResp) GetRefreshTokenExpiresAt() time.Time`

GetRefreshTokenExpiresAt returns the RefreshTokenExpiresAt field if non-nil, zero value otherwise.

### GetRefreshTokenExpiresAtOk

`func (o *LoginResp) GetRefreshTokenExpiresAtOk() (*time.Time, bool)`

GetRefreshTokenExpiresAtOk returns a tuple with the RefreshTokenExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenExpiresAt

`func (o *LoginResp) SetRefreshTokenExpiresAt(v time.Time)`

SetRefreshTokenExpiresAt sets RefreshTokenExpiresAt field to given value.

### HasRefreshTokenExpiresAt

`func (o *LoginResp) HasRefreshTokenExpiresAt() bool`

HasRefreshTokenExpiresAt returns a boolean if a field has been set.

### GetSessionExpiresAt

`func (o *LoginResp) GetSessionExpiresAt() time.Time`

GetSessionExpiresAt returns the SessionExpiresAt field if non-nil, zero value otherwise.

### GetSessionExpiresAtOk

`func (o *LoginResp) GetSessionExpiresAtOk() (*time.Time, bool)`

GetSessionExpiresAtOk returns a tuple with the SessionExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionExpiresAt

`func (o *LoginResp) SetSessionExpiresAt(v time.Time)`

SetSessionExpiresAt sets SessionExpiresAt field to given value.

### HasSessionExpiresAt

`func (o *LoginResp) HasSessionExpiresAt() bool`

HasSessionExpiresAt returns a boolean if a field has been set.

### GetAccount

`func (o *LoginResp) GetAccount() RespAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *LoginResp) GetAccountOk() (*RespAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *LoginResp) SetAccount(v RespAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *LoginResp) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


