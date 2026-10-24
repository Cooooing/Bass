# ListGlobalChatBarragesResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**PageResp**](PageResp.md) |  | [optional] 
**Rows** | Pointer to [**[]GlobalChatBarrage**](GlobalChatBarrage.md) |  | [optional] 

## Methods

### NewListGlobalChatBarragesResp

`func NewListGlobalChatBarragesResp() *ListGlobalChatBarragesResp`

NewListGlobalChatBarragesResp instantiates a new ListGlobalChatBarragesResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListGlobalChatBarragesRespWithDefaults

`func NewListGlobalChatBarragesRespWithDefaults() *ListGlobalChatBarragesResp`

NewListGlobalChatBarragesRespWithDefaults instantiates a new ListGlobalChatBarragesResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ListGlobalChatBarragesResp) GetPage() PageResp`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListGlobalChatBarragesResp) GetPageOk() (*PageResp, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListGlobalChatBarragesResp) SetPage(v PageResp)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListGlobalChatBarragesResp) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetRows

`func (o *ListGlobalChatBarragesResp) GetRows() []GlobalChatBarrage`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *ListGlobalChatBarragesResp) GetRowsOk() (*[]GlobalChatBarrage, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *ListGlobalChatBarragesResp) SetRows(v []GlobalChatBarrage)`

SetRows sets Rows field to given value.

### HasRows

`func (o *ListGlobalChatBarragesResp) HasRows() bool`

HasRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


