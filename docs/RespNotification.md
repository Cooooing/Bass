# RespNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | 通知 ID | [optional] 
**EventId** | Pointer to **string** | 来源事件 ID | [optional] 
**ReceiverId** | Pointer to **string** | 接收账号 ID | [optional] 
**NotificationType** | Pointer to **string** | 通知类型 | [optional] 
**Title** | Pointer to **string** | 通知标题 | [optional] 
**Content** | Pointer to **string** | 通知内容 | [optional] 
**ReadAt** | Pointer to **time.Time** | 已读时间 | [optional] 
**CreatedAt** | Pointer to **time.Time** | 创建时间 | [optional] 
**UpdatedAt** | Pointer to **time.Time** | 更新时间 | [optional] 

## Methods

### NewRespNotification

`func NewRespNotification() *RespNotification`

NewRespNotification instantiates a new RespNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRespNotificationWithDefaults

`func NewRespNotificationWithDefaults() *RespNotification`

NewRespNotificationWithDefaults instantiates a new RespNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RespNotification) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RespNotification) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RespNotification) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RespNotification) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEventId

`func (o *RespNotification) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *RespNotification) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *RespNotification) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *RespNotification) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetReceiverId

`func (o *RespNotification) GetReceiverId() string`

GetReceiverId returns the ReceiverId field if non-nil, zero value otherwise.

### GetReceiverIdOk

`func (o *RespNotification) GetReceiverIdOk() (*string, bool)`

GetReceiverIdOk returns a tuple with the ReceiverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverId

`func (o *RespNotification) SetReceiverId(v string)`

SetReceiverId sets ReceiverId field to given value.

### HasReceiverId

`func (o *RespNotification) HasReceiverId() bool`

HasReceiverId returns a boolean if a field has been set.

### GetNotificationType

`func (o *RespNotification) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *RespNotification) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *RespNotification) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.

### HasNotificationType

`func (o *RespNotification) HasNotificationType() bool`

HasNotificationType returns a boolean if a field has been set.

### GetTitle

`func (o *RespNotification) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RespNotification) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RespNotification) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *RespNotification) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetContent

`func (o *RespNotification) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *RespNotification) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *RespNotification) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *RespNotification) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetReadAt

`func (o *RespNotification) GetReadAt() time.Time`

GetReadAt returns the ReadAt field if non-nil, zero value otherwise.

### GetReadAtOk

`func (o *RespNotification) GetReadAtOk() (*time.Time, bool)`

GetReadAtOk returns a tuple with the ReadAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadAt

`func (o *RespNotification) SetReadAt(v time.Time)`

SetReadAt sets ReadAt field to given value.

### HasReadAt

`func (o *RespNotification) HasReadAt() bool`

HasReadAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RespNotification) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RespNotification) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RespNotification) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RespNotification) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RespNotification) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RespNotification) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RespNotification) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RespNotification) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


