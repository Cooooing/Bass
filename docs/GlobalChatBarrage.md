# GlobalChatBarrage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**PointsCost** | Pointer to **string** |  | [optional] 
**Sender** | Pointer to [**GlobalChatMember**](GlobalChatMember.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGlobalChatBarrage

`func NewGlobalChatBarrage() *GlobalChatBarrage`

NewGlobalChatBarrage instantiates a new GlobalChatBarrage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalChatBarrageWithDefaults

`func NewGlobalChatBarrageWithDefaults() *GlobalChatBarrage`

NewGlobalChatBarrageWithDefaults instantiates a new GlobalChatBarrage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GlobalChatBarrage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GlobalChatBarrage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GlobalChatBarrage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GlobalChatBarrage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *GlobalChatBarrage) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GlobalChatBarrage) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GlobalChatBarrage) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *GlobalChatBarrage) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetContent

`func (o *GlobalChatBarrage) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GlobalChatBarrage) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GlobalChatBarrage) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *GlobalChatBarrage) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetPointsCost

`func (o *GlobalChatBarrage) GetPointsCost() string`

GetPointsCost returns the PointsCost field if non-nil, zero value otherwise.

### GetPointsCostOk

`func (o *GlobalChatBarrage) GetPointsCostOk() (*string, bool)`

GetPointsCostOk returns a tuple with the PointsCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointsCost

`func (o *GlobalChatBarrage) SetPointsCost(v string)`

SetPointsCost sets PointsCost field to given value.

### HasPointsCost

`func (o *GlobalChatBarrage) HasPointsCost() bool`

HasPointsCost returns a boolean if a field has been set.

### GetSender

`func (o *GlobalChatBarrage) GetSender() GlobalChatMember`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *GlobalChatBarrage) GetSenderOk() (*GlobalChatMember, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *GlobalChatBarrage) SetSender(v GlobalChatMember)`

SetSender sets Sender field to given value.

### HasSender

`func (o *GlobalChatBarrage) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GlobalChatBarrage) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GlobalChatBarrage) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GlobalChatBarrage) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GlobalChatBarrage) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


