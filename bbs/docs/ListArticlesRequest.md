# ListArticlesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**PageRequest**](PageRequest.md) |  | [optional] 
**Query** | Pointer to [**ArticleQuery**](ArticleQuery.md) |  | [optional] 

## Methods

### NewListArticlesRequest

`func NewListArticlesRequest() *ListArticlesRequest`

NewListArticlesRequest instantiates a new ListArticlesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListArticlesRequestWithDefaults

`func NewListArticlesRequestWithDefaults() *ListArticlesRequest`

NewListArticlesRequestWithDefaults instantiates a new ListArticlesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ListArticlesRequest) GetPage() PageRequest`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListArticlesRequest) GetPageOk() (*PageRequest, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListArticlesRequest) SetPage(v PageRequest)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListArticlesRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetQuery

`func (o *ListArticlesRequest) GetQuery() ArticleQuery`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ListArticlesRequest) GetQueryOk() (*ArticleQuery, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ListArticlesRequest) SetQuery(v ArticleQuery)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *ListArticlesRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


