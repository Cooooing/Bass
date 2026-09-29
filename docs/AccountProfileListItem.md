# AccountProfileListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**AccountProfile**](AccountProfile.md) |  | [optional] 
**ViewerRelation** | Pointer to [**ProfileRelation**](ProfileRelation.md) |  | [optional] 

## Methods

### NewAccountProfileListItem

`func NewAccountProfileListItem() *AccountProfileListItem`

NewAccountProfileListItem instantiates a new AccountProfileListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountProfileListItemWithDefaults

`func NewAccountProfileListItemWithDefaults() *AccountProfileListItem`

NewAccountProfileListItemWithDefaults instantiates a new AccountProfileListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *AccountProfileListItem) GetAccount() AccountProfile`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AccountProfileListItem) GetAccountOk() (*AccountProfile, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AccountProfileListItem) SetAccount(v AccountProfile)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AccountProfileListItem) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetViewerRelation

`func (o *AccountProfileListItem) GetViewerRelation() ProfileRelation`

GetViewerRelation returns the ViewerRelation field if non-nil, zero value otherwise.

### GetViewerRelationOk

`func (o *AccountProfileListItem) GetViewerRelationOk() (*ProfileRelation, bool)`

GetViewerRelationOk returns a tuple with the ViewerRelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerRelation

`func (o *AccountProfileListItem) SetViewerRelation(v ProfileRelation)`

SetViewerRelation sets ViewerRelation field to given value.

### HasViewerRelation

`func (o *AccountProfileListItem) HasViewerRelation() bool`

HasViewerRelation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


