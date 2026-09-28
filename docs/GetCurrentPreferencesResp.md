# GetCurrentPreferencesResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Preference** | Pointer to [**RespPreference**](RespPreference.md) |  | [optional] 

## Methods

### NewGetCurrentPreferencesResp

`func NewGetCurrentPreferencesResp() *GetCurrentPreferencesResp`

NewGetCurrentPreferencesResp instantiates a new GetCurrentPreferencesResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCurrentPreferencesRespWithDefaults

`func NewGetCurrentPreferencesRespWithDefaults() *GetCurrentPreferencesResp`

NewGetCurrentPreferencesRespWithDefaults instantiates a new GetCurrentPreferencesResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreference

`func (o *GetCurrentPreferencesResp) GetPreference() RespPreference`

GetPreference returns the Preference field if non-nil, zero value otherwise.

### GetPreferenceOk

`func (o *GetCurrentPreferencesResp) GetPreferenceOk() (*RespPreference, bool)`

GetPreferenceOk returns a tuple with the Preference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreference

`func (o *GetCurrentPreferencesResp) SetPreference(v RespPreference)`

SetPreference sets Preference field to given value.

### HasPreference

`func (o *GetCurrentPreferencesResp) HasPreference() bool`

HasPreference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


