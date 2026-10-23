# SearchMentionAccountsResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rows** | Pointer to [**[]AccountProfile**](AccountProfile.md) |  | [optional] 
**Page** | Pointer to [**PageResp**](PageResp.md) |  | [optional] 

## Methods

### NewSearchMentionAccountsResp

`func NewSearchMentionAccountsResp() *SearchMentionAccountsResp`

NewSearchMentionAccountsResp instantiates a new SearchMentionAccountsResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchMentionAccountsRespWithDefaults

`func NewSearchMentionAccountsRespWithDefaults() *SearchMentionAccountsResp`

NewSearchMentionAccountsRespWithDefaults instantiates a new SearchMentionAccountsResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *SearchMentionAccountsResp) GetRows() []AccountProfile`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *SearchMentionAccountsResp) GetRowsOk() (*[]AccountProfile, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *SearchMentionAccountsResp) SetRows(v []AccountProfile)`

SetRows sets Rows field to given value.

### HasRows

`func (o *SearchMentionAccountsResp) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetPage

`func (o *SearchMentionAccountsResp) GetPage() PageResp`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *SearchMentionAccountsResp) GetPageOk() (*PageResp, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *SearchMentionAccountsResp) SetPage(v PageResp)`

SetPage sets Page field to given value.

### HasPage

`func (o *SearchMentionAccountsResp) HasPage() bool`

HasPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


