# ListDomainsReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**PageReq**](PageReq.md) |  | [optional] 
**Query** | Pointer to [**ReqDomainQuery**](ReqDomainQuery.md) |  | [optional] 

## Methods

### NewListDomainsReq

`func NewListDomainsReq() *ListDomainsReq`

NewListDomainsReq instantiates a new ListDomainsReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDomainsReqWithDefaults

`func NewListDomainsReqWithDefaults() *ListDomainsReq`

NewListDomainsReqWithDefaults instantiates a new ListDomainsReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ListDomainsReq) GetPage() PageReq`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListDomainsReq) GetPageOk() (*PageReq, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListDomainsReq) SetPage(v PageReq)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListDomainsReq) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetQuery

`func (o *ListDomainsReq) GetQuery() ReqDomainQuery`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ListDomainsReq) GetQueryOk() (*ReqDomainQuery, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ListDomainsReq) SetQuery(v ReqDomainQuery)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *ListDomainsReq) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


