# ListCommentsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**PageRequest**](PageRequest.md) | 分页参数。 | [optional] 
**Query** | Pointer to [**CommentQuery**](CommentQuery.md) | 查询条件。 | [optional] 

## Methods

### NewListCommentsRequest

`func NewListCommentsRequest() *ListCommentsRequest`

NewListCommentsRequest instantiates a new ListCommentsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCommentsRequestWithDefaults

`func NewListCommentsRequestWithDefaults() *ListCommentsRequest`

NewListCommentsRequestWithDefaults instantiates a new ListCommentsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ListCommentsRequest) GetPage() PageRequest`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListCommentsRequest) GetPageOk() (*PageRequest, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListCommentsRequest) SetPage(v PageRequest)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListCommentsRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetQuery

`func (o *ListCommentsRequest) GetQuery() CommentQuery`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ListCommentsRequest) GetQueryOk() (*CommentQuery, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ListCommentsRequest) SetQuery(v CommentQuery)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *ListCommentsRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


