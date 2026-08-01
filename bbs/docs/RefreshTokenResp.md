# RefreshTokenResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** |  | [optional] 
**RefreshToken** | Pointer to **string** |  | [optional] 
**AccessTokenExpiresAt** | Pointer to **time.Time** |  | [optional] 
**RefreshTokenExpiresAt** | Pointer to **time.Time** |  | [optional] 
**SessionExpiresAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRefreshTokenResp

`func NewRefreshTokenResp() *RefreshTokenResp`

NewRefreshTokenResp instantiates a new RefreshTokenResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefreshTokenRespWithDefaults

`func NewRefreshTokenRespWithDefaults() *RefreshTokenResp`

NewRefreshTokenRespWithDefaults instantiates a new RefreshTokenResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *RefreshTokenResp) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *RefreshTokenResp) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *RefreshTokenResp) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *RefreshTokenResp) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetRefreshToken

`func (o *RefreshTokenResp) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *RefreshTokenResp) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *RefreshTokenResp) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *RefreshTokenResp) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetAccessTokenExpiresAt

`func (o *RefreshTokenResp) GetAccessTokenExpiresAt() time.Time`

GetAccessTokenExpiresAt returns the AccessTokenExpiresAt field if non-nil, zero value otherwise.

### GetAccessTokenExpiresAtOk

`func (o *RefreshTokenResp) GetAccessTokenExpiresAtOk() (*time.Time, bool)`

GetAccessTokenExpiresAtOk returns a tuple with the AccessTokenExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenExpiresAt

`func (o *RefreshTokenResp) SetAccessTokenExpiresAt(v time.Time)`

SetAccessTokenExpiresAt sets AccessTokenExpiresAt field to given value.

### HasAccessTokenExpiresAt

`func (o *RefreshTokenResp) HasAccessTokenExpiresAt() bool`

HasAccessTokenExpiresAt returns a boolean if a field has been set.

### GetRefreshTokenExpiresAt

`func (o *RefreshTokenResp) GetRefreshTokenExpiresAt() time.Time`

GetRefreshTokenExpiresAt returns the RefreshTokenExpiresAt field if non-nil, zero value otherwise.

### GetRefreshTokenExpiresAtOk

`func (o *RefreshTokenResp) GetRefreshTokenExpiresAtOk() (*time.Time, bool)`

GetRefreshTokenExpiresAtOk returns a tuple with the RefreshTokenExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenExpiresAt

`func (o *RefreshTokenResp) SetRefreshTokenExpiresAt(v time.Time)`

SetRefreshTokenExpiresAt sets RefreshTokenExpiresAt field to given value.

### HasRefreshTokenExpiresAt

`func (o *RefreshTokenResp) HasRefreshTokenExpiresAt() bool`

HasRefreshTokenExpiresAt returns a boolean if a field has been set.

### GetSessionExpiresAt

`func (o *RefreshTokenResp) GetSessionExpiresAt() time.Time`

GetSessionExpiresAt returns the SessionExpiresAt field if non-nil, zero value otherwise.

### GetSessionExpiresAtOk

`func (o *RefreshTokenResp) GetSessionExpiresAtOk() (*time.Time, bool)`

GetSessionExpiresAtOk returns a tuple with the SessionExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionExpiresAt

`func (o *RefreshTokenResp) SetSessionExpiresAt(v time.Time)`

SetSessionExpiresAt sets SessionExpiresAt field to given value.

### HasSessionExpiresAt

`func (o *RefreshTokenResp) HasSessionExpiresAt() bool`

HasSessionExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


