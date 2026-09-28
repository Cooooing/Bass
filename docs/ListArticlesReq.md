# ListArticlesReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**PageReq**](PageReq.md) |  | [optional] 
**Query** | Pointer to [**ReqArticleQuery**](ReqArticleQuery.md) |  | [optional] 

## Methods

### NewListArticlesReq

`func NewListArticlesReq() *ListArticlesReq`

NewListArticlesReq instantiates a new ListArticlesReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListArticlesReqWithDefaults

`func NewListArticlesReqWithDefaults() *ListArticlesReq`

NewListArticlesReqWithDefaults instantiates a new ListArticlesReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ListArticlesReq) GetPage() PageReq`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListArticlesReq) GetPageOk() (*PageReq, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListArticlesReq) SetPage(v PageReq)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListArticlesReq) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetQuery

`func (o *ListArticlesReq) GetQuery() ReqArticleQuery`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ListArticlesReq) GetQueryOk() (*ReqArticleQuery, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ListArticlesReq) SetQuery(v ReqArticleQuery)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *ListArticlesReq) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


