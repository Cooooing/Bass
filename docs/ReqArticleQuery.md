# ReqArticleQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TagId** | Pointer to **string** |  | [optional] 
**DomainId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Order** | Pointer to **string** |  | [optional] 
**Keyword** | Pointer to **string** |  | [optional] 
**AuthorId** | Pointer to **string** |  | [optional] 
**PublishStatus** | Pointer to **string** |  | [optional] 
**PublishStatuses** | Pointer to **[]string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Visibilities** | Pointer to **[]string** |  | [optional] 
**Restriction** | Pointer to **string** |  | [optional] 
**Restrictions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewReqArticleQuery

`func NewReqArticleQuery() *ReqArticleQuery`

NewReqArticleQuery instantiates a new ReqArticleQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReqArticleQueryWithDefaults

`func NewReqArticleQueryWithDefaults() *ReqArticleQuery`

NewReqArticleQueryWithDefaults instantiates a new ReqArticleQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagId

`func (o *ReqArticleQuery) GetTagId() string`

GetTagId returns the TagId field if non-nil, zero value otherwise.

### GetTagIdOk

`func (o *ReqArticleQuery) GetTagIdOk() (*string, bool)`

GetTagIdOk returns a tuple with the TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagId

`func (o *ReqArticleQuery) SetTagId(v string)`

SetTagId sets TagId field to given value.

### HasTagId

`func (o *ReqArticleQuery) HasTagId() bool`

HasTagId returns a boolean if a field has been set.

### GetDomainId

`func (o *ReqArticleQuery) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *ReqArticleQuery) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *ReqArticleQuery) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *ReqArticleQuery) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetType

`func (o *ReqArticleQuery) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReqArticleQuery) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReqArticleQuery) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ReqArticleQuery) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOrder

`func (o *ReqArticleQuery) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ReqArticleQuery) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ReqArticleQuery) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ReqArticleQuery) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetKeyword

`func (o *ReqArticleQuery) GetKeyword() string`

GetKeyword returns the Keyword field if non-nil, zero value otherwise.

### GetKeywordOk

`func (o *ReqArticleQuery) GetKeywordOk() (*string, bool)`

GetKeywordOk returns a tuple with the Keyword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyword

`func (o *ReqArticleQuery) SetKeyword(v string)`

SetKeyword sets Keyword field to given value.

### HasKeyword

`func (o *ReqArticleQuery) HasKeyword() bool`

HasKeyword returns a boolean if a field has been set.

### GetAuthorId

`func (o *ReqArticleQuery) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *ReqArticleQuery) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *ReqArticleQuery) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *ReqArticleQuery) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetPublishStatus

`func (o *ReqArticleQuery) GetPublishStatus() string`

GetPublishStatus returns the PublishStatus field if non-nil, zero value otherwise.

### GetPublishStatusOk

`func (o *ReqArticleQuery) GetPublishStatusOk() (*string, bool)`

GetPublishStatusOk returns a tuple with the PublishStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishStatus

`func (o *ReqArticleQuery) SetPublishStatus(v string)`

SetPublishStatus sets PublishStatus field to given value.

### HasPublishStatus

`func (o *ReqArticleQuery) HasPublishStatus() bool`

HasPublishStatus returns a boolean if a field has been set.

### GetPublishStatuses

`func (o *ReqArticleQuery) GetPublishStatuses() []string`

GetPublishStatuses returns the PublishStatuses field if non-nil, zero value otherwise.

### GetPublishStatusesOk

`func (o *ReqArticleQuery) GetPublishStatusesOk() (*[]string, bool)`

GetPublishStatusesOk returns a tuple with the PublishStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishStatuses

`func (o *ReqArticleQuery) SetPublishStatuses(v []string)`

SetPublishStatuses sets PublishStatuses field to given value.

### HasPublishStatuses

`func (o *ReqArticleQuery) HasPublishStatuses() bool`

HasPublishStatuses returns a boolean if a field has been set.

### GetVisibility

`func (o *ReqArticleQuery) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ReqArticleQuery) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ReqArticleQuery) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ReqArticleQuery) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetVisibilities

`func (o *ReqArticleQuery) GetVisibilities() []string`

GetVisibilities returns the Visibilities field if non-nil, zero value otherwise.

### GetVisibilitiesOk

`func (o *ReqArticleQuery) GetVisibilitiesOk() (*[]string, bool)`

GetVisibilitiesOk returns a tuple with the Visibilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibilities

`func (o *ReqArticleQuery) SetVisibilities(v []string)`

SetVisibilities sets Visibilities field to given value.

### HasVisibilities

`func (o *ReqArticleQuery) HasVisibilities() bool`

HasVisibilities returns a boolean if a field has been set.

### GetRestriction

`func (o *ReqArticleQuery) GetRestriction() string`

GetRestriction returns the Restriction field if non-nil, zero value otherwise.

### GetRestrictionOk

`func (o *ReqArticleQuery) GetRestrictionOk() (*string, bool)`

GetRestrictionOk returns a tuple with the Restriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestriction

`func (o *ReqArticleQuery) SetRestriction(v string)`

SetRestriction sets Restriction field to given value.

### HasRestriction

`func (o *ReqArticleQuery) HasRestriction() bool`

HasRestriction returns a boolean if a field has been set.

### GetRestrictions

`func (o *ReqArticleQuery) GetRestrictions() []string`

GetRestrictions returns the Restrictions field if non-nil, zero value otherwise.

### GetRestrictionsOk

`func (o *ReqArticleQuery) GetRestrictionsOk() (*[]string, bool)`

GetRestrictionsOk returns a tuple with the Restrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictions

`func (o *ReqArticleQuery) SetRestrictions(v []string)`

SetRestrictions sets Restrictions field to given value.

### HasRestrictions

`func (o *ReqArticleQuery) HasRestrictions() bool`

HasRestrictions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


