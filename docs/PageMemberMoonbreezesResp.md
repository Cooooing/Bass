# PageMemberMoonbreezesResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rows** | Pointer to [**[]Moonbreeze**](Moonbreeze.md) | 协议字段。 | [optional] 
**NextCursor** | Pointer to **string** | 协议字段。 | [optional] 

## Methods

### NewPageMemberMoonbreezesResp

`func NewPageMemberMoonbreezesResp() *PageMemberMoonbreezesResp`

NewPageMemberMoonbreezesResp instantiates a new PageMemberMoonbreezesResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageMemberMoonbreezesRespWithDefaults

`func NewPageMemberMoonbreezesRespWithDefaults() *PageMemberMoonbreezesResp`

NewPageMemberMoonbreezesRespWithDefaults instantiates a new PageMemberMoonbreezesResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *PageMemberMoonbreezesResp) GetRows() []Moonbreeze`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *PageMemberMoonbreezesResp) GetRowsOk() (*[]Moonbreeze, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *PageMemberMoonbreezesResp) SetRows(v []Moonbreeze)`

SetRows sets Rows field to given value.

### HasRows

`func (o *PageMemberMoonbreezesResp) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetNextCursor

`func (o *PageMemberMoonbreezesResp) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *PageMemberMoonbreezesResp) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *PageMemberMoonbreezesResp) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *PageMemberMoonbreezesResp) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


