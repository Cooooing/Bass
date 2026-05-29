# ListNotificationsReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**PageReply**](PageReply.md) |  | [optional] 
**Rows** | Pointer to [**[]Notification**](Notification.md) |  | [optional] 

## Methods

### NewListNotificationsReply

`func NewListNotificationsReply() *ListNotificationsReply`

NewListNotificationsReply instantiates a new ListNotificationsReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNotificationsReplyWithDefaults

`func NewListNotificationsReplyWithDefaults() *ListNotificationsReply`

NewListNotificationsReplyWithDefaults instantiates a new ListNotificationsReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ListNotificationsReply) GetPage() PageReply`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListNotificationsReply) GetPageOk() (*PageReply, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListNotificationsReply) SetPage(v PageReply)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListNotificationsReply) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetRows

`func (o *ListNotificationsReply) GetRows() []Notification`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *ListNotificationsReply) GetRowsOk() (*[]Notification, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *ListNotificationsReply) SetRows(v []Notification)`

SetRows sets Rows field to given value.

### HasRows

`func (o *ListNotificationsReply) HasRows() bool`

HasRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


