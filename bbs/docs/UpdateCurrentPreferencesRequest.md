# UpdateCurrentPreferencesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timezone** | Pointer to **string** | 时区。 | [optional] 
**Theme** | Pointer to **string** | 桌面端主题。 | [optional] 
**MobileTheme** | Pointer to **string** | 移动端主题。 | [optional] 
**Language** | Pointer to **string** | 界面语言。 | [optional] 

## Methods

### NewUpdateCurrentPreferencesRequest

`func NewUpdateCurrentPreferencesRequest() *UpdateCurrentPreferencesRequest`

NewUpdateCurrentPreferencesRequest instantiates a new UpdateCurrentPreferencesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCurrentPreferencesRequestWithDefaults

`func NewUpdateCurrentPreferencesRequestWithDefaults() *UpdateCurrentPreferencesRequest`

NewUpdateCurrentPreferencesRequestWithDefaults instantiates a new UpdateCurrentPreferencesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimezone

`func (o *UpdateCurrentPreferencesRequest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *UpdateCurrentPreferencesRequest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *UpdateCurrentPreferencesRequest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *UpdateCurrentPreferencesRequest) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetTheme

`func (o *UpdateCurrentPreferencesRequest) GetTheme() string`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *UpdateCurrentPreferencesRequest) GetThemeOk() (*string, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *UpdateCurrentPreferencesRequest) SetTheme(v string)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *UpdateCurrentPreferencesRequest) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### GetMobileTheme

`func (o *UpdateCurrentPreferencesRequest) GetMobileTheme() string`

GetMobileTheme returns the MobileTheme field if non-nil, zero value otherwise.

### GetMobileThemeOk

`func (o *UpdateCurrentPreferencesRequest) GetMobileThemeOk() (*string, bool)`

GetMobileThemeOk returns a tuple with the MobileTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileTheme

`func (o *UpdateCurrentPreferencesRequest) SetMobileTheme(v string)`

SetMobileTheme sets MobileTheme field to given value.

### HasMobileTheme

`func (o *UpdateCurrentPreferencesRequest) HasMobileTheme() bool`

HasMobileTheme returns a boolean if a field has been set.

### GetLanguage

`func (o *UpdateCurrentPreferencesRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *UpdateCurrentPreferencesRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *UpdateCurrentPreferencesRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *UpdateCurrentPreferencesRequest) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


