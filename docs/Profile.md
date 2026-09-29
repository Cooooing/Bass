# Profile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**AccountProfile**](AccountProfile.md) |  | [optional] 
**BackgroundUrl** | Pointer to **string** |  | [optional] 
**Location** | Pointer to [**ProfileLocation**](ProfileLocation.md) |  | [optional] 
**LastSuccessLoginAt** | Pointer to **time.Time** |  | [optional] 
**Visibility** | Pointer to [**ProfileVisibility**](ProfileVisibility.md) |  | [optional] 
**ViewerRelation** | Pointer to [**ProfileRelation**](ProfileRelation.md) |  | [optional] 

## Methods

### NewProfile

`func NewProfile() *Profile`

NewProfile instantiates a new Profile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileWithDefaults

`func NewProfileWithDefaults() *Profile`

NewProfileWithDefaults instantiates a new Profile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *Profile) GetAccount() AccountProfile`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Profile) GetAccountOk() (*AccountProfile, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Profile) SetAccount(v AccountProfile)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Profile) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetBackgroundUrl

`func (o *Profile) GetBackgroundUrl() string`

GetBackgroundUrl returns the BackgroundUrl field if non-nil, zero value otherwise.

### GetBackgroundUrlOk

`func (o *Profile) GetBackgroundUrlOk() (*string, bool)`

GetBackgroundUrlOk returns a tuple with the BackgroundUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundUrl

`func (o *Profile) SetBackgroundUrl(v string)`

SetBackgroundUrl sets BackgroundUrl field to given value.

### HasBackgroundUrl

`func (o *Profile) HasBackgroundUrl() bool`

HasBackgroundUrl returns a boolean if a field has been set.

### GetLocation

`func (o *Profile) GetLocation() ProfileLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Profile) GetLocationOk() (*ProfileLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Profile) SetLocation(v ProfileLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Profile) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetLastSuccessLoginAt

`func (o *Profile) GetLastSuccessLoginAt() time.Time`

GetLastSuccessLoginAt returns the LastSuccessLoginAt field if non-nil, zero value otherwise.

### GetLastSuccessLoginAtOk

`func (o *Profile) GetLastSuccessLoginAtOk() (*time.Time, bool)`

GetLastSuccessLoginAtOk returns a tuple with the LastSuccessLoginAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessLoginAt

`func (o *Profile) SetLastSuccessLoginAt(v time.Time)`

SetLastSuccessLoginAt sets LastSuccessLoginAt field to given value.

### HasLastSuccessLoginAt

`func (o *Profile) HasLastSuccessLoginAt() bool`

HasLastSuccessLoginAt returns a boolean if a field has been set.

### GetVisibility

`func (o *Profile) GetVisibility() ProfileVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Profile) GetVisibilityOk() (*ProfileVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Profile) SetVisibility(v ProfileVisibility)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *Profile) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetViewerRelation

`func (o *Profile) GetViewerRelation() ProfileRelation`

GetViewerRelation returns the ViewerRelation field if non-nil, zero value otherwise.

### GetViewerRelationOk

`func (o *Profile) GetViewerRelationOk() (*ProfileRelation, bool)`

GetViewerRelationOk returns a tuple with the ViewerRelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerRelation

`func (o *Profile) SetViewerRelation(v ProfileRelation)`

SetViewerRelation sets ViewerRelation field to given value.

### HasViewerRelation

`func (o *Profile) HasViewerRelation() bool`

HasViewerRelation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


