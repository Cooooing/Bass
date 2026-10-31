# ListPointRecordsResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rows** | Pointer to [**[]RespPointRecord**](RespPointRecord.md) |  | [optional] 
**Page** | Pointer to [**PageResp**](PageResp.md) |  | [optional] 

## Methods

### NewListPointRecordsResp

`func NewListPointRecordsResp() *ListPointRecordsResp`

NewListPointRecordsResp instantiates a new ListPointRecordsResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPointRecordsRespWithDefaults

`func NewListPointRecordsRespWithDefaults() *ListPointRecordsResp`

NewListPointRecordsRespWithDefaults instantiates a new ListPointRecordsResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *ListPointRecordsResp) GetRows() []RespPointRecord`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *ListPointRecordsResp) GetRowsOk() (*[]RespPointRecord, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *ListPointRecordsResp) SetRows(v []RespPointRecord)`

SetRows sets Rows field to given value.

### HasRows

`func (o *ListPointRecordsResp) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetPage

`func (o *ListPointRecordsResp) GetPage() PageResp`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListPointRecordsResp) GetPageOk() (*PageResp, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListPointRecordsResp) SetPage(v PageResp)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListPointRecordsResp) HasPage() bool`

HasPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


