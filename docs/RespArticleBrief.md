# RespArticleBrief

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**AuthorUser** | Pointer to [**RespAccountProfile**](RespAccountProfile.md) |  | [optional] 
**CoverImageUrl** | Pointer to **string** |  | [optional] 
**PublishStatus** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Restriction** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRespArticleBrief

`func NewRespArticleBrief() *RespArticleBrief`

NewRespArticleBrief instantiates a new RespArticleBrief object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRespArticleBriefWithDefaults

`func NewRespArticleBriefWithDefaults() *RespArticleBrief`

NewRespArticleBriefWithDefaults instantiates a new RespArticleBrief object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RespArticleBrief) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RespArticleBrief) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RespArticleBrief) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RespArticleBrief) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *RespArticleBrief) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RespArticleBrief) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RespArticleBrief) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *RespArticleBrief) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *RespArticleBrief) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RespArticleBrief) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RespArticleBrief) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RespArticleBrief) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAuthorUser

`func (o *RespArticleBrief) GetAuthorUser() RespAccountProfile`

GetAuthorUser returns the AuthorUser field if non-nil, zero value otherwise.

### GetAuthorUserOk

`func (o *RespArticleBrief) GetAuthorUserOk() (*RespAccountProfile, bool)`

GetAuthorUserOk returns a tuple with the AuthorUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorUser

`func (o *RespArticleBrief) SetAuthorUser(v RespAccountProfile)`

SetAuthorUser sets AuthorUser field to given value.

### HasAuthorUser

`func (o *RespArticleBrief) HasAuthorUser() bool`

HasAuthorUser returns a boolean if a field has been set.

### GetCoverImageUrl

`func (o *RespArticleBrief) GetCoverImageUrl() string`

GetCoverImageUrl returns the CoverImageUrl field if non-nil, zero value otherwise.

### GetCoverImageUrlOk

`func (o *RespArticleBrief) GetCoverImageUrlOk() (*string, bool)`

GetCoverImageUrlOk returns a tuple with the CoverImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverImageUrl

`func (o *RespArticleBrief) SetCoverImageUrl(v string)`

SetCoverImageUrl sets CoverImageUrl field to given value.

### HasCoverImageUrl

`func (o *RespArticleBrief) HasCoverImageUrl() bool`

HasCoverImageUrl returns a boolean if a field has been set.

### GetPublishStatus

`func (o *RespArticleBrief) GetPublishStatus() string`

GetPublishStatus returns the PublishStatus field if non-nil, zero value otherwise.

### GetPublishStatusOk

`func (o *RespArticleBrief) GetPublishStatusOk() (*string, bool)`

GetPublishStatusOk returns a tuple with the PublishStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishStatus

`func (o *RespArticleBrief) SetPublishStatus(v string)`

SetPublishStatus sets PublishStatus field to given value.

### HasPublishStatus

`func (o *RespArticleBrief) HasPublishStatus() bool`

HasPublishStatus returns a boolean if a field has been set.

### GetVisibility

`func (o *RespArticleBrief) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *RespArticleBrief) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *RespArticleBrief) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *RespArticleBrief) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetRestriction

`func (o *RespArticleBrief) GetRestriction() string`

GetRestriction returns the Restriction field if non-nil, zero value otherwise.

### GetRestrictionOk

`func (o *RespArticleBrief) GetRestrictionOk() (*string, bool)`

GetRestrictionOk returns a tuple with the Restriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestriction

`func (o *RespArticleBrief) SetRestriction(v string)`

SetRestriction sets Restriction field to given value.

### HasRestriction

`func (o *RespArticleBrief) HasRestriction() bool`

HasRestriction returns a boolean if a field has been set.

### GetCreatedBy

`func (o *RespArticleBrief) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RespArticleBrief) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RespArticleBrief) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *RespArticleBrief) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *RespArticleBrief) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *RespArticleBrief) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *RespArticleBrief) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *RespArticleBrief) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RespArticleBrief) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RespArticleBrief) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RespArticleBrief) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RespArticleBrief) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RespArticleBrief) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RespArticleBrief) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RespArticleBrief) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RespArticleBrief) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


