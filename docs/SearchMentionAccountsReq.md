# SearchMentionAccountsReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to **string** |  | [optional] 
**Page** | Pointer to [**PageReq**](PageReq.md) |  | [optional] 

## Methods

### NewSearchMentionAccountsReq

`func NewSearchMentionAccountsReq() *SearchMentionAccountsReq`

NewSearchMentionAccountsReq instantiates a new SearchMentionAccountsReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchMentionAccountsReqWithDefaults

`func NewSearchMentionAccountsReqWithDefaults() *SearchMentionAccountsReq`

NewSearchMentionAccountsReqWithDefaults instantiates a new SearchMentionAccountsReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *SearchMentionAccountsReq) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SearchMentionAccountsReq) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SearchMentionAccountsReq) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *SearchMentionAccountsReq) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetPage

`func (o *SearchMentionAccountsReq) GetPage() PageReq`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *SearchMentionAccountsReq) GetPageOk() (*PageReq, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *SearchMentionAccountsReq) SetPage(v PageReq)`

SetPage sets Page field to given value.

### HasPage

`func (o *SearchMentionAccountsReq) HasPage() bool`

HasPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


