# ListDomainsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**PageRequest**](PageRequest.md) | 分页参数。 | [optional] 
**Query** | Pointer to [**DomainQuery**](DomainQuery.md) | 查询条件。 | [optional] 

## Methods

### NewListDomainsRequest

`func NewListDomainsRequest() *ListDomainsRequest`

NewListDomainsRequest instantiates a new ListDomainsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDomainsRequestWithDefaults

`func NewListDomainsRequestWithDefaults() *ListDomainsRequest`

NewListDomainsRequestWithDefaults instantiates a new ListDomainsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ListDomainsRequest) GetPage() PageRequest`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListDomainsRequest) GetPageOk() (*PageRequest, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListDomainsRequest) SetPage(v PageRequest)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListDomainsRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetQuery

`func (o *ListDomainsRequest) GetQuery() DomainQuery`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ListDomainsRequest) GetQueryOk() (*DomainQuery, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ListDomainsRequest) SetQuery(v DomainQuery)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *ListDomainsRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


