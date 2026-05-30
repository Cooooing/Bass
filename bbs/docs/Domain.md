# Domain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | 板块 ID。 | [optional] 
**Name** | Pointer to **string** | 板块名称。 | [optional] 
**Description** | Pointer to **string** | 板块描述。 | [optional] 
**Status** | Pointer to **string** | 板块状态。 | [optional] 
**Url** | Pointer to **string** | 板块 URL。 | [optional] 
**Icon** | Pointer to **string** | 板块图标。 | [optional] 
**IsNav** | Pointer to **bool** | 是否在导航中展示。 | [optional] 
**CreatedBy** | Pointer to **string** | 创建账号 ID。 | [optional] 
**UpdatedBy** | Pointer to **string** | 更新账号 ID。 | [optional] 
**CreatedAt** | Pointer to **string** | 创建时间。 | [optional] 
**UpdatedAt** | Pointer to **string** | 更新时间。 | [optional] 

## Methods

### NewDomain

`func NewDomain() *Domain`

NewDomain instantiates a new Domain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainWithDefaults

`func NewDomainWithDefaults() *Domain`

NewDomainWithDefaults instantiates a new Domain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Domain) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Domain) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Domain) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Domain) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Domain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Domain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Domain) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Domain) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Domain) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Domain) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Domain) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Domain) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *Domain) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Domain) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Domain) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Domain) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUrl

`func (o *Domain) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Domain) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Domain) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Domain) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetIcon

`func (o *Domain) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *Domain) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *Domain) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *Domain) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetIsNav

`func (o *Domain) GetIsNav() bool`

GetIsNav returns the IsNav field if non-nil, zero value otherwise.

### GetIsNavOk

`func (o *Domain) GetIsNavOk() (*bool, bool)`

GetIsNavOk returns a tuple with the IsNav field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNav

`func (o *Domain) SetIsNav(v bool)`

SetIsNav sets IsNav field to given value.

### HasIsNav

`func (o *Domain) HasIsNav() bool`

HasIsNav returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Domain) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Domain) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Domain) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Domain) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Domain) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Domain) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Domain) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Domain) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Domain) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Domain) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Domain) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Domain) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Domain) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Domain) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Domain) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Domain) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


